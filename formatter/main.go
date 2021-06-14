package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Indosaram/gomunkong-backend/formatter/file_io"

	"github.com/gin-gonic/gin"
)

func runFormatter(input file_io.Input) string {
	var formatted_code string
	var port string

	switch input.Lang {
	case "python":
		port = os.Getenv("LANG_SERVERS_PYTHON_SERVICE_PORT")
	case "javascript":
		port = os.Getenv("LANG_SERVERS_JAVASCRIPT_SERVICE_PORT")
	case "golang":
		port = os.Getenv("LANG_SERVERS_GOLANG_SERVICE_PORT")
	case "java":
		port = os.Getenv("LANG_SERVERS_JAVA_SERVICE_PORT")
	default:
		return ""
	}
	fmt.Println("Selected port for", input.Lang, "is", port)
	formatted_code = Format(port, input.Code, input.Formatter, input.Args)
	return formatted_code
}

func main() {
	g := gin.Default()
	g.GET("/:lang/:formatter/:code", func(ctx *gin.Context) {
		input := file_io.Input{
			Lang:      ctx.Param("lang"),
			Formatter: ctx.Param("formatter"),
			Code:      ctx.Param("code"),
		}
		formatted_code := runFormatter(input)

		fmt.Println(formatted_code)
		ctx.JSON(http.StatusOK, gin.H{"result": formatted_code})
	})

	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
