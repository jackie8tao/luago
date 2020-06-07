package main

import (
	"log"
	"strconv"
)

func main() {
	val, err := strconv.ParseUint("1E-02", 10, 64)
	if err != nil {
		log.Println(err)
	}
	log.Println(val)
}
