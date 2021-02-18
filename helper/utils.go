package helper

import (
    "regexp"
    "strconv"
)

func CheckPositiveDigit(str string) (int) {
	var digitRegexp = regexp.MustCompile(`^[0-9]+$`)

    if digitRegexp.MatchString(str) {
        var n, err = strconv.Atoi(str)

        if err == nil && n > 0 {
            return n
        }
    }

    return -1
}