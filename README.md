# Message Broker / queues for Go Bolo Framework

## Congiguration

### NSQ:

- NSQ_ADDR "127.0.0.1:4150"
- NSQ_LOOKUPD_ADDR "127.0.0.1:4161"

### RabbitMQ

## Usage

### With nsq services

main.go
```go
package main

import (
	"fmt"
	"os"

	"github.com/go-bolo/bolo"
	"github.com/sirupsen/logrus"
	"gitlab.com/www.monitordomercado.com.br/mm/msgbroker"
	"gitlab.com/www.monitordomercado.com.br/mm/msgbroker_nsq"

)

func main() {
	logrus.Debug("Starting main")

	var err error

	app := bolo.Init()

  // Register the NSQ plugin to load its feature in the app
	app.RegisterPlugin(msgbroker.NewPlugin(&msgbroker.PluginCfgs{
		Client: msgbroker_nsq.NewNSQClient(&msgbroker_nsq.NSQClientCfg{}),
	}))
	app.RegisterPlugin(project.NewPlugin(&project.PluginCfgs{}))

	err = app.Bootstrap()
	handleExecutionError(err)

	cmd.Execute()

	handleExecutionError(nil)
}

func handleExecutionError(err error) {
	if err != nil {
		os.Exit(1)
	}
}
```

## TODO!

- add tests
