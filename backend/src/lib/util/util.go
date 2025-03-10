package util

import (
	"auto_proforma/src/handlers"
	"net/http"
)

func SendString(rw http.ResponseWriter, text string) {
	rw.Header().Set("Content-Type", "text/plain")
	_, err := rw.Write([]byte(text))
	if err != nil {
		panic(err)
	}
}

func SendExcel(rw http.ResponseWriter, text string) {

	file, err := handlers.CreateExcel(text)
	if err != nil {
		panic(err)
	}

	rw.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	rw.Header().Set("Content-Disposition", `attachment; filename="file.xlsx"`)

	if err := file.Write(rw); err != nil {
		http.Error(rw, "Failed to generate Excel file", http.StatusInternalServerError)
	}
}
