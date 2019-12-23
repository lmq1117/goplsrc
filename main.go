package main

import "fmt"

type Celsius float64    // 摄氏温度
type Fahrenheit float64 // 华氏温度

const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度
	FreezingC     Celsius = 0       // 结冰点温度
	BoilingC      Celsius = 100     // 沸水温度
)

func main() {
	//fmt.Println(BoilingC - FreezingC)
	//boilingF := CToF(BoilingC)
	//fmt.Printf("%g\n",boilingF)
	var a float64 = 100
	var b float64 = 0
	var c float64 = 36.5
	var d float64 = 38
	var e float64 = 42
	fmt.Printf("摄氏温度%g度是华氏温度%g度\n", Celsius(a), CToF(Celsius(a)))
	fmt.Printf("摄氏温度%g度是华氏温度%g度\n", Celsius(b), CToF(Celsius(b)))
	fmt.Printf("摄氏温度%g度是华氏温度%g度\n", Celsius(c), CToF(Celsius(c)))
	fmt.Printf("摄氏温度%g度是华氏温度%g度\n", Celsius(d), CToF(Celsius(d)))
	fmt.Printf("摄氏温度%g度是华氏温度%g度\n", Celsius(e), CToF(Celsius(e)))

}

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
