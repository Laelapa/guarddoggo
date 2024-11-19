package guarddoggo

import (
	"errors"
	"time"
)

type rt struct {
	tokenSize int           // in bytes, ideally >= 32
	lifetime  time.Duration // e.g. 7 * 24 * time.Hour for a week
}

func (s *rt) validate() error {

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
