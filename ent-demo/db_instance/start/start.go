package main

import (
	"context"
	"github.com/minicloudsky/golang-in-action/ent-demo/db_instance/ent"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func CreateEntClient() *ent.Client {
	client, err := ent.Open("mysql",
		"root:root@tcp(localhost:3306)/ent?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}

	return client
}

var client = CreateEntClient()

func main() {
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
