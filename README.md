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


### Middleware

```go
	e := echo.New()
    secured := e.Group("/secure")
    secured.Use(keycloak_middleware.TokenAuthMiddleware(resultConfig))
```

