// Package guarddoggo provides functions to easily work with jwt and refresh tokens.
// Provides only basic functionality such as token creation and validation.
// Preconfigured with basic defaults that should be good for personal/hobby projects.
package guarddoggo

import (
	"errors"
	"fmt"
	"time"
)

type doggo struct {
	jwt *jwt
	rt  *rt
}

// Get yourself a trusty guard dog. Give it a cute nickname.
func Adopt() *doggo {
	return &doggo{}
}

// Train your doggo to fetch fresh JWTs and sniff out rotten ones.
// If your instructions to the dog trainer make no sense you might get an error back, as well as your untrained doggo.
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

// Train your doggo to fetch fresh refresh tokens.
// If your instructions to the dog trainer are nonsensical you'll get returned an error along with your untrained doggo.
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

func (d *doggo) JWT() (*jwt, error) {

	if d.jwt == nil {
		return nil, errors.New("this doggo hasn't undergone JWT training")
	}

	return d.jwt, nil
}

func (d *doggo) RT() (*rt, error) {

	if d.rt == nil {
		return nil, errors.New("this doggo isn't RT certified")
	}

	return d.rt, nil
}
