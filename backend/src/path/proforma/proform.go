package proforma

import (
	"auto_proforma/src/handlers"
	"net/http"
)

func SaveExcel(rw http.ResponseWriter, r *http.Request) {
	//data, err := io.ReadAll(r.Body)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(string(data))
	//fmt.Println(r.Body)
	codes := "82710PS\nA4100P\nA4210P\nA4220P\n50181P\n51492P\n17310PS\n20220PS\n20310P\n20370PS\n20060P\n20133P\n58872PS\n58860P\n25590P\n27T70P\n50162P\n52600P\n54171P\n54181P\n63321P\n63331P\n20050P"

	file, err := handlers.CreateProforma(codes)
	if err != nil {
		panic(err)
	}
	filepath := "output.xlsx"

	if err := file.SaveAs(filepath); err != nil {
		panic(err)
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("OK"))
}
