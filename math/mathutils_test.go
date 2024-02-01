package mathutils

import (
	"fmt"
	"testing"
)

func TestMax(t *testing.T) {
	max1 := Max(1, 2)
	fmt.Println("较大值是: ", max1)
	max2 := Max(1.2, 3.4)
	fmt.Println("较大值是: ", max2)
}

func TestMin(t *testing.T) {
	max1 := Min(-1, -2)
	fmt.Println("较小值是: ", max1)
	max2 := Min(-1.2, -3.4)
	fmt.Println("较小值是: ", max2)
}
