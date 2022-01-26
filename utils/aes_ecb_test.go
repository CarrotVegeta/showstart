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
	a := `XZ7oPYA9M5X1aWyQIUsydkqPZVvd/2Md/F3AZz7o68PhyUvDOz50wTm3lqMTrx8S28qEvTwmvtgu59UWWax5tbabB2ugevjQJQ7V3h+ZF1nwQZHMf53Dh8ufnKyb3e7Np4BRw+g4ANN+gOpgPKPMn1Bqf600j+P7sfG71I3/79DXuO2FSAvlGPIlM3ihugDd38CgthE1pj/qu7o4nM/eyDFsXR/l13q6oHxyN5TwVMZ3i75LM4ls8fQjyZQ2+rQpRvvCn2IrWBT3Yb3ddmpT2cmXKC2bqyrWK4nh5VFyvY+6ynQV6IArqNfrYlODAXENYhdHZYt0aywKOv1jmkpoz/UNeXE7p3fSZxvlRAlqmS7vDsy3ooZRvS0FDh/4NZZYGYd84mxMOMHtB7+n+3bMxCtHSdMWzO9KoYdit5KrWfyseUu0E8QXeJFJjgcs+SJ3/HpFgrDyDdwgsPIunvj994TZ5I8jCu26JEv/mJjm19lHjKBOJPN1RTZfM7x0kccL99qiUZ9Tfeyz/xHHxdUmy2+4l5BE1WbodqYu2PMvD+fUo8w+2XAw3VQXgkNtaxGQaWdVLh+tpP6TKpX3WFoLNHVaIBqWJWjslpkGL3QLLL8MVVuMaul+AcJLmpv1rueNi9JymF6SZsXUmL24Bxz3xrEXRN3px8OzZnIQEitozkZSU6aBmI9X7AwMAS8ZTStzq06Mbd559aGAeYCc6FIF1yHPH2Za0WHk6TmMvE2PJ93GZGIwOdaFMzQVG299y21MNzFQqeXgNT1WghVIbP0hRKKfQ1E3PJEwZcbpEtkw/lTVzI5SJRlwQ+Ie9MuPpo6E6RcrsQchzvzET3ts93sspL/ljyo+W6WbcWVYxLpPNtBpt7U7fT2FUS1pTCfUWVryqOX/36GItLFZqnAmHaNt1A==`
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
