// package renderers

// type Toast struct {
// 	*BaseRenderer
// }

// func NewToast() *Toast {
// 	t := &Toast{
// 		BaseRenderer: NewBaseRenderer(),
// 	}
// 	t.Set("type", "toast")
// 	return t
// }

// /**
//  * 内容
//  */
// func (t *Toast) Body(value interface{}) *Toast {
// 	t.Set("body", value)
// 	return t
// }

// /**
//  * 是否显示关闭按钮
//  */
// func (t *Toast) CloseButton(value bool) *Toast {
// 	t.Set("closeButton", value)
// 	return t
// }

// /**
//  * 轻提示内容
//  */
// func (t *Toast) Items(value interface{}) *Toast {
// 	t.Set("items", value)
// 	return t
// }

// /**
//  * 展示图标，可选'info'/'success'/'error'/'warning'
//  */
// func (t *Toast) Level(value interface{}) *Toast {
// 	t.Set("level", value)
// 	return t
// }

// /**
//  * 提示显示位置，可选值: top-right | top-center | top-left | bottom-center | bottom-left | bottom-right | center
//  */
// func (t *Toast) Position(value interface{}) *Toast {
// 	t.Set("position", value)
// 	return t
// }

// /**
//  * 是否显示图标
//  */
// func (t *Toast) ShowIcon(value bool) *Toast {
// 	t.Set("showIcon", value)
// 	return t
// }

// /**
//  * 持续时间
//  */
// func (t *Toast) Timeout(value interface{}) *Toast {
// 	t.Set("timeout", value)
// 	return t
// }

// /**
//  * 标题
//  */
// func (t *Toast) Title(value interface{}) *Toast {
// 	t.Set("title", value)
// 	return t
// }

package com.wcz0.renderers;

/**
 * Toast
 * @author wcz0
 */
public class Toast extends BaseRenderer {

	Toast() {
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