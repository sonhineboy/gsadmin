<template>
    <el-container>
        <el-container>
            <el-header>
                <div class="left-panel">
                    <!-- <el-button v-if="this.$AUTH('user.add')" type="primary" icon="el-icon-plus" @click="add"></el-button>
                    <el-button v-if="this.$AUTH('user.del')" type="danger" plain icon="el-icon-delete"
                        :disabled="selection.length == 0" @click="batch_del"></el-button> -->
                    <!-- <el-button type="primary" plain :disabled="selection.length==0">分配角色</el-button>
                          <el-button type="primary" plain :disabled="selection.length==0">密码重置</el-button> -->
                </div>
                <div class="right-panel">
                    <div class="right-panel-search">
                        <el-date-picker style="min-width: 300px;" start-placeholder="开始时间" end-placeholder="结束时间"
                            v-model="search.date" type="daterange" format="YYYY/MM/DD"
                            value-format="YYYY-MM-DD"></el-date-picker>
                        <el-input style="max-width: 200px;" v-model="search.user_name" placeholder="用户名"
                            clearable></el-input>
                        <el-button type="primary" icon="el-icon-search" @click="upsearch"></el-button>
                    </div>
                </div>
            </el-header>
            <el-main class="nopadding">
                <scTable ref="table" :apiObj="apiObj" @selection-change="selectionChange" stripe remoteSort remoteFilter>
                    <!-- <el-table-column label="ID" prop="id" width="80"></el-table-column> -->
                    <el-table-column label="#" type="index" width="50"></el-table-column>
                    <el-table-column label="用户名" prop="user_name" width="80"></el-table-column>
                    <el-table-column label="方法" prop="method" width="80"></el-table-column>
                    <el-table-column label="名称" prop="path_name" width="140"></el-table-column>
                    <el-table-column label="访问路径" prop="url_path" width="140"></el-table-column>
                    <el-table-column label="IP" prop="ip" width="120"></el-table-column>
                    <el-table-column label="数据" prop="do_data"></el-table-column>
                    <el-table-column label="创建时间" prop="created_at" width="150"></el-table-column>
                    <!-- <el-table-column label="操作" fixed="right" align="right" width="160"> -->
                    <!-- <template #default="scope">
                            <el-button-group>
                                <el-button text type="primary" size="small">查看</el-button>
                            </el-button-group>
                        </template> -->
                    <!-- </el-table-column> -->
                </scTable>
            </el-main>
        </el-container>
    </el-container>
</template>
  
<script>

export default {
    name: "systemLog",

    data() {
        return {
            apiObj: this.$API.system.systemLog.list,
            selection: [],
            search: {
                user_name: null,
                date: null
            },

        };
    },
    watch: {

    },
    mounted() {
        // this.getGroup()

        // console.log(this.$AUTH("user.add"))
    },
    methods: {
        upsearch() {

            console.log(this.search.date)

            let where = {};
            if (this.search.user_name) {
                where = { user_name: this.search.user_name };
            }

            if (this.search.date) {


                Object.assign(where, {
                    created_at: {
                        begin: this.search.date[0],
                        end: this.search.date[1]
                    }
                })

            }
            this.$refs.table.reload({ where: where }, 1);
        }

    },
};
</script>
  
<style></style>
  