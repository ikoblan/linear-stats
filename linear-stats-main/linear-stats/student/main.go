package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Использование: go run your-program.go <файл>")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var y []float64
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		val, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			log.Fatal(err)
		}
		y = append(y, val)
	}

	if len(y) == 0 {
		log.Fatal("Файл пуст")
	}

	// x = 0,1,2...
	n := float64(len(y))
	x := make([]float64, len(y))
	for i := range x {
		x[i] = float64(i)
	}

	// Считаем суммы
	var sumX, sumY, sumXY, sumX2, sumY2 float64

	for i := 0; i < len(y); i++ {
		sumX += x[i]
		sumY += y[i]
		sumXY += x[i] * y[i]
		sumX2 += x[i] * x[i]
		sumY2 += y[i] * y[i]
	}

	// Линейная регрессия (y = mx + b)
	a := (n*sumXY - sumX*sumY) / (n*sumX2 - sumX*sumX)
	b := (sumY - a*sumX) / n

	// Корреляция Пирсона
	r := (n*sumXY - sumX*sumY) /
		math.Sqrt((n*sumX2 - sumX*sumX) * (n*sumY2 - sumY*sumY))

	// Вывод
	fmt.Printf("Linear Regression Line: y = %.6fx + %.6f\n", a, b)
	fmt.Printf("Pearson Correlation Coefficient: %.10f\n", r)
}
