package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/drive/v3"
)

//Use Service account
func ServiceAccount(secretFile string) *http.Client {
	b, err := ioutil.ReadFile(secretFile)
	if err != nil {
		log.Fatal("error while reading the credential file", err)
	}
	var s = struct {
		Email      string `json:"client_email"`
		PrivateKey string `json:"private_key"`
	}{}
	json.Unmarshal(b, &s)
	config := &jwt.Config{
		Email:      s.Email,
		PrivateKey: []byte(s.PrivateKey),
		Scopes: []string{
			drive.DriveScope,
		},
		TokenURL: google.JWTTokenURL,
	}
	client := config.Client(context.Background())
	return client
}

// func createFolder(service *drive.Service, name string, parentId string) (*drive.File, error) {
// 	d := &drive.File{
// 		Name:     name,
// 		MimeType: "application/vnd.google-apps.folder",
// 		Parents:  []string{parentId},
// 	}

// 	file, err := service.Files.Create(d).Do()

// 	if err != nil {
// 		log.Println("Could not create dir: " + err.Error())
// 		return nil, err
// 	}

// 	return file, nil
// }

func createFile(service *drive.Service, name string, mimeType string, content io.Reader, parentId string) (*drive.File, error) {
	f := &drive.File{
		MimeType: mimeType,
		Name:     name,
		Parents:  []string{parentId},
	}
	file, err := service.Files.Create(f).Media(content).Do()

	if err != nil {
		log.Println("Could not create file: " + err.Error())
		return nil, err
	}

	return file, nil
}

func main() {
	// Step 1: Open  file
	f, err := os.Open("test.txt")

	if err != nil {
		panic(fmt.Sprintf("cannot open file: %v", err))
	}

	defer f.Close()

	// Step 2: Get the Google Drive service
	client := ServiceAccount("client_secret.json")

	srv, err := drive.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve drive Client %v", err)
	}

	// Step 3: Create directory
	// dir, err := createFolder(srv, "New Folder", "root")

	// if err != nil {
	// 	panic(fmt.Sprintf("Could not create dir: %v\n", err))
	// }

	//give your folder id here in which you want to upload or create new directory
	folderId := ""

	// Step 4: create the file and upload
	file, err := createFile(srv, f.Name(), "application/octet-stream", f, folderId)

	if err != nil {
		panic(fmt.Sprintf("Could not create file: %v\n", err))
	}

	fmt.Printf("File '%s' successfully uploaded", file.Name)
	fmt.Printf("\nFile Id: '%s' ", file.Id)

}
