<template>
  <el-container>
    <el-header></el-header>
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
        <el-table-column label="昵称" prop="nick_name"></el-table-column>
        <el-table-column label="姓名" prop="real_name"></el-table-column>
        <el-table-column label="年龄" prop="age"></el-table-column>
        <el-table-column label="状态" prop="status"></el-table-column>
        <el-table-column label="在线" prop="online"></el-table-column>
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
      </scTable>
    </el-main>
  </el-container>
</template>
<script>
export default {
  name: "userMember",
  components: {},
  data() {
    return {
      apiObj: this.$API.userMember.index,
    };
  },
  methods: {
    selectionChange(items) {
      console.log(items);
    },
    actions(type, index, data) {
      console.log(type, index, data);
    },
  },
};
</script>