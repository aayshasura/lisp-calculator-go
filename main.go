package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Stack []string

func (st *Stack) IsEmpty() bool {
	return len(*st) == 0
}
func (st *Stack) Push(v string) {
	*st = append(*st, v)
}

func (st *Stack) Pop() string {
	if st.IsEmpty() {
		return ""
	}
	top := (*st)[len(*st)-1]
	*st = (*st)[:len(*st)-1]
	return top
}

func (st *Stack) Operate(op string) int {

	ans, _ := strconv.Atoi(st.Pop())
	n := len(*st)
	switch op {
	case "+":
		for i := 0; i < n; i++ {
			num, _ := strconv.Atoi(st.Pop())
			ans += num
		}
	case "-":
		for i := 0; i < n; i++ {
			num, _ := strconv.Atoi(st.Pop())
			ans -= num
		}
	case "*":
		for i := 0; i < n; i++ {
			num, _ := strconv.Atoi(st.Pop())
			ans *= num
		}
	case "/":
		num, _ := strconv.Atoi(st.Pop())
		ans /= num
	}
	return ans
}
func main() {
	expression := "( - 1 2 5 )"
	// st := Stack{}
	expressionArray := strings.Split(expression, " ")
	n := len(expressionArray)
	i := n - 1
	for {
		a := expressionArray[i]
		if a == ")" {
			j := i - 1
			tempStack := Stack{}
			for {
				b := expressionArray[j]
				j--
				if b == "(" {
					op := tempStack.Pop()
					for {
						if tempStack.IsEmpty() {
							break
						}
						fmt.Println(tempStack.Operate(op))
					}
					break
				} else {
					tempStack.Push(b)
				}
			}
			i = j
		}
		if i == -1 {
			break
		}

	}

}
