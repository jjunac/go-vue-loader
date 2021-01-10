package vueloader

import (
	"encoding/xml"
	"fmt"
	"io"
)

type Compiler struct {
	decoder xml.Decoder
}

func NewCompiler(r io.Reader) Compiler {
    c := Compiler{decoder: *xml.NewDecoder(r)};
    c.decoder.Strict = false;
    return c
}

func (c *Compiler) Compile() (string, error) {
	component := struct {
        Name     string
        Template string
        Script   string
        Style    string
    }{}
	component.Name = "test"

	for {
		t, err := c.decoder.Token()
		if t == nil {
			break
		}
		switch node := t.(type) {
		case xml.StartElement:
			switch {
			case node.Name.Local == "template":
				component.Template, err = c.decodeRawContent(&node)
			case node.Name.Local == "script":
				component.Script, err = c.decodeRawContent(&node)
			case node.Name.Local == "style":
				component.Style, err = c.decodeRawContent(&node)
			}
		}
		if err != nil {
			return "", err
		}
	}

	return fmt.Sprintf("Vue.component('%s', {\n%s,\ntemplate: `%s`\n});", component.Name, component.Script, component.Template), nil
}

func (c *Compiler)  decodeRawContent(node *xml.StartElement) (string, error) {
	rawContainer := struct {
		RawContent string `xml:",innerxml"`
	}{}
	err := c.decoder.DecodeElement(&rawContainer, node)
	if err != nil {
		return "", err
	}
	return rawContainer.RawContent, nil
}
