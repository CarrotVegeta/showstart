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
	a := `XZ7oPYA9M5X1aWyQIUsydkqPZVvd/2Md/F3AZz7o68PhyUvDOz50wTm3lqMTrx8S28qEvTwmvtgu59UWWax5tbabB2ugevjQJQ7V3h+ZF1l+lrXZlfMMRX8tyo6RLNLtp4BRw+g4ANN+gOpgPKPMn12AUxtBXk+/eTxaMJ9VOZBESmZWvorE9Ev8FYuBs5uXLXxT8ZsrMxRjKHmnvrtmfSB/tYNp7lH9Oca6edg8mzJ3i75LM4ls8fQjyZQ2+rQpvaU30zLz1JSQLkg90s2Ac/CSgyQyrH37tXPrYAYB+e9/GfgtumrQlKiBuojD+xA2swdxa5mUko02U4AmGD7caEbtTPDnPEelfecCJghaWvoMNZheJXRAYnXI9Wlms44+DDgidihW/EWFEVVB00hyyGOR7ROSKD+eSUZtbzwObvcoV2aVHUpK9yKlE+ktU5HtgU4CnUBNHU+9WJx7Qi5CrHCz7FQdys2+8LGFqMWYnHKy7dXO4QLouIJWAIwWodMK4OJM7O8hGUhKEpLFOufN0cWyExUOlG8tiI+KJ92LmS7rFguUZGB/899VJBlC7SOVq47pGQSNCSCeXBlBZLJghVSXyDrfYumUpM2IKLoLDlDW9b0DLOPl+yW8cNbF8q8jKAvOot+AdyDHu7pCsvhx8dUQtZoGzcdqtqyJaaf8iTbNWG0sDOiNgmFLxj+CwlG5tmkRyzRLio1G5VDSV1sZpeTmRX+XDck3dzpa+JyYrVDf0cP0zsAWJKpBOMdRnUqpR5PzvG8YmVZuq/BXeV+R0Ubn7QdWevzVoQH+uKaCvvyBRGFF2L1LITjdZi2CsmNB`
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
