package v1

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/microsoft"
)

var (
	// Reemplaza estos con los valores de Azure
	ClientID     = "TU_CLIENT_ID"
	ClientSecret = "TU_CLIENT_SECRET"
	RedirectURI  = "http://localhost:8080/callback"

	Scopes = []string{
		"Files.ReadWrite",
		"offline_access",
	}

	conf = &oauth2.Config{
		ClientID:     ClientID,
		ClientSecret: ClientSecret,
		RedirectURL:  RedirectURI,
		Scopes:       Scopes,
		Endpoint:     microsoft.AzureADEndpoint("common"),
	}
)

// Abre navegador y espera el callback
func Authorize() *oauth2.Token {
	url := conf.AuthCodeURL("state123", oauth2.AccessTypeOffline)

	fmt.Println("Abre esta URL en el navegador:\n" + url)

	codeCh := make(chan string)

	// Servidor local para capturar el código
	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		fmt.Fprintf(w, "¡Código recibido! Ya puedes cerrar esta ventana.")
		codeCh <- code
	})
	go http.ListenAndServe(":8080", nil)

	code := <-codeCh

	token, err := conf.Exchange(context.Background(), code)
	if err != nil {
		log.Fatal("Error intercambiando código:", err)
	}
	return token
}


