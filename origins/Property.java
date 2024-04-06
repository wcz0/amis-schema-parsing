package com.wcz0.renderers;

/**
 * Property
 * @author wcz0
 */
public class Property extends BaseRenderer {

	public Property() {
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