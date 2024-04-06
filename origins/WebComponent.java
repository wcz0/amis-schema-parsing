package com.wcz0.renderers;

/**
 * WebComponent
 * @author wcz0
 */
public class WebComponent extends BaseRenderer {

	public WebComponent(){
		this.set("type", "web_component");
	}
	

	/**
	 * 内容
	 */
	public WebComponent body(Object value) {
		return (WebComponent) this.set("body", value);
	}

	/**
	 * 属性
	 */
	public WebComponent props(Object value) {
		return (WebComponent) this.set("props", value);
	}

	/**
	 * 标签
	 */
	public WebComponent tag(Object value) {
		return (WebComponent) this.set("tag", value);
	}

	/**
	 * 指定为 web_component 渲染器。
	 */
	public WebComponent type(Object value) {
		return (WebComponent) this.set("type", value);
	}
}