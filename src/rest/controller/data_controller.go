package controller

import (
	"encoding/xml"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/setcy/spider/src/post"
	"github.com/setcy/spider/src/rule"
	"golang.org/x/tools/blog/atom"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

type DataController struct {
	RuleConfig *rule.RuleConfig
	DB         *gorm.DB
}

func (d DataController) GetDataHandler(c *gin.Context) {
	if c.Query("tag") != "" {
		ruleItem := d.RuleConfig.SiteRule.TagFilter(c.Query("tag"))
		if ruleItem == nil {
			c.JSON(MakeErrorJSON(404, 40400, "error tag"))
			return
		}
		RangeDataWithoutNotify(d.DB, ruleItem)
		c.JSON(MakeSuccessJSON(nil, "success"))
	}
	c.JSON(MakeErrorJSON(404, 40400, "error tag"))
	return
}

func (d DataController) GetTagHandler(c *gin.Context) {
	var ruleArray = d.RuleConfig.SiteRule.ToRuleArrayItem(new(rule.RuleItem))
	c.JSON(MakeSuccessJSON(ruleArray.Item, "success"))
	return
}

func (d DataController) GetNoticeHandler(c *gin.Context) {
	tag := c.Query("tag")
	fmt.Print(tag)
	ruleItem := d.RuleConfig.SiteRule.TagFilter(tag)
	if ruleItem == nil {
		c.JSON(MakeErrorJSON(404, 40400, "error tag"))
		return
	}
	page, _ := strconv.Atoi(c.Query("page"))
	rpp, _ := strconv.Atoi(c.Query("rpp"))
	_, resp := post.GetPostsFromDB(d.DB, tag, page, rpp)
	c.JSON(MakeSuccessJSON(&map[string]any{
		"Title": ruleItem.Site,
		"Url":   ruleItem.Url,
		"Tag":   ruleItem.Tag,
		"data":  resp,
	}, "success"))
	return
}

func (d DataController) GetFeedHandler(c *gin.Context) {
	tag := c.Query("tag")
	ruleItem := d.RuleConfig.SiteRule.TagFilter(tag)
	if ruleItem == nil {
		c.JSON(404, &map[string]any{
			"error": 40400,
			"msg":   "error tag",
		})
		return
	}
	data, _ := post.GetPostsFromDB(d.DB, tag, 1, 20)
	c.XML(http.StatusOK, atom.Feed{
		XMLName: xml.Name{},
		Title:   ruleItem.Site,
		ID:      "hdu_" + ruleItem.Tag,
		Link:    []atom.Link{{Href: ruleItem.Url}},
		Entry:   post.Posts(data).ToAtomFeeds(),
		Updated: atom.Time(time.Now()),
	})
	return
}
