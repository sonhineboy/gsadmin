import config from "@/config"
import http from "@/utils/request"

export default {
	user: {
		add: {
			url: `${config.API_URL}/user/add`,
			name: "添加用户",
			post: async function(data){
				return await http.post(this.url,data);
			}
		},
  up:{
   url: `${config.API_URL}/user/update`,
			name: "更新信息",
			post: async function(data){
				return await http.post(this.url,data);
			}
  },
		del:{
   url: `${config.API_URL}/user/del`,
			name: "删除用户",
			post: async function(data){
				return await http.post(this.url,data);
			}
  }
 }
}