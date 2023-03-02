package logf

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

var Logger *log.Logger
var LogFile string

func init() {
	tmpdir := filepath.Join(os.TempDir(), "logf")
	err := os.MkdirAll(tmpdir, 0o777)
	if err != nil {
		panic(err)
	}

    LogFile := filepath.Join(tmpdir, time.Now().Format("060102_150405"))

	f, err := os.Create(LogFile)
	if err != nil {
		panic(err)
	}

	Logger = log.New(f, "", log.LstdFlags|log.Lshortfile)
}

var P = Print

func Print(v ...any) {
	Logger.Print(v...)
}

var Pln = Println

func Println(v ...any) {
	Logger.Println(v...)
}

var Pf = Printf

func Printf(format string, v ...any) {
	Logger.Printf(format, v...)
}
