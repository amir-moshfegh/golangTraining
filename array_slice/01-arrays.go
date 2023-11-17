package main

import (
	"fmt"
	"strings"
)

const seperator = "="
const lenAllArray = 5

func printIntArray(numbers [lenAllArray]int) {
	fmt.Printf("%T: %d (len=%d cap=%d)\n", numbers, numbers, len(numbers), cap(numbers))
}

func appendItemToStringArray(arr [lenAllArray]string) [lenAllArray]string {
	fmt.Printf("array in func is: %q, in %p memory address.\n", arr, &arr)
	arr[0] = "Amir"
	return arr
}

func appendItemToStringArrayWithReferenceMethod(arr *[lenAllArray]string) {
	fmt.Printf("array in func with referance is: %q, in %p memory address.\n", arr, &arr)
	arr[0] = "Moshfegh"
}

func main() {
	/*
			زمانی که یک آرایه تعریف می شود به همان میزان طولی که برایش در نظر گرفته می شود از نوع آرایه خانه خالی در نظر گرفته میش ود
			مثلا متغییر a یک آرایه با تعداد 5 صفر در خانه حافظه ایجاد می کند و
			متغییر b یک آرایه با تعداد 5 خانه حاوی "" در حافظه ایجاد می کنه
		اگر یکی از خانه های آرایه پر شود فقط همان خانه مقدار دارد و الباقی همان مقدار اولیه را دارد
	*/
	var a [lenAllArray]int
	a[0] = 25
	a[1] = 45
	printIntArray(a)

	/*
		برای پاس دادن آرایه به یک تابع. کل اون آرایه در تابع ک‍پی میشه و اگر بخوایم آدرسشو بفرستیم بایستی از * & استفاده کنیم
		در این روش می توانیم شماره خانه حافظه رو هم مشخص کنیم مثلا در مثال زیر در خانه سوم آرایه گفتیم که حرف c قرار بگیرد
		arr1 := [5]int{1:10,2:40}
		[0 10 40 0 0]
	*/
	//Short variable declarations
	b := [lenAllArray]string{"a", "b", 3: "c"}
	fmt.Printf("array init is: %q, in %p memory address.\n", b, &b)

	// send array to func with value
	b = appendItemToStringArray(b)
	fmt.Printf("array after change is: %q, in %p memory address.\n", b, &b)

	// send array to func with reference
	appendItemToStringArrayWithReferenceMethod(&b)
	fmt.Printf("array after change is: %q, in %p memory address.\n", b, &b)

	fmt.Printf("%T: %q (len=%d cap=%d)\n", b, b, len(b), cap(b))
	fmt.Println(strings.Repeat(seperator, 30))

	/*
		در زمان کپی کردن آرایه. دو آرایه هیچ نسبتی با هم ندارند و در دو آدرس متفاوت ذخیره می شوند
		پس تغییرات هر یک مربوط به خود آن است و در دیگری تاثیری ندارد
	*/
	// ... makes the compiler determine the length
	c := [...]rune{'a', 'b', 'c', 'd', 'f'}
	d := c //g copy of h; g is a new array
	d[0] = 'f'
	fmt.Printf("%T: %d (len=%d cap=%d)\n", c, c, len(c), cap(c))
	fmt.Printf("%T: %d (len=%d cap=%d)\n", d, d, len(d), cap(d))
	for j := 0; j < len(c); j++ {
		fmt.Printf("array index %d value is %s and binery value is %b \n", j, string(c[j]), c[j])
	}
	fmt.Println(strings.Repeat(seperator, 30))

	// It is possible to create multidimensional arrays.
	s := [2][2]string{
		{"Amir", "Sima"},
		{"Javad", "Nima"}, // add comma in the end of sentence for simple rules
	}
	t := [2][1]string{
		{"sisi"},
		{"rabna"},
	}
	fmt.Println(s)
	fmt.Println(t)

	/*
		در آرایه از سورت هم نمی تون استفاده کرد
		در آرایه از دستور make و append نمی توان استفاده کرد
	*/
}
