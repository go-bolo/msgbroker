package msgbroker

import "github.com/go-bolo/bolo"

type PluginCfgs struct {
	Client Client
}

func NewPlugin(cfg *PluginCfgs) *MSGBrokerPlugin {
	p := MSGBrokerPlugin{
		Name:   "msg-broker",
		Client: cfg.Client,
	}

	if cfg.Client == nil {
		panic("cfg.Client is required and should be one of: rabbitmq, nsq clients")
	}

	return &p
}

type MSGBrokerPlugin struct {
	Name   string
	Client Client
}

func (p *MSGBrokerPlugin) GetName() string {
	return p.Name
}

func (p *MSGBrokerPlugin) Init(app bolo.App) error {
	return p.Client.Init(app)
}
