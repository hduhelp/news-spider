package config

import (
	"github.com/pelletier/go-toml"
	"github.com/setcy/spider/logger"
	"reflect"
)

var sInstanceMap = make(map[string]*TConfig)

type TConfig struct {
	mRootTree *toml.Tree
}

func InitConfig(alias string, configPath string) (*TConfig, error) {
	rootTree, err := toml.LoadFile(configPath)
	if err != nil {
		return nil, err
	}
	var tConfig = &TConfig{mRootTree: rootTree}
	sInstanceMap[alias] = tConfig

	return tConfig, nil
}

func MustInitConfig(alias string, configPath string) *TConfig {
	tConfig, err := InitConfig(alias, configPath)
	if err != nil {
		logger.Error.Panicln(err)
		return nil
	}
	return tConfig
}

func (config TConfig) GetRootTree() *toml.Tree {
	return config.mRootTree
}

// Unmarshall 在包含模块名的父节点上获取功能节点的struct或值，$outRef必须传引用
func (config TConfig) Unmarshall(moduleName string, nodeName string, outRef any) error {
	var moduleTree = config.GetRootTree().Get(moduleName).(*toml.Tree)
	var treeType = reflect.TypeOf(moduleTree)
	node := moduleTree.Get(nodeName)
	if reflect.TypeOf(node).ConvertibleTo(treeType) {
		var nodeTree = node.(*toml.Tree)
		err := nodeTree.Unmarshal(outRef)
		return err
	} else {
		reflect.ValueOf(outRef).Elem().Set(reflect.ValueOf(node))
		return nil
	}

}

func GetAsTree(current *toml.Tree, key string) *toml.Tree {
	return current.Get(key).(*toml.Tree)
}
