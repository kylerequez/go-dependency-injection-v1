package main

import (
	"fmt"

	"github.com/kylerequez/go-dependency-injection/initializers"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	fmt.Println(":::-:::\t\tGO DEPENDENCY INJECTION EXAMPLE\t\t:::-:::")

	var dependencies *initializers.Dependencies = initializers.InitializeDependencies()

	dependencies.UserController.GetAllUsers()
}
