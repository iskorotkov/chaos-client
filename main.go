package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"time"
)

var verb = flag.String("verb", "get", "Makes client execute request with given verb")
var rate = flag.Int("rate", 0, "Rate at which to send messages to the server")
var port = flag.Int("port", 8811, "Port of server to send requests to")
var host = flag.String("host", "localhost", "Host of server to send requests to")

func main() {
	flag.Parse()

	address := getAddress()
	if *rate > 0 {
		ticker := time.NewTicker(time.Second / time.Duration(*rate))
		sendRequestsOnTimer(address, ticker.C)

	} else {
		err := sendRequest(address)
		if err != nil {
			panic(err)
		}
	}
}

func sendRequestsOnTimer(address string, ch <-chan time.Time) {
	for {
		err := sendRequest(address)
		if err != nil {
			panic(err)
		}

		<-ch
	}
}

func sendRequest(address string) error {
	t := time.Now()
	defer func() {
		fmt.Printf("Time elapsed: %v\n", time.Since(t))
	}()

	switch *verb {
	case "get":
		return getCounter(address)
	case "post":
		return incCounter(address)
	default:
		return errors.New("no correct verb was specified")
	}
}

func getCounter(address string) error {
	r, err := http.Get(address)
	if err != nil {
		fmt.Printf("Error occured when reaching to server: %v\n", err)
		return err
	}

	defer func() { _ = r.Body.Close() }()

	buf := make([]byte, 512)
	num, err := r.Body.Read(buf)
	if err != nil && err != io.EOF {
		fmt.Printf("Error occured when reading response body: %v\n", err)
		return err
	}

	s := string(buf[:num])
	fmt.Printf("r = %v\n", s)
	return nil
}

func incCounter(address string) error {
	r, err := http.Post(address, "", nil)
	if err != nil {
		fmt.Printf("Error occured when reaching to server: %v\n", err)
		return err
	}

	defer func() { _ = r.Body.Close() }()

	fmt.Println("Counter has been incremented")
	return nil
}

func getAddress() string {
	return fmt.Sprintf("http://%v:%v/counter", *host, *port)
}
