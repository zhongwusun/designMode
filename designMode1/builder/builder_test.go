package builder

import (
	"fmt"
	"testing"
)

// 生成器模式
func TestBuilder(t *testing.T) {
	pancakeCook := NewPancakeCook(NewNormalPancakeBuilder())
	fmt.Printf("摊一个普通煎饼 %#v\n", pancakeCook.MakePancake())

	pancakeCook.SetPancakeBuilder(NewHealthyPancakeBuilder())
	fmt.Printf("摊一个健康的加量煎饼 %#v\n", pancakeCook.MakeBigPancake())
}
