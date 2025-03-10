package main

import (
	"auto_proforma/src/app"
	"auto_proforma/src/lib/make_handle"
	"auto_proforma/src/path/home"
	"auto_proforma/src/path/proforma"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	//fsAdmin := http.FileServer(http.Dir(app.Settings.HomeDir + "/assets/"))
	//http.Handle("/assets/", http.StripPrefix("/assets/", fsAdmin))
	http.HandleFunc("/getExcel", make_handle.MakeHandle(home.GetExcel))
	http.HandleFunc("/getTable", make_handle.MakeHandle(home.SendCodeTable))
	http.HandleFunc("/dev", make_handle.MakeHandle(proforma.SaveExcel))
	http.HandleFunc("/getOEM", make_handle.MakeHandle(home.GetSuparCodeWithOEM))

	//corsHandler := cors.New(cors.Options{
	//	AllowedOrigins:   []string{"http://localhost:3000"}, // Allow frontend
	//	AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
	//	AllowedHeaders:   []string{"Content-Type", "Authorization"},
	//	AllowCredentials: true,
	//})

	//myCors := &MyCors{Cors: corsHandler}

	srv := &http.Server{
		Addr:        app.Settings.Host + ":" + app.Settings.Port,
		ReadTimeout: 30 * time.Second,
		//Handler:      myCors.corsWrapper(),
		WriteTimeout: 30 * time.Second,
	}

	fmt.Printf("Backend server starting server on port: %s\n", app.Settings.Port)
	err := srv.ListenAndServe()

	if err != nil {
		log.Println(err.Error())
	}

	defer app.DisconnectDb()
}
