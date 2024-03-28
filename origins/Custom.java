package com.wcz0.renderers;

/**
 * Custom
 * @author wcz0
 */
public class Custom extends BaseRenderer {

    Custom() {
        this.set("type", "custom");
    }

    /**
     * 自定义渲染器
     */
    public Custom custom(Object value) {
        return (Custom) this.set("custom", value);
    }
}