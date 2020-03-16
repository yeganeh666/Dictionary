package main

import (
	"errors"
	"fmt"


	//	"github.com/inancgumus/screen"
	//"os"
	//"os/exec"
	"sort"
	"strings"
)


//Menue ,print menue!
func Menue() {

//	screen.Clear()

	fmt.Println("\n\t  ** MENUE **")
	fmt.Println("   _______________________________")
	fmt.Println("\t1: Insert A New Word ")
	fmt.Println("   _______________________________")
	fmt.Println("\t2: Find A Word ")
	fmt.Println("   _______________________________")
	fmt.Println("\t3: Remove A Word  ")
	fmt.Println("   _______________________________")
	fmt.Println("\t4: Show List  ")
	fmt.Println("   _______________________________")
	fmt.Println("\t5: Exit  ")
	fmt.Println("   _______________________________")
}
func main() {


	Dic:=make(map[string]string)

	for  {

		var n string

		Menue()
		
		fmt.Scanln(&n)

		switch n {

		case "1":
			Insert(Dic)
			break

		case "2":
			var word string
			fmt.Scan(&word)
			Find(word,Dic)
			Find2(word,Dic)

		case "3":
			var word string
			fmt.Scan(&word)
			delete(Dic, word)

		case "4":
			Show(Dic)

		case "5":
			break
		}
	}
}

//Find ,find word in dic!
func Find(w string,Dic map[string]string) error {

	_, ok := Dic[w]
	if ok {
		return nil
	} else {
		return errors.New("NOT FOUND IN WORDS")
	}
}

//Insert ,insert new word or definition!
func Insert(Dic map[string]string) {

	var word string
	var def string

	fmt.Println("new word?")
	fmt.Scan(&word)

	fmt.Println("definition(s)?")
	fmt.Scan(&def)

	if err := Find(word,Dic); err != nil {
		Dic[word] = def
		fmt.Println("\n Added")
	} else {
		//Dic[word] = append(Dic[word] ,def)
		Dic[word] += def
		fmt.Println("\n Updated")
	}
}
func Find2(value string,Dic map[string]string) error {

	var err error
	for key, v := range Dic {
		if strings.ContainsAny(Dic[key], value) == true {
			fmt.Println(v)
			err = nil
		}
	}
	return err
}

//Show ,show dic!
func Show(Dic map[string]string) {

	//sort
	keys := make([]string, 0, len(Dic))
	for k := range Dic {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	//print
	for _, k := range keys {
		fmt.Println(k, Dic[k],",")
	}
}
