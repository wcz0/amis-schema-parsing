package com.wcz0.renderers;

/**
 * Markdown
 * @author wcz0
 */
public class Markdown extends BaseRenderer {

	public Markdown() {
		this.set("type", "markdown");
	}


	/**
	 * 类名
	 */
	public Markdown className(Object value) {
		return (Markdown) this.set("className", value);
	}

	/**
	 * 名称
	 */
	public Markdown name(Object value) {
		return (Markdown) this.set("name", value);
	}

	/**
	 * 外部地址
	 */
	public Markdown src(Object value) {
		return (Markdown) this.set("src", value);
	}

	/**
	 * 静态值
	 */
	public Markdown value(Object value) {
		return (Markdown) this.set("value", value);
	}
}