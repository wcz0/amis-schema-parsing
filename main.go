package main

import (
    "fmt"
    "os"
    "path/filepath"
    "regexp"
    "sort"
    "strings"

    "amis_schema_parsing/pkg/util"
)

// PhpClass 表示从 PHP 文件解析出来的类信息
type PhpClass struct {
    Name        string
    Parent      string
    Methods     map[string]string // methodName => schemaKey
    TypeDefault string            // 构造函数中设置的默认 type 值
}

func main() {
    if err := generateFromPhpRenderers("./php-renderers", "./dist"); err != nil {
        fmt.Println("生成失败:", err)
        os.Exit(1)
    }
    fmt.Println("生成完成 -> ./dist")
}

// generateFromPhpRenderers 解析指定目录中的 PHP 渲染器类，并生成对应的 Go 渲染器文件
func generateFromPhpRenderers(phpDir, outDir string) error {
    // 1) 扫描 PHP 文件，构建类图
    classes, err := parsePhpClasses(phpDir)
    if err != nil {
        return err
    }

    // 2) 继承展开：把父类的方法与默认 type 透传到子类
    expanded := expandInheritance(classes)

    // 3) 输出到目标目录
    renderersDir := filepath.Join(outDir, "renderers")
    if err := util.InitDir(renderersDir); err != nil {
        return err
    }

    // 3.1) 写入基础 renderer（包含 BaseRenderer、isArrayOfArrays、mapOfArrays）
    if err := writeBaseRenderer(renderersDir); err != nil {
        return err
    }

    // 3.2) 按类生成
    for _, cls := range expanded {
        if cls.Name == "BaseRenderer" { // 跳过基础类
            continue
        }
        if err := writeRendererFile(renderersDir, cls, expanded); err != nil {
            return err
        }
    }

    // 3.3) 生成 amis.go 便捷函数文件
    if err := writeAmisFile(outDir, expanded); err != nil {
        return err
    }

    return nil
}

// parsePhpClasses 解析目录中的所有 PHP 文件
func parsePhpClasses(dir string) (map[string]*PhpClass, error) {
    classes := make(map[string]*PhpClass)

    // 正则：类定义、构造函数内默认 type、方法里的 set('key')
    reClass := regexp.MustCompile(`(?m)^\s*class\s+([A-Za-z_][A-Za-z0-9_]*)\s+extends\s+([A-Za-z_][A-Za-z0-9_]*)`)
    reCtorType := regexp.MustCompile(`(?s)function\s+__construct\s*\(\s*\)\s*\{.*?\$this->set\(\s*'type'\s*,\s*'([^']+)'\s*\)`) // dot-all
    reMethod := regexp.MustCompile(`(?s)public\s+function\s+([A-Za-z_][A-Za-z0-9_]*)\s*\([^)]*\)\s*\{[^{}]*?\$this->set\(\s*'([^']+)'`) // 只收集直接 set 的方法

    err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if info.IsDir() {
            return nil
        }
        if !strings.HasSuffix(strings.ToLower(info.Name()), ".php") {
            return nil
        }
        // 读取文件
        bs, err := os.ReadFile(path)
        if err != nil {
            return err
        }
        content := string(bs)

        // 匹配类
        m := reClass.FindStringSubmatch(content)
        if len(m) < 3 {
            return nil // 非类文件或不符合预期
        }
        name := m[1]
        parent := m[2]

        // 构造类对象
        c := &PhpClass{
            Name:        name,
            Parent:      parent,
            Methods:     map[string]string{},
            TypeDefault: "",
        }

        // 默认 type
        if mt := reCtorType.FindStringSubmatch(content); len(mt) == 2 {
            c.TypeDefault = mt[1]
        }

        // 方法收集
        for _, mm := range reMethod.FindAllStringSubmatch(content, -1) {
            if len(mm) == 3 {
                methodName := mm[1]
                key := mm[2]
                // 过滤魔术、构造、一些非属性方法
                if methodName == "__construct" || strings.HasPrefix(methodName, "__") {
                    continue
                }
                c.Methods[methodName] = key
            }
        }

        classes[name] = c
        return nil
    })

    if err != nil {
        return nil, err
    }
    return classes, nil
}

// expandInheritance 将父类的方法和默认 type 传递到子类（子类同名方法优先覆盖）
func expandInheritance(classes map[string]*PhpClass) map[string]*PhpClass {
    // 递归收集函数
    var collect func(name string, visited map[string]bool) (map[string]string, string)
    collect = func(name string, visited map[string]bool) (map[string]string, string) {
        if visited[name] {
            return map[string]string{}, ""
        }
        visited[name] = true

        cls, ok := classes[name]
        if !ok {
            return map[string]string{}, ""
        }

        methods := map[string]string{}
        t := cls.TypeDefault

        if cls.Parent != "" {
            pm, pt := collect(cls.Parent, visited)
            for k, v := range pm {
                methods[k] = v
            }
            if t == "" {
                t = pt
            }
        }

        // 覆盖父类同名方法
        for k, v := range cls.Methods {
            methods[k] = v
        }

        // 返回合并结果
        return methods, t
    }

    expanded := make(map[string]*PhpClass)
    for name := range classes {
        methods, t := collect(name, map[string]bool{})
        // 复制类信息
        base := *classes[name]
        base.Methods = methods
        base.TypeDefault = t
        expanded[name] = &base
    }
    return expanded
}

// writeBaseRenderer 生成基础 renderer 文件
func writeBaseRenderer(outDir string) error {
    content := `package renderers

import (
    "encoding/json"
)

type BaseRenderer struct {
    Type       string ` + "`json:\"type\"`" + `
    AmisSchema map[string]interface{}
}

func NewBaseRenderer() *BaseRenderer {
    return &BaseRenderer{
        AmisSchema: make(map[string]interface{}),
    }
}

func (b *BaseRenderer) ToJson() (string, error) {
    bytes, err := json.Marshal(b.AmisSchema)
    if err != nil {
        return "", err
    }
    return string(bytes), nil
}

func (b *BaseRenderer) Set(name string, value interface{}) *BaseRenderer {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    b.AmisSchema[name] = value
    return b
}

func (b *BaseRenderer) MarshalJSON() ([]byte, error) {
    byte, err := json.Marshal(b.AmisSchema)
    if err != nil {
        return nil, err
    }
    return byte, nil
}

func (b *BaseRenderer) ToArray() map[string]interface{} {
    return b.AmisSchema
}

// isArrayOfArrays checks if a slice of interfaces contains only slices of interfaces.
func isArrayOfArrays(v []interface{}) bool {
    for _, item := range v {
        if _, ok := item.([]interface{}); !ok {
            return false
        }
    }
    return true
}

// mapOfArrays converts a slice of slices of interfaces to a map of interfaces.
func mapOfArrays(v []interface{}) map[string]interface{} {
    r := make(map[string]interface{})
    for _, v := range v {
        array := v.([]interface{})
        if len(array) >= 2 {
            key, ok1 := array[0].(string)
            if ok1 {
                r[key] = array[1]
            }
        }
    }
    return r
}
`
    return os.WriteFile(filepath.Join(outDir, "base_renderer.go"), []byte(content), 0644)
}

// writeRendererFile 为单个类生成渲染器文件
func writeRendererFile(outDir string, cls *PhpClass, allClasses map[string]*PhpClass) error {
    pkg := "renderers"
    structName := cls.Name

    // 文件名: pascal -> snake
    fileName := util.PascalToSnake(structName) + ".go"

    var sb strings.Builder
    sb.WriteString("package ")
    sb.WriteString(pkg)
    sb.WriteString("\n\n")

    // 结构体定义与构造函数
    sb.WriteString("type ")
    sb.WriteString(structName)
    sb.WriteString(" struct {\n\t*BaseRenderer\n}\n\n")

    // Component 类的特殊处理
    if structName == "Component" {
        // Component 的构造函数可以接受 type 参数
        sb.WriteString("func New")
        sb.WriteString(structName)
        sb.WriteString("(typeStr string) *")
        sb.WriteString(structName)
        sb.WriteString(" {\n\ta := &")
        sb.WriteString(structName)
        sb.WriteString("{\n\t\tBaseRenderer: NewBaseRenderer(),\n\t}\n")
        sb.WriteString("\ta.Set(\"type\", typeStr)\n")
        sb.WriteString("\treturn a\n}\n\n")
    } else {
        // 普通组件的构造函数
        sb.WriteString("func New")
        sb.WriteString(structName)
        sb.WriteString("() *")
        sb.WriteString(structName)
        sb.WriteString(" {\n\ta := &")
        sb.WriteString(structName)
        sb.WriteString("{\n\t\tBaseRenderer: NewBaseRenderer(),\n\t}\n\n")
        if cls.TypeDefault != "" {
            sb.WriteString("\ta.Set(\"type\", \"")
            sb.WriteString(cls.TypeDefault)
            sb.WriteString("\")\n")
        }
        sb.WriteString("\treturn a\n}\n\n")
    }

    // 类型化 Set 方法，便于链式调用
    sb.WriteString("func (a *")
    sb.WriteString(structName)
    sb.WriteString(") Set(name string, value interface{}) *")
    sb.WriteString(structName)
    sb.WriteString(" {\n\tif name == \"map\" {\n\t\tif v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {\n\t\t\tvalue = mapOfArrays(v)\n\t\t}\n\t}\n\ta.AmisSchema[name] = value\n\treturn a\n}\n")

    // Component 类需要添加 SetType 方法
    if structName == "Component" {
        sb.WriteString("\nfunc (a *")
        sb.WriteString(structName)
        sb.WriteString(") SetType(typeStr string) *")
        sb.WriteString(structName)
        sb.WriteString(" {\n\ta.Set(\"type\", typeStr)\n\treturn a\n}\n")
    }

    // 收集要生成的方法
    var allMethods map[string]string
    if structName == "Component" {
        // Component 类需要包含所有组件的方法
        allMethods = make(map[string]string)
        for _, otherCls := range allClasses {
            if otherCls.Name != "BaseRenderer" && otherCls.Name != "Component" {
                for method, key := range otherCls.Methods {
                    // 避免重复方法，优先保留第一个
                    if _, exists := allMethods[method]; !exists {
                        allMethods[method] = key
                    }
                }
            }
        }
        // 添加 Component 自己的方法，但排除 setType（因为我们已经手动添加了 SetType）
        for method, key := range cls.Methods {
            if strings.ToLower(method) != "settype" {
                allMethods[method] = key
            }
        }
    } else {
        // 普通组件只使用自己的方法
        allMethods = cls.Methods
    }

    // 排序方法，输出稳定
    methodNames := make([]string, 0, len(allMethods))
    for m := range allMethods {
        methodNames = append(methodNames, m)
    }
    sort.Strings(methodNames)

    // 为每个方法生成链式配置器
    for _, m := range methodNames {
        key := allMethods[m]
        methodGo := toPascal(m)
        sb.WriteString("\nfunc (a *")
        sb.WriteString(structName)
        sb.WriteString(") ")
        sb.WriteString(methodGo)
        sb.WriteString("(value interface{}) *")
        sb.WriteString(structName)
        sb.WriteString(" {\n\ta.Set(\"")
        sb.WriteString(key)
        sb.WriteString("\", value)\n\treturn a\n}\n")
    }

    return os.WriteFile(filepath.Join(outDir, fileName), []byte(sb.String()), 0644)
}

// writeAmisFile 生成 amis.go 便捷函数文件
func writeAmisFile(outDir string, classes map[string]*PhpClass) error {
    var sb strings.Builder

    // 包声明和导入
    sb.WriteString("package gamis\n\n")
    sb.WriteString("import \"github.com/wcz0/gamis/renderers\"\n\n")

    // 收集所有类名并排序
    var classNames []string
    for name := range classes {
        if name != "BaseRenderer" {
            classNames = append(classNames, name)
        }
    }
    sort.Strings(classNames)

    // 为每个类生成便捷函数
    for _, name := range classNames {
        if name == "Component" {
            // Component 的特殊便捷函数，可以接受 type 参数
            sb.WriteString("func ")
            sb.WriteString(name)
            sb.WriteString("(typeStr string) *renderers.")
            sb.WriteString(name)
            sb.WriteString(" {\n\treturn renderers.New")
            sb.WriteString(name)
            sb.WriteString("(typeStr)\n}\n\n")
        } else {
            // 普通组件的便捷函数
            sb.WriteString("func ")
            sb.WriteString(name)
            sb.WriteString("() *renderers.")
            sb.WriteString(name)
            sb.WriteString(" {\n\treturn renderers.New")
            sb.WriteString(name)
            sb.WriteString("()\n}\n\n")
        }
    }

    return os.WriteFile(filepath.Join(outDir, "amis.go"), []byte(sb.String()), 0644)
}

// toPascal 将驼峰/下划线方法名转为 Go 导出方法名
func toPascal(s string) string {
    if s == "" {
        return s
    }
    // 如果本身包含下划线，按下划线切
    if strings.Contains(s, "_") {
        parts := strings.Split(s, "_")
        for i := range parts {
            if parts[i] == "" {
                continue
            }
            parts[i] = strings.ToUpper(parts[i][:1]) + parts[i][1:]
        }
        return strings.Join(parts, "")
    }
    // 对已有的驼峰直接首字母大写
    return strings.ToUpper(s[:1]) + s[1:]
}
