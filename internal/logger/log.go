package syslogger

import "go.uber.org/zap"

func InitLogger(l *zap.Logger) error {
	cfg := zap.NewDevelopmentConfig()

	var err error
	l, err = cfg.Build()
	if err != nil {
		return err
	}

	return nil
}
