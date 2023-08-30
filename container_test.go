package injector

import (
	"context"
	"fmt"
	"testing"
)

type A struct {
	C1 *C `inject:""`
	C2 *C `inject:"alias:zhangsan"`
	C3 *C `inject:"alias:wuwei"`
	B  *B `inject:"alias:lisi,opts:NR"`
}

func (a *A) Provide(context.Context) interface{} {
	return a
}

type B struct {
	Name string
	Tab  string
}

func (b *B) Provide(context.Context) interface{} {
	return b
}

type C struct {
	Name string
	Age  int
}

func (c *C) Provide(context.Context) interface{} {
	return c
}

func TestInject(t *testing.T) {
	container := New()
	if err := container.Provide(&C{
		Name: "lisi",
		Age:  22,
	}); err != nil {
		t.Fatalf(err.Error())
		return
	}
	if err := container.Provide(&C{
		Name: "zhangsan",
		Age:  28,
	}, WithAlias("zhangsan")); err != nil {
		t.Fatalf(err.Error())
		return
	}
	if err := container.Provide(&C{
		Name: "wuwei",
		Age:  30,
	}, WithAlias("wuwei")); err != nil {
		t.Fatalf(err.Error())
		return
	}
	if err := container.Provide(&B{
		Name: "b",
		Tab:  "b",
	}, WithAlias("lisi")); err != nil {
		t.Fatalf(err.Error())
		return
	}
	a := &A{}
	if err := container.Provide(a); err != nil {
		t.Fatalf(err.Error())
		return
	}
	if err := container.Populate(); err != nil {
		t.Fatalf(err.Error())
	}
	//if err := container.Populate(); err != nil {
	//	t.Fatalf(err.Error())
	//}
	fmt.Println(a.C1, a.C2, a.C3, a.B)
	//fmt.Println(a.C2, a.B)
}
