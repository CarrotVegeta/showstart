package utils

import (
	"encoding/base64"
	"fmt"
	"log"
	"testing"
)

//测试AES ECB 加密解密
func TestEncryptDecrypt(t *testing.T) {
	s := `{"action":"/wap/common/info","method":"GET","query":{"st_flpv":"Uri4jL8p8d4Ed1t55MI9","sign":"991022effe9469c62beb5a3161712899","trackPath":"","terminal":"wap"},"body":null,"qtime":1642734525057,"ranstr":"64Cm5eMP"}`
	b := `0RGF99CtUajPF0Ny`
	c, err := AesECBEncrypt([]byte(s), []byte(b))
	if err != nil {
		log.Fatal(err.Error())
	}
	a := base64.StdEncoding.EncodeToString(c)
	d, err := base64.StdEncoding.DecodeString(a)
	if err != nil {
		log.Fatal(err.Error())
	}
	e, err := AesECBDecrypt(d, []byte(b))
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(e))
}
func TestAesECBDecrypt(t *testing.T) {
	a := `XZ7oPYA9M5X1aWyQIUsydkqPZVvd/2Md/F3AZz7o68PhyUvDOz50wTm3lqMTrx8S28qEvTwmvtgu59UWWax5tbabB2ugevjQJQ7V3h+ZF1l+lrXZlfMMRX8tyo6RLNLtp4BRw+g4ANN+gOpgPKPMn/V2oj3pI2xiKg9WAZQCUu2zHhr+mvWZ0bORSN/FRq3KeVcrEH+4FeIj5c4tmYyaVLRwjB9zAwGQlUCfUgVxZXV3i75LM4ls8fQjyZQ2+rQpvaU30zLz1JSQLkg90s2Ac/CSgyQyrH37tXPrYAYB+e8DcqmejXAEbs1DPYDHbgPlz/tD4zNU9tPT132w1fc460btTPDnPEelfecCJghaWvoMNZheJXRAYnXI9Wlms44+DDgidihW/EWFEVVB00hyyGOR7ROSKD+eSUZtbzwObvcoV2aVHUpK9yKlE+ktU5HtgU4CnUBNHU+9WJx7Qi5CrHCz7FQdys2+8LGFqMWYnHKy7dXO4QLouIJWAIwWodMKP3+OAFFCLJR0TvbnFUS9Q8WyExUOlG8tiI+KJ92LmS6B0bFnBRyxzdEW8r/C21NHjV+9x1+G0lqrQOFOwRXlqlSXyDrfYumUpM2IKLoLDlDcThTG4pOJbo3+p9ozX6IXwIbjDFD4BVahmiSQtMVXWsxrEkE/JaBdQNg20/iNZqR/KxBN6hX7QHaCtOMU1dUK6Kfg2r/hUpnXDn+U+pjT5+TmRX+XDck3dzpa+JyYrVDf0cP0zsAWJKpBOMdRnUqpR5PzvG8YmVZuq/BXeV+R0RcbgVbIEQlVyRXD0cKcOVKwo2AK3+ZjXx2/aoUD/Vax`
	b := `0RGF99CtUajPF0Ny`
	c, err := base64.StdEncoding.DecodeString(a)
	if err != nil {
		log.Fatal(err.Error())
	}
	d, err := AesECBDecrypt(c, []byte(b))
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(d))
}
