package service

//
//const loginMod = "login"
//
//type User struct {
//	AreaCode   string `json:"areaCode"`
//	Phone      string `json:"phone"`
//	VerifyCode string `json:"verifyCode"`
//	StFlpv     string `json:"st_flpv"`
//	FormToken  string `json:"formToken"`
//	Captcha    string `json:"captcha"`
//	Sign       string `json:"sign"`
//	TrackPath  string `json:"trackPath"`
//	Terminal   string `json:"terminal"`
//}
//
//func (u *User) Open(s *engine.Server) {
//	s.RegisterMod(u, loginMod)
//}
//func (u *User) Request() {
//	u.Login()
//}
//func (u *User) Login() {
//	m := service.Login()
//	log.Println(m)
//}
//func Login() map[string]interface{} {
//	stflpv := "T67bztUXga8VJRqrYQvQ"
//	lp := &models.LoginParam{
//		AreaCode:   "86_CN",
//		Phone:      "13118115155",
//		VerifyCode: "658922",
//		StFlpv:     stflpv,
//		FormToken:  stflpv,
//		Captcha:    "",
//		Sign:       "",
//		TrackPath:  "",
//		Terminal:   pkg.Terminal,
//	}
//	return RequestWithBodyParam(pkg.WapUserLogin, pkg.A, "POST", lp)
//}
