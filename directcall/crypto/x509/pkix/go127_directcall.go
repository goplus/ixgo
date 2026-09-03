// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package pkix

import (
	q "crypto/x509/pkix"

	"github.com/goplus/ixgo"
	time "time"
)

func init() {
	ixgo.RegisterDirectCalls("crypto/x509/pkix", map[string]ixgo.DirectCallAdapter{
		"(*CertificateList).HasExpired": method_ptr_CertificateList_HasExpired,
		"(*Name).FillFromRDNSequence":   method_ptr_Name_FillFromRDNSequence,
		"(*Name).String":                method_ptr_Name_String,
		"(*Name).ToRDNSequence":         method_ptr_Name_ToRDNSequence,
		"(*RDNSequence).String":         method_ptr_RDNSequence_String,
		"(Name).String":                 method_Name_String,
		"(Name).ToRDNSequence":          method_Name_ToRDNSequence,
		"(RDNSequence).String":          method_RDNSequence_String,
	})
}

func method_ptr_CertificateList_HasExpired(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.CertificateList).HasExpired(ixgo.DirectCallArg[*q.CertificateList](ctx, 0), ixgo.DirectCallArg[time.Time](ctx, 1)))
}

func method_ptr_Name_FillFromRDNSequence(ctx ixgo.DirectCallContext) {
	(*q.Name).FillFromRDNSequence(ixgo.DirectCallArg[*q.Name](ctx, 0), ixgo.DirectCallArg[*q.RDNSequence](ctx, 1))
}

func method_Name_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Name.String(ixgo.DirectCallArg[q.Name](ctx, 0)))
}

func method_ptr_Name_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Name).String(ixgo.DirectCallArg[*q.Name](ctx, 0)))
}

func method_Name_ToRDNSequence(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Name.ToRDNSequence(ixgo.DirectCallArg[q.Name](ctx, 0)))
}

func method_ptr_Name_ToRDNSequence(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Name).ToRDNSequence(ixgo.DirectCallArg[*q.Name](ctx, 0)))
}

func method_RDNSequence_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.RDNSequence.String(ixgo.DirectCallArg[q.RDNSequence](ctx, 0)))
}

func method_ptr_RDNSequence_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.RDNSequence).String(ixgo.DirectCallArg[*q.RDNSequence](ctx, 0)))
}
