package home

import (
	"auto_proforma/src/lib/util"
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
