<template>
	<div>
		<p>{{ this.$data.dataEditor.Name }}</p>
		<div class="main-page">
			<div class="sub-page">
				<p v-html="this.$data.dataEditor.Form"></p>
			</div>
		</div>
	</div>
</template>
<script>
export default {
	name: "Editor",
	data() {
		return {
			dataEditor: {},
		};
	},
	mounted() {
		this.getEditor();
	},
	methods: {
		async getEditor() {
			console.log(this.$route.params.id);
			await fetch(`http://localhost:8000/api/get/form/${this.$route.params.id}`)
				.then((res) => res.json())
				.then((data) => (this.$data.dataEditor = data));
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
	box-shadow: 0 0 0.5cm rgba(0, 0, 0, 0.5);
}
.sub-page {
	padding: 1cm;
	height: 297mm;
}
</style>
