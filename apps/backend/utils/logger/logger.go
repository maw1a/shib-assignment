package logger

import (
	"fmt"
	"io"
	"strings"

	"github.com/gin-gonic/gin"
)

var Writer io.Writer = gin.DefaultWriter

func logger(level string, format string, values ...any) {
	if gin.IsDebugging() {
		if !strings.HasSuffix(format, "\n") {
			format += "\n"
		}
		fmt.Fprintf(Writer, fmt.Sprintf("[GIN-%s] ", level)+format, values...)
	}
}

func Info(format string, values ...any) {
	logger("info", format, values...)
}

func Debug(format string, values ...any) {
	logger("debug", format, values...)
}

func Error(format string, values ...any) {
	logger("error", format, values...)
}
