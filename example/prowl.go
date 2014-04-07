package main

import (
	"flag"
	"fmt"
	"github.com/rem7/goprowl"
	"os"
	"strings"
)

var apikey, application, event, priority, url string

func init() {
	flag.StringVar(&apikey, "apikey", "", "Your API key")
	flag.StringVar(&application, "app", "goprowl", "Your application name")
	flag.StringVar(&event, "event", "", "Prowl event")
	flag.StringVar(&priority, "pri", "0", "Prowl priority (-2 to 2)")
	flag.StringVar(&apikey, "url", "", "URL to send")
}

func main() {
	flag.Parse()

	p := goprowl.Goprowl{}
	if err := p.RegisterKey(apikey); err != nil {
		fmt.Fprintf(os.Stderr, "Error registering key:  %v\n", err)
		os.Exit(1)
	}

	n := goprowl.Notification{
		Application: application,
		Description: strings.Join(flag.Args(), " "),
		Event:       event,
		Priority:    priority,
		Url:         url,
	}

	if err := p.Push(&n); err != nil {
		fmt.Fprintf(os.Stderr, "Error sending message:  %v\n", err)
		os.Exit(1)
	}
}
