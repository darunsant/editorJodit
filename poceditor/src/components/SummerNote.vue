<template>
    <div :id="element">
        <slot></slot>
    </div>
</template>

<script>
    const $ = require("jquery")
    import "summernote/dist/summernote-lite.min.css";
    import "summernote/dist/summernote-lite.min";
    export default {
        name: "Summernote",
        data() {
            return {
                element: `summernote_${new Date().getTime()}`
            };
        },
        props: {
            config: {
                type: Object
            },
            ready: {
                type: Function
            }
        },
        mounted() {
            this.$nextTick(() => {
                let defaultConfig = {
                    //lang: "es-ES",
                    placeholder: "placeholder...",
                    height: 300,
                    focus: true,
                    callbacks: {
                        onInit: () => {
                            this.$emit("ready", this);
                        },
                        onChange: contents => {
                            this.$emit("change", contents);
                        },
                        onEnter: () => {
                            this.$emit("enter");
                        },
                        onFocus: () => {
                            this.$emit("focus");
                        },
                        onBlur: () => {
                            this.$emit("blur");
                        },
                        onKeyup: e => {
                            this.$emit("keyup", e);
                        },
                        onKeydown: e => {
                            this.$emit("keydown", e);
                        },
                        onPaste: e => {
                            this.$emit("paste", e);
                        },
                    }
                };
                let config = Object.assign({}, defaultConfig, this.config);
                console.log('config',this.config);
                $("#" + this.element).summernote(config);
            });
        },
        methods: {
            summernote(...args) {
                $("#" + this.element).summernote(...args);
            },
            ui() {
                return $.summernote.ui;
            }
        }
    };
</script>
<style scoped>
    .summernote-img {
        width: 100%;
    }
</style>