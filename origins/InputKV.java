package	com.wcz0.renderers;

/**
 * InputKV
 * @author wcz0
 */
public class InputKV extends BaseRenderer {

	public InputKV() {
		this.set("type", "input-kv");
	}


	/**
	 * 数据录入配置，自动填充或者参照录入
	 */
	public InputKV autoFill(Object value) {
		return (InputKV) this.set("autoFill", value);
	}

	/**
	 * 表单最外层类名
	 */
	public InputKV className(Object value) {
		return (InputKV) this.set("className", value);
	}

	/**
	 * 默认值
	 */
	public InputKV defaultValue(Object value) {
		return (InputKV) this.set("defaultValue", value);
	}

	/**
	 * 表单项描述
	 */
	public InputKV description(Object value) {
		return (InputKV) this.set("description", value);
	}

	/**
	 * 是否禁用
	 */
	public InputKV disabled(boolean value) {
		return (InputKV) this.set("disabled", value);
	}

	/**
	 * 当前表单项是否禁用的条件
	 */
	public InputKV disabledOn(Object value) {
		return (InputKV) this.set("disabledOn", value);
	}

	/**
	 * 是否可以拖拽排序
	 */
	public InputKV draggable(boolean value) {
		return (InputKV) this.set("draggable", value);
	}

	/**
	 * 是否内联
	 */
	public InputKV inline(boolean value) {
		return (InputKV) this.set("inline", value);
	}

	/**
	 * 表单控制器类名
	 */
	public InputKV inputClassName(Object value) {
		return (InputKV) this.set("inputClassName", value);
	}

	/**
	 * key 的提示信息的
	 */
	public InputKV keyPlaceholder(Object value) {
		return (InputKV) this.set("keyPlaceholder", value);
	}

	/**
	 * 表单项标签
	 */
	public InputKV label(Object value) {
		return (InputKV) this.set("label", value);
	}

	/**
	 * 表单项标签对齐方式，默认右对齐，仅在 mode为horizontal 时生效
	 */
	public InputKV labelAlign(Object value) {
		return (InputKV) this.set("labelAlign", value);
	}

	/**
	 * label的类名
	 */
	public InputKV labelClassName(Object value) {
		return (InputKV) this.set("labelClassName", value);
	}

	/**
	 * 表单项标签描述
	 */
	public InputKV labelRemark(Object value) {
		return (InputKV) this.set("labelRemark", value);
	}

	/**
	 * 字段名，指定该表单项提交时的 key
	 */
	public InputKV name(Object value) {
		return (InputKV) this.set("name", value);
	}

	/**
	 * 表单项描述
	 */
	public InputKV placeholder(Object value) {
		return (InputKV) this.set("placeholder", value);
	}

	/**
	 * 是否必填
	 */
	public InputKV required(boolean value) {
		return (InputKV) this.set("required", value);
	}

	/**
	 * 通过表达式来配置当前表单项是否为必填。
	 */
	public InputKV requiredOn(Object value) {
		return (InputKV) this.set("requiredOn", value);
	}

	/**
	 * 是否该表单项值发生变化时就提交当前表单。
	 */
	public InputKV submitOnChange(boolean value) {
		return (InputKV) this.set("submitOnChange", value);
	}

	/**
	 * 指定为 input-kv 渲染器。
	 */
	public InputKV type(Object value) {
		return (InputKV) this.set("type", value);
	}

	/**
	 * 表单校验接口
	 */
	public InputKV validateApi(Object value) {
		return (InputKV) this.set("validateApi", value);
	}

	/**
	 * 表单项值格式验证，支持设置多个，多个规则用英文逗号隔开。
	 */
	public InputKV validations(Object value) {
		return (InputKV) this.set("validations", value);
	}

	/**
	 * 表单默认值
	 */
	public InputKV value(Object value) {
		return (InputKV) this.set("value", value);
	}

	/**
	 * value 的提示信息的
	 */
	public InputKV valuePlaceholder(Object value) {
		return (InputKV) this.set("valuePlaceholder", value);
	}

	/**
	 * 值类型
	 */
	public InputKV valueType(Object value) {
		return (InputKV) this.set("valueType", value);
	}

	/**
	 * 是否可见
	 */
	public InputKV visible(boolean value) {
		return (InputKV) this.set("visible", value);
	}

	/**
	 * 当前表单项是否禁用的条件
	 */
	public InputKV visibleOn(Object value) {
		return (InputKV) this.set("visibleOn", value);
	}
}

