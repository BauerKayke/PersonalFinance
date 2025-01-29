package application

import (
	"backend/config"
	"backend/internal/handlers"
	"backend/internal/middleware"
	"backend/internal/repositories"
	"backend/internal/routes"
	"backend/internal/services"
	"github.com/go-chi/chi/v5"
	middleware2 "github.com/go-chi/chi/v5/middleware"
	"net/http"
	"time"
)

type App struct {
	Server *http.Server
}

func NewApp() *App {
	cfg := config.LoadConfig()

	db := config.Init(cfg)

	transactionsRepository := repositories.NewTransactionRepository(db)
	transactionsService := services.NewTransactionService(transactionsRepository)
	transactionsHandler := handlers.NewTransactionHandler(transactionsService)
	transactionRouter := routes.NewTransactionRouter(transactionsHandler)

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)
	userRouter := routes.NewUserRouter(userHandler)

	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authHandler := handlers.NewAuthHandler(authService)
	authRouter := routes.NewAuthRouter(authHandler)

	creditCardRepository := repositories.NewCreditCardRepository(db)
	creditCardService := services.NewCreditCardService(creditCardRepository)
	creditCardHandler := handlers.NewCreditCardHandler(creditCardService)
	creditCardRouter := routes.NewCreditCardRouter(creditCardHandler)

	bankAccountRepository := repositories.NewBankAccountRepository(db)
	bankAccountService := services.NewBankAccountService(bankAccountRepository)
	bankAccountHandler := handlers.NewBankAccountHandler(bankAccountService)
	bankAccountRouter := routes.NewBankAccountRouter(bankAccountHandler)

	router := chi.NewRouter()

	router.Use(middleware2.Logger)
	router.Use(middleware2.Recoverer)

	authMiddleware := middleware.NewAuthMiddleware(authService)

	bankAccountRouter.RegisterRoutes(router, authMiddleware)
	userRouter.RegisterRoutes(router, authMiddleware)
	creditCardRouter.RegisterRoutes(router, authMiddleware)
	transactionRouter.RegisterRoutes(router, authMiddleware)
	authRouter.RegisterRoutes(router)

	// Servidor com o roteador configurado
	return &App{Server: &http.Server{
		Addr:         ":8080",
		Handler:      middleware.RateLimitMiddleware(router),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second},
	}
}

func (app *App) Run() error {
	return app.Server.ListenAndServe()
}
