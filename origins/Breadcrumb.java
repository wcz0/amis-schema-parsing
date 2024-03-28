package com.wcz0.renderers;

/**
 * Breadcrumb
 * @author wcz0
 */
public class Breadcrumb extends BaseRenderer {

    Breadcrumb(){
        this.set("type", "breadcrumb");
    }

    /**
     * 外层类名
     */
    public Breadcrumb className(Object value) {
        return (Breadcrumb) this.set("className", value);
    }

    /**
     * 下拉菜单类名
     */
    public Breadcrumb dropdownClassName(Object value) {
        return (Breadcrumb) this.set("dropdownClassName", value);
    }

    /**
     * 下拉菜单项类名
     */
    public Breadcrumb dropdownItemClassName(Object value) {
        return (Breadcrumb) this.set("dropdownItemClassName", value);
    }

    /**
     * 导航项类名
     */
    public Breadcrumb itemClassName(Object value) {
        return (Breadcrumb) this.set("itemClassName", value);
    }

    /**
     * 文本
     */
    public Breadcrumb items(Object value) {
        return (Breadcrumb) this.set("items", value);
    }

    /**
     * 最大展示长度
     */
    public Breadcrumb labelMaxLength(Object value) {
        return (Breadcrumb) this.set("labelMaxLength", value);
    }

    /**
     * 分隔符
     */
    public Breadcrumb separator(Object value) {
        return (Breadcrumb) this.set("separator", value);
    }

    /**
     * 分割符类名
     */
    public Breadcrumb separatorClassName(Object value) {
        return (Breadcrumb) this.set("separatorClassName", value);
    }

    /**
     * 动态数据
     */
    public Breadcrumb source(Object value) {
        return (Breadcrumb) this.set("source", value);
    }

    /**
     * 浮窗提示位置
     */
    public Breadcrumb tooltipPosition(Object value) {
        return (Breadcrumb) this.set("tooltipPosition", value);
    }

    /**
     * 指定为 breadcrumb 渲染器。
     */
    public Breadcrumb type(Object value) {
        return (Breadcrumb) this.set("type", value);
    }

}

