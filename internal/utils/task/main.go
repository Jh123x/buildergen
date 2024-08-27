package main

import (
	"fmt"

	"github.com/Jh123x/buildergen/internal/consts"
	"github.com/Jh123x/buildergen/internal/utils"
)

func main() {
	counter := make(map[int]string, 25)
	for _, word := range consts.Keywords {
		hashVal := utils.Attempt2(word)
		counter[hashVal] = word
	}
	fmt.Println(counter)
	_, counterMp := is_valid_mod(73, counter)

	fmt.Printf("[]string{")
	for i := 0; i < 73; i++ {
		fmt.Printf("\"%s\",", counterMp[i])
	}
	fmt.Println("}")
}

func is_valid_mod(mod int, counter map[int]string) (bool, map[int]string) {
	current := map[int]string{}
	for k, v := range counter {
		key := k % mod
		if _, ok := current[key]; ok {
			return false, nil
		}
		current[key] = v
	}
	return true, current
}
