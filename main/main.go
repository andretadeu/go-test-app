package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func connectToDb() (*mongo.Client, error) {
	var uri string
	if uri = os.Getenv("MONGODB_URI"); uri == "" {
		return nil, fmt.Errorf("You must set your 'MONGODB_URI' environment variable. See\n\t https://docs.mongodb.com/drivers/go/current/usage-examples/")
	}
	serverApi := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverApi)
	return mongo.Connect(opts)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	http.HandleFunc("/", handler)

	_, err := connectToDb()
	if err != nil {
		log.Panic().Err(err).Msg("mongodb connection error.")
	}

	log.Info().Msg("listening to 8080...")
	log.Fatal().Err(http.ListenAndServe(":8080", nil)).Msg("error while listening on port 8080.")
}
