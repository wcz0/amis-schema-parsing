// package renderers

// type TableView struct {
// 	*BaseRenderer
// }

// func NewTableView() *TableView {
// 	t := &TableView{
// 		BaseRenderer: NewBaseRenderer(),
// 	}
// 	t.Set("type", "table_view")
// 	return t
// }

// func (t *TableView) Type(value interface{}) *TableView {
// 	t.Set("type", value)
// 	return t
// }

package com.wcz0.renderers;

/**
 * TableView
 * @author wcz0
 */
public class TableView extends BaseRenderer {

	TableView() {
		this.set("type", "table_view");
	}

	/**
	 * 指定为 table_view 渲染器。
	 */
	public TableView type(Object value) {
		return (TableView) this.set("type", value);
	}
}