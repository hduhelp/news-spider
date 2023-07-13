package controller

import (
	"github.com/setcy/spider/src/post"
	"github.com/setcy/spider/src/rule"
	"github.com/setcy/spider/src/subscribe"
	"gorm.io/gorm"
)

func RangeData(db *gorm.DB, item *rule.RuleItem) {
	if len(item.Item) == 0 {
		GetPosts(db, item)
	}
	num := post.GetPostsNumFromDB(db, item)
	for _, q := range item.Item {
		RangeData(db, q)
	}
	num2 := post.GetPostsNumFromDB(db, item)
	post.GetNewPost(db, item, num2-num)
	subscribe.GetSubscribesByRule(db, item)
}

func RangeDataWithoutNotify(db *gorm.DB, item *rule.RuleItem) {
	if len(item.Item) == 0 {
		GetPosts(db, item)
	}
	for _, q := range item.Item {
		RangeData(db, q)
	}
}

func GetPosts(db *gorm.DB, item *rule.RuleItem) {
	num := post.GetPostsNumFromDB(db, item)
	posts := post.GetUrls(item)
	for _, v := range posts {
		v.Init(item)
		v.GetPageHTML()
		v.GetGoQueryDocument()
		v.GetContent()
		v.GetTitle()
		v.GetDateStr()
		v.GetTime()
		v.Fix()
		post.SavePost(db, v)
	}
	num2 := post.GetPostsNumFromDB(db, item)
	posts2 := post.GetNewPost(db, item, num2-num)
	subscribe.GetSubscribesByRule(db, item)
	println(item.Tag, len(posts), len(posts2))
}
