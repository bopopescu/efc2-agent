package collector

import "github.com/efc2/efc2-agent/common/plugin"

// Checker XXX
type Checker func(conf plugin.InitConfig) plugin.Plugin

// Plugins XXX
var Plugins = map[string]Checker{}

// Add XXX
func Add(name string, checker Checker) {
	Plugins[name] = checker
}
