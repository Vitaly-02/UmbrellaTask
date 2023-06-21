package tools

import "log"

func MinorError(err error, msg string) bool {
	if err != nil {
		log.Printf("%s: %s\n", msg, err)
		return true
	}
	return false
}

func FatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
