package main

import (
	"fmt"
	"log"
	"os"
	"github.com/jjunac/go-vue-loader"
)

func main() {
	vueFile, err := os.Open("examples/data/Hello.vue")
	if err != nil {
		log.Fatal(err)
	}
	defer vueFile.Close()

	compiler := vueloader.NewCompiler(vueFile)
	comp, err := compiler.Compile()
	fmt.Println(comp)
}
