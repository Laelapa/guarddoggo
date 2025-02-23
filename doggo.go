// Package guarddoggo is your friendly neighborhood auth companion.
// It provides a playful wrapper around JWT and refresh token functionality,
// perfect for keeping your authentication tokens on a tight leash.
package guarddoggo

import (
	"errors"
	"fmt"
	"time"
)

// doggo is your faithful companion in token management.
// Can be trained in JWT creation/validation and refresh token generation.
type doggo struct {
	jwt *jwt
	rt  *rt
}

// Adopt gets you a fresh new guard doggo from the authentication shelter.
// Train it in JWT and/or RT capabilities to make it truly yours.
func Adopt() *doggo {
	return &doggo{}
}

// TrainedInJWT teaches your doggo the art of JWT handling.
// Provide a secret for signing, an issuer name, and how long the tokens should live.
// Your doggo will learn to both create new JWTs and verify existing ones.
func (d *doggo) TrainedInJWT(secret string, issuer string, lifetime time.Duration) (*doggo, error) {

	c := jwt{
		secret:   secret,
		issuer:   issuer,
		lifetime: lifetime,
	}

	if err := c.validateConfig(); err != nil {
		return d, fmt.Errorf("couldn't give doggo jwt capabilites: %w", err)
	}

	d.jwt = &c
	return d, nil
}

// TrainedInRT sends your doggo to refresh token training school.
// Specify how big the tokens should be and how long they should last.
// Your doggo will learn to generate cryptographically secure refresh tokens.
func (d *doggo) TrainedInRT(tokenSize int, lifetime time.Duration) (*doggo, error) {

	c := rt{
		tokenSize: tokenSize,
		lifetime:  lifetime,
	}

	if err := c.validateConfig(); err != nil {
		return d, fmt.Errorf("couldn't give doggo rt capabilites: %w", err)
	}

	d.rt = &c
	return d, nil
}

// JWT returns your doggo's JWT training certificate.
// If your doggo hasn't been trained in JWT handling, it will let you know.
func (d *doggo) JWT() (*jwt, error) {

	if d.jwt == nil {
		return nil, errors.New("this doggo hasn't undergone JWT training")
	}

	return d.jwt, nil
}

// RT returns your doggo's refresh token training certificate.
// If your doggo hasn't been trained in refresh tokens, it will let you know.
func (d *doggo) RT() (*rt, error) {

	if d.rt == nil {
		return nil, errors.New("this doggo isn't RT certified")
	}

	return d.rt, nil
}
