package filex

import (
	"os"

	"gopkg.in/yaml.v2"
)

func MustLoad(filename string, v interface{}) {
	content, err := os.ReadFile(filename)

	if err != nil {
		panic("read file " + filename + " error:" + err.Error())
	}
	err = yaml.Unmarshal(content, v)

	if err != nil {
		panic("parse config content error:" + err.Error())
	}
}
