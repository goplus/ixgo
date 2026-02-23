// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26
// +build go1.26

package time

import (
	q "time"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "time",
		Path: "time",
		Deps: map[string]string{
			"errors":               "errors",
			"internal/bytealg":     "bytealg",
			"internal/godebug":     "godebug",
			"internal/stringslite": "stringslite",
			"math/bits":            "bits",
			"runtime":              "runtime",
			"sync":                 "sync",
			"syscall":              "syscall",
			"unsafe":               "unsafe",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Duration":   reflect.TypeOf((*q.Duration)(nil)).Elem(),
			"Location":   reflect.TypeOf((*q.Location)(nil)).Elem(),
			"Month":      reflect.TypeOf((*q.Month)(nil)).Elem(),
			"ParseError": reflect.TypeOf((*q.ParseError)(nil)).Elem(),
			"Ticker":     reflect.TypeOf((*q.Ticker)(nil)).Elem(),
			"Time":       reflect.TypeOf((*q.Time)(nil)).Elem(),
			"Timer":      reflect.TypeOf((*q.Timer)(nil)).Elem(),
			"Weekday":    reflect.TypeOf((*q.Weekday)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"Local": reflect.ValueOf(&q.Local),
			"UTC":   reflect.ValueOf(&q.UTC),
		},
		Funcs: map[string]reflect.Value{
			"After":                  reflect.ValueOf(q.After),
			"AfterFunc":              reflect.ValueOf(q.AfterFunc),
			"Date":                   reflect.ValueOf(q.Date),
			"FixedZone":              reflect.ValueOf(q.FixedZone),
			"LoadLocation":           reflect.ValueOf(q.LoadLocation),
			"LoadLocationFromTZData": reflect.ValueOf(q.LoadLocationFromTZData),
			"NewTicker":              reflect.ValueOf(q.NewTicker),
			"NewTimer":               reflect.ValueOf(q.NewTimer),
			"Now":                    reflect.ValueOf(q.Now),
			"Parse":                  reflect.ValueOf(q.Parse),
			"ParseDuration":          reflect.ValueOf(q.ParseDuration),
			"ParseInLocation":        reflect.ValueOf(q.ParseInLocation),
			"Since":                  reflect.ValueOf(q.Since),
			"Sleep":                  reflect.ValueOf(q.Sleep),
			"Tick":                   reflect.ValueOf(q.Tick),
			"Unix":                   reflect.ValueOf(q.Unix),
			"UnixMicro":              reflect.ValueOf(q.UnixMicro),
			"UnixMilli":              reflect.ValueOf(q.UnixMilli),
			"Until":                  reflect.ValueOf(q.Until),
		},
		TypedConsts: map[string]ixgo.TypedConst{
			"April":       {reflect.TypeOf(q.April), constant.MakeInt64(int64(q.April))},
			"August":      {reflect.TypeOf(q.August), constant.MakeInt64(int64(q.August))},
			"December":    {reflect.TypeOf(q.December), constant.MakeInt64(int64(q.December))},
			"February":    {reflect.TypeOf(q.February), constant.MakeInt64(int64(q.February))},
			"Friday":      {reflect.TypeOf(q.Friday), constant.MakeInt64(int64(q.Friday))},
			"Hour":        {reflect.TypeOf(q.Hour), constant.MakeInt64(int64(q.Hour))},
			"January":     {reflect.TypeOf(q.January), constant.MakeInt64(int64(q.January))},
			"July":        {reflect.TypeOf(q.July), constant.MakeInt64(int64(q.July))},
			"June":        {reflect.TypeOf(q.June), constant.MakeInt64(int64(q.June))},
			"March":       {reflect.TypeOf(q.March), constant.MakeInt64(int64(q.March))},
			"May":         {reflect.TypeOf(q.May), constant.MakeInt64(int64(q.May))},
			"Microsecond": {reflect.TypeOf(q.Microsecond), constant.MakeInt64(int64(q.Microsecond))},
			"Millisecond": {reflect.TypeOf(q.Millisecond), constant.MakeInt64(int64(q.Millisecond))},
			"Minute":      {reflect.TypeOf(q.Minute), constant.MakeInt64(int64(q.Minute))},
			"Monday":      {reflect.TypeOf(q.Monday), constant.MakeInt64(int64(q.Monday))},
			"Nanosecond":  {reflect.TypeOf(q.Nanosecond), constant.MakeInt64(int64(q.Nanosecond))},
			"November":    {reflect.TypeOf(q.November), constant.MakeInt64(int64(q.November))},
			"October":     {reflect.TypeOf(q.October), constant.MakeInt64(int64(q.October))},
			"Saturday":    {reflect.TypeOf(q.Saturday), constant.MakeInt64(int64(q.Saturday))},
			"Second":      {reflect.TypeOf(q.Second), constant.MakeInt64(int64(q.Second))},
			"September":   {reflect.TypeOf(q.September), constant.MakeInt64(int64(q.September))},
			"Sunday":      {reflect.TypeOf(q.Sunday), constant.MakeInt64(int64(q.Sunday))},
			"Thursday":    {reflect.TypeOf(q.Thursday), constant.MakeInt64(int64(q.Thursday))},
			"Tuesday":     {reflect.TypeOf(q.Tuesday), constant.MakeInt64(int64(q.Tuesday))},
			"Wednesday":   {reflect.TypeOf(q.Wednesday), constant.MakeInt64(int64(q.Wednesday))},
		},
		UntypedConsts: map[string]ixgo.UntypedConst{
			"ANSIC":       {"untyped string", constant.MakeString(string(q.ANSIC))},
			"DateOnly":    {"untyped string", constant.MakeString(string(q.DateOnly))},
			"DateTime":    {"untyped string", constant.MakeString(string(q.DateTime))},
			"Kitchen":     {"untyped string", constant.MakeString(string(q.Kitchen))},
			"Layout":      {"untyped string", constant.MakeString(string(q.Layout))},
			"RFC1123":     {"untyped string", constant.MakeString(string(q.RFC1123))},
			"RFC1123Z":    {"untyped string", constant.MakeString(string(q.RFC1123Z))},
			"RFC3339":     {"untyped string", constant.MakeString(string(q.RFC3339))},
			"RFC3339Nano": {"untyped string", constant.MakeString(string(q.RFC3339Nano))},
			"RFC822":      {"untyped string", constant.MakeString(string(q.RFC822))},
			"RFC822Z":     {"untyped string", constant.MakeString(string(q.RFC822Z))},
			"RFC850":      {"untyped string", constant.MakeString(string(q.RFC850))},
			"RubyDate":    {"untyped string", constant.MakeString(string(q.RubyDate))},
			"Stamp":       {"untyped string", constant.MakeString(string(q.Stamp))},
			"StampMicro":  {"untyped string", constant.MakeString(string(q.StampMicro))},
			"StampMilli":  {"untyped string", constant.MakeString(string(q.StampMilli))},
			"StampNano":   {"untyped string", constant.MakeString(string(q.StampNano))},
			"TimeOnly":    {"untyped string", constant.MakeString(string(q.TimeOnly))},
			"UnixDate":    {"untyped string", constant.MakeString(string(q.UnixDate))},
		},
	})
}
