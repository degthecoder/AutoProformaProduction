package main

import (
	"auto_proforma/src/app"
	"auto_proforma/src/lib/make_handle"
	"auto_proforma/view/home"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/getExcel", make_handle.MakeHandle(home.GetExcel))
	http.HandleFunc("/getTable", make_handle.MakeHandle(home.SendCodeTable))

	srv := &http.Server{
		Addr:        app.Settings.Host + ":" + app.Settings.Port,
		ReadTimeout: 30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	fmt.Printf("Backend server starting server on port: %s\n", app.Settings.Port)
	err := srv.ListenAndServe()

	if err != nil {
		log.Println(err.Error())
	}

	defer app.DisconnectDb()
}
