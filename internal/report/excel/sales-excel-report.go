package excel

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/henrybravo/micro-report/internal/repositories"
	"github.com/henrybravo/micro-report/pkg/files"
	v1 "github.com/henrybravo/micro-report/protos/gen/go/v1"
	"github.com/xuri/excelize/v2"
)

type SalesGenerator struct{}

func NewSalesGenerator() *SalesGenerator {
	return &SalesGenerator{}
}

func (e *SalesGenerator) GenerateSalesReport(business *repositories.Business, period string, sales []*v1.SalesReport) (path string, err error) {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	defaultSheetName := f.GetSheetName(0)
	sheetName := "Reporte_de_ventas"
	f.SetSheetName(defaultSheetName, sheetName)

	f.SetActiveSheet(0)
	lastRow := len(sales) + 7 // 7 es el número de filas de encabezado

	columns := []ColumnSheetStyle{
		{"A", 9}, {"B", 8}, {"C", 8}, {"D", 3}, {"E", 5},
		{"F", 6}, {"G", 3}, {"H", 10}, {"I", 35}, {"J", 10},
		{"K", 12}, {"L", 9}, {"M", 13}, {"N", 11}, {"O", 5},
		{"P", 7}, {"Q", 6}, {"R", 7}, {"S", 8}, {"T", 11},
		{"U", 6}, {"V", 9}, {"W", 3}, {"X", 6}, {"Y", 10},
	}

	if err := setSheetStyles(f, columns, sheetName); err != nil {
		return "", err
	}

	if err := setSalesExcelSheetStyle(f, sheetName, lastRow); err != nil {
		return "", err
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
	if err := createTitle(f, business, period, sheetName); err != nil {
		return "", err
	}
	if err := createHeaders(f, headers, sheetName); err != nil {
		return "", err
	}
	if err := e.fillSalesData(f, sheetName, sales); err != nil {
		return "", err
	}
	if err := addTotalRow(f, sheetName, lastRow); err != nil {
		return "", err
	}

	path = files.GenerateUniqueNameFile("xlsx")
	err = f.SaveAs("tmp/" + path)
	if err != nil {
		log.Printf("Error al guardar el archivo: %s", err.Error())
	} else {
		log.Printf("Archivo guardado en: %s", "tmp/"+path)
		files.RemoveAfter("tmp/"+path, 5*time.Minute)
	}
	return
}

func (e *SalesGenerator) fillSalesData(f *excelize.File, sheetName string, sales []*v1.SalesReport) error {
	var wg sync.WaitGroup
	rowCh := make(chan struct {
		row  int
		sale *v1.SalesReport
	}, len(sales))

	// Goroutine para llenar datos en paralelo
	for i := 0; i < 4; i++ { // 4 es el número de goroutines, ajusta según sea necesario
		wg.Add(1)
		go func() {
			defer wg.Done()
			for row := range rowCh {
				if err := e.setSalesRow(f, sheetName, row.row, row.sale); err != nil {
					log.Printf("Error al establecer el valor en la celda: %s", err.Error())
				}
			}
		}()
	}

	// Alimenta el canal con los datos
	for index, sale := range sales {
		rowCh <- struct {
			row  int
			sale *v1.SalesReport
		}{row: index + 7, sale: sale}
	}
	close(rowCh)

	wg.Wait()
	return nil
}

func (e *SalesGenerator) setSalesRow(f *excelize.File, sheetName string, row int, sale *v1.SalesReport) error {

	cells := []struct {
		col string
		val interface{}
	}{
		{"A", sale.Cuo}, {"B", sale.FechaEmision}, {"C", sale.FechaVencimiento},
		{"D", sale.CodigoTipoCdp}, {"E", sale.Serie}, {"F", sale.Correlativo},
		{"G", sale.CodTipoDocIdentidad}, {"H", sale.NumDocIdentidadClient},
		{"I", sale.RazonSocial}, {"J", sale.MtoValFactExpo}, {"K", sale.BaseIvap},
		{"L", sale.Igv}, {"M", sale.Exonerada}, {"N", sale.Inafecta},
		{"O", sale.Isc}, {"P", sale.Base}, {"Q", sale.Ivap}, {"R", sale.Icbper},
		{"S", sale.Otros}, {"T", sale.Total}, {"U", sale.TipoCambio},
		{"V", sale.FecEmisionMod}, {"W", sale.CodigoTipoCdpMod},
		{"X", sale.NumSerieCdpMod}, {"Y", sale.NumCdpMod},
	}

	for _, cell := range cells {
		cellRef := fmt.Sprintf("%s%d", cell.col, row)
		value := cell.val
		if cell.col == "U" {
			if num, ok := value.(float32); ok && num == 1 {
				value = ""
			}
		} else {
			if num, ok := value.(float32); ok && num == 0 {
				value = ""
			}
		}
		if err := f.SetCellValue(sheetName, cellRef, value); err != nil {
			log.Printf("Error al establecer el valor en la celda %s: %s", cell.col, err.Error())
			return err
		}
	}

	return nil
}

func createTitle(f *excelize.File, business *repositories.Business, period string, sheetName string) error {
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

func setSalesExcelSheetStyle(f *excelize.File, sheetName string, lastRow int) error {
	showGridLines := false
	showZeros := false
	if err := f.SetSheetView(sheetName, 0, &excelize.ViewOptions{
		ShowGridLines: &showGridLines,
		ShowZeros:     &showZeros,
	}); err != nil {
		return err
	}

	orientation := "landscape"
	scale := uint(65)
	if err := f.SetPageLayout(sheetName, &excelize.PageLayoutOptions{
		Orientation: &orientation,
		AdjustTo: &scale,
	}); err != nil {
		return err
	}

	topMargin := 0.25
	bottomMargin := 0.25
	leftMargin := 0.2
	rightMargin := 0.2
	headerMargin := 0.2
	footerMargin := 0.2
	if err := f.SetPageMargins(sheetName, &excelize.PageLayoutMarginsOptions{
		Top:    &topMargin,
		Bottom: &bottomMargin,
		Left:   &leftMargin,
		Right:  &rightMargin,
		Header: &headerMargin,
		Footer: &footerMargin,
	}); err != nil {
		return err
	}

	if err := setSalesRowStyle(f, sheetName, lastRow); err != nil {
		return err
	}

	return nil
}

func setSalesRowStyle(f *excelize.File, sheetName string, lastRow int) error {
	exp := "#,##0.00_);(#,##0.00)"
	fontConfig := &excelize.Font{
		Family: "Arial Narrow",
		Size:   8,
	}
	numberStyle, _ := f.NewStyle(&excelize.Style{
		CustomNumFmt: &exp,
		Font:         fontConfig,
	})
	fontStyle, _ := f.NewStyle(&excelize.Style{
		Font: fontConfig,
		Alignment: &excelize.Alignment{
            WrapText: true,
        },
	})
	totalRowStyle, rightAlignStyle, err := setTotalRowStyle(f)
	if err != nil {
		return err
	}
	// Define un slice con las columnas que deben usar el numberStyle
	numberStyleColumns := []string{"J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T"}

	// Convierte el slice a un mapa para una búsqueda rápida
	numberStyleMap := make(map[string]bool)
	for _, col := range numberStyleColumns {
		numberStyleMap[col] = true
	}

	// Aplica el estilo correspondiente a cada celda en numberStyleColumns
	for _, col := range numberStyleColumns {
		startCell := fmt.Sprintf("%s7", col)
		endCell := fmt.Sprintf("%s%d", col, lastRow)
		if err := f.SetCellStyle(sheetName, startCell, endCell, numberStyle); err != nil {
			log.Printf("Error al establecer el estilo en la columna %s: %s", col, err.Error())
			return err
		}
	}

	// Aplica el estilo de fuente a las columnas de "A" a "Y" excepto las que están en numberStyleColumns
	for col := 'A'; col <= 'Y'; col++ {
		colStr := string(col)
		if !numberStyleMap[colStr] {
			startCell := fmt.Sprintf("%s7", colStr)
			endCell := fmt.Sprintf("%s%d", colStr, lastRow)
			if err := f.SetCellStyle(sheetName, startCell, endCell, fontStyle); err != nil {
				log.Printf("Error al establecer el estilo en la columna %s: %s", colStr, err.Error())
				return err
			}
		}
	}

	// Aplica el estilo de alineamiento a la derecha a la columna "I" en la fila de totales
	if err := f.SetCellStyle(sheetName, fmt.Sprintf("I%d", lastRow), fmt.Sprintf("I%d", lastRow), rightAlignStyle); err != nil {
		log.Printf("Error al establecer el estilo de alineamiento a la derecha en la columna I: %s", err.Error())
		return err
	}

	// Aplica el estilo a la fila de totales
	if err := f.SetCellStyle(sheetName, fmt.Sprintf("J%d", lastRow), fmt.Sprintf("T%d", lastRow), totalRowStyle); err != nil {
		log.Printf("Error al establecer el estilo en la fila de totales: %s", err.Error())
		return err
	}
	
	return nil
}

func setTotalRowStyle(f *excelize.File) (int, int, error) {
	// Define el estilo de fuente común
	fontStyle := &excelize.Font{
		Bold:   true,
		Size:   8,
		Family: "Arial Narrow",
	}

	// Define el estilo personalizado para la fila de totales
	totalStyle, err := f.NewStyle(&excelize.Style{
		Font: fontStyle,
		Border: []excelize.Border{
			{Type: "top", Style: 1, Color: "000000"},
			{Type: "bottom", Style: 6, Color: "000000"},
		},
	})
	if err != nil {
		return 0, 0, err
	}

	rightAlignStyle, err := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "right",
			Vertical:   "center",
		},
		Font: fontStyle,
	})
	if err != nil {
		return 0, 0, err
	}

	return totalStyle, rightAlignStyle, nil
}

func addTotalRow(f *excelize.File, sheetName string, row int) error {

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

	// Establece los valores las celdas
	for _, cell := range cells {
		cellRef := fmt.Sprintf("%s%d", cell.col, row)
		if cell.col == "I" {
			if err := f.SetCellValue(sheetName, cellRef, cell.val); err != nil {
				return err
			}
		} else {
			if err := f.SetCellFormula(sheetName, cellRef, cell.val); err != nil {
				return err
			}
		}
	}

	return nil
}
