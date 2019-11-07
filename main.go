package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result struct {
	strikes int
	balls   int
}

func main() {
	rand.Seed(time.Now().UnixNano())
	//Make 3 random numbers
	numbers := MakeNumbers()

	cnt := 0
	for {
		cnt++
		//Get input from keyboard
		inputNumbers := InputNumbers()

		//compare the numbers
		result := CompareNumbers(numbers, inputNumbers)

		PrintResult(result)

		//is it 3 strikes?
		if IsGameEnd(result) {
			// game over
			break
		}
	}

	//game over, how many time did user guess?
	fmt.Printf("You got the answer in %d times.\n", cnt)

}

func MakeNumbers() [3]int {

	//return any 3 numbers that are from 0-9
	var rst [3]int

	for i := 0; i < 3; i++ {
		for {
			n := rand.Intn(10)
			duplicated := false
			for j := 0; j < i; j++ {
				if rst[j] == n {
					//if the number is the same as before, then get a new random number
					duplicated = true
					break
				}
			}
			if !duplicated {
				rst[i] = n
				break
			}
		}
	}

	//fmt.Println(rst)
	return rst
}

func InputNumbers() [3]int {
	//키보드로부터 0-9사이의 겹치지 않는 입력을 받아 반환한다.

	var rst [3]int

	for {
		fmt.Println("Choose 3 different numbers between 0~9.")
		var no int
		_, err := fmt.Scanf("%d\n", &no)
		if err != nil {
			fmt.Println("Something is wrong.")
			continue
		}

		success := true
		idx := 0
		for no > 0 {
			n := no % 10
			no = no / 10

			duplicated := false
			for j := 0; j < idx; j++ {
				if rst[j] == n {
					duplicated = true
					break
				}
			}
			if duplicated {
				fmt.Println("Each numbers have to be all different.")
				success = false
				break
			}
			if idx >= 3 {
				fmt.Println("You insert more than 3 numbers.")
				success = false
				break
			}

			rst[idx] = n
			idx++
		}

		if success && idx < 3 {
			fmt.Println("Choose 3 numbers.")
			success = false
		}

		if !success {
			continue
		}
		break
	}

	rst[0], rst[2] = rst[2], rst[0]
	//fmt.Println(rst)
	return rst
}

func CompareNumbers(numbers, inputNumbers [3]int) Result {
	//두개의 숫자 3개를 비교해서 결과를 반환한다.
	strikes := 0
	balls := 0

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if numbers[i] == inputNumbers[j] {
				if i == j {
					strikes++
				} else {
					balls++
				}
				break
			}
		}
	}

	return Result{strikes, balls}
}

func PrintResult(result Result) {
	fmt.Printf("%dS%dB\n", result.strikes, result.balls)
}

func IsGameEnd(result Result) bool {
	//비교 결과가 3스트라이크인지 확인
	return result.strikes == 3
}
