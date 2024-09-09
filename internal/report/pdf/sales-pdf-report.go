package pdf

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"time"

	"github.com/henrybravo/micro-report/internal/repositories"
	"github.com/henrybravo/micro-report/pkg/files"
	v1 "github.com/henrybravo/micro-report/protos/gen/go/v1"
	"github.com/signintech/gopdf"
	"strings"
)

type SalesGenerator struct{}

func NewSalesGenerator() *SalesGenerator {
	return &SalesGenerator{}
}

type layout struct {
	marginX      float64
	marginY      float64
	headerPageH  float64
	pageW        float64
	pageH        float64
	headerTableH float64
	rowTableH    float64
	textH        float64
	cuoW         float64

	cpeInfoW   float64
	cpeFecEmiW float64
	cpeFecVenW float64
	cpeTipoW   float64
	cpeSerieW  float64
	cpeNumW    float64

	clienteInfoW float64
	cliDocTipoW  float64
	cliDocNumW   float64
	cliApeNomW   float64

	valFacOExpW float64
	baseImpW    float64
	igvW        float64

	totalExoInaW float64
	totalExoW    float64
	totalInaW    float64

	iscW float64

	opGravIvapW float64
	opBaseW     float64
	opIVAPW     float64

	icbW      float64
	otrosW    float64
	impTotalW float64
	tcW       float64

	refComW   float64
	refComFec float64
	refComTip float64
	refComSer float64
	refComNum float64
}

func (p *SalesGenerator) initializeLayout() layout {
	return layout{

		headerPageH:  2,
		marginX:      0.5,
		marginY:      0.35,
		pageW:        29.7,
		pageH:        21.0,
		headerTableH: 1.5,
		rowTableH:    0.35,
		textH:        0.18,
		cuoW:         1.4,

		cpeInfoW:   4.1,
		cpeFecEmiW: 4.1 / 4 + 0.15,
		cpeFecVenW: 4.1 / 4 + 0.15,
		cpeTipoW:   4.1 / 6 - 0.3,
		cpeSerieW:  4.1 / 6,
		cpeNumW:    4.1 / 6,

		clienteInfoW: 6.5,
		cliDocTipoW:  6.5 / 8,
		cliDocNumW:   6.5 / 4.8,
		cliApeNomW:   2 * 6.5 / 3,

		valFacOExpW: 1.3,
		baseImpW:    1.4,
		igvW:        1.1,

		totalExoInaW: 2.6,
		totalExoW:    2.6 / 2,
		totalInaW:    2.6 / 2,

		iscW: 0.7,

		opGravIvapW: 2.0,
		opBaseW:     2.0 / 2,
		opIVAPW:     2.0 / 2,

		icbW:      0.9,
		otrosW:    1.3,
		impTotalW: 1.2,
		tcW:       0.9,

		refComW:   3.3,
		refComFec: 3.3 / 4 + 0.2,
		refComTip: 3.3 / 4 - 0.2,
		refComSer: 3.3 / 4,
		refComNum: 3.3 / 4,
	}
}
func (p *SalesGenerator) GenerateSalesReport(business *repositories.Business, sales []*v1.SalesReport, period string) (path string, err error) {
	layout := p.initializeLayout()
	pdf := gopdf.GoPdf{}
	pdf.SetCompressLevel(9) //compress content streams
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4Landscape, Unit: gopdf.UnitCM})
	err = pdf.AddTTFFont("arial", "./fonts/ARIAL.TTF")
	err = pdf.AddTTFFont("arialB", "./fonts/ARIALBD.TTF")
	if err != nil {
		log.Print(err.Error())
		return
	}
	err = generatePage(business, period, &pdf, layout)
	var (
		sumMtoValFactExpo float32
		sumBase           float32
		sumIgv            float32
		sumExonerada      float32
		sumInafecta       float32
		sumIsc            float32
		sumBaseIvap       float32
		sumIcbper         float32
		sumOtros          float32
		sumMtoTotalCp     float32
	)

	bandComings := false
	lensales := len(sales)
	for _, sale := range sales {
		lensales--
		if checkNewPage(&pdf, layout) {
			err = generatePage(business, period, &pdf, layout)
			if err != nil {
				return "", err
			}
			bandComings = true
		}
		numCells := determineNumCells(layout, &pdf, sale)
		locationY := pdf.GetY() + layout.rowTableH
		if numCells > 1 {
			locationY = pdf.GetY() + layout.rowTableH + float64(numCells - 1) * layout.textH
		}
		if bandComings {
			//for comming true and for going false
			locationY += layout.rowTableH
			err = generateComingsAndGoings(&pdf, locationY, layout, true, float64(numCells), sumMtoValFactExpo, sumBase, sumIgv, sumExonerada, sumInafecta, sumIsc, sumBaseIvap, sumIcbper, sumOtros, sumMtoTotalCp)
			if err != nil {
				return "", err
			}
			bandComings = false
		}
		err = generateRowTable(&pdf, sale, locationY, layout, numCells)
		sumMtoValFactExpo += sale.MtoValFactExpo
		sumBase += sale.Base
		sumIgv += sale.Igv
		sumExonerada += sale.Exonerada
		sumInafecta += sale.Inafecta
		sumIsc += sale.Isc
		sumBaseIvap += sale.BaseIvap
		sumIcbper += sale.Icbper
		sumOtros += sale.Otros
		sumMtoTotalCp += sale.MtoTotalCp

		//check for goings (van)
		if pdf.GetY()+2*layout.rowTableH*2+layout.marginY > layout.pageH && lensales > 0 {
			err = generateComingsAndGoings(&pdf, locationY, layout, false, 1, sumMtoValFactExpo, sumBase, sumIgv, sumExonerada, sumInafecta, sumIsc, sumBaseIvap, sumIcbper, sumOtros, sumMtoTotalCp)
		}
		//for total
		if lensales == 0 {
			if checkNewPage(&pdf, layout) {
				err = generatePage(business, period, &pdf, layout)
				if err != nil {
					return "", err
				}
			}
			err = addTotal(&pdf, locationY, layout, sumMtoValFactExpo, sumBase, sumIgv, sumExonerada, sumInafecta, sumIsc, sumBaseIvap, sumIcbper, sumOtros, sumMtoTotalCp)
			if err != nil {
				return "", err
			}
		}
	}
	if err != nil {
		log.Print(err.Error())
		return
	}
	path = files.GenerateUniqueNameFile("pdf")
	err = pdf.WritePdf("tmp/" + path)
	if err != nil {
		fmt.Println("Error writing file:", err)
	} else {
		fmt.Println("PDF created send successfully")
		files.RemoveAfter("tmp/"+path, 5*time.Minute)
	}
	return
}
func checkNewPage(pdf *gopdf.GoPdf, layout layout) bool {
	return pdf.GetY()+2*layout.rowTableH+2*layout.marginY > layout.pageH
}
func determineNumCells(layout layout, pdf *gopdf.GoPdf, sale *v1.SalesReport) int {
	razonSocialRows := calculateRows(layout.cliApeNomW, pdf, sale.RazonSocial)
	cuoRows := calculateRows(layout.cuoW, pdf, sale.Cuo)
	rows := []int{razonSocialRows, cuoRows}
	maxRows := rows[0]
	for _, row := range rows[1:] {
		maxRows = int(math.Max(float64(maxRows), float64(row)))
	}
	return maxRows
}
func calculateRows(length float64, pdf *gopdf.GoPdf, text string) int {
	cellMaxCharacters,_ := pdf.MeasureTextWidth(text)
	return int(math.Ceil(cellMaxCharacters/ length))
}
func generatePage(business *repositories.Business, period string, pdf *gopdf.GoPdf, layout layout) error {
	pdf.AddPage()
	pdf.SetXY(0, layout.marginY)
	err := generateHeaderPage(business, period, pdf, layout)
	pdf.SetXY(layout.marginX, layout.headerPageH)
	err = generateHeaderTable(pdf, layout)
	pdf.SetXY(layout.marginX, layout.headerPageH+layout.headerTableH)
	return err
}
func generateHeaderPage(business *repositories.Business, period string, pdf *gopdf.GoPdf, layout layout) (err error) {
	err = pdf.SetFont("arialB", "", 10)
	rect := &gopdf.Rect{
		H: 0.5,
		W: layout.pageW,
	}
	cellOptionCenter := gopdf.CellOption{
		Align: gopdf.Middle | gopdf.Center,
	}
	err = pdf.CellWithOption(rect, business.BusinessName, cellOptionCenter)
	pdf.SetXY(0, 0.7)
	err = pdf.CellWithOption(rect, "R.U.C.: "+business.RUC, cellOptionCenter)
	pdf.SetXY(0, 1.1)
	err = pdf.CellWithOption(rect, business.Address, cellOptionCenter)
	pdf.SetXY(0, 1.5)
	err = pdf.CellWithOption(rect, "REGISTRO DE VENTAS DEL MES DE "+period, cellOptionCenter)
	return
}
func generateHeaderTable(pdf *gopdf.GoPdf, layout layout) error {
	err := pdf.SetFont("arial", "", 5.5)
	cellOptionAllBorderCenter := gopdf.CellOption{
		Border: gopdf.AllBorders,
		Align:  gopdf.Middle | gopdf.Center,
	}
	cellOptionCenter := gopdf.CellOption{
		Align: gopdf.Middle | gopdf.Center,
	}
	cellOptionBorderRBCenter := gopdf.CellOption{
		Border: gopdf.Right | gopdf.Bottom,
		Align:  gopdf.Middle | gopdf.Center,
	}
	cellOptionBorderRCenter := gopdf.CellOption{
		Border: gopdf.Right,
		Align:  gopdf.Middle | gopdf.Center,
	}
	cellOptionBorderRTCenter := gopdf.CellOption{
		Border: gopdf.Right | gopdf.Top,
		Align:  gopdf.Middle | gopdf.Center,
	}

	rect := &gopdf.Rect{
		H: layout.headerTableH,
		W: layout.cuoW,
	}
	err = pdf.CellWithOption(rect, "CUO", cellOptionAllBorderCenter)
	rect = &gopdf.Rect{
		H: layout.headerTableH / 2,
		W: layout.cpeInfoW,
	}
	err = pdf.CellWithOption(rect, "DATOS DE CP", cellOptionAllBorderCenter)
	pdf.SetXY(layout.cuoW+layout.marginX, (layout.headerTableH/2)+layout.headerPageH)
	rect = &gopdf.Rect{
		H: layout.headerTableH / 4,
		W: layout.cpeFecEmiW,
	}
	err = pdf.CellWithOption(rect, "FECHA DE", cellOptionBorderRCenter)
	pdf.SetXY(pdf.GetX()-layout.cpeFecEmiW, pdf.GetY()+layout.headerTableH/4)
	err = pdf.CellWithOption(rect, "EMISIÓN", cellOptionBorderRBCenter)
	pdf.SetXY(pdf.GetX(), pdf.GetY()-layout.headerTableH/4)
	rect = &gopdf.Rect{
		H: layout.headerTableH / 4,
		W: layout.cpeFecVenW,
	}
	err = pdf.CellWithOption(rect, "FECHA DE", cellOptionCenter)
	pdf.SetXY(pdf.GetX()-layout.cpeFecVenW, pdf.GetY()+layout.headerTableH/4)
	err = pdf.CellWithOption(rect, "VENCIMIENTO", cellOptionBorderRBCenter)
	pdf.SetXY(pdf.GetX(), pdf.GetY()-layout.headerTableH/4)
	rect = &gopdf.Rect{
		H: layout.headerTableH / 2,
		W: layout.cpeTipoW,
	}
	err = pdf.CellWithOption(rect, "TIPO", cellOptionAllBorderCenter)
	rect = &gopdf.Rect{
		H: layout.headerTableH / 2,
		W: layout.cpeSerieW,
	}
	err = pdf.CellWithOption(rect, "SERIE", cellOptionAllBorderCenter)
	rect = &gopdf.Rect{
		H: layout.headerTableH / 2,
		W: layout.cpeNumW,
	}
	err = pdf.CellWithOption(rect, "NÚMERO", cellOptionAllBorderCenter)
	pdf.SetXY(layout.cuoW+layout.cpeInfoW+layout.marginX, layout.headerPageH)
	err = pdf.CellWithOption(&gopdf.Rect{
		H: layout.headerTableH / 4,
		W: layout.clienteInfoW,
	}, "INFORMACIÓN DEL CLIENTE", cellOptionAllBorderCenter)
	pdf.SetXY(pdf.GetX()-layout.clienteInfoW, layout.headerPageH+layout.headerTableH/4)
	rect = &gopdf.Rect{
		H: layout.headerTableH / 4,
		W: layout.clienteInfoW / 3,
	}
	err = pdf.CellWithOption(rect, "DOCUMENTO DE", cellOptionCenter)
	pdf.SetXY(pdf.GetX()-layout.clienteInfoW/3, layout.headerPageH+layout.headerTableH/2)
	err = pdf.CellWithOption(rect, "IDENTIDAD", cellOptionCenter)
	pdf.SetXY(pdf.GetX()-layout.clienteInfoW/3, layout.headerPageH+3*layout.headerTableH/4)
	rect = &gopdf.Rect{
		H: layout.headerTableH / 4,
		W: layout.clienteInfoW / 8,
	}
	err = pdf.CellWithOption(rect, "TIPO", cellOptionAllBorderCenter)
	rect = &gopdf.Rect{
		H: layout.headerTableH / 4,
		W: layout.clienteInfoW / 4.8,
	}
	err = pdf.CellWithOption(rect, "NÚMERO", cellOptionAllBorderCenter)
	pdf.SetXY(pdf.GetX(), layout.headerPageH+layout.headerTableH/4)
	rect = &gopdf.Rect{
		H: 3 * layout.headerTableH / 4,
		W: 2 * layout.clienteInfoW / 3,
	}
	err = pdf.CellWithOption(rect, "APELLIDOS Y NOMBRES O RAZÓN SOCIAL", cellOptionAllBorderCenter)
	pdf.SetY(layout.headerPageH)
	rect = &gopdf.Rect{
		H: layout.headerTableH / 4,
		W: layout.valFacOExpW,
	}
	err = pdf.CellWithOption(rect, "VALOR", cellOptionBorderRTCenter)
	pdf.SetXY(pdf.GetX()-layout.valFacOExpW, layout.headerPageH+layout.headerTableH/4)
	err = pdf.CellWithOption(rect, "FACTURADO O", cellOptionBorderRCenter)
	pdf.SetXY(pdf.GetX()-layout.valFacOExpW, layout.headerPageH+layout.headerTableH/2)
	err = pdf.CellWithOption(rect, "DE", cellOptionBorderRCenter)
	pdf.SetXY(pdf.GetX()-layout.valFacOExpW, layout.headerPageH+3*layout.headerTableH/4)
	err = pdf.CellWithOption(rect, "EXPORTACIÓN", cellOptionBorderRBCenter)
	pdf.SetY(layout.headerPageH)
	rect = &gopdf.Rect{
		H: layout.headerTableH / 4,
		W: layout.baseImpW,
	}
	err = pdf.CellWithOption(rect, "BASE IMPONIBLE", cellOptionBorderRTCenter)
	pdf.SetXY(pdf.GetX()-layout.baseImpW, pdf.GetY()+layout.headerTableH/4)
	err = pdf.CellWithOption(rect, "DE LA", cellOptionCenter)
	pdf.SetXY(pdf.GetX()-layout.baseImpW, pdf.GetY()+layout.headerTableH/4)
	err = pdf.CellWithOption(rect, "OPERACIÓN", cellOptionCenter)
	pdf.SetXY(pdf.GetX()-layout.baseImpW, pdf.GetY()+layout.headerTableH/4)
	err = pdf.CellWithOption(rect, "GRAVADA", cellOptionBorderRBCenter)
	pdf.SetY(layout.headerPageH)
	rect = &gopdf.Rect{
		H: layout.headerTableH,
		W: layout.igvW,
	}
	err = pdf.CellWithOption(rect, "IGV Y/O IPM", cellOptionAllBorderCenter)
	rect = &gopdf.Rect{
		H: layout.headerTableH / 4,
		W: layout.totalExoInaW,
	}
	err = pdf.CellWithOption(rect, "VALOR TOTAL DE LA", cellOptionBorderRTCenter)
	pdf.SetXY(pdf.GetX()-layout.totalExoInaW, pdf.GetY()+layout.headerTableH/4)
	err = pdf.CellWithOption(rect, "OPERACIÓN EXONERADA", cellOptionCenter)
	pdf.SetXY(pdf.GetX()-layout.totalExoInaW, pdf.GetY()+layout.headerTableH/4)
	err = pdf.CellWithOption(rect, "O INAFECTA", cellOptionCenter)
	pdf.SetXY(pdf.GetX()-layout.totalExoInaW, pdf.GetY()+layout.headerTableH/4)
	rect = &gopdf.Rect{
		H: layout.headerTableH / 4,
		W: layout.totalExoInaW / 2,
	}
	err = pdf.CellWithOption(rect, "EXONERADA", cellOptionAllBorderCenter)
	err = pdf.CellWithOption(rect, "INAFECTA", cellOptionAllBorderCenter)
	pdf.SetY(layout.headerPageH)
	rect = &gopdf.Rect{
		H: layout.headerTableH,
		W: layout.iscW,
	}
	err = pdf.CellWithOption(rect, "ISC", cellOptionAllBorderCenter)
	rect = &gopdf.Rect{
		H: layout.headerTableH / 4,
		W: layout.opGravIvapW,
	}
	err = pdf.CellWithOption(rect, "OPERACIÓN GRAVADA", cellOptionBorderRTCenter)
	pdf.SetXY(pdf.GetX()-layout.opGravIvapW, pdf.GetY()+layout.headerTableH/4)
	err = pdf.CellWithOption(rect, "CON EL IVAP", cellOptionCenter)
	pdf.SetXY(pdf.GetX()-layout.opGravIvapW, pdf.GetY()+layout.headerTableH/4)
	rect = &gopdf.Rect{
		H: layout.headerTableH / 4,
		W: layout.opGravIvapW / 2,
	}
	err = pdf.CellWithOption(rect, "BASE", cellOptionBorderRTCenter)
	pdf.SetXY(pdf.GetX()-(layout.opGravIvapW/2), pdf.GetY()+layout.headerTableH/4)
	err = pdf.CellWithOption(rect, "IMPONIBLE", cellOptionBorderRBCenter)
	pdf.SetY(layout.headerPageH + layout.headerTableH/2)
	err = pdf.CellWithOption(&gopdf.Rect{
		H: layout.headerTableH / 2,
		W: layout.opGravIvapW / 2,
	}, "IVAP", cellOptionAllBorderCenter)
	pdf.SetY(layout.headerPageH)
	err = pdf.CellWithOption(&gopdf.Rect{
		H: layout.headerTableH,
		W: layout.icbW,
	}, "ICB PER", cellOptionAllBorderCenter)
	rect = &gopdf.Rect{
		H: layout.headerTableH / 4,
		W: layout.otrosW,
	}
	err = pdf.CellWithOption(rect, "OTROS", cellOptionBorderRTCenter)
	pdf.SetXY(pdf.GetX()-layout.otrosW, pdf.GetY()+layout.headerTableH/4)
	err = pdf.CellWithOption(rect, "TRIBUTOS", cellOptionBorderRCenter)
	pdf.SetXY(pdf.GetX()-layout.otrosW, pdf.GetY()+layout.headerTableH/4)
	err = pdf.CellWithOption(rect, "Y", cellOptionBorderRCenter)
	pdf.SetXY(pdf.GetX()-layout.otrosW, pdf.GetY()+layout.headerTableH/4)
	err = pdf.CellWithOption(rect, "CARGOS", cellOptionBorderRBCenter)
	pdf.SetY(layout.headerPageH)
	rect = &gopdf.Rect{
		H: layout.headerTableH / 2,
		W: layout.impTotalW,
	}
	err = pdf.CellWithOption(rect, "IMPORTE", cellOptionBorderRTCenter)
	pdf.SetXY(pdf.GetX()-layout.impTotalW, pdf.GetY()+layout.headerTableH/2)
	err = pdf.CellWithOption(rect, "TOTAL", cellOptionBorderRBCenter)
	pdf.SetY(layout.headerPageH)
	rect = &gopdf.Rect{
		H: layout.headerTableH / 3,
		W: layout.tcW,
	}
	err = pdf.CellWithOption(rect, "TIPO", cellOptionBorderRTCenter)
	pdf.SetXY(pdf.GetX()-layout.tcW, pdf.GetY()+layout.headerTableH/3)
	err = pdf.CellWithOption(rect, "DE", cellOptionBorderRCenter)
	pdf.SetXY(pdf.GetX()-layout.tcW, pdf.GetY()+layout.headerTableH/3)
	err = pdf.CellWithOption(rect, "CAMBIO", cellOptionBorderRBCenter)
	pdf.SetY(layout.headerPageH)
	rect = &gopdf.Rect{
		H: layout.headerTableH / 4,
		W: layout.refComW,
	}
	err = pdf.CellWithOption(rect, "REFERENCIA DEL COMPROBANTE", cellOptionBorderRTCenter)
	pdf.SetXY(pdf.GetX()-layout.refComW, pdf.GetY()+layout.headerTableH/4)
	err = pdf.CellWithOption(rect, "DE PAGO O DOCUMENTO ORIGINAL", cellOptionBorderRCenter)
	pdf.SetXY(pdf.GetX()-layout.refComW, pdf.GetY()+layout.headerTableH/4)
	err = pdf.CellWithOption(rect, "QUE SE MODIFICA", cellOptionBorderRCenter)
	pdf.SetXY(pdf.GetX()-layout.refComW, pdf.GetY()+layout.headerTableH/4)
	rect = &gopdf.Rect{
		H: layout.headerTableH / 4,
		W: layout.refComFec,
	}
	err = pdf.CellWithOption(rect, "FECHA", cellOptionAllBorderCenter)
	rect = &gopdf.Rect{
		H: layout.headerTableH / 4,
		W: layout.refComTip,
	}
	err = pdf.CellWithOption(rect, "TIPO", cellOptionAllBorderCenter)
	rect = &gopdf.Rect{
		H: layout.headerTableH / 4,
		W: layout.refComSer,
	}
	err = pdf.CellWithOption(rect, "SERIE", cellOptionAllBorderCenter)
	rect = &gopdf.Rect{
		H: layout.headerTableH / 4,
		W: layout.refComNum,
	}
	err = pdf.CellWithOption(rect, "NÚMERO", cellOptionAllBorderCenter)
	return err
}
func generateRowTable(pdf *gopdf.GoPdf, sale *v1.SalesReport, locationY float64, layout layout, numCells int) error {
	err := pdf.SetFont("arial", "", 7)
	rowMiddle := locationY + layout.textH
	if numCells > 1 {
		rowMiddle = locationY + (layout.rowTableH + layout.textH*float64(numCells-1) - layout.textH - layout.textH*float64(numCells-1))/2
	}
	rowW := layout.pageW - layout.marginX*2
	marginText := 0.1
	currentWriteW := layout.marginX + marginText
	pdf.SetXY(layout.marginX, locationY)
	pdf.SetStrokeColor(222, 219, 218)
	pdf.SetLineWidth(0)
	cellOptionBottom := gopdf.CellOption{
		Border: gopdf.Bottom,
	}
	rect := &gopdf.Rect{
		H: layout.rowTableH,
		W: rowW,
	}
	err = pdf.CellWithOption(rect, "", cellOptionBottom)
	pdf.SetXY(currentWriteW, rowMiddle)
	numRowsCuo := calculateRows(layout.cuoW, pdf, sale.Cuo)
	if numRowsCuo > 1{
		pdf.SetXY(currentWriteW, rowMiddle - (layout.textH * float64(numRowsCuo-1) + marginText*float64(numRowsCuo-1)))
		rect := &gopdf.Rect{
			H: layout.rowTableH * float64(numCells),
			W: layout.cuoW,
		}
		err = pdf.MultiCell(rect, sale.Cuo)
	}else{
		err = pdf.Text(sale.Cuo)
	}
	currentWriteW += layout.cuoW
	pdf.SetXY(currentWriteW, rowMiddle)
	err = pdf.Text(sale.FecEmision)
	currentWriteW += layout.cpeFecEmiW
	pdf.SetXY(currentWriteW, rowMiddle)
	err = pdf.Text(sale.FechaVencimiento)
	currentWriteW += layout.cpeFecVenW
	alignCenter(pdf, sale.CodigoTipoCdp, layout.cpeTipoW, currentWriteW, rowMiddle, marginText)
	currentWriteW += layout.cpeTipoW
	alignCenter(pdf, sale.NumSerieCdp, layout.cpeSerieW, currentWriteW, rowMiddle, marginText)
	currentWriteW += layout.cpeSerieW
	alignCenter(pdf, sale.NumCdp, layout.cpeNumW, currentWriteW, rowMiddle, marginText)
	currentWriteW += layout.cpeNumW
	alignCenter(pdf, sale.CodTipoDocIdentidad, layout.cliDocTipoW, currentWriteW, rowMiddle, marginText)
	currentWriteW += layout.cliDocTipoW
	pdf.SetXY(currentWriteW, rowMiddle)
	err = pdf.Text(sale.NumDocIdentidadClient)
	currentWriteW += layout.cliDocNumW
	numRowsRazonSocial := calculateRows(layout.cliApeNomW, pdf, sale.RazonSocial)
	if numRowsRazonSocial > 1 {
		pdf.SetXY(currentWriteW, rowMiddle - (layout.textH * float64(numRowsRazonSocial-1) + marginText*float64(numRowsRazonSocial-1)))
		rect := &gopdf.Rect{
			H: layout.rowTableH * float64(numCells),
			W: layout.cliApeNomW,
		}
		err = pdf.MultiCell(rect, sale.RazonSocial)
	} else {
		pdf.SetXY(currentWriteW, rowMiddle)
		err = pdf.Text(sale.RazonSocial)
	}
	currentWriteW += layout.cliApeNomW
	alignRight(pdf, sale.MtoValFactExpo, layout.valFacOExpW, currentWriteW, rowMiddle, marginText)
	currentWriteW += layout.valFacOExpW
	alignRight(pdf, sale.Base, layout.baseImpW, currentWriteW, rowMiddle, marginText)
	currentWriteW += layout.baseImpW
	alignRight(pdf, sale.Igv, layout.igvW, currentWriteW, rowMiddle, marginText)
	currentWriteW += layout.igvW
	alignRight(pdf, sale.Exonerada, layout.totalExoW, currentWriteW, rowMiddle, marginText)
	currentWriteW += layout.totalExoW
	alignRight(pdf, sale.Inafecta, layout.totalInaW, currentWriteW, rowMiddle, marginText)
	currentWriteW += layout.totalInaW
	alignRight(pdf, sale.Isc, layout.iscW, currentWriteW, rowMiddle, marginText)
	currentWriteW += layout.iscW
	alignRight(pdf, sale.Base, layout.opBaseW, currentWriteW, rowMiddle, marginText)
	currentWriteW += layout.opBaseW
	alignRight(pdf, sale.BaseIvap, layout.opIVAPW, currentWriteW, rowMiddle, marginText)
	currentWriteW += layout.opIVAPW
	alignRight(pdf, sale.Icbper, layout.icbW, currentWriteW, rowMiddle, marginText)
	currentWriteW += layout.icbW
	alignRight(pdf, sale.Otros, layout.otrosW, currentWriteW, rowMiddle, marginText)
	currentWriteW += layout.otrosW
	alignRight(pdf, sale.MtoTotalCp, layout.impTotalW, currentWriteW, rowMiddle, marginText)
	currentWriteW += layout.impTotalW
	pdf.SetXY(currentWriteW, rowMiddle)
	if sale.TipoCambio == 1 {
		err = pdf.Text("")
	} else {
		err = pdf.Text(fmt.Sprintf("%.2f", sale.TipoCambio))
	}
	currentWriteW += layout.tcW
	alignCenter(pdf, sale.FecEmisionMod, layout.refComFec, currentWriteW, rowMiddle, marginText)
	currentWriteW += layout.refComFec
	alignCenter(pdf, sale.CodigoTipoCdpMod, layout.refComTip, currentWriteW, rowMiddle, marginText)
	currentWriteW += layout.refComTip
	alignCenter(pdf, sale.NumSerieCdpMod, layout.refComSer, currentWriteW, rowMiddle, marginText)
	currentWriteW += layout.refComSer
	alignCenter(pdf, sale.NumCdpMod, layout.refComNum, currentWriteW, rowMiddle, marginText)
	return err
}
func generateComingsAndGoings(pdf *gopdf.GoPdf, locationY float64, layout layout, goingComing bool, numCells float64,
	sumMtoValFactExpo float32, sumBase float32,
	sumIgv float32, sumExonerada float32, sumInafecta float32,
	sumIsc float32, sumBaseIvap float32, sumIcbper float32, sumOtros float32, sumMtoTotalCp float32) error {
	if !goingComing {
		locationY += layout.rowTableH
	} else {
		offset := 0.2
		if numCells > 1 {
			offset = 0.02 * (numCells - 1)
		}
		locationY -= layout.rowTableH*numCells + offset
	}
	err := pdf.SetFont("arialB", "", 7)
	if err != nil {
		log.Print(err.Error())
	}
	rowMiddle := locationY + 2*layout.rowTableH/3
	if numCells > 1 {
		rowMiddle = locationY + 2*layout.rowTableH/4
	}
	marginText := 0.075
	currentWriteW := layout.marginX + marginText
	writeCell := func(width float64, text string, alignRigth bool) {
		pdf.SetXY(currentWriteW, locationY)
		cellOption := gopdf.CellOption{
			Align: gopdf.Left | gopdf.Middle,
		}
		if !goingComing && text != "" && text != "VAN" {
			pdf.SetStrokeColor(50, 50, 50)
			cellOption.Border = gopdf.Top
		}
		if goingComing {
			pdf.SetStrokeColor(222, 219, 218)
			cellOption.Border = gopdf.Bottom
		}
		rect := &gopdf.Rect{
			H: layout.rowTableH,
			W: width,
		}
		err = pdf.CellWithOption(rect, "", cellOption)
		if err != nil {
			log.Fatalf("Error writing cell: %v", err)
		}
		if goingComing {
			pdf.SetXY(currentWriteW, rowMiddle)
		} else {
			pdf.SetXY(currentWriteW, rowMiddle+0.03)
		}
		if alignRigth {
			value, err := strconv.ParseFloat(text, 32)
			if err != nil {
				log.Fatalf("Error converting text to float32: %v", err)
			}
			alignRight(pdf, float32(value), width, currentWriteW, rowMiddle, marginText)
		} else {
			err = pdf.Text(text)

			if err != nil {
				log.Fatalf("Error writing text: %v", err)
			}
		}
		currentWriteW += width
	}
	pdf.SetLineWidth(0)
	writeCell(layout.cuoW, "", false)
	writeCell(layout.cpeFecEmiW, "", false)
	writeCell(layout.cpeFecVenW, "", false)
	writeCell(layout.cpeTipoW, "", false)
	writeCell(layout.cpeSerieW, "", false)
	writeCell(layout.cpeNumW, "", false)
	writeCell(layout.cliDocTipoW, "", false)
	writeCell(layout.cliDocNumW, "", false)
	writeCell(layout.cliApeNomW, func() string {
		if !goingComing {
			return "VAN"
		}
		return "VIENEN"
	}(), false)
	writeCell(layout.valFacOExpW, fmt.Sprintf("%.2f", sumMtoValFactExpo), true)
	writeCell(layout.baseImpW, fmt.Sprintf("%.2f", sumBase), true)
	writeCell(layout.igvW, fmt.Sprintf("%.2f", sumIgv), true)
	writeCell(layout.totalExoW, fmt.Sprintf("%.2f", sumExonerada), true)
	writeCell(layout.totalInaW, fmt.Sprintf("%.2f", sumInafecta), true)
	writeCell(layout.iscW, fmt.Sprintf("%.2f", sumIsc), true)
	writeCell(layout.opBaseW, fmt.Sprintf("%.2f", sumBase), true)
	writeCell(layout.opIVAPW, fmt.Sprintf("%.2f", sumBaseIvap), true)
	writeCell(layout.icbW, fmt.Sprintf("%.2f", sumIcbper), true)
	writeCell(layout.otrosW, fmt.Sprintf("%.2f", float32(100)), true)
	writeCell(layout.impTotalW, fmt.Sprintf("%.2f", sumMtoTotalCp), true)
	writeCell(layout.tcW, "", false)
	writeCell(layout.refComFec, "", false)
	writeCell(layout.refComTip, "", false)
	writeCell(layout.refComSer, "", false)

	return err
}
func addTotal(pdf *gopdf.GoPdf, locationY float64, layout layout, sumMtoValFactExpo float32, sumBase float32,
	sumIgv float32, sumExonerada float32, sumInafecta float32, sumIsc float32, sumBaseIvap float32,
	sumIcbper float32, sumOtros float32, sumMtoTotalCp float32) error {
	err := pdf.SetFont("arialB", "", 7)
	if err != nil {
		log.Print(err.Error())
	}
	rowMiddle := locationY + 2*layout.rowTableH - layout.textH/3
	marginText := 0.075
	currentWriteW := layout.marginX + marginText
	writeCell := func(width float64, text string, borderTop bool, alignRigth bool) {
		margin := 0.0
		if borderTop {
			margin = 0.03
		}
		pdf.SetXY(currentWriteW, locationY+margin)
		cellOption := gopdf.CellOption{
			Align: gopdf.Left | gopdf.Middle,
		}
		if text != "" {
			if borderTop || text != "TOTAL" {
				pdf.SetStrokeColor(50, 50, 50)
				cellOption.Border = gopdf.Bottom
			}
		}
		rect := &gopdf.Rect{
			H: layout.rowTableH,
			W: width,
		}
		err = pdf.CellWithOption(rect, "", cellOption)
		if err != nil {
			log.Fatalf("Error writing cell: %v", err)
		}
		pdf.SetXY(currentWriteW, rowMiddle)
		if !borderTop {
			if alignRigth {
				value, err := strconv.ParseFloat(text, 32)
				if err != nil {
					log.Fatalf("Error converting text to float: %v", err)
				}
				alignRight(pdf, float32(value), width, currentWriteW, rowMiddle, marginText)
			} else if text == "TOTAL" {
				err = pdf.Text(text)
			}
			currentWriteW += width
		}

		if text != "" && borderTop {
			currentWriteW += width
			margin := 0.06
			pdf.SetXY(currentWriteW-width, locationY+layout.rowTableH+margin)
			cellOptionTop := gopdf.CellOption{
				Border: gopdf.Top,
			}
			rectTop := &gopdf.Rect{
				H: layout.rowTableH,
				W: width,
			}
			err = pdf.CellWithOption(rectTop, "", cellOptionTop)
			if err != nil {
				log.Fatalf("Error writing top border cell: %v", err)
			}
		} else if text == "" && borderTop {
			currentWriteW += width
		}
	}
	pdf.SetLineWidth(0)
	writeCell(layout.cuoW, "", false, false)
	writeCell(layout.cpeFecEmiW, "", false, false)
	writeCell(layout.cpeFecVenW, "", false, false)
	writeCell(layout.cpeTipoW, "", false, false)
	writeCell(layout.cpeSerieW, "", false, false)
	writeCell(layout.cpeNumW, "", false, false)
	writeCell(layout.cliDocTipoW, "", false, false)
	writeCell(layout.cliDocNumW, "", false, false)
	writeCell(layout.cliApeNomW, "TOTAL", false, false)
	writeCell(layout.valFacOExpW, fmt.Sprintf("%.2f", sumMtoValFactExpo), false, true)
	writeCell(layout.baseImpW, fmt.Sprintf("%.2f", sumBase), false, true)
	writeCell(layout.igvW, fmt.Sprintf("%.2f", sumIgv), false, true)
	writeCell(layout.totalExoW, fmt.Sprintf("%.2f", sumExonerada), false, true)
	writeCell(layout.totalInaW, fmt.Sprintf("%.2f", sumInafecta), false, true)
	writeCell(layout.iscW, fmt.Sprintf("%.2f", sumIsc), false, true)
	writeCell(layout.opBaseW, fmt.Sprintf("%.2f", sumBase), false, true)
	writeCell(layout.opIVAPW, fmt.Sprintf("%.2f", sumBaseIvap), false, true)
	writeCell(layout.icbW, fmt.Sprintf("%.2f", sumIcbper), false, true)
	writeCell(layout.otrosW, fmt.Sprintf("%.2f", sumOtros), false, true)
	writeCell(layout.impTotalW, fmt.Sprintf("%.2f", sumMtoTotalCp), false, true)
	writeCell(layout.tcW, "", false, false)
	writeCell(layout.refComFec, "", false, false)
	writeCell(layout.refComTip, "", false, false)
	writeCell(layout.refComSer, "", false, false)
	// Mover a la siguiente fila
	locationY += layout.rowTableH
	currentWriteW = layout.marginX + marginText
	// Escribir en la nueva fila con borde superior
	writeCell(layout.cuoW, "", true, false)
	writeCell(layout.cpeFecEmiW, "", true, false)
	writeCell(layout.cpeFecVenW, "", true, false)
	writeCell(layout.cpeTipoW, "", true, false)
	writeCell(layout.cpeSerieW, "", true, false)
	writeCell(layout.cpeNumW, "", true, false)
	writeCell(layout.cliDocTipoW, "", true, false)
	writeCell(layout.cliDocNumW, "", true, false)
	writeCell(layout.cliApeNomW, "", true, false)
	writeCell(layout.valFacOExpW, "0", true, false)
	writeCell(layout.baseImpW, "0", true, false)
	writeCell(layout.igvW, "0", true, false)
	writeCell(layout.totalExoW, "0", true, false)
	writeCell(layout.totalInaW, "0", true, false)
	writeCell(layout.iscW, "0", true, false)
	writeCell(layout.opBaseW, "0", true, false)
	writeCell(layout.opIVAPW, "0", true, false)
	writeCell(layout.icbW, "0", true, false)
	writeCell(layout.otrosW, "0", true, false)
	writeCell(layout.impTotalW, "0", true, false)
	writeCell(layout.tcW, "", true, false)
	writeCell(layout.refComFec, "", true, false)
	writeCell(layout.refComTip, "", true, false)
	writeCell(layout.refComSer, "", true, false)

	return err
}
func formatWithCommasAndDecimals(value float64) string {
    parts := strings.Split(fmt.Sprintf("%.2f", value), ".")
    integerPart := parts[0]
    decimalPart := parts[1]

    var result strings.Builder
    n := len(integerPart)
    for i, digit := range integerPart {
        if i > 0 && (n-i)%3 == 0 {
            result.WriteRune(',')
        }
        result.WriteRune(digit)
    }
    return result.String() + "." + decimalPart
}
func alignCenter(pdf *gopdf.GoPdf, value string, width float64, currentWriteW float64, rowMiddle float64, marginText float64) {
	textWidth, _ := pdf.MeasureTextWidth(value)
	pdf.SetXY(currentWriteW+(width-textWidth)/2 - marginText, rowMiddle)
	pdf.Text(value)
}
func alignRight(pdf *gopdf.GoPdf, value float32, width float64, currentWriteW float64, rowMiddle float64, marginText float64) {
    if value == 0 {
        return
    }
    text := formatWithCommasAndDecimals(float64(value))
    textWidth, _ := pdf.MeasureTextWidth(text)
    pdf.SetXY(currentWriteW+width-textWidth-marginText, rowMiddle)
    pdf.Text(text)
}
