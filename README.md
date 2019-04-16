Pub/Sub
===

basic implementation of pub sub pattern for golang with model sharing

The main idea is that inside the hemstreet folder, there are two modules ( foo / bar ). These
two modules want to interact with each other but not create a circular dependency. So we go through the hub/hub ( sorry for the stutter ) package
to be the mediator while using the models folder for a common package that does not import either of the two so we can have the same models.

* `hemstreet/main_test.go` sets up some scenarios
* `hemstreet/models/*` is the repo for models that need to be shared between packages
* `hemstreet/foo/*` and `hemstreet/bar/*` are the packages that wish to talk to eachother
* `hemstreet/hub/hub.go` is the exchange/broker for everything