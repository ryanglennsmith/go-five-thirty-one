package googleapi

import (
	"context"
	"fmt"
	"go-five-thirty-one/internal/util"
	"io"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

type DriveService struct {
	SecretsPath string
}

func NewDriveService(secretsPath string) *DriveService {
	return &DriveService{
		SecretsPath: secretsPath,
	}
}
func (ds *DriveService) DownloadFile(fileID string, filePath string) error {
	done := util.StartLoadingIndicator()
	driveService, err := ds.getService()
	if err != nil {
		return fmt.Errorf("unable to get drive service: %v", err)
	}

	remoteFile, err := driveService.Files.Export(fileID, "text/csv").Download()
	if err != nil {
		return fmt.Errorf("unable to download file: %v", err)
	}

	defer remoteFile.Body.Close()

	localFile, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("unable to create file: %v", err)
	}

	defer localFile.Close()

	_, err = io.Copy(localFile, remoteFile.Body)
	if err != nil {
		return fmt.Errorf("unable to copy file: %v", err)
	}
	done <- true
	return nil
}

func (ds *DriveService) UpdateFile(fileID string, localPath string) error {
	driveService, err := ds.getService()
	if err != nil {
		return fmt.Errorf("unable to get drive service: %v", err)
	}

	file, err := os.Open(localPath)
	if err != nil {
		return fmt.Errorf("unable to open file: %v", err)
	}

	_, err = driveService.Files.Update(fileID, nil).Media(file).Do()
	if err != nil {
		return fmt.Errorf("unable to update file: %v", err)
	}

	return nil
}

func (ds *DriveService) getService() (*drive.Service, error) {
	ctx := context.Background()
	path := ds.SecretsPath + "/credentials.json"
	byteArr, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to read client secret file: %v", err)
	}
	//drive scope is the big-bang
	config, err := google.ConfigFromJSON(byteArr, drive.DriveScope)
	if err != nil {
		return nil, fmt.Errorf("unable to parse client secret file to config: %v", err)
	}
	client := ds.getClient(config)

	driveService, err := drive.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve drive client: %v", err)
	}
	return driveService, nil
}


func (ds *DriveService) getClient(config *oauth2.Config) *http.Client {
	tokenFile := ds.SecretsPath + "/token.json"
	token, err := getTokenFromFile(tokenFile)
	if err != nil {
		token = getTokenFromWeb(config)
		saveToken(tokenFile, token)
	}
	return config.Client(context.Background(), token)
}







