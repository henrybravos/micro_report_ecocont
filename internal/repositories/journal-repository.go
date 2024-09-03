package repositories

import (
	"context"
	"github.com/henrybravo/micro-report/pkg/db"
	v1 "github.com/henrybravo/micro-report/protos/gen/go/v1"
	"log"
)

type JournalRepository struct {
	Connection *db.Connection
}

func NewJournalRepository(connection *db.Connection) *JournalRepository {
	return &JournalRepository{Connection: connection}
}
func (r *JournalRepository) GetJournalEntries(businessId, year, month string, isConsolidated, includeClose, includeCuBa bool) ([]*v1.JournalEntry, error) {

	sIc := ""
	if !includeClose {
		sIc = " AND (d.identificador NOT LIKE 'C%'"
		if !includeCuBa {
			sIc += " AND o.glosa <> 'Cierre de Cuentas de Balance'"
		}
		sIc += ")"
	}
	operatorConsolidated := "<="
	if !isConsolidated {
		operatorConsolidated = "="
	}
	query := `
		SELECT p2.codigo, --0
		       p2.denominacion, --1
 				COALESCE(TO_CHAR(p2.created_at, 'DD/MM/YYYY'),'') AS "created_at",
 				COALESCE(TO_CHAR(p2.updated_at, 'DD/MM/YYYY'),'') AS "created_at",
 				COALESCE(TO_CHAR(p2.deleted_at, 'DD/MM/YYYY'),'') AS "created_at",
		       p2.empresa_id, --5 
		       p2.id,
		       COALESCE(efe,''), --7
		       COALESCE(smvn,''),  --8
		       COALESCE(smve,''), --9
		       COALESCE(p2.tipo,''), --10
		       COALESCE(p2.tipo_id::text,''), --11
		       esf, --12
		       ern, --13
		       erf, --14
		       efle, --15
		       ecp, -- 16
		       COALESCE(p2.cierre_id::text,''), --17
		       COALESCE(p2.sequence,0), --18
		       COALESCE(SUM(
			       CASE
			           WHEN o.observaciones = 'inv_ini' AND o.estado_le <> '2' THEN ROUND(d.debe * COALESCE(NULLIF(pg.tipo_cambio, 1), o.tipo_cambio), 4)
			           ELSE NULL
			       END
		       ),0) AS iidebe,
		       COALESCE(SUM(
			       CASE
			           WHEN o.observaciones = 'inv_ini' AND o.estado_le <> '2' THEN ROUND(d.haber * COALESCE(NULLIF(pg.tipo_cambio, 1), o.tipo_cambio), 4)
			           ELSE NULL
			       END
		       ),0) AS iihaber,
		       COALESCE(SUM(
			       CASE
			           WHEN (o.observaciones <> 'inv_ini' OR o.observaciones IS NULL) AND o.estado_le <> '2' THEN ROUND(d.debe * COALESCE(NULLIF(pg.tipo_cambio, 1), o.tipo_cambio), 4)
			           ELSE NULL
			       END
		       ),0) AS debe,
		       COALESCE(SUM(
			       CASE
			           WHEN (o.observaciones <> 'inv_ini' OR o.observaciones IS NULL) AND o.estado_le <> '2' THEN ROUND(d.haber * COALESCE(NULLIF(pg.tipo_cambio, 1), o.tipo_cambio), 4)
			           ELSE NULL
			       END
		       ),0) AS haber
		FROM operaciones o
		         INNER JOIN diarios d ON o.id = d.operacion_id
		         LEFT JOIN pagos pg ON d.referencia_id = pg.id
		         INNER JOIN pcge p2 ON d.pcge_id = p2.id
		         INNER JOIN locales l ON o.local_id = l.id
		WHERE l.empresa_id = $1
		  AND SUBSTRING(d.periodo, 1, 4) = $2
		  AND SUBSTRING(d.periodo, 6, 8) ` + operatorConsolidated + ` $3
		  ` + sIc + `
		  AND o.comprobante_id::text <> ALL(ARRAY['a6062ae0-15a4-11ec-8fec-77a5f80a0a28', '1daedb70-a779-11eb-84c1-40b0344a6892'])
		  AND o.tipo_operacion::text <> ALL(ARRAY['5913663b-a77a-11eb-8923-40b0344a6892', '59133fcc-a77a-11eb-8919-40b0344a6892'])
		  AND o.estado_le <> '2'
		  AND d.deleted_at IS NULL
		  AND o.deleted_at IS NULL
		GROUP BY p2.codigo, p2.denominacion, p2.created_at, p2.updated_at, p2.deleted_at, p2.empresa_id, p2.id, efe, smvn, smve, p2.tipo, tipo_id, esf, ern, erf, efle, ecp, cierre_id, sequence
		ORDER BY p2.codigo`

	rows, err := r.Connection.Pool.Query(context.Background(), query, businessId, year, month)
	if err != nil {
		log.Fatalf("Query failed: %v", err)
	}
	log.Println(businessId, year, month)
	defer rows.Close()
	var journalEntries []*v1.JournalEntry
	for rows.Next() {
		var journal v1.JournalEntry
		if err := rows.Scan(
			&journal.Codigo,
			&journal.Denominacion,
			&journal.CreatedAt,
			&journal.UpdatedAt,
			&journal.DeletedAt,
			&journal.EmpresaId,
			&journal.Id,
			&journal.Efe,
			&journal.Smvn,
			&journal.Smve,
			&journal.Tipo,
			&journal.TipoId,
			&journal.Esf,
			&journal.Ern,
			&journal.Erf,
			&journal.Efle,
			&journal.Ecp,
			&journal.CierreId,
			&journal.Sequence,
			&journal.Iidebe,
			&journal.Iihaber,
			&journal.Debe,
			&journal.Haber,
		); err != nil {
			return nil, err
		}
		journalEntries = append(journalEntries, &journal)
	}
	return journalEntries, nil
}
