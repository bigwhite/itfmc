package main

import (
	"errors"
	"fmt"
	"os"
	"text/scanner"

	mystack "github.com/bigwhite/itfmc/stack"
)

// $find $GOROOT -name '*.go'|grep -v "_test.go" |xargs cat| grep -v ^$|grep -v "^[[:space:]]*\/\/"|itfmc 2>/dev/null

func nextIdentTokenText(s *scanner.Scanner) (string, error) {
	for {
		switch s.Scan() {
		case scanner.EOF:
			return "", errors.New("EOF")
		case scanner.Ident:
			return s.TokenText(), nil
		}
	}
}

func typeEnd(s *scanner.Scanner) error {
	stk := mystack.NewStack()
	var txt string

	for {
		switch s.Scan() {
		case scanner.EOF:
			return errors.New("EOF")
		default:
			txt = s.TokenText()
			if txt == "}" {
				stk.Pop() // pop "{"
				if stk.Empty() {
					return nil
				}
			}

			if txt == "{" {
				stk.Push(txt) //push "{"
			}
		}
	}
}

func main() {
	var a [100]int
	var m = make(map[string]int)
	var s scanner.Scanner
	s.Init(os.Stdin)
	for {
		switch s.Scan() {
		case scanner.EOF:
			total := 0
			for i := range a {
				if a[i] != 0 {
					fmt.Printf("a[%d] = %d\n", i, a[i])
					total += a[i]
				}
			}
			fmt.Printf("interfaces in total = %d\n", total)

			total = 0
			for range m {
				total++
			}
			fmt.Printf("interfaces in total sort by name = %d\n", total)
			return // all done
		case scanner.Ident:
			//fmt.Println(s.TokenText())
			tok := s.TokenText()
			if tok == "type" {
				start := s.Pos().Line
				//fmt.Println("found type: start line = ", start)
				name, err := nextIdentTokenText(&s) //skip interface name
				if err != nil {
					fmt.Println("EOF: return")
					return
				}

				t1, err := nextIdentTokenText(&s)
				if err != nil {
					fmt.Println("EOF: return")
					return
				}

				//fmt.Println("t1=", t1)
				if t1 == "interface" {
					err = typeEnd(&s)
					if err != nil {
						fmt.Println("EOF: return")
						return
					}

					// found
					end := s.Pos().Line
					//fmt.Println("found }: end line = ", end)
					c := end - start
					if c <= 0 || c >= 100 {
						break
					}
					a[c-1] += 1
					//fmt.Println("interface name is", name)
					m[name] = c - 1
					//fmt.Printf("===%s : %d\n", name, c-1)
				}
			}
		}
	}

}
