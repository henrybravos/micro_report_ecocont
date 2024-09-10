package repositories

import (
	"context"
	"fmt"
	"github.com/henrybravo/micro-report/pkg/db"
	v1 "github.com/henrybravo/micro-report/protos/gen/go/v1"
	"strings"
)

type KardexRepository struct {
	Connection *db.Connection
}

func NewKardexRepository(connection *db.Connection) *KardexRepository {
	return &KardexRepository{Connection: connection}
}
func (k *KardexRepository) GetReportKardex(localID, startDate, endDate, productID string, isNotes, perPeriod bool) ([]*v1.KardexValued, error) {
	nv := "a6062ae0-15a4-11ec-8fec-77a5f80a0a28"
	if isNotes {
		nv = "fab8b709-c736-4fd3-9e14-727fe2ab8eed"
	}
	orderBy := " ORDER BY fecha ASC, libro ASC, numero ASC"
	query := "select incluir_guias_kardex from locales where id=$1"
	row := k.Connection.Pool.QueryRow(context.Background(), query, localID)
	var includeGuides bool
	err := row.Scan(&includeGuides)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}
	locales := []string{"4d827b06-1820-11ed-b22a-000c298054cc", "e80cae84-6bb4-11ed-bf23-000c2901d02f"}
	isOnline := false
	for _, locale := range locales {
		if localID == locale {
			isOnline = true
			break
		}
	}

	var queryOnline string
	if isOnline {
		for i, locale := range locales {
			if i > 0 {
				queryOnline += " OR "
			}
			queryOnline += fmt.Sprintf("o.local_id='%s'", locale)
		}
	}

	query = `
		SELECT 
		CONCAT(tmed.codigo, ' - ', tmed.nombre) as medida, --0
		CONCAT(inv.codigo, ' - ', inv.descripcion) as inventario,
		CONCAT(p.codigo, ' - ', p.descripcion) as existencia, --2
		o.fecha_emision::text as fecha, 
		tco.codigo as tipo, 
		o.serie as serie, 
		o.correlativo as numero, 
		tm.codigo as tipo_operacion, 
		m.cantidad as cantidad,
		m.costo_unitario as costo, --9
		COALESCE(o.tipo_operacion_kardex, '') as kardex, 
		o.codigo_libro as libro, 
		COALESCE(m.equivalencia, 0) as equivalencia, 
		cat.nombre as categoria, 
		p.id as producto_id, 
		COALESCE(p.codigo_barras, '') as codigo_barras, 
		o.tipo_cambio as tipo_cambio
		FROM movimientos m 
		INNER JOIN operaciones o ON m.operacion_id=o.id 
		INNER JOIN t_tipo_operacion tm ON tm.id= o.tipo_operacion 
		INNER JOIN t_comprobantes tc ON tc.id= o.comprobante_id 
		INNER JOIN productos p ON p.id= m.producto_id 
		INNER JOIN t_inventarios inv ON inv.id= p.inventario_id 
		INNER JOIN t_medidas tmed ON tmed.id=p.medida_id 
		INNER JOIN t_comprobantes tco ON tco.id=o.comprobante_id
		INNER JOIN categorias cat ON cat.id=p.categoria_id
		WHERE (m.anticipo = false OR m.anticipo IS NULL) 
		AND p.deleted_at IS NULL AND o.fecha_emision >= $2 AND o.fecha_emision <= $3 AND `

	if isOnline {
		query += fmt.Sprintf("(%s) ", queryOnline)
	} else {
		query += "o.local_id = $1 "
	}
	query += `
		AND o.comprobante_id <> $4 
		AND o.codigo_libro IS NOT NULL 
		AND COALESCE(o.serie, '') NOT ILIKE 'V%' 
		AND COALESCE(o.serie, '') NOT ILIKE 'P%' 
		AND COALESCE(o.serie, '') NOT ILIKE 'O%' 
		AND o.deleted_at IS NULL 
		AND m.deleted_at IS NULL 
		AND o.estado_le = '1'`
	if perPeriod {
		query = strings.Replace(query, "o.fecha_emision >= $2 AND o.fecha_emision <= $3", "o.periodo >= $2 AND o.periodo <= $3", 1)
	}
	if productID != "" {
		query += fmt.Sprintf(" AND p.id = '%s'", productID)
	}
	if !includeGuides {
		query += " AND o.comprobante_id != '1daedb70-a779-11eb-84c1-40b0344a6892'"
	}
	query += orderBy
	rows, err := k.Connection.Pool.Query(context.Background(), query, localID, startDate, endDate, nv)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}
	defer rows.Close()
	var results []*v1.KardexValued
	for rows.Next() {
		var kardex v1.KardexValued
		err := rows.Scan(
			&kardex.Medida, &kardex.Inventario, &kardex.Existencia, &kardex.FechaEmision, &kardex.Tipo, &kardex.Serie, &kardex.Numero,
			&kardex.TipoOpercion, &kardex.Cantidad, &kardex.Costo, &kardex.Kardex, &kardex.Libro, &kardex.Equivalencia,
			&kardex.Categoria, &kardex.ProductoId, &kardex.CodigoBarras, &kardex.TipoCambio,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		results = append(results, &kardex)
	}

	return results, nil
}
