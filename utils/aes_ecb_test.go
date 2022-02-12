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
	a := `XZ7oPYA9M5X1aWyQIUsydlM/KogoVPC2F4hclDEl6yOjnnt/0yiEBEHnRWpS722gqt9uGTC3xhuAeg359u0Xq84JigxGArlmESA2O2HAvayRPFdLiF48A7JvG3rG2AWrUGmWKvrmRb1Csi3tH1m/GZMtmxnqtaf2PM2NmDc+LjLE4UXwR80XwVGXA4sq9lYb4mOTpiuJzLh4xPXzQY/pq+fCaRPV/e+exue8qRkOwKHeznf/k0YWVA8Fs21D5tZzPBYw0EF+eRyGBaFaCvvyZYjE2vrn623EAgeUbHj+c0LUUUSlIRBFusgUuSXHZiFb6wUthdsyTUcx7h0X8SaaIP6LpzS6BeS/ZuNWVQJ5UT2shy4RR4wKMnY87ZEI2lWm`
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
