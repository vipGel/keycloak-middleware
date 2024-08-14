package keycloak_middleware

import "github.com/Nerzal/gocloak/v13"

type Config struct {
	KeycloakURL   string
	Realm         string
	ClientID      string
	ClientSecret  string
	GocloakClient *gocloak.GoCloak
}

// this is a function that set config

func NewConfig(keycloakURL, realm, clientID, clientSecret string) *Config {
	return &Config{
		KeycloakURL:   keycloakURL,
		Realm:         realm,
		ClientID:      clientID,
		ClientSecret:  clientSecret,
		GocloakClient: gocloak.NewClient(keycloakURL),
	}
}
