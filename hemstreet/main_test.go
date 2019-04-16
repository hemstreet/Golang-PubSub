package hemstreet

import (
	"UnionStation/hemstreet/bar"
	"UnionStation/hemstreet/foo"
	"UnionStation/hemstreet/hub"
	"UnionStation/hemstreet/models"
	"fmt"
	"log"
	"testing"
)

func TestPubSub(t *testing.T) {
	fmt.Println(hub.Hub)

	barSub := bar.NewSubscriber(&bar.Sub{
		ID:    "bar",
		Topic: "bar-topic",
	})

	fooSub := foo.NewSubscriber(&foo.Sub{
		ID:    "foo",
		Topic: "foo-topic",
	})

	barPub := bar.NewPublisher(&bar.Pub{
		ID:    "bar",
		Topic: "bar-topic",
	})

	fooPub := foo.NewPublisher(&foo.Pub{
		ID:    "foo",
		Topic: "foo-topic",
	})

	if err := barPub.Publish(models.A{
		Message: "bar",
		Test: "tings",
	}); err != nil {
		log.Println("error from barPub")
	}

	if err := fooPub.Publish(models.A{
		Message: "foo",
		Test: "tings",
	}); err != nil {
		log.Println("error from fooPub")
	}

	barSub.StopListening()

	if err := barPub.Publish(models.A{
		Message: "bar",
		Test: "things",
	}); err != nil {
		log.Println("error from barPub")
	}

	fooSub.StopListening()

	if err := fooPub.Publish(models.A{
		Message: "foo",
		Test: "things",
	}); err != nil {
		log.Println("error from fooPub")
	}
}

func TestPubSubMultiple(t *testing.T) {

	barSub := bar.NewSubscriber(&bar.Sub{
		ID:    "bar",
		Topic: "bar-topic",
	})

	barSub2 := bar.NewSubscriber(&bar.Sub{
		ID:    "bar2",
		Topic: "bar-topic",
	})

	barSub3 := bar.NewSubscriber(&bar.Sub{
		ID:    "bar3",
		Topic: "bar-topic",
	})


	barPub := bar.NewPublisher(&bar.Pub{
		ID:    "bar",
		Topic: "bar-topic",
	})

	if err := barPub.Publish(models.A{
		Message: "bar",
		Test: "tings",
	}); err != nil {
		log.Println("error from barPub")
	}

	barSub.StopListening()
	barSub2.StopListening()

	if err := barPub.Publish(models.A{
		Message: "bar",
		Test: "tings",
	}); err != nil {
		log.Println("error from barPub")
	}

	barSub3.StopListening()
}