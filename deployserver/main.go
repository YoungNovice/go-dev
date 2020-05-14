package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func main() {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, `this is deploy server`)
		cmd := exec.Command("sh", "./deploy.sh")
		err := cmd.Start()
		if err != nil {
			fmt.Fprint(writer, err)
			log.Fatal(err)
		}
		err = cmd.Wait()
		if err != nil {
			fmt.Fprint(writer, err)
		}
	})
	http.ListenAndServe(":5000", nil)
}
