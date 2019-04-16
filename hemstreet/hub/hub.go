package hub

import (
	"fmt"
	"log"
)

type hub struct {}

type Publisher interface {
	GetID() string
	GetTopic() string
}

type Subscriber interface {
	GetID() string
	GetTopic() string
	OnReceiveSubscriberMessage(message interface{})
	StopListening()
}

var Hub *hub

func init() {
	Hub = &hub{}
}

var listeners = map[string]map[string]Subscriber{}

// Register adds interface to be notified of messages from the publisher
func (h *hub) Register(sub Subscriber) {
	_, ok := listeners[sub.GetTopic()]

	if !ok {
		listeners[sub.GetTopic()] = map[string]Subscriber{}
	}

	listeners[sub.GetTopic()][sub.GetID()] = sub

	log.Println(fmt.Sprintf("New subscriber for topic (%s) by id (%s). Current subscriber count count (%d)", sub.GetTopic(), sub.GetID(), len(listeners[sub.GetTopic()])))
}

// Unregister removes subscriber from list to be notified
func (h *hub) Unregister(sub Subscriber) {
	_, ok := listeners[sub.GetTopic()]

	if !ok {
		log.Println("subscriber not found, this is weird")
		return
	}

	delete(listeners[sub.GetTopic()], sub.GetID())

	log.Println(fmt.Sprintf("Removed subscriber for topic (%s) by id (%s). Current subscriber count count (%d)", sub.GetTopic(), sub.GetID(), len(listeners[sub.GetTopic()])))
}

// Publish emits a message for all subscribers to get
func (h *hub) Publish(pub Publisher, message interface{}) (err error) {

	ll, ok := listeners[pub.GetTopic()]

	if !ok {
		log.Println(fmt.Sprintf("No current listeners to topic %s", pub.GetTopic()))
		return nil
	}

	log.Println(fmt.Sprintf("Publish to message topic %s. %d listeners", pub.GetTopic(), len(ll)))
	for _, l := range ll {
		l.OnReceiveSubscriberMessage(message)
	}

	return nil
}