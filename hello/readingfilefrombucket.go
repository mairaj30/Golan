package main

import (
	"context"
	"fmt"
	"io/ioutil"

	"cloud.google.com/go/storage"
)

func main() {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		// TODO: handle error.
	}
	rc, err := client.Bucket("my-bucket").Object("my-object").NewReader(ctx)
	if err != nil {
		// TODO: handle error.
	}
	slurp, err := ioutil.ReadAll(rc)
	rc.Close()
	if err != nil {
		// TODO: handle error.
	}
	fmt.Println("file contents:", slurp)
}
