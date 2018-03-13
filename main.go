package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	min int = 1
	max int = 10
	maxScore int = 10
)

func main(){
	var userScore int = 0
	var pcScore int =0

	for userScore < maxScore && pcScore < maxScore{
		userNumber := getUserNumber()
		calculatedNumber := getCalculateNumber()

		if userNumber == calculatedNumber {
			fmt.Printf("El numero %d es correcto\n", userNumber)
			userNumber++
		}else{
			fmt.Printf("El numero %d es incorrecto, el numero correcto era %d\n", userNumber, calculatedNumber)
			pcScore++
		}

		fmt.Printf("Marcador actual:\n\tUsuario: %d\n\tPC: %d\n", userScore, pcScore)
		time.Sleep(time.Second)
	}
}

func getUserNumber() int {
	var userNumber int = 0
	fmt.Printf("Escribe un numero entero entre %d y %d: ", min, max)
	fmt.Scanf("%d", &userNumber)	
	return userNumber
}

func getCalculateNumber() int {
	rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}
