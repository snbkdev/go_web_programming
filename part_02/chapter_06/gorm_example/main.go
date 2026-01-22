package main

import (
    "fmt"
    "log"
    "os"
    "strconv"
	"time"
    
    "github.com/joho/godotenv"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var Db *gorm.DB

func init() {
    if err := godotenv.Load(); err != nil {
        log.Println("Warning: .env file not found")
    }
    
    dsn := fmt.Sprintf(
        "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
        getEnv("DB_HOST", "localhost"),
        getEnvAsInt("DB_PORT", 5432),
        getEnv("DB_USER", "postgres"),
        getEnv("DB_PASSWORD", ""),
        getEnv("DB_NAME", "testdb"),
        getEnv("DB_SSLMODE", "disable"),
    )
    
    var err error
    Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    
    sqlDB, err := Db.DB()
    if err != nil {
        log.Fatalf("Failed to get database instance: %v", err)
    }
    
    if err := sqlDB.Ping(); err != nil {
        log.Fatalf("Failed to ping database: %v", err)
    }
    
    sqlDB.SetMaxIdleConns(10)
    sqlDB.SetMaxOpenConns(100)
	
	Db.AutoMigrate(&Post{}, &Comment{})

    log.Println("✅ Successfully connected to database with GORM!")
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
    if value := os.Getenv(key); value != "" {
        if intVal, err := strconv.Atoi(value); err == nil {
            return intVal
        }
    }
    return defaultValue
}

type Post struct {
	Id int
	Content string
	AuthorName string `sql:"not null"`
	Comments []Comment
	CreatedAt time.Time
}

type Comment struct {
	Id int
	Content string
	Author string `sql:"not null"`
	PostId int `sql:"index"`
	CreatedAt time.Time
}

func main() {
	post := Post{Content: "World!!!", AuthorName: "First User"}
	fmt.Println(post)

	Db.Create(&post)
	fmt.Println(post)

	comment := Comment{Content: "Excellent post!!!", Author: "Second"}
	Db.Model(&post).Association("Comments").Append(&comment)

	var readPost Post
	Db.Where("author_name = ?", "First User").First(&readPost)

	var comments []Comment
	Db.Model(&readPost).Association("Comments").Find(&comments)

	if len(comments) > 0 {
    	fmt.Println(comments[0])
	} else {
    	fmt.Println("No comments found")
	}
}