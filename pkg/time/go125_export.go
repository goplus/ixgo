// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

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
			"April":       {Typ: reflect.TypeOf(q.April), Value: constant.MakeInt64(int64(q.April))},
			"August":      {Typ: reflect.TypeOf(q.August), Value: constant.MakeInt64(int64(q.August))},
			"December":    {Typ: reflect.TypeOf(q.December), Value: constant.MakeInt64(int64(q.December))},
			"February":    {Typ: reflect.TypeOf(q.February), Value: constant.MakeInt64(int64(q.February))},
			"Friday":      {Typ: reflect.TypeOf(q.Friday), Value: constant.MakeInt64(int64(q.Friday))},
			"Hour":        {Typ: reflect.TypeOf(q.Hour), Value: constant.MakeInt64(int64(q.Hour))},
			"January":     {Typ: reflect.TypeOf(q.January), Value: constant.MakeInt64(int64(q.January))},
			"July":        {Typ: reflect.TypeOf(q.July), Value: constant.MakeInt64(int64(q.July))},
			"June":        {Typ: reflect.TypeOf(q.June), Value: constant.MakeInt64(int64(q.June))},
			"March":       {Typ: reflect.TypeOf(q.March), Value: constant.MakeInt64(int64(q.March))},
			"May":         {Typ: reflect.TypeOf(q.May), Value: constant.MakeInt64(int64(q.May))},
			"Microsecond": {Typ: reflect.TypeOf(q.Microsecond), Value: constant.MakeInt64(int64(q.Microsecond))},
			"Millisecond": {Typ: reflect.TypeOf(q.Millisecond), Value: constant.MakeInt64(int64(q.Millisecond))},
			"Minute":      {Typ: reflect.TypeOf(q.Minute), Value: constant.MakeInt64(int64(q.Minute))},
			"Monday":      {Typ: reflect.TypeOf(q.Monday), Value: constant.MakeInt64(int64(q.Monday))},
			"Nanosecond":  {Typ: reflect.TypeOf(q.Nanosecond), Value: constant.MakeInt64(int64(q.Nanosecond))},
			"November":    {Typ: reflect.TypeOf(q.November), Value: constant.MakeInt64(int64(q.November))},
			"October":     {Typ: reflect.TypeOf(q.October), Value: constant.MakeInt64(int64(q.October))},
			"Saturday":    {Typ: reflect.TypeOf(q.Saturday), Value: constant.MakeInt64(int64(q.Saturday))},
			"Second":      {Typ: reflect.TypeOf(q.Second), Value: constant.MakeInt64(int64(q.Second))},
			"September":   {Typ: reflect.TypeOf(q.September), Value: constant.MakeInt64(int64(q.September))},
			"Sunday":      {Typ: reflect.TypeOf(q.Sunday), Value: constant.MakeInt64(int64(q.Sunday))},
			"Thursday":    {Typ: reflect.TypeOf(q.Thursday), Value: constant.MakeInt64(int64(q.Thursday))},
			"Tuesday":     {Typ: reflect.TypeOf(q.Tuesday), Value: constant.MakeInt64(int64(q.Tuesday))},
			"Wednesday":   {Typ: reflect.TypeOf(q.Wednesday), Value: constant.MakeInt64(int64(q.Wednesday))},
		},
		UntypedConsts: map[string]ixgo.UntypedConst{
			"ANSIC":       {Typ: "untyped string", Value: constant.MakeString(string(q.ANSIC))},
			"DateOnly":    {Typ: "untyped string", Value: constant.MakeString(string(q.DateOnly))},
			"DateTime":    {Typ: "untyped string", Value: constant.MakeString(string(q.DateTime))},
			"Kitchen":     {Typ: "untyped string", Value: constant.MakeString(string(q.Kitchen))},
			"Layout":      {Typ: "untyped string", Value: constant.MakeString(string(q.Layout))},
			"RFC1123":     {Typ: "untyped string", Value: constant.MakeString(string(q.RFC1123))},
			"RFC1123Z":    {Typ: "untyped string", Value: constant.MakeString(string(q.RFC1123Z))},
			"RFC3339":     {Typ: "untyped string", Value: constant.MakeString(string(q.RFC3339))},
			"RFC3339Nano": {Typ: "untyped string", Value: constant.MakeString(string(q.RFC3339Nano))},
			"RFC822":      {Typ: "untyped string", Value: constant.MakeString(string(q.RFC822))},
			"RFC822Z":     {Typ: "untyped string", Value: constant.MakeString(string(q.RFC822Z))},
			"RFC850":      {Typ: "untyped string", Value: constant.MakeString(string(q.RFC850))},
			"RubyDate":    {Typ: "untyped string", Value: constant.MakeString(string(q.RubyDate))},
			"Stamp":       {Typ: "untyped string", Value: constant.MakeString(string(q.Stamp))},
			"StampMicro":  {Typ: "untyped string", Value: constant.MakeString(string(q.StampMicro))},
			"StampMilli":  {Typ: "untyped string", Value: constant.MakeString(string(q.StampMilli))},
			"StampNano":   {Typ: "untyped string", Value: constant.MakeString(string(q.StampNano))},
			"TimeOnly":    {Typ: "untyped string", Value: constant.MakeString(string(q.TimeOnly))},
			"UnixDate":    {Typ: "untyped string", Value: constant.MakeString(string(q.UnixDate))},
		},
	})
}
