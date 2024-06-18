package system

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeptList(c *gin.Context) {
	data := "{\"code\":200,\"data\":[{\"id\":\"1\",\"parentId\":\"0\",\"label\":\"华南分部\",\"date\":\"2022-10-10 08:00:00\",\"remark\":\"\",\"status\":1,\"sort\":1,\"children\":[{\"id\":\"11\",\"parentId\":\"1\",\"label\":\"售前客服部\",\"date\":\"2022-10-10 08:00:00\",\"remark\":\"\",\"status\":1,\"sort\":2},{\"id\":\"12\",\"parentId\":\"1\",\"label\":\"技术研发部\",\"date\":\"2022-10-10 08:00:00\",\"remark\":\"软件开发&测试\",\"status\":0,\"sort\":3}]},{\"id\":\"2\",\"parentId\":\"0\",\"label\":\"华东分部\",\"date\":\"2022-10-10 08:00:00\",\"remark\":\"\",\"status\":1,\"sort\":4,\"children\":[{\"id\":\"21\",\"parentId\":\"2\",\"label\":\"售前客服部\",\"date\":\"2022-10-10 08:00:00\",\"remark\":\"\",\"status\":1,\"sort\":5},{\"id\":\"22\",\"parentId\":\"2\",\"label\":\"技术研发部\",\"date\":\"2022-10-10 08:00:00\",\"remark\":\"\",\"status\":1,\"sort\":6}]}],\"message\":\"\"}"

	var dataMap map[string]interface{}
	_ = json.Unmarshal([]byte(data), &dataMap)
	c.JSON(http.StatusOK, dataMap)
}

func DemoUser(c *gin.Context) {
	user := "{\"code\":200,\"data\":{\"total\":41,\"rows\":[{\"id\":\"410000199512025445\",\"name\":\"用户登录\",\"url\":\"/oauth/token\",\"type\":\"GET\",\"code\":\"401\",\"cip\":\"95.214.92.71\",\"user\":\"魏磊\",\"time\":\"2012-09-05 18:08:06\"},{\"id\":\"520000198407304275\",\"name\":\"用户登录\",\"url\":\"/oauth/token\",\"type\":\"GET\",\"code\":\"500\",\"cip\":\"129.166.168.115\",\"user\":\"史平\",\"time\":\"1993-08-25 05:08:41\"},{\"id\":\"230000199511014097\",\"name\":\"用户登录\",\"url\":\"/oauth/token\",\"type\":\"GET\",\"code\":\"200\",\"cip\":\"186.160.119.210\",\"user\":\"于杰\",\"time\":\"1998-02-09 13:42:41\"},{\"id\":\"990000197205236780\",\"name\":\"用户登录\",\"url\":\"/oauth/token\",\"type\":\"POST\",\"code\":\"401\",\"cip\":\"193.14.94.222\",\"user\":\"田娟\",\"time\":\"2009-11-05 12:37:58\"},{\"id\":\"640000200911201176\",\"name\":\"用户登录\",\"url\":\"/oauth/token\",\"type\":\"GET\",\"code\":\"500\",\"cip\":\"12.69.226.121\",\"user\":\"邵涛\",\"time\":\"1989-08-25 05:33:06\"},{\"id\":\"710000198709077149\",\"name\":\"用户登录\",\"url\":\"/oauth/token\",\"type\":\"GET\",\"code\":\"500\",\"cip\":\"116.42.15.149\",\"user\":\"魏娟\",\"time\":\"2008-09-07 10:35:41\"},{\"id\":\"360000197302144442\",\"name\":\"用户登录\",\"url\":\"/oauth/token\",\"type\":\"GET\",\"code\":\"500\",\"cip\":\"239.114.89.252\",\"user\":\"何敏\",\"time\":\"2016-07-14 20:41:53\"},{\"id\":\"500000199407048831\",\"name\":\"用户登录\",\"url\":\"/oauth/token\",\"type\":\"GET\",\"code\":\"401\",\"cip\":\"95.203.182.52\",\"user\":\"傅刚\",\"time\":\"1995-03-09 23:53:50\"},{\"id\":\"450000198307033289\",\"name\":\"用户登录\",\"url\":\"/oauth/token\",\"type\":\"GET\",\"code\":\"500\",\"cip\":\"60.137.195.112\",\"user\":\"康伟\",\"time\":\"1970-01-13 19:18:19\"},{\"id\":\"220000200908305857\",\"name\":\"用户登录\",\"url\":\"/oauth/token\",\"type\":\"POST\",\"code\":\"500\",\"cip\":\"106.167.1.227\",\"user\":\"何秀英\",\"time\":\"1976-01-01 14:27:22\"},{\"id\":\"450000201411302578\",\"name\":\"用户登录\",\"url\":\"/oauth/token\",\"type\":\"GET\",\"code\":\"401\",\"cip\":\"241.235.41.98\",\"user\":\"吴涛\",\"time\":\"2003-10-22 18:03:07\"},{\"id\":\"81000020100227047X\",\"name\":\"用户登录\",\"url\":\"/oauth/token\",\"type\":\"GET\",\"code\":\"200\",\"cip\":\"120.218.167.233\",\"user\":\"毛军\",\"time\":\"1983-12-27 22:13:35\"},{\"id\":\"420000198411242296\",\"name\":\"用户登录\",\"url\":\"/oauth/token\",\"type\":\"GET\",\"code\":\"500\",\"cip\":\"123.130.209.70\",\"user\":\"张霞\",\"time\":\"2008-06-17 01:36:50\"},{\"id\":\"460000201605269026\",\"name\":\"用户登录\",\"url\":\"/oauth/token\",\"type\":\"POST\",\"code\":\"500\",\"cip\":\"79.243.107.181\",\"user\":\"段超\",\"time\":\"1980-01-16 19:59:55\"},{\"id\":\"610000200506056280\",\"name\":\"用户登录\",\"url\":\"/oauth/token\",\"type\":\"POST\",\"code\":\"200\",\"cip\":\"14.23.92.219\",\"user\":\"宋涛\",\"time\":\"1986-10-19 11:28:05\"},{\"id\":\"370000197405268159\",\"name\":\"用户登录\",\"url\":\"/oauth/token\",\"type\":\"POST\",\"code\":\"500\",\"cip\":\"233.81.94.91\",\"user\":\"石勇\",\"time\":\"2012-04-29 21:58:34\"},{\"id\":\"440000200407105727\",\"name\":\"用户登录\",\"url\":\"/oauth/token\",\"type\":\"POST\",\"code\":\"200\",\"cip\":\"178.29.69.222\",\"user\":\"贾超\",\"time\":\"1989-09-28 16:26:08\"},{\"id\":\"45000019760731722X\",\"name\":\"用户登录\",\"url\":\"/oauth/token\",\"type\":\"GET\",\"code\":\"500\",\"cip\":\"117.35.28.120\",\"user\":\"汤强\",\"time\":\"1975-05-24 14:29:46\"}]},\"message\":\"\"}"
	var userMap map[string]interface{}
	_ = json.Unmarshal([]byte(user), &userMap)
	c.JSON(http.StatusOK, userMap)

}

func OrderDemo(c *gin.Context) {

	//
	//model := models.Order{
	//	UserName: "zs",
	//	Age:      10,
	//}
	//
	//err := re.Insert(&model)
	//if err != nil {
	//	c.JSON(http.StatusOK, gin.H{
	//		"code":  202,
	//		"error": err.Error(),
	//	})
	//	return
	//}
	//c.JSON(http.StatusOK, model)

	//type request struct {
	//	A string `json:"a" binding:"required"`
	//	B string `json:"b"`
	//}
	//var model models.Order
	//err := c.ShouldBind(&model)
	//if err != nil {
	//
	//	c.JSON(http.StatusOK, gin.H{
	//		"error": err.Error(),
	//	})
	//	return
	//}

}
