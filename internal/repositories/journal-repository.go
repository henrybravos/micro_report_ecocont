package repositories

import (
	"context"
	"fmt"
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

func (r *JournalRepository) GetLfJournals(businessID, year, month string, isConsolidated, includeClose, includeCuBa bool, typeReq string) ([]map[string]interface{}, error) {
	var sIc string
	if !includeClose {
		sIc = " AND (d.identificador not like 'C%'"
		if !includeCuBa {
			sIc += " AND o.glosa <> 'Cierre de Cuentas de Balance'"
		}
		sIc += ")"
	}

	o51v61 := ""
	if typeReq == "060100" {
		o51v61 = "i4,"
	}
	query := fmt.Sprintf(`
        SELECT
            o.id AS id, --0
            concat(o.cuo, '-', d.tipo) AS i2, --1
            d.identificador AS i3, --2
            p2.codigo AS i4, --3
            p2.denominacion AS denominacion, --4
            COALESCE(d.unidad_operacion, '') AS i5, --5
            COALESCE(d.centro_costo,'') AS i6, --6
            tm.codigo AS i7, --7
            COALESCE(td.codigo,'') AS i8, --8
            COALESCE(p.numero, '') AS i9, --9
            COALESCE(tc.codigo,'') AS i10, --10
            COALESCE(o.serie,'') AS i11, --11
            COALESCE(o.correlativo,'') AS i12,	--12
            COALESCE(o.fecha_contable::text,'') AS i13, --13
            COALESCE(o.fecha_vencimiento::text,'') AS i14, --14
            COALESCE(o.fecha_emision::text, '') AS i15, --15
            COALESCE(pg.fecha::text, '') AS i15pg,  --16
            o.glosa AS i16o, --17
            COALESCE(pg.glosa,'') AS i16, --18
            COALESCE(o.glosa_referencia,'') AS i17,
            COALESCE(d.debe,0) AS i18, --20
            COALESCE(d.haber,0) AS i19,
            o.dato_estructurado AS i20, --22
            COALESCE(d.estado::text,'') AS i21, --23
            o.estado_le AS estado_le, --24
            o.tipo_cambio AS tipo_cambio,
            o.codigo_libro AS codigo_libro,
            o.periodo AS periodo, --27
            o.cuo AS cuo, --28
            o.observaciones AS observaciones, --29
            d.pcge_id AS pcge_id,
            COALESCE( pg.tipo_cambio,0) AS p_tipo_cambio --31
        FROM operaciones o
        INNER JOIN diarios d ON o.id = d.operacion_id
        LEFT JOIN pagos pg ON d.referencia_id = pg.id
        LEFT JOIN personas p ON o.persona_asociado_id = p.id
        LEFT JOIN t_documentos td ON p.documento_id = td.id
        LEFT JOIN t_comprobantes tc ON o.comprobante_id = tc.id
        INNER JOIN t_monedas tm ON o.moneda_id = tm.id
        INNER JOIN pcge p2 ON d.pcge_id = p2.id
        INNER JOIN locales l ON o.local_id = l.id
        WHERE l.empresa_id = $1
        AND SUBSTRING(d.periodo, 1, 4) = $2
        AND SUBSTRING(d.periodo, 6, 8) %s $3
        %s
        AND o.comprobante_id::text <> ALL(ARRAY['a6062ae0-15a4-11ec-8fec-77a5f80a0a28', '1daedb70-a779-11eb-84c1-40b0344a6892'])
        AND o.tipo_operacion::text <> ALL(ARRAY['5913663b-a77a-11eb-8923-40b0344a6892', '59133fcc-a77a-11eb-8919-40b0344a6892'])
        AND d.deleted_at IS NULL
        AND o.deleted_at IS NULL
        ORDER BY %s o.fecha_emision, i2, d.identificador
    `, map[bool]string{true: "<=", false: "="}[isConsolidated], sIc, o51v61)

	rows, err := r.Connection.Pool.Query(context.Background(), query, businessID, year, month)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var results []map[string]interface{}

	for rows.Next() {
		var id, i2, i3, i4, denominacion, i5, i6, i7, i8, i9, i10 string
		var i11, i12, i13, i14, i15, i15pg, i16o, i16, i17, i20, estadoLe, tipoCambio, codigoLibro, periodo, cuo, observaciones, i21 string
		var pcgeId string
		var i18, i19, pTipoCambio float64
		err := rows.Scan(
			&id, &i2, &i3, &i4, &denominacion, &i5, &i6, &i7, &i8, &i9, &i10,
			&i11, &i12, &i13, &i14, &i15, &i15pg, &i16o, &i16, &i17, &i18, &i19, &i20, &i21,
			&estadoLe, &tipoCambio, &codigoLibro, &periodo, &cuo, &observaciones, &pcgeId, &pTipoCambio,
		)
		if err != nil {
			return nil, err
		}
		row := map[string]interface{}{
			"id":            id,
			"i2":            i2,
			"i3":            i3,
			"i4":            i4,
			"denominacion":  denominacion,
			"i5":            i5,
			"i6":            i6,
			"i7":            i7,
			"i8":            i8,
			"i9":            i9,
			"i10":           i10,
			"i11":           i11,
			"i12":           i12,
			"i13":           i13,
			"i14":           i14,
			"i15":           i15,
			"i15pg":         i15pg,
			"i16o":          i16o,
			"i16":           i16,
			"i17":           i17,
			"i18":           i18,
			"i19":           i19,
			"i20":           i20,
			"i21":           i21,
			"estado_le":     estadoLe,
			"tipo_cambio":   tipoCambio,
			"codigo_libro":  codigoLibro,
			"periodo":       periodo,
			"cuo":           cuo,
			"observaciones": observaciones,
			"pcge_id":       pcgeId,
			"p_tipo_cambio": pTipoCambio,
		}

		results = append(results, row)

	}

	return results, nil
}
