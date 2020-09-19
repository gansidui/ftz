# ftz
繁体字转换


```golang
package main

import (
	"fmt"

	"github.com/gansidui/ftz"
)

func main() {
	// 简体 --> 繁体
	fmt.Println(ftz.SimplifiedToTraditional("我爱学习"))

	// 繁体 --> 简体
	fmt.Println(ftz.TraditionalToSimplified("我愛吃飯"))
}

```