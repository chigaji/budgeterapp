package utils

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const logDateFormat = "2006-01-02 15:04:05"

type customHTTPLoggerStruct struct {
	Format           string
	CustomTimeFormat string
}

type customHTTPLogger struct {
	customHTTPLoggerStruct
	handlerFuc echo.MiddlewareFunc
}

func (c *customHTTPLogger) setCustomHTTPLogger(config ...customHTTPLoggerStruct) {

	var handlerFunc echo.MiddlewareFunc

	if config == nil {
		handlerFunc = middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format:           "[${time_custom}] [${method}] - ${uri} => ${status}\n",
			CustomTimeFormat: "2006-01-02 15:04:05.000",
		})
	} else {
		handlerFunc = middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format:           config[0].Format,
			CustomTimeFormat: config[0].CustomTimeFormat,
		})
	}
	c.handlerFuc = handlerFunc
}

func NewCustomHTTPLogger() echo.MiddlewareFunc {
	customHttpLogger := customHTTPLogger{}
	customHttpLogger.setCustomHTTPLogger()

	return customHttpLogger.handlerFuc
}

// CustomLogger is a custom logger that logs date, package name and log levels
type CustomLogger struct {
	// logger *log.Logger
	echo.Logger
	packageName string
}

func NewCustomLogger(packageName string) *CustomLogger {
	// NewCustomHTTPLogger()

	return &CustomLogger{
		// logger: log.New(os.Stdout, "", 0),
		Logger:      echo.New().Logger,
		packageName: packageName,
	}
}

func (c *CustomLogger) Output() io.Writer {
	return os.Stdout
}

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
