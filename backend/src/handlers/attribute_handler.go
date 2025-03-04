package handlers

import (
	"auto_proforma/src/app"
	"database/sql"
	"fmt"
	"strings"
)

type Supap struct {
	SuparCode    string `json:"suparCode"`
	MakeModel    string `json:"makeModel"`
	Type         string `json:"type"`
	OriginalCode string `json:"originalCode"`
}

func GetAttributes(codes string) []Supap {
	suparCodes := strings.Fields(codes)

	placeholders := make([]string, len(suparCodes))
	args := make([]interface{}, len(suparCodes))

	for i, code := range suparCodes {
		placeholders[i] = fmt.Sprintf("@param%d", i)
		args[i] = sql.Named(fmt.Sprintf("param%d", i), code)
	}

	query := fmt.Sprintf(`SELECT
			stok.INGISIM,
			stok.CINSI,
			stok.STOK_KODU,
			spec.ORJINAL_KODU
		FROM
			AA_STOK_1 AS stok
		LEFT JOIN
			AA_Spesifikasyon AS spec
		ON
			stok.STOK_KODU = spec.STOK_KODU
		WHERE
			stok.STOK_KODU IN (%s) ORDER BY stok.STOK_KODU ASC;`, strings.Join(placeholders, ","))

	rows, err := app.Db.Query(query, args...)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var supaps []Supap
	for rows.Next() {
		var supap Supap
		if err := rows.Scan(&supap.MakeModel, &supap.Type, &supap.SuparCode, &supap.OriginalCode); err != nil {
			panic(err)
		}
		supaps = append(supaps, supap)
	}

	return supaps
}
