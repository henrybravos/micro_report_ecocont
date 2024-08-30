package excel

import (
	"fmt"
	"github.com/henrybravo/micro-report/pkg/files"
	v1 "github.com/henrybravo/micro-report/protos/gen/go/v1"
	"github.com/xuri/excelize/v2"
	"log"
	"sync"
	"time"
)

type SalesGenerator struct{}

func NewSalesGenerator() *SalesGenerator {
	return &SalesGenerator{}
}

func (e *SalesGenerator) GenerateSalesReport(sales []*v1.SalesReport) (path string, err error) {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	sheetName := "Reporte_de_ventas"
	index, err := f.NewSheet(sheetName)
	if err != nil {
		return "", err
	}
	f.SetActiveSheet(index)
	columns := []ColumnSheetStyle{
		{"A", 14}, {"B", 14}, {"C", 14}, {"D", 5}, {"E", 6},
		{"F", 7}, {"G", 5}, {"H", 13}, {"I", 43}, {"J", 10},
		{"K", 12}, {"L", 9}, {"M", 13}, {"N", 11}, {"O", 5},
		{"P", 7}, {"Q", 6}, {"R", 7}, {"S", 8}, {"T", 11},
		{"U", 6}, {"V", 10}, {"W", 6}, {"X", 7}, {"Y", 10},
	}

	if err := setSheetStyles(f, columns, sheetName); err != nil {
		return "", err
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
	if err := createHeaders(f, headers, sheetName); err != nil {
		return "", err
	}
	if err := e.fillSalesData(f, sheetName, sales); err != nil {
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
		}{row: index + 10, sale: sale}
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
		if err := f.SetCellValue(sheetName, fmt.Sprintf("%s%d", cell.col, row), cell.val); err != nil {
			log.Printf("Error al establecer el valor en la celda %s: %s", cell.col, err.Error())
			return err
		}
	}

	return nil
}
