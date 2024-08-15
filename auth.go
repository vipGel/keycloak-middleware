package keycloak_middleware

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/Nerzal/gocloak/v13"
	"github.com/golang-jwt/jwt/v4"
)

func ValidateToken(
	ctx context.Context,
	config *Config,
	tokenString string,
) (*jwt.Token, error) {
	certs, err := config.GocloakClient.GetCerts(ctx, config.Realm)
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(tokenString, keyFunc(*certs.Keys))
	if err != nil || !token.Valid {
		return nil, err
	}

	//claims, ok := token.Claims.(jwt.MapClaims)
	//if !ok || !token.Valid {
	//	return nil, nil, fmt.Errorf("invalid token")
	//}

	return token, nil
}

// GetToken
// Need to enable "Client authentication" and "Service accounts roles" marks in client settings to make it work
func GetToken(
	ctx context.Context,
	config *Config,
) (string, error) {
	token, err := config.GocloakClient.LoginClient(ctx, config.ClientID, config.ClientSecret, config.Realm)
	if err != nil {
		return "", err
	}
	return token.AccessToken, nil
}

func keyFunc(keys []gocloak.CertResponseKey) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		kid := token.Header["kid"].(string)
		for _, cert := range keys { // Dereference the pointer to slice
			if cert.Kid == nil || *cert.Kid != kid {
				continue
			}
			if cert.X5c == nil || len(*cert.X5c) <= 0 {
				continue
			}
			certPEM := fmt.Sprintf(
				"-----BEGIN CERTIFICATE-----\n%s\n-----END CERTIFICATE-----",
				(*cert.X5c)[0],
			)
			block, _ := pem.Decode([]byte(certPEM))
			if block == nil {
				return nil, fmt.Errorf("failed to parse certificate PEM")
			}
			parsedCert, err := x509.ParseCertificate(block.Bytes)
			if err != nil {
				return nil, fmt.Errorf("failed to parse certificate: %v", err)
			}
			return parsedCert.PublicKey, nil
		}
		return nil, fmt.Errorf("unable to find appropriate key")
	}
}
