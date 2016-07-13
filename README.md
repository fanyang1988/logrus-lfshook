# logrus-lfshook
local file log hook for logrus, support log segmentation by time or file size

## install

```
go get "github.com/fanyang1988/logrus-lfshook"
```

## usage

the hook create a gorountine, so it need close when app exit.
```
    import (
        "github.com/Sirupsen/logrus"
        "github.com/fanyang1988/logrus-lfshook"
    )

    func main() {
        // config for hook
        cfg := logrus_lfshook.HookConfig{}

        // add config for a log level
        cfg.AddLevel(
            logrus.InfoLevel, // log level
            "./info.log",     // file path for log
            3000,             // time for auto flush(ms)
            300)              // size for auto flush(kb)

        hook := logrus_lfshook.NewHook(cfg)

        // to close hook to flush before exit
        defer hook.Close()

        logrus.AddHook(hook)
        logrus.SetLevel(logrus.DebugLevel)
        logrus.WithField("c", i).Info("logs")
    }
```

## TODO
- auto flush by size
- segmentation by time or file size
