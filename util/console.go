package util

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

func Info(s string) {
	color.Cyan(logPrefix()+"%s\n", s)
}

func Error(s string, err error) {
	color.Red(logPrefix()+"%s: %e\n", s, err)
}

func Print(s string, fg color.Attribute) {
	color.New(fg).Printf(logPrefix()+"%s\n", s)
}

func logPrefix() string {
	return fmt.Sprintf("[KiKiBot %s] ", time.Now().Format("15:04:05"))
}
