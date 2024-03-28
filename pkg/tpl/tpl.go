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
	t.Content = `package com.wcz0.renderers;`
}

// Replace 处理
func (t *Tpl) Replace() {
	t.init()
	t.doNamespace()
	t.doClass()
	t.doDoc()
	t.replaceUrl()
}

// doNamespace 处理命名空间
func (t *Tpl) doNamespace() {
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
public class ` + t.ClassName + ` extends BaseRenderer {

	publice ` + t.ClassName + `() {
		this.set("type", "` + util.PascalToCamel(t.ClassName) + `");
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
	properties := t.Data.(map[string]interface{})["properties"].(map[string]interface{})

	// 补充方法
	properties = addSomeMethod(t.ClassName, properties)

	for key, value := range properties {
		if key != "$ref" {
			t.Content += `
	/**
	`
			if value.(map[string]interface{})["description"] != nil {
				t.Content += " * " + util.ClearLineBreak(value.(map[string]interface{})["description"].(string)) + "\n"
			}

			// 将enum中的值拼接到description后面
			if value.(map[string]interface{})["enum"] != nil {
				t.Content += "    * 可选值: "
				for _, v := range value.(map[string]interface{})["enum"].([]interface{}) {
					// 获取v的类型
					switch v.(type) {
					case bool:
						t.Content += fmt.Sprintf("%v", v.(bool)) + " | "
					default:
						t.Content += fmt.Sprintf("%v", v) + " | "
					}
				}
				t.Content = strings.TrimSuffix(t.Content, " | ") // 去掉最后的 " | "
				t.Content += "\n"
			}

			t.Content += `    */
	public ` + t.ClassName + ` ` + key + `(Object value) {
		return (` + t.ClassName + `) this.set("` + key + `", value);
	}
	`
		}
	}

	t.Content += "}\n"

	t.Content = t.Content[:len(t.Content)-1]
}

// doContent 处理内容
func (t *Tpl) doContent() {
	if t.Data.(map[string]interface{})["properties"] == nil {
		return
	}

	// properties := t.Data.(map[string]interface{})["properties"].(map[string]interface{})
	// PHP 映射包
	// for key, value := range properties {
	// 	// 必须的属性
	// 	_required := t.Data.(map[string]interface{})["required"]
	// 	if _required != nil {
	// 		for _, v := range _required.([]interface{}) {
	// 			// 查找对应属性的 const 或 enum
	// 			if key == v.(string) {
	// 				if value.(map[string]any)["const"] != nil {
	// 					t.Content += "    public string $" + v.(string) + " = '" + value.(map[string]any)["const"].(string) + "';"
	// 					t.Content += "\n"
	// 				} else {
	// 					if value.(map[string]any)["enum"] != nil {
	// 						t.Content += "    public string $" + v.(string) + " = '" + value.(map[string]any)["enum"].([]interface{})[0].(string) + "';"
	// 						t.Content += "\n"
	// 					}
	// 				}
	// 			}
	// 		}
	// 	}
	// }
	// PHP 映射包2
	// 	t.Content += "class AjaxAction extends BaseRenderer\n{\n"
	// t.Content += "    public function __construct()\n    {\n"

	// for key, value := range properties {
	// 	// 必须的属性
	// 	_required := t.Data.(map[string]interface{})["required"]
	// 	if _required != nil {
	// 		for _, v := range _required.([]interface{}) {
	// 			// 查找对应属性的 const 或 enum
	// 			if key == v.(string) {
	// 				if value.(map[string]any)["const"] != nil {
	// 					t.Content += "        $this->set('" + v.(string) + "', '" + value.(map[string]any)["const"].(string) + "');\n"
	// 				} else {
	// 					if value.(map[string]any)["enum"] != nil {
	// 						t.Content += "        $this->set('" + v.(string) + "', '" + value.(map[string]any)["enum"].([]interface{})[0].(string) + "');\n"
	// 					}
	// 				}
	// 			}
	// 		}
	// 	}
	// }

	// t.Content += "    }\n}"
	// GO 映射包

	// for key, value := range properties {
	// 	// 必须的属性
	// 	_required := t.Data.(map[string]interface{})["required"]
	// 	if _required != nil {
	// 		for _, v := range _required.([]interface{}) {
	// 			// 查找对应属性的 const 或 enum
	// 			if key == v.(string) {
	// 				if value.(map[string]interface{})["const"] != nil {
	// 					t.Content += "    a.Set(\"" + v.(string) + "\", \"" + value.(map[string]interface{})["const"].(string) + "\")\n"
	// 				} else {
	// 					if value.(map[string]interface{})["enum"] != nil {
	// 						t.Content += "    a.Set(\"" + v.(string) + "\", \"" + value.(map[string]interface{})["enum"].([]interface{})[0].(string) + "\")\n"
	// 					}
	// 				}
	// 			}
	// 		}
	// 	}
	// }
	// t.Content += "}\n"

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
