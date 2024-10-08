/*
Антон учит Степана делить массивы одинаковой длины n друг на друга следующим образом:
если a[i]​ нацело делится на b[i]​ для всех 1≤i≤n, то массив a делится на массив b.

В качестве упражнения, Степан всегда использует массив [1,2,…,n] в качестве делителя.
Теперь он хочет посчитать количество возможных массивов a, которые делятся на заданный делитель.
Но Антон догадался, что таких массивов бесконечно много, поэтому для Степана
он наложил на построение массива a следующие ограничение: l[i]​≤a[i]​≤r[i]​ для всех 1≤i≤n,
где l и r — подготовленные заранее Антоном массивы.

Но после введения таких ограничений, посчитать ответ оказалось слишком трудно, поэтому ребята просят вас им помочь.
*/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	DIFER := int64(math.Pow10(9) + 7)
	Tinput, _ := in.ReadString('\n')
	t, err := strconv.Atoi(strings.TrimSpace(Tinput))
	if err != nil {
		fmt.Fprint(out, "T AToI error")
		return
	}
	for tries := 0; tries < t; tries++ {
		Ninput, _ := in.ReadString('\n')
		n, err := strconv.Atoi(strings.TrimSpace(Ninput))
		if err != nil {
			fmt.Fprint(out, "n AToI error")
			return
		}
		aArray := ReadArray(in, n)
		bArray := ReadArray(in, n)
		var res int64 = 1
		for k := 0; k < n; k++ {
			res *= CountKrat(aArray[k], bArray[k], k+1)
			// res *= (bArray[k] - aArray[k] + 1) / (k + 1)
		}
		// fmt.Fprint(out, res%DIFER, "\n")
		fmt.Fprint(out, res%DIFER, "\n")
	}
}

func ReadArray(reader *bufio.Reader, length int) []int {
	input, _ := reader.ReadString('\n')
	strNumbers := strings.Split(strings.TrimSpace(input), " ")

	numbers := make([]int, length)
	for i, strNum := range strNumbers {
		num, _ := strconv.Atoi(strNum)
		numbers[i] = num
	}
	return numbers
}

func CountKrat(a int, b int, num int) int64 {
	var ctr int64 = 0
	for i := a; i <= b; i++ {
		if i%num == 0 {
			ctr++
		}
	}
	return ctr
}
