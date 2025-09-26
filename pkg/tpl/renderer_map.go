package tpl

import (
	"amis_schema_parsing/pkg/util"
)

// RendererMap 生成 RendererMap
func RendererMap(schemaType interface{}, definitions map[string]interface{}) string {
	list := schemaType.(map[string]interface{})["enum"]

	// 所有类型
	var listString []string

	for _, v := range list.([]interface{}) {
		listString = append(listString, v.(string))
	}

	// 类型对应的类名
	data := make(map[string]string)

	for key, value := range definitions {
		className := util.GetClassName(key)

		if valueMap, ok := value.(map[string]interface{}); ok {
			for _, k := range extractSchemaTypes(valueMap) {
				data[k] = className
			}
		}
	}

	// 补充额外的
	data = appendExtra(data)

	// 文件头
	content := header()

	for _, v := range listString {
		content += "\t\t'" + v + "' => "

		if data[v] != "" {
			content += data[v] + "::class"
		} else {
			content += "Component::class"
		}
		content += ",\n"
	}

	content += footer()

	return content
}

func GetMap(schemaType interface{}, definitions map[string]interface{}) map[string]string {
	list := schemaType.(map[string]interface{})["enum"]

	// 所有类型
	var listString []string

	for _, v := range list.([]interface{}) {
		listString = append(listString, v.(string))
	}

	// 类型对应的类名
	data := make(map[string]string)

	for key, value := range definitions {
		className := util.GetClassName(key)

		if valueMap, ok := value.(map[string]interface{}); ok {
			for _, k := range extractSchemaTypes(valueMap) {
				data[k] = className
			}
		}
	}

	return appendExtra(data)
}

// 提取 schema 中的类型字符串，兼容新旧结构
func extractSchemaTypes(valueMap map[string]interface{}) []string {
	types := []string{}

	// 处理新版 schema 结构 (allOf)
	if allOf, ok := valueMap["allOf"].([]interface{}); ok {
		for _, item := range allOf {
			if itemMap, ok := item.(map[string]interface{}); ok {
				// 处理 allOf 中的 properties
				if props, exists := itemMap["properties"]; exists && props != nil {
					extractFromProperties(props, &types)
				}
				
				// 处理 allOf 中的 required 和 type
				if typeVal, exists := itemMap["type"]; exists && typeVal != nil {
					if typeStr, ok := typeVal.(string); ok && typeStr == "object" {
						continue // 跳过 object 类型
					}
					
					// 处理 const 类型
					if constVal, exists := itemMap["const"]; exists && constVal != nil {
						if constStr, ok := constVal.(string); ok {
							appendUnique(&types, constStr)
						}
					}
				}
			}
		}
	}

	// 处理旧版 schema 结构 (直接 properties)
	if props, exists := valueMap["properties"]; exists && props != nil {
		extractFromProperties(props, &types)
	}

	// 处理顶层 type 字段
	if typeVal, exists := valueMap["type"]; exists && typeVal != nil {
		if typeMap, ok := typeVal.(map[string]interface{}); ok {
			// const 情况
			if c, ok := typeMap["const"].(string); ok && c != "" {
				appendUnique(&types, c)
			}
			// enum 情况
			if enumVal, ok := typeMap["enum"].([]interface{}); ok && len(enumVal) > 0 {
				if s, ok := enumVal[0].(string); ok {
					appendUnique(&types, s)
				}
			}
		} else if typeStr, ok := typeVal.(string); ok && typeStr != "object" {
			// 直接是字符串类型值
			appendUnique(&types, typeStr)
		}
	}

	return types
}

// 从 properties 节点提取类型
func extractFromProperties(props interface{}, types *[]string) {
	if propsMap, ok := props.(map[string]interface{}); ok {
		// 处理 properties.type 字段
		if typeVal, exists := propsMap["type"]; exists && typeVal != nil {
			if typeMap, ok := typeVal.(map[string]interface{}); ok {
				// const 情况
				if c, ok := typeMap["const"].(string); ok && c != "" {
					appendUnique(types, c)
				}
				// enum 情况，取第一个即可（与旧逻辑保持一致）
				if enumVal, ok := typeMap["enum"].([]interface{}); ok && len(enumVal) > 0 {
					if s, ok := enumVal[0].(string); ok {
						appendUnique(types, s)
					}
				}
				// anyOf 情况
				if anyOf, ok := typeMap["anyOf"].([]interface{}); ok {
					for _, v := range anyOf {
						if m, ok := v.(map[string]interface{}); ok {
							if c2, ok := m["const"].(string); ok && c2 != "" {
								appendUnique(types, c2)
							}
							if enum2, ok := m["enum"].([]interface{}); ok && len(enum2) > 0 {
								if s2, ok := enum2[0].(string); ok {
									appendUnique(types, s2)
								}
							}
						}
					}
				}
			} else if typeStr, ok := typeVal.(string); ok && typeStr != "object" {
				// 直接是字符串类型值
				appendUnique(types, typeStr)
			}
		}
	}
}

// 去重追加
func appendUnique(arr *[]string, val string) {
	for _, v := range *arr {
		if v == val {
			return
		}
	}
	*arr = append(*arr, val)
}

// 文件头
func header() string {
	return `<?php

namespace Slowlyo\OwlAdmin\Renderers;

class RendererMap
{
	public static array $map = [
`
}

// 文件尾
func footer() string {
	return `	];
}
`
}

// 追加额外的类型
func appendExtra(data map[string]string) map[string]string {
	data["breadcrumb"] = "Breadcrumb"
	data["custom"] = "Custom"
	data["date"] = "Date"
	data["static-date"] = "StaticExactControl"
	data["datetime"] = "Date"
	data["static-datetime"] = "StaticExactControl"
	data["time"] = "Date"
	data["static-time"] = "StaticExactControl"
	data["month"] = "Date"
	data["static-month"] = "StaticExactControl"
	data["flex"] = "Flex"
	data["image"] = "Image"
	data["static-image"] = "StaticExactControl"
	data["list"] = "ListRenderer"
	data["map"] = "Mapping"
	data["mapping"] = "Mapping"
	data["markdown"] = "Markdown"
	data["property"] = "Property"
	data["qrcode"] = "QRCode"
	data["qr-code"] = "QRCode"
	data["barcode"] = "Barcode"
	data["table"] = "Table"
	data["static-table"] = "StaticExactControl"
	data["html"] = "Html"
	data["tpl"] = "Tpl"
	data["web-component"] = "WebComponent"
	data["button"] = "VanillaAction"
	data["submit"] = "VanillaAction"
	data["reset"] = "VanillaAction"
	data["chart-radios"] = "ChartRadios"
	data["input-date-range"] = "DateRangeControl"
	data["input-time-range"] = "DateRangeControl"
	data["input-datetime-range"] = "DateRangeControl"
	data["input-excel"] = "InputExcel"
	data["editor"] = "EditorControl"
	data["bat-editor"] = "EditorControl"
	data["c-editor"] = "EditorControl"
	data["coffeescript-editor"] = "EditorControl"
	data["cpp-editor"] = "EditorControl"
	data["csharp-editor"] = "EditorControl"
	data["css-editor"] = "EditorControl"
	data["dockerfile-editor"] = "EditorControl"
	data["fsharp-editor"] = "EditorControl"
	data["go-editor"] = "EditorControl"
	data["handlebars-editor"] = "EditorControl"
	data["html-editor"] = "EditorControl"
	data["ini-editor"] = "EditorControl"
	data["java-editor"] = "EditorControl"
	data["javascript-editor"] = "EditorControl"
	data["json-editor"] = "EditorControl"
	data["less-editor"] = "EditorControl"
	data["lua-editor"] = "EditorControl"
	data["markdown-editor"] = "EditorControl"
	data["msdax-editor"] = "EditorControl"
	data["objective-c-editor"] = "EditorControl"
	data["php-editor"] = "EditorControl"
	data["plaintext-editor"] = "EditorControl"
	data["postiats-editor"] = "EditorControl"
	data["powershell-editor"] = "EditorControl"
	data["pug-editor"] = "EditorControl"
	data["python-editor"] = "EditorControl"
	data["r-editor"] = "EditorControl"
	data["razor-editor"] = "EditorControl"
	data["ruby-editor"] = "EditorControl"
	data["sb-editor"] = "EditorControl"
	data["scss-editor"] = "EditorControl"
	data["sol-editor"] = "EditorControl"
	data["sql-editor"] = "EditorControl"
	data["swift-editor"] = "EditorControl"
	data["typescript-editor"] = "EditorControl"
	data["vb-editor"] = "EditorControl"
	data["xml-editor"] = "EditorControl"
	data["yaml-editor"] = "EditorControl"
	data["fieldset"] = "FieldSetControl"
	data["fieldSet"] = "FieldSetControl"
	data["input-text"] = "TextControl"
	data["input-password"] = "TextControl"
	data["input-email"] = "TextControl"
	data["input-url"] = "TextControl"
	data["native-date"] = "TextControl"
	data["native-time"] = "TextControl"
	data["native-number"] = "TextControl"
	data["table-view"] = "TableView"
	data["grid-nav"] = "GridNav"
	data["select"] = "SelectControl"
	data["multi-select"] = "SelectControl"
	data["code"] = "Code"
	data["plain"] = "Plain"

	return data
}
