package com.wcz0.renderers;

/**
 * TableView
 * @author wcz0
 */
public class TableView extends BaseRenderer {

	public TableView() {
		this.set("type", "table_view");
	}
	

	/**
	 * 指定为 table_view 渲染器。
	 */
	public TableView type(Object value) {
		return (TableView) this.set("type", value);
	}
}