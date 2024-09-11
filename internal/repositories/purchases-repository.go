package repositories

import (
	"context"
	"fmt"
	"github.com/henrybravo/micro-report/pkg/db"
	v1 "github.com/henrybravo/micro-report/protos/gen/go/v1"
)

type PurchaseRepository struct {
	Connection *db.Connection
}

func NewPurchaseRepository(connection *db.Connection) *PurchaseRepository {
	return &PurchaseRepository{Connection: connection}
}
func (p *PurchaseRepository) GetPurchasesByBusinessID(businessID, period string) ([]*v1.PurchaseReport, error) {
	query := `
		SELECT 
			o.id as id, --0
			o.periodo as i1, --1
			o.cuo as i2, --2
			c.identificador_linea as i3, --3
			o.fecha_emision::text as i4, --4
			COALESCE(o.fecha_vencimiento::text, '') as i5, --5
			tco.codigo as codTipoCDP,
			o.serie as numSerieCDP,
			COALESCE(c.dua, '') as annCDP, --8
			o.correlativo as numCDP,
			COALESCE(c.numero_final, '') as numCDPRangoFinal,
			td.codigo as codTipoDocIdentidadProveedor,
			p.numero as numDocIdentidadProveedor,
			p.razon_social as nomRazonSocialProveedor,
			c.base_1 as mtoBIGravadaDG,
			c.igv_1 as mtoIgvIpmDG,
			c.base_2 as mtoBIGravadaDGNG,
			c.igv_2 as mtoIgvIpmDGNG,
			c.base_3 as mtoBIGravadaDNG,
			c.igv_3 as mtoIgvIpmDNG,
			c.no_gravada as mtoValorAdqNG,
			c.isc as mtoISC,
			c.otros as mtoOtrosTrib,
			c.total as mtoTotalCp,
			tm.codigo as codMoneda,
			o.tipo_cambio as mtoTipoCambio,
			COALESCE(c.fecha_cdpm::text, '') as fecEmisionMod, --26
			COALESCE(tico.codigo, '') as codTipoCDPMod, --27
			COALESCE(c.serie_cdpm, '') as numSerieCDPMod,
			COALESCE(c.dua_cdpm, '') as codDepAduanera,
			COALESCE(c.correlativo_cdpm, '') as numCDPMod,
			COALESCE(c.fecha_detraccion::text, '')    as i31,
			COALESCE(c.numero_detraccion, '')   as i32,
       		COALESCE(c.marca_cdp_retencion, '') as i33,
       		COALESCE(bbss.codigo, '')          as i34,
       		COALESCE(c.contrato, '')            as i35,
 			COALESCE(c.error_1, '')             as i36,
       		COALESCE(c.error_2 , '')            as i37,
       		COALESCE(c.error_3, '')             as i38,
       		COALESCE(c.error_4,'')             as i39,
            COALESCE(c.medio_pago, '')         as i40, --40
       		COALESCE(o.importe_cdp_regimen_sunat, 0)  as itotal,
       		COALESCE(o.tipo_cdp_regimen_sunat::text, '')    as afectacion, --42
			c.icbper as mtoIcbp,
			c.estado as codEstadoComprobante
		FROM operaciones o
			INNER JOIN compras c ON o.id = c.operacion_id
			INNER JOIN personas p ON p.id = o.persona_asociado_id
			INNER JOIN t_comprobantes tco ON tco.id = o.comprobante_id
			INNER JOIN t_monedas tm ON tm.id = o.moneda_id
			INNER JOIN t_documentos td ON td.id = p.documento_id
			LEFT JOIN t_comprobantes tico ON tico.id = c.comprobante_cdpm_id
			LEFT JOIN t_bbss_adquiridos bbss ON c.bbss_id = bbss.id
			INNER JOIN locales l ON l.id = o.local_id
		WHERE o.periodo = $2 
			AND l.empresa_id = $1
			AND o.cuo IS NOT NULL
			AND o.tipo_operacion::text <> ALL(ARRAY[
				'59133fc4-a77a-11eb-8911-40b0344a6892', 
				'59136636-a77a-11eb-891e-40b0344a6892',
				'59136639-a77a-11eb-8921-40b0344a6892', 
				'5913663a-a77a-11eb-8922-40b0344a6892', 
				'5913663b-a77a-11eb-8923-40b0344a6892', 
				'5913663c-a77a-11eb-8924-40b0344a6892',
				'59138d1d-a77a-11eb-8928-40b0344a6892', 
				'59138d1f-a77a-11eb-892a-40b0344a6892', 
				'59138d20-a77a-11eb-892b-40b0344a6892',
				'59138d22-a77a-11eb-892d-40b0344a6892', 
				'59136632-a77a-11eb-891a-40b0344a6892'
			])
			AND o.comprobante_id::text <> ALL(ARRAY[
				'a6062ae0-15a4-11ec-8fec-77a5f80a0a28', 
				'1daedb70-a779-11eb-84c1-40b0344a6892', 
				'1daeb476-a779-11eb-84b8-40b0344a6892', 
				'1daf9e37-a779-11eb-84f1-40b0344a6892', 
				'1daf9e39-a779-11eb-84f3-40b0344a6892', 
				'c384c63a-31f1-11ec-9c4c-2b2d93307b9f'
			])
			AND o.codigo_libro = '080100'
			AND o.deleted_at IS NULL 
			AND o.referencia_id IS NULL
		ORDER BY o.fecha_emision, o.cuo ASC;
	`
	rows, err := p.Connection.Pool.Query(context.Background(), query, businessID, period)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()
	var purchases []*v1.PurchaseReport
	for rows.Next() {
		var purchase v1.PurchaseReport
		err := rows.Scan(
			&purchase.OId,
			&purchase.OPeriodo,
			&purchase.OCuo,
			&purchase.CIdentificadorLinea,
			&purchase.OFechaEmision,
			&purchase.OFechaVencimiento,
			&purchase.TcoCodigo,
			&purchase.OSerie,
			&purchase.CDua,
			&purchase.OCorrelativo,
			&purchase.CNumeroFinal,
			&purchase.TdCodigo,
			&purchase.PNumero,
			&purchase.PRazonSocial,
			&purchase.CBase_1,
			&purchase.CIgv_1,
			&purchase.CBase_2,
			&purchase.CIgv_2,
			&purchase.CBase_3,
			&purchase.CIgv_3,
			&purchase.CNoGravada,
			&purchase.CIsc,
			&purchase.COtros,
			&purchase.CTotal,
			&purchase.TmCodigo,
			&purchase.OTipoCambio,
			&purchase.CFechaCdpm,
			&purchase.TicoCodigo,
			&purchase.CSerieCdpm,
			&purchase.CDuaCdpm,
			&purchase.CCorrelativoCdpm,
			&purchase.CFechaDetraccion,
			&purchase.CNumeroDetraccion,
			&purchase.CMarcaCdpRetencion,
			&purchase.BbssCodigo,
			&purchase.CContrato,
			&purchase.CError_1,
			&purchase.CError_2,
			&purchase.CError_3,
			&purchase.CError_4,
			&purchase.CMedioPago,
			&purchase.OImporteCdpRegimenSunat,
			&purchase.OTipoCdpRegimenSunat,
			&purchase.CIcbper,
			&purchase.CEstado,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		purchases = append(purchases, &purchase)

	}
	return purchases, nil
}
