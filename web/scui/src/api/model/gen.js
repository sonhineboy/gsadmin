import config from "@/config"
import http from "@/utils/request"

export default {

 tables: {
  url: `${config.API_URL}/gen/tables`,
  name: "获取表名",
  get: async function (params) {
   return await http.get(this.url, params);
  }
 },
 genField: {
  url: `${config.API_URL}/gen/fields`,
  name: "生成字段",
  get: async function (params) {
   return await http.get(this.url, params)
  }
 },
 genCode: {
  url: `${config.API_URL}/gen/genCode`,
  name: "生成代码",
  post: async function (params) {
   return await http.post(this.url, params)
  }
 }


}