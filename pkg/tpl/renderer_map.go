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
		// 类名
		className := key

		// 如果是以 Schema 或者 Object 结尾的则去掉后缀
		if util.IsEndWith(key, "Schema") || util.IsEndWith(key, "Object") {
			className = key[:len(key)-6]
		}

		// 关键字冲突
		if className == "List" {
			className = "ListRenderer"
		}

		// 判断是否存在 value.properties.type.const , 如果存在则取出值
		if value.(map[string]interface{})["properties"] != nil {
			valueProperties := value.(map[string]interface{})["properties"].(map[string]interface{})

			if valueProperties["type"] != nil {
				valueType := valueProperties["type"].(map[string]interface{})

				if valueType["const"] != nil {
					data[valueType["const"].(string)] = className
				}
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

		// 判断是否存在 value.properties.type.const , 如果存在则取出值
		if value.(map[string]interface{})["properties"] != nil {
			valueProperties := value.(map[string]interface{})["properties"].(map[string]interface{})

			if valueProperties["type"] != nil {
				valueType := valueProperties["type"].(map[string]interface{})

				if valueType["const"] != nil {
					data[valueType["const"].(string)] = className
				}
			}
		}
	}

	return appendExtra(data)
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
