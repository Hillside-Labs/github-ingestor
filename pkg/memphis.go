package pkg

import (
	"log"

	"github.com/memphisdev/memphis.go"
)

type GithubProducer struct {
	conn     *memphis.Conn
	producer *memphis.Producer
	l        *log.Logger
}

func NewProducer(l *log.Logger) *GithubProducer {
	conn, err := memphis.Connect("aws-us-east-1.cloud.memphis.dev", "github_ingestor", memphis.Password("mh2-JK69M5"), memphis.AccountId(223674564))
	if err != nil {
		l.Println("Producer failed to connect:", err.Error())
	}

	p, err := conn.CreateProducer("github-events", "github_ingestor")
	if err != nil {
		l.Println("Producer failed to start:", err.Error())
	}
	return &GithubProducer{conn: conn, producer: p, l: l}
}

func (gp *GithubProducer) PushEvent(message interface{}) {
	err := gp.producer.Produce(message, memphis.AsyncProduce())
	if err != nil {
		gp.l.Printf("Produce failed: %v\n", err)
	}
}

func (gp *GithubProducer) Close() {
	gp.conn.Close()
}
