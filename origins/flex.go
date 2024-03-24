package renderers

type Flex struct {
	*BaseRenderer
}

func NewFlex() *Flex {
	f := &Flex{
		BaseRenderer: NewBaseRenderer(),
	}
	f.Set("type", "flex")
	return f
}

/**
 * stretch, start, flex-start, flex-end, end, center, baseline
 */
func (f *Flex) AlignItems(value interface{}) *Flex {
	f.Set("alignItems", value)
	return f
}

/**
 * css 类名
 */
func (f *Flex) ClassName(value interface{}) *Flex {
	f.Set("className", value)
	return f
}

/**
 * 布局方向: column, row
 */
func (f *Flex) Direction(value interface{}) *Flex {
	f.Set("direction", value)
	return f
}

/**
 * 子元素, 数组表示多个
 */
func (f *Flex) Items(value interface{}) *Flex {
	f.Set("items", value)
	return f
}

/**
 * start, flex-start, center, end, flex-end, space-around, space-between, space-evenly
 */
func (f *Flex) Justify(value interface{}) *Flex {
	f.Set("justify", value)
	return f
}

/**
 * 自定义样式
 */
func (f *Flex) Style(value interface{}) *Flex {
	f.Set("style", value)
	return f
}

/**
 * 指定为 flex 渲染器。
 */
func (f *Flex) Type(value interface{}) *Flex {
	f.Set("type", value)
	return f
}
