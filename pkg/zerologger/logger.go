package zerologger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/luyasr/mpush/pkg/utils"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"path"
	"runtime/debug"
	"strings"
	"time"
)

var (
	Console = NewConsoleLog()
)

type LogWithOptions struct {
	Dir string `json:"dir"`
}

func init() {
	zerolog.TimeFieldFormat = time.DateTime
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
}

func NewConsoleLog() zerolog.Logger {
	return zerolog.New(console()).With().Timestamp().Logger()
}

func NewFileLog(name string, opts LogWithOptions) zerolog.Logger {
	return zerolog.New(file(name, opts)).With().Timestamp().Caller().Logger()
}

func NewMultiLog(name string, opts LogWithOptions) zerolog.Logger {
	multi := zerolog.MultiLevelWriter(console(), file(name, opts))
	return zerolog.New(multi).With().Timestamp().Logger()
}

func console() zerolog.ConsoleWriter {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.DateTime}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s |", i))
	}

	return output
}

func file(name string, opts LogWithOptions) *os.File {
	var err error
	var filepath string
	var filename strings.Builder

	if opts.Dir == "" {
		filepath = path.Join(utils.RootPath(), "log", name)
		err = os.MkdirAll(filepath, os.ModePerm)
	} else {
		filepath = path.Join(opts.Dir, name)
		err = os.MkdirAll(filepath, os.ModePerm)
	}
	if err != nil {
		panic(err)
	}

	filename.WriteString(time.Now().Format("2006-01-02-15"))
	filename.WriteString(".log")
	f := path.Join(filepath, filename.String())
	openFile, err := os.OpenFile(f, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
	}

	return openFile
}

// GinLogger receive gin framework default log
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		cost := time.Since(start)

		Console.Info().
			Int("status", c.Writer.Status()).
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Str("query", c.Request.URL.RawQuery).
			Str("ip", c.ClientIP()).
			Str("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()).
			Dur("cost", cost).
			Send()
	}
}

// GinRecovery Recover any panic that may occur in the project and use zerolog to record relevant logs
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					Console.Error().
						Any("errors", err).
						Str("request", string(httpRequest)).Send()
					// If the connection is dead, we can't write a status to it.
					_ = c.Error(err.(error)) // nolint: err check
					c.Abort()
					return
				}

				if stack {
					Console.Error().
						Any("errors", err).
						Str("request", string(httpRequest)).
						Str("stack", string(debug.Stack())).Send()
				} else {
					Console.Error().
						Any("errors", err).
						Str("request", string(httpRequest)).Send()
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
