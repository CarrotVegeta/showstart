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
	a := `RdKnoWKXGxZP0zYw1ohcYgXTfjEC1Cncgujqi9fqCOyuXom9Na+R+iaGA+mDcMTpQPxvhilxXHezPX3QkvrkQxYe4j+mlMx7ATo9jWdTdel4CXS6F6PojNUzM6+jidIlrVUBTYyWyZBHKG1wWCgz7Ct2ouoIskswsZHQ/ImUS0dqDGYwfS4abrtIFbmTV/ii73iWYw+EOO0fwYYnjiGYciD0Y3J92luijenWEM7BcGpZOUxNcwfcnzi6VpLNXfF68p7sCMqF0ufbw9vXj8+rrRCYPYvwDupHrbP4nRl3gpgfpJrrsPgUqBsJdzrE8dkOxa3NAQk53I1szhb3yIjig3brCtN7nP0G96aRA8g/O/VrwXvchpR0AymIOP8VfOLU8es0eNoIKPoG95zUUEcIlfgWdzC7ukEmPwAoAudsCWCsVk2Jw/ifPGHXEh5q3LGMH7J/5yV0u3HVixy7qnc3N8QI4VUTXpWk/VZ4oxeC/QOR9hUwDFrtQdHnd6d2gJZyaXF2Kb1DNB5SLIGLgNmlTYECZVFYM/C6Hf4Tl7LggQ3k5kV/lw3JN3c6WvicmK1Q39HD9M7AFiSqQTjHUZ1KqUeT87xvGJlWbqvwV3lfkdEayPA01Bf5SvGpM74Utfm3fMxL3DAgy+q9vYUiw2ZY8A==`
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
