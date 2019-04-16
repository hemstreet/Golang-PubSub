package bar

import (
	"log"

	"UnionStation/hemstreet/hub"
	"UnionStation/hemstreet/models"
)

type Pub struct {
	ID string
	Topic string
}

func (p *Pub) GetID() string {
	return p.ID
}

func (p *Pub) GetTopic() string {
	return p.Topic
}

func (p *Pub) Publish(message models.A) (err error) {

	var msg interface{}
	msg = message

	if err := hub.Hub.Publish(p, msg); err != nil {
		log.Println("error during publish")
	}

	return nil
}

func NewPublisher(p *Pub) *Pub {
	return p
}
