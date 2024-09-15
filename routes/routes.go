package routes

import (
	"go-backend/controllers"
	
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)

	// File routes
	app.Post("/api/upload", controllers.UploadFile)
	app.Get("/api/files", controllers.ListFiles)
	app.Get("/api/share/:file_id", controllers.ShareFile)

}
