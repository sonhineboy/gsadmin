package test

import (
	"fmt"
	"github.com/sonhineboy/gsadmin/service/app/models"
	"reflect"
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
