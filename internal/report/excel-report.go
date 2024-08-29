package report

import (
	"bytes"
	"fmt"
	"log"
	"sync"

	"github.com/henrybravo/micro-report/internal/repositories"
	"github.com/xuri/excelize/v2"
)

type ExcelGenerator struct{}

func NewExcelGenerator() *ExcelGenerator {
	return &ExcelGenerator{}
}

type HeaderTitleCol struct {
	topLeft     string
	bottomRight string
	title       string
}

func (e *ExcelGenerator) GenerateSalesReport(business repositories.Business, sales []repositories.SalesReport, period string) (*bytes.Buffer, error) {

	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	sheetName := "Reporte_de_ventas"
	index, err := f.NewSheet(sheetName)
	if err != nil {
		return nil, err
	}
	f.SetActiveSheet(index)

	if err := setSheetStyles(f, sheetName); err != nil {
		return nil, err
	}

	if err := createTitle(f, business, period, sheetName); err != nil {
		return nil, err
	}

	if err := createHeaders(f, sheetName); err != nil {
		return nil, err
	}

	lastRow, err := fillSalesData(f, sheetName, sales)
	if err != nil {
		return nil, err
	}

	if err := addTotalRow(f, sheetName, lastRow); err != nil {
		return nil, err
	}

	buff, err := f.WriteToBuffer()
	if err != nil {
		return nil, err
	}
	return buff, nil
}

func setSheetStyles(f *excelize.File, sheetName string) error {
	//quitar cuadrícula a la hoja
	showGridLines := false
	if err := f.SetSheetView(sheetName, 0, &excelize.ViewOptions{
		ShowGridLines: &showGridLines,
	}); err != nil {
		return err
	}

	columns := []struct {
		col   string
		width float64
	}{
		{"A", 14}, {"B", 14}, {"C", 14}, {"D", 5}, {"E", 6},
		{"F", 7}, {"G", 5}, {"H", 13}, {"I", 43}, {"J", 10},
		{"K", 12}, {"L", 9}, {"M", 13}, {"N", 11}, {"O", 5},
		{"P", 10}, {"Q", 6}, {"R", 7}, {"S", 8}, {"T", 11},
		{"U", 6}, {"V", 10}, {"W", 6}, {"X", 7}, {"Y", 10},
	}

	for _, col := range columns {
		if err := f.SetColWidth(sheetName, col.col, col.col, col.width); err != nil {
			return err
		}
	}

	return nil
}

func createTitle(f *excelize.File, business repositories.Business, period string, sheetName string) error {
	titleStyle, err := createTitleStyle(f)
	if err != nil {
		return err
	}

	businessName := fmt.Sprintf("%s\nR.U.C.:%s\n%s\nREGISTRO DE VENTAS DEL MES DE %s", business.BusinessName, business.RUC, business.Address, period)
	if err := f.MergeCell(sheetName, "A1", "Y1"); err != nil {
		return err
	}
	if err := f.SetCellStyle(sheetName, "A1", "Y1", titleStyle); err != nil {
		return err
	}
	if err := f.SetCellStr(sheetName, "A1", businessName); err != nil {
		return err
	}
	if err := f.SetRowHeight(sheetName, 1, 80); err != nil {
		return err
	}

	return nil
}

func createTitleStyle(f *excelize.File) (int, error) {
	return f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold:   true,
			Size:   11,
			Family: "Arial Narrow",
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
			WrapText:   true,
		},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 0},
			{Type: "top", Color: "000000", Style: 0},
			{Type: "bottom", Color: "000000", Style: 0},
			{Type: "right", Color: "000000", Style: 0},
		},
	})
}

func createHeaders(f *excelize.File, sheetName string) error {
	headerStyle, err := createHeaderStyle(f)
	if err != nil {
		return err
	}

	headers := []HeaderTitleCol{
		{"A2", "A6", "CUO"},
		{"B2", "F4", "COMPROBANTE DE PAGO O DOCUMENTO"},
		{"B5", "B6", "FECHA DE EMISIÓN"},
		{"C5", "C6", "FECHA DE VENCIMIENTO"},
		{"D5", "D6", "TIPO"},
		{"E5", "E6", "SERIE"},
		{"F5", "F6", "NUMERO"},
		{"G2", "I2", "INFORMACIÓN DEL CLIENTE"},
		{"G3", "H4", "DOCUMENTO DE IDENTIDAD"},
		{"G5", "G6", "TIPO"},
		{"H5", "H6", "NÚMERO"},
		{"I3", "I6", "APELLIDOS Y NOMBRES O RAZÓN SOCIAL"},
		{"J2", "J6", "VALOR FACTURADO O DE EXPORTACIÓN"},
		{"K2", "K6", "BASE IMPONIBLE DE LA OPERACIÓN GRAVADA"},
		{"L2", "L6", "IGV Y/O IPM"},
		{"M2", "N4", "VALOR TOTAL DE LA OPERACIÓN EXONERADA O INAFECTA"},
		{"M5", "M6", "EXONERADA"},
		{"N5", "N6", "INAFECTA"},
		{"O2", "O6", "ISC"},
		{"P2", "Q4", "OPERACIÓN GRAVADA CON EL IVAP"},
		{"P5", "P6", "BASE IMPONIBLE"},
		{"Q5", "Q6", "IVAP"},
		{"R2", "R6", "ICBPER"},
		{"S2", "S6", "OTROS TRIBUTOS Y CARGOS"},
		{"T2", "T6", "IMPORTE TOTAL"},
		{"U2", "U6", "TIPO DE CAMBIO"},
		{"V2", "Y4", "REFERENCIA DEL COMPROBANTE DE PAGO O DOCUMENTO ORIGINAL QUE SE MODIFICA"},
		{"V5", "V6", "FECHA"},
		{"W5", "W6", "TIPO"},
		{"X5", "X6", "SERIE"},
		{"Y5", "Y6", "NÚMERO"},
	}

	for _, header := range headers {
		if err := createHeaderTitle(f, headerStyle, sheetName, header); err != nil {
			return err
		}
	}

	return nil
}

func createHeaderTitle(f *excelize.File, style int, sheetName string, props HeaderTitleCol) error {
	if err := f.MergeCell(sheetName, props.topLeft, props.bottomRight); err != nil {
		return err
	}
	if err := f.SetCellStyle(sheetName, props.topLeft, props.bottomRight, style); err != nil {
		return err
	}
	return f.SetCellStr(sheetName, props.topLeft, props.title)
}

func createHeaderStyle(f *excelize.File) (int, error) {
	borderStyle := []excelize.Border{
		{Type: "left", Color: "000000", Style: 1},
		{Type: "top", Color: "000000", Style: 1},
		{Type: "bottom", Color: "000000", Style: 1},
		{Type: "right", Color: "000000", Style: 1},
	}
	alignmentStyle := &excelize.Alignment{
		Horizontal: "center",
		Vertical:   "center",
		WrapText:   true,
	}
	fontStyle := &excelize.Font{
		Family: "Arial Narrow",
		Size:   7.5,
		Bold:   false,
	}
	return f.NewStyle(&excelize.Style{
		Alignment: alignmentStyle,
		Border:    borderStyle,
		Font:      fontStyle,
	})
}

func fillSalesData(f *excelize.File, sheetName string, sales []repositories.SalesReport) (int, error) {
	var wg sync.WaitGroup
	rowCh := make(chan struct {
		row  int
		sale repositories.SalesReport
	}, len(sales))

	// Goroutine para llenar datos en paralelo
	for i := 0; i < 4; i++ { // 4 es el número de goroutines, ajusta según sea necesario
		wg.Add(1)
		go func() {
			defer wg.Done()
			for row := range rowCh {
				if err := setSalesRow(f, sheetName, row.row, row.sale); err != nil {
					log.Printf("Error al establecer el valor en la celda: %s", err.Error())
				}
			}
		}()
	}

	// Alimenta el canal con los datos
	for index, sale := range sales {
		rowCh <- struct {
			row  int
			sale repositories.SalesReport
		}{row: index + 7, sale: sale}
	}
	close(rowCh)

	wg.Wait()
	return len(sales) + 7, nil
}

func setSalesRow(f *excelize.File, sheetName string, row int, sale repositories.SalesReport) error {
	exp := "#,##0.00_);(#,##0.00)"
	// Define la configuración de la fuente
	fontConfig := &excelize.Font{
		Family: "Arial Narrow",
		Size:   8,
	}

	numberStyle, _ := f.NewStyle(&excelize.Style{
		CustomNumFmt: &exp,
		Font:         fontConfig,
	})

	// Define el estilo de solo fuente
	fontStyle, _ := f.NewStyle(&excelize.Style{
		Font: fontConfig,
	})

	fechaEmisionStr := sale.FechaEmision.Time.Format("02/01/2006")
	fechaVencimientoStr := ""
	if sale.FechaVencimiento.Valid {
		fechaVencimientoStr = sale.FechaVencimiento.Time.Format("02/01/2006")
	}

	cells := []struct {
		col string
		val interface{}
	}{
		{"A", sale.Cuo}, {"B", fechaEmisionStr}, {"C", fechaVencimientoStr},
		{"D", sale.CodigoTipoCDP}, {"E", sale.Serie}, {"F", sale.Correlativo},
		{"G", sale.CodTipoDocIdentidad}, {"H", sale.NumDocIdentidadClient},
		{"I", sale.RazonSocial}, {"J", sale.MtoValFactExpo}, {"K", sale.BaseIVAP},
		{"L", sale.IGV}, {"M", sale.Exonerada}, {"N", sale.Inafecta},
		{"O", sale.ISC}, {"P", sale.Base}, {"Q", sale.IVAP}, {"R", sale.ICBPER},
		{"S", sale.Otros}, {"T", sale.Total}, {"U", sale.TipoCambio},
		{"V", sale.FecEmisionMod.String}, {"W", sale.CodigoTipoCDPMod.String},
		{"X", sale.NumSerieCDPMod.String}, {"Y", sale.NumCDPMod.String},
	}

	// for _, cell := range cells {
	// 	if err := f.SetCellValue(sheetName, fmt.Sprintf("%s%d", cell.col, row), cell.val); err != nil {
	// 		log.Printf("Error al establecer el valor en la celda %s: %s", cell.col, err.Error())
	// 		return err
	// 	}
	// }
	// Define un mapa con las columnas que deben usar el numberStyle
	numberStyleColumns := map[string]bool{
		"J": true, "K": true,
		"L": true, "M": true, "N": true, "O": true, "P": true, "Q": true, "R": true, "S": true, "T": true,
	}

	for _, cell := range cells {
		cellRef := fmt.Sprintf("%s%d", cell.col, row)
		value := cell.val
		if cell.col == "U" {
			if num, ok := value.(float64); ok && num == 1 {
				value = ""
			}
		} else {
			if num, ok := value.(float64); ok && num == 0 {
				value = ""
			}
		}
		if err := f.SetCellValue(sheetName, cellRef, value); err != nil {
			log.Printf("Error al establecer el valor en la celda %s: %s", cell.col, err.Error())
			return err
		}
		// Aplica el estilo correspondiente a cada celda
		if _, ok := numberStyleColumns[cell.col]; ok {
			if err := f.SetCellStyle(sheetName, cellRef, cellRef, numberStyle); err != nil {
				log.Printf("Error al establecer el estilo en la celda %s: %s", cell.col, err.Error())
				return err
			}
		} else {
			if err := f.SetCellStyle(sheetName, cellRef, cellRef, fontStyle); err != nil {
				log.Printf("Error al establecer el estilo en la celda %s: %s", cell.col, err.Error())
				return err
			}
		}

	}

	return nil
}

func addTotalRow(f *excelize.File, sheetName string, row int) error {
	// Define el estilo personalizado para la fila de totales
	totalStyle, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold:   true,
			Size:   8,
			Family: "Arial Narrow",
		},
		Alignment: &excelize.Alignment{
			Horizontal: "right",
		},
	})
	if err != nil {
		return err
	}

	// Define el estilo personalizado para los números
	exp := "#,##0.00;-#,##0.00;0"
	numberStyle, err := f.NewStyle(&excelize.Style{
		CustomNumFmt: &exp,
		Alignment: &excelize.Alignment{
			Horizontal: "right",
		},
		Font: &excelize.Font{
			Bold:   true,
			Family: "Arial Narrow",
			Size:   8,
		},
		Border: []excelize.Border{
			{Type: "top", Style: 1, Color: "000000"},
			{Type: "bottom", Style: 6, Color: "000000"},
		},
	})
	if err != nil {
		return err
	}

	// Define las celdas y sus fórmulas
	cells := []struct {
		col string
		val string
	}{
		{"I", "TOTAL"}, {"J", fmt.Sprintf("SUM(J7:J%d)", row-1)},
		{"K", fmt.Sprintf("SUM(K7:K%d)", row-1)}, {"L", fmt.Sprintf("SUM(L7:L%d)", row-1)},
		{"M", fmt.Sprintf("SUM(M7:M%d)", row-1)}, {"N", fmt.Sprintf("SUM(N7:N%d)", row-1)},
		{"O", fmt.Sprintf("SUM(O7:O%d)", row-1)}, {"P", fmt.Sprintf("SUM(P7:P%d)", row-1)},
		{"Q", fmt.Sprintf("SUM(Q7:Q%d)", row-1)}, {"R", fmt.Sprintf("SUM(R7:R%d)", row-1)},
		{"S", fmt.Sprintf("SUM(S7:S%d)", row-1)}, {"T", fmt.Sprintf("SUM(T7:T%d)", row-1)},
	}

	// Establece los valores y aplica el estilo a las celdas
	for _, cell := range cells {
		cellRef := fmt.Sprintf("%s%d", cell.col, row)
		if cell.col == "A" || cell.col == "I" {
			if err := f.SetCellValue(sheetName, cellRef, cell.val); err != nil {
				return err
			}
			if err := f.SetCellStyle(sheetName, cellRef, cellRef, totalStyle); err != nil {
				return err
			}
		} else {
			if err := f.SetCellFormula(sheetName, cellRef, cell.val); err != nil {
				return err
			}
			if cell.val != "0" {
				if err := f.SetCellStyle(sheetName, cellRef, cellRef, numberStyle); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
