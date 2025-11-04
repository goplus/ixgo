package x

//go:generate go run ../cmd/qexp -outdir ../pkg ./testdeps
//go:generate go run ../cmd/qexp -outdir ../pkg ./testlog
//go:generate go run ../cmd/qexp -outdir ../pkg ./race
//go:generate go run ../cmd/qexp -outdir ../pkg -src ./abi
