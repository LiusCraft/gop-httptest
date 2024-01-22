module github.com/liuscraft/gop-httptest

go 1.18

require (
    github.com/qiniu/httptest v1.0.3
	github.com/qiniu/dyn v1.3.0
	github.com/qiniu/x v1.10.5
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
)

// replace github.com/qiniu/httptest => ../httptest
// replace github.com/qiniu/dyn => ../dyn