package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

var Dic map[string]string

//Menue ,print menue!
func Menue() {

	fmt.Println("\n\t  ** MENUE **")
	fmt.Println("\n   _______________________________\n")
	fmt.Println("\t1: Insert A New Word \n")
	fmt.Println("\n   _______________________________\n")
	fmt.Println("\t2: Find A Word \n")
	fmt.Println("\n   _______________________________\n")
	fmt.Println("\t3: Remove A Word  \n")
	fmt.Println("\n   _______________________________\n")
	fmt.Println("\t4: Show List  \n")
	fmt.Println("\n   _______________________________\n")
	fmt.Println("\t5: Exit  \n")
	fmt.Println("\n   _______________________________\n")
}
func main() {

	var t bool

	for t {

		var n string

		Menue()

		fmt.Scan(&n)

		switch n {

		case "1":
			Insert()

		case "2":
			var word string
			fmt.Scan(&word)
			Find(word)
			Find2(word)

		case "3":
			var word string
			fmt.Scan(&word)
			delete(Dic, word)

		case "4":
			Show()

		case "5":
			t = false
		}
	}
}

//Find ,find word in dic!
func Find(w string) error {

	_, ok := Dic[w]
	if ok {
		return nil
	} else {
		return errors.New("NOT FOUND IN WORDS")
	}
}

//Insert ,insert new word or definition!
func Insert() {

	var word string
	var def string

	fmt.Println("new word?")
	fmt.Scan(&word)

	fmt.Println("definition(s)?")
	fmt.Scan(&def)

	if err := Find(word); err != nil {
		Dic[word] = def
		fmt.Println("\n Added")
	} else {
		Dic[word] = def
		fmt.Println("\n Updated")
	}
}
func Find2(value string) error {

	var err error
	for key, _ := range Dic {
		if strings.ContainsAny(Dic[key], value) == true {
			fmt.Println(key)
			err = nil
		}
	}
	return err
}
func Scond_case() {

}

//Show ,show dic!
func Show() {

	//sort
	keys := make([]string, 0, len(Dic))
	for k := range Dic {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	//print
	for _, k := range keys {
		fmt.Println(k, Dic[k])
	}
}
