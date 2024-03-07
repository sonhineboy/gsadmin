package test

import (
	"fmt"
	"github.com/sonhineboy/gsadmin/service/app"
	"github.com/sonhineboy/gsadmin/service/app/models"
	"github.com/sonhineboy/gsadmin/service/app/repositorys"
	"testing"
)

func TestSelect(t *testing.T) {
	app.TestLoad()
	testRe := repositorys.NewTestRepository()
	var data []models.Article
	page := testRe.Page(1, 10, "id", &data)
	fmt.Println(page)

	//add
	//var data models.Article
	//data.Title = "标题"
	//data.Cid = 2
	//data.CommentCount = 10
	//data.Content = "这里是内容"
	//data.Desc = "这里是描述"
	//data.Img = "xxx"
	//data.ReadCount = 10
	//
	//err :=testRe.Add(&data)
	//if err !=nil {
	//	t.Error(err)
	//}

	//updata
	//db :=testRe.UpdateById(4,map[string]interface{}{"read_count":40})
	//fmt.Println(db.RowsAffected)

	//updataByWhere
	//db := testRe.UpdateByWhere(map[string]interface{}{"read_count": 50}, map[string]interface{}{"read_count": 40})
	//fmt.Println(db.RowsAffected)

}
