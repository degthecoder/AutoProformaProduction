package home

import (
	"auto_proforma/src/handlers"
	"auto_proforma/src/lib/util"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetExcel(rw http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
	fmt.Println(r.Body)

	util.SendExcel(rw, string(data))
}

func SendCodeTable(rw http.ResponseWriter, req *http.Request) {
	data, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	codes := handlers.GetAttributes(string(data))

	jsonData, err := json.Marshal(codes)
	if err != nil {
		http.Error(rw, "Error converting to JSON", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	_, jsonerr := rw.Write(jsonData)
	if jsonerr != nil {
		http.Error(rw, "Error converting to JSON", http.StatusInternalServerError)
	}
}
