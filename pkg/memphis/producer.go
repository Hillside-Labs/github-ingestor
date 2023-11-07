package memphis

import (
	"log"

	"github.com/memphisdev/memphis.go"
)

type ProducerConfig struct {
	accountId int
	host      string
	username  string
	password  string
}

func NewProducerConfig(accountId int, host, username, pass string) *ProducerConfig {
	return &ProducerConfig{accountId: accountId, host: host, username: username, password: pass}
}

type GithubProducer struct {
	conn     *memphis.Conn
	producer *memphis.Producer
	l        *log.Logger
}

func NewProducer(pc *ProducerConfig, l *log.Logger) *GithubProducer {
	conn, err := memphis.Connect(pc.host, pc.username, memphis.Password(pc.password), memphis.AccountId(pc.accountId))

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
