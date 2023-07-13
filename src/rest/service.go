package rest

import (
	"github.com/BurntSushi/toml"
	"github.com/setcy/spider/config"
	"github.com/setcy/spider/db"
	"github.com/setcy/spider/src/rule"
	"gorm.io/gorm"
)

type Service struct {
	DB        *gorm.DB
	conf      *ServiceConfig
	Rule      *rule.RuleConfig
	RuleArray rule.RuleArray
}

func (s *Service) Init() {
	var configTree = config.MustInitConfig("service", "config/conf.toml")
	s.conf = new(ServiceConfig)
	var err = configTree.GetRootTree().Unmarshal(s.conf)
	if err != nil {
		panic(err)
	}
	s.DB = db.InitDB()

	var ruleConfig = new(rule.RuleConfig)
	_, err = toml.DecodeFile("config/rule/rules.toml", ruleConfig)
	ruleConfig.SiteRule.Item, err = rule.LoadRule("config/rule")
	ruleConfig.SiteRule.UpdateItems()
	if err != nil {
		panic(err)
	}

	s.Rule = ruleConfig
	s.RuleArray.SiteRule = ruleConfig.SiteRule.ToRuleArrayItem(new(rule.RuleItem)) //获取所有tag需要的rule结构

	s.initRouter()
}
