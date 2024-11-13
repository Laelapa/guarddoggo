package guarddoggo

import "time"

type rt struct {
	tokenSize int           // in bytes, ideally >= 32
	lifetime  time.Duration // e.g. 7 * 24 * time.Hour for a week
}

func (s *rt) validate() error {

	// TODO: implement

	return nil
}
