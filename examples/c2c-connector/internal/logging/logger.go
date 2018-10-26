package logging

import (
	"context"
	"github.com/SmartThingsOSS/smartapp-go/pkg/smartappcore"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"time"
)

func LoggingContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		correlationId := c.GetHeader(smartappcore.CorrelationHeader)
		if correlationId == "" {
			correlationId = uuid.NewV4().String()
		}
		c.Set(smartappcore.CorrelationKey, correlationId)
		c.Next()
	}
}

func RequestLogger(timeFormat string, utc bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		end := time.Now()
		latency := end.Sub(start)
		if utc {
			end = end.UTC()
		}

		logger := Logger(c)

		entry := logger.WithFields(log.Fields{
			"status":     c.Writer.Status(),
			"method":     c.Request.Method,
			"path":       path,
			"query":      query,
			"ip":         c.ClientIP(),
			"user-agent": c.Request.UserAgent(),
			"time":       end.Format(timeFormat),
			"latency":    latency,
		})

		if len(c.Errors) > 0 {
			// Append error field if this is an erroneous request.
			entry.Error(c.Errors.String())
		} else {
			entry.Info()
		}
	}
}

// Logger creates a new Ginrus logger with a UUID included
func Logger(c context.Context) *log.Entry {
	loggingId, ok := c.Value(smartappcore.CorrelationKey).(string)
	if !ok {
		return log.StandardLogger().WithField(smartappcore.CorrelationKey, "")
	}
	return log.StandardLogger().WithField(smartappcore.CorrelationKey, loggingId)
}
