module github.com/ktmdan/offheap

go 1.16

replace github.com/glycerine/gommap => github.com/tysonmote/gommap v0.0.2

replace launchpad.net/gocheck => gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c

replace github.com/OneOfOne/xxhash => github.com/ktmdan/xxhash-64 v0.0.0-20230502181827-2feccca7449e

replace github.com/OneOfOne/xxhash/native => github.com/ktmdan/xxhash-64/native v0.0.0-20230502181827-2feccca7449e

require (
	github.com/glycerine/go-capnproto v0.0.0-20190118050403-2d07de3aa7fc
	github.com/glycerine/goconvey v0.0.0-20190410193231-58a59202ab31
	github.com/glycerine/rbtree v0.0.0-20190406191118-ceb71889d809 // indirect
	github.com/gopherjs/gopherjs v1.17.2 // indirect
	github.com/ktmdan/xxhash-64 v0.0.0-20230502181827-2feccca7449e
	github.com/shurcooL/go-goon v1.0.0
	github.com/tinylib/msgp v1.1.8
	github.com/tysonmote/gommap v0.0.2
)
