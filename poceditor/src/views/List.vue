<template>
	<div>
		<div v-for="(item,index) in this.$data.list" :key="index">
			<h4 class="editorList" @click="changePage(item.id)"><b>{{item.Name}}</b></h4>
		</div>
	</div>
</template>
<script>
export default {
	name: "List",
	data() {
		return {
			list: [],
		};
	},
	mounted() {
		this.getList();
	},
	methods: {
		async getList() {
			await fetch(`http://localhost:8000/api/get/form`)
				.then((res) => res.json())
				.then((data) => (this.$data.list = data));
		},
		changePage(value){
			this.$router.push({name:'list',params:{id:value}})
		}
	},
};
</script>

<style lang="css">
.editorList{
	cursor: pointer;
}
</style>
