// File path: cmd/csp-report-api/main.go
package main

import (
	"context"
	"fmt"
	"github.com/achimgrolimund/csp-report-api/api/middleware"
	"github.com/achimgrolimund/csp-report-api/pkg/api"
	"github.com/achimgrolimund/csp-report-api/pkg/domain/csp_report"
	"github.com/achimgrolimund/csp-report-api/pkg/infrastructure/persistence/mongodb"
	"github.com/achimgrolimund/csp-report-api/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"os"
	"time"
)

func main() {

	log, err := logger.NewLogger()
	if err != nil {
		panic(err)
	}
	log.Info("Starting CSP Report API")

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbName := os.Getenv("DB_NAME")
	dbCollection := os.Getenv("DB_COLLECTION")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbEndpoint := os.Getenv("DB_ENDPOINT")
	dbPort := os.Getenv("DB_PORT")

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", dbUser, dbPassword, dbEndpoint, dbPort)

	// Create a MongoDB client
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Failed to create client: %v", zap.Error(err))
	}

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err = mongo.Connect(ctx)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB: %v", zap.Error(err))
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Failed to ping MongoDB: %v", zap.Error(err))
	}
	log.Info("Connected to MongoDB")

	// Create a CSP report repository
	repo := mongodb.NewCSPReportRepository(client, dbName, dbCollection)
	log.Info("CSP Report Repository created",
		zap.String("db_name", dbName),
		zap.String("db_collection", dbCollection))

	r := gin.New()

	r.Use(middleware.ZapLogger(log))

	// Use the CORS validator middleware
	r.Use(middleware.CORSValidator())
	log.Info("CORS Validator middleware registered")

	// Create a CSP report service
	service := csp_report.NewReportService(repo)
	log.Info("CSP Report Service created")
	handler := api.NewHandler(service)
	log.Info("API Handler created")

	// Handle POST request to /report
	r.POST("/report", middleware.AuthMiddleware(), handler.SaveReportHandler)
	log.Info("POST /report handler registered")

	// Handle GET request to /health
	r.GET("/health", api.HealthCheckHandler)
	log.Info("GET /health handler registered")

	log.Info("Starting server on port 8080")
	err = r.Run(":8080")
	if err != nil {
		log.Fatal("Failed to start server: %v", zap.Error(err))
		return
	} // listen and serve on 0.0.0.0:8080

}
