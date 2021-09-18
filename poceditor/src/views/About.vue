<template>
	<div id="app">
		<div>
			<!-- <textarea class="summernote"></textarea> -->
			
			<summer-note
				:config="summernoteConfig"
				@ready="readySummernote"
			></summer-note>
		</div>
		<div class="preview">
			<div v-html="contents"></div>
		</div>
	</div>
</template>
<script>
import SummerNote from "@/components/SummerNote";

export default {
	name: "App",
	components: {
		SummerNote,
	},
	data() {
		return {
			summernoteConfig: {
				placeholder: "",
				tabsize: 2,
				height: 200,

				toolbar: [
					["font", ["bold", "underline", "clear"]],
					["font", ["superscript", "subscript"]],
					["para", ["ul", "ol", "paragraph"]],
					["table", ["table"]],
					["insert", ["picture", "hr", "grid"]],
					["view", ["codeview"]],
				],
			},
			contents: "",
		};
	},
	methods: {
		readySummernote(editor) {
			editor.summernote("code", this.contents);
			editor.$on("change", (contents) => {
				console.log("contents", contents);
				this.contents = contents;
			});
		},
	},
};
</script>

<style>
#app {
	font-family: "Avenir", Helvetica, Arial, sans-serif;
	-webkit-font-smoothing: antialiased;
	-moz-osx-font-smoothing: grayscale;
	text-align: center;
	color: #2c3e50;
	margin-top: 60px;
}
</style>
