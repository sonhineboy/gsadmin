<template>
  <el-container>
    <el-main class="nopadding">
      <div class="diy-gen">
        <el-form :model="form" label-width="auto" style="max-width: 80%">
          <el-row>
            <el-form-item label="表名">
              <el-col :span="8">
                <el-input
                  v-model="form.tableDiyName"
                  placeholder="请填写表名"
                  style="max-width: 95%"
                ></el-input
              ></el-col>

              <el-col :span="8">
                <el-select
                  v-model="form.tableName"
                  placeholder="已有数据表"
                  clearable
                  @change="tablesChange"
                >
                  <el-option
                    v-for="item in tableNames"
                    :key="item.value"
                    :label="item.label"
                    :value="item.value"
                  ></el-option>
                </el-select>
              </el-col>
              <el-col :span="8"
                ><div style="padding-left: 8px">
                  <el-button type="success" @click="genFields"
                    >生成字段</el-button
                  >
                </div></el-col
              >
            </el-form-item>
          </el-row>
          <el-row>
            <el-col :span="8">
              <el-form-item label="包名">
                <el-input
                  v-model="form.controllerPackage"
                  placeholder="默认system"
                ></el-input>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="8">
              <el-form-item label="菜单">
                <el-input
                  v-model="form.menuName"
                  placeholder="菜单名称"
                ></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="10">
              <p style="color: red; line-height: 33px; text-indent: 6px">
                <el-icon><el-icon-WarningFilled /></el-icon
                >&nbsp;如果生成菜单，必填项！
              </p>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="24">
              <el-checkbox-group v-model="form.checkbox">
                <el-checkbox label="生成Model" value="genModel"></el-checkbox>
                <el-checkbox
                  label="生成Request"
                  value="genRequest"
                ></el-checkbox>
                <el-checkbox
                  label="生成Controller"
                  value="genController"
                ></el-checkbox>
                <el-checkbox
                  label="生成Repository"
                  value="genRepository"
                ></el-checkbox>
                <el-checkbox
                  label="生成前端模板"
                  value="genRepository"
                ></el-checkbox>
                <el-checkbox label="生成数据库" value="genDb"></el-checkbox>
                <el-checkbox label="生成路由" value="genRouter"></el-checkbox>
                <el-checkbox label="生成菜单" value="genRouter"></el-checkbox>
              </el-checkbox-group>
            </el-col>
          </el-row>
        </el-form>
      </div>
      <div style="height: 10px; background-color: #f6f8f9; width: 100%"></div>
      <el-container>
        <el-header>
          <div class="left-panel">
            <el-button
              type="primary"
              @click="openNewDrawer()"
              icon="ElIconPlus"
            />
          </div>
        </el-header>
        <el-main class="nopadding">
          <el-table
            :data="tableData"
            class="table"
            style="width: 100%"
            row-key="id"
          >
            <el-table-column prop="name" label="字段名"></el-table-column>
            <el-table-column prop="transform" label="翻译"></el-table-column>
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
            <el-table-column prop="primary" label="主键">
              <template #default="scope">
                <el-checkbox
                  disabled
                  :model-value="scope.row.primary"
                ></el-checkbox>
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
        </el-main>
      </el-container>
    </el-main>
    <el-footer>
      <div style="float: right">
        <el-button type="primary" @click="genCode">生成代码</el-button>
      </div></el-footer
    >
  </el-container>

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
            <el-form-item label="主键" prop="primary">
              <el-checkbox
                v-model="formField.primary"
                label="主键"
              ></el-checkbox>
            </el-form-item>
            <el-form-item label="Json" prop="json">
              <el-input v-model="formField.json" placeholder=""></el-input>
            </el-form-item>
            <el-form-item label="翻译" prop="transform">
              <el-input v-model="formField.transform" placeholder=""></el-input>
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

<style scoped>
.diy-gen {
  background-color: #fff;
  padding: 15px;
}
</style>


<script>
import sortTable from "sortablejs";

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
        transform: [
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
      tableData: [],
      drawer: false,
      drawerType: null,
      form: {
        checkbox: [],
        tableName: "",
        tableDiyName: "",
        controllerPackage: "system",
        menuName: "",
      },
      tableNames: [],
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
        { value: "image", label: "image" },
      ],
      filedIndex: [
        { value: "Null", labelL: "Null" },
        { value: "UNIQUE", labelL: "UNIQUE" },
        { value: "NORMAL", labelL: "NORMAL" },
        { value: "FULLTEXT", labelL: "FULLTEXT" },
      ],
    };
  },
  created: function () {
    let _that = this;
    this.$API.gen.tables.get().then(function (re) {
      _that.tableNames = re.data;
    });
  },
  mounted() {
    // 获取el-table DOM
    const el = document.querySelector(".el-table__body-wrapper  table tbody");

    let _that = this;
    //
    sortTable.create(el, {
      animation: 300,
      //拖动结束
      onEnd: function (evt) {
        console.log("new-->", evt.newIndex, "old-->", evt.oldIndex, "evt", evt);

        let old = _that.tableData[evt.oldIndex];
        // _that.tableData[evt.oldIndex] = _that.tableData[evt.newIndex];
        // _that.tableData[evt.newIndex] = old;

        _that.tableData.splice(evt.oldIndex, 1);
        _that.tableData.splice(evt.newIndex, 0, old);

        var newArray = _that.tableData.slice(0);
        _that.tableData = [];

        _that.$nextTick(function () {
          _that.tableData = newArray;
        });

        console.log(_that.tableData);
      },
    });
  },
  methods: {
    genCode() {
      if (this.tableData.length === 0) {
        this.$message.error("字段信息不能为空");
        return;
      }

      if (this.form.checkbox.length === 0) {
        this.$message.error("生成内容至少选择一项");
        return;
      }

      if (!this.form.tableDiyName) {
        this.$message.error("表名必填");
        return;
      }
      if (!this.form.controllerPackage) {
        this.$message.error("报名必须");
        return;
      }
      let data = {};

      data = this.$TOOL.objCopy(this.form);
      let _that = this;
      Object.assign(data, { fields: this.tableData });
      this.$API.gen.genCode.post(data).then((re) => {
        if (re.code === 422) {
          _that.$message.error(re.msg);
        } else {
          _that.$message.success(re.msg);
          _that.$API.system.menu.myMenus.reLoad();
        }
      });
    },
    genFields() {
      if (!this.form.tableName) {
        this.$message.error("请先选择表名");
        return;
      }

      let _that = this;

      this.$API.gen.genField
        .get({ table_name: _that.form.tableName })
        .then((re) => {
          _that.tableData = re.data;
        });
    },
    tablesChange(v) {
      this.form.tableDiyName = v;
    },
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

          console.log(this.formField);
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
