package main

/* Скрипт для работьі с координатами вещественных чисел,
преобразующий полярные координаты в декартовы
*/

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"runtime" //доступ к св-вам прогр. времени выполнения (опр. платформы)
)

type polar struct {
	radius float64
	phi    float64
}
type cartesian struct {
	x float64
	y float64
}

var prompt = "Enter a radius and an angle (in degrees), e.g., 12.5 90, " + "or %s to quit."

func init() {
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
	} else { // Unix-подобная система
		prompt = fmt.Sprintf(prompt, "Ctrl+D")
	}
}

func createSolver(questions chan polar) chan cartesian {
	answers := make(chan cartesian)
	go func() {
		for {
			polarCoord := <-questions
			phi := polarCoord.phi * math.Pi / 180.0 //преобр. градусов в радианы
			x := polarCoord.radius * math.Cos(phi)
			y := polarCoord.radius * math.Sin(phi)
			answers <- cartesian{x, y}
		}
	}()
	return answers
}

const result = "Polar radius=%.02fmm q=%.02fdeg | Cartesian x=%.02fmm  y=%.02fmm \n"

func interact(questions chan polar, answers chan cartesian) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(prompt)
	for {
		fmt.Printf("Radius and angle: ")
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		var radius, phi float64
		if _, err := fmt.Sscanf(line, "%f %f", &radius, &phi); err != nil {
			fmt.Fprintln(os.Stderr, "Invalid input")
			continue
		}
		questions <- polar{radius, phi}
		coord := <-answers
		fmt.Printf(result, radius, phi, coord.x, coord.y)
	}
	fmt.Println()
}

func main() {
	questions := make(chan polar)
	defer close(questions)
	answers := createSolver(questions)
	defer close(answers)
	interact(questions, answers)
}
