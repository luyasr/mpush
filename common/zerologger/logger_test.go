package zerologger

import (
	"errors"
	"testing"
)

func TestNewConsoleLog(t *testing.T) {
	Console.Error().Stack().Err(errors.New("404")).Msg("hello world")
}
