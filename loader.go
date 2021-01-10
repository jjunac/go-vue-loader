package vueloader

import (
	"os"

	"github.com/jjunac/go-vue-loader/internal"
)


type Loader struct {
}

func NewLoader() Loader {
	return Loader{}
}

func (l *Loader) GetComponent(path string) (string, error) {
	vueFile, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer vueFile.Close()

	compiler := NewCompiler(vueFile)
	comp, err := compiler.Compile(componentNameFromPath(path))
	if err != nil {
		return "", err
	}
	return comp, nil
}

func componentNameFromPath(path string) string {
	return internal.PascalCaseToKebabCase(internal.FilenameWithoutExtension(path));
}
