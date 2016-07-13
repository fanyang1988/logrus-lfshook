package logrus_lfshook

import (
	"github.com/Sirupsen/logrus"
	"sync"
)

type hook struct {
	config      HookConfig
	levels      []logrus.Level
	cache       logCache
	logDataChan chan logMsg
	wg          sync.WaitGroup
}

func NewHook(cfg HookConfig) *hook {
	nh := &hook{
		config: cfg,
	}
	nh.start()
	return nh
}

func (h *hook) Fire(entry *logrus.Entry) error {
	d, err := entry.String()
	if err != nil {
		return err
	}
	h.logDataChan <- logMsg{
		level: entry.Level,
		data:  []byte(d),
	}
	return nil
}

func (h *hook) Levels() []logrus.Level {
	return h.levels
}
