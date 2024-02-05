package initializers

import (
	"fmt"
	"log"
	"os"

	"github.com/kylerequez/go-dependency-injection/controllers"
	"github.com/kylerequez/go-dependency-injection/repository"
	"github.com/kylerequez/go-dependency-injection/services"

	"github.com/joho/godotenv"
)

type Dependencies struct {
	UserController *controllers.UserController
}

func LoadEnvVariables() {
	fmt.Println(":::-:::\t\tLoading Environment Variables\t\t:::-:::")
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}
	fmt.Println(":::-:::\t\tLoaded Environment Variables\t\t:::-:::")
}

func InitializeDependencies() *Dependencies {
	var dsn string = GenerateDsn()

	var ur *repository.UserRepository = repository.NewUserRepository(dsn)
	var us *services.UserService = services.NewUserService(ur)
	var uc *controllers.UserController = controllers.NewUserController(us)

	return &Dependencies{
		UserController: uc,
	}
}

func GenerateDsn() (dsn string) {
	var host string = os.Getenv("DB_HOST")
	var user string = os.Getenv("DB_USER")
	var password string = os.Getenv("DB_PASSWORD")
	var name string = os.Getenv("DB_NAME")
	var port string = os.Getenv("DB_PORT")
	var mode string = os.Getenv("DB_SSLMODE")
	var timezone string = os.Getenv("DB_TIMEZONE")

	var dsnTemplate string = os.Getenv("DB_URL")
	return fmt.Sprintf(dsnTemplate, host, user, password, name, port, mode, timezone)
}
