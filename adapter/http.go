package adapter

import (
	"desafio-nu/core/domain"
	"desafio-nu/core/ports"
	_ "desafio-nu/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func Router(useCase ports.OperationUseCase) {
	g := gin.Default()
	v1 := g.Group("/api/v1")

	healthz(v1)
	calc(useCase, v1)
	swagger(g)

	err := g.Run(":8080")
	if err != nil {
		panic(err)
	}
}

func swagger(g *gin.Engine) gin.IRoutes {
	return g.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

// @Summary      healthz
// @Description  Responds running
// @Tags         healthZ
// @Produce      json
// @Success      200
// @Router       /healthz [get]
func healthz(g *gin.RouterGroup) gin.IRoutes {
	return g.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, `running`)
	})
}

// @Summary      Calculate Capital Gain
// @Description  Calculate Capital Gain
// @Tags         Calc
// @Produce      json
// @Param        operations  body      []domain.Oper  true  "Operations JSON"
// @Success      200   {array}  domain.FeeResponse
// @Router       /calc [post]
func calc(useCase ports.OperationUseCase, g *gin.RouterGroup) gin.IRoutes {
	return g.POST("/calc", func(c *gin.Context) {
		var input []*domain.Oper

		err := c.ShouldBindJSON(&input)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		feeResults, err := useCase.CalcCapitalGain(input)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, feeResults)
	})
}
