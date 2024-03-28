package com.wcz0.renderers;

/**
 * ChartRadios
 * @author wcz0
 */
public class ChartRadios extends BaseRenderer {

	ChartRadios() {
		this.set("type", "chart-radios");
	}

	/**
	 * 图表数值字段名
	 */
	public ChartRadios chartValueField(Object value) {
		return (ChartRadios) this.set("chartValueField", value);
	}

	/**
	 * 图表配置
	 */
	public ChartRadios config(Object value) {
		return (ChartRadios) this.set("config", value);
	}

	/**
	 * 高亮的时候是否显示 tooltip
	 */
	public ChartRadios showTooltipOnHighlight(boolean value) {
		return (ChartRadios) this.set("showTooltipOnHighlight", value);
	}

	/**
	 * 指定为 chart-radios 渲染器。
	 */
	public ChartRadios type(Object value) {
		return (ChartRadios) this.set("type", value);
	}
}