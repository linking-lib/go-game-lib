package async

import "github.com/sirupsen/logrus"

func pCall(fn func()) {
	defer func() {
		if err := recover(); err != nil {
			logrus.Errorf("async/pCall: Error=%v", err)
		}
	}()

	fn()
}

func Run(fn func()) {
	go pCall(fn)
}
