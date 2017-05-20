package cmd

import "log"

func fatalOnError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
