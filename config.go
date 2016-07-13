package logrus_lfshook

import (
	"github.com/BurntSushi/toml"
	. "github.com/Sirupsen/logrus"
)

// HookConfig config info for hook, can load as toml
type HookConfig struct {
	FlushTime  int64 // ms
	FlushCount int64
	Levels     map[string]levelConfig
	levels     map[Level]levelConfig
}

type levelConfig struct {
	Path            string
	TimeIntervalSeg int64 // ms
	SizeIntervalSeg int64 // kb
}

// LoadFromToml load config info by "github.com/BurntSushi/toml"
func (h *HookConfig) LoadFromToml(datas []byte) error {
	err := toml.Unmarshal(datas, h)
	if err != nil {
		return err
	}

	// build levels,
	// conv level from string to logrus.Level
	h.levels = make(map[Level]levelConfig,
		len(h.Levels))
	for levelStr, info := range h.Levels {
		level, err := ParseLevel(levelStr)
		if err != nil {
			return err
		}
		h.levels[level] = info
	}

	return nil
}

// AddLevel add config info for a log level
func (h *HookConfig) AddLevel(
	level Level,
	path string,
	time4seg int64,
	size4seg int64) {
	if h.levels == nil {
		h.levels = make(map[Level]levelConfig, 8)
		h.Levels = make(map[string]levelConfig, 8)
	}
	info := levelConfig{
		path,
		time4seg,
		size4seg,
	}
	h.Levels[level.String()] = info
	h.levels[level] = info
}
