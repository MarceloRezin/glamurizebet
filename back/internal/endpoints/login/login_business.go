package login

import (
	"fmt"
	appcontext "glamurizebet/internal/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type login_business struct {
	appContext *appcontext.AppContext
}

func (lb *login_business) logar(c *gin.Context) {

	var l = login{
		Email: "a@a.com", Senha: "1234",
	}

	var r string
	lb.appContext.DB.Raw("select now()").Scan(&r)
	fmt.Println(r)

	c.IndentedJSON(http.StatusOK, l)
}
