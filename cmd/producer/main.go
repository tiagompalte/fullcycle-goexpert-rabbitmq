package main

import (
	"fmt"
	"net/http"

	"github.com/tiagompalte/fullcycle-goexpert-rabbitmq/rabbitmq"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		message := r.URL.Query().Get("message")
		if message == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = rabbitmq.Publish(ctx, ch, message)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	})

	fmt.Println("Listening in port 8080")
	http.ListenAndServe(":8080", nil)
}
