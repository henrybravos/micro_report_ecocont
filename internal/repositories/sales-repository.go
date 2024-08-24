package repositories

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/henrybravo/micro-report/pkg/db"
	"github.com/jackc/pgx/v5"
)

type SalesRepository struct {
	Connection *db.Connection
}

func NewSalesRepository(connection *db.Connection) *SalesRepository {
	return &SalesRepository{Connection: connection}
}

type SalesReport struct {
	ID                     uuid.UUID      // 1
	Periodo                string         // 2
	Cuo                    string         // 3
	IdentificadorLinea     string         // 4
	FechaEmision           sql.NullTime   // 5
	FecEmision             string         // 6
	FechaVencimiento       sql.NullTime   // 7
	FecVencPag             sql.NullString // 8
	CodigoTipoCDP          string         // 9
	CodTipoCDP             string         // 10
	Serie                  string         // 11
	NumSerieCDP            string         // 12
	Correlativo            string         // 13
	NumCDP                 string         // 14
	NumeroFinal            sql.NullString // 15
	CodigoTipoDocIdentidad string         // 16
	CodTipoDocIdentidad    string         // 17
	NumDocIdentidad        string         // 18
	NumDocIdentidadClient  string         // 19
	RazonSocial            string         // 20
	NomRazonSocialCliente  string         // 21
	Exportacion            float64        // 22
	MtoValFactExpo         float64        // 23
	Base                   float64        // 24
	MtoBIGravada           float64        // 25
	DescBase               float64        // 26
	MtoDsctoBI             float64        // 27
	IGV                    float64        // 28
	MtoIGV                 float64        // 29
	DescIGV                float64        // 30
	MtoDsctoIGV            float64        // 31
	Exonerada              float64        // 32
	MtoExonerado           float64        // 33
	Inafecta               float64        // 34
	MtoInafecto            float64        // 35
	ISC                    float64        // 36
	MtoISC                 float64        // 37
	BaseIVAP               float64        // 38
	MtoBIIvap              float64        // 39
	IVAP                   float64        // 40
	MtoIvap                float64        // 41
	Otros                  float64        // 42
	MtoOtrosTrib           float64        // 43
	Total                  float64        // 44
	MtoTotalCP             float64        // 45
	CodigoMoneda           string         // 46
	CodMoneda              string         // 47
	TipoCambio             float64        // 48
	MtoTipoCambio          float64        // 49
	FechaCDPM              sql.NullTime   // 50
	FecEmisionMod          sql.NullString // 51
	CodigoTipoCDPMod       sql.NullString // 52
	CodTipoCDPMod          sql.NullString // 53
	NumSerieCDPMod         sql.NullString // 54
	NumCDPMod              sql.NullString // 55
	Numero                 sql.NullString // 56
	NumCDPMod2             sql.NullString // 57
	IdentificadorContrato  sql.NullString // 58
	Error1                 sql.NullString // 59
	Identificador          sql.NullString // 60
	EstadoOperacion        sql.NullString // 61
	CodEstadoComprobante   sql.NullString // 62
	ICBPER                 float64        // 63
	MtoIcbp                float64        // 64
	EstadoCPE              sql.NullString // 65
	Observaciones          sql.NullString // 66
}

func (r *SalesRepository) GetSalesReports(companyID string, period string, pagination PaginationParams) ([]SalesReport, *Pagination, error) {
	// Count total records
	var totalCount int
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
          AND o.comprobante_id::text <> ALL(ARRAY['a6062ae0-15a4-11ec-8fec-77a5f80a0a28', '1daedb70-a779-11eb-84c1-40b0344a6892', '1daeb476-a779-11eb-84b8-40b0344a6892'])
          AND o.tipo_operacion::text <> ALL(ARRAY['59133fcb-a77a-11eb-8918-40b0344a6892', 
                  '59133fcc-a77a-11eb-8919-40b0344a6892', '5913b410-a77a-11eb-8934-40b0344a6892',
                  '59136634-a77a-11eb-891c-40b0344a6892', '59136635-a77a-11eb-891d-40b0344a6892',
                  '59138d1a-a77a-11eb-8925-40b0344a6892', '59138d1e-a77a-11eb-8929-40b0344a6892',
                  '59138d21-a77a-11eb-892c-40b0344a6892', '59138d23-a77a-11eb-892e-40b0344a6892',
                  '5913b40e-a77a-11eb-8932-40b0344a6892', '5913b40f-a77a-11eb-8933-40b0344a6892',
                  '59136632-a77a-11eb-891a-40b0344a6892'])
          AND o.deleted_at IS NULL
          AND o.referencia_id IS NULL `

	countQuery := `
        SELECT COUNT(*)
        FROM operaciones o 
        ` + joinAndWhere

	err := r.Connection.Pool.QueryRow(context.Background(), countQuery, companyID, period).Scan(&totalCount)
	if err != nil {
		return nil, nil, err
	}

	totalPages := 0
	if pagination.Limit > 0 {
		totalPages = (totalCount + pagination.Limit - 1) / pagination.Limit
	}

	if pagination.Offset > totalCount {
		return nil, &Pagination{
			TotalCount: totalCount,
			TotalPages: totalPages,
			Err:        "Offset out of range",
		}, nil
	}
	query := `
    SELECT
        o.id,
        o.periodo,
        o.cuo,
        v.identificador_linea,
        o.fecha_emision,
        TO_CHAR(o.fecha_emision, 'DD/MM/YYYY') AS "fecEmision",
        o.fecha_vencimiento,
        TO_CHAR(o.fecha_vencimiento, 'DD/MM/YYYY') AS "fecVencPag",
        tco.codigo,
        tco.codigo AS "codTipoCDP",
        o.serie,
        o.serie AS "numSerieCDP",
        o.correlativo,
        o.correlativo AS "numCDP",
        v.numero_final,
        td.codigo,
        td.codigo AS "codTipoDocIdentidad",
        p.numero,
        p.numero AS "numDocIdentidad",
        p.razon_social,
        p.razon_social AS "nomRazonSocialCliente",
        v.exportacion,
        v.exportacion AS "mtoValFactExpo",
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
        v.fecha_cdpm,
        TO_CHAR(v.fecha_cdpm, 'DD/MM/YYYY') AS "fecEmisionMod",
        tico.codigo,
        tico.codigo AS "codTipoCDPMod",
        v.serie_cdpm,
        v.serie_cdpm AS "numSerieCDPMod",
        v.numero,
        v.numero AS "numCDPMod",
        v.identificador_contrato,
        v.error_1,
        v.identificador,
        v.estado_operacion,
        v.estado_operacion AS "codEstadoComprobante",
        v.icbper,
        v.icbper AS "mtoIcbp",
        v.estado_cpe,
        CASE
            WHEN e.ruc = '10095595761' OR e.ruc = '20523537009' OR e.ruc = '10093714135' OR e.ruc = '20601613884' THEN
                o.observaciones
        END AS observaciones
    FROM operaciones o   
    ` + joinAndWhere + ` ORDER BY o.fecha_emision DESC `
	var rows pgx.Rows
	if pagination.Pagination && pagination.Limit > 0 && pagination.Offset >= 0 {
		query += `LIMIT $3 OFFSET $4`
		rows, err = r.Connection.Pool.Query(context.Background(), query, companyID, period, pagination.Limit, pagination.Offset)
	} else {
		rows, err = r.Connection.Pool.Query(context.Background(), query, companyID, period)
	}

	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	var sales []SalesReport
	for rows.Next() {
		var reporte SalesReport
		if err := rows.Scan(
			&reporte.ID,                     // 1
			&reporte.Periodo,                // 2
			&reporte.Cuo,                    // 3
			&reporte.IdentificadorLinea,     // 4
			&reporte.FechaEmision,           // 5
			&reporte.FecEmision,             // 6
			&reporte.FechaVencimiento,       // 7
			&reporte.FecVencPag,             // 8
			&reporte.CodigoTipoCDP,          // 9
			&reporte.CodTipoCDP,             // 10
			&reporte.Serie,                  // 11
			&reporte.NumSerieCDP,            // 12
			&reporte.Correlativo,            // 13
			&reporte.NumCDP,                 // 14
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
			&reporte.MtoBIGravada,           // 25
			&reporte.DescBase,               // 26
			&reporte.MtoDsctoBI,             // 27
			&reporte.IGV,                    // 28
			&reporte.MtoIGV,                 // 29
			&reporte.DescIGV,                // 30
			&reporte.MtoDsctoIGV,            // 31
			&reporte.Exonerada,              // 32
			&reporte.MtoExonerado,           // 33
			&reporte.Inafecta,               // 34
			&reporte.MtoInafecto,            // 35
			&reporte.ISC,                    // 36
			&reporte.MtoISC,                 // 37
			&reporte.BaseIVAP,               // 38
			&reporte.MtoBIIvap,              // 39
			&reporte.IVAP,                   // 40
			&reporte.MtoIvap,                // 41
			&reporte.Otros,                  // 42
			&reporte.MtoOtrosTrib,           // 43
			&reporte.Total,                  // 44
			&reporte.MtoTotalCP,             // 45
			&reporte.CodigoMoneda,           // 46
			&reporte.CodMoneda,              // 47
			&reporte.TipoCambio,             // 48
			&reporte.MtoTipoCambio,          // 49
			&reporte.FechaCDPM,              // 50
			&reporte.FecEmisionMod,          // 51
			&reporte.CodigoTipoCDPMod,       // 52
			&reporte.CodTipoCDPMod,          // 53
			&reporte.NumSerieCDPMod,         // 54
			&reporte.NumCDPMod,              // 55
			&reporte.Numero,                 // 56
			&reporte.NumCDPMod2,             // 57
			&reporte.IdentificadorContrato,  // 58
			&reporte.Error1,                 // 59
			&reporte.Identificador,          // 60
			&reporte.EstadoOperacion,        // 61
			&reporte.CodEstadoComprobante,   // 62
			&reporte.ICBPER,                 // 63
			&reporte.MtoIcbp,                // 64
			&reporte.EstadoCPE,              // 65
			&reporte.Observaciones,          // 66
		); err != nil {
			return nil, nil, err
		}
		sales = append(sales, reporte)
	}

	return sales, &Pagination{TotalCount: totalCount, TotalPages: totalPages}, nil
}
