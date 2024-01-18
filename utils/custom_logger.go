package utils

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

const logDateFormat = "2006-01-02 15:04:05"
const httpLogDateFormat = "2024-01-16T14:52:36.326"

// type CustomHTTPLogger struct{}

// func NewCustomHTTPLogger() *CustomHTTPLogger {
// 	return &CustomHTTPLogger{}
// }

// func (c *CustomHTTPLogger) Log() {}

func CustomHTTPLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()
		err := next(c)
		timestamp := time.Now().Format(httpLogDateFormat)
		logEntry := fmt.Sprintf("[%s] [%d] %s - %s",
			timestamp,
			c.Response().Status,
			c.Request().RequestURI,
			time.Since(start))

		fmt.Println(logEntry)

		return err

	}
}

// CustomLogger is a custom logger that logs date, package name and log levels
type CustomLogger struct {
	// logger *log.Logger
	echo.Logger
	packageName string
}

func NewCustomLogger(packageName string) *CustomLogger {
	return &CustomLogger{
		// logger: log.New(os.Stdout, "", 0),
		Logger:      echo.New().Logger,
		packageName: packageName,
	}
}

func (c *CustomLogger) Output() io.Writer {
	return os.Stdout
}

// // Log logs date, package name and log levels and messages
// func (c *CustomLogger) Log(packageName, level, message string) {
// 	logEntry := strings.Join([]string{
// 		level,
// 		packageName,
// 		message,
// 	}, " ")

// 	c.logger.Printf("[%s] %s\n", logDateFormat, logEntry)
// }

// Log logs date, package name and log levels and messages
// the message can be of any type
func (c *CustomLogger) Log(msg string) {
	packageName := c.packageName
	timestamp := time.Now().Format(logDateFormat)
	logLevel := "INFO"
	logEntry := fmt.Sprintf("[%s] [%s] %s - %s",
		timestamp,
		logLevel,
		packageName,
		msg)

	c.Logger.Output().Write([]byte(logEntry + "\n"))
}
