package server

import (
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	jwtware "github.com/gofiber/jwt/v3"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"minibank/auth"
	authhttp "minibank/auth/delivery/http"
	authrepo "minibank/auth/repository"
	"minibank/auth/usecase"
	"minibank/transaction"
	transhttp "minibank/transaction/delivery/http"
	transrepo "minibank/transaction/repository"
	transuc "minibank/transaction/usecase"
)

type App struct {
	authUC  auth.UseCase
	transUC transaction.UseCase
}

func NewApp() (*App, error) {
	db, err := initDb()
	if err != nil {
		return nil, err
	}

	authRepo := authrepo.NewUserRepository(db)
	transRepo := transrepo.NewTransactionRepo(db)

	return &App{
		authUC: usecase.NewAuthUseCase(
			authRepo,
			viper.GetString("HASH_SALT"),
			[]byte(viper.GetString("SIGNING_KEY")),
			viper.GetDuration("TOKEN_TTL"),
		),
		transUC: transuc.NewTransUseCase(transRepo),
	}, nil
}

func ping() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(200).JSON(&fiber.Map{
			"msg": "pong",
		})
	}
}

func (a *App) Run(port string) error {
	// Init handler
	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())
	app.Get("/", ping())
	app.Get("/ping", ping())

	authhttp.RegisterHTTPEndpoints(app, a.authUC)

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		},
		SigningKey: []byte(viper.GetString("SIGNING_KEY")),
	}))

	app.Use(authhttp.CurrentUser(a.authUC))

	transhttp.RegisterHTTPEndpoints(app, a.transUC)

	addr := fmt.Sprintf(":%d", viper.GetInt("PORT"))
	err := app.Listen(addr)

	if err != nil {
		return err
	}

	return nil
}

func initDb() (*sql.DB, error) {
	db, err := sql.Open("postgres", viper.GetString("DATABASE_URL"))

	if err != nil {
		return nil, fmt.Errorf("sql.Open %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("db.Ping %w", err)
	}

	return db, nil
}
