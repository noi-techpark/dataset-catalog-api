// © 2024 NOI Techpark <digital@noi.bz.it>
// SPDX-License-Identifier: AGPL-3.0-or-later

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"opendatahub.com/dataset-catalog-api/handlers"
)

func init() {
	// Load environment variables from .env if available.
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func main() {
  mode := os.Getenv("GIN_MODE")
	if mode == "" {
		// Se non impostato, usa la modalità "release"
		gin.SetMode(gin.ReleaseMode)
	} else {
		// Altrimenti, usa il valore specificato nell'ambiente
		gin.SetMode(mode)
	}
	router := gin.Default()

	// Load HTML templates from the "templates" directory.
	router.LoadHTMLGlob("templates/*.html")

	// Register the index route using the dedicated handler.
	router.GET("/", handlers.IndexHandler)

	// Register other endpoints.
	router.GET("/dcat", handlers.DcatGinHandler)
	router.GET("/odps", handlers.ODPSGinHandler)
	router.GET("/odps30", handlers.ODPS30GinHandler)
	router.GET("/odps30/:uuid", handlers.ODPS30DetailGinHandler)
	router.GET("/odps31", handlers.ODPS31GinHandler)
	router.GET("/odps31/:uuid", handlers.ODPS31DetailGinHandler)

  // Register the new /healthcheck endpoint.
	router.GET("/healthcheck", handlers.HealthcheckHandler)

	fmt.Println("Server running on :8878")
	log.Fatal(router.Run(":8878"))
}
