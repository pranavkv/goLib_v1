module github.com/pranavkv/golib_v1/BaseHandler

go 1.18

replace github.com/pranavkv/golib_v1/libError => ../libError

replace github.com/pranavkv/golib_v1/LibUtils => ../LibUtils

replace github.com/pranavkv/golib_v1/LibData => ../LibData

require (
	github.com/pranavkv/golib_v1/LibUtils v0.0.0-00010101000000-000000000000
	github.com/pranavkv/golib_v1/libError v0.0.0-00010101000000-000000000000
	github.com/pranavkv/golib_v1/LibData v0.0.0-00010101000000-000000000000
)

require (
	github.com/sirupsen/logrus v1.9.0 // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
)


