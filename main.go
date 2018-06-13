package main

import (
	"doc/cmd"
	"log"
)

func main() {
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
	cmd.Execute()
}
