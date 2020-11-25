package db

import (
	"testing"
)

func TestConnection(t *testing.T) {
	if mongoConnection != nil {
		t.Errorf("Connection needs to be init first")
	}
}

func TestGettingConnection(t *testing.T) {
	connection := GetConnection()
	if connection == nil {
		t.Errorf("Connection needs a value after init")
	}
}
