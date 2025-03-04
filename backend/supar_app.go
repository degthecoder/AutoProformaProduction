package main

import (
	"auto_proforma/src/app"
	"auto_proforma/src/lib/make_handle"
	"auto_proforma/view/home"
	"fmt"
	"github.com/rs/cors"
	"log"
	"net/http"
	"time"
)

func main() {
	//fsAdmin := http.FileServer(http.Dir(app.Settings.HomeDir + "/assets/"))
	//http.Handle("/assets/", http.StripPrefix("/assets/", fsAdmin))

	http.HandleFunc("/getExcel", make_handle.MakeHandle(home.GetExcel))

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Allow frontend
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	srv := &http.Server{
		Addr:         app.Settings.Host + ":" + app.Settings.Port,
		ReadTimeout:  30 * time.Second,
		Handler:      corsHandler,
		WriteTimeout: 30 * time.Second,
	}

	fmt.Printf("Backend server starting server on port: %s\n", app.Settings.Port)
	err := srv.ListenAndServe()

	if err != nil {
		log.Println(err.Error())
	}

	defer app.DisconnectDb()
}
