package main

// https://github.com/andcloudio/go-concurrency-exercises
// “{ [] ( ) }” should return 1 (Openers and closers should be correctly paired)
//“{(([()] ()))}” should return 1 (Nesting should be supported)
//“some {text()[]} ok” should return 1 (Text and spaces should be ignored)
//“({()})” should return 1 (There isn’t hierarchy between symbols)
//“{ [(] ) }” should return 0 (Symbols should be correctly nested)
//“{ [ }” should return 0 (Symbols should allways close)

/*func main() {

	string1 := "{ [( ] ) }"
	runeJavascript := []rune(strings.ReplaceAll(string1, " ", ""))
	fmt.Println(validateCharacter(runeJavascript))
}

func validateCharacter(runeJavascript []rune) error {
	for i, c := range runeJavascript {

		// validamos que el string este lleno
		if i+1 > len(runeJavascript) {
			break
		}

		// obtenemos el cierre del caracter
		closeCharacter := getInitialANdLastCharacter(string(c))

		// validamos que el caracter actual sea un { o [ o (
		if closeCharacter == "" {
			continue
		}

		arrayValidator := runeJavascript[i+1:]
		foundPosition, err := verifyExistsCharacterClose(arrayValidator, closeCharacter)

		if err != nil {
			return err
		}
		runeJavascript = append(runeJavascript[:foundPosition], runeJavascript[foundPosition+1:]...)
	}

	return nil
}

func verifyExistsCharacterClose(arrayValidator []rune, closeCharacter string) (int, error) {
	for j, rr := range arrayValidator {
		if string(rr) == closeCharacter {
			fmt.Println("encontro", string(arrayValidator), closeCharacter)
			// eliminamos la posicion del array principal
			return j + 1, nil
		}
	}

	return 0, errors.New("error no found close character")
}

// validamos dado el caracter
func getInitialANdLastCharacter(character string) string {
	switch character {
	case "{":
		return "}"
	case "(":
		return ")"
	case "[":
		return "]"
	default:
		return ""
	}
}*/
