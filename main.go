package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "template/docs"
)

// @title           Space Targets API
// @version         1.0
// @description     Space targets and their distance from Earth
// @BasePath        /
func main() {
	pool := newPool()
	defer pool.Close()

	if err := runMigrations(connectionString()); err != nil {
		log.Fatalf("migration failed: %v", err)
	}

	r := gin.Default()
	r.GET("/api/targets", listTargets(pool))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/", func(c *gin.Context) { c.Redirect(302, "/swagger/index.html") })

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(r.Run(":" + port))
}
