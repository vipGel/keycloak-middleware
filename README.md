This module base on: [Keycloak Secured Go Application](https://github.com/manimovassagh/go-keycloack/tree/main)

# Keycloak Secured Go Module

This module provides keycloak token validation and middleware to Echo framework.

## Features

- Token validation
- Token auth middleware
- Client auth middleware

## Usage

### Installation

```shell
go get github.com/vipGel/keycloak-middleware
```

### Importing

```go
import "github.com/Nerzal/gocloak/v13"
```


### Config

```go
keycloak_middleware.NewConfig(
    // Url of deployed Keycloak. Example "http://localhost:8080"
    KeycloakURL,
    // Name of the Keycloak Realm
    Realm,
    // Name of the Keycloak Client id
    ClientID,
    // Secret key of the Client. Optional.
    // Can be found in "Credentials" tab of client
    // if "Client authentication" is enabled.
    ClientSecret, 
)
```

### Middleware

```go
e := echo.New()
secured := e.Group("/secure")
secured.Use(keycloak_middleware.TokenAuthMiddleware(resultConfig))
```

