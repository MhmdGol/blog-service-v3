package logger

import "go.uber.org/zap"

type CustomLogger struct {
	*zap.Logger
}

func NewCustomLogger() (*CustomLogger, error) {
	cfg := zap.NewDevelopmentConfig()

	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	return &CustomLogger{logger}, nil
}

func (c *CustomLogger) Close() error {
	return c.Sync()
}
