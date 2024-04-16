package main

import (
	"context"
	"log"

	spanner "cloud.google.com/go/spanner"
	"google.golang.org/api/iterator"
)

func main() {
	// Create a new client
	client, err := spanner.NewClient(context.Background(), "projects/<project>/instances/<instance>")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// Define the query
	query := `SELECT * FROM <table>`

	// Execute the query
	iter := client.Single().Query(context.Background(), spanner.NewStatement(query))

	// Iterate through the rows
	for {
		row, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to read row: %v", err)
		}

		// Print the values
		for i, col := range row.Columns {
			log.Printf("Column %d: %s", i, col)
		}
	}
}
In this example, we first create a new client for the Cloud Spanner database using the spanner.NewClient function. We then define the query that we want to execute and use the client.Single().Query function to execute the query.

We then iterate through the rows using the iterator.Next function and print the values of each column.

Note that yo
