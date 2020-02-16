package main

import (
	"math/rand"
	"os"
	"strconv"
	"testing"
)

func randPort() string {
	return strconv.Itoa(rand.Intn(10000))
}

func TestGetOrElseReturnsDefault(t *testing.T) {
	var defaultValue = randPort()
	var port = getOrElse("PORT", defaultValue)
	if port != defaultValue {
		t.Fail()
	}
}

func TestGetOrElseDelegatesToLookupEnv(t *testing.T) {
	var expectedPort = randPort()
	os.Setenv("PORT", expectedPort)
	var port = getOrElse("PORT", randPort())
	if port != expectedPort {
		t.Fail()
	}
}
