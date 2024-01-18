package httptest

import (
	"fmt"
	"strings"
	"strconv"
)

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
	// p.Test = new(Test)
}

func Gopt_App_Main(app Apper) {
	app.initApp()
	app.(interface{ MainEntry() }).MainEntry()
}

func (p * App) Get__0(url string) {
	p.Test.Get(url)
}

func (p * App) Get__1() {
	p.Test.Get(p.Test.gUrl)
}

func (p *App) Post__0(url string, body map[string][]string) {
	p.Test.Post(url, body)
}

func (p *App) Post__1() {
	p.Test.Post(p.Test.gUrl, p.Test.body)
}

func (p *App) Post__2(url string) {
	p.Test.Post(url, p.Test.body)
}

func (p *App) Url(url string) {
	p.Test.gUrl = url
}

func (p *App) Body(body map[string][]string) {
	p.Test.body = body
}

func (p *App) CaseName(caseName string) {
	p.Test.caseName = caseName
}

func (p *App) Ret(code int) {
	if(p.Test.result.StatusCode == code) {
		fmt.Println(p.Test.caseName + ": pass")
	} else {
		fmt.Println(p.Test.caseName + ": fail")
	}
}

func (p *App) Match__0(src, tgt string) {
	if(strings.Compare(src, tgt) == 0) {
		fmt.Println(src + " " + tgt + ": matched")
	} else {
		fmt.Println(src + " " + tgt + ": not matched")
	}
}

func (p *App) Match__1(src, tgt int) {
	if(src == tgt) {
		fmt.Println(strconv.Itoa(src) + " " + strconv.Itoa(tgt) + ": matched")
	} else {
		fmt.Println(strconv.Itoa(src) + " " + strconv.Itoa(tgt) + ": not matched")
	}
}

// func (p *App) Run(name string) {
// 	fmt.Println(name)
// }
