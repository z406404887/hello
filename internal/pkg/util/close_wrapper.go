package util

import (
	"log"
)

type closer interface {
	Close() error
}

func Close(closer Closer) {
	err := closer.Close()
	if err != nil {
		log.Fatal("close stmt failed. %v", err)
	}
}
