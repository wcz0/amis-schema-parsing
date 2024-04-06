package com.wcz0.renderers;

/**
 * Toast
 * @author wcz0
 */
public class Toast extends BaseRenderer {

	public Toast() {
		this.set("type", "toast");
	}


	/**
	 * 内容
	 */
	public Toast body(Object value) {
		return (Toast) this.set("body", value);
	}

	/**
	 * 是否显示关闭按钮
	 */
	public Toast closeButton(boolean value) {
		return (Toast) this.set("closeButton", value);
	}

	/**
	 * 轻提示内容
	 */
	public Toast items(Object value) {
		return (Toast) this.set("items", value);
	}

	/**
	 * 展示图标，可选'info'/'success'/'error'/'warning'
	 */
	public Toast level(Object value) {
		return (Toast) this.set("level", value);
	}

	/**
	 * 提示显示位置，可选值: top-right | top-center | top-left | bottom-center | bottom-left | bottom-right | center
	 */
	public Toast position(Object value) {
		return (Toast) this.set("position", value);
	}

	/**
	 * 是否显示图标
	 */
	public Toast showIcon(boolean value) {
		return (Toast) this.set("showIcon", value);
	}

	/**
	 * 持续时间
	 */
	public Toast timeout(Object value) {
		return (Toast) this.set("timeout", value);
	}

	/**
	 * 标题
	 */
	public Toast title(Object value) {
		return (Toast) this.set("title", value);
	}
}