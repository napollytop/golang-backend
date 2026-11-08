package routes


import(
"github.com/napollytop/manga/internal/handler"
"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(app *fiber.App, authHandler *handler.AuthHandler){
	api := app.Group("api/v1")

	api.Post("/register", authHandler.Register)
	api.Post("/login", authHandler.Login)
}

func GetProducts(app *fiber.App, ProductHandler *handler.ProductHandler){
	apiProduct := app.Group("api/v1/product")
	apiProduct.Get("get-products", ProductHandler.GetProducts)
	apiProduct.Post("/post-products", ProductHandler.PostProduct)
	apiProduct.Delete("/delete-products/:id", ProductHandler.DeleteProduct)
}
