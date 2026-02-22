package runtime

import "runtime"

func AddCleanup[T, S any](ptr *T, cleanup func(S), arg S) (r runtime.Cleanup) {
	println(`ixgo warning: runtime.AddCleanup not impl. need build -tags=linknamefix -ldflags="-checklinkname=0"`)
	return
}
