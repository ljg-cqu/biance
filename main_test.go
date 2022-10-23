package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMe(t *testing.T) {
	type Token []string
	tokens := []string{"1", "3"}

	tokensBytes, _ := json.Marshal(tokens)
	fmt.Println(string(tokensBytes))
}

//func ParseFile(filename string) *Config {
//	return verifyConfig(parseContent(readFile(filename)))
//}
//
//func readFile(filename string) []byte {
//	content, err := ioutil.ReadFile(filename)
//	if err != nil {
//		panic(errors.Wrapf(err, "failed to read file %q", filename))
//	}
//	return content
//}
