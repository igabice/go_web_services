package main

import (
	"io/ioutil"
	"net/http"
	"log"
	"fmt"
)

func main(){


	http.HandleFunc("/", func(rw http.ResponseWriter, r*http.Request){
		log.Println("Hello World")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil{
			log.Println(err)
		http.Error(rw, "Oooops", http.StatusBadRequest)
			return
		}
		log.Printf("Data %s\n", d)
		fmt.Fprintf(rw, "Hwllo %s", d)
	})
	
	http.HandleFunc("/join", func(http.ResponseWriter, *http.Request){
		log.Println("Hello World, come join me!")
	})

	http.ListenAndServe(":9090", nil)

}