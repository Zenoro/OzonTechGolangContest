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
	"os"
	"strconv"
	"strings"
)

const MOD = 1_000_000_007

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

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
		fmt.Fprint(out, countArrays(n, aArray, bArray), "\n")
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

func countArrays(n int, l []int, r []int) int64 {
	var totalCount int64 = 1

	for i := 1; i <= n; i++ {
		li := l[i-1]
		ri := r[i-1]

		start := (li + i - 1) / i * i
		end := ri / i * i

		if start > end {
			return 0
		}

		count := (end-start)/i + 1
		totalCount = (totalCount * int64(count)) % MOD
	}
	return totalCount
}
