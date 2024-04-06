package com.wcz0.renderers;

/**
 * GridNav
 * @author wcz0
 */
public class GridNav extends BaseRenderer {

	public GridNav() {
		this.set("type", "grid-nav");
	}


	/**
	 * 是否显示列表项边框
	 */
	public GridNav border(boolean value) {
		return (GridNav) this.set("border", value);
	}

	/**
	 * 是否将列表项内容居中显示
	 */
	public GridNav center(boolean value) {
		return (GridNav) this.set("center", value);
	}

	/**
	 * 外层 CSS 类名
	 */
	public GridNav className(Object value) {
		return (GridNav) this.set("className", value);
	}

	/**
	 * 列数
	 */
	public GridNav columnNum(Object value) {
		return (GridNav) this.set("columnNum", value);
	}

	/**
	 * 列表项内容排列的方向，可选值为 horizontal 、vertical
	 */
	public GridNav direction(Object value) {
		return (GridNav) this.set("direction", value);
	}

	/**
	 * 列表项之间的间距，默认单位为px
	 */
	public GridNav gutter(Object value) {
		return (GridNav) this.set("gutter", value);
	}

	/**
	 * 图标宽度占比，单位%
	 */
	public GridNav iconRatio(Object value) {
		return (GridNav) this.set("iconRatio", value);
	}

	/**
	 * 列表项 css 类名
	 */
	public GridNav itemClassName(Object value) {
		return (GridNav) this.set("itemClassName", value);
	}

	/**
	 * 列表项图标
	 */
	public GridNav options(Object value) {
		return (GridNav) this.set("options", value);
	}

	/**
	 * 是否调换图标和文本的位置
	 */
	public GridNav reverse(boolean value) {
		return (GridNav) this.set("reverse", value);
	}

	/**
	 * 数据源
	 */
	public GridNav source(Object value) {
		return (GridNav) this.set("source", value);
	}

	/**
	 * 是否将列表项固定为正方形
	 */
	public GridNav square(boolean value) {
		return (GridNav) this.set("square", value);
	}

	/**
	 * 指定为 grid-nav 渲染器。
	 */
	public GridNav type(Object value) {
		return (GridNav) this.set("type", value);
	}

	/**
	 * 图片数组
	 */
	public GridNav value(Object value) {
		return (GridNav) this.set("value", value);
	}

}