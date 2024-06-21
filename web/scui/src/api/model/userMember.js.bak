import config from "@/config"
import http from "@/utils/request"

export default {
 index: {
  url: `${config.API_URL}/userMember/index`,
  name: "列表",
  get: async function (params) {
   return await http.get(this.url, params)
  }
 },
 get: {
  url: `${config.API_URL}/userMember/`,
  name: "单条信息",
  get: async function (id) {
   return await http.get(this.url + id)
  }
 },
 save: {
  url: `${config.API_URL}/userMember/save`,
  name: "添加信息",
  post: async function (params) {
   return http.post(this.url, params)
  },
 },
 edit: {
  url: `${config.API_URL}/userMember/edit/`,
  name: "编辑信息",
  post: async function (id, params) {
   return http.post(this.url + id, params)
  },
 },
 delete: {
  url: `${config.API_URL}/userMember/delete`,
  name: "删除信息",
  post: async function (params) {
   return http.post(this.url, params)
  },
 }
}