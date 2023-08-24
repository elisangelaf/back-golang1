package main

import (
	"github.com/gin-gonic/gin"
)
func main() {
	//Instancia de rotas do gin
	r := gin.Default()

	//Rota para o endpoint /ping
	r.GET("/ping", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "pong",})
		})

		r.GET("/health", func(c *gin.Context){
			c.JSON(200, gin.H{
				"message": "aplication is runing",})
			})

		//porta que o servidor irá escutar
		r.Run(":8080")
}
