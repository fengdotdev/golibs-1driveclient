package v1

import (
	"context"
	"encoding/json"
	"fmt"

	"golang.org/x/oauth2"
)

func ListRootFiles(token *oauth2.Token) error {
	client := conf.Client(context.Background(), token)

	resp, err := client.Get("https://graph.microsoft.com/v1.0/me/drive/root/children")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}

	fmt.Println("Archivos del root:")
	for _, item := range result["value"].([]interface{}) {
		file := item.(map[string]interface{})
		fmt.Println(" -", file["name"])
	}
	return nil
}
