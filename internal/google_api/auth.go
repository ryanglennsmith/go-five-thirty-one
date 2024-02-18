package googleapi

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"golang.org/x/oauth2"
)
func saveToken(tokenFile string, token *oauth2.Token) {
	fmt.Printf("saving token to file %v\n", tokenFile)
	file, err := os.OpenFile(tokenFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("problem saving token to file %v", err)
	}
	defer file.Close()
	json.NewEncoder(file).Encode(token)
}

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	url := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)

	fmt.Printf("go to this link and type the auth code: %v\n", url)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatalf("bad read of auth code %v", err)
	}

	token, err := config.Exchange(context.TODO(), code)
	if err != nil {
		log.Fatalf("bad exchange of token %v", err)
	}
	return token
}

func getTokenFromFile(tokenFile string) (*oauth2.Token, error) {
	file, err := os.Open(tokenFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	token := &oauth2.Token{}
	err = json.NewDecoder(file).Decode(token)
	return token, err
}
