package guarddoggo

import "time"

type jwt struct {
	secret   string
	issuer   string
	lifetime time.Duration
}

func (s *jwt) validate() error {

	// TODO: implement

	return nil
}
