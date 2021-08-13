package nats_streaming

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/batchcorp/plumber/dproxy"
)

// Dynamic starts up a new GRPC client connected to the dProxy service and receives a stream of outbound replay messages
// which are then written to the message bus.
func (n *NatsStreaming) Dynamic(ctx context.Context) error {
	llog := logrus.WithField("pkg", "nats-streaming/dynamic")

	// Start up dynamic connection
	grpc, err := dproxy.New(n.Options, "Nats Streaming")
	if err != nil {
		return errors.Wrap(err, "could not establish connection to Batch")
	}

	go grpc.Start()

	// Continually loop looking for messages on the channel.
MAIN:
	for {
		select {
		case outbound := <-grpc.OutboundMessageCh:
			if err := n.stanClient.Publish(n.Options.NatsStreaming.Channel, outbound.Blob); err != nil {
				llog.Errorf("Unable to replay message: %s", err)
				break
			}

			llog.Debugf("Replayed message to NATS streaming channel '%s' for replay '%s'",
				n.Options.NatsStreaming.Channel, outbound.ReplayId)

		case <-ctx.Done():
			llog.Debug("context closed")
			break MAIN
		}
	}

	llog.Debug("exiting")

	return nil
}
