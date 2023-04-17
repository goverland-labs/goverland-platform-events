package examples

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/nats-io/nats.go"

	"github.com/goverland-labs/platform-events/events"
	"github.com/goverland-labs/platform-events/pkg/natsclient"
)

func main() {
	conn, err := nats.Connect(
		nats.DefaultURL,
		nats.RetryOnFailedConnect(true),
		nats.MaxReconnects(10),
		nats.ReconnectWait(time.Second),
	)
	if err != nil {
		log.Fatal("connect:", err)
	}
	defer conn.Close()

	subject := events.SubjectProposalCreated
	pr, err := natsclient.NewProducer(conn, subject)
	if err != nil {
		log.Fatal("new producer:", err)
	}

	go func() {
		for i := 0; i < 5; i++ {
			pl := events.ProposalPayload{ID: fmt.Sprintf("id-%d", i), Title: fmt.Sprintf("title #%d", i*10)}
			if err = pr.PublishJSON(context.Background(), pl); err != nil {
				log.Fatal("publish:", err)
			}
		}
	}()

	var handler events.ProposalHandler = func(payload events.ProposalPayload) error {
		fmt.Printf("message from nats: %s / %s \n", payload.ID, payload.Title)

		return nil
	}

	group := "example"
	con, err := natsclient.NewConsumer(context.Background(), conn, group, subject, handler)
	if err != nil {
		log.Fatal("new consumer:", err)
	}

	defer func() {
		err = con.Close()
		if err != nil {
			log.Fatal("consumer close:", err)
		}
	}()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		<-time.After(time.Second * 2)

		fmt.Println("wait is done")
		wg.Done()
	}(&wg)
	wg.Wait()

	log.Printf("all done")
}
