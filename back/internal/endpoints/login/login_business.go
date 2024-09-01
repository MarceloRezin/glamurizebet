package login

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func logar(c *gin.Context) {
	var l = login{
		Email: "a@a.com", Senha: "1234",
	}

	fmt.Println(l)

	c.IndentedJSON(http.StatusOK, l)
}
