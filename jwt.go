package guarddoggo

import (
	"errors"
	"time"

	gjwt "github.com/golang-jwt/jwt/v5"
)

type jwt struct {
	secret   string
	issuer   string
	lifetime time.Duration
}

func (s *jwt) validate() error {

	// TODO: implement

	return nil
}

// Command the doggo to fetch you a brand new jwt
func (s *jwt) Fetch(userID string) (string, error) {
	claims := gjwt.RegisteredClaims{
		Issuer:    s.issuer,
		IssuedAt:  gjwt.NewNumericDate(time.Now().UTC()),
		ExpiresAt: gjwt.NewNumericDate(time.Now().UTC().Add(s.lifetime)),
		Subject:   userID,
	}

	token := gjwt.NewWithClaims(gjwt.SigningMethodHS256, claims)

	// TODO: Asymmetric option

	signedToken, err := token.SignedString([]byte(s.secret))
	if err != nil {
		return "", errors.New("failed to sign token: " + err.Error())
	}

	return signedToken, nil
}
