package guarddoggo

import (
	"errors"
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
func (d *doggo) TrainedInJWT(secret string, issuer string, lifetime time.Duration) (*doggo, error) {

	c := jwt{
		secret: secret,
		issuer: issuer,
		lifetime: lifetime,
	}

	if err := c.validate(); err != nil {
		return d, errors.New("couldn't give doggo jwt capabilites: " + err.Error())
	}

	d.jwt = &c
	return d, nil
}

func (d *doggo) TrainedInRT(tokenSize int, lifetime time.Duration) (*doggo, error) {

	c := rt{
		tokenSize: tokenSize,
		lifetime: lifetime,
	}

	if err := c.validate(); err != nil {
		return d, errors.New("couldn't give doggo rt capabilites: " + err.Error())
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
