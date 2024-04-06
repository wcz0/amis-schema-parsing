package com.wcz0.renderers;

/**
 * InputKVS
 * @author wcz0
 */
public class InputKVS extends BaseRenderer {

	public InputKVS() {
		this.set("type", "input-kvs");
	}


	/**
	 * 数据录入配置，自动填充或者参照录入
	 */
	public InputKVS autoFill(Object value) {
		return (InputKVS) this.set("autoFill", value);
	}

	/**
	 * 表单最外层类名
	 */
	public InputKVS className(Object value) {
		return (InputKVS) this.set("className", value);
	}

	/**
	 * 表单项描述
	 */
	public InputKVS description(Object value) {
		return (InputKVS) this.set("description", value);
	}

	/**
	 * 是否禁用
	 */
	public InputKVS disabled(boolean value) {
		return (InputKVS) this.set("disabled", value);
	}

	/**
	 * 当前表单项是否禁用的条件
	 */
	public InputKVS disabledOn(Object value) {
		return (InputKVS) this.set("disabledOn", value);
	}

	/**
	 * 是否内联
	 */
	public InputKVS inline(boolean value) {
		return (InputKVS) this.set("inline", value);
	}

	/**
	 * 表单控制器类名
	 */
	public InputKVS inputClassName(Object value) {
		return (InputKVS) this.set("inputClassName", value);
	}

	/**
	 * 表单项标签
	 */
	public InputKVS label(Object value) {
		return (InputKVS) this.set("label", value);
	}

	/**
	 * 表单项标签对齐方式，默认右对齐，仅在 mode为horizontal 时生效
	 */
	public InputKVS labelAlign(Object value) {
		return (InputKVS) this.set("labelAlign", value);
	}

	/**
	 * label 的类名
	 */
	public InputKVS labelClassName(Object value) {
		return (InputKVS) this.set("labelClassName", value);
	}

	/**
	 * 表单项标签描述
	 */
	public InputKVS labelRemark(Object value) {
		return (InputKVS) this.set("labelRemark", value);
	}

	/**
	 * 字段名，指定该表单项
	 * 提交时的 key
	 */
	public InputKVS name(Object value) {
		return (InputKVS) this.set("name", value);
	}

	/**
	 * 表单项描述
	 */
	public InputKVS placeholder(Object value) {
		return (InputKVS) this.set("placeholder", value);
	}

	/**
	 * 是否必填
	 */
	public InputKVS required(boolean value) {
		return (InputKVS) this.set("required", value);
	}

	/**
	 * 通过表达式来配置当前表单项是否为必填。
	 */
	public InputKVS requiredOn(Object value) {
		return (InputKVS) this.set("requiredOn", value);
	}

	/**
	 * 是否该表单项值发生变化时就提交当前表单。
	 */
	public InputKVS submitOnChange(boolean value) {
		return (InputKVS) this.set("submitOnChange", value);
	}

	/**
	 * 指定为 input-kvs 渲染器。
	 */
	public InputKVS type(Object value) {
		return (InputKVS) this.set("type", value);
	}

	/**
	 * 表单校验接口
	 */
	public InputKVS validateApi(Object value) {
		return (InputKVS) this.set("validateApi", value);
	}

	/**
	 * 表单项值格式验证，支持设置多个，多个规则用英文逗号隔开。
	 */
	public InputKVS validations(Object value) {
		return (InputKVS) this.set("validations", value);
	}

	/**
	 * 表单默认值
	 */
	public InputKVS value(Object value) {
		return (InputKVS) this.set("value", value);
	}

	/**
	 * 是否可见
	 */
	public InputKVS visible(boolean value) {
		return (InputKVS) this.set("visible", value);
	}

	/**
	 * 当前表单项是否禁用的条件
	 */
	public InputKVS visibleOn(Object value) {
		return (InputKVS) this.set("visibleOn", value);
	}

}
