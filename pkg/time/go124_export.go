// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package time

import (
	q "time"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("time", func() *ixgo.Package {
		return &ixgo.Package{
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
			Vars: map[string]interface{}{
				"Local": &q.Local,
				"UTC":   &q.UTC,
			},
			Funcs: map[string]interface{}{
				"After":                  q.After,
				"AfterFunc":              q.AfterFunc,
				"Date":                   q.Date,
				"FixedZone":              q.FixedZone,
				"LoadLocation":           q.LoadLocation,
				"LoadLocationFromTZData": q.LoadLocationFromTZData,
				"NewTicker":              q.NewTicker,
				"NewTimer":               q.NewTimer,
				"Now":                    q.Now,
				"Parse":                  q.Parse,
				"ParseDuration":          q.ParseDuration,
				"ParseInLocation":        q.ParseInLocation,
				"Since":                  q.Since,
				"Sleep":                  q.Sleep,
				"Tick":                   q.Tick,
				"Unix":                   q.Unix,
				"UnixMicro":              q.UnixMicro,
				"UnixMilli":              q.UnixMilli,
				"Until":                  q.Until,
			},
			TypedConsts: map[string]interface{}{
				"April":       q.April,
				"August":      q.August,
				"December":    q.December,
				"February":    q.February,
				"Friday":      q.Friday,
				"Hour":        q.Hour,
				"January":     q.January,
				"July":        q.July,
				"June":        q.June,
				"March":       q.March,
				"May":         q.May,
				"Microsecond": q.Microsecond,
				"Millisecond": q.Millisecond,
				"Minute":      q.Minute,
				"Monday":      q.Monday,
				"Nanosecond":  q.Nanosecond,
				"November":    q.November,
				"October":     q.October,
				"Saturday":    q.Saturday,
				"Second":      q.Second,
				"September":   q.September,
				"Sunday":      q.Sunday,
				"Thursday":    q.Thursday,
				"Tuesday":     q.Tuesday,
				"Wednesday":   q.Wednesday,
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
		}
	})
}
