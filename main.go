package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

type CustomLogger struct {
	logger *log.Logger
}

func NewCustomLogger() *CustomLogger {
	return &CustomLogger{
		logger: log.New(os.Stdout, "", 0),
	}
}

func (cl *CustomLogger) Info(message string, additionalValue interface{}) {
	_, file, line, _ := runtime.Caller(1)
	folderName := filepath.Base(filepath.Dir(file))

	logMessage := fmt.Sprintf("[%s]INFO: %s\n\tFOLDER: %s\n\tFILE: %s:%d\n\tINFO VALUE: %v",
		time.Now().Format("2006/01/02 15:04:05"), message, folderName, file, line, additionalValue)

	cl.logger.Print(logMessage)
}

func (cl *CustomLogger) Error(message string, additionalValue interface{}) {
	_, file, line, _ := runtime.Caller(1)
	folderName := filepath.Base(filepath.Dir(file))

	logMessage := fmt.Sprintf("[%s]ERROR: %s\n\tFOLDER: %s\n\tFILE: %s:%d\n\tERROR VALUE: %v",
		time.Now().Format("2006/01/02 15:04:05"), message, folderName, file, line, additionalValue)

	cl.logger.Print(logMessage)
}

func (cl *CustomLogger) InfoURL(message string, u *url.URL) {
	_, file, line, _ := runtime.Caller(1)
	folderName := filepath.Base(filepath.Dir(file))

	var address string
	if u.Host != "" {
		address = u.Host
	} else {
		address = "N/A"
	}

	scheme := "UNKNOWN"
	if u.Scheme != "" {
		scheme = strings.ToUpper(u.Scheme)
	}

	logMessage := fmt.Sprintf("[%s]INFO: %s\n\tURL: %s\n\tADDRESS: %s\n\tFOLDER: %s\n\tFILE: %s:%d\n\tSCHEME: %s",
		time.Now().Format("2006/01/02 15:04:05"), message, u.String(), address, folderName, file, line, scheme)

	cl.logger.Print(logMessage)
}
