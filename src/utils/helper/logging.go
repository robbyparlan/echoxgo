package utils

import (
	"bytes"
	ctx "echoxgo/src/utils/constants"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/sirupsen/logrus"
)

// Initialize the logger
var Logger *logrus.Logger
var LogWriteFile io.Writer

func init() {
	// Create a new logger instance
	Logger = logrus.New()

	if !strings.HasSuffix(os.Args[0], ".test") && (len(os.Args) <= 1 || strings.Contains(os.Args[1], "test.run")) {
		LogWriteFile = io.MultiWriter(os.Stdout, LogFile())

		Logger.SetOutput(LogWriteFile)

		Logger.SetFormatter(&logrus.JSONFormatter{})
		Logger.SetLevel(logrus.InfoLevel)
	}
}

func LogFile() *os.File {
	filename := ctx.FolderLog + ctx.PrefixLogFilename()

	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err.Error())
	}
	return f
}

// Custom response writer to capture response data
type ResponseWriter struct {
	http.ResponseWriter
	body *bytes.Buffer
}

func (rw *ResponseWriter) Write(p []byte) (int, error) {
	rw.body.Write(p)
	return rw.ResponseWriter.Write(p)
}

func LoggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get Content-Type
		contentType := c.Request().Header.Get("Content-Type")

		if !strings.HasPrefix(contentType, "multipart/") {
			// Capture the request body if not a file upload
			var bodyBytes []byte
			if c.Request().Body != nil {
				bodyBytes, _ = io.ReadAll(c.Request().Body)

				// Unmarshal the body to a map
				var bodyMap map[string]interface{}
				if err := json.Unmarshal(bodyBytes, &bodyMap); err != nil {
					// If there's an error unmarshaling, just use the raw body
					bodyMap = map[string]interface{}{
						"raw": string(bodyBytes),
					}
				}

				// Marshal the map to a JSON string
				bodyJSON, _ := json.Marshal(bodyMap)

				// Log the request
				Logger.WithFields(logrus.Fields{
					"method": c.Request().Method,
					"uri":    c.Request().URL.Path,
					"query":  c.Request().URL.RawQuery,
					"body":   string(bodyJSON),
				}).Info("Request")

				// Restore the body to the request
				c.Request().Body = io.NopCloser(bytes.NewReader(bodyBytes))
			}
		} else {
			// Parse multipart form data
			if err := c.Request().ParseMultipartForm(32 << 20); err != nil {
				Logger.WithFields(logrus.Fields{
					"error": err.Error(),
				}).Error("Failed to parse multipart form")
				return err
			}

			// Iterate over file headers
			for _, fileHeaders := range c.Request().MultipartForm.File {
				for _, fileHeader := range fileHeaders {
					Logger.WithFields(logrus.Fields{
						"method":   c.Request().Method,
						"uri":      c.Request().URL.Path,
						"query":    c.Request().URL.RawQuery,
						"body":     "file upload detected, body not logged",
						"filename": fileHeader.Filename,
					}).Info("Request")
				}
			}
		}

		// Capture the response
		rw := &ResponseWriter{
			ResponseWriter: c.Response().Writer,
			body:           new(bytes.Buffer),
		}
		c.Response().Writer = rw

		// Call the next middleware/handler
		err := next(c)
		if err != nil {
			Logger.WithFields(logrus.Fields{
				"error": err.Error(),
			}).Error("Response")
			return err
		}

		// Log the response
		Logger.WithFields(logrus.Fields{
			"status": c.Response().Status,
			"body":   rw.body.String(),
		}).Info("Response")

		return nil
	}
}
