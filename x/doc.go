package x

//go:generate go run ../cmd/qexp -lazy -outdir ../pkg ./testdeps
//go:generate go run ../cmd/qexp -lazy -outdir ../pkg ./testlog
//go:generate go run ../cmd/qexp -lazy -outdir ../pkg ./race
//go:generate go run ../cmd/qexp -lazy -outdir ../pkg ./abi
//go:generate go run ../cmd/qexp -lazy -outdir ../pkg ./goarch
//go:generate go run ../cmd/qexp -lazy -outdir ../pkg ./stringslite
//go:generate go run ../cmd/qexp -lazy -outdir ../pkg -src ./isync
