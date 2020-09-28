package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Monkey struct{}

type Compile struct {
	Program string `json:"program"`
}

func (*Monkey) Compile(c *gin.Context) {
	var req Compile
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	/*input := antlr.NewInputStream(req.Program)
	lexer := parser.NewCompiler(input)*/

	c.JSON(200, gin.H{
		"message": req.Program,
	})
}
