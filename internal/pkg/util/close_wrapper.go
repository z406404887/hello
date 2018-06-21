package util

import (
	"log"
)

type closer interface {
	Close() error
}

func Close(c closer) {
	err := c.Close()
	if err != nil {
		log.Fatalf("close stmt failed. %v", err)
	}
}
