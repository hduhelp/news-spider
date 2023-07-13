package post

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/parnurzeal/gorequest"
	"github.com/setcy/spider/src/rule"
	"golang.org/x/tools/blog/atom"
	"io"
	"log"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Posts []*Post

func GetUrls(ruleItem *rule.RuleItem) Posts {
	resp, _, err := gorequest.New().Post(ruleItem.Url).End()
	if err != nil || resp.StatusCode != 200 {
		return nil
	}
	doc, err2 := goquery.NewDocumentFromReader(resp.Body)
	posts := make([]*Post, 0)
	if err2 != nil || doc == nil {
		return nil
	}
	siteUrl, _ := url.Parse(ruleItem.Url)
	doc.Find(ruleItem.ListRuleUrl).Each(func(i int, g *goquery.Selection) {
		postUrlString, _ := g.Attr("href")
		postUrl, _ := url.Parse(postUrlString)
		if postUrl.Host == "" {
			postUrl.Host = siteUrl.Host
		}
		postUrl.Scheme = siteUrl.Scheme
		posts = append(posts,
			&Post{
				Rule: ruleItem,
				Url:  postUrl.String(),
			})
	})
	doc.Find(ruleItem.ListRulesTitle).Each(func(i int, g *goquery.Selection) {
		posts[i].Title = g.Text()
	})
	doc.Find(ruleItem.ListRulesDate).Each(func(i int, g *goquery.Selection) {
		posts[i].DateString = g.Text()
	})
	return posts
}

func (p *Post) GetTime() {
	var _time time.Time
	if p.DateString == "" {
		return
	}
	p.DateString = "20" + p.DateString
	dateString := regexp.MustCompile(`[0-9][0-9][0-9][0-9]-[0-9][0-9]-[0-9][0-9]`).FindStringSubmatch(p.DateString)
	if dateString != nil {
		_time, _ = time.Parse("2006-01-02", dateString[0])
	}
	if dateString == nil {
		dateString = regexp.MustCompile(`[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9]`).FindStringSubmatch(p.DateString)
		_time, _ = time.Parse("2006/01/02", dateString[0])
	}
	p.Date = _time
	p.DateString = _time.Format("2006-01-02")
}

func (p Posts) ToAtomFeeds() []*atom.Entry {
	entries := make([]*atom.Entry, 0)
	for _, v := range p {
		entry := &atom.Entry{
			Title:     v.Title,
			Link:      []atom.Link{{Href: v.Url}},
			Updated:   atom.Time(v.UpdatedAt),
			Published: atom.Time(v.Date),
			Content:   &atom.Text{Body: v.Content},
			ID:        v.Tag + "-" + strconv.Itoa(int(v.ID)),
		}
		entries = append(entries, entry)
	}
	return entries
}

func (p *Post) GetContent() {
	if !(strings.Contains(p.Url, "html") || strings.Contains(p.Url, "htm")) {
		p.Content = p.Url
	}

	data := p.Document.Find(p.Rule.ContentRule)

	noAttr := []string{
		"style",
		"lang",
		"original-src",
		"sudyfile-attr",
		"data-layer",
	}

	data.Find("[class]").Each(func(i int, s *goquery.Selection) {
		s.RemoveClass()
	})
	data.Find("xml").Each(func(i int, s *goquery.Selection) {
		s.Remove()
	})
	for _, v := range noAttr {
		data.Find("[" + v + "]").Each(func(i int, s *goquery.Selection) {
			s.RemoveAttr(v)
		})
	}

	postUrl, _ := url.Parse(p.Url)

	urlArr := []string{
		"src",
		"href",
	}

	for _, v := range urlArr {
		data.Find("[" + v + "]").Each(func(i int, s *goquery.Selection) {
			src, _ := s.Attr(v)
			srcUrl, err := url.Parse(src)
			if err != nil {
				fmt.Println(err)
			}
			if srcUrl.Host == "" {
				n := postUrl
				n.Path = srcUrl.Path
				n.RawQuery = srcUrl.RawQuery
				srcUrl = n
			}
			s.SetAttr(v, srcUrl.String())
		})
	}

	p.Content, _ = data.Html()
}

func (p *Post) GetTitle() *Post {
	data := p.Document.Find(p.Rule.TittleRule)
	if data.Text() != "" {
		p.Title = data.Text()
	}
	return p
}

func (p *Post) GetDateStr() *Post {
	data := p.Document.Find(p.Rule.DateRule)
	if data.Text() != "" {
		p.DateString = data.Text()
	}
	return p
}

func (p *Post) Init(r *rule.RuleItem) *Post {
	p.Rule = r
	p.Tag = r.Tag
	return p
}

func (p *Post) GetPageHTML() *Post {
	resp, _, err := gorequest.New().Post(p.Url).End()
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		fmt.Println("err")
	}
	p.PageHTML = io.Reader(resp.Body)
	return p
}

func (p *Post) GetGoQueryDocument() *Post {
	doc, err2 := goquery.NewDocumentFromReader(p.PageHTML)
	if err2 != nil {
		log.Fatal(err2)
	}
	p.Document = doc
	return p
}

func (p *Post) Fix() *Post {
	if p.Content == "" {
		p.Content = p.Url
	}
	return p
}
