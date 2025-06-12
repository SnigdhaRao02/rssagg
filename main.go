package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/SnigdhaRao02/rssagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	//fmt.Println("Let's get started - again")

	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found!")
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("DB_URL is not found!")
	}

	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Cannot connect to Database!")
	}

	db := database.New(conn)
	apiCfg := apiConfig{
		DB: db,
	}

	go startScraping(db, 10, time.Minute)

	router := chi.NewRouter()

	//implementing cors
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1router := chi.NewRouter()
	v1router.Get("/health", handlerReadiness)
	v1router.Get("/error", handlerErr)

	v1router.Post("/users", apiCfg.handlerCreateUser)
	v1router.Get("/users", apiCfg.handlerGetUser)

	v1router.Post("/feeds", apiCfg.handlerCreateFeed)
	v1router.Get("/feeds", apiCfg.handlerGetAllFeeds)

	v1router.Post("/feed_follows", apiCfg.handlerCreateFeedFollow)
	v1router.Get("/feed_follows", apiCfg.handlerGetAllFeedFollows)
	v1router.Delete("/feed_follows/{feedFollowID}", apiCfg.handlerDeleteFeedFollows)

	v1router.Get("/posts", apiCfg.handlerGetPostsForUser)

	router.Mount("/v1", v1router) //so full path: /v1/health

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server starting on port %v", portString)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal((err))
	}

	fmt.Println("Port:", portString)
}
