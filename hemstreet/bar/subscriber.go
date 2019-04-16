package bar

import (
	"fmt"
	"log"

	"UnionStation/hemstreet/hub"
	"UnionStation/hemstreet/models"
)

type Sub struct {
	ID string
	Topic string
}

func (s *Sub) GetID() string {
	return s.ID
}

func (s *Sub) GetTopic() string {
	return s.Topic
}

func (s *Sub) OnReceiveSubscriberMessage(v interface{}) {
	a, ok := v.(models.A)
	if !ok {
		// Uh Oh! type unknown
		return
	}

	log.Println(fmt.Sprintf("got message for sub by topic (%s) and id (%s): (%s)", s.ID, s.Topic, a))
}

func (s *Sub) StopListening() {
	hub.Hub.Unregister(s)
}

func NewSubscriber(s *Sub) *Sub {
	hub.Hub.Register(s)
	return s
}
