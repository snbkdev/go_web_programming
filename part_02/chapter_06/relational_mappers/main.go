package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

type Post struct {
	Id int
	Content string
	AuthorName string `db:"author"`
}

var Db *sqlx.DB

func init() {
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
    
    Db, err = sqlx.Open("postgres", connStr)
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    
    err = Db.Ping()
    if err != nil {
        log.Fatalf("Failed to ping database: %v", err)
    }
    
    log.Println("Successfully connected to database!")
}

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

func GetPost(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRowx("select id, content, author from posts where id = $1", id).StructScan(&post)
	if err != nil {
		return
	}

	return
}

func (post *Post) Create() (err error) {
	err = Db.QueryRow("insert into posts (content, author) values($1, $2) returning id", post.Content, post.AuthorName).Scan(&post.Id)
	return
}

func main() {
	post := Post{Content: "Hello!", AuthorName: "Armor"}
	post.Create()
	fmt.Println(post)
}