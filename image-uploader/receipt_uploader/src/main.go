package main

import (
	"github.com/ermiaswalelgne/receipt_uploader/src/controllers/accountcontroller"
	"github.com/ermiaswalelgne/receipt_uploader/src/controllers/imagecontroller"
	"log"
	"net/http"
)

const uploadPath = "./tmp"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/account", accountcontroller.Index)
	mux.HandleFunc("/account/index", accountcontroller.Index)
	mux.HandleFunc("/account/login", accountcontroller.Login)
	mux.HandleFunc("/receipt/uploadimage", imagecontroller.ImageUploader)
	mux.HandleFunc("/receipt/upload", imagecontroller.UploadHandler)
	log.Print("Server started on localhost:4550, login using /account to for uploading images")

	fs := http.FileServer(http.Dir(uploadPath))
	http.Handle("/files/", http.StripPrefix("/files", fs))

	if err := http.ListenAndServe(":4550", mux); err != nil {
		log.Fatal(err)
	}

}
