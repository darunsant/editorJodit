<template>
	<vue-editor v-model="content" :editor-toolbar="customToolbar" />
</template>

<script>
// @ is an alias to /src
import { VueEditor } from "vue2-editor";

export default {
	name: "Home",
	components: { VueEditor },

	data: () => ({
		content: "<h1>Some initial content</h1>",
		customToolbar: [
			["bold", "italic", "underline", "strike"],
			[{ list: "ordered" }, { list: "bullet" }],
			["image", "code-block",'blockquote',],
			[{ direction: "rtl" }],
			[{ size: ["small", false, "large", "huge"] }],
			[{ header: [1, 2, 3, 4, 5, 6, false] }],
			[{ color: [] }, { background: [] }],
			[{ font: [] }],
			[{ align: [] }],
		],
	}),
};
</script>
