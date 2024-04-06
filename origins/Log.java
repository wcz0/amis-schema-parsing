package com.wcz0.renderers;

/**
 * Log
 * @author wcz0
 */
public class Log extends BaseRenderer {

	public Log() {
		this.set("type", "log");
	}


	/**
	 * 是否自动滚动到底部
	 */
	public Log autoScroll(Object value) {
		return (Log) this.set("autoScroll", value);
	}

	/**
	 * 外层 CSS 类名
	 */
	public Log className(Object value) {
		return (Log) this.set("className", value);
	}

	/**
	 * 关闭 ANSI 颜色支持
	 */
	public Log disableColor(Object value) {
		return (Log) this.set("disableColor", value);
	}

	/**
	 * 返回内容的字符编码
	 */
	public Log encoding(Object value) {
		return (Log) this.set("encoding", value);
	}

	/**
	 * 展示区域高度
	 */
	public Log height(Object value) {
		return (Log) this.set("height", value);
	}

	/**
	 * 最大显示行数
	 */
	public Log maxLength(Object value) {
		return (Log) this.set("maxLength", value);
	}

	/**
	 * 可选日志操作：['stop', 'clear', 'showLineNumber', 'filter']
	 */
	public Log operation(Object value) {
		return (Log) this.set("operation", value);
	}

	/**
	 * 加载中的文字
	 */
	public Log placeholder(Object value) {
		return (Log) this.set("placeholder", value);
	}

	/**
	 * 设置每行高度，将会开启虚拟渲染
	 */
	public Log rowHeight(Object value) {
		return (Log) this.set("rowHeight", value);
	}

	/**
	 * 接口
	 */
	public Log source(Object value) {
		return (Log) this.set("source", value);
	}

	/**
	 * 指定为 log 渲染器。
	 */
	public Log type(Object value) {
		return (Log) this.set("type", value);
	}
}