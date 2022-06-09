package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/meajudaaqui/user-service/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("user-api")
	{
		//ROTAS RELACIONADAS A LISTAGEM DE USUÁRIOS
		users := main.Group("usuarios")
		{
			users.GET("/", controllers.MostrarUsers)
		}

		//ROTAS RELACIONADAS CADASTRO DE USUÁRIOS
		signup := main.Group("novo")
		{
			signup.POST("/", controllers.CriarUser)
		}

		//ROTAS RELACIONADAS A AUTENTICAÇÃO E AUTORIZAÇÃO
		auth := main.Group("login")
		{
			auth.POST("/", controllers.Autenticar)
		}
	}

	return router
}
