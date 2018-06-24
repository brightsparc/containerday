package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis"
)

func main() {
	var addr = flag.String("db", "db:6379", "Redis host")
	var host = flag.String("host", ":5000", "Redis host")
	var ping = flag.Bool("ping", false, "Ping command")
	flag.Parse()

	client := redis.NewClient(&redis.Options{
		Addr:     *addr,
		Password: "",
		DB:       0,
	})

	if *ping {
		fmt.Printf("Connecting to %s\n", *addr)
		pong, err := client.Ping().Result()
		if err != nil {
			panic(err)
		}
		fmt.Println(pong)
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		result, err := client.Incr("counter").Result()
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, "Hello, %d\n", result)
	})

	fmt.Printf("Listening on %s\n", *host)
	log.Fatal(http.ListenAndServe(*host, nil))
}
