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
	for _, sale := range sales {
		if pdf.GetY()+layout.rowTableH+layout.marginY > layout.pageH {
			err = generatePage(business, period, &pdf, layout)
		}
		locationY := pdf.GetY() + layout.rowTableH
		err = generateRowTable(&pdf, sale, locationY, layout)
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
	if lenRazon > 40 {
		pdf.SetXY(currentWriteW, rowMiddle-0.25)
		err = pdf.SetFont("arial", "", 3)
		rect = &gopdf.Rect{
			H: layout.rowTableH * 3,
			W: layout.cliApeNomW,
		}
		err = pdf.MultiCell(rect, sale.RazonSocial)
		err = pdf.SetFont("arial", "", 4.5)
	} else {
		err = pdf.Text(sale.RazonSocial)
	}
	currentWriteW += layout.cliApeNomW
	pdf.SetXY(currentWriteW, rowMiddle)
	err = pdf.Text(fmt.Sprintf("%.2f", sale.MtoValFactExpo))
	currentWriteW += layout.valFacOExpW
	pdf.SetXY(currentWriteW, rowMiddle)
	err = pdf.Text(fmt.Sprintf("%.2f", sale.Base))
	currentWriteW += layout.baseImpW
	pdf.SetXY(currentWriteW, rowMiddle)
	err = pdf.Text(fmt.Sprintf("%.2f", sale.Igv))
	currentWriteW += layout.igvW
	pdf.SetXY(currentWriteW, rowMiddle)
	err = pdf.Text(fmt.Sprintf("%.2f", sale.Exonerada))
	currentWriteW += layout.totalExoW
	pdf.SetXY(currentWriteW, rowMiddle)
	err = pdf.Text(fmt.Sprintf("%.2f", sale.Inafecta))
	currentWriteW += layout.totalInaW
	pdf.SetXY(currentWriteW, rowMiddle)
	err = pdf.Text(fmt.Sprintf("%.2f", sale.Isc))
	currentWriteW += layout.iscW
	pdf.SetXY(currentWriteW, rowMiddle)
	err = pdf.Text(fmt.Sprintf("%.2f", sale.Base))
	currentWriteW += layout.opBaseW
	pdf.SetXY(currentWriteW, rowMiddle)
	err = pdf.Text(fmt.Sprintf("%.2f", sale.BaseIvap))
	currentWriteW += layout.opIVAPW
	pdf.SetXY(currentWriteW, rowMiddle)
	err = pdf.Text(fmt.Sprintf("%.2f", sale.Icbper))
	currentWriteW += layout.icbW
	pdf.SetXY(currentWriteW, rowMiddle)
	err = pdf.Text(fmt.Sprintf("%.2f", sale.Otros))
	currentWriteW += layout.otrosW
	pdf.SetXY(currentWriteW, rowMiddle)
	err = pdf.Text(fmt.Sprintf("%.2f", sale.MtoTotalCp))
	currentWriteW += layout.impTotalW
	pdf.SetXY(currentWriteW, rowMiddle)
	err = pdf.Text(fmt.Sprintf("%.2f", sale.TipoCambio))
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
