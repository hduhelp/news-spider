package post

import (
	pageable "github.com/FDKevin0/gorm-pageable"
	"github.com/setcy/spider/src/rule"
	"gorm.io/gorm"
)

func SavePost(db *gorm.DB, post *Post) {
	data := new(Post)
	tx := db.Begin()
	tx.Model(&Post{}).Where(&Post{Url: post.Url, Tag: post.Tag}).Find(&data)
	if data.ID != 0 {
		tx.Model(&Post{}).Where(&Post{Url: post.Url}).
			Updates(map[string]any{
				"content":     post.Content,
				"title":       post.Title,
				"date":        post.Date,
				"date_string": post.DateString})
		tx.Commit()
		return
	}
	tx.Save(post)
	tx.Commit()
}

func GetPostsFromDB(db *gorm.DB, Tag string, Page int, rpp int) ([]*Post, *pageable.Response) {
	if Page == 0 {
		Page = 1
	}
	if rpp == 0 {
		rpp = 25
	}
	data := make([]*Post, 0)
	handler := db.Model(&Post{}).
		Where("TAG LIKE ?", Tag+"%").
		Limit(25).Order("date desc")
	resp, err := pageable.PageQuery(Page, rpp, handler, &data)
	if err != nil {
		panic(err)
	}
	return data, resp
}

func GetPostsNumFromDB(db *gorm.DB, ruleItem *rule.RuleItem) int64 {
	var count int64
	db.Model(&Post{}).
		Where("TAG LIKE ?", ruleItem.Tag+"%").
		Count(&count)
	return count
}

func GetNewPost(db *gorm.DB, ruleItem *rule.RuleItem, Num int64) []*Post {
	data := make([]*Post, 0)
	db.Model(&Post{}).
		Where("TAG LIKE ?", ruleItem.Tag+"%").
		Limit(int(Num)).Order("id desc").Find(&data)
	return data
}
