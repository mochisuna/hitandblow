package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const NUM_COL = 6
const NUM_BALL = 4

func permutation(pat *[]string, index *int, s string) string {
	if len(s) >= NUM_BALL {
		(*pat)[*index] = s
		*index = *index + 1
		return s
	}
	for i := 0; i < NUM_COL; i++ {
		permutation(pat, index, fmt.Sprintf("%v%v", s, i))
	}
	return ""
}

func count(val string) []int {
	cnt := make([]int, NUM_COL)
	for _, p := range val {
		cnt[int(p-'0')] += 1
	}
	return cnt
}

func listup(pat *[]string, cnt *[][]int, candidate *[]bool, sel string, hit, blow int) {
	cnt_sel := count(sel)
	for i := 0; i < len(*cnt); i++ {
		ch_hit, ch_blow := 0, 0
		for j, s := range sel {
			if int(s) == int((*pat)[i][j]) {
				ch_hit++
			}
		}
		if hit != ch_hit {
			(*candidate)[i] = false
			continue
		}
		for j := 0; j < len(cnt_sel); j++ {
			for a, b := cnt_sel[j], (*cnt)[i][j]; ; a, b = a-1, b-1 {
				if a == 0 || b == 0 {
					break
				}
				ch_blow++
			}
		}
		if blow != ch_blow-hit {
			(*candidate)[i] = false
		}
	}
}

func main() {
	pat := make([]string, int(math.Pow(NUM_COL, NUM_BALL)))
	cnt := make([][]int, int(math.Pow(NUM_COL, NUM_BALL)))
	index := 0
	permutation(&pat, &index, "")
	index = 0
	for _, p := range pat {
		cnt[index] = count(p)
		index++
	}

	sc := bufio.NewScanner(os.Stdin)
	candidate := make([]bool, len(pat))
	for i := 0; i < len(candidate); i++ {
		candidate[i] = true
	}
	for {
		fmt.Println("number hit blow")
		sc.Scan()
		inputs := strings.Split(sc.Text(), " ")
		if len(inputs) != 3 {
			continue
		}

		hit, _ := strconv.Atoi(inputs[1])
		if hit == NUM_BALL {
			fmt.Printf("Answer is %v\n", inputs[0])
			return
		}
		blow, _ := strconv.Atoi(inputs[2])
		listup(&pat, &cnt, &candidate, inputs[0], hit, blow)
		fmt.Println("--- LIST ---")
		for i, f := range candidate {
			if f {
				fmt.Println(pat[i])
			}
		}
		fmt.Println("------------")
	}
}
