package routes

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/real-time-vocab-quiz/server/configs"
	"github.com/real-time-vocab-quiz/server/controllers"
	"github.com/real-time-vocab-quiz/server/dataservices"
	"github.com/real-time-vocab-quiz/server/services"
	"github.com/redis/go-redis/v9"
)

func Setup(r *gin.Engine, holder *dataservices.Holder, rdb *redis.Client) {
	setupCORs(r)

	// Init services and controllers
	services := services.NewServices(holder)
	controllers := controllers.NewControllers(services)

	apiV1Group := r.Group("/api/v1")
	newQuizRouter(apiV1Group, controllers.QuizController)

}

func setupCORs(r *gin.Engine) {
	// Setup Cors
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = configs.GetCorsAllowedOrigins()
	corsConfig.AllowHeaders = configs.GetCorsAllowedHeaders()
	corsConfig.AllowMethods = configs.GetCorsAllowedMethods()
	corsConfig.ExposeHeaders = configs.GetCorsExposedHeaders()
	corsConfig.MaxAge = time.Second * time.Duration(configs.GetCorsMaxAge())

	r.Use(cors.New(corsConfig))
}
