package config

import "strings"

const (
	MacOS = "darwin"
)

type GetReplaceNameFn func(string) string

type ReplaceContentItem struct {
	Key string
	Fn  GetReplaceNameFn
}

type NewProjectInfo struct {
	TplRepo        string
	ReplaceContent []ReplaceContentItem
	ReplaceFile    map[string]string
	ReplaceDir     map[string]string
}

var NewOdinInfo NewProjectInfo = NewProjectInfo{
	TplRepo:        "https://github.com/tal-tech/odin.git",
	ReplaceContent: []ReplaceContentItem{ReplaceContentItem{"encoding", skipTemplateName}, ReplaceContentItem{"odinPlugin", skipTemplateNameOdin}, ReplaceContentItem{"odin", DefaultReplaceName}, ReplaceContentItem{"Odin", TitleReplaceName}, ReplaceContentItem{"#TemplateName#", recoverEncoding}, ReplaceContentItem{"#TemplateName2#", recoverOdinPlugin}},
	ReplaceFile:    map[string]string{"odin": ""},
	ReplaceDir:     map[string]string{"odin": ""},
}
var NewHsfInfo NewProjectInfo = NewProjectInfo{
	TplRepo:        "git@gitlab.alibaba-inc.com:amap-car-lbs/hsfDemoProject.git",
	ReplaceContent: []ReplaceContentItem{{"hsfDemoProject", DefaultReplaceName}, {"hsfDemoProject", TitleReplaceName}},
	ReplaceDir:     map[string]string{"hsfDemoProject": ""},
	ReplaceFile:    map[string]string{"hsfDemoProject.release": ".release"},
}

var NewGaeaInfo NewProjectInfo = NewProjectInfo{
	TplRepo:        "https://github.com/tal-tech/gaea.git",
	ReplaceContent: []ReplaceContentItem{ReplaceContentItem{"gaea", DefaultReplaceName}, ReplaceContentItem{"Gaea", TitleReplaceName}},
	ReplaceFile:    map[string]string{"gaea": ""},
	ReplaceDir:     map[string]string{"gaea": ""},
}

var NewTritonInfo NewProjectInfo = NewProjectInfo{
	TplRepo:        "https://github.com/tal-tech/triton.git",
	ReplaceContent: []ReplaceContentItem{ReplaceContentItem{"triton", DefaultReplaceName}},
	ReplaceFile:    map[string]string{},
	ReplaceDir:     map[string]string{},
}

var NewPanInfo NewProjectInfo = NewProjectInfo{
	TplRepo:        "https://github.com/tal-tech/pan.git",
	ReplaceContent: []ReplaceContentItem{{"panic", skipTemplateName}, {"pan", DefaultReplaceName}, {"#TemplateName#", recoverPanic}},
	ReplaceFile:    map[string]string{},
	ReplaceDir:     map[string]string{},
}

var NewJobInfo NewProjectInfo = NewProjectInfo{
	TplRepo:        "https://github.com/tal-tech/hera.git",
	ReplaceContent: []ReplaceContentItem{},
	ReplaceFile:    map[string]string{},
	ReplaceDir:     map[string]string{},
}

func DefaultReplaceName(in string) string {
	return in
}

func TitleReplaceName(in string) string {
	return strings.Title(in)
}

func skipTemplateName(in string) string {
	return "#TemplateName#"
}

func skipTemplateNameOdin(in string) string {
	return "#TemplateName2#"
}

func recoverEncoding(in string) string {
	return "encoding"
}

func recoverPanic(in string) string {
	return "panic"
}

func recoverOdinPlugin(in string) string {
	return "odinPlugin"
}

func DefaultXormConfig() map[string]string {
	return map[string]string{
		"lang":    "go",
		"genJson": "1",
		"prefix":  "cos_",
	}
}
