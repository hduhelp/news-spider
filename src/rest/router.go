package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/setcy/spider/src/rest/controller"
)

func (s *Service) initRouter() {
	engine := gin.New()

	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(controller.MakeSuccessJSON(nil, "pong"))
	})

	var dataController = controller.DataController{RuleConfig: s.Rule, DB: s.DB}
	engine.GET("/refresh", dataController.GetDataHandler)
	engine.GET("/tags", dataController.GetTagHandler)
	engine.GET("/feed", dataController.GetFeedHandler)
	engine.GET("/channel", dataController.GetNoticeHandler)

	var subController = controller.SubscribeController{}
	engine.GET("/sub", subController.GetSubscribeHandler)
	engine.POST("/sub", subController.SubscribeHandler)
	engine.DELETE("/sub", subController.UnSubscribeHandler)

	err := engine.Run(":20000")
	if err != nil {
		panic(err)
	}
}
