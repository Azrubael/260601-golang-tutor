/* 2022-09-01 Консольное приложение для нахождения корней квадратного уравления
   с использованием стандартной формулы
   x = (-b +/- sqrt(b^2 - 4*a*c) / 2 / a
   и _комплексных чисел_
*/

package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"os"
	"path/filepath"
	"strconv"
)

func solver(a, b, c float64) (x1, x2 complex128) {
	// Convert everything to complex once
	ax := complex(a, 0)
	bx := complex(b, 0)
	cx := complex(c, 0)

	D := bx*bx - 4*ax*cx // Discriminant as complex
	x1 = (-bx + cmplx.Sqrt(D)) / (2 * ax)
	x2 = (-bx - cmplx.Sqrt(D)) / (2 * ax)

	return x1, x2
}

func toReal(c complex128) bool {
	if math.Abs(imag(c)) > 1e-12 {
		return false
	}
	return true
}

func main() {
	fmt.Println("Решатель квадратных уравнений вида Аx^2+Вx+С=0 ,")
	var X1, X2 complex128
	var A, B, C float64 = 0, 0, 0
	var err1, err2, err3 error
	Lengh := len(os.Args)
	switch Lengh {
	case 1:
		{
			fmt.Printf("Для запуска требуется ввод %s и затем через пробел три вещественных коэффициента А, В и С>\n\n", filepath.Base(os.Args[0]))
			os.Exit(1)
		}
	case 2:
		{
			A, err1 = strconv.ParseFloat(os.Args[1], 64)
			B = 0
			C = 0
		}
	case 3:
		{
			A, err1 = strconv.ParseFloat(os.Args[1], 64)
			B, err2 = strconv.ParseFloat(os.Args[2], 64)
			C = 0
		}
	case 4:
		{
			A, err1 = strconv.ParseFloat(os.Args[1], 64)
			B, err2 = strconv.ParseFloat(os.Args[2], 64)
			C, err3 = strconv.ParseFloat(os.Args[3], 64)
		}
	}
	if (err1 != nil) || (err2 != nil) || (err3 != nil) {
		fmt.Printf("Внимание! Требуется ввод %s и трех вещественных коэффициентов А, В и С\n\n", filepath.Base(os.Args[0]))
		os.Exit(2)
	}

	fmt.Println("Введены коэффициенты >\t\t  A = ", A, "    B = ", B, "    C = ", C)
	fmt.Println("Решается уравнение вида >\t ", A, "*x^2 + ", B, "*x + ", C, " = 0")
	X1, X2 = solver(A, B, C)
	if toReal(X1) && toReal(X2) {
		fmt.Printf("Получено два вещественных корня:   X1 = %#v , \t X2 = %#v \n", real(X1), real(X2))
	} else {
		fmt.Printf("Получены значения корней:   X1 = %#v , \t X2 = %#v \n", X1, X2)
	}
	fmt.Println()
	os.Exit(0)

}
