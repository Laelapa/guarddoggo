package guarddoggo

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"time"
)

type rt struct {
	// tokenSize determines how many bytes of randomness to fetch (min 16, recommended 32+)
	tokenSize int
	// lifetime determines how long the refresh token should remain valid, e.g. 7 * 24 * time.Hour for a week
	lifetime time.Duration
}

// Fetch commands your trained doggo to fetch you a brand new refresh token.
// If something goes wrong, your doggo will fetch you the error.
func (s *rt) Fetch() (rt string, err error) {

	rtBytes := make([]byte, s.tokenSize)

	_, err = rand.Read(rtBytes)
	if err != nil {
		return "", errors.New("error while generating refresh token: " + err.Error())
	}

	rt = hex.EncodeToString(rtBytes)
	return rt, nil
}

// validateConfig ensures your doggo's RT training parameters are reasonable.
// Checks for minimum token size and positive lifetime values.
func (s *rt) validateConfig() error {

	var errs []error

	if s.tokenSize < 16 {
		errs = append(errs, errors.New("the refresh token length should be at least 16 bytes, ideally 32 bytes or larger"))
	}
	if s.lifetime <= 0 {
		errs = append(errs, errors.New("refresh token lifetime must be a positive value"))
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}
