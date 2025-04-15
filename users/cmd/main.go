package main

import (
	pb "github.com/mhirii/chronobreak/users/internal/api/users"
	"github.com/mhirii/chronobreak/users/internal/handlers"
	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()

	authHandler := handlers.NewAuthHandler()

	// register route greet
	app.GET("/greet", func(ctx *gofr.Context) (any, error) {
		return "Hello World!", nil
	})

	pb.RegisterAuthServiceServerWithGofr(app, pb.NewAuthServiceGoFrServer())

	app.POST("/auth/login", authHandler.Login)
	app.POST("/auth/signup", authHandler.Signup)
	app.POST("/auth/refresh", authHandler.RefreshToken)
	app.GET("/auth/user", authHandler.GetUser)
	app.GET("/auth/user/:id", authHandler.GetUserByID)
	app.POST("/auth/logout", authHandler.Logout)

	app.Run()
}
