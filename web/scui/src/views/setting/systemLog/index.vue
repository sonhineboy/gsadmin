<template>
  <el-container>
    <el-container>
      <el-header>
        <div class="left-panel"></div>
        <div class="right-panel">
          <div class="right-panel-search">
            <el-date-picker
              style="min-width: 300px"
              start-placeholder="开始时间"
              end-placeholder="结束时间"
              v-model="search.date"
              type="daterange"
              format="YYYY/MM/DD"
              value-format="YYYY-MM-DD"
            ></el-date-picker>
            <el-input
              style="max-width: 200px"
              v-model="search.user_name"
              placeholder="用户名"
              clearable
            ></el-input>
            <el-button
              type="primary"
              icon="el-icon-search"
              @click="upsearch"
            ></el-button>
          </div>
        </div>
      </el-header>
      <el-main class="nopadding">
        <scTable ref="table" :apiObj="apiObj" stripe remoteSort remoteFilter>
          <el-table-column label="#" type="index" width="50"></el-table-column>
          <el-table-column
            label="用户名"
            prop="user_name"
            width="80"
          ></el-table-column>
          <el-table-column
            label="方法"
            prop="method"
            width="80"
          ></el-table-column>
          <el-table-column
            label="名称"
            prop="path_name"
            width="140"
          ></el-table-column>
          <el-table-column
            label="访问路径"
            prop="url_path"
            width="260"
          ></el-table-column>
          <el-table-column label="IP" prop="ip" width="120"></el-table-column>
          <el-table-column label="数据" prop="do_data"></el-table-column>
          <el-table-column
            label="创建时间"
            prop="created_at"
            width="150"
          ></el-table-column>
        </scTable>
      </el-main>
    </el-container>
  </el-container>
</template>
  
<script>
import scEcharts from "@/components/scEcharts";
export default {
  name: "systemLog",
  components: {
    scEcharts,
  },

  data() {
    return {
      apiObj: this.$API.system.systemLog.list,
      selection: [],
      search: {
        user_name: null,
        date: null,
      },
    };
  },
  watch: {},
  mounted() {
    // this.getGroup()
    // console.log(this.$AUTH("user.add"))
  },
  methods: {
    upsearch() {
      console.log(this.search.date);

      let where = {};
      if (this.search.user_name) {
        where = { user_name: this.search.user_name };
      }

      if (this.search.date) {
        Object.assign(where, {
          created_at: {
            begin: this.search.date[0],
            end: this.search.date[1],
          },
        });
      }
      this.$refs.table.reload({ where: where }, 1);
    },
  },
};
</script>
  
<style></style>
  