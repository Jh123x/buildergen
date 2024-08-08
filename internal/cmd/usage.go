package cmd

import (
	"flag"
	"log"
)

func GetUsage() {
	flag.VisitAll(func(f *flag.Flag) {
		log.Printf("%s: %s\n", f.Name, f.Usage)
	})
}
