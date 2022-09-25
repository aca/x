package noti_test

import (
	"testing"

	"github.com/aca/x/noti"
)

func TestSend(t *testing.T) {
    noti.Send("test")
}
