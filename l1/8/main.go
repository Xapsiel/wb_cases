package main

import "fmt"

type BitReplacer struct {
	b int64
}

func (b *BitReplacer) Set(index uint8, value int64) {
	mask := int64(1)<<index - 1

	if value == 1 {
		fmt.Println(b.b | mask)
		return
	}
	fmt.Println(b.b &^ mask)
}

func main() {
	b := &BitReplacer{
		b: 4,
	}
	b.Set(1, 0)
}
