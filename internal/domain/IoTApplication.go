package domain

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/google/uuid"
)

type IoTApplication struct {
	Id          string       `json:"id"`
	Name        string       `json:"name"`
	UserId      string       `json:"userId"`
	Credentials *Credentials `json:"credentials"`
}

func NewIoTApplication(userId string, appName string) (*IoTApplication, error) {
	creds, err := generateApplicationCredentials()
	if err != nil {
		return nil, err
	}
	return &IoTApplication{
		Id:          uuid.New().String(),
		Name:        appName,
		UserId:      userId,
		Credentials: creds,
	}, nil
}

func generateApplicationCredentials() (*Credentials, error) {
	clientID := uuid.New().String()
	clientSecretBytes := make([]byte, 16)
	_, err := rand.Read(clientSecretBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to generate random bytes: %w", err)
	}
	clientSecret := base64.StdEncoding.EncodeToString(clientSecretBytes)

	return &Credentials{
		ClientId:     clientID,
		ClientSecret: clientSecret,
	}, nil
}

type Credentials struct {
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}
