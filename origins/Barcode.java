package com.wcz0.renderers;

/**
 * Barcode
 * @author wcz0
 */
class Barcode extends BaseRenderer {

	public Barcode() {
		this.set("type", "barcode");
	}

	/**
	 * 外层类名
	 */
	public Barcode className(Object value) {
		return (Barcode) this.set("className", value);
	}

	/**
	 * 指定为 barcode 渲染器。
	 */
	public Barcode type(Object value) {
		return (Barcode) this.set("type", value);
	}
}
