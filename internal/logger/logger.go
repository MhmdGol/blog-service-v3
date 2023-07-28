package logger

import "go.uber.org/zap"

func InitLogger() (*zap.Logger, error) {
	return zap.Must(zap.NewDevelopment()), nil
}
