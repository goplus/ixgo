// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package time

import (
	q "time"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("time", map[string]ixgo.DirectCallAdapter{
		"(*Duration).Abs":          method_ptr_Duration_Abs,
		"(*Duration).Hours":        method_ptr_Duration_Hours,
		"(*Duration).Microseconds": method_ptr_Duration_Microseconds,
		"(*Duration).Milliseconds": method_ptr_Duration_Milliseconds,
		"(*Duration).Minutes":      method_ptr_Duration_Minutes,
		"(*Duration).Nanoseconds":  method_ptr_Duration_Nanoseconds,
		"(*Duration).Round":        method_ptr_Duration_Round,
		"(*Duration).Seconds":      method_ptr_Duration_Seconds,
		"(*Duration).String":       method_ptr_Duration_String,
		"(*Duration).Truncate":     method_ptr_Duration_Truncate,
		"(*Location).String":       method_ptr_Location_String,
		"(*Month).String":          method_ptr_Month_String,
		"(*ParseError).Error":      method_ptr_ParseError_Error,
		"(*Ticker).Reset":          method_ptr_Ticker_Reset,
		"(*Ticker).Stop":           method_ptr_Ticker_Stop,
		"(*Time).Add":              method_ptr_Time_Add,
		"(*Time).AddDate":          method_ptr_Time_AddDate,
		"(*Time).After":            method_ptr_Time_After,
		"(*Time).AppendFormat":     method_ptr_Time_AppendFormat,
		"(*Time).Before":           method_ptr_Time_Before,
		"(*Time).Compare":          method_ptr_Time_Compare,
		"(*Time).Day":              method_ptr_Time_Day,
		"(*Time).Equal":            method_ptr_Time_Equal,
		"(*Time).Format":           method_ptr_Time_Format,
		"(*Time).GoString":         method_ptr_Time_GoString,
		"(*Time).GobDecode":        method_ptr_Time_GobDecode,
		"(*Time).Hour":             method_ptr_Time_Hour,
		"(*Time).In":               method_ptr_Time_In,
		"(*Time).IsDST":            method_ptr_Time_IsDST,
		"(*Time).IsZero":           method_ptr_Time_IsZero,
		"(*Time).Local":            method_ptr_Time_Local,
		"(*Time).Location":         method_ptr_Time_Location,
		"(*Time).Minute":           method_ptr_Time_Minute,
		"(*Time).Month":            method_ptr_Time_Month,
		"(*Time).Nanosecond":       method_ptr_Time_Nanosecond,
		"(*Time).Round":            method_ptr_Time_Round,
		"(*Time).Second":           method_ptr_Time_Second,
		"(*Time).String":           method_ptr_Time_String,
		"(*Time).Sub":              method_ptr_Time_Sub,
		"(*Time).Truncate":         method_ptr_Time_Truncate,
		"(*Time).UTC":              method_ptr_Time_UTC,
		"(*Time).Unix":             method_ptr_Time_Unix,
		"(*Time).UnixMicro":        method_ptr_Time_UnixMicro,
		"(*Time).UnixMilli":        method_ptr_Time_UnixMilli,
		"(*Time).UnixNano":         method_ptr_Time_UnixNano,
		"(*Time).UnmarshalBinary":  method_ptr_Time_UnmarshalBinary,
		"(*Time).UnmarshalJSON":    method_ptr_Time_UnmarshalJSON,
		"(*Time).UnmarshalText":    method_ptr_Time_UnmarshalText,
		"(*Time).Weekday":          method_ptr_Time_Weekday,
		"(*Time).Year":             method_ptr_Time_Year,
		"(*Time).YearDay":          method_ptr_Time_YearDay,
		"(*Timer).Reset":           method_ptr_Timer_Reset,
		"(*Timer).Stop":            method_ptr_Timer_Stop,
		"(*Weekday).String":        method_ptr_Weekday_String,
		"(Duration).Abs":           method_Duration_Abs,
		"(Duration).Hours":         method_Duration_Hours,
		"(Duration).Microseconds":  method_Duration_Microseconds,
		"(Duration).Milliseconds":  method_Duration_Milliseconds,
		"(Duration).Minutes":       method_Duration_Minutes,
		"(Duration).Nanoseconds":   method_Duration_Nanoseconds,
		"(Duration).Round":         method_Duration_Round,
		"(Duration).Seconds":       method_Duration_Seconds,
		"(Duration).String":        method_Duration_String,
		"(Duration).Truncate":      method_Duration_Truncate,
		"(Month).String":           method_Month_String,
		"(Time).Add":               method_Time_Add,
		"(Time).AddDate":           method_Time_AddDate,
		"(Time).After":             method_Time_After,
		"(Time).AppendFormat":      method_Time_AppendFormat,
		"(Time).Before":            method_Time_Before,
		"(Time).Compare":           method_Time_Compare,
		"(Time).Day":               method_Time_Day,
		"(Time).Equal":             method_Time_Equal,
		"(Time).Format":            method_Time_Format,
		"(Time).GoString":          method_Time_GoString,
		"(Time).Hour":              method_Time_Hour,
		"(Time).In":                method_Time_In,
		"(Time).IsDST":             method_Time_IsDST,
		"(Time).IsZero":            method_Time_IsZero,
		"(Time).Local":             method_Time_Local,
		"(Time).Location":          method_Time_Location,
		"(Time).Minute":            method_Time_Minute,
		"(Time).Month":             method_Time_Month,
		"(Time).Nanosecond":        method_Time_Nanosecond,
		"(Time).Round":             method_Time_Round,
		"(Time).Second":            method_Time_Second,
		"(Time).String":            method_Time_String,
		"(Time).Sub":               method_Time_Sub,
		"(Time).Truncate":          method_Time_Truncate,
		"(Time).UTC":               method_Time_UTC,
		"(Time).Unix":              method_Time_Unix,
		"(Time).UnixMicro":         method_Time_UnixMicro,
		"(Time).UnixMilli":         method_Time_UnixMilli,
		"(Time).UnixNano":          method_Time_UnixNano,
		"(Time).Weekday":           method_Time_Weekday,
		"(Time).Year":              method_Time_Year,
		"(Time).YearDay":           method_Time_YearDay,
		"(Weekday).String":         method_Weekday_String,
		"After":                    func_After,
		"AfterFunc":                func_AfterFunc,
		"Date":                     func_Date,
		"FixedZone":                func_FixedZone,
		"NewTicker":                func_NewTicker,
		"NewTimer":                 func_NewTimer,
		"Now":                      func_Now,
		"Since":                    func_Since,
		"Sleep":                    func_Sleep,
		"Tick":                     func_Tick,
		"Unix":                     func_Unix,
		"UnixMicro":                func_UnixMicro,
		"UnixMilli":                func_UnixMilli,
		"Until":                    func_Until,
	})
}

func func_After(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.After(ixgo.DirectCallArg[q.Duration](ctx, 0)))
}

func func_AfterFunc(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.AfterFunc(ixgo.DirectCallArg[q.Duration](ctx, 0), ixgo.DirectCallArg[func()](ctx, 1)))
}

func func_Date(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Date(ixgo.DirectCallArg[int](ctx, 0), ixgo.DirectCallArg[q.Month](ctx, 1), ixgo.DirectCallArg[int](ctx, 2), ixgo.DirectCallArg[int](ctx, 3), ixgo.DirectCallArg[int](ctx, 4), ixgo.DirectCallArg[int](ctx, 5), ixgo.DirectCallArg[int](ctx, 6), ixgo.DirectCallArg[*q.Location](ctx, 7)))
}

func method_Duration_Abs(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Duration.Abs(ixgo.DirectCallArg[q.Duration](ctx, 0)))
}

func method_ptr_Duration_Abs(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Duration).Abs(ixgo.DirectCallArg[*q.Duration](ctx, 0)))
}

func method_Duration_Hours(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Duration.Hours(ixgo.DirectCallArg[q.Duration](ctx, 0)))
}

func method_ptr_Duration_Hours(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Duration).Hours(ixgo.DirectCallArg[*q.Duration](ctx, 0)))
}

func method_Duration_Microseconds(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Duration.Microseconds(ixgo.DirectCallArg[q.Duration](ctx, 0)))
}

func method_ptr_Duration_Microseconds(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Duration).Microseconds(ixgo.DirectCallArg[*q.Duration](ctx, 0)))
}

func method_Duration_Milliseconds(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Duration.Milliseconds(ixgo.DirectCallArg[q.Duration](ctx, 0)))
}

func method_ptr_Duration_Milliseconds(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Duration).Milliseconds(ixgo.DirectCallArg[*q.Duration](ctx, 0)))
}

func method_Duration_Minutes(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Duration.Minutes(ixgo.DirectCallArg[q.Duration](ctx, 0)))
}

func method_ptr_Duration_Minutes(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Duration).Minutes(ixgo.DirectCallArg[*q.Duration](ctx, 0)))
}

func method_Duration_Nanoseconds(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Duration.Nanoseconds(ixgo.DirectCallArg[q.Duration](ctx, 0)))
}

func method_ptr_Duration_Nanoseconds(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Duration).Nanoseconds(ixgo.DirectCallArg[*q.Duration](ctx, 0)))
}

func method_Duration_Round(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Duration.Round(ixgo.DirectCallArg[q.Duration](ctx, 0), ixgo.DirectCallArg[q.Duration](ctx, 1)))
}

func method_ptr_Duration_Round(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Duration).Round(ixgo.DirectCallArg[*q.Duration](ctx, 0), ixgo.DirectCallArg[q.Duration](ctx, 1)))
}

func method_Duration_Seconds(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Duration.Seconds(ixgo.DirectCallArg[q.Duration](ctx, 0)))
}

func method_ptr_Duration_Seconds(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Duration).Seconds(ixgo.DirectCallArg[*q.Duration](ctx, 0)))
}

func method_Duration_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Duration.String(ixgo.DirectCallArg[q.Duration](ctx, 0)))
}

func method_ptr_Duration_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Duration).String(ixgo.DirectCallArg[*q.Duration](ctx, 0)))
}

func method_Duration_Truncate(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Duration.Truncate(ixgo.DirectCallArg[q.Duration](ctx, 0), ixgo.DirectCallArg[q.Duration](ctx, 1)))
}

func method_ptr_Duration_Truncate(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Duration).Truncate(ixgo.DirectCallArg[*q.Duration](ctx, 0), ixgo.DirectCallArg[q.Duration](ctx, 1)))
}

func func_FixedZone(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.FixedZone(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func method_ptr_Location_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Location).String(ixgo.DirectCallArg[*q.Location](ctx, 0)))
}

func method_Month_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Month.String(ixgo.DirectCallArg[q.Month](ctx, 0)))
}

func method_ptr_Month_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Month).String(ixgo.DirectCallArg[*q.Month](ctx, 0)))
}

func func_NewTicker(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewTicker(ixgo.DirectCallArg[q.Duration](ctx, 0)))
}

func func_NewTimer(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewTimer(ixgo.DirectCallArg[q.Duration](ctx, 0)))
}

func func_Now(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Now())
}

func method_ptr_ParseError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ParseError).Error(ixgo.DirectCallArg[*q.ParseError](ctx, 0)))
}

func func_Since(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Since(ixgo.DirectCallArg[q.Time](ctx, 0)))
}

func func_Sleep(ctx ixgo.DirectCallContext) {
	q.Sleep(ixgo.DirectCallArg[q.Duration](ctx, 0))
}

func func_Tick(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Tick(ixgo.DirectCallArg[q.Duration](ctx, 0)))
}

func method_ptr_Ticker_Reset(ctx ixgo.DirectCallContext) {
	(*q.Ticker).Reset(ixgo.DirectCallArg[*q.Ticker](ctx, 0), ixgo.DirectCallArg[q.Duration](ctx, 1))
}

func method_ptr_Ticker_Stop(ctx ixgo.DirectCallContext) {
	(*q.Ticker).Stop(ixgo.DirectCallArg[*q.Ticker](ctx, 0))
}

func method_Time_Add(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Time.Add(ixgo.DirectCallArg[q.Time](ctx, 0), ixgo.DirectCallArg[q.Duration](ctx, 1)))
}

func method_ptr_Time_Add(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Time).Add(ixgo.DirectCallArg[*q.Time](ctx, 0), ixgo.DirectCallArg[q.Duration](ctx, 1)))
}

func method_Time_AddDate(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Time.AddDate(ixgo.DirectCallArg[q.Time](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2), ixgo.DirectCallArg[int](ctx, 3)))
}

func method_ptr_Time_AddDate(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Time).AddDate(ixgo.DirectCallArg[*q.Time](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2), ixgo.DirectCallArg[int](ctx, 3)))
}

func method_Time_After(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Time.After(ixgo.DirectCallArg[q.Time](ctx, 0), ixgo.DirectCallArg[q.Time](ctx, 1)))
}

func method_ptr_Time_After(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Time).After(ixgo.DirectCallArg[*q.Time](ctx, 0), ixgo.DirectCallArg[q.Time](ctx, 1)))
}

func method_Time_AppendFormat(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Time.AppendFormat(ixgo.DirectCallArg[q.Time](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1), ixgo.DirectCallArg[string](ctx, 2)))
}

func method_ptr_Time_AppendFormat(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Time).AppendFormat(ixgo.DirectCallArg[*q.Time](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1), ixgo.DirectCallArg[string](ctx, 2)))
}

func method_Time_Before(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Time.Before(ixgo.DirectCallArg[q.Time](ctx, 0), ixgo.DirectCallArg[q.Time](ctx, 1)))
}

func method_ptr_Time_Before(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Time).Before(ixgo.DirectCallArg[*q.Time](ctx, 0), ixgo.DirectCallArg[q.Time](ctx, 1)))
}

func method_Time_Compare(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Time.Compare(ixgo.DirectCallArg[q.Time](ctx, 0), ixgo.DirectCallArg[q.Time](ctx, 1)))
}

func method_ptr_Time_Compare(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Time).Compare(ixgo.DirectCallArg[*q.Time](ctx, 0), ixgo.DirectCallArg[q.Time](ctx, 1)))
}

func method_Time_Day(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Time.Day(ixgo.DirectCallArg[q.Time](ctx, 0)))
}

func method_ptr_Time_Day(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Time).Day(ixgo.DirectCallArg[*q.Time](ctx, 0)))
}

func method_Time_Equal(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Time.Equal(ixgo.DirectCallArg[q.Time](ctx, 0), ixgo.DirectCallArg[q.Time](ctx, 1)))
}

func method_ptr_Time_Equal(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Time).Equal(ixgo.DirectCallArg[*q.Time](ctx, 0), ixgo.DirectCallArg[q.Time](ctx, 1)))
}

func method_Time_Format(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Time.Format(ixgo.DirectCallArg[q.Time](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_Time_Format(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Time).Format(ixgo.DirectCallArg[*q.Time](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_Time_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Time.GoString(ixgo.DirectCallArg[q.Time](ctx, 0)))
}

func method_ptr_Time_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Time).GoString(ixgo.DirectCallArg[*q.Time](ctx, 0)))
}

func method_ptr_Time_GobDecode(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Time).GobDecode(ixgo.DirectCallArg[*q.Time](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func method_Time_Hour(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Time.Hour(ixgo.DirectCallArg[q.Time](ctx, 0)))
}

func method_ptr_Time_Hour(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Time).Hour(ixgo.DirectCallArg[*q.Time](ctx, 0)))
}

func method_Time_In(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Time.In(ixgo.DirectCallArg[q.Time](ctx, 0), ixgo.DirectCallArg[*q.Location](ctx, 1)))
}

func method_ptr_Time_In(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Time).In(ixgo.DirectCallArg[*q.Time](ctx, 0), ixgo.DirectCallArg[*q.Location](ctx, 1)))
}

func method_Time_IsDST(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Time.IsDST(ixgo.DirectCallArg[q.Time](ctx, 0)))
}

func method_ptr_Time_IsDST(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Time).IsDST(ixgo.DirectCallArg[*q.Time](ctx, 0)))
}

func method_Time_IsZero(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Time.IsZero(ixgo.DirectCallArg[q.Time](ctx, 0)))
}

func method_ptr_Time_IsZero(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Time).IsZero(ixgo.DirectCallArg[*q.Time](ctx, 0)))
}

func method_Time_Local(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Time.Local(ixgo.DirectCallArg[q.Time](ctx, 0)))
}

func method_ptr_Time_Local(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Time).Local(ixgo.DirectCallArg[*q.Time](ctx, 0)))
}

func method_Time_Location(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Time.Location(ixgo.DirectCallArg[q.Time](ctx, 0)))
}

func method_ptr_Time_Location(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Time).Location(ixgo.DirectCallArg[*q.Time](ctx, 0)))
}

func method_Time_Minute(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Time.Minute(ixgo.DirectCallArg[q.Time](ctx, 0)))
}

func method_ptr_Time_Minute(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Time).Minute(ixgo.DirectCallArg[*q.Time](ctx, 0)))
}

func method_Time_Month(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Time.Month(ixgo.DirectCallArg[q.Time](ctx, 0)))
}

func method_ptr_Time_Month(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Time).Month(ixgo.DirectCallArg[*q.Time](ctx, 0)))
}

func method_Time_Nanosecond(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Time.Nanosecond(ixgo.DirectCallArg[q.Time](ctx, 0)))
}

func method_ptr_Time_Nanosecond(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Time).Nanosecond(ixgo.DirectCallArg[*q.Time](ctx, 0)))
}

func method_Time_Round(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Time.Round(ixgo.DirectCallArg[q.Time](ctx, 0), ixgo.DirectCallArg[q.Duration](ctx, 1)))
}

func method_ptr_Time_Round(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Time).Round(ixgo.DirectCallArg[*q.Time](ctx, 0), ixgo.DirectCallArg[q.Duration](ctx, 1)))
}

func method_Time_Second(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Time.Second(ixgo.DirectCallArg[q.Time](ctx, 0)))
}

func method_ptr_Time_Second(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Time).Second(ixgo.DirectCallArg[*q.Time](ctx, 0)))
}

func method_Time_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Time.String(ixgo.DirectCallArg[q.Time](ctx, 0)))
}

func method_ptr_Time_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Time).String(ixgo.DirectCallArg[*q.Time](ctx, 0)))
}

func method_Time_Sub(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Time.Sub(ixgo.DirectCallArg[q.Time](ctx, 0), ixgo.DirectCallArg[q.Time](ctx, 1)))
}

func method_ptr_Time_Sub(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Time).Sub(ixgo.DirectCallArg[*q.Time](ctx, 0), ixgo.DirectCallArg[q.Time](ctx, 1)))
}

func method_Time_Truncate(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Time.Truncate(ixgo.DirectCallArg[q.Time](ctx, 0), ixgo.DirectCallArg[q.Duration](ctx, 1)))
}

func method_ptr_Time_Truncate(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Time).Truncate(ixgo.DirectCallArg[*q.Time](ctx, 0), ixgo.DirectCallArg[q.Duration](ctx, 1)))
}

func method_Time_UTC(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Time.UTC(ixgo.DirectCallArg[q.Time](ctx, 0)))
}

func method_ptr_Time_UTC(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Time).UTC(ixgo.DirectCallArg[*q.Time](ctx, 0)))
}

func method_Time_Unix(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Time.Unix(ixgo.DirectCallArg[q.Time](ctx, 0)))
}

func method_ptr_Time_Unix(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Time).Unix(ixgo.DirectCallArg[*q.Time](ctx, 0)))
}

func method_Time_UnixMicro(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Time.UnixMicro(ixgo.DirectCallArg[q.Time](ctx, 0)))
}

func method_ptr_Time_UnixMicro(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Time).UnixMicro(ixgo.DirectCallArg[*q.Time](ctx, 0)))
}

func method_Time_UnixMilli(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Time.UnixMilli(ixgo.DirectCallArg[q.Time](ctx, 0)))
}

func method_ptr_Time_UnixMilli(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Time).UnixMilli(ixgo.DirectCallArg[*q.Time](ctx, 0)))
}

func method_Time_UnixNano(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Time.UnixNano(ixgo.DirectCallArg[q.Time](ctx, 0)))
}

func method_ptr_Time_UnixNano(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Time).UnixNano(ixgo.DirectCallArg[*q.Time](ctx, 0)))
}

func method_ptr_Time_UnmarshalBinary(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Time).UnmarshalBinary(ixgo.DirectCallArg[*q.Time](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func method_ptr_Time_UnmarshalJSON(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Time).UnmarshalJSON(ixgo.DirectCallArg[*q.Time](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func method_ptr_Time_UnmarshalText(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Time).UnmarshalText(ixgo.DirectCallArg[*q.Time](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func method_Time_Weekday(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Time.Weekday(ixgo.DirectCallArg[q.Time](ctx, 0)))
}

func method_ptr_Time_Weekday(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Time).Weekday(ixgo.DirectCallArg[*q.Time](ctx, 0)))
}

func method_Time_Year(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Time.Year(ixgo.DirectCallArg[q.Time](ctx, 0)))
}

func method_ptr_Time_Year(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Time).Year(ixgo.DirectCallArg[*q.Time](ctx, 0)))
}

func method_Time_YearDay(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Time.YearDay(ixgo.DirectCallArg[q.Time](ctx, 0)))
}

func method_ptr_Time_YearDay(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Time).YearDay(ixgo.DirectCallArg[*q.Time](ctx, 0)))
}

func method_ptr_Timer_Reset(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Timer).Reset(ixgo.DirectCallArg[*q.Timer](ctx, 0), ixgo.DirectCallArg[q.Duration](ctx, 1)))
}

func method_ptr_Timer_Stop(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Timer).Stop(ixgo.DirectCallArg[*q.Timer](ctx, 0)))
}

func func_Unix(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Unix(ixgo.DirectCallArg[int64](ctx, 0), ixgo.DirectCallArg[int64](ctx, 1)))
}

func func_UnixMicro(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.UnixMicro(ixgo.DirectCallArg[int64](ctx, 0)))
}

func func_UnixMilli(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.UnixMilli(ixgo.DirectCallArg[int64](ctx, 0)))
}

func func_Until(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Until(ixgo.DirectCallArg[q.Time](ctx, 0)))
}

func method_Weekday_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Weekday.String(ixgo.DirectCallArg[q.Weekday](ctx, 0)))
}

func method_ptr_Weekday_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Weekday).String(ixgo.DirectCallArg[*q.Weekday](ctx, 0)))
}
