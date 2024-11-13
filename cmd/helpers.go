package cmd

import log "github.com/sirupsen/logrus"

type Err struct {
	E error
}

func (e *Err) Handle() {
	if e.E != nil {
		log.Error(e.E)
	}
}

func (e *Err) HandleWarn() {
	if e.E != nil {
		log.Warn(e.E)
	}
}

func (e *Err) HandleError() {
	if e.E != nil {
		log.Error(e.E)
	}
}

func (e *Err) HandleFatal() {
	if e.E != nil {
		log.Fatal(e.E)
	}
}
