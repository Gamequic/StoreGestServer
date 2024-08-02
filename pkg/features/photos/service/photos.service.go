package photoservice

import (
	"net/http"
	"os"
	"storegestserver/utils/middlewares"
	"strings"

	"github.com/google/uuid"
)

// Create the static folder
func InitPhotosService() {
	dirName := "static"

	// Create the directory with 0755 permissions (readable and executable by everyone, writable by the owner)
	err := os.Mkdir(dirName, 0755)
	if err != nil {
		if err.Error() != "mkdir static: file exists" { // If the error is "folder exists" ignore
			panic(err)
		}
	}
}

// CRUD Operations

func Create(r *http.Request) string {
	// Generate a UUIDv4
	UUID := uuid.New()

	// Parse the multipart form, with a limit of 10 MB
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		panic(middlewares.GormError{Code: http.StatusBadRequest, Message: "Unable to parse form", IsGorm: true})
	}

	// Get the file from the form
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		panic(middlewares.GormError{Code: http.StatusBadRequest, Message: "Unable to get file", IsGorm: true})
	}
	defer file.Close()

	// Validate the file type (Content-Type)
	contentType := fileHeader.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "image/") {
		panic(middlewares.GormError{Code: http.StatusBadRequest, Message: "Invalid file type", IsGorm: true})
	}

	// Define a map of allowed file extensions
	var allowedExtensions = map[string]string{
		".jpg":  ".jpg",
		".jpeg": ".jpg",
		".png":  ".png",
		".gif":  ".gif",
	}

	// Extract the file extension and add it to the file name
	fileName := fileHeader.Filename
	extension := ""
	for ext := range allowedExtensions {
		if strings.HasSuffix(fileName, ext) {
			extension = allowedExtensions[ext]
			break
		}
	}

	// Check if the extension is valid
	if extension == "" {
		panic(middlewares.GormError{Code: http.StatusBadRequest, Message: "Unsupported file format: " + fileName, IsGorm: true})
	}

	// Create a new file in the server's file system
	filePath := "static/" + UUID.String() + extension
	dst, err := os.Create(filePath)
	if err != nil {
		panic(middlewares.GormError{Code: http.StatusInternalServerError, Message: "Error creating file:" + err.Error(), IsGorm: true})
	}
	defer dst.Close()

	// Copy the uploaded file to the destination file
	_, err = dst.ReadFrom(file)
	if err != nil {
		panic(middlewares.GormError{Code: http.StatusInternalServerError, Message: "Unable to save file:" + err.Error(), IsGorm: true})
	}

	return UUID.String() + extension
}

func Delete(photo string) {
	filePath := "static/" + photo

	err := os.Remove(filePath)
	if err != nil {
		if err.Error() == "remove "+filePath+": no such file or directory" {
			panic(middlewares.GormError{Code: http.StatusNotFound, Message: "File not found:" + err.Error(), IsGorm: true})
		} else {
			panic(middlewares.GormError{Code: http.StatusInternalServerError, Message: "Unable to delete file:" + err.Error(), IsGorm: true})
		}
	}
}
