package midi

import (
	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

// initialize the logger
func init() {
	log.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	log.SetOutput(colorable.NewColorableStdout())
	log.SetLevel(logrus.DebugLevel)
}
