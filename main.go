package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thegera4/events-rest-api/db"
	"github.com/thegera4/events-rest-api/routes"
)

//NOTE 1 framework for rest api: go get -u github.com/gin-gonic/gin
//NOTE 2 use the sql package from go + a sqldriver (in this case for sqlite3) go get modernc.org/sqlite
//NOTE 3 run "set CGO_ENABLED=1" if you are on windows before "go run ." because sqlite3 needs this
//NOTE 4 run "go get github.com/golang-jwt/jwt/v5" to get the jwt package
//NOTE 5 run "go get golang.org/x/crypto/bcrypt" to get the bcrypt package
func main() {
	db.InitDB()
	
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}