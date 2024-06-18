<template>
  <el-container v-loading="loading">
    <el-header>
      <div class="left-panel">
        <el-button type="primary" icon="el-icon-plus" @click="add"></el-button>

        <el-button
          type="danger"
          plain
          icon="el-icon-delete"
          :disabled="selection.length == 0"
          @click="batch_del"
        ></el-button>
      </div>
      <div class="right-panel">
        <div class="right-panel-search">
          <el-input v-model="searchValue" placeholder="id" clearable></el-input>
          <el-button
            type="primary"
            icon="el-icon-search"
            @click="search"
          ></el-button>
        </div>
      </div>
    </el-header>
    <el-main class="nopadding">
      <scTable
        ref="table"
        :apiObj="apiObj"
        @selection-change="selectionChange"
        stripe
        remoteSort
        remoteFilter
      >
        <el-table-column type="selection" width="50"></el-table-column>
        <el-table-column label="#" type="index" width="50"></el-table-column>
        <el-table-column label="ID" prop="id" width="50"></el-table-column>
        <el-table-column label="昵称" prop="nick_name"></el-table-column><el-table-column label="真实姓名" prop="real_name"></el-table-column><el-table-column label="年龄" prop="age"></el-table-column><el-table-column label="状态" prop="status"></el-table-column><el-table-column label="在线" prop="online"></el-table-column>
        <el-table-column label="时间" prop="created_at"></el-table-column>
        <el-table-column label="操作" fixed="right" width="100">
          <template #default="scope">
            <el-link
              type="primary"
              @click="actions('view', scope.$index, scope.row)"
              icon="ElIconView"
            />
            &nbsp;
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
      </scTable>
    </el-main>
  </el-container>
  <my-form ref="myForm" @confirm="confirm" v-model="v"></my-form>
</template>
<script>
import myForm from "./form";
export default {
  name: "userMember",
  components: { myForm },
  data() {
    return {
      apiObj: this.$API.userMember.index,
      selection: [],
      searchValue: "",
      loading: false,
      v: true,
    };
  },
  mounted() {},
  methods: {
    selectionChange(items) {
      this.selection = items;
    },
    actions(type, index, data) {
      console.log(type, index, data);
      switch (type) {
        case "del":
          this.del(data, index);
          break;

        case "edit":
          this.edit(data);
          break;
        case "view":
          this.$refs.myForm.open("view", data);
          break;
        default:
          break;
      }
    },
    async del(row, index) {
      var reqData = { ids: [row.id] };
      var res = await this.$API.userMember.delete.post(reqData);
      if (res.code == 200) {
        this.$refs.table.tableData.splice(index, 1);
        this.$message.success("删除成功");
      } else {
        this.$alert(res.message, "提示", { type: "error" });
      }
    },
    edit(row) {
      this.$refs.myForm.open("edit", row);
    },
    batch_del() {
      var ids = [];
      var _that = this;
      this.selection.forEach((v) => {
        ids.unshift(v.id);
      });
      this.$confirm("确定删除选中的"+this.selection.length+"项吗？", "提示", {
        type: "warning",
      })
        .then(async () => {
          _that.loading = true;
          let res = await _that.$API.userMember.delete.post({ ids: ids });
          _that.loading = false;

          if (res.code === 200) {
            _that.$message.success("删除成功");
            _that.$refs.table.refresh();
          } else {
            _that.$message.error("删除失败");
          }
        })
        .catch(() => {});
    },
    add() {
      this.$refs.myForm.open();
    },
    search() {
      this.$refs.table.reload({ where: { id: this.searchValue } });
    },
    async confirm(data) {
      let _that = this;
      let re = { code: "422", msg: "未知事件", message: "未知事件" };
      switch (this.$refs.myForm.mode) {
        case "add":
          re = await this.$API.userMember.save.post(data);
          break;
        case "edit":
          re = await this.$API.userMember.edit.post(data.id, data);
          break;

        default:
          console.log(this.$refs.myForm.mode, data);
          break;
      }
      if (re.code === 200) {
        _that.$message.success("操作成功");
        _that.$refs.table.refresh();
        _that.$refs.myForm.close();
      } else {
        _that.$message.error("操作失败:" + re.msg);
      }
    },
  },
};
</script>
