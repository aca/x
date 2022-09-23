package log_test

import (
	"testing"

	"github.com/aca/x/log"
)

func TestLog(t *testing.T) {
	log.Infof("hello")
	log.Debugf("hello")
	log.Errorf("hello")
	log.Fatalf("fatal")
	log.Errorf("hello")
}
