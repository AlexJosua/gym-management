package routes

import (
	"gym-membership/handlers"
	"gym-membership/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// ----------------- API GROUP -----------------
	api := r.Group("/api")
	{
		// ===== Auth =====
		api.POST("/users/register", handlers.Register)
		api.POST("/users/login", handlers.Login)

		// ===== Memberships =====
		memberships := api.Group("/memberships")
		memberships.Use(middleware.AuthMiddleware())
		{
			memberships.POST("", handlers.CreateMembership)
			memberships.GET("", handlers.GetMemberships)
			memberships.PUT("/:id", handlers.UpdateMembership)
			memberships.DELETE("/:id", handlers.DeleteMembership)
		}

		// ===== Trainers =====
		trainers := api.Group("/trainers")
		trainers.Use(middleware.AuthMiddleware())
		{
			trainers.POST("", handlers.CreateTrainer)
			trainers.GET("", handlers.GetTrainers)
			trainers.PUT("/:id", handlers.UpdateTrainer)
			trainers.DELETE("/:id", handlers.DeleteTrainer)
		}

		// ===== Workout Sessions =====
		workouts := api.Group("/workouts")
		workouts.Use(middleware.AuthMiddleware())
		{
			workouts.POST("", handlers.CreateWorkoutSession)
			workouts.GET("", handlers.GetWorkoutSessions)
			workouts.PUT("/:id", handlers.UpdateWorkoutSession)
			workouts.DELETE("/:id", handlers.DeleteWorkoutSession)
		}
	}

	return r
}
