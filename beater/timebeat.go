package beater

import (
	"fmt"
	"time"

	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/logp"

	"github.com/davdjl/timebeat/config"
)

// timebeat configuration.
type timebeat struct {
	done   chan struct{}
	config config.Config
	client beat.Client
}

// New creates an instance of timebeat.
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	c := config.DefaultConfig
	if err := cfg.Unpack(&c); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &timebeat{
		done:   make(chan struct{}),
		config: c,
	}
	return bt, nil
}

// Run starts timebeat.
func (bt *timebeat) Run(b *beat.Beat) error {
	logp.Info("timebeat is running! Hit CTRL-C to stop it.")

	var err error
	bt.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}

	ticker := time.NewTicker(bt.config.Period)
	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}

		event := beat.Event{
			Timestamp: time.Now(),
			Fields: common.MapStr{
				"type":    b.Info.Name,
				"files": getTime(bt.config.Path),
			},
		}
		bt.client.Publish(event)
		logp.Info("Event sent")
	}
}

// Stop stops timebeat.
func (bt *timebeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
