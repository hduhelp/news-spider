package rule

import (
	"fmt"
	"regexp"
)

type RuleConfig struct {
	SiteRule *RuleItem
}

type RuleItem struct {
	Type           string `toml:"Type"`
	Site           string
	Section        string
	Url            string `json:"-"`
	ListRuleUrl    string `json:"-"`
	ListRulesTitle string `json:"-"`
	ListRulesDate  string `json:"-"`
	ContentRule    string `json:"-"`
	TittleRule     string
	DateRule       string
	Tag            string
	Item           map[string]*RuleItem `json:",omitempty"`
}

type RuleArray struct {
	SiteRule *RuleArrayItem
}

type RuleArrayItem struct {
	Type           string
	Site           string
	Section        string
	URL            string `json:"-"`
	ListRuleUrl    string `json:"-"`
	ListRulesTitle string `json:"-"`
	ListRulesDate  string `json:"-"`
	ContentRule    string `json:"-"`
	Tag            string
	Item           []*RuleArrayItem `json:"Items,omitempty"`
}

func (r *RuleItem) ToRuleArrayItem(father *RuleItem) *RuleArrayItem {
	var kids []*RuleArrayItem
	if len(r.Item) != 0 {
		for _, v := range r.Item {
			kids = append(kids, v.ToRuleArrayItem(r))
		}
	}
	return &RuleArrayItem{
		Type:    r.Type,
		Site:    father.Site + "·" + r.Site,
		Section: r.Section,
		Tag:     r.Tag,
		Item:    kids,
	}
}

func (r *RuleItem) TagFilter(tag string) *RuleItem {
	tagArray := regexp.MustCompile(`[A-Za-z]{1,30}`).FindAllStringSubmatch(tag, -1)
	rule := r
	fmt.Println(tagArray)
	for _, path := range tagArray {
		rule = rule.Item[path[0]]
	}
	if rule.Tag != tag {
		return nil
	}
	return rule
}

func (r *RuleItem) ToReptileRule() *RuleItem {
	data := new(RuleItem)
	data2 := r
	rule := data2
	tag := regexp.MustCompile(`[A-Za-z]{2,30}`).FindAllStringSubmatch(r.Tag, -1)
	for _, path := range tag {
		rule = rule.Item[path[0]]
		if rule.ContentRule != "" {
			data.ContentRule = rule.ContentRule
		}
		if rule.ListRulesDate != "" {
			data.ListRulesDate = rule.ListRulesDate
		}
		if rule.ListRulesTitle != "" {
			data.ListRulesTitle = rule.ListRulesTitle
		}
		if rule.ListRuleUrl != "" {
			data.ListRuleUrl = rule.ListRuleUrl
		}
		if rule.TittleRule != "" {
			data.TittleRule = rule.TittleRule
		}
		if rule.DateRule != "" {
			data.DateRule = rule.DateRule
		}
	}
	rule.ListRuleUrl = data.ListRuleUrl
	rule.ListRulesTitle = data.ListRulesTitle
	rule.ListRulesDate = data.ListRulesDate
	rule.ContentRule = data.ContentRule
	rule.DateRule = data.DateRule
	rule.TittleRule = data.TittleRule
	return rule
}

func (r *RuleItem) UpdateItems() {
	for _, v := range r.Item {
		if v.ContentRule == "" {
			v.ContentRule = r.ContentRule
		}
		if v.ListRulesDate == "" {
			v.ListRulesDate = r.ListRulesDate
		}
		if v.ListRulesTitle == "" {
			v.ListRulesTitle = r.ListRulesTitle
		}
		if v.ListRuleUrl == "" {
			v.ListRuleUrl = r.ListRuleUrl
		}
		if v.DateRule == "" {
			v.DateRule = r.DateRule
		}
		if v.TittleRule == "" {
			v.TittleRule = r.TittleRule
		}
		v.UpdateItems()
	}
}
