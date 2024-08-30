package repositories

import (
	"context"
	"errors"
	"github.com/henrybravo/micro-report/pkg/db"
	v1 "github.com/henrybravo/micro-report/protos/gen/go/v1"
	"github.com/jackc/pgx/v5"
)

type SalesRepository struct {
	Connection *db.Connection
}

func NewSalesRepository(connection *db.Connection) *SalesRepository {
	return &SalesRepository{Connection: connection}
}

func (r *SalesRepository) GetSalesReports(companyID string, period string, pagination *v1.Pagination) ([]*v1.SalesReport, *v1.Pagination, error) {
	// Count total records
	var totalCount int32
	joinAndWhere := `
     INNER JOIN personas p ON o.persona_asociado_id = p.id
	INNER JOIN ventas v ON o.id = v.operacion_id
	INNER JOIN t_tipo_operacion tto ON o.tipo_operacion = tto.id
	INNER JOIN t_comprobantes tco ON o.comprobante_id = tco.id
	INNER JOIN t_documentos td ON p.documento_id = td.id
	INNER JOIN t_monedas tmo ON o.moneda_id = tmo.id
	LEFT JOIN t_comprobantes tico ON v.tipo_cdpm = tico.id
	INNER JOIN locales l ON l.id = o.local_id
	INNER JOIN empresas e ON e.id = l.empresa_id
	WHERE o.periodo = $2
	  AND e.id = $1
	  AND v.estado_cpe <> '4'
	  AND o.comprobante_id NOT IN ('a6062ae0-15a4-11ec-8fec-77a5f80a0a28', '1daedb70-a779-11eb-84c1-40b0344a6892', '1daeb476-a779-11eb-84b8-40b0344a6892')
	  AND o.tipo_operacion NOT IN ('59133fcb-a77a-11eb-8918-40b0344a6892', '59133fcc-a77a-11eb-8919-40b0344a6892', '5913b410-a77a-11eb-8934-40b0344a6892', '59136634-a77a-11eb-891c-40b0344a6892', '59136635-a77a-11eb-891d-40b0344a6892', '59138d1a-a77a-11eb-8925-40b0344a6892', '59138d1e-a77a-11eb-8929-40b0344a6892', '59138d21-a77a-11eb-892c-40b0344a6892', '59138d23-a77a-11eb-892e-40b0344a6892', '5913b40e-a77a-11eb-8932-40b0344a6892', '5913b40f-a77a-11eb-8933-40b0344a6892', '59136632-a77a-11eb-891a-40b0344a6892')
	  AND o.deleted_at IS NULL
	  AND o.referencia_id IS NULL `

	countQuery := `
        SELECT COUNT(*)
        FROM operaciones o 
        ` + joinAndWhere
	totalPages := 0

	if pagination != nil && pagination.PageSize > 0 && pagination.Offset >= 0 {
		err := r.Connection.Pool.QueryRow(context.Background(), countQuery, companyID, period).Scan(&totalCount)
		if err != nil {
			return nil, nil, err
		}
		totalPages = int((totalCount + pagination.PageSize - 1) / pagination.PageSize)
	}

	if pagination != nil && pagination.Offset > totalCount {
		return nil, nil, errors.New("offset out of range")
	}
	query := `
    SELECT
        o.id, --0
        o.periodo,
        o.cuo,
        v.identificador_linea,
        COALESCE(TO_CHAR(o.fecha_emision, 'DD/MM/YYYY'),'') AS "fecha_emision", --4
        COALESCE(TO_CHAR(o.fecha_emision, 'DD/MM/YYYY'),'') AS "fecEmision",
        COALESCE(TO_CHAR(o.fecha_vencimiento, 'DD/MM/YYYY'),'') AS "fecha_vencimiento",
        COALESCE(TO_CHAR(o.fecha_vencimiento, 'DD/MM/YYYY'),'') AS "fecVencPag",
        tco.codigo, --8
        tco.codigo AS "codTipoCDP",
        o.serie,
        o.serie AS "numSerieCDP",
        o.correlativo,
        o.correlativo AS "numCDP",
        COALESCE(v.numero_final,0),
        td.codigo,
        td.codigo AS "codTipoDocIdentidad",
        p.numero,
        p.numero AS "numDocIdentidad",
        p.razon_social,
        p.razon_social AS "nomRazonSocialCliente",
        v.exportacion,
        v.exportacion AS "mtoValFactExpo", -- 22
        v.base,
        v.base AS "mtoBIGravada",
        v.desc_base,
        v.desc_base AS "mtoDsctoBI",
        v.igv,
        v.igv AS "mtoIGV",
        v.desc_igv,
        v.desc_igv AS "mtoDsctoIGV",
        v.exonerada,
        v.exonerada AS "mtoExonerado",
        v.inafecta,
        v.inafecta AS "mtoInafecto",
        v.isc,
        v.isc AS "mtoISC",
        v.base_ivap,
        v.base_ivap AS "mtoBIIvap",
        v.ivap,
        v.ivap AS "mtoIvap",
        v.otros,
        v.otros AS "mtoOtrosTrib",
        v.total,
        v.total AS "mtoTotalCP",
        tmo.codigo,
        tmo.codigo AS "codMoneda",
        o.tipo_cambio,
        o.tipo_cambio AS "mtoTipoCambio",
        COALESCE( TO_CHAR(v.fecha_cdpm, 'DD/MM/YYYY'),''),
        COALESCE(TO_CHAR(v.fecha_cdpm, 'DD/MM/YYYY'),'') AS "fecEmisionMod",
        COALESCE(tico.codigo,''), -- 51
        COALESCE(tico.codigo,'') AS "codTipoCDPMod",
        COALESCE(v.serie_cdpm,''), -- 53
      	COALESCE(v.serie_cdpm,'') AS "numSerieCDPMod",
        COALESCE(v.numero,''), --55
       	COALESCE(v.numero,'') AS "numCDPMod", --56
    	COALESCE(v.identificador_contrato,''), --57
        COALESCE(v.error_1,FALSE),
        COALESCE(v.identificador,FALSE), --59
        v.estado_operacion, --60
        v.estado_operacion AS "codEstadoComprobante",	--61
        v.icbper, --62
        v.icbper AS "mtoIcbp", --63
        v.estado_cpe, --64
         CASE
			WHEN e.ruc IN ('10095595761', '20523537009', '10093714135', '20601613884') THEN o.observaciones
			ELSE ''
    	END AS observaciones --65
    FROM operaciones o   
    ` + joinAndWhere + ` ORDER BY o.cuo ASC `
	var rows pgx.Rows
	var err error
	if pagination != nil && pagination.PageSize > 0 && pagination.Offset >= 0 {
		query += `LIMIT $3 OFFSET $4`
		rows, err = r.Connection.Pool.Query(context.Background(), query, companyID, period, pagination.PageSize, pagination.Offset)
	} else {
		rows, err = r.Connection.Pool.Query(context.Background(), query, companyID, period)
	}

	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	var sales []*v1.SalesReport
	for rows.Next() {
		var reporte v1.SalesReport
		if err := rows.Scan(
			&reporte.Id,                     // 1
			&reporte.Periodo,                // 2
			&reporte.Cuo,                    // 3
			&reporte.IdentificadorLinea,     // 4
			&reporte.FechaEmision,           // 5
			&reporte.FecEmision,             // 6
			&reporte.FechaVencimiento,       // 7
			&reporte.FecVencPag,             // 8
			&reporte.CodigoTipoCdp,          // 9
			&reporte.CodTipoCdp,             // 10
			&reporte.Serie,                  // 11
			&reporte.NumSerieCdp,            // 12
			&reporte.Correlativo,            // 13
			&reporte.NumCdp,                 // 14
			&reporte.NumeroFinal,            // 15
			&reporte.CodigoTipoDocIdentidad, // 16
			&reporte.CodTipoDocIdentidad,    // 17
			&reporte.NumDocIdentidad,        // 18
			&reporte.NumDocIdentidadClient,  // 19
			&reporte.RazonSocial,            // 20
			&reporte.NomRazonSocialCliente,  // 21
			&reporte.Exportacion,            // 22
			&reporte.MtoValFactExpo,         // 23
			&reporte.Base,                   // 24
			&reporte.MtoBiGravada,           // 25
			&reporte.DescBase,               // 26
			&reporte.MtoDsctoBi,             // 27
			&reporte.Igv,                    // 28
			&reporte.MtoIgv,                 // 29
			&reporte.DescIgv,                // 30
			&reporte.MtoDsctoIgv,            // 31
			&reporte.Exonerada,              // 32
			&reporte.MtoExonerado,           // 33
			&reporte.Inafecta,               // 34
			&reporte.MtoInafecto,            // 35
			&reporte.Isc,                    // 36
			&reporte.MtoIsc,                 // 37
			&reporte.BaseIvap,               // 38
			&reporte.MtoBIIvap,              // 39
			&reporte.Ivap,                   // 40
			&reporte.MtoIvap,                // 41
			&reporte.Otros,                  // 42
			&reporte.MtoOtrosTrib,           // 43
			&reporte.Total,                  // 44
			&reporte.MtoTotalCp,             // 45
			&reporte.CodigoMoneda,           // 46
			&reporte.CodMoneda,              // 47
			&reporte.TipoCambio,             // 48
			&reporte.MtoTipoCambio,          // 49
			&reporte.FechaCdpm,              // 50
			&reporte.FecEmisionMod,          // 51
			&reporte.CodigoTipoCdpMod,       // 52
			&reporte.CodTipoCdpMod,          // 53
			&reporte.NumSerieCdpMod,         // 54
			&reporte.NumCdpMod,              // 55
			&reporte.Numero,                 // 56
			&reporte.NumCdpMod2,             // 57
			&reporte.IdentificadorContrato,  // 58
			&reporte.Error1,                 // 59
			&reporte.Identificador,          // 60
			&reporte.EstadoOperacion,        // 61
			&reporte.CodEstadoComprobante,   // 62
			&reporte.Icbper,                 // 63
			&reporte.MtoIcbp,                // 64
			&reporte.EstadoCpe,              // 65
			&reporte.Observaciones,          // 66
		); err != nil {
			return nil, nil, err
		}
		sales = append(sales, &reporte)
	}

	return sales, &v1.Pagination{TotalCount: int32(totalCount), TotalPages: int32(totalPages)}, nil
}
