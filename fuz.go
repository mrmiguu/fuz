package fuz

func main() {
	fuz.Has("AlexIs2Cool", "is") // true
	fuz.Has("AlexIs2Cool", "xis") // false
	fuz.Has("Alexis2Cool", "xis") // true
	fuz.Has("Alexis2Cool", "oo") // false
}

func Sub(sub, s string) bool{
	if len(sub) == 0 {
		return -1, errors.New("empty substring")
	}
	expr := `(?i)[^` + string(sub[0]) + `]*`
	for x, r := range sub {
		end := `.*`
		if xplus1 := x + 1; xplus1 < len(sub) {
			end = `[^` + string(sub[xplus1]) + `]*`
		}
		expr += string(r) + end
	}

	n := -1
	matches := []string{""}
	for i, str := range strs {
		l, L := len(str), len(matches[0])
		if l == 0 || !regexp.MustCompile(expr).MatchString(str) {
			continue
		} else if 0 < L && L < l {
			continue
		}

		if l == L {
			matches = append(matches, str)
		} else {
			matches = []string{str}
			n = i
		}
	}

	if len(matches[0]) == 0 {
		return -1, errors.New("no match")
	} else if len(matches) > 1 {
		return -1, errors.New(sub + "=" + fmt.Sprint(matches) + " (too many matches)")
	}
	return n, nil
}

// func abbrev(sub string, strs ...string) (int, error) {
// 	if len(sub) == 0 {
// 		return -1, errors.New("empty substring")
// 	}
// 	expr := `(?i)[^` + string(sub[0]) + `]*`
// 	for x, r := range sub {
// 		end := `.*`
// 		if xplus1 := x + 1; xplus1 < len(sub) {
// 			end = `[^` + string(sub[xplus1]) + `]*`
// 		}
// 		expr += string(r) + end
// 	}

// 	n := -1
// 	matches := []string{""}
// 	for i, str := range strs {
// 		l, L := len(str), len(matches[0])
// 		if l == 0 || !regexp.MustCompile(expr).MatchString(str) {
// 			continue
// 		} else if 0 < L && L < l {
// 			continue
// 		}

// 		if l == L {
// 			matches = append(matches, str)
// 		} else {
// 			matches = []string{str}
// 			n = i
// 		}
// 	}

// 	if len(matches[0]) == 0 {
// 		return -1, errors.New("no match")
// 	} else if len(matches) > 1 {
// 		return -1, errors.New(sub + "=" + fmt.Sprint(matches) + " (too many matches)")
// 	}
// 	return n, nil
// }