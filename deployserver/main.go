package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func main() {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(writer, `this is deploy server`)
		cmd := exec.Command("sh", "./deploy.sh")
		err := cmd.Start()
		if err != nil {
			log.Fatal(err)
		}
		err = cmd.Wait()
	})
	http.ListenAndServe("5000", nil)
}
