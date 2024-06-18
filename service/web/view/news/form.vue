<template>
  <div>
    <el-drawer v-model="visible" :title="title" size="40%" @opened="opened"
      ><el-divider style="margin: 0px" /><el-container>
        <el-main>
          <el-form
            :model="form"
            label-position="right"
            label-width="auto"
            style="max-width: 90%"
            :disabled="disabled"
          >
			
            <el-form-item label="标题">
				
			  	<el-input v-model="form.title" placeholder="标题"></el-input>
				
            </el-form-item>
			
            <el-form-item label="作者">
				
			  	<el-input v-model="form.author" placeholder="作者"></el-input>
				
            </el-form-item>
			
            <el-form-item label="内容">
				
				<el-input  :autosize="{ minRows: 4, maxRows: 8 }" v-model="form.content" type="textarea" />
				
            </el-form-item>
			
		</el-form></el-main>
        <el-footer v-show="mode !== 'view'">
          <div style="float: right">
            <el-button @click="resetForm">重置</el-button>
            <el-button type="primary" @click="confirm">确认</el-button>
          </div>
        </el-footer>
      </el-container>
    </el-drawer>
  </div>
</template>
<script>
export default {
  data() {
    return {
      visible: false,
      mode: "add",
      title: "添加",
      form: {
		title:"",author:"",content:"",
      },
      resetFormData: {},
      disabled: true,
    };
  },
  methods: {
    resetForm() {
      this.form = this.$TOOL.objCopy(this.resetFormData);
    },
    opened() {
      this.resetFormData = this.$TOOL.objCopy(this.form);
    },
    open(
      mode = "add",
      data = {
title:"",author:"",content:"",
		}
    ) {
      this.doTitle(mode);
      this.doDisabled(mode);
      this.doMode(mode);
      this.form = this.$TOOL.objCopy(data);
      this.visible = true;
    },
    doDisabled(mode) {
      switch (mode) {
        case "add":
        case "edit":
          this.disabled = false;
          break;
        case "view":
          this.disabled = true;
          break;
        default:
          break;
      }
    },
    confirm() {
      this.$emit("confirm", this.form, this.visible);
    },
    doTitle(type) {
      switch (type) {
        case "add":
          this.title = "添加信息";
          break;
        case "edit":
          this.title = "编辑信息";
          break;
        case "view":
          this.title = "查看信息";
          break;
        default:
          break;
      }
    },
    doMode(mode) {
      this.mode = mode;
    },
    close() {
      this.visible = false;
    },
  },
};
</script>