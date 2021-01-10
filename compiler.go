package vueloader

import (
	"encoding/xml"
	"fmt"
	"io"
	"regexp"
)

var reProperties = regexp.MustCompile(`(?s)\s*module\.exports\s*=\s*{(.*)};?`)

type Compiler struct {
	decoder xml.Decoder
}

func NewCompiler(r io.Reader) Compiler {
    c := Compiler{*xml.NewDecoder(r)};
    c.decoder.Strict = false;
    return c
}

func (c *Compiler) Compile(name string) (string, error) {
	component := struct {
        Name     	string
        Template 	string
        Properties  string
        Style    	string
    }{}
	component.Name = name

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
				var script string
				script, err = c.decodeRawContent(&node)
				component.Properties = reProperties.FindStringSubmatch(script)[1]
			case node.Name.Local == "style":
				component.Style, err = c.decodeRawContent(&node)
			}
		}
		if err != nil {
			return "", err
		}
	}

	return fmt.Sprintf("Vue.component('%s', {%s,\ntemplate: `%s`\n});", component.Name, component.Properties, component.Template), nil
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
