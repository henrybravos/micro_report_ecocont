package report

import (
	"bytes"
	"fmt"
	"github.com/henrybravo/micro-report/internal/repositories"
	"github.com/xuri/excelize/v2"
	"log"
	"sync"
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

func (e *ExcelGenerator) GenerateSalesReport(sales []repositories.SalesReport) (*bytes.Buffer, error) {
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

	if err := createHeaders(f, sheetName); err != nil {
		return nil, err
	}

	if err := fillSalesData(f, sheetName, sales); err != nil {
		return nil, err
	}
	buff, err := f.WriteToBuffer()
	if err != nil {
		return nil, err
	}
	return buff, nil
}

func setSheetStyles(f *excelize.File, sheetName string) error {
	columns := []struct {
		col   string
		width float64
	}{
		{"A", 14}, {"B", 14}, {"C", 14}, {"D", 5}, {"E", 6},
		{"F", 7}, {"G", 5}, {"H", 13}, {"I", 43}, {"J", 10},
		{"K", 12}, {"L", 9}, {"M", 13}, {"N", 11}, {"O", 5},
		{"P", 7}, {"Q", 6}, {"R", 7}, {"S", 8}, {"T", 11},
		{"U", 6}, {"V", 10}, {"W", 6}, {"X", 7}, {"Y", 10},
	}

	for _, col := range columns {
		if err := f.SetColWidth(sheetName, col.col, col.col, col.width); err != nil {
			return err
		}
	}

	return nil
}

func createHeaders(f *excelize.File, sheetName string) error {
	headerStyle, err := createHeaderStyle(f)
	if err != nil {
		return err
	}

	headers := []HeaderTitleCol{
		{"A5", "A9", "CUO"},
		{"B5", "F7", "COMPROBANTE DE PAGO O DOCUMENTO"},
		{"B8", "B9", "FECHA DE EMISIÓN"},
		{"C8", "C9", "FECHA DE VENCIMIENTO"},
		{"D8", "D9", "TIPO"},
		{"E8", "E9", "SERIE"},
		{"F8", "F9", "NUMERO"},
		{"G5", "I5", "INFORMACIÓN DEL CLIENTE"},
		{"G6", "H7", "DOCUMENTO DE IDENTIDAD"},
		{"G8", "G9", "TIPO"},
		{"H8", "H9", "NÚMERO"},
		{"I6", "I9", "APELLIDOS Y NOMBRES O RAZÓN SOCIAL"},
		{"J5", "J9", "VALOR FACTURADO O DE EXPORTACIÓN"},
		{"K5", "K9", "BASE IMPONIBLE DE LA OPERACIÓN GRAVADA"},
		{"L5", "L9", "IGV Y/O IPM"},
		{"M5", "N7", "VALOR TOTAL DE LA OPERACIÓN EXONERADA O INAFECTA"},
		{"M8", "M9", "EXONERADA"},
		{"N8", "N9", "INAFECTA"},
		{"O5", "O9", "ISC"},
		{"P5", "Q7", "OPERACIÓN GRAVADA CON EL IVAP"},
		{"P8", "P9", "BASE IMPONIBLE"},
		{"Q8", "Q9", "IVAP"},
		{"R5", "R9", "ICBPER"},
		{"S5", "S9", "OTROS TRIBUTOS Y CARGOS"},
		{"T5", "T9", "IMPORTE TOTAL"},
		{"U5", "U9", "TIPO DE CAMBIO"},
		{"V5", "Y7", "REFERENCIA DEL COMPROBANTE DE PAGO O DOCUMENTO ORIGINAL QUE SE MODIFICA"},
		{"V8", "V9", "FECHA"},
		{"W8", "W9", "TIPO"},
		{"X8", "X9", "SERIE"},
		{"Y8", "Y9", "NÚMERO"},
	}
	for _, header := range headers {
		if err := createHeaderTitle(f, headerStyle, sheetName, header); err != nil {
			return err
		}
	}

	return nil
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
		Size:   8,
	}
	return f.NewStyle(&excelize.Style{
		Alignment: alignmentStyle,
		Border:    borderStyle,
		Font:      fontStyle,
	})
}

func fillSalesData(f *excelize.File, sheetName string, sales []repositories.SalesReport) error {
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
		}{row: index + 10, sale: sale}
	}
	close(rowCh)

	wg.Wait()
	return nil
}

func setSalesRow(f *excelize.File, sheetName string, row int, sale repositories.SalesReport) error {
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

	for _, cell := range cells {
		if err := f.SetCellValue(sheetName, fmt.Sprintf("%s%d", cell.col, row), cell.val); err != nil {
			log.Printf("Error al establecer el valor en la celda %s: %s", cell.col, err.Error())
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
