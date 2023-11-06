package pkg

import (
	"log"
	"os"
	"strconv"

	"github.com/memphisdev/memphis.go"
)

type GithubProducer struct {
	conn     *memphis.Conn
	producer *memphis.Producer
	l        *log.Logger
}

func NewProducer(l *log.Logger) *GithubProducer {
	memphis_acc_id, _ := strconv.Atoi(os.Getenv("MEMPHIS_ACCOUNT_ID"))
	conn, err := memphis.Connect(os.Getenv("MEMPHIS_HOST"), os.Getenv("MEMPHIS_USERNAME"),
		memphis.Password(os.Getenv("MEMPHIS_PASSWORD")), memphis.AccountId(memphis_acc_id))

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
