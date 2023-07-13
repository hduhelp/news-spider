package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/setcy/spider/src/rule"
	"github.com/setcy/spider/src/subscribe"
	"gorm.io/gorm"
)

type SubscribeController struct {
	RuleConfig *rule.RuleConfig
	DB         *gorm.DB
}

func (s SubscribeController) SubscribeHandler(c *gin.Context) {
	tag := c.Query("tag")
	staffID := c.GetHeader("staffId")
	ruleItem := s.RuleConfig.SiteRule.TagFilter(tag)
	if ruleItem == nil {
		c.JSON(MakeErrorJSON(404, 40400, "error tag"))
		return
	}
	err := subscribe.SaveSubscribeInDatabase(s.DB, staffID, ruleItem)
	if err != nil {
		c.JSON(MakeErrorJSON(403, 40301, err))
		return
	}
	c.JSON(MakeSuccessJSON(nil, "success"))
	return
}

func (s SubscribeController) UnSubscribeHandler(c *gin.Context) {
	tag := c.Query("tag")
	staffID := c.GetHeader("staffId")
	ruleItem := s.RuleConfig.SiteRule.TagFilter(tag)
	if ruleItem == nil {
		c.JSON(MakeErrorJSON(404, 40400, "error tag"))
		return
	}
	err := subscribe.DeleteSubscribeInDatabase(s.DB, staffID, ruleItem)
	if err != nil {
		c.JSON(MakeErrorJSON(403, 40301, err))
		return
	}
	c.JSON(MakeSuccessJSON(nil, "success"))
	return
}

func (s SubscribeController) GetSubscribeHandler(c *gin.Context) {
	staffID := c.GetHeader("staffId")
	data := subscribe.GetSubscribesByStaffID(s.DB, staffID)
	c.JSON(MakeSuccessJSON(data, "success"))
	return
}
