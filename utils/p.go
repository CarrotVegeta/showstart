package utils

import "math"

const str = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateUrlApi(userID, eventID, apiID int) string {
	i := "000000" + generateIndexOfChar(userID)
	i = i[len(i)-6:]
	A := "0000" + generateIndexOfChar(eventID)
	A = A[len(A)-4:]
	P := "00" + generateIndexOfChar(apiID)
	P = P[len(P)-2:]
	e := P + A + i
	return e
}
func generateIndexOfChar(k int) string {
	var n string
	r := make(map[int]string)
	i := make(map[string]int)
	for t := 0; t < len(str); t++ {
		n = string(str[t])
		r[t] = n
		i[n] = t
	}
	return encode(k, r)
}
func encode(e int, n map[int]string) string {
	if e == 0 {
		return n[0]
	}
	var r string
	i := len(str)
	for e > 0 {
		r = n[e%i] + r
		e = int(math.Floor(float64(e / i)))
	}
	return r
}
