package main

import (
	"flag"
	"fmt"
	"os/exec"
	"math/rand"
	"strconv"
	"strings"
)

var hosts = flag.String("hosts", "", "hosts to test")

func main() {
	flag.Parse()

	hs := strings.Split(*hosts, ",")
	results := make(chan error)

	for _, h := range hs {
		go func(host string) {
			d := rand.Int()
			o, err := exec.Command("gcloud", "compute", "ssh", host, "--", "mkdir", "/tmp/rep" + strconv.Itoa(d)).CombinedOutput()
			if err != nil {
				results <- fmt.Errorf("%v %s", err, o)
				return
			}
			o, err = exec.Command("gcloud", "compute", "ssh", host, "--", "rmdir", "/tmp/rep" + strconv.Itoa(d)).CombinedOutput()
			if err != nil {
				results <- fmt.Errorf("%v %s", err, o)
				return
			}

			results <- nil
		}(h)
	}

	for i := 0; i < len(hs); i++ {
		if err := <-results; err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}
