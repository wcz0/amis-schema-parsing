// package renderers

// type WebComponent struct {
// 	*BaseRenderer
// }

// func NewWebComponent() *WebComponent {
// 	w := &WebComponent{
// 		BaseRenderer: NewBaseRenderer(),
// 	}
// 	w.Set("type", "web_component")
// 	return w
// }

// func (w *WebComponent) body(value interface{}) *WebComponent {
// 	w.Set("body", value)
// 	return w
// }

// func (w *WebComponent) Props(value interface{}) *WebComponent {
// 	w.Set("props", value)
// 	return w
// }

// func (w *WebComponent) Tag(value interface{}) *WebComponent {
// 	w.Set("tag", value)
// 	return w
// }

// func (w *WebComponent) Type(value interface{}) *WebComponent {
// 	w.Set("type", value)
// 	return w
// }

package com.wcz0.renderers;

/**
 * WebComponent
 * @author wcz0
 */
public class WebComponent extends BaseRenderer {

	WebComponent(){
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