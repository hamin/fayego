package fayeserver

import (
	"fmt"
  // "io"
	// "io/ioutil"
	"net/http"
	// "os"
	//"encoding/json"
	// "text/template"
	// "time"
)

func serveLongPolling(f *FayeServer, w http.ResponseWriter, r *http.Request) {
	fmt.Println("PLZ WORK ", r.FormValue("message"))
	jsonMessage := r.FormValue("message")
	jsonpParam := r.FormValue("jsonp")
	fmt.Println("JSONP VALUE ", jsonpParam)
	fmt.Println("THIS IS JSON MESSAGE ", jsonMessage)

	// parse message url param

	// instantiate channel
	send := make(chan string)
	response, error := f.HandleMessage([]byte(jsonMessage), send)

	if error != nil {
		// We have to figure what correct error response is for HTTP for faye
		fmt.Println("HTTP SERVER ERROR: ", error)
		return
	} else {
		finalResponse := jsonpParam + "(" + string(response) + ");"
		fmt.Println("THIS IS OUR HTTP RESPONSE: %v", finalResponse)
		fmt.Println("THIS IS THE W HEADERS: ", w.Header())
		//n, err := io.WriteString(w, finalResponse)
		fmt.Fprint(w, finalResponse)
		//fmt.Println("sending.")
		//send <- finalResponse
		//fmt.Println("THIS IS N: ", n)
		//fmt.Println("THIS IS ERROR: ", err)
	}

	/*
	for {
		select {
		case <-time.After(60e9):
			fmt.Println("IT CAME TO TIMEOUT!!!!")
			io.WriteString(w, "Timeout!\n")
		case msg := <-send:
			fmt.Println("IT CAME TO MSG SEND PART")
			io.WriteString(w, msg)
		}	
	}*/

}
