package logrus_lfshook

import (
	"github.com/Sirupsen/logrus"
	"os"
	"errors"
)

var (
	ErrUnConfigLevel = errors.New("ErrUnConfigLevel")
)

const (
	kb        = 1024
	cacheSize = 32 * kb
)

type logCache struct {
	logData map[logrus.Level][]byte
}

func (l *logCache) init(config *HookConfig) {
	l.logData = make(map[logrus.Level][]byte,
		len(config.levels))

	for level, _ := range config.levels {
		l.logData[level] = make([]byte, 0, cacheSize)
	}
}

func (l *logCache)cache(level logrus.Level, d []byte){
	data, ok := l.logData[level]
	if !ok{
		return
	}
	data = append(data, d...)
	l.logData[level] = data
}

func (l *logCache) flush(level logrus.Level, path string)error{
	data, ok := l.logData[level]
	if !ok{
		return ErrUnConfigLevel
	}

	fd, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		l.logData[level] = data[0:0]
		return err
	}
	defer fd.Close()
	_, err = fd.Write(data)
	l.logData[level] = data[0:0]
	return err
}