package tpl

import "amis_schema_parsing/pkg/util"

// addSomeMethod 补充一些方法
func addSomeMethod(className string, properties map[string]interface{}) map[string]interface{} {
	// Control 结尾的类
	if util.IsEndWith(className, "Control") {
		properties["width"] = map[string]interface{}{
			"description": "在Table中调整宽度",
		}
	}

	switch className {
	case "SelectControl":
		properties["labelField"] = map[string]interface{}{
			"description": "设置label字段",
		}
		properties["valueField"] = map[string]interface{}{
			"description": "设置value字段",
		}
		properties["searchable"] = map[string]interface{}{
			"description": "是否可搜索",
		}
		properties["autoCheckChildren"] = map[string]interface{}{
			"description": "是否自动选中子节点",
		}
	case "TreeSelectControl":
		properties["labelField"] = map[string]interface{}{
			"description": "设置label字段",
		}
		properties["valueField"] = map[string]interface{}{
			"description": "设置value字段",
		}
		properties["searchable"] = map[string]interface{}{
			"description": "是否可搜索",
		}
		properties["autoCheckChildren"] = map[string]interface{}{
			"description": "是否自动选中子节点",
		}
	case "Operation":
		properties["label"] = map[string]interface{}{
			"description": "设置label",
		}
	case "UrlAction":
		properties["link"] = map[string]interface{}{
			"description": "设置链接",
		}
	case "Table":
		properties["data"] = map[string]interface{}{
			"description": "设置数据",
		}
	case "CRUDTable":
		properties["data"] = map[string]interface{}{
			"description": "设置数据",
		}
	case "DrawerAction":
		properties["align"] = map[string]interface{}{
			"description": "设置对齐方式",
		}
	}

	return properties
}
