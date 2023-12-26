package utils

import "log"

func Check(e error) {
	if e != nil {
			log.Fatalf("An error occurred: %v", e)
	}
}
