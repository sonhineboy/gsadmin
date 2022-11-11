<template>
	<el-card shadow="hover" header="关于项目" class="item-background">
		<p>GS Admin=gin+scui 它是golang 开发的一个企业级后台。遵循MIT开源协议。前端框架是scui,SCUI基于 Vue3、elementPlus 持续性的提供独家组件和丰富的业务模板帮助你快速搭建企业级中后台前端任务。后端框架是gin,Gin是一个golang的微框架，封装比较优雅，具有快速灵活，容错方便等特点。</p>
		<p>
			<a href='https://gitee.com/kevn/gsadmin' target="_blank">
				<img src='https://gitee.com/kevn/gsadmin/badge/star.svg?theme=dark' alt='star' style="vertical-align: middle">
			</a>
			<a href='https://github.com/sonhineboy/gsadmin' target="_blank">
				<img src='https://img.shields.io/github/stars/sonhineboy/gsadmin.svg?logo=GitHub' alt='star' style="vertical-align: middle">
			</a>
		</p>
	</el-card>
</template>

<script>
	export default {
		title: "关于项目",
		icon: "el-icon-setting",
		description: "点个星星支持一下",
		data() {
			return {

			}
		}
	}
</script>

<style scoped>
	.item-background p {color: #999;margin-top:10px;line-height: 1.8;}
</style>
