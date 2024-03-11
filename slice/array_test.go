package Test

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"testing"
)

func TestArray(t *testing.T) {
	var myArray [10]int
	for i := 0; i < len(myArray); i++ {
		myArray[i] = 1
	}

	for i := 0; i < len(myArray); i++ {
		fmt.Println(myArray[i])
	}

	for index, value := range myArray {
		fmt.Println(index, "--", value)
	}

	for _, value := range myArray {
		fmt.Println(value)
	}
}

func TestArray2(t *testing.T) {
	var a [3]int
	fmt.Println(a)
	var b = [...]int{1, 2, 3, 4}
	fmt.Println(b)
	//冒号的意思 下标：数值
	var c = [...]int{4: 200, 5: 300, 1: 100, 0: 20}
	fmt.Println(c)
}
func TestArray3(t *testing.T) {
	// 字符串数组
	var s1 = [2]string{"hello", "world"}
	var s2 = [...]string{"你好", "世界"}
	var s3 = [...]string{1: "世界", 0: "你好"}

	// 结构体数组
	var line1 [2]image.Point
	var line2 = [...]image.Point{image.Point{X: 0, Y: 0}, image.Point{X: 1, Y: 1}}
	var line3 = [...]image.Point{{0, 0}, {1, 1}}

	// 图像解码器数组
	var decoder1 [2]func(io.Reader) (image.Image, error)
	var decoder2 = [...]func(io.Reader) (image.Image, error){
		png.Decode,
		jpeg.Decode,
	}

	// 接口数组
	var unknown1 [2]interface{}
	var unknown2 = [...]interface{}{123, "你好"}

	// 管道数组
	var chanList = [2]chan int{}

	fmt.Println(s1, s2, s3, line1, line2, line3, decoder1, decoder2, unknown1, unknown2, chanList)
}
