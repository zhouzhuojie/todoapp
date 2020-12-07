package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	a := NewAppWithDefaults()
	logrus.Fatal(a.Start())
}
