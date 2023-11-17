package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	/*
		در ایجاد یک اسلایس نیازی به وجود تعداد خانه حافظه وجود ندارد
		اسلایس فقط یک آدرس خانه حافظه. ظرفیت و طول هستش و از طریق همین سه عدد کنترل میشه و بر خلاف ارایه نسبی کنترل می گردد.
		مثلا اسلایس زیر دقیقا به یک جا اشاره می کنند
	*/
	a := []string{"a", "b", "c", "d e", "f"}
	b := a
	b[0] = "amir"
	c := append(a, b...)

	fmt.Printf("a: %q %p\n", a, a)
	fmt.Printf("b: %q %p\n", b, b)
	fmt.Printf("c: %q %p\n", c, c)

	/*
		با دستور make می شود یک اسلایس با ظرفیت و طول دلخواه ایجاد کرد و برای مقدار دهی حتما بایستی بعد از دستور make از آدرس استفاده کرد
		مثلا
		a:=make([]int,3)
		سه خانه ایجاد می کنه و برای مقدار دهی آن بایستی بگین
		a[0]=1
		a[2]=2
		اگر دستور append استفاده کنیم مقداری که به این اسلایس اضافه می شه از خانه چهارم به بعد اضافه خواهد شد
	*/
	b = append(b, "100", "200", "300")
	fmt.Printf("%q %p\n", b, b)

	changeSliceInFunc(b)
	fmt.Printf("%q %p\n", b, b)

	/*
		وقتی یک اسلایس از یک ارایه می گیریم در واقع به خانه های آرایه اشاره می کنیم و هر تغییری در اسلایس در واقع در خانه های آرایه تغییر ایجاد می کند
	*/
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("array before change")
	fmt.Println("array:", arr)
	s := arr[0:5]
	s[0] = 100
	s[1] = 200
	fmt.Println("array:", arr)
	fmt.Println("slice:", s)

	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s2 := make([]int, len(s1))
	var s3 []int
	/*
			برای استفاده از دستور کپی که یک اسلایس جدید می سازد حتما بایستی از دستور make برای ساخت اسلایس استفاده کنیم و اندازه طولش برابر با اسلایس اول باشه
		ولی برای append اینطور نیست
	*/
	copy(s2, s1)
	s3 = append(s3, s2...)
	fmt.Println(s1, s2)
	for i := range s2 {
		s2[i] *= 10
	}
	fmt.Println(s1, s2, s3)

	/*
		وقتی یک اسلایس با دستور make ساخته می شود nil نیست ولی در حالت عادی مقدار nil دارد
	*/
	var snil1 []int
	snil2 := make([]int, 0)
	fmt.Println(snil1 == nil, snil2 == nil, snil1, snil2)

	/*
		پکیچ slice امکانات خوبی برای کار با اسلایس در اختیار قرار می دهد
	*/
	//	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("find 9 in slice", slices.Contains(s1, 9))
	fmt.Println("find 19 in slice", slices.Contains(s1, 19))
	fmt.Println("max number is", slices.Max(s1))

	/*
		برای سورت کردن به روش دلخواه می توان از این روش استفاده کرد
	*/
	as := []int{1, 2, 3, 4, 5, 6, 7, 8}
	sort.Sort(AmirSlice(as))
	fmt.Println(as)
}

type AmirSlice []int

func (a AmirSlice) Len() int           { return len(a) }
func (a AmirSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a AmirSlice) Less(i, j int) bool { return a[i] > a[j] }

func changeSliceInFunc(slice []string) {
	slice[0] = "change value in func without return it"
}
