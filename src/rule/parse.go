package rule

import (
	"github.com/pelletier/go-toml"
	"github.com/setcy/spider/config"
	"github.com/setcy/spider/logger"
	"os"
	"path/filepath"
	"regexp"
)

type Site struct {
	ItemMap map[string]SiteItem
}

type SiteItem struct {
	Type string
}

func LoadRule(ruleDir string) (map[string]*RuleItem, error) {
	var resultMap = make(map[string]*RuleItem)
	var err = filepath.Walk(ruleDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		var ruleFileNameExpr = regexp.MustCompile("^rule_.+\\.toml$")
		if !ruleFileNameExpr.MatchString(info.Name()) {
			return nil
		}
		ruleMap, err := ParseRule(path)
		if err != nil {
			return err
		}
		for school, rule := range ruleMap {
			resultMap[school] = rule
		}
		return nil
	})
	return resultMap, err
}

func ParseRule(rulePath string) (map[string]*RuleItem, error) {
	var ruleMap = make(map[string]*RuleItem)
	tree, err := toml.LoadFile(rulePath)
	if err != nil {
		panic(err)
	}
	var schoolListTree = config.GetAsTree(config.GetAsTree(tree, "SiteRule"), "Item")
	for _, school := range schoolListTree.Keys() {
		var schoolTree = config.GetAsTree(schoolListTree, school)
		var ruleItem RuleItem
		err := schoolTree.Unmarshal(&ruleItem)
		if err != nil {
			return nil, err
		}
		ruleItem.UpdateItems() //统一模板，将列html信息复制到下级Item
		ruleMap[school] = &ruleItem
		logger.Info.Println(ruleItem.Site, ruleItem)
	}

	var treeMap = schoolListTree.ToMap()
	logger.Info.Println(treeMap)
	return ruleMap, nil
}
