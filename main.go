package main

import (
	"amis_schema_parsing/pkg/tpl"
	"amis_schema_parsing/pkg/util"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func main() {
	generateRenderClass()
	//getResidueData()
}

// getResidueData 获取剩余的数据
func getResidueData() {
	definitions := ReadJsonFile()

	for key, _ := range definitions {
		if util.IsEndWith(key, "Schema") || util.IsEndWith(key, "Object") || key == "Option" {
			delete(definitions, key)
		}
	}

	file, _ := json.MarshalIndent(definitions, "", " ")

	_ = os.WriteFile("./schema_new.json", file, 0644)
}

// generateRenderClass 生成render类
func generateRenderClass() {
	timeStart := time.Now()

	definitions := ReadJsonFile()

	// 排除的
	exclude := []string{
		"Option",
		"VanillaAction",
		"SchemaMessage",
		"ListenerAction",
		"ImageToolbarAction",
		"NavOverflow",
		"QRCodeImageSettings",
		"ComboCondition",
	}

	// 只保留 key 以 Schema 结尾 | Object 结尾 | 不在 exclude 中的
	for key, _ := range definitions {
		if !util.IsEndWith(key, "Schema") && !util.IsEndWith(key, "Object") && !util.IsInArray(key, exclude) {
			delete(definitions, key)
		}
	}

	// 只保留 properties 不为空的
	for key, value := range definitions {
		if value.(map[string]interface{})["properties"] == nil {
			delete(definitions, key)
			continue
		}
	}

	schemaType := ReadJsonFile()["SchemaType"]

	schemaMap := tpl.GetMap(schemaType, definitions)

	// 写入文件 生成所有 Render 类
	WriteFile(definitions, schemaMap)

	// 生成映射
	// rendererMap := tpl.RendererMap(schemaType, definitions)

	// _ = os.WriteFile("./renderers/RendererMap.php", []byte(rendererMap), 0644)

	// 生成 amisClass, amisMake()
	amisClass := tpl.AmisClassGenerate()

	_ = os.WriteFile("./dist/amis.go", []byte(amisClass), 0644)

	fmt.Println("耗时：", time.Since(timeStart))
}

// WriteFile 写入文件
func WriteFile(data map[string]interface{}, schemaMap map[string]string) {
	// 检查目录
	util.InitDir("./dist/renderers")

	// 先复制补充文件
	files, _ := os.ReadDir("./origins")

	for _, file := range files {
		util.CopyFile("./origins/"+file.Name(), "./dist/renderers/"+file.Name())
	}

	// 生成文件
	for key, value := range data {
		// 类名
		className := util.GetClassName(key)

		// 替换数据
		t := tpl.Tpl{
			ClassName: className,
			Data:      value,
			SchemaMap: schemaMap,
		}

		// 生成内容
		t.Replace()

		// 写入文件
		_ = os.WriteFile("./dist/renderers/"+util.PascalToSnake(className)+".go", []byte(t.Content), 0644)
	}
}

// ReadJsonFile 读取json文件
func ReadJsonFile() map[string]interface{} {
	fileContent, err := os.ReadFile("./schema-new.json")

	if err != nil {
		fmt.Println(err)
	}

	// 解析json内容
	var data map[string]interface{}
	err = json.Unmarshal(fileContent, &data)

	if err != nil {
		fmt.Println(err)
	}

	// 确保返回的是map[string]interface{}类型
	if definitions, ok := data["definitions"].(map[string]interface{}); ok {
		return definitions
	} else {
		fmt.Println("警告: 无法将definitions转换为map[string]interface{}")
		return make(map[string]interface{})
	}
}
