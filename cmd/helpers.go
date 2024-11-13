package cmd

import log "github.com/sirupsen/logrus"

type Err struct {
	E error
}

func (e *Err) Handle() {
	if e.E != nil {
		log.Fatal(e.E)
	}
}
