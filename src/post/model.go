package post

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/setcy/spider/src/rule"
	"gorm.io/gorm"
	"io"
	"time"
)

type Post struct {
	gorm.Model `json:"-"`
	Content    string `sql:"type:text;"`
	Url        string
	DateString string `sql:"-" json:"-"`
	Date       time.Time
	Tag        string
	Title      string
	Rule       *rule.RuleItem    `json:"-" gorm:"-"`
	PageHTML   io.Reader         `json:"-" gorm:"-"`
	Document   *goquery.Document `json:"-" gorm:"-"`
}
