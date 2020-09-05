package util

import (
	"log"
)

func CheckErr(err error, level string) error {
	if level == "fatal" {
		log.Fatalln(err)

	} else {
		log.Println(err)
	}
	return err
}
