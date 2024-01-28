package main

import (
	"time"
	"github.com/gin-gonic/gin"
	"github.com/thegera4/events-rest-api/db"
	"github.com/thegera4/events-rest-api/routes"
	"github.com/gin-contrib/cors"
)

//NOTE 1 framework for rest api: go get -u github.com/gin-gonic/gin
//NOTE 2 use the sql package from go + a sqldriver (in this case for sqlite3) go get modernc.org/sqlite
//NOTE 3 run "set CGO_ENABLED=1" if you are on windows before "go run ." because sqlite3 needs this
//NOTE 4 run "go get github.com/golang-jwt/jwt/v5" to get the jwt package
//NOTE 5 run "go get golang.org/x/crypto/bcrypt" to get the bcrypt package
//NOTE 6 run "go get github.com/gin-contrib/cors" to get the cors package

func main() {
	db.InitDB()
	
	server := gin.Default()

	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "DELETE", "GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
		  return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	routes.RegisterRoutes(server)

	server.Run(":8080")
}