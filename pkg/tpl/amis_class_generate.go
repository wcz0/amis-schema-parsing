package tpl

import (
	"amis_schema_parsing/pkg/util"
	"strings"
)

// AmisClassGenerate 生成 amisClass
func AmisClassGenerate() string {
	// 读取 Renderers 目录下的文件名

	files := util.ReadDir("renderers")

	content := amisClassHeader()

	for _, v := range files {
		fileName := strings.ReplaceAll(v, ".go", "")
		className := util.SnakeToPascal(fileName)

		content += `

	func (a *Amis) ` + className + `() *` + className + `{
		return New` + className + `();
	}`
	}

	return content
}

// amisClassHeader amisClass 头部
func amisClassHeader() string {
	return `package renderers

type Amis struct {}

func NewAmis() *Amis {
	return &Amis{}
}
`
}
