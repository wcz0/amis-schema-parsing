package tpl

import (
	"amis_schema_parsing/pkg/util"
	"fmt"
	"strings"
)

type Tpl struct {
	Content   string
	Data      interface{}
	ClassName string
	SchemaMap map[string]string
}

var NeedUploadMap = []string{
	"FileControl",
	"ImageControl",
	"RichTextControl",
}

// init 文件头
func (t *Tpl) init() {
	t.Content = `package renderers

`
}

// Replace 处理
func (t *Tpl) Replace() {
	t.init()
	// t.doNamespace()
	t.doClass()
	t.doDoc()
	t.replaceUrl()
}

// doNamespace 处理命名空间
func (t *Tpl) doNamespace() {
	t.Content += `
`

	// 判断是否需要上传文件
	// 	for _, v := range NeedUploadMap {
	// 		if t.ClassName == v {
	// 			t.Content += `
	// use Slowlyo\OwlAdmin\Traits\Uploader;
	// `
	// 		}
	// 	}
}

// doDoc 处理文档
func (t *Tpl) doDoc() {

	t.doMethod()
}

// doClass 处理类
func (t *Tpl) doClass() {
	t.Content += `
/**
`
	if t.Data.(map[string]interface{})["description"] != nil {
		t.Content += " * " + util.ClearLineBreak(t.Data.(map[string]interface{})["description"].(string)) + "\n"
	}

	t.Content += `
 * @author wcz0
 * @version 6.2.2
 */
`

	t.Content += `type ` + t.ClassName + ` struct {
	*BaseRenderer
}

`
	// 判断是否需要上传文件
	// 	for _, v := range NeedUploadMap {
	// 		if t.ClassName == v {
	// 			t.Content += `
	// 	use Uploader;

	// `
	// 		}
	// 	}

	t.doContent()
}

// doMethod 处理方法
func (t *Tpl) doMethod() {
	dataMap, ok := t.Data.(map[string]interface{})
	if !ok {
		fmt.Println("警告: 无法将Data转换为map[string]interface{}")
		return
	}

	propertiesVal, exists := dataMap["properties"]
	if !exists || propertiesVal == nil {
		fmt.Println("警告: properties不存在或为nil")
		return
	}

	properties, ok := propertiesVal.(map[string]interface{})
	if !ok {
		fmt.Println("警告: 无法将properties转换为map[string]interface{}, 实际类型:", fmt.Sprintf("%T", propertiesVal))
		return
	}

	// 补充方法
	properties = addSomeMethod(t.ClassName, properties)

	exist := []string{}

	// GO 映射包
	for key, value := range properties {
		for _, v := range exist {
			if v == key {
				continue
			}
		}

		if key != "$ref" && key != "$$id"  {
			exist = append(exist, key)
			// if t.ClassName == "RatingControl" && key == "readOnly" {
			// 	continue
			// }
			// if t.ClassName == "Tpl" && key == "testIdBuilder" {
			// 	continue
			// }

			t.Content += "/**\n"
			
			valueMap, ok := value.(map[string]interface{})
			if !ok {
				// 处理value不是map的情况
				t.Content += " * " + key + "\n"
			} else {
				// 处理description
				if descVal, exists := valueMap["description"]; exists && descVal != nil {
					if descStr, ok := descVal.(string); ok {
						t.Content += " * " + util.ClearLineBreak(descStr) + "\n"
					}
				}

				// 将enum中的值拼接到description后面
				if enumVal, exists := valueMap["enum"]; exists && enumVal != nil {
					if enumArr, ok := enumVal.([]interface{}); ok {
						t.Content += " * 可选值: "
						for _, v := range enumArr {
							// 获取v的类型
							switch v := v.(type) {
							case bool:
								t.Content += fmt.Sprintf("%v", v) + " | "
							default:
								t.Content += fmt.Sprintf("%v", v) + " | "
							}
						}
						t.Content = strings.TrimSuffix(t.Content, " | ") // 去掉最后的 " | "
						t.Content += "\n"
					}
				}
			}

			t.Content += " */\n"
			t.Content += "func (a *" + t.ClassName + ") " + strings.Title(key) + "(value interface{}) *" + t.ClassName + " {\n"
			t.Content += "    a.Set(\"" + key + "\", value)\n"
			t.Content += "    return a\n"
			t.Content += "}\n\n"
		}
	}

	if len(t.Content) > 0 {
		t.Content = t.Content[:len(t.Content)-1]
	}
}

// doContent 处理内容
func (t *Tpl) doContent() {
	dataMap, ok := t.Data.(map[string]interface{})
	if !ok {
		fmt.Println("警告: 无法将Data转换为map[string]interface{}")
		return
	}

	propertiesVal, exists := dataMap["properties"]
	if !exists || propertiesVal == nil {
		fmt.Println("警告: properties不存在或为nil")
		return
	}

	properties, ok := propertiesVal.(map[string]interface{})
	if !ok {
		fmt.Println("警告: 无法将properties转换为map[string]interface{}, 实际类型:", fmt.Sprintf("%T", propertiesVal))
		return
	}

	// GO 映射包
	t.Content += "func New" + t.ClassName + "() *" + t.ClassName + " {\n"
	t.Content += "    a := &" + t.ClassName + "{\n"
	t.Content += "        BaseRenderer: NewBaseRenderer(),\n"
	t.Content += "    }\n"
	t.Content += "\n"

	for key, value := range properties {
		// 必须的属性
		requiredVal, exists := dataMap["required"]
		if exists && requiredVal != nil {
			requiredArr, ok := requiredVal.([]interface{})
			if ok {
				for _, v := range requiredArr {
					reqStr, ok := v.(string)
					if !ok {
						continue
					}
					
					// 查找对应属性的 const 或 enum
					if key == reqStr {
						valueMap, ok := value.(map[string]interface{})
						if !ok {
							continue
						}
						
						constVal, hasConst := valueMap["const"]
						if hasConst && constVal != nil {
							if constStr, ok := constVal.(string); ok {
								t.Content += "    a.Set(\"" + reqStr + "\", \"" + constStr + "\")\n"
							}
						} else {
							enumVal, hasEnum := valueMap["enum"]
							if hasEnum && enumVal != nil {
								enumArr, ok := enumVal.([]interface{})
								if ok && len(enumArr) > 0 {
									if enumStr, ok := enumArr[0].(string); ok {
										t.Content += "    a.Set(\"" + reqStr + "\", \"" + enumStr + "\")\n"
									}
								}
							}
						}
					}
				}
			}
		}
	}

	t.Content += "    return a\n"
	t.Content += "}\n"
	t.Content += "\n\n"
	t.Content += "func (a *" + t.ClassName + ") Set(name string, value interface{}) *" + t.ClassName + " {\n"
	t.Content += "    if name == \"map\" {\n"
	t.Content += "        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {\n"
	t.Content += "            value = mapOfArrays(v)\n"
	t.Content += "        }\n"
	t.Content += "    }\n"
	t.Content += "    a.AmisSchema[name] = value\n"
	t.Content += "    return a\n"
	t.Content += "}\n"
	t.Content += "\n"



	// TODO: 文件上传
	// 	// 文件上传路径 默认值
	// 	if t.ClassName == "FileControl" {
	// 		t.Content += `
	//     public function __construct()
	//     {
	//         $this->receiver($this->uploadFilePath());
	//     }
	// `
	// 	}

	// 	// 图片上传路径 默认值
	// 	if t.ClassName == "ImageControl" {
	// 		t.Content += `
	//     public function __construct()
	//     {
	//         $this->receiver($this->uploadImagePath());
	//     }
	// `
	// 	}

	// 	// 富文本默认值
	// 	if t.ClassName == "RichTextControl" {
	// 		t.Content += `
	//     public function __construct()
	//     {
	//         $this->receiver($this->uploadRichPath());
	//         $this->videoReceiver($this->uploadRichPath());
	//         // tinymce 暂时无法使用, 等待官方修复, 默认使用 froala
	//         $this->vendor('froala');
	//     }
	// `
	// 	}

	// 表格默认值
	// 	if t.ClassName == "TableColumn" {
	// 		t.Content += `
	// 	public function type($type)
	// 	{
	// 		$this->type = $type;

	// 		$instance = $this;
	// `
	// 		for k, v := range t.SchemaMap {
	// 			t.Content += fmt.Sprintf(`
	//         /** @noinspection all */
	// 		if ($type == '%s') { /** @var %s $instance */ }
	// `, k, v)
	// 		}

	// 		t.Content += `
	// 		return $instance;
	// 	}
	// `
	// 	}

	t.Content = t.Content[:len(t.Content)-1]
}

// 替换错误的文档地址
func (t *Tpl) replaceUrl() {
	t.Content = strings.Replace(t.Content, "https://baidu.gitee.io/amis/docs", "https://aisuda.bce.baidu.com/amis/zh-CN", -1)
}
