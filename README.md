# guarddoggo
A faithful Go package for JWT and refresh token management. Keeps your auth tokens on a tight leash.

## About

This project intends to be little more than a silly-themed wrapper around parts of [golang-jwt](https://github.com/golang-jwt/jwt), preconfigured with some *hopefully* sensible defaults for the basic functions that a casual/personal project might need.

## 🦮 Features
- Simple JWT creation and validation ([golang-jwt/SigningMethodHS256](https://pkg.go.dev/github.com/golang-jwt/jwt/v5@v5.2.1#SigningMethodHS256))
- Refresh token generation
- Basic verification that token configuration is valid
- Minimal dependencies

## Installation
1. Pull the package with:
```bash
go get github.com/Laelapa/guarddoggo
```
2. then import it in your code with:
```go
import "github.com/Laelapa/guarddoggo"
```

## 🐾 Quick Start

Adopt your very own guard doggo,
even consider giving it an affectionate nickname:
```go
azor := guarddoggo.Adopt()
```

Train it in JWT handling:
```go
azor, err := azor.TrainedInJWT(
    "your-jwt-secret-here",
    "your-app-name",
    24*time.Hour, // Example token lifetime
)
if err != nil {
    // Handle training error
}
```

Command it to fetch you a fresh JWT:
```go
token, err := azor.JWT().Fetch("user123")
if err != nil {
    // Handle token creation error
}
```

Have it sniff out spoiled jwts:
```go
userID, err := azor.JWT().Sniff(token)
if err != nil {
    // Handle invalid token error
}
```

Train it in refresh token hunting:
```go
azor, err = azor.TrainedInRT(
    32,             // Example token size in bytes
    7*24*time.Hour, // Example token lifetime
)
if err != nil {
    // Handle training error
}
```


Command it to fetch you a new refresh token:
```go
refreshToken, err := azor.RT().Fetch()
if err != nil {
    // Handle refresh token creation error
}
```
## TODOs

Stuff that might be added in the future if I end up needing it in my other projects:
- Custom claims for jwts
- Assymetric signing option
- More utility functions
