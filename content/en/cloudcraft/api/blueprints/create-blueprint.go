package main

import (
	"context"
	"log"
	"os"

	"github.com/DataDog/cloudcraft-go"
)

func main() {
	key, ok := os.LookupEnv("CLOUDCRAFT_API_KEY")
	if !ok {
		log.Fatal("missing env var: CLOUDCRAFT_API_KEY")
	}

	// Create new Config to be initialize a Client.
	cfg := cloudcraft.NewConfig(key)

	// Create a new Client instance with the given Config.
	client, err := cloudcraft.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// List all blueprints in an account.
	blueprint, _, err := client.Blueprint.Create(
		context.Background(),
		&cloudcraft.Blueprint{
			Name: "My new blueprint",
		}
	)
	if err != nil {
		log.Fatal(err)
	}

	// Print the name of blueprint we created
	log.Println(blueprint.Name)
}
