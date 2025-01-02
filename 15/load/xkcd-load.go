package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

const URL = "https://xkcd.com/" // 571/info.0.json

func scrape(i int) []byte {

	resp, err := http.Get(URL + strconv.Itoa(i) + "/info.0.json")

	if err != nil {
		fmt.Fprintf(os.Stderr, "stopped reading: %s\n", err)
		os.Exit(-1)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Fprintf(os.Stderr, "skipping %d: got %d\n", i, resp.StatusCode)
		return nil
	}

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "bad body: %s\n", err)
		os.Exit(-1)
	}

	return data
}

func main() {

	var (
		output io.WriteCloser = os.Stdout
		err    error
		cnt    int
		fails  int
		data   []byte
	)

	if len(os.Args) > 0 {
		output, err = os.Create(os.Args[1])

		if err != nil {
			log.Fatal("Something wet wrong : ", err.Error())
		}

		defer output.Close()
	}

	fmt.Fprint(output, "[")
	defer fmt.Fprint(output, "]")

	for i := 1; fails < 2; i++ {

		if data = scrape(i); data == nil {
			fails++
			continue
		}

		if cnt > 0 {
			fmt.Fprint(output, ",")
		}

		if _, err := io.Copy(output, bytes.NewBuffer(data)); err != nil {
			fmt.Fprintf(os.Stderr, "stopped: %s", err)
			os.Exit(-1)
		}

		fails = 0
		cnt++
	}
	fmt.Fprintf(os.Stderr, "read %d comics\n", cnt)
}
