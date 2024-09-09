package repositories

import (
	"context"
	"fmt"
	"github.com/henrybravo/micro-report/pkg/db"
	v1 "github.com/henrybravo/micro-report/protos/gen/go/v1"
)

type BankBookRepository struct {
	Connection *db.Connection
}

func NewBankBookRepository(connection *db.Connection) *BankBookRepository {
	return &BankBookRepository{Connection: connection}
}

func (b *BankBookRepository) GetLBankBalance(businessID string, year string, month string, financialAccountIDs []string) (banks []*v1.BankBalance, err error) {
	query := `
		SELECT b.cuenta_bancaria_id, 
		       SUM((COALESCE(b.debe, 0) - COALESCE(b.haber, 0)) * 
		          COALESCE(NULLIF(pg.tipo_cambio, 1), o.tipo_cambio)) AS total
		FROM operaciones o
		INNER JOIN bancos b ON o.id = b.operacion_id
		LEFT JOIN pagos pg ON b.referencia_id = pg.id
		INNER JOIN locales l ON o.local_id = l.id
		WHERE l.empresa_id = $1
          AND SUBSTRING(b.periodo, 1, 4) = $2
		  AND SUBSTRING(b.periodo, 6, 8) = $3
		  AND b.cuenta_bancaria_id::text = ANY($4)
		  AND o.comprobante_id::text <> ALL(ARRAY['a6062ae0-15a4-11ec-8fec-77a5f80a0a28', '1daedb70-a779-11eb-84c1-40b0344a6892'])
		  AND b.deleted_at IS NULL
		  AND o.deleted_at IS NULL
		  AND o.estado_le = '1'
		GROUP BY b.cuenta_bancaria_id`

	rows, err := b.Connection.Pool.Query(context.Background(), query, businessID, year, month, financialAccountIDs)
	if err != nil {
		return nil, fmt.Errorf("query failed: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var bank v1.BankBalance
		if err := rows.Scan(&bank.Id, &bank.Amount); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		banks = append(banks, &bank)
	}
	return
}
func (b *BankBookRepository) GetLBanks(businessID string, year string, month string, financialAccountIDs []string) (banks []*v1.LBank, err error) {
	query := `
		SELECT b.id, --0
		       b.periodo,
		       concat(o.cuo, '-', b.tipo),
		       b.identificador,
		       teb.codigo,
		       cf.cuenta,
		       COALESCE(o.fecha_emision::text,''), --6
		       COALESCE(pg.fecha::text, ''), --7
		       tmp.codigo,
		       COALESCE(pg.glosa,''),
		       o.glosa,
		       COALESCE(td.codigo,''), -- 11
		       COALESCE(p.numero,''),
		       COALESCE(p.razon_social,''), --13
		       b.numero_transaccion,
		       COALESCE(b.debe,0), --15
		       COALESCE(b.haber,0),
		       b.estado, --17
		       pc.codigo,
		       pc.denominacion, --19
		       o.estado_le, --20
		       o.observaciones,
		       o.tipo_cambio,
		       COALESCE(pg.tipo_cambio, 0), --23
		      CASE
		        WHEN pg.fecha IS NOT NULL THEN  COALESCE(pg.fecha::text,'')
		        ELSE COALESCE(o.fecha_emision::text,'')
		       END as i6oo
		FROM operaciones o
		INNER JOIN bancos b ON o.id = b.operacion_id
		LEFT JOIN pagos pg ON b.referencia_id = pg.id
		INNER JOIN cuentas_financieras cf ON b.cuenta_bancaria_id = cf.id
		INNER JOIN t_entidad_financiera teb ON cf.banco_id = teb.id
		LEFT JOIN t_comprobantes tc ON o.comprobante_id = tc.id
		INNER JOIN t_medio_pago tmp ON b.medio_pago_id = tmp.id
		LEFT JOIN personas p ON o.persona_asociado_id = p.id
		LEFT JOIN t_documentos td ON p.documento_id = td.id
		INNER JOIN pcge pc ON b.pcge_id = pc.id
		INNER JOIN locales l ON o.local_id = l.id
		WHERE l.empresa_id = $1
		  AND SUBSTRING(b.periodo, 1, 4) = $2
		  AND SUBSTRING(b.periodo, 6, 8) = $3
		  AND b.cuenta_bancaria_id::text = ANY($4)
		  AND o.comprobante_id::text <> ALL(ARRAY['a6062ae0-15a4-11ec-8fec-77a5f80a0a28', '1daedb70-a779-11eb-84c1-40b0344a6892'])
		  AND b.deleted_at IS NULL
		  AND o.deleted_at IS NULL
		ORDER BY i6oo, o.cuo`
	rows, err := b.Connection.Pool.Query(context.Background(), query, businessID, year, month, financialAccountIDs)
	if err != nil {
		return nil, fmt.Errorf("query failed: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var bank v1.LBank
		if err := rows.Scan(
			&bank.BId,
			&bank.BPeriodo,
			&bank.OCuoTipo,
			&bank.BIdentificador,
			&bank.TefCodigo,
			&bank.CfCuenta,
			&bank.OFechaEmion,
			&bank.PgFecha,
			&bank.TmpCodigo,
			&bank.PgGlosa,
			&bank.OGlosa,
			&bank.TdCodigo,
			&bank.PNumero,
			&bank.PRazonSocial,
			&bank.BNumeroTransaccion,
			&bank.BDebe,
			&bank.BHaber,
			&bank.BEstado,
			&bank.PcCodigo,
			&bank.PcDenominacion,
			&bank.OEstadoLe,
			&bank.OObservaciones,
			&bank.OTipoCambio,
			&bank.PgTipoCambio,
			&bank.PgFechaOFechaEmision,
		); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		banks = append(banks, &bank)
	}
	return
}
