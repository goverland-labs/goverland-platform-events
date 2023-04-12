package natsclient

import (
	"fmt"

	"github.com/nats-io/nats-server/v2/server"
	natsserver "github.com/nats-io/nats-server/v2/test"
	"github.com/nats-io/nats.go"
)

const testPort = 8369

func runNatsServer(port int) (*server.Server, *nats.Conn, error) {
	opts := natsserver.DefaultTestOptions
	opts.Port = port
	opts.JetStream = true

	s := natsserver.RunServer(&opts)

	sUrl := fmt.Sprintf("nats://127.0.0.1:%d", port)
	nc, err := nats.Connect(sUrl)

	return s, nc, err
}
