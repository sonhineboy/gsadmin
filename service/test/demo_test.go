package test

import (
	"bufio"
	"fmt"
	"github.com/sonhineboy/gsadmin/service/app/models"
	"github.com/sonhineboy/gsadminGen"
	"github.com/sonhineboy/gsadminGen/pkg"
	"os"
	"reflect"
	"strconv"
	"sync"
	"testing"
)

func TestStringTo(t *testing.T) {

	b := sync.WaitGroup{}
	a := "xxx"
	b.Add(1)
	go func() {
		fmt.Println(a)
		b.Done()
	}()

	b.Wait()
}

type b struct {
	Name string
}

func (receiver b) GetC() {

	fmt.Printf("%p", &receiver)
	receiver.Name = "xcc"
}

func (receiver b) GetName() string {
	fmt.Printf("%p", &receiver)
	return receiver.Name

}

func newB() *b {
	return &b{}
}

func TestP(t *testing.T) {

	role := models.Role{}

	var a *int
	var c *int
	var b int
	b = 3
	a = &b
	c = a
	fmt.Println(a, "---", c)

	//model.ID = 45
	model, _ := reflect.TypeOf(role).FieldByName("GAD_MODEL")

	f := model.Type.Field(0)

	fmt.Println(model.Type.NumField())
	fmt.Println(f.Tag.Get("gorm"), f.Tag.Get("json"))

	fmt.Println(reflect.TypeOf(&model))
}

func TestVersion(t *testing.T) {

	str := "he,llo, , world! Hello 你好"
	m := make(map[string]int)

	// s为临时字符串，用于拼接每个字母，组成单词
	s := ""

	//  temp用于统计每个单词中的字母的个数
	temp := 0

	// 遍历这个str，获取每个字符char
	for _, char := range str {
		// 当char是26个字母的大写或小写
		if (char >= 65 && char <= 90) || (char >= 97 && char <= 122) {
			temp++
			s += string(char)
		} else { // 如果遍历到的字符 不是26个字母
			if temp != 0 { // 此处判断用于防止将非26个字母的字符加入到map中
				m[s]++ // 将这个单词添加到map中
				s = "" // 将临时字符串s清空
			}
			temp = 0 // 每个单词中字母的个数清理
		}
	}

	// 防止当字符串str的最后一个字符为字母时，不会将最后一个单词添加进map
	if s != "" {
		m[s]++
	}

	for k, v := range m {
		fmt.Printf("%s:%d\n", k, v)
	}

}

// deleteMarkOfWord 删除单词末尾的标点符号
func deleteMarkOfWord(word string) string {
	// 从单词末端开始遍历，删除标点符号
	for i := len(word) - 1; i >= 0; i-- {
		if (word[i] >= 97 && word[i] <= 122) || (word[i] >= 65 && word[i] <= 90) {
			break
		} else {
			word = word[:len(word)-1]
		}
	}
	return word
}

func TestSocket(t *testing.T) {
	//print(test(3, 6))

	a := make([]int, 2, 5)
	fmt.Println(a[0:])
}

func test(b int, c int) (a int) {
	a = b + c
	return
}

func TestSlice(t *testing.T) {
	defer func() {

		fmt.Println("defer :-->1")
	}()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	defer func() {
		fmt.Println("defer:--->3")
	}()

	car := make([]int, 0, 10)

	for i := 0; i <= 9; i++ {
		car = append(car, i)
	}

	//car = nil
	fmt.Println("长度--》", len(car))

}

func TestDemo(t *testing.T) {

	id, err := strconv.Atoi("asfd")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(id)
}

func TestGenModel(t *testing.T) {

	/**

	default:""
	describe:""
	index:"Null"
	isNull:true
	json:"content"
	name:"content"
	primary:false
	transform:""
	type:"longtext"

	*/

	fields := []pkg.Field{
		{
			Name:      "userName",
			Json:      "user_name",
			Default:   "",
			Describe:  "用户名",
			Primary:   false,
			Index:     "UNIQUE",
			IsNull:    false,
			Type:      "varchar",
			Transform: "用户名",
		},
		{
			Name:      "age",
			Json:      "age",
			Default:   "0",
			Describe:  "年龄",
			Primary:   false,
			Index:     "UNIQUE",
			IsNull:    true,
			Type:      "int",
			Transform: "年龄",
		},
	}

	v := pkg.TableModal{
		Name:   "user_member",
		Fields: fields,
	}

	var err error

	fileName := "user_member"

	err = gsadminGen.GenController("../app/controllers/demo/"+gsadminGen.UnderToConvertSoreLow(fileName)+"Controller.go", v, "demo")

	if err != nil {
		t.Fatal(err.Error())
	}

	err = gsadminGen.GenModel("../app/models/"+gsadminGen.UnderToConvertSoreLow(fileName)+".go", v)

	if err != nil {
		t.Fatal(err.Error())
	}

	err = gsadminGen.GenRequest("../app/requests/"+gsadminGen.UnderToConvertSoreLow(fileName)+"Request.go", v)

	if err != nil {
		t.Fatal(err.Error())
	}

	err = gsadminGen.GenRepository("../app/repositorys/"+gsadminGen.UnderToConvertSoreLow(fileName)+"Repository.go", v)

	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestHasSlice(t *testing.T) {
	//fmt.Println(dir)
}

func TestRead(t *testing.T) {

	// 打开文件
	file, err := os.Open("orm_test.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// 创建一个新的 scanner
	scanner := bufio.NewScanner(file)

	i := 1

	// 按行读取文件
	for scanner.Scan() {
		line := scanner.Text()     // 获取当前行的内容
		fmt.Println(i, "  ", line) // 打印当前行

		i++
	}

	// 检查读取过程中是否出现错误
	if err := scanner.Err(); err != nil {
		fmt.Println(i, "  ", err)

	}
}
