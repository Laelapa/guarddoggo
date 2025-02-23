# guarddoggo
A faithful Go package for JWT and refresh token management. Keeps your auth tokens on a tight leash.

> [!IMPORTANT]
> This is a **work in progress**. No useful functionality implemented yet.

## About

This project intends to be little more than a silly-themed wrapper around parts of [golang-jwt](https://github.com/golang-jwt/jwt), preconfigured with some *hopefully* sensible defaults for the basic functions that a casual/personal project might need.

## Installation
```bash
go get github.com/Laelapa/guarddoggo
```

## 🐾 Quick Start
```go
// Adopt your very own authentication guard doggo
doggo := guarddoggo.Adopt()

// Train it in JWT handling
doggo, err := doggo.TrainedInJWT(
    "your-secret-key-here",
    "your-app-name",
    24*time.Hour, // Token lifetime
)
if err != nil {
    // Handle training error
}

// Train it in refresh token generation
doggo, err = doggo.TrainedInRT(
    32,           // Token size in bytes
    7*24*time.Hour, // Token lifetime
)
if err != nil {
    // Handle training error
}

// Command it to fetch you a fresh JWT
token, err := doggo.JWT().Fetch("user123")
if err != nil {
    // Handle token creation error
}

// Have it sniff out invalid tokens
userID, err := doggo.JWT().Sniff(token)
if err != nil {
    // Handle invalid token error
}

// Command it to fetch you a new refresh token
refreshToken, err := doggo.RT().Fetch()
if err != nil {
    // Handle refresh token creation error
}
```

## 🦮 Features
- Playful, dog-themed API
- Simple JWT creation and validation
- Secure refresh token generation
- Configurable token lifetimes
- Minimal dependencies

