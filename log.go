package logrus_lfshook

import (
	"github.com/Sirupsen/logrus"
	"time"
)

const (
	defaultTimeForFlush = 5000
)

type logMsg struct {
	data    []byte
	level   logrus.Level
	resChan chan error
}

func (h *hook) start() {
	for level, _ := range h.config.levels {
		h.levels = append(h.levels, level)
	}

	h.cache.init(&h.config)
	h.logDataChan = make(chan logMsg, 128)
	h.wg.Add(1)
	go func() {
		defer h.wg.Done()
		timeFlushSecond := h.config.FlushTime
		if timeFlushSecond == 0{
			timeFlushSecond = defaultTimeForFlush
		}
		timeChan := time.After(
			time.Duration(timeFlushSecond) *
				time.Millisecond)
		for {
			select {
			case msg, ok := <-h.logDataChan:
				if !ok {
					// close
					h.flushAll()
					return
				}
				if msg.data != nil && len(msg.data) > 0 {
					h.cache.cache(msg.level, msg.data)
				} else {
					h.cache.flush(
						msg.level,
						// TODO change path by time or size
						h.config.levels[msg.level].Path)
				}
			case <-timeChan:
				h.flushAll()
			}
		}

	}()
}

func (h *hook)flushAll(){
	for _, l := range h.levels {
		h.cache.flush(l,
			// TODO change path by time or size
			h.config.levels[l].Path)
	}
}

func (h *hook) Close() {
	close(h.logDataChan)
	h.wg.Wait()
}

func (h *hook) Flush() {
	resChan := make(chan error, 1)
	h.logDataChan <- logMsg{
		data:    nil,
		resChan: resChan,
	}
	<-resChan
}

func (h *hook) Reload(cfg HookConfig) {
	h.Close()
	h.config = cfg
	h.start()
}
