package zaplog

import "testing"

func TestZapLogger(t *testing.T) {
	logger := InitZapLogger()
	logger.Info("test zap logger")
	logger.Debug("test debug zap logger")
	logger.Warn("test warn zap logger")
	logger.Error("test error zap logger")
}
