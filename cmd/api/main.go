package main

import (
	"sendMail/cmd/api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.Routes(r)

	r.Run(":3000")
}
