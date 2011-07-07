

package goprowl

import(
    "http"
    "fmt"
)

const (
	API_SERVER  = "https://api.prowlapp.com"
	ADD_PATH    = "/publicapi/add"
    API_URL       = API_SERVER + ADD_PATH
)

func Push(apikey string, application string, event string, description string, priority string){
    
   apikeyList := []string{apikey}
   applicationList := []string{application}
   eventList := []string{event}
   descriptionList := []string{description}
   priorityList := []string{priority}
   vals := http.Values{"apikey" : apikeyList, 
                              "application" : applicationList,
                               "description" : descriptionList,
                               "event" : eventList,
                               "priority" : priorityList}
     
    r, err := http.PostForm(API_URL, vals)
    
    if err != nil {fmt.Printf("%s\n", err)}
    
    // Parse the response to see if it worked.
    //fmt.Printf("%v\n", r)
    
}
