package module

import "go.uber.org/zap"

func T() {
	zap.L().Info("hello world")
}
