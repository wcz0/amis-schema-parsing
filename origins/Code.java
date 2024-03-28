package com.wcz0.renderers;

/**
 * Code
 * @author wcz0
 */
public class Code extends BaseRenderer{

	Code() {
		this.set("type", "code");
	}

	/**
	 * 外层 CSS 类名
	 */
	public Code className(Object value) {
		return (Code) this.set("className", value);
	}

	/**
	 * 主题，还有 'vs-dark'
	 */
	public Code editorTheme(Object value) {
		return (Code) this.set("editorTheme", value);
	}

	/**
	 * 所使用的高亮语言，默认是 plaintext
	 */
	public Code language(Object value) {
		return (Code) this.set("language", value);
	}

	/**
	 * 在其他组件中，时，用作变量映射
	 */
	public Code name(Object value) {
		return (Code) this.set("name", value);
	}

	/**
	 * 默认 tab 大小
	 */
	public Code tabSize(Object value) {
		return (Code) this.set("tabSize", value);
	}

	/**
	 * 指定为 code 渲染器。
	 */
	public Code type(Object value) {
		return (Code) this.set("type", value);
	}

	/**
	 * 显示的颜色值
	 */
	public Code value(Object value) {
		return (Code) this.set("value", value);
	}

	/**
	 * 是否折行
	 */
	public Code wordWrap(boolean value) {
		return (Code) this.set("wordWrap", value);
	}
}