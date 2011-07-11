package goprowl

import (
	"http"
	"fmt"
)

const (
	API_SERVER = "https://api.prowlapp.com"
	ADD_PATH   = "/publicapi/add"
	API_URL    = API_SERVER + ADD_PATH
)


type Goprowl struct {
	apikeys []string
}

func (gp *Goprowl) RegisterKey(key string) {

	if len(key) != 40 {

		fmt.Printf("Error, Apikey must be 40 characters long.\n")
		// raise
	}

	gp.apikeys = append(gp.apikeys, key)

}

func (gp *Goprowl) DelKey(key string) {
}

func (gp *Goprowl) Push(app string, event string, description string, priority string) {

	ch := make(chan string)

	for _, apikey := range gp.apikeys {

		apikeyList := []string{apikey}
		applicationList := []string{app}
		eventList := []string{event}
		descriptionList := []string{description}
		priorityList := []string{priority}
		vals := http.Values{"apikey": apikeyList,
			"application": applicationList,
			"description": descriptionList,
			"event":       eventList,
			"priority":    priorityList}

		// overkill?
		go func(key string) {
			r, err := http.PostForm(API_URL, vals)

			if err != nil {
				fmt.Printf("%s\n", err)
				ch <- key
			} else {
				if r.StatusCode != 200 {
					ch <- key
				} else {
					ch <- ""
				}
			}

		}(apikey)

	}


	//fmt.Printf("Waiting...\n")
	for i := 0; ; i++ {

		if i == len(gp.apikeys) {
			break
		}

		rc := <-ch
		if rc != "" {
			fmt.Printf("The following key failed: %s\n", rc)
		}

	}

}
