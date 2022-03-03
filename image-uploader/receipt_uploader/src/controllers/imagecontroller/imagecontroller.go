package imagecontroller

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	sessions "github.com/ermiaswalelgne/receipt_uploader/src/sessions"
)

const MAX_UPLOAD_SIZE = 1024 * 1024 // 1MB

func ImageUploader(res http.ResponseWriter, req *http.Request) {
	// get user name from the session
	session, err := sessions.Get(req)
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
	username := session.Values["username"]

	data := map[string]interface{}{
		"username": username,
	}
	tmp, _ := template.ParseFiles("views/accountcontroller/uploadimage.html")
	tmp.Execute(res, data)
}

func UploadHandler(res http.ResponseWriter, req *http.Request) {
	// get user name from the session
	session, err := sessions.Get(req)
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
	username := session.Values["username"]

	if req.URL.Path != "/receipt/upload" {
		http.Error(res, "404 not found.", http.StatusNotFound)
		return
	}
	if req.Method != "POST" {
		http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(res, "Uploading images ...\n")

	//restrict the maximum size to be uploaded using http.MaxBytesReade method to avoid uncessary  wasting of server resources.*/
	req.Body = http.MaxBytesReader(res, req.Body, MAX_UPLOAD_SIZE)
	if err := req.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		http.Error(res, "The image size is  too larg. Please choose an an image < 1MB", http.StatusBadRequest)
		return
	}

	//Save the uploaded image
	// The argument to FormFile must match the name attribute
	// of the file input on the frontend
	file, fileHeader, err := req.FormFile("file")
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	defer file.Close()

	buff := make([]byte, 512)
	_, err = file.Read(buff)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	filetype := http.DetectContentType(buff)
	if filetype != "image/jpeg" && filetype != "image/png" {
		{
			http.Error(res, "The provided file format is not allowed. Please upload a JPEG or PNG image", http.StatusBadRequest)
			return
		}
	}
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create the uploaded-receipts folder if it doesn't
	err = os.MkdirAll("./uploaded-receipts", os.ModePerm)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a new file in the uploaded-receipts directory
	dst, err := os.Create(fmt.Sprintf("./uploaded-receipts/%d%s%s", time.Now().UnixNano(), "_", fileHeader.Filename))
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	defer dst.Close()

	// Copy the uploaded file to the filesystem
	// at the specified destination
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	// report image infromation
	fmt.Println("Image info")
	fmt.Println("Image Name:", fileHeader.Filename)
	fmt.Println("Image size:", fileHeader.Size)
	fmt.Println("Image Type:", fileHeader.Header.Get("Content-Type"))

	fmt.Fprintf(res, "Thanks, your receipt has upload successfully")
	log.Printf("User: %s has upload a recipt: %s successfully", username, fileHeader.Filename)
}
