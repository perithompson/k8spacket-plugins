package connections_log

import "log"

var LOGGER log.Logger

func BuildLogger() {
	LOGGER = *log.New(log.Writer(), "[connections plugin] ", log.LstdFlags|log.Lmsgprefix)
}
