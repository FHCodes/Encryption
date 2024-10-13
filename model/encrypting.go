package model

func Encrypt(phrase *string) {

	var slice []rune

	for _, c := range *phrase {
		slice = append(slice, c+3)
	}
	*phrase = string(slice)
}
