package httptest

import "fmt"

var (
	GopPackage = true
)

type Apper interface {
	initApp()
}

type App struct {
	*Test
}

func (p *App) initApp() {
	p.Test = &Test{}
}

func Gopt_App_Main(app Apper) {
	app.initApp()
	app.(interface{ MainEntry() }).MainEntry()
}

func (p * App) get(url string) {
	p.Test.Get(url)
}

// func (p *App) post(url string) {
// 	p.Test.Post(url)
// }

// func (p * App) get__0() {
// 	p.Test.Get(p.Test.url)
// }

// func (p *App) post__0() {
// 	p.Test.Post(p.Test.url)
// }

func (p *App) url(url string) {
	p.Test.url = url
}

func (p *App) ret(code int) {
	if(p.Test.result.StatusCode == code) {
		fmt.Println("pass")
	} else {
		fmt.Println("fail")
	}
}

func (p *App) Run(name string) {
	fmt.Println(name)
}
