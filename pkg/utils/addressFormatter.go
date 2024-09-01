package utils

import "strings"

func FormatCard(card string) string {
	card = strings.ReplaceAll(card, " ", "")
	card = strings.ReplaceAll(card, "-", "")
	card = strings.Trim(card, " ")
	return card
}

func FormatQiwi(card string) string {
	card = strings.ReplaceAll(card, "(", "")
	card = strings.ReplaceAll(card, ")", "")
	card = strings.ReplaceAll(card, "+", "")
	card = strings.ReplaceAll(card, " ", "")
	card = strings.Trim(card, " ")
	card = strings.ReplaceAll(card, "-", "")
	card = strings.Trim(card, " ")
	return card
}
