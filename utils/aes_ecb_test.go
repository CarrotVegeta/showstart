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
	a := `RdKnoWKXGxZP0zYw1ohcYtCFoaCEsWtat9s7r3I+rGjhyUvDOz50wTm3lqMTrx8S0hL+6lTYOsuc9RU7P8zg9r5gtsrftqIo+2608gt9tZ5Gyf4KCCAdGGJ5oCdZKMJ85pfkl7Fv8+OYj5jrvVJlCJ5dN2lxPpxboJwm6f4oyHai/1MzhnPkSy0dieu3eWI+I8duMQ5OewR5B4+4ReY0gKOVTWzS+ucw9PMdA++GMNNVnSJbO5dGpiuL6RLWt30b+vJYgPj50hsyy+3uyxP7+1SapeqN2lV8E3DrPpHNp38=`
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
