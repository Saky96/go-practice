package main

import "fmt"

type MapOfString map[string]string

func (m MapOfString) output() {
	fmt.Println(m)
}

func main() {
	os := MapOfString{
		"mac":     "sequoia",
		"windows": "11",
	}

	fmt.Println(os)

	os["linux"] = "ubuntu"

	//fmt.Println(os)

	os.output()

	for k, v := range os {
		fmt.Println(k, v)
	}
}
