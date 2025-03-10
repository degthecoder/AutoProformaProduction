package handlers

import (
	"auto_proforma/src/app"
	"database/sql"
	"fmt"
	"net/http"
	"strings"
)

type Supap struct {
	OrderNo      string `json:"orderNo"`
	SuparCode    string `json:"suparCode"`
	MakeModel    string `json:"makeModel"`
	Type         string `json:"type"`
	OriginalCode string `json:"originalCode"`
	ItemNo       string `json:"itemNo"`
	Quantity     string `json:"quantity"`
	UnitPrice    string `json:"unitPrice"`
	UrunGrupKodu string `json:"urunGrupKodu"`
}

func GetSpecification(codes string) []Supap {
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
			spec.ORJINAL_KODU,
			spec.URUN_GRUP_KODU
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
		if err := rows.Scan(&supap.MakeModel, &supap.Type, &supap.SuparCode, &supap.OriginalCode, &supap.UrunGrupKodu); err != nil {
			panic(err)
		}
		supaps = append(supaps, supap)
	}

	return supaps
}

func GetSuparCodesFromOEM(oem string) []Supap {
	oemNo := strings.Fields(oem)

	var placeholders []string
	args := make([]interface{}, len(oemNo)*2)

	for i, code := range oemNo {
		if code == "" { // Skip empty values
			continue
		}

		paramName := fmt.Sprintf("param%d", i)
		placeholders = append(placeholders, fmt.Sprintf("@%s", paramName))
		args[i] = sql.Named(fmt.Sprintf("param%d", i), code)

		if strings.Contains(code, "-") {
			noHyphenCode := strings.ReplaceAll(code, "-", "")
			paramNameNoHyphen := fmt.Sprintf("param%d_nohyphen", i)
			placeholders = append(placeholders, fmt.Sprintf("@%s", paramNameNoHyphen))
			args = append(args, sql.Named(paramNameNoHyphen, noHyphenCode))
		}
	}

	query := fmt.Sprintf(`SELECT
			spec.ORJINAL_KODU,
			spec.STOK_KODU,
			spec.URUN_GRUP_KODU
		FROM
			AA_Spesifikasyon AS spec
		WHERE
			spec.ORJINAL_KODU IN (%s) ORDER BY spec.STOK_KODU ASC;`, strings.Join(placeholders, ","))

	rows, err := app.Db.Query(query, args...)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var supaps []Supap
	for rows.Next() {
		var supap Supap
		if err := rows.Scan(&supap.OriginalCode, &supap.SuparCode, &supap.UrunGrupKodu); err != nil {
			panic(err)
		}
		supaps = append(supaps, supap)
	}
	return supaps
}

func GetProformAttributes(w http.ResponseWriter, r *http.Request) {

}
