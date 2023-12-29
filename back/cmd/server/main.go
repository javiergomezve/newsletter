package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"newsletter-back/config"
	"newsletter-back/controllers"
	"newsletter-back/data"
	"newsletter-back/repositories"
	"newsletter-back/services"
)

type BaseController interface {
	RouteSetup(rg *gin.RouterGroup)
}

func setRoutes(router *gin.Engine, routes map[string]BaseController) {
	for route, controller := range routes {
		group := router.Group(route)
		controller.RouteSetup(group)
	}
}

func Run() error {
	env, err := config.NewEnv()
	if err != nil {
		return err
	}

	pgDatabase, err := data.NewPostgresDatabase(data.PostgresDatabaseOptions{
		Username: env.DbUsername,
		Password: env.DbPassword,
		Db:       env.DbDb,
		Host:     env.DbHost,
		Port:     env.DbPort,
		SslMode:  env.DbSslMode,
	})
	if err != nil {
		return err
	}

	crypt := config.NewBcryptAdapter()
	jwt := config.NewJwtAdapter(env.AccessTokenPrivateKey, env.AccessTokenPublicKey, env.AccessTokenExpiredIn)
	s3 := config.NewAWSS3Uploader(env.AwsS3Region, env.AwsS3Bucket, env.AwsAccessKeyID, env.AwsSecretAccessKey)
	ses, _ := config.NewSESService(env.AwsS3Region, "newsletter.javiergomezve@gmail.com", env.FrontendUrl)
	//credentialsFile := "path/to/credentials.json"
	//gmailService, err := NewGmailService(credentialsFile)
	//if err != nil {
	//	log.Fatalf("Failed to create GmailService: %v", err)
	//}

	pgAuthRepository := repositories.NewPostgresAuthRepository(pgDatabase)
	authService := services.NewAuthService(pgAuthRepository, jwt.GenerateToken, crypt.ComparePasswords)

	newsletterRepository := repositories.NewPgNewsletterRepository(pgDatabase)
	newsletterService := services.NewNewsletterService(newsletterRepository)

	listRepository := repositories.NewPgListRepository(pgDatabase)
	listService := services.NewListService(listRepository)

	recipientRepository := repositories.NewPgRecipientRepository(pgDatabase)
	recipientService := services.NewRecipientService(recipientRepository)

	mediaRepository := repositories.NewPgMediaRepository(pgDatabase)
	mediaService := services.NewMediaService(mediaRepository, s3)

	gin.SetMode(gin.ReleaseMode)
	svr := gin.New()
	svr.MaxMultipartMemory = 8 << 20 // 8 MiB

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowCredentials = true
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"}
	corsConfig.MaxAge = 300
	corsConfig.ExposeHeaders = []string{"*"}

	svr.Use(cors.New(corsConfig))

	routes := make(map[string]BaseController)
	routes["/api/auth"] = controllers.NewAuthController(authService)
	// TODO: protect with authentication
	routes["/api/lists"] = controllers.NewListController(listService)
	routes["/api/recipients"] = controllers.NewRecipientController(recipientService)
	routes["/api/media"] = controllers.NewMediaController(mediaService)
	routes["/api/newsletters"] = controllers.NewNewsletterController(newsletterService, ses)

	setRoutes(svr, routes)

	return svr.Run(":" + env.HttpPort)
}

func main() {
	log.Println("Server running")

	if err := Run(); err != nil {
		log.Fatal("Error starting the REST API: ", err)
	}

}
