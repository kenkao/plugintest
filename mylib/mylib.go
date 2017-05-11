package mylib

import(
	"plugin"
)

type MyLib struct {
}

var lib *MyLib

func New() *MyLib {
	if lib == nil {
		lib = new(MyLib)
	}
	return lib
}

func (lib *MyLib) Add(a int, b int) int {
	_, err := plugin.Open("xx.so")
	if err != nil {
		// 插件加载失败
		return 0
	}
	return a + b
}
