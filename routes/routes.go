package routes

import (
	"github.com/gin-gonic/gin"

	"hello-go/controllers"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")

	// Portfolio routes
	portfolios := api.Group("/portfolios")
	{
		portfolios.GET("/", controllers.GetAllPortfolios)
		portfolios.GET("/:id", controllers.GetPortfolio)
		portfolios.POST("/", controllers.CreatePortfolio)
		portfolios.PUT("/:id", controllers.UpdatePortfolio)
		portfolios.DELETE("/:id", controllers.DeletePortfolio)
	}

	// Blog posts routes
	posts := api.Group("/posts")
	{
		posts.GET("/", controllers.GetAllBlogPosts)
		posts.GET("/:id", controllers.GetBlogPost)
		posts.POST("/", controllers.CreateBlogPost)
		posts.PUT("/:id", controllers.UpdateBlogPost)
		posts.DELETE("/:id", controllers.DeleteBlogPost)
	}

	// Resume routes
	resume := api.Group("/resume")
	{
		resume.GET("/", controllers.GetAllResumes)
		resume.GET("/:id", controllers.GetResume)
		resume.POST("/", controllers.CreateResume)
		resume.PUT("/:id", controllers.UpdateResume)
		resume.DELETE("/:id", controllers.DeleteResume)
	}

	// Social media routes
	social := api.Group("/social")
	{
		social.GET("/", controllers.GetAllSocialMedia)
		social.GET("/:id", controllers.GetSocialMedia)
		social.POST("/", controllers.CreateSocialMedia)
		social.PUT("/:id", controllers.UpdateSocialMedia)
		social.DELETE("/:id", controllers.DeleteSocialMedia)
	}
}
