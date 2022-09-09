package main

import (
	"fmt"
)

// https://github.com/andcloudio/go-concurrency-exercises
// “{ [] ( ) }” should return 1 (Openers and closers should be correctly paired)
//“{(([()] ()))}” should return 1 (Nesting should be supported)
//“some {text() is[]} ok” should return 1 (Text and spaces should be ignored)
//“({()})” should return 1 (There isn’t hierarchy between symbols)
//“{ [(] ) }” should return 0 (Symbols should be correctly nested)
//“{ [ }” should return 0 (Symbols should allways close)

var s = []string{
	"{ [] ( ) }",
	"{(([()] ()))}",
	"some {text() is[]} ok",
	"({()})",
	"{ [(] ) }",
	"{ [ }",
}

func main() {
	fmt.Println(s)

	fmt.Printf("%s", "Guillermo")
	// https://www.educative.io/answers/the-valid-parentheses-problem
	validateChars := []string{
		//"{ [] ( ) }",
		"{(([()] ()))}",
		//"some {text() is[]} ok",
		//"({()})",
		//"{ [(] ) }",
		//"{ [ }",
	}

	for _, vChars := range validateChars {
		charactersValidator := sanitize([]rune(vChars))
		lenCharacter := len(charactersValidator)
		fmt.Println(validateChar(charactersValidator, lenCharacter))
	}
}

func validateChar(charactersValidator []rune, lenCharacter int) string {
	for i, r := range charactersValidator {
		closeCharacter := getInitialANdLastCharacter(r)

		if closeCharacter == "" {
			continue
		}

		if string(charactersValidator[lenCharacter-(i+1)]) != closeCharacter && string(charactersValidator[i+1]) != closeCharacter {
			fmt.Println(string(r), string(charactersValidator[lenCharacter-(i+1)]), string(charactersValidator[i+1]))
			return "error validation"
		}
	}

	return "ok"
}

// delete character number, string and spacing
func sanitize(char []rune) []rune {
	var newRune []rune
	// convert char to rune
	for _, c := range char {
		if c == '{' || c == '}' || c == '(' || c == ')' || c == '[' || c == ']' {
			newRune = append(newRune, c)
		}
	}
	return newRune
}

// validate character is open
func getInitialANdLastCharacter(character rune) string {
	switch string(character) {
	case "{":
		return "}"
	case "(":
		return ")"
	case "[":
		return "]"
	default:
		return ""
	}
}
