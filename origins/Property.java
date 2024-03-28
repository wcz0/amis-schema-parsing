// package renderers

// type Property struct {
// 	*BaseRenderer
// }

// func NewProperty() *Property {
// 	p := &Property{
// 		BaseRenderer: NewBaseRenderer(),
// 	}
// 	p.Set("type", "property")
// 	return p
// }

// /**
//  * 外层 dom 的类名
//  */
// func (p *Property) ClassName(value interface{}) *Property {
// 	p.Set("className", value)
// 	return p
// }

// /**
//  * 每行几列
//  */
// func (p *Property) Column(value interface{}) *Property {
// 	p.Set("column", value)
// 	return p
// }

// /**
//  * 属性值的样式
//  */
// func (p *Property) ContentStyle(value interface{}) *Property {
// 	p.Set("contentStyle", value)
// 	return p
// }

// /**
//  *
//  */
// func (p *Property) Items(value interface{}) *Property {
// 	p.Set("items", value)
// 	return p
// }

// /**
//  * 属性名的样式
//  */
// func (p *Property) LabelStyle(value interface{}) *Property {
// 	p.Set("labelStyle", value)
// 	return p
// }

// /**
//  * 显示模式，目前只有 'table' 和 'simple'
//  */
// func (p *Property) Mode(value interface{}) *Property {
// 	p.Set("mode", value)
// 	return p
// }

// /**
//  * 模式下属性名和值之间的分隔符
//  */
// func (p *Property) Separator(value interface{}) *Property {
// 	p.Set("separator", value)
// 	return p
// }

// /**
//  * 数据源
//  */
// func (p *Property) Source(value interface{}) *Property {
// 	p.Set("source", value)
// 	return p
// }

// /**
//  * 外层 dom 的样式
//  */
// func (p *Property) Style(value interface{}) *Property {
// 	p.Set("style", value)
// 	return p
// }

// /**
//  * 标题
//  */
// func (p *Property) Title(value interface{}) *Property {
// 	p.Set("title", value)
// 	return p
// }

// /**
//  * 指定为 property 渲染器。
//  */
// func (p *Property) Type(value interface{}) *Property {
// 	p.Set("type", value)
// 	return p
// }

package com.wcz0.renderers;

/**
 * Property
 * @author wcz0
 */
public class Property extends BaseRenderer {

	Property() {
		this.set("type", "property");
	}

	/**
	 * 外层 dom 的类名
	 */
	public Property className(Object value) {
		return (Property) this.set("className", value);
	}

	/**
	 * 每行几列
	 */
	public Property column(Object value) {
		return (Property) this.set("column", value);
	}

	/**
	 * 属性值的样式
	 */
	public Property contentStyle(Object value) {
		return (Property) this.set("contentStyle", value);
	}

	/**
	 *
	 */
	public Property items(Object value) {
		return (Property) this.set("items", value);
	}

	/**
	 * 属性名的样式
	 */
	public Property labelStyle(Object value) {
		return (Property) this.set("labelStyle", value);
	}

	/**
	 * 显示模式，目前只有 'table' 和 'simple'
	 */
	public Property mode(Object value) {
		return (Property) this.set("mode", value);
	}

	/**
	 * 模式下属性名和值之间的分隔符
	 */
	public Property separator(Object value) {
		return (Property) this.set("separator", value);
	}

	/**
	 * 数据源
	 */
	public Property source(Object value) {
		return (Property) this.set("source", value);
	}

	/**
	 * 外层 dom 的样式
	 */
	public Property style(Object value) {
		return (Property) this.set("style", value);
	}

	/**
	 * 标题
	 */
	public Property title(Object value) {
		return (Property) this.set("title", value);
	}

	/**
	 * 指定为 property 渲染器。
	 */
	public Property type(Object value) {
		return (Property) this.set("type", value);
	}
}