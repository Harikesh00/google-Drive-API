package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

func authenticateServiceAccount(ctx context.Context, serviceAccountFile string) (*drive.Service, error) {
	b, err := os.ReadFile(serviceAccountFile)
	if err != nil {
		return nil, fmt.Errorf("unable to read client secret file: %v", err)
	}

	config, err := google.JWTConfigFromJSON(b, drive.DriveFileScope)
	if err != nil {
		return nil, fmt.Errorf("unable to parse client secret file to config: %v", err)
	}

	client := config.Client(ctx)

	return drive.NewService(ctx, option.WithHTTPClient(client))
}

func uploadFile(srv *drive.Service, filePath, folderID string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("unable to open file: %v", err)
	}
	defer file.Close()

	fileMetadata := &drive.File{
		MimeType: "application/octet-stream",
		Name:     "test.txt",
		Parents:  []string{folderID}, // Replace with the actual folder ID
	}

	driveFile, err := srv.Files.Create(fileMetadata).Media(file).Do()
	if err != nil {
		return fmt.Errorf("unable to create file: %v", err)
	}

	fmt.Printf("File ID: %s\n", driveFile.Id)
	return nil
}

func main() {
	ctx := context.Background()
	serviceAccountFile := "client_secret.json" // Replace with the correct path

	srv, err := authenticateServiceAccount(ctx, serviceAccountFile)
	if err != nil {
		log.Fatalf("Error initializing Drive service: %v", err)
	}

	filePath := "test.txt"                          // Ensure the correct file path
	folderID := "1vZLmkm7ny9euaGLlIu7FmRecQbS3JDZ3" // Replace with the actual folder ID

	if err := uploadFile(srv, filePath, folderID); err != nil {
		log.Fatalf("Error uploading file: %v", err)
	}
}
