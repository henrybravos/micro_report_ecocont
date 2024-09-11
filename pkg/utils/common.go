package utils

import "log"

func CheckErr(err error) {
	if err != nil {
		log.Printf("Error: %v", err)
		panic(err)
	}
}
