package x

//go:generate go run ../cmd/qexp -lazy -outdir ../pkg ./testdeps
//go:generate go run ../cmd/qexp -lazy -outdir ../pkg ./testlog
//go:generate go run ../cmd/qexp -outdir ../pkg ./race
//go:generate go run ../cmd/qexp -outdir ../pkg -addtags "//go:build !llgo" ./abi
//go:generate go run ../cmd/qexp -outdir ../pkg -tags "llgo" -addtags "//go:build llgo" -filename "export_llgo" ./abi
//go:generate go run ../cmd/qexp -outdir ../pkg ./goarch
//go:generate go run ../cmd/qexp -outdir ../pkg ./stringslite
//go:generate go run ../cmd/qexp -outdir ../pkg -src ./isync
