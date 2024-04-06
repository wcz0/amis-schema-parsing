package com.wcz0.renderers;

/**
 * SvgIcon
 * @author wcz0
 */
public class SvgIcon extends BaseRenderer {

	public SvgIcon() {
		this.set("type", "svg-icon");
	}


	/**
	 * 外层类名
	 */
	public SvgIcon className(Object value) {
		return (SvgIcon) this.set("className", value);
	}

	/**
	 * 图标
	 */
	public SvgIcon icon(Object value) {
		return (SvgIcon) this.set("icon", value);
	}

	/**
	 * 指定为 svg-icon 渲染器。
	 */
	public SvgIcon type(Object value) {
		return (SvgIcon) this.set("type", value);
	}
}