package v1

import (
	"encoding/json"
	"os"
)

type DataClient struct {
	CLientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectURI  string `json:"redirect_uri"`
}

func NewMockDataClient() *DataClient {
	return &DataClient{
		CLientID:     "mock_client_id",
		ClientSecret: "mock_client_secret",
		RedirectURI:  "http://localhost:8080/callback",
	}
}

const DATAJSONFILE = "data.json"

func WriteEmptyFileDataClient() error {
	mock := NewMockDataClient()
	data, err := json.Marshal(mock)
	if err != nil {
		return err
	}

	file, err := os.Create(DATAJSONFILE)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func ReadDataClient() (*DataClient, error) {
	file, err := os.Open(DATAJSONFILE)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var dataClient DataClient
	err = decoder.Decode(&dataClient)
	if err != nil {
		return nil, err
	}

	return &dataClient, nil
}

func UpdateConfigFromFile() error {
	dataClient, err := ReadDataClient()
	if err != nil {
		return err
	}

	UpdateConfig(dataClient.CLientID, dataClient.ClientSecret, dataClient.RedirectURI)

	return nil
}
