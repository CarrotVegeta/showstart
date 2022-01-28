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
	a := `XZ7oPYA9M5X1aWyQIUsydkqPZVvd/2Md/F3AZz7o68PhyUvDOz50wTm3lqMTrx8S28qEvTwmvtgu59UWWax5tbabB2ugevjQJQ7V3h+ZF1nhPKEBawHCVdbzCGgF+s0Tp4BRw+g4ANN+gOpgPKPMn+4M07JcMS5urgfrWOYylFk3oaDIBRaQR6U48tGAM2e1pSQKx5nBUuMGQ86xJ7zsdLONF57WPxPjtNROD7UF0St3i75LM4ls8fQjyZQ2+rQpRvvCn2IrWBT3Yb3ddmpT2cmXKC2bqyrWK4nh5VFyvY+6ynQV6IArqNfrYlODAXENYhdHZYt0aywKOv1jmkpozwhVVCNGZhWb7dIU2pfvdrybTY7GXdY4NUQy720109ngA3lmuHEZWVq1VxV1z+d+SLA+1w1gNKydaDNfIUYkOgXO66t+SdgLCzc3X7Lr3RUVdRmx/1IITyHDeVr0bhlBC6phR48FUkCjmfWr9OGa4BD+hEsIImLzznWv7kxz0E/ckFczJCkRIA7MVE0+mvNePPkZ2YsI6o62JTRw+EyVxdS+w8HFvvoE085ALFFIlInTUhA1YryLy/wqei/u2mYqKfpB43Rq6DipiZWHDuPxYRyhyplP8U0Q7c29VRbTm9X2XB4Sy6jKaxf54+GeUOo470VcxwqhjbJ+htJ4hgCCC+N+X+8lerzgPT+C72oBHgI2AgCzxdD0vjgkGPnoCMyTz5Q4Y5NsGqSDRMip++lHdN5chjo4KLO1tvHVuQbKj6JZVgPzSi0Vp4wFmgp1+wVk4nIdSMK56/w1WLHktO9HYw+pb8hg+3jHqDyf4jLSx2La+NGO6QZ/MrwUXSRBUKxSmg==`
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
