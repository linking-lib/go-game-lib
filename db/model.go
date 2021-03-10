package db

import (
	"github.com/linking-lib/go-game-lib/infrastructure/mysql"
)

const asyncTaskBacklog = 128

var (
	chWrite  chan interface{} // async write channel
	chUpdate chan interface{} // async update channel
)

func envInit() {
	// async task
	go func() {
		for {
			select {
			case t, ok := <-chWrite:
				if !ok {
					return
				}
				mysql.MAdd(t)

			case t, ok := <-chUpdate:
				if !ok {
					return
				}
				mysql.MSave(t)
			}
		}
	}()
}

func MustStartup() func() {
	chWrite = make(chan interface{}, asyncTaskBacklog)
	chUpdate = make(chan interface{}, asyncTaskBacklog)
	envInit()
	closer := func() {
		close(chWrite)
		close(chUpdate)
	}
	return closer
}
