# Image-uploader

This is REST API for uploading images from the user.

The API has an HTTP server

##  The Backend
 Starts by creating an HTTP server using the net/http package.

 All the source files are under src directory
<ul>
<li> To start the HTTP server run `go run main.go`. This will run the server on localhost port 4550.
To login to the system use  `http://localhost:4500/account `. This page has a login form to authenticate users. 
For test purposes username and password are hardcoded as testuser/secret, though in real-world this has to be fetched from database

The login page has basic error handling such as checking empty username and passwords.

<<li> Once the user has successfully login users can upload images.
 The maximum image size is set to 1MB, allowing image extensions re JPG and PNG.
 Images are saved under uploaded-receipts folder.
</ul>

## Front End
<ul>
<li> Front end  fies are under view folder, whre login forma and image uploading pages are stored
</ul>
