package logger

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
)

// Custom logger with module name, line number, and date
var Logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

// Log function with module name, line number, and date
func Log(message string) {
	moduleName := callingModuleName()
	_, file, line, _ := runtime.Caller(1)
	Logger.Printf("[%s:%d] %s\n", moduleName, line, message)
}

// callingModuleName retrieves the calling module's name
func callingModuleName() string {
	pc, _, _, ok := runtime.Caller(2)
	if !ok {
		return "unknown"
	}
	fullPath := runtime.FuncForPC(pc).Name()
	// The module path is the part before the last occurrence of "/"
	lastSlashIndex := strings.LastIndex(fullPath, "/")
	if lastSlashIndex == -1 {
		return "unknown"
	}
	modulePath := fullPath[:lastSlashIndex]
	// The module name is the part after the last occurrence of "/"
	lastSlashIndex = strings.LastIndex(modulePath, "/")
	if lastSlashIndex == -1 {
		return "unknown"
	}
	moduleName := modulePath[lastSlashIndex+1:]
	return moduleName
}
