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
	a := `RdKnoWKXGxZP0zYw1ohcYoD+r1dSc1H64zrbklj5gRibwQgXYgZwzLh/Jr700HO92htHb6maa6SvKmCFLYmRYGpN1+mP+x3pzJOn+9secCttFaob9ulQTqhf5bC2lIDkRMp+JNzSgy96nvANad/flZi5p5dkzoLF6o9Sreqh669kdMJMGsYyFO2QHpThZGTvAsKyzLMSGcOXJTCOU0m8LT04xQLN4mjI/5JIHeBwji6ximBdFj8SBQHQKn26+4UdATju0SJr6dHKfrahZ5OOwYrBJWjCRtGpI3wkTG0/VsxwPJNG0LSa3yaYCcr/IcPmjhXd4y89gfqNLTNaeRvcOClAa5z7T9zTA5fn817oGvUthZVNt9NSC3ipPC2m/4M9VN17Ybtmx8iYNIPZ8AK2sFntezdHk/UCAmOhHYXJ/Tc2s9zxY3wBdhm+gnjsmckR`
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
