package httptest

var (
	GopPackage = true
)

type Apper interface {
	initApp()
}

type App struct {
	*Test
	*Cmd
}

func (p *App) initApp() {
	p.Cmd = &Cmd{}
}

func Gopt_App_Main(app Apper) {
	app.initApp()
	app.(interface{ MainEntry() }).MainEntry()
}

func (p *App) Assert(data bool) {
	p.addCode("assert")
	if data {
		p.addData("assert", "1")
	} else {
		p.addData("assert", "0")
	}
}

func (p *App) SubCmd(data string) string {
	return data
}

// func (p *App) GetVar(data string) {
// return dyn.Get(p.vars, key)
// }

func (p *App) Case(data ...string) {
	p.addCode("case")
	p.addData("case", data...)
}

func (p *App) Auth(data ...string) {
	p.addCode("auth")
	p.addData("auth", data...)
}

func (p *App) Match(data ...string) {
	p.addCode("match")
	p.addData("match", data...)
}

func (p *App) Get(data ...string) {
	p.addCode("get")
	p.addData("get", data...)
}

func (p *App) Post(data ...string) {
	p.addCode("post")
	p.addData("post", data...)
}

func (p *App) Ret(data ...string) {
	p.addCode("ret")
	p.addData("ret", data...)
}

func (p *App) Echo(data ...string) {
	p.addCode("echo")
	p.addData("echo", data...)
}

func (p *App) Run(stage string) {
	p.Cmd.Run(stage)
	p.Cmd.Code = make([]string, 0)
	p.Cmd.Data = make([][]string, 0)
}

func (p *App) addCode(code string) {
	p.Cmd.Code = append(p.Cmd.Code, code)
}

func (p *App) addData(code string, data ...string) {
	p.Cmd.Data = append(p.Cmd.Data, append([]string{code}, data...))
}
