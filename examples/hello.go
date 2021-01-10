package main

import (
	"fmt"
	"github.com/jjunac/go-vue-loader"
	"log"
	"os"
)

func main() {
	vueFile, err := os.Open("examples/components/Hello.vue")
	if err != nil {
		log.Fatal(err)
	}
	defer vueFile.Close()

	compiler := vueloader.NewCompiler(vueFile)
	comp, err := compiler.Compile("hello")
	fmt.Println(comp)
}
