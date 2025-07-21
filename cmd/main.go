package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/Ararat25/api-marketplace/config"
	_ "github.com/Ararat25/api-marketplace/docs"
	"github.com/Ararat25/api-marketplace/internal/controller"
	"github.com/Ararat25/api-marketplace/internal/database"
	middle "github.com/Ararat25/api-marketplace/internal/middleware"
	"github.com/Ararat25/api-marketplace/internal/model"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

const configPath = "config.yml" // путь до файла конфигурации

var (
	authService   *model.AuthService
	marketService *model.MarketplaceService
)

// @title Auth AuthService API
// @version 1.0
// @description This is an authentication service with JWT
// @host localhost:8080
// @BasePath /
func main() {
	handler, conf := initApp()

	r := initRouter(handler)

	hostPort := fmt.Sprintf("%s:%d", conf.Server.Host, conf.Server.Port)

	log.Printf("Server starting on %s", hostPort)
	err := http.ListenAndServe(hostPort, r)
	if err != nil {
		log.Fatalf("Start server error: %s", err.Error())
	}
}

// initApp инициализирует конфигурацию, подключение к базе данных и сервисы приложения
func initApp() (*controller.Handler, *config.Config) {
	err := config.LoadEnvVariables()
	if err != nil {
		log.Fatalf("error loading env variables: %v\n", err)
	}

	passwordSalt, tokenSalt, dbHost, dbUser, dbPassword, dbName, dbPortString :=
		os.Getenv("PASSWORD_SALT"),
		os.Getenv("TOKEN_SALT"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT")

	if passwordSalt == "" || tokenSalt == "" || dbHost == "" || dbUser == "" || dbPassword == "" || dbName == "" || dbPortString == "" {
		log.Fatalln("not all environment variables are set")
	}

	dbPort, _ := strconv.Atoi(dbPortString)

	conf, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatalf("error loading config file: %v\n", err)
	}

	err = database.ConnectDB(dbHost, dbUser, dbPassword, dbName, dbPort)
	if err != nil {
		log.Fatalf("error connecting to database: %v\n", err)
	}

	authService = model.NewAuthService([]byte(passwordSalt), []byte(tokenSalt), conf.Server.AccessTokenTTl, conf.Server.RefreshTokenTTl, database.DB.Db)
	marketService = model.NewMarketplaceService(database.DB.Db)

	handler := controller.NewHandler(authService, marketService)

	return handler, conf
}

// initRouter настраивает маршруты и middleware для сервера
func initRouter(handler *controller.Handler) *chi.Mux {
	r := chi.NewRouter()

	middlew := middle.NewMiddleware(authService)

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middlew.JsonHeader)

	r.Get("/api/docs/*", httpSwagger.WrapHandler)
	r.Post("/api/register", handler.Register)
	r.Post("/api/auth", handler.Auth)
	r.Post("/api/refresh", handler.RefreshToken)
	r.Post("/api/logout", handler.Logout)

	r.Group(func(r chi.Router) {
		r.Use(middlew.CheckAuth)
		r.Post("/api/create", handler.CreateAd)
	})

	r.Post("/api/ads", handler.ListAd)

	return r
}
