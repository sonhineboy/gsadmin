<template>
  <el-container
    ><el-main class="noPadding">
      <el-card shadow="never">
        <template #header> 代码生成器 </template>
        <el-form :model="form" label-width="auto" style="max-width: 80%">
          <el-row>
            <el-form-item label="表名">
              <el-col :span="6">
                <el-input
                  v-model="form.tableDiyName"
                  placeholder="请填写表名"
                  style="max-width: 95%"
                ></el-input>
              </el-col>

              <el-col :span="10">
                <el-select v-model="form.tableName" placeholder="已有数据表">
                </el-select>
              </el-col>
            </el-form-item>
          </el-row>
          <el-row>
            <el-col :span="24">
              <el-checkbox-group v-model="form.checkbox">
                <el-checkbox label="自动创建表" value="a"></el-checkbox>
                <el-checkbox label="b" value="b"></el-checkbox>
              </el-checkbox-group>
            </el-col>
          </el-row>
        </el-form>
        <el-divider />
        <el-button type="primary" @click="openNewDrawer()" icon="ElIconPlus" />
        <el-table :data="tableData" style="width: 100%">
          <el-table-column prop="name" label="字段名"></el-table-column>
          <el-table-column prop="transfor" label="翻译"></el-table-column>
          <el-table-column prop="type" label="类型"></el-table-column>
          <el-table-column prop="isNull" label="是否空值">
            <template #default="scope">
              <el-switch v-model="scope.row.isNull" disabled />
            </template>
          </el-table-column>
          <el-table-column prop="index" label="索引"></el-table-column>
          <el-table-column prop="json" label="Json"></el-table-column>
          <el-table-column prop="default" label="默认值">
            <template #default="scope">
              <div v-if="scope.row.default">
                {{ scope.row.default }}
              </div>
              <div v-else>-</div>
            </template>
          </el-table-column>
          <el-table-column prop="describe" label="描述">
            <template #default="scope">
              <div v-if="scope.row.describe">
                {{ scope.row.describe }}
              </div>
              <div v-else>-</div>
            </template>
          </el-table-column>
          <el-table-column prop="primapy" label="主键">
            <template #default="scope">
            
            <el-checkbox disabled :model-value="scope.row.primapy"></el-checkbox>
            </template>

          </el-table-column>
          <el-table-column label="操作" fixed="right" width="100">
            <template #default="scope">
              <el-link
                type="primary"
                @click="actions('edit', scope.$index, scope.row)"
                icon="ElIconEdit"
              />
              &nbsp;
              <el-popconfirm
                title="确定要删除吗?"
                @confirm="actions('del', scope.$index, scope.row)"
              >
                <template #reference>
                  <el-link type="primary" icon="ElIconDelete" />
                </template>
              </el-popconfirm>
            </template>
          </el-table-column>
          <template #empty>
            <el-empty description="站暂无数据" />
          </template>
        </el-table>
      </el-card> </el-main
  ></el-container>

  <el-drawer
    v-model="drawer"
    title="详细信息"
    size="40%"
    @closed="drawerClosed()"
  >
    <template #default>
      <el-divider style="margin: 0px" />
      <el-container>
        <el-main>
          <el-form
            :model="formField"
            label-position="left"
            label-width="auto"
            ref="formfieldRef"
            :rules="rules"
          >
            <el-form-item label="字段名" prop="name">
              <el-input v-model="formField.name" placeholder=""></el-input>
            </el-form-item>
            <el-form-item label="主键" prop="primapy">
              <el-checkbox
                v-model="formField.primapy"
                label="主键"
              ></el-checkbox>
            </el-form-item>
            <el-form-item label="Json" prop="json">
              <el-input v-model="formField.json" placeholder=""></el-input>
            </el-form-item>
            <el-form-item label="翻译" prop="transfor">
              <el-input v-model="formField.transfor" placeholder=""></el-input>
            </el-form-item>
            <el-form-item label="类型" prop="type">
              <el-select v-model="formField.type" placeholder="请选择类型">
                <el-option
                  v-for="item in fieldType"
                  :value="item.value"
                  :label="item.label"
                  :key="item.value"
                />
              </el-select>
            </el-form-item>

            <el-form-item label="是否空值" prop="inNul">
              <el-switch v-model="formField.isNull" />
            </el-form-item>

            <el-form-item label="索引" prop="index">
              <el-select v-model="formField.index" placeholder="">
                <el-option
                  v-for="item in filedIndex"
                  :value="item.value"
                  :label="item.label"
                  :key="item.value"
                />
              </el-select>
            </el-form-item>

            <el-form-item label="默认值" prop="default">
              <el-input v-model="formField.default" placeholder="" />
            </el-form-item>
            <el-form-item label="描述" prop="describe">
              <el-input
                type="textarea"
                v-model="formField.describe"
                placeholder=""
              />
            </el-form-item>
          </el-form>
        </el-main>
        <el-footer>
          <div style="float: right">
            <el-button @click="resetForm()">重置</el-button>
            <el-button @click="fieldFormSumit()" type="primary">确认</el-button>
          </div>
        </el-footer>
      </el-container>
    </template>
  </el-drawer>
</template>
<script>
export default {
  name: "gen",
  data() {
    return {
      rules: {
        name: [
          {
            required: true,
            message: this.$TOOL.validate.transform(
              this.$t("validate.required"),
              "字段名"
            ),
            trigger: "blur",
          },
        ],
        json: [
          {
            required: true,
            message: this.$TOOL.validate.transform(
              this.$t("validate.required"),
              "json"
            ),
            trigger: "blur",
          },
        ],
        transfor: [
          {
            required: true,
            message: this.$TOOL.validate.transform(
              this.$t("validate.required"),
              "翻译"
            ),
            trigger: "blur",
          },
        ],
        type: [
          {
            required: true,
            message: this.$TOOL.validate.transform(
              this.$t("validate.required"),
              "类型"
            ),
            trigger: "blur",
          },
        ],
      },
      tableData: [
      ],
      drawer: false,
      drawerType: null,
      form: {
        checkbox: [],
        tableName: "",
        tableDiyName: "",
      },
      formField: {
        index: "",
        type: "",
      },
      formFieldReset: {},
      tableDataEditIndex: null,
      fieldType: [
        { value: "varchar", label: "varchar" },
        { value: "text", label: "text" },
        { value: "timestamp", label: "timestamp" },
        { value: "bigint", label: "bigint" },
        { value: "int", label: "int" },
        { value: "tinyint", label: "tinyint" },
        { value: "float", label: "float" },
        { value: "decimal", label: "decimal" },
      ],
      filedIndex: [
        { value: "Null", labelL: "Null" },
        { value: "UNIQUE", labelL: "UNIQUE" },
        { value: "NORMAL", labelL: "NORMAL" },
        { value: "FULLTEXT", labelL: "FULLTEXT" },
      ],
    };
  },
  methods: {
    actions(type, index, row) {
      const _that = this;
      switch (type) {
        case "del":
          this.del(index);
          _that.$message.success("删除完成");
          break;
        case "edit":
          this.edite(index, row);
          break;
        case "detail":
          this.detail(row);
          break;
      }
      console.log(type, index, row);
    },
    del(index) {
      this.tableData.splice(index, 1);
    },
    edite(index, row) {
      this.drawer = true;
      this.formField = this.$TOOL.objCopy(row);
      this.formFieldReset = this.$TOOL.objCopy(row);
      this.drawerType = "edit";
      this.tableDataEditIndex = index;
    },
    detail(row) {
      console.log(row);
    },
    resetForm() {
      this.formField = this.$TOOL.objCopy(this.formFieldReset);
    },
    drawerClosed() {
      this.formField = {};
      this.formFieldReset = {};
      this.$refs["formfieldRef"].clearValidate();
      this.tableDataEditInde = null;
      this.drawerType = null;
    },
    openNewDrawer() {
      this.drawer = true;
      this.drawerType = "new";
    },
    fieldFormSumit() {
      this.$refs["formfieldRef"].validate((valid) => {
        if (valid) {
          if (this.drawerType == "edit" && this.tableDataEditIndex !== null) {
            this.updateItem(this.formField);
            this.drawer = false;
          }

          if (this.drawerType == "new") {
            this.addItem(this.formField);
            this.drawer = false;
          }
        } else {
          this.$message.error("请完善字段信息！");
        }
      });
    },
    addItem(row) {
      this.tableData.push(row);
    },
    updateItem(row) {
      this.tableData[this.tableDataEditIndex] = this.$TOOL.objCopy(row);
    },
  },
};
</script>
