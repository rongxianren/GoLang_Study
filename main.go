package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"helloworld/person"
)

func main() {

	person := person.Person{Name: "kloud", Age: 30}

	fmt.Println(person)

	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
}
