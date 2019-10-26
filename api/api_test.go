package api

import (
	"context"
	"fmt"
	"github.com/mchirico/tlib/util"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestMainListen(t *testing.T) {
	defer util.NewTlib().ConstructDir()()

	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(4*time.Second))
	defer cancel()

	go func() {
		api := NewAPIfile("./dummy.csv")
		api.MainListen(ctx)
		for {
			select {

			case <-ctx.Done():
				return
			}
		}
	}()

	res, err := http.Get("http://localhost:3000/")
	if err != nil {
		log.Fatal(err)
	}
	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	if !strings.Contains(string(result), "welcome") {
		t.Fatalf("Did we get a server?")
	}

	fmt.Printf("%s\n", result)

	time.Sleep(5 * time.Second)
}
