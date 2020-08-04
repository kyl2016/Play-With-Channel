# logrus

## AddHook 

## graylog new hook

Usage
The hook must be configured with:

A Graylog GELF UDP address (a "ip:port" string).
an optional hash with extra global fields. These fields will be included in all messages sent to Graylog

```cassandraql
package main

import (
    log "github.com/sirupsen/logrus"
    "github.com/gemnasium/logrus-graylog-hook/v3"
    )

func main() {
    hook := graylog.NewGraylogHook("<graylog_ip>:<graylog_port>", map[string]interface{}{"this": "is logged every time"})
    log.AddHook(hook)
    log.Info("some logging message")
}
```

[logrus-graylog-hook](https://github.com/gemnasium/logrus-graylog-hook)

