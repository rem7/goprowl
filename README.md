Goprowl
===

A wrapper for Prowl, Growl-like iPhone push notifications, written in Go.

Originally written by Yanko D Sanchez Bolanos, 07/12/2011.

Usage
---

	var p goprowl.Goprowl

	err := p.RegisterKey("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	if err != nil {
		fmt.Println("Unable to register key! - " + + err.Error())
		return
	}

	n := &goprowl.Notification{
	    Application : "Foo",
	    Description: "Foobar!",
	    Event : "Bar",
		Priority : "1",
		Providerkey : "",
		Url: "www.foobar.com",		
	}

	err = p.Push(n)
	if err != nil {
		fmt.Println("Unable to send Prowl notification! - " + + err.Error())
	}
	
	err = p.DelKey("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	if err != nil {
		fmt.Println("Unable to remove key! - " + err.Error())
	}