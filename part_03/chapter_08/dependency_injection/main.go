package main

import (
	"database/sql"
	"encoding/json"
	_ "github.com/lib/pq"
	"net/http"
	"path"
	"log"
	"strconv"
	"os"
	"fmt"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
    if valueStr, exists := os.LookupEnv(key); exists {
        if value, err := strconv.Atoi(valueStr); err == nil {
            return value
        }
    }
    return defaultValue
}

func main() {
	err := godotenv.Load()
    if err != nil {
        log.Println("Warning: .env file not found, using system environment variables")
    }
    
    dbHost := getEnv("DB_HOST", "localhost")
    dbPort := getEnvAsInt("DB_PORT", 5432)
    dbUser := getEnv("DB_USER", "postgres")
    dbPassword := getEnv("DB_PASSWORD", "")
    dbName := getEnv("DB_NAME", "testdb")
    dbSSLMode := getEnv("DB_SSLMODE", "disable")
    
    connStr := fmt.Sprintf(
        "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
        dbHost, dbPort, dbUser, dbPassword, dbName, dbSSLMode,
    )
    
    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
	err = DB.Ping()
    if err != nil {
        log.Fatalf("Failed to ping database: %v", err)
    }
	if err != nil {
		panic(err)
	}

	server := http.Server{
		Addr: "127.0.0.1:8081",
	}
	http.HandleFunc("/post/", handleRequest(&Post{Db: DB}))
	server.ListenAndServe()
}

func handleRequest(t Text) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		switch r.Method {
		case "GET":
			err = handleGet(w, r, t)
		case "POST":
			err = handlePost(w, r, t)
		case "PUT":
			err = handlePut(w, r, t)
		case "DELETE":
			err = handleDelete(w, r, t)
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func handleGet(w http.ResponseWriter, r *http.Request, post Text) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	err = post.fetch(id)
	if err != nil {
		return
	}
	output, err := json.MarshalIndent(post, "", "\t\t")
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

func handlePost(w http.ResponseWriter, r *http.Request, post Text) (err error) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	json.Unmarshal(body, post)
	err = post.create()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func handlePut(w http.ResponseWriter, r *http.Request, post Text) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	err = post.fetch(id)
	if err != nil {
		return
	}
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	json.Unmarshal(body, post)
	err = post.update()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func handleDelete(w http.ResponseWriter, r *http.Request, post Text) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	err = post.fetch(id)
	if err != nil {
		return
	}
	err = post.delete()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}