package com.wcz0.renderers;

/**
 * Flex
 * @author wcz0
 */
public class Flex extends BaseRenderer {

	Flex() {
		this.set("type", "flex");
	}

	/**
	 * stretch, start, flex-start, flex-end, end, center, baseline
	 */
	public Flex alignItems(Object value) {
		return (Flex) this.set("alignItems", value);
	}

	/**
	 * css 类名
	 */
	public Flex className(Object value) {
		return (Flex) this.set("className", value);
	}

	/**
	 * 布局方向: column, row
	 */
	public Flex direction(Object value) {
		return (Flex) this.set("direction", value);
	}

	/**
	 * 子元素, 数组表示多个
	 */
	public Flex items(Object value) {
		return (Flex) this.set("items", value);
	}

	/**
	 * start, flex-start, center, end, flex-end, space-around, space-between, space-evenly
	 */
	public Flex justify(Object value) {
		return (Flex) this.set("justify", value);
	}

	/**
	 * 自定义样式
	 */
	public Flex style(Object value) {
		return (Flex) this.set("style", value);
	}

	/**
	 * 指定为 flex 渲染器。
	 */
	public Flex type(Object value) {
		return (Flex) this.set("type", value);
	}
}