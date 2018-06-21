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
		log.Fatal("close stmt failed. %v", err)
	}
}
