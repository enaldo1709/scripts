package fileutils

import (
	"math"
	"slices"
	"strings"
)

var allowedCharacters = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h',
	'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q',
	'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func LooksLike(s1, s2 string) bool {
	n1 := normalizeString(strings.ToLower(s1))
	n2 := normalizeString(strings.ToLower(s2))

	lengthDifference := math.Abs(float64(len(n1) - len(n2)))
	totalRunes := math.Max(float64(len(n1)), float64(len(n2)))

	coincidenceRunes := 0
	ns2 := n2
	for _, r := range n1 {
		if strings.ContainsRune(ns2, r) {
			coincidenceRunes++
			ns2 = strings.Replace(ns2, string(r), "", 1)
		}
	}

	return n1 == n2 || (lengthDifference < 3 &&
		(math.Abs(totalRunes-float64(coincidenceRunes))/totalRunes) < 0.15)
}

func normalizeString(s string) string {
	normalizedArr := []rune{}
	for _, r := range s {
		if !slices.Contains(allowedCharacters, r) {
			continue
		}
		normalizedArr = append(normalizedArr, r)
	}
	return string(normalizedArr)
}

func GetNameAndExtension(f string) (string, string) {
	lastDot := strings.LastIndex(f, ".")
	return f[:lastDot], f[lastDot+1:]
}
