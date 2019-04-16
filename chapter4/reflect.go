package chapter4

import (
	"bytes"
	"fmt"
	"reflect"
	"unsafe"
)

type ReflectUsage struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (ref ReflectUsage) String() string {
	return fmt.Sprintf("Name: %s, Age: %d", ref.Name, ref.Age)
}

func (ref *ReflectUsage) AddAge(add int) int {
	ref.Age += add
	return ref.Age
}

func (ref ReflectUsage) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("Name: %s", ref.Name))
	return buf.Bytes(), nil
}

func Usage() {
	var example ReflectUsage
	example.Name = "XieWei"
	example.Age = 20

	// 获取类型的两种方法
	typ := reflect.TypeOf(example)
	fmt.Printf("%T\n", example)
	fmt.Println(typ)

	// 获取值的两种方法
	val := reflect.ValueOf(example)
	fmt.Printf("%#v\n", example)
	fmt.Printf("%v\n", example)
	fmt.Println(val)

	// 结构体包含：字段(标签、值）、方法
	fmt.Println(typ.NumField(), typ.NumMethod())
	fmt.Println(val.NumField(), val.NumMethod())

	// 通过 type 获取标签属性
	fmt.Println(typ.FieldByName("Name"))
	fmt.Println(typ.FieldByName("Age"))

	// 通过 value 获取标签值
	fmt.Println(val.FieldByName("Name"))
	fmt.Println(val.FieldByName("Age"))

	// 函数的个数是根据传值的才计数，函数的排序根据函数名称的字母，函数可以调用
	fmt.Println(typ.NumMethod(), typ.Method(0))
	fmt.Println(typ.NumMethod(), typ.Method(1))

	methodOne := val.Method(1)
	args := make([]reflect.Value, 0)
	result := methodOne.Call(args)
	fmt.Println(result)

	methodTwo := val.MethodByName("MarshalJSON")
	argsTwo := make([]reflect.Value, 0)
	resultTwo := methodTwo.Call(argsTwo)
	fmt.Println(string(resultTwo[0].Bytes()))

	// 可以重新对结构体赋值操作，前提是获得指针
	valCanSet := reflect.ValueOf(&example)
	ptr := valCanSet.Elem()
	ptr.FieldByName("Age").SetInt(100)
	fmt.Println(example)

}

func UnsafeUsage() {
	var example ReflectUsage
	example.Name = "XieWei"
	example.Age = 20

	typ := reflect.TypeOf(unsafe.Sizeof(example))
	fmt.Println(typ)

	fmt.Println(unsafe.Sizeof(example))

	ptr := unsafe.Pointer(&example) //  第一个字段地址
	fmt.Println(ptr)
	fmt.Println(*(*string)(ptr)) //  强制类型转换成第一个字段类型，获取值

	ptrOfSecondField := unsafe.Pointer(uintptr(ptr) + unsafe.Offsetof(example.Age))
	fmt.Println(ptrOfSecondField)
	fmt.Println(*(*int)(ptrOfSecondField))

	*(*int)(ptrOfSecondField) = 32
	fmt.Println(example)

}
