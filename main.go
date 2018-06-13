package main

import (
	"log"

	"github.com/richardcane/doc/cmd"
)

func main() {
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
	cmd.Execute()
}
