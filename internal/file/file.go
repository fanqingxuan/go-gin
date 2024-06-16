package filex

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func MustLoad(filename string, v interface{}) error {
	content, err := os.ReadFile(filename)

	if err != nil {
		return fmt.Errorf("read file error," + err.Error())
	}
	err = yaml.Unmarshal(content, v)

	if err != nil {
		return fmt.Errorf("parse config content error:" + err.Error())
	}
	return nil
}
