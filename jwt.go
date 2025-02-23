package guarddoggo

import (
	"errors"
	"fmt"
	"time"

	gjwt "github.com/golang-jwt/jwt/v5"
)

type jwt struct {
	secret   string
	issuer   string
	lifetime time.Duration // e.g. 7 * 24 * time.Hour for a week
}

// Fetch commands your trained doggo to go fetch you a new JWT.
// Give it a userID and it'll return a signed token ready for use.
// If something goes wrong, your doggo will whimper back an error.
func (s *jwt) Fetch(userID string) (signedToken string, err error) {

	// TODO: consider the best format for userID or leave unopinionated string

	claims := gjwt.RegisteredClaims{
		Issuer:    s.issuer,
		IssuedAt:  gjwt.NewNumericDate(time.Now().UTC()),
		ExpiresAt: gjwt.NewNumericDate(time.Now().UTC().Add(s.lifetime)),
		Subject:   userID,
	}

	token := gjwt.NewWithClaims(gjwt.SigningMethodHS256, claims)

	// TODO: Asymmetric option

	signedToken, err = token.SignedString([]byte(s.secret))
	if err != nil {
		return "", errors.New("failed to sign token: " + err.Error())
	}

	return signedToken, nil
}

// Sniff lets your doggo inspect a JWT token for validity.
// If the token is valid, it'll retrieve the subject (userID) for you.
// If something smells fishy, your doggo will bark back an error.
func (s *jwt) Sniff(tokenStr string) (subject string, err error) {
	token, err := gjwt.ParseWithClaims(tokenStr, &gjwt.RegisteredClaims{}, func(token *gjwt.Token) (interface{}, error) {

		if token.Method != gjwt.SigningMethodHS256 {
			return nil, fmt.Errorf("wrong signing method: %v", token.Header["alg"])
		}

		// TODO: Assymetric option

		return []byte(s.secret), nil
	})
	if err != nil {
		return "", err
	}

	subject, err = token.Claims.GetSubject()
	if err != nil {
		return "", errors.New("failed to extract the 'subject' claim from the jwt" + err.Error())
	}

	return subject, nil
}

// validateConfig makes sure your doggo's JWT training parameters make sense.
// Checks for proper secret length and positive lifetime values.
func (s *jwt) validateConfig() error {

	var errs []error

	if len(s.secret) < 16 {
		errs = append(errs, errors.New("jwt secret must be at least 16 characters, ideally >= 32"))
	}
	if s.lifetime <= 0 {
		errs = append(errs, errors.New("jwt lifetime must be a positive value"))
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}
