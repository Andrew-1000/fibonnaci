package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"azure-devops/algorithms"
)
func main() {
	
	var port = flag.String("port", ":8081", "Port to use")

	flag.Parse()

	declare := gin.Default()
	declare.GET("/fibon", func(c *gin.Context) {
		fibResult := []int64{}

		fib := algorithms.NewFibonacciIterator()

		for fib.Next() {
			fibResult = append(fibResult, fib.Value())
			if fib.Value() == 34 {
				break
			}
		}
		c.JSON(200, fibResult)

	})
	declare.Run(*port)
	
}