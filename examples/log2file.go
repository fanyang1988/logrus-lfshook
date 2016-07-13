package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/fanyang1988/logrus-lfshook"
)

func main() {
	cfg := logrus_lfshook.HookConfig{}
	cfg.AddLevel(logrus.InfoLevel, "./info.log", 3000, 300)
	cfg.AddLevel(logrus.DebugLevel, "./debug.log", 3000, 300)
	hook := logrus_lfshook.NewHook(cfg)
	defer hook.Close()
	logrus.AddHook(hook)
	logrus.SetLevel(logrus.DebugLevel)
	for i := 0; i < 100000; i++ {
		logrus.WithField("c", i).
			Info("logs", "info")
		logrus.WithField("c", i).
			Debug("logs", "debug")
	}
}
