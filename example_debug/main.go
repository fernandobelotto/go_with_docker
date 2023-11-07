package main

import (
	// "context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	// "github.com/go-redis/redis/v8"
	// "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// func init() {

// 	redisClient := redis.NewClient(&redis.Options{
// 		Addr:     "redis:6379",
// 		Password: "",
// 		DB:       0,
// 	})

// 	_, err := redisClient.Ping(context.Background()).Result()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Connected to Redis")

// 	cfg := mysql.Config{
// 		User:   "root",
// 		Passwd: "secret",
// 		Net:    "tcp",
// 		Addr:   "mysql:3306",
// 		DBName: "dbname",
// 	}

// 	db, err = sql.Open("mysql", cfg.FormatDSN())

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	pingErr := db.Ping()
// 	if pingErr != nil {
// 		log.Fatal(pingErr)
// 	}
// 	fmt.Println("Connected with MySQL!")
// }

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Println("Endpoint Hit: helloHandler")
	fmt.Fprintf(w, "Hello World! 2")
}

func main() {
	http.HandleFunc("/", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
