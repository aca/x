package log_test

import (
	"testing"

	"github.com/aca/x/log"
)

func TestLog(t *testing.T) {
	log.Infof("hello1")
	log.Debugf("hello2")
	log.Debugf("hello2")
	log.Errorf("hello")
	log.Fatalf("fatal")
	log.Errorf("hello")
}
