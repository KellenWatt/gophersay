package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

var MAX = 40

func main() {
	in := bufio.NewReader(os.Stdin)
	maxlength := 0
	lines := make([]string, 1)
	text := ""
	if len(os.Args) == 1 || os.Args[1] == "-" {
		text,_ = in.ReadString('\n')
		if len(text) != 0 {
			text = text[:len(text)-1]
		}
	}else {
		for _,v := range os.Args[1:] {
			text += v + " "
		}
		text = text[:len(text)-1]
	}
	words := strings.Split(text, " ")
	for k,w := range words {
		if len(w) > MAX {
		s := make([]string, 0)
			for j:=0; j<(len(w)/MAX)+1; j++ {
				if j != (len(w)/MAX) {
					s = append(s, w[j*MAX : (j+1)*MAX])
				}else {
					s = append(s, w[j*MAX : len(w)])
				}
			}
		n := k
		if k == len(words)-1 {
			n++
		}
		w2 := words[n:]
		words = append(words[:k], s...)
		words = append(words, w2...)
		}
	}
	
	i := 0
	for _,w := range words {
		length := len(lines[i])
		if length <= MAX {
			lines[i] += (w + " ")
			length += len(w)+1
			if length >= MAX {
				lines = append(lines, "")
				i++
			}
			if length > maxlength {
				maxlength = length
			}
		}
	}
	fmt.Print(" ")
	for j:=0; j<maxlength+1; j++ {
		fmt.Print("_")
	}
	fmt.Println(" ")
	for _,l := range lines {
		fmt.Print("( ", l[:len(l)-1])
		for j:=len(l)-1; j<maxlength; j++ {
			fmt.Print(" ")
		}
		fmt.Println(")")
	}
	fmt.Print(" ")
	for j:=0; j<maxlength+1; j++ {
		if j == 1 {
			fmt.Print("\\")
		}else{
			fmt.Print("‾")
		}
	}
	fmt.Println()
	fmt.Println("   \\")
	fmt.Println("    \\ ___")
	fmt.Println("     (0_0)")
//	fmt.Println("     | \" |")
	fmt.Println("    (| \" |)")
	fmt.Println("     || ||")
	fmt.Println("     U---U")
}
