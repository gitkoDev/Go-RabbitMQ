package helpers

import "log"

func FailOnError( msg string, err error) {
	if err != nil {
		log.Fatal(msg, err)
	}
}