package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gogogo/internal/repository/user"
	srvsuser "gogogo/internal/service/v1/user"
	"gogogo/pkg/config"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	SqlConn *gorm.DB
)

func init() {
	err := godotenv.Load("env/dev.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbConfig := os.Getenv("SQLSTRING")
	SqlConn = config.CreateDBConnection(dbConfig)
}

func RouterV1(app *fiber.App) (err error) {
	api := app.Group("/api/v1")

	err = Handler(api)
	if err != nil {
		return err
	}

	return
}

func Handler(group fiber.Router) (err error) {
	DbGoRepo := user.UserTable(SqlConn)
	UserHandler := srvsuser.NewUserHandler(DbGoRepo)

	userRoute := group.Group("/user")
	userRoute.Get("/", UserHandler.GetAllUser)
	return
}
