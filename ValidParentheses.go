package main

import (
	"fmt"
)

func main() {
	s := "(){}])"
	valid := isValid(s)

	fmt.Println("Given String is:")
	fmt.Print(valid)
}

type ValidStack struct {
	vStack []string
}

func Constructor() ValidStack {
	return ValidStack{}
}

func (this *ValidStack) Pop() {
	if len(this.vStack) > 0 {
		this.vStack = this.vStack[0 : len(this.vStack)-1]
	}
}

func (this *ValidStack) Push(val string) {

	this.vStack = append(this.vStack, val)

}

func (this *ValidStack) Top() string {
	return this.vStack[len(this.vStack)-1]
}

func (this *ValidStack) bracketMap(b string) string {
	var bMap = make(map[string]string)
	bMap[")"] = "("
	bMap["}"] = "{"
	bMap["]"] = "["
	return bMap[b]
}

func isValid(s string) bool {

	//Create obj from Constructor
	obj := Constructor()

	if len(s) == 1 {
		return false
	}

	if len(s) == 0 {
		return true
	} else {
		i := 0
		for _, ch := range s {
			// Unicode 123, 125 is {}; 40,41 is (); 91,93 is []
			if ch == 123 || ch == 91 || ch == 40 {
				obj.Push(string(ch))
				i++
			} else {
				//encountered closing at first
				if len(obj.vStack) == 0 {
					return false
				}
				//Find pair
				braces := obj.bracketMap(string(ch))
				if len(obj.vStack) > 0 {
					if braces == obj.Top() {
						fmt.Println("Popped this guy:")
						fmt.Println(obj.Top())
						obj.Pop()
						i--
					} else {
						return false
					}
				}
			}
		}
		// if all pairs removed
		if len(obj.vStack) == 0 {
			return true
		}
		//fmt.Println(s)
		fmt.Println(obj.vStack)
	}
	return false
}
