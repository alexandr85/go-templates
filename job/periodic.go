package job

import "time"

type periodic struct {
	action func()
	onStop func()

	frequency      time.Duration
	runImmediately bool

	stop chan struct{}
	exit chan struct{}
}

func NewPeriodicJob(action, onStop func(), frequency time.Duration, runImmediately bool) *periodic {
	return &periodic{
		action:         action,
		onStop:         onStop,
		frequency:      frequency,
		runImmediately: runImmediately,
		stop:           make(chan struct{}),
		exit:           make(chan struct{}),
	}
}

func (j *periodic) Run() {
	if j.runImmediately {
		j.action()
	}

	ticker := time.NewTicker(j.frequency)

	for {
		select {
		case <-ticker.C:
			j.action()

		case <-j.stop:
			ticker.Stop()
			j.finalize()
			j.exit <- struct{}{}
			return
		}
	}
}

func (j *periodic) Stop() {
	j.stop <- struct{}{}
	<-j.exit
}

func (j *periodic) finalize() {
	if j.onStop != nil {
		j.onStop()
	}
}
