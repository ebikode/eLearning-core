package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/ebikode/eLearning-core/config"
	storage "github.com/ebikode/eLearning-core/storage/mysql"
	"github.com/ebikode/eLearning-core/translation"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"

	// ut "github.com/ebikode/eLearning-core/utils"
	"github.com/rs/cors"
)

func main() {
	// This is set once to handle all random string generation in the application
	rand.Seed(time.Now().UnixNano())

	errEnv := godotenv.Load()

	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialising application configurations
	configuration, err := config.New()
	if err != nil {
		log.Panicln("Configuration error", err, configuration)
	}

	// initialize DB
	dbConfig := storage.New(configuration)
	mdb, err := dbConfig.InitDB()

	// if an error occurred while initialising db
	if err != nil {
		log.Println(err)
	}

	tErr := translation.NewTranslationBundle()
	log.Println(tErr)

	router := InitRoutes(configuration.Constants, mdb)

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route) // Walk and print out all routes
		return nil
	}
	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging err: %s\n", err.Error()) // panic if there is an error
	}

	// Set sport monks api and initiate jobs
	InitJobs(mdb)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{
			"*",
			// "http://localhost:3000",
			// "https://whitelabel.pexportal.com",
			// "https://laughing-keller-a1692a.netlify.app",
		},
		AllowCredentials: true,
		AllowedHeaders:   []string{"AppKey", "AccountKey", "Allow", "Authorization", "Content-Type"},
		AllowedMethods:   []string{"GET", "POST", "PUT"},
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})

	handler := c.Handler(router)

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))

	// Serve the application
	log.Println("Serving application at PORT :" + port)
	log.Fatal(http.ListenAndServe(port, handler)) // Note, the port is usually gotten from the environment.

}
