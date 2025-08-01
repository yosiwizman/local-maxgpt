package p2p

import (
	"os"
	"strings"
)

var logLevel = strings.ToLower(os.Getenv("MAXGPT_P2P_LOGLEVEL"))

const (
	logLevelDebug = "debug"
	logLevelInfo  = "info"
)

func init() {
	if logLevel == "" {
		logLevel = logLevelInfo
	}
}
