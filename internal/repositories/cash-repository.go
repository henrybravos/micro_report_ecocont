package repositories

import (
	"context"
	"github.com/henrybravo/micro-report/pkg/db"
	v1 "github.com/henrybravo/micro-report/protos/gen/go/v1"
)

type CashRepository struct {
	Connection *db.Connection
}

func NewCashRepository(connection *db.Connection) *CashRepository {
	return &CashRepository{Connection: connection}
}

func (r *CashRepository) GetCashBalance(businessId, year, month string, accountsIDs []string) (accounts []*v1.AccountBalance, err error) {
	query := `
			SELECT c.cuenta_financiera_id, 
				   SUM((COALESCE(c.debe, 0) - COALESCE(c.haber, 0)) * 
					   COALESCE(NULLIF(pg.tipo_cambio, 1), o.tipo_cambio)) AS total
			FROM operaciones o
			INNER JOIN cajas c ON o.id = c.operacion_id
			LEFT JOIN pagos pg ON c.referencia_id = pg.id
			INNER JOIN locales l ON o.local_id = l.id
			WHERE l.empresa_id = $1
			  AND c.periodo <= $2 || '-' || $3
			  AND c.cuenta_financiera_id = ANY($4)
			  AND o.comprobante_id::text <> ALL(ARRAY['a6062ae0-15a4-11ec-8fec-77a5f80a0a28', '1daedb70-a779-11eb-84c1-40b0344a6892'])
			  AND c.deleted_at IS NULL
			  AND o.deleted_at IS NULL
			  AND o.estado_le = '1'
			GROUP BY c.cuenta_financiera_id`
	rows, err := r.Connection.Pool.Query(context.Background(), query, businessId, year, month, accountsIDs)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var account v1.AccountBalance
		err = rows.Scan(&account.AccountId, &account.Amount)
		if err != nil {
			return
		}
		accounts = append(accounts, &account)
	}
	return
}

func (r *CashRepository) GetLCash(businessId, year, month string, accountsIDs []string) (cashes []*v1.LCash, err error) {
	query := `
		SELECT c.id, --0
			   c.periodo,
			   CONCAT(o.cuo, '-', c.tipo),
			   c.identificador,
			   pc.codigo,
			   pc.denominacion,
			   COALESCE(c.centro_costo_id::text,''), --6
			   tm.codigo                 ,
			   tc.codigo                 ,
			   COALESCE(o.serie, '0000') ,
			   COALESCE(o.correlativo, '0'), --10
			   COALESCE(o.fecha_contable::text, '')          , --11
			   COALESCE(o.fecha_vencimiento::text, '')       , --12
			   COALESCE(o.fecha_emision::text, '')           , --13
			   COALESCE(pg.fecha::text, '')                 , --14
			   COALESCE(pg.glosa::text, '')               ,
			   o.glosa                   ,
			   COALESCE(o.glosa_referencia,'')        ,	--17
			   COALESCE(c.debe,0)                    ,
			   COALESCE(c.haber,0)                   , --19
			   o.dato_estructurado       ,
			   c.estado                  , --21
			   o.estado_le               ,	--22
			   o.observaciones           , --23
			   o.tipo_cambio             ,
			   o.codigo_libro            , --25
			   o.periodo                 , --26
			   COALESCE(pg.tipo_cambio ,0)           , --27
			   COALESCE(COALESCE(pg.fecha::text, o.fecha_emision::text),'') AS "i3oo" --28
		FROM operaciones o
		INNER JOIN cajas c ON o.id = c.operacion_id
		LEFT JOIN pagos pg ON c.referencia_id = pg.id
		INNER JOIN t_monedas tm ON o.moneda_id = tm.id
		LEFT JOIN t_comprobantes tc ON o.comprobante_id = tc.id
		INNER JOIN pcge pc ON c.pcge_id = pc.id
		INNER JOIN locales l ON o.local_id = l.id
		WHERE l.empresa_id = $1
		  AND SUBSTRING(c.periodo, 1, 4) = $2
		  AND SUBSTRING(c.periodo, 6, 8) = $3 
		  AND c.cuenta_financiera_id::text = ANY($4)
		  AND o.comprobante_id::text <> ALL(ARRAY['a6062ae0-15a4-11ec-8fec-77a5f80a0a28', '1daedb70-a779-11eb-84c1-40b0344a6892'])
		  AND c.deleted_at IS NULL
		  AND o.deleted_at IS NULL
		ORDER BY "i3oo", o.cuo`
	rows, err := r.Connection.Pool.Query(context.Background(), query, businessId, year, month, accountsIDs)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var lCash v1.LCash
		err = rows.Scan(
			&lCash.CId,
			&lCash.CPeriodo,
			&lCash.OCuoTipo,
			&lCash.CIdentificador,
			&lCash.PcCodigo,
			&lCash.PcDenominacion,
			&lCash.CCentroCostoId,
			&lCash.TmCodigo,
			&lCash.TcCodigo,
			&lCash.OSerie,
			&lCash.OCorrelativo,
			&lCash.OFechaContable,
			&lCash.OFechaVencimiento,
			&lCash.OFechaEmision,
			&lCash.PgFecha,
			&lCash.PgGlosa,
			&lCash.OGlosa,
			&lCash.OGlosaReferencia,
			&lCash.CDebe,
			&lCash.CHaber,
			&lCash.ODatoEstructurado,
			&lCash.CEstado,
			&lCash.OEstadoLe,
			&lCash.OObservaciones,
			&lCash.OTipoCambio,
			&lCash.OCodigoLibro,
			&lCash.OPeriodo,
			&lCash.PgTipoCambio,
			&lCash.PgFechaOFechaEmision,
		)
		if err != nil {
			return
		}
		cashes = append(cashes, &lCash)
	}
	return
}
