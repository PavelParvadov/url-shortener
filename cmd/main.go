package main

import (
	"fmt"
	"net/http"
	"url/configs"
	_ "url/configs"
	"url/internal/auth"
	"url/internal/link"
	"url/internal/stat"
	"url/internal/user"
	"url/pkg/db"
	"url/pkg/event"
	"url/pkg/middleware"
)

func App() http.Handler {
	config := configs.LoadConfig()
	DB := db.NewDb(config)
	router := http.NewServeMux()
	//repositories
	LinkRepository := link.NewLinkRepository(DB)
	UserRepository := user.NewUserRepository(DB)
	StatRepository := stat.NewStatRepository(DB)

	//паттерн шины событий тут
	EventBus := event.NewEventBus()

	//services

	authService := auth.NewAuthService(UserRepository)
	statService := stat.NewStatService(&stat.StatServiceDeps{
		EventBus:       EventBus,
		StatRepository: StatRepository,
	})
	go statService.AddClick()

	//handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      config,
		AuthService: authService,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: LinkRepository,
		Config:         config,
		EventBus:       EventBus,
	})

	stat.NewStatHandler(router, stat.StatHandlerDeps{
		StatRepository: StatRepository,
		Config:         config,
	})
	//middlewares
	stack := middleware.Chain(middleware.CORS, middleware.Logging)
	return stack(router)
}

func main() {
	server := http.Server{
		Addr:    ":8080",
		Handler: App(),
	}
	fmt.Println("Сервер запущен")
	err := server.ListenAndServe()

	if err != nil {
		return
	}

}
