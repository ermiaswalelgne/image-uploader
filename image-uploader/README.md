# Reciept-uploader

This is REST API for uploading images from user.

The API has an http server

##  The Backend
 Start off by creating an HTTP server using the net/http package.

 All the soruces files are under src directory
<ul>
<li> To start the http server run `go run main.go` Thi will reun the server on localhost port 4550.
To login to the system use  `http://localhost:4500/account `. This page have  have a loging form to authenticate users.
For test purose username and password are hardcoded as testuser/secret

The login page has basic error handlng such as checking empty username and passowrd

<<li> Once the user has successfully login users acan upload images.
 The maximum image sie is set to 1MB, allowd image extentionas re JPG and PNG.
 Images are saved under uploaded-receipts folder.
</ul>

## Front End
<ul>
<li> Front end  fies are under view folder, whre login forma and image uploading pages are stored
</ul>