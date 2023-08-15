# Image Uploader

Welcome to the Image Uploader repository. This project offers a REST API for users to conveniently upload images.

## Backend Details
The backend starts by establishing an HTTP server using the `net/http` package.

All the source files reside within the `src` directory.

- To launch the HTTP server, execute `go run main.go`. This command will initiate the server on `localhost`, port `4550`.
- To access the system, navigate to `http://localhost:4500/account`. This page presents a login form for user authentication.
- For testing purposes, the username and password are hardcoded as `testuser/secret`. In a real-world scenario, this data should be fetched from a database.

The login page incorporates basic error handling, including checks for empty username and passwords.

- Once users successfully log in, they gain the ability to upload images.
- Image size is limited to 1MB, with supported extensions including JPG and PNG.
- Uploaded images are stored within the `uploaded-receipts` folder.

## Frontend Information
Frontend files are located in the `view` folder. Here, you will find pages containing login forms and image upload interfaces.
