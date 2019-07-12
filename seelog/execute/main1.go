package main

import "go-start/seelog/util"

func main() {
	log := util.GetLoggerInstance().Logger
	log.Info("seelog test begin")

	for i := 0; i < 1; i++ {
		log.Tracef("hello seelog trace, i = %d", i)
		log.Debugf("hello seelog debug, i = %d", i)
		log.Infof("hello seelog info, i = %d", i)
		log.Warnf("hello seelog warn, i = %d", i)
		log.Errorf("hello seelog error, i = %d", i)
		log.Criticalf("hello seelog critical, i = %d", i)
	}

	log.Debug("seelog test end")
}
