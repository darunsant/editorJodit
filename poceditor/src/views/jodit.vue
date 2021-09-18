<template>
	<div class="main-page">
		<input v-model="editorname" placeholder="Name" />
		<jodit-editor v-model="content" />
		<button @click="saveJodit">Save</button>
	</div>
</template>

<script>
import "jodit/build/jodit.min.css";
import { JoditEditor } from "jodit-vue";
export default {
	name: "app",

	components: { JoditEditor },
	watch: {
		content() {
			// console.log('content',this.content);
			if (this.content.includes(`<table style="width: 100%;">`)) {
				this.content = this.content.replaceAll(
					`<table style="width: 100%;">`,
					`<table border="1" style="width: 100%;border-collapse:collapse;">`
				);
			}
		},
	},
	data() {
		return {
			content: "",
			editorname: "",
		};
	},
	methods: {
		async saveJodit() {
            const objData = {
                Name:this.$data.editorname,
                Form:this.$data.content
            }
			const response = await fetch(`http://localhost:8000/api/add/form`, {
				method: "POST",
				headers: { "Content-Type": "application/json" },
				body: JSON.stringify(objData),
			});
			console.log("save", response);

		},
	},
};
</script>

<style lang="css">
.main-page {
	width: 210mm;
	min-height: 297mm;
	margin: 10mm auto;
	background: white;
}
</style>