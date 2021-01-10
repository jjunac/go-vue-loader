package main

import (
	"fmt"
	"github.com/jjunac/go-vue-loader"
)

func main() {
	loader := vueloader.NewLoader()
	fmt.Println(loader.GetComponent("examples/components/Hello.vue"))
	fmt.Println(loader.GetComponent("examples/components/Test.vue"))
	fmt.Println(loader.GetComponent("examples/components/BaseInputText.vue"))
	fmt.Println(loader.GetComponent("examples/components/TodoList.vue"))
}
