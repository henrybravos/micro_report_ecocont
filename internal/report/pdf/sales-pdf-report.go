package pdf

import (
	"fmt"
	"github.com/henrybravo/micro-report/internal/repositories"
	"github.com/henrybravo/micro-report/pkg/files"
	v1 "github.com/henrybravo/micro-report/protos/gen/go/v1"
	"github.com/signintech/gopdf"
	"log"
	"time"
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

		headerPageH:  2.0,
		marginX:      0.5,
		marginY:      0.5,
		pageW:        29.7,
		pageH:        21.0,
		headerTableH: 1.5,
		rowTableH:    0.25,
		cuoW:         1.2,

		cpeInfoW:   4.3,
		cpeFecEmiW: 4.3 / 4,
		cpeFecVenW: 4.3 / 4,
		cpeTipoW:   4.3 / 6,
		cpeSerieW:  4.3 / 6,
		cpeNumW:    4.3 / 6,

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
		tcW:       0.5,

		refComW:   3.7,
		refComFec: 3.7 / 4,
		refComTip: 3.7 / 4,
		refComSer: 3.7 / 4,
		refComNum: 3.7 / 4,
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
	//calcular totales
	sumMtoValFactExpo := float32(0)
	sumBase := float32(0)
	sumIgv := float32(0)
	sumExonerada := float32(0)
	sumInafecta := float32(0)
	sumIsc := float32(0)
	sumBaseIvap := float32(0)
	sumIcbper := float32(0)
	sumOtros := float32(0)
	sumMtoTotalCp := float32(0)

	bandComings := false
	lensales := len(sales)
	for _, sale := range sales {
		lensales -= 1
		if pdf.GetY()+2*layout.rowTableH+2*layout.marginY > layout.pageH {
			err = generatePage(business, period, &pdf, layout)
			if err != nil {
				return "", err
			}
			bandComings = true
		}
		numCells := determineNumCells(sale)
		locationY := pdf.GetY() + layout.rowTableH*float64(numCells)

		if bandComings {
			//for comming true and for going false
			err = generateComingsAndGoings(&pdf, locationY, layout, true, sumMtoValFactExpo, sumBase, sumIgv, sumExonerada, sumInafecta, sumIsc, sumBaseIvap, sumIcbper, sumOtros, sumMtoTotalCp)
			if err != nil {
				return "", err
			}
			bandComings = false
			//locationY = pdf.GetY() + layout.rowTableH*float64(numCells)
		}
		err = generateRowTable(&pdf, sale, locationY, layout)

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
			err = generateComingsAndGoings(&pdf, locationY, layout, false, sumMtoValFactExpo, sumBase, sumIgv, sumExonerada, sumInafecta, sumIsc, sumBaseIvap, sumIcbper, sumOtros, sumMtoTotalCp)
		}
		//for total
		if lensales == 0 {
			if pdf.GetY()+layout.rowTableH+layout.marginY > layout.pageH {
				err = generatePage(business, period, &pdf, layout)
				if err != nil {
					return "", err
				}
			}
			err = addTotal(&pdf, locationY, layout, sumMtoValFactExpo, sumBase, sumIgv, sumExonerada, sumInafecta, sumIsc, sumBaseIvap, sumIcbper, sumOtros, sumMtoTotalCp)

			if err != nil {
				log.Print(err.Error())
				return
			}
		}
	}
	if err != nil {
		log.Print(err.Error())
		return
	}
	//add total, primero verifica que haya espacio para agregar el total

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
func determineNumCells(sale *v1.SalesReport) int {
	numCells := len(sale.RazonSocial)/50 + 1
	return numCells
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
	err = pdf.SetFont("arialB", "", 6)
	rect := &gopdf.Rect{
		H: 0.5,
		W: layout.pageW,
	}
	cellOptionCenter := gopdf.CellOption{
		Align: gopdf.Middle | gopdf.Center,
	}
	err = pdf.CellWithOption(rect, business.BusinessName, cellOptionCenter)
	pdf.SetXY(0, 0.75)
	err = pdf.CellWithOption(rect, "R.U.C.: "+business.RUC, cellOptionCenter)
	pdf.SetXY(0, 1.25)
	err = pdf.CellWithOption(rect, business.Address, cellOptionCenter)
	pdf.SetXY(0, 1.5)
	err = pdf.CellWithOption(rect, "REGISTRO DE VENTAS DEL MES DE "+period, cellOptionCenter)
	return
}
func generateHeaderTable(pdf *gopdf.GoPdf, layout layout) error {
	err := pdf.SetFont("arialB", "", 4.5)
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
		W: layout.cpeInfoW / 4,
	}
	err = pdf.CellWithOption(rect, "FECHA DE", cellOptionBorderRCenter)
	pdf.SetXY(pdf.GetX()-layout.cpeInfoW/4, pdf.GetY()+layout.headerTableH/4)
	err = pdf.CellWithOption(rect, "EMISIÓN", cellOptionBorderRBCenter)
	pdf.SetXY(pdf.GetX(), pdf.GetY()-layout.headerTableH/4)
	err = pdf.CellWithOption(rect, "FECHA DE", cellOptionCenter)
	pdf.SetXY(pdf.GetX()-layout.cpeInfoW/4, pdf.GetY()+layout.headerTableH/4)
	err = pdf.CellWithOption(rect, "VENCIMIENTO", cellOptionBorderRBCenter)
	pdf.SetXY(pdf.GetX(), pdf.GetY()-layout.headerTableH/4)
	rect = &gopdf.Rect{
		H: layout.headerTableH / 2,
		W: layout.cpeInfoW / 6,
	}
	err = pdf.CellWithOption(rect, "TIPO", cellOptionAllBorderCenter)
	err = pdf.CellWithOption(rect, "SERIE", cellOptionAllBorderCenter)
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
		H: layout.headerTableH / 4,
		W: layout.tcW,
	}
	err = pdf.CellWithOption(rect, "TIPO", cellOptionBorderRTCenter)
	pdf.SetXY(pdf.GetX()-layout.tcW, pdf.GetY()+layout.headerTableH/4)
	err = pdf.CellWithOption(rect, "DE", cellOptionBorderRCenter)
	pdf.SetXY(pdf.GetX()-layout.tcW, pdf.GetY()+layout.headerTableH/4)
	err = pdf.CellWithOption(rect, "CAM", cellOptionBorderRCenter)
	pdf.SetXY(pdf.GetX()-layout.tcW, pdf.GetY()+layout.headerTableH/4)
	err = pdf.CellWithOption(rect, "BIO", cellOptionBorderRBCenter)
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
		W: layout.refComW / 4,
	}
	err = pdf.CellWithOption(rect, "FECHA", cellOptionAllBorderCenter)
	err = pdf.CellWithOption(rect, "TIPO", cellOptionAllBorderCenter)
	err = pdf.CellWithOption(rect, "SERIE", cellOptionAllBorderCenter)
	err = pdf.CellWithOption(rect, "NÚMERO", cellOptionAllBorderCenter)
	return err
}
func generateRowTable(pdf *gopdf.GoPdf, sale *v1.SalesReport, locationY float64, layout layout) error {
	err := pdf.SetFont("arial", "", 4.5)
	rowMiddle := locationY + 2*layout.rowTableH/3
	rowW := layout.pageW - layout.marginX*2
	marginText := 0.075
	currentWriteW := layout.marginX + marginText
	pdf.SetXY(layout.marginX, locationY)
	pdf.SetStrokeColor(222, 219, 218)
	pdf.SetLineWidth(0)
	cellOptionBottom := gopdf.CellOption{
		Border: gopdf.Bottom,
		Align:  gopdf.Left | gopdf.Middle,
	}
	rect := &gopdf.Rect{
		H: layout.rowTableH,
		W: rowW,
	}
	err = pdf.CellWithOption(rect, "", cellOptionBottom)
	pdf.SetXY(currentWriteW, rowMiddle)
	err = pdf.Text(sale.Cuo)
	currentWriteW += layout.cuoW
	pdf.SetXY(currentWriteW, rowMiddle)
	err = pdf.Text(sale.FecEmision)
	currentWriteW += layout.cpeFecEmiW
	pdf.SetXY(currentWriteW, rowMiddle)
	err = pdf.Text(sale.FechaVencimiento)
	currentWriteW += layout.cpeFecVenW
	pdf.SetXY(currentWriteW, rowMiddle)
	err = pdf.Text(sale.CodigoTipoCdp)
	currentWriteW += layout.cpeTipoW
	pdf.SetXY(currentWriteW, rowMiddle)
	err = pdf.Text(sale.NumSerieCdp)
	currentWriteW += layout.cpeSerieW
	pdf.SetXY(currentWriteW, rowMiddle)
	err = pdf.Text(sale.NumCdp)
	currentWriteW += layout.cpeNumW
	pdf.SetXY(currentWriteW, rowMiddle)
	err = pdf.Text("  " + sale.CodTipoDocIdentidad)
	currentWriteW += layout.cliDocTipoW
	pdf.SetXY(currentWriteW, rowMiddle)
	err = pdf.Text(sale.NumDocIdentidadClient)
	currentWriteW += layout.cliDocNumW
	pdf.SetXY(currentWriteW, rowMiddle)
	lenRazon := len(sale.RazonSocial)
	if lenRazon > 50 {
		pdf.SetXY(currentWriteW, rowMiddle-0.25)
		rect := &gopdf.Rect{
			H: layout.rowTableH * 3,
			W: layout.cliApeNomW,
		}
		err = pdf.MultiCell(rect, sale.RazonSocial)
	} else {
		err = pdf.Text(sale.RazonSocial)
	}
	currentWriteW += layout.cliApeNomW
	pdf.SetXY(currentWriteW, rowMiddle)
	if sale.MtoValFactExpo == 0 {
		err = pdf.Text("")
	} else {
		err = pdf.Text(fmt.Sprintf("%.2f", sale.MtoValFactExpo))
	}
	currentWriteW += layout.valFacOExpW
	pdf.SetXY(currentWriteW, rowMiddle)
	if sale.Base == 0 {
		err = pdf.Text("")
	} else {
		err = pdf.Text(fmt.Sprintf("%.2f", sale.Base))
	}
	currentWriteW += layout.baseImpW
	pdf.SetXY(currentWriteW, rowMiddle)
	if sale.Igv == 0 {
		err = pdf.Text("")
	} else {
		err = pdf.Text(fmt.Sprintf("%.2f", sale.Igv))
	}
	currentWriteW += layout.igvW
	pdf.SetXY(currentWriteW, rowMiddle)
	if sale.Exonerada == 0 {
		err = pdf.Text("")
	} else {
		err = pdf.Text(fmt.Sprintf("%.2f", sale.Exonerada))
	}
	currentWriteW += layout.totalExoW
	pdf.SetXY(currentWriteW, rowMiddle)
	if sale.Inafecta == 0 {
		err = pdf.Text("")
	} else {
		err = pdf.Text(fmt.Sprintf("%.2f", sale.Inafecta))
	}
	currentWriteW += layout.totalInaW
	pdf.SetXY(currentWriteW, rowMiddle)
	if sale.Isc == 0 {
		err = pdf.Text("")
	} else {
		err = pdf.Text(fmt.Sprintf("%.2f", sale.Isc))
	}
	currentWriteW += layout.iscW
	pdf.SetXY(currentWriteW, rowMiddle)
	if sale.Base == 0 {
		err = pdf.Text("")
	} else {
		err = pdf.Text(fmt.Sprintf("%.2f", sale.Base))
	}
	currentWriteW += layout.opBaseW
	pdf.SetXY(currentWriteW, rowMiddle)
	if sale.BaseIvap == 0 {
		err = pdf.Text("")
	} else {
		err = pdf.Text(fmt.Sprintf("%.2f", sale.BaseIvap))
	}
	currentWriteW += layout.opIVAPW
	pdf.SetXY(currentWriteW, rowMiddle)
	if sale.Icbper == 0 {
		err = pdf.Text("")
	} else {
		err = pdf.Text(fmt.Sprintf("%.2f", sale.Icbper))
	}
	currentWriteW += layout.icbW
	pdf.SetXY(currentWriteW, rowMiddle)
	if sale.Otros == 0 {
		err = pdf.Text("")
	} else {
		err = pdf.Text(fmt.Sprintf("%.2f", sale.Otros))
	}
	currentWriteW += layout.otrosW
	pdf.SetXY(currentWriteW, rowMiddle)
	if sale.MtoTotalCp == 0 {
		err = pdf.Text("")
	} else {
		err = pdf.Text(fmt.Sprintf("%.2f", sale.MtoTotalCp))
	}
	currentWriteW += layout.impTotalW
	pdf.SetXY(currentWriteW, rowMiddle)
	if sale.TipoCambio == 1 {
		err = pdf.Text("")
	} else {
		err = pdf.Text(fmt.Sprintf("%.2f", sale.TipoCambio))
	}
	currentWriteW += layout.tcW
	pdf.SetXY(currentWriteW, rowMiddle)
	err = pdf.Text(sale.FecEmisionMod)
	currentWriteW += layout.refComFec
	pdf.SetXY(currentWriteW, rowMiddle)
	err = pdf.Text(sale.CodigoTipoCdpMod)
	currentWriteW += layout.refComTip
	pdf.SetXY(currentWriteW, rowMiddle)
	err = pdf.Text(sale.NumSerieCdpMod)
	currentWriteW += layout.refComSer
	pdf.SetXY(currentWriteW, rowMiddle)
	err = pdf.Text(sale.NumCdpMod)
	return err
}

func generateComingsAndGoings(pdf *gopdf.GoPdf, locationY float64, layout layout, goingComing bool,
	sumMtoValFactExpo float32, sumBase float32,
	sumIgv float32, sumExonerada float32, sumInafecta float32,
	sumIsc float32, sumBaseIvap float32, sumIcbper float32, sumOtros float32, sumMtoTotalCp float32) error {
	if !goingComing {
		locationY += layout.rowTableH
	} else {
		locationY -= layout.rowTableH
	}
	err := pdf.SetFont("arialB", "", 4.5)
	if err != nil {
		log.Print(err.Error())
	}
	rowMiddle := locationY + 2*layout.rowTableH/3
	marginText := 0.075
	currentWriteW := layout.marginX + marginText

	// Función para escribir una celda con borde
	writeCell := func(width float64, text string) {
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
		pdf.SetXY(currentWriteW, rowMiddle)
		err = pdf.Text(text)
		if err != nil {
			log.Fatalf("Error writing text: %v", err)
		}
		currentWriteW += width
	}
	pdf.SetLineWidth(0)

	// Escribir celdas con borde
	// currentWriteW, err = writeCell(pdf, currentWriteW, locationY, rowMiddle, layout.cuoW, "", goingComing, layout)
	// currentWriteW, err = writeCell(pdf, currentWriteW, locationY, rowMiddle, layout.cpeFecEmiW, "", goingComing, layout)
	// currentWriteW, err = writeCell(pdf, currentWriteW, locationY, rowMiddle, layout.cpeFecVenW, "", goingComing, layout)
	// currentWriteW, err = writeCell(pdf, currentWriteW, locationY, rowMiddle, layout.cpeTipoW, "", goingComing, layout)
	// currentWriteW, err = writeCell(pdf, currentWriteW, locationY, rowMiddle, layout.cpeSerieW, "", goingComing, layout)
	// currentWriteW, err = writeCell(pdf, currentWriteW, locationY, rowMiddle, layout.cpeNumW, "", goingComing, layout)
	// currentWriteW, err = writeCell(pdf, currentWriteW, locationY, rowMiddle, layout.cliDocTipoW, "", goingComing, layout)
	// currentWriteW, err = writeCell(pdf, currentWriteW, locationY, rowMiddle, layout.cliDocNumW, "", goingComing, layout)
	// currentWriteW, err = writeCell(pdf, currentWriteW, locationY, rowMiddle, layout.cliApeNomW, func() string {
	// 	if !goingComing {
	// 		return "VAN"
	// 	}
	// 	return "VIENEN"
	// }(), goingComing, layout)
	// currentWriteW, err = writeCell(pdf, currentWriteW, locationY, rowMiddle, layout.valFacOExpW, fmt.Sprintf("%.2f", sumMtoValFactExpo), goingComing, layout)
	// currentWriteW, err = writeCell(pdf, currentWriteW, locationY, rowMiddle, layout.baseImpW, fmt.Sprintf("%.2f", sumBase), goingComing, layout)
	// currentWriteW, err = writeCell(pdf, currentWriteW, locationY, rowMiddle, layout.igvW, fmt.Sprintf("%.2f", sumIgv), goingComing, layout)
	// currentWriteW, err = writeCell(pdf, currentWriteW, locationY, rowMiddle, layout.totalExoW, fmt.Sprintf("%.2f", sumExonerada), goingComing, layout)
	// currentWriteW, err = writeCell(pdf, currentWriteW, locationY, rowMiddle, layout.totalInaW, fmt.Sprintf("%.2f", sumInafecta), goingComing, layout)
	// currentWriteW, err = writeCell(pdf, currentWriteW, locationY, rowMiddle, layout.iscW, fmt.Sprintf("%.2f", sumIsc), goingComing, layout)
	// currentWriteW, err = writeCell(pdf, currentWriteW, locationY, rowMiddle, layout.opBaseW, fmt.Sprintf("%.2f", sumBase), goingComing, layout)
	// currentWriteW, err = writeCell(pdf, currentWriteW, locationY, rowMiddle, layout.opIVAPW, fmt.Sprintf("%.2f", sumBaseIvap), goingComing, layout)
	// currentWriteW, err = writeCell(pdf, currentWriteW, locationY, rowMiddle, layout.icbW, fmt.Sprintf("%.2f", sumIcbper), goingComing, layout)
	// currentWriteW, err = writeCell(pdf, currentWriteW, locationY, rowMiddle, layout.otrosW, fmt.Sprintf("%.2f", sumOtros), goingComing, layout)
	// currentWriteW, err = writeCell(pdf, currentWriteW, locationY, rowMiddle, layout.impTotalW, fmt.Sprintf("%.2f", sumMtoTotalCp), goingComing, layout)
	// currentWriteW, err = writeCell(pdf, currentWriteW, locationY, rowMiddle, layout.tcW, "", goingComing, layout)
	// currentWriteW, err = writeCell(pdf, currentWriteW, locationY, rowMiddle, layout.refComFec, "", goingComing, layout)
	// currentWriteW, err = writeCell(pdf, currentWriteW, locationY, rowMiddle, layout.refComTip, "", goingComing, layout)
	// currentWriteW, err = writeCell(pdf, currentWriteW, locationY, rowMiddle, layout.refComSer, "", goingComing, layout)

	writeCell(layout.cuoW, "")
	writeCell(layout.cpeFecEmiW, "")
	writeCell(layout.cpeFecVenW, "")
	writeCell(layout.cpeTipoW, "")
	writeCell(layout.cpeSerieW, "")
	writeCell(layout.cpeNumW, "")
	writeCell(layout.cliDocTipoW, "")
	writeCell(layout.cliDocNumW, "")
	writeCell(layout.cliApeNomW, func() string {
		if !goingComing {
			return "VAN"
		}
		return "VIENEN"
	}())
	writeCell(layout.valFacOExpW, fmt.Sprintf("%.2f", sumMtoValFactExpo))
	writeCell(layout.baseImpW, fmt.Sprintf("%.2f", sumBase))
	writeCell(layout.igvW, fmt.Sprintf("%.2f", sumIgv))
	writeCell(layout.totalExoW, fmt.Sprintf("%.2f", sumExonerada))
	writeCell(layout.totalInaW, fmt.Sprintf("%.2f", sumInafecta))
	writeCell(layout.iscW, fmt.Sprintf("%.2f", sumIsc))
	writeCell(layout.opBaseW, fmt.Sprintf("%.2f", sumBase))
	writeCell(layout.opIVAPW, fmt.Sprintf("%.2f", sumBaseIvap))
	writeCell(layout.icbW, fmt.Sprintf("%.2f", sumIcbper))
	writeCell(layout.otrosW, fmt.Sprintf("%.2f", sumOtros))
	writeCell(layout.impTotalW, fmt.Sprintf("%.2f", sumMtoTotalCp))
	writeCell(layout.tcW, "")
	writeCell(layout.refComFec, "")
	writeCell(layout.refComTip, "")
	writeCell(layout.refComSer, "")

	return err
}

func addTotal(pdf *gopdf.GoPdf, locationY float64, layout layout, sumMtoValFactExpo float32, sumBase float32,
	sumIgv float32, sumExonerada float32, sumInafecta float32, sumIsc float32, sumBaseIvap float32,
	sumIcbper float32, sumOtros float32, sumMtoTotalCp float32) error {
	err := pdf.SetFont("arialB", "", 4.5)
	if err != nil {
		log.Print(err.Error())
	}

	rowMiddle := locationY + 2*layout.rowTableH
	marginText := 0.075
	currentWriteW := layout.marginX 

	writeCell := func(width float64, text string, borderTop bool) {
		pdf.SetXY(currentWriteW, locationY)
		cellOption := gopdf.CellOption{
			Align: gopdf.Left | gopdf.Middle,
		}

		if !borderTop && text != "" && text != "TOTAL" {
			pdf.SetStrokeColor(50, 50, 50)
			cellOption.Border = gopdf.Bottom
		}
		if borderTop && text != "" {
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
		pdf.SetXY(currentWriteW, rowMiddle)
		if !borderTop{
			err = pdf.Text(text)
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
	writeCell(layout.cliApeNomW, "TOTAL", false)
	writeCell(layout.valFacOExpW, fmt.Sprintf("%.2f", sumMtoValFactExpo), false)
	writeCell(layout.baseImpW, fmt.Sprintf("%.2f", sumBase), false)
	writeCell(layout.igvW, fmt.Sprintf("%.2f", sumIgv), false)
	writeCell(layout.totalExoW, fmt.Sprintf("%.2f", sumExonerada), false)
	writeCell(layout.totalInaW, fmt.Sprintf("%.2f", sumInafecta), false)
	writeCell(layout.iscW, fmt.Sprintf("%.2f", sumIsc), false)
	writeCell(layout.opBaseW, fmt.Sprintf("%.2f", sumBase), false)
	writeCell(layout.opIVAPW, fmt.Sprintf("%.2f", sumBaseIvap), false)
	writeCell(layout.icbW, fmt.Sprintf("%.2f", sumIcbper), false)
	writeCell(layout.otrosW, fmt.Sprintf("%.2f", sumOtros), false)
	writeCell(layout.impTotalW, fmt.Sprintf("%.2f", sumMtoTotalCp), false)
	writeCell(layout.tcW, "", false)
	writeCell(layout.refComFec, "", false)
	writeCell(layout.refComTip, "", false)
	writeCell(layout.refComSer, "", false)

	// Mover a la siguiente fila
	locationY += layout.rowTableH
	currentWriteW = layout.marginX + marginText
	// Escribir en la nueva fila con borde superior
	writeCell(layout.cuoW, "", true)
	writeCell(layout.cpeFecEmiW, "", true)
	writeCell(layout.cpeFecVenW, "", true)
	writeCell(layout.cpeTipoW, "", true)
	writeCell(layout.cpeSerieW, "", true)
	writeCell(layout.cpeNumW, "", true)
	writeCell(layout.cliDocTipoW, "", true)
	writeCell(layout.cliDocNumW, "", true)
	writeCell(layout.cliApeNomW, "", true)
	writeCell(layout.valFacOExpW, "0", true)
	writeCell(layout.baseImpW, "0", true)
	writeCell(layout.igvW, "0", true)
	writeCell(layout.totalExoW, "0", true)
	writeCell(layout.totalInaW, "0", true)
	writeCell(layout.iscW, "0", true)
	writeCell(layout.opBaseW, "0", true)
	writeCell(layout.opIVAPW, "0", true)
	writeCell(layout.icbW, "0", true)
	writeCell(layout.otrosW, "0", true)
	writeCell(layout.impTotalW, "0", true)
	writeCell(layout.tcW, "", true)
	writeCell(layout.refComFec, "", true)
	writeCell(layout.refComTip, "", true)
	writeCell(layout.refComSer, "", true)

	return err
}
