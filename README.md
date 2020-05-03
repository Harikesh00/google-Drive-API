# Google-Drive-API
## how to upload file to google drive using golang

### Using simple command line
- step 1: Create API key from https://console.developers.google.com and download secret file.
                 
- step 2: If google drive API not enabled,enable from https://console.developers.google.com/apis/library

- step 3: give the path of secret file (eg. credential.json) in the code.

- step 4: download neccesary packages given in the code

- step 5: Run the code `go run upload.go`

### Using service account
   **Why we need serivce account**
> In server-to-server communication. when we want to upload without user intervention.
   
- step 1: To upload using service account create service account from https://console.developers.google.com/iam-admin/serviceaccounts

- step 2: Download the secret file and give the path in code.

- step 3: create a folder on google drive and share it with new service account email (client_email) you got in secret file.

     **Note:** If you will not share folder it will not visible to you afteer upload.
 
 - step 4: Run the code `go run upload.go`

## For more info 
- https://developers.google.com/drive/api/v3/quickstart/go
- https://developers.google.com/identity/protocols/oauth2/service-account
