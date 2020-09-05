package util

import (
	"log"
)

func CheckErr(err error, level string) {
	if level == "fatal" {
		log.Fatalln(err)
	} else {
		log.Println(err)
	}
}
