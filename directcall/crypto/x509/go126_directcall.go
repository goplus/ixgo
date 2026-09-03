// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26 && !go1.27
// +build go1.26,!go1.27

package x509

import (
	q "crypto/x509"

	rsa "crypto/rsa"
	pkix "crypto/x509/pkix"
	asn1 "encoding/asn1"
	pem "encoding/pem"
	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("crypto/x509", map[string]ixgo.DirectCallAdapter{
		"(*CertPool).AddCert":                  method_ptr_CertPool_AddCert,
		"(*CertPool).AddCertWithConstraint":    method_ptr_CertPool_AddCertWithConstraint,
		"(*CertPool).AppendCertsFromPEM":       method_ptr_CertPool_AppendCertsFromPEM,
		"(*CertPool).Clone":                    method_ptr_CertPool_Clone,
		"(*CertPool).Equal":                    method_ptr_CertPool_Equal,
		"(*CertPool).Subjects":                 method_ptr_CertPool_Subjects,
		"(*Certificate).CheckCRLSignature":     method_ptr_Certificate_CheckCRLSignature,
		"(*Certificate).CheckSignature":        method_ptr_Certificate_CheckSignature,
		"(*Certificate).CheckSignatureFrom":    method_ptr_Certificate_CheckSignatureFrom,
		"(*Certificate).Equal":                 method_ptr_Certificate_Equal,
		"(*Certificate).VerifyHostname":        method_ptr_Certificate_VerifyHostname,
		"(*CertificateInvalidError).Error":     method_ptr_CertificateInvalidError_Error,
		"(*CertificateRequest).CheckSignature": method_ptr_CertificateRequest_CheckSignature,
		"(*ConstraintViolationError).Error":    method_ptr_ConstraintViolationError_Error,
		"(*ExtKeyUsage).OID":                   method_ptr_ExtKeyUsage_OID,
		"(*ExtKeyUsage).String":                method_ptr_ExtKeyUsage_String,
		"(*HostnameError).Error":               method_ptr_HostnameError_Error,
		"(*InsecureAlgorithmError).Error":      method_ptr_InsecureAlgorithmError_Error,
		"(*KeyUsage).String":                   method_ptr_KeyUsage_String,
		"(*OID).Equal":                         method_ptr_OID_Equal,
		"(*OID).EqualASN1OID":                  method_ptr_OID_EqualASN1OID,
		"(*OID).String":                        method_ptr_OID_String,
		"(*OID).UnmarshalBinary":               method_ptr_OID_UnmarshalBinary,
		"(*OID).UnmarshalText":                 method_ptr_OID_UnmarshalText,
		"(*PublicKeyAlgorithm).String":         method_ptr_PublicKeyAlgorithm_String,
		"(*RevocationList).CheckSignatureFrom": method_ptr_RevocationList_CheckSignatureFrom,
		"(*SignatureAlgorithm).String":         method_ptr_SignatureAlgorithm_String,
		"(*SystemRootsError).Error":            method_ptr_SystemRootsError_Error,
		"(*SystemRootsError).Unwrap":           method_ptr_SystemRootsError_Unwrap,
		"(*UnhandledCriticalExtension).Error":  method_ptr_UnhandledCriticalExtension_Error,
		"(*UnknownAuthorityError).Error":       method_ptr_UnknownAuthorityError_Error,
		"(CertificateInvalidError).Error":      method_CertificateInvalidError_Error,
		"(ConstraintViolationError).Error":     method_ConstraintViolationError_Error,
		"(ExtKeyUsage).OID":                    method_ExtKeyUsage_OID,
		"(ExtKeyUsage).String":                 method_ExtKeyUsage_String,
		"(HostnameError).Error":                method_HostnameError_Error,
		"(InsecureAlgorithmError).Error":       method_InsecureAlgorithmError_Error,
		"(KeyUsage).String":                    method_KeyUsage_String,
		"(OID).Equal":                          method_OID_Equal,
		"(OID).EqualASN1OID":                   method_OID_EqualASN1OID,
		"(OID).String":                         method_OID_String,
		"(PublicKeyAlgorithm).String":          method_PublicKeyAlgorithm_String,
		"(SignatureAlgorithm).String":          method_SignatureAlgorithm_String,
		"(SystemRootsError).Error":             method_SystemRootsError_Error,
		"(SystemRootsError).Unwrap":            method_SystemRootsError_Unwrap,
		"(UnhandledCriticalExtension).Error":   method_UnhandledCriticalExtension_Error,
		"(UnknownAuthorityError).Error":        method_UnknownAuthorityError_Error,
		"IsEncryptedPEMBlock":                  func_IsEncryptedPEMBlock,
		"MarshalPKCS1PrivateKey":               func_MarshalPKCS1PrivateKey,
		"MarshalPKCS1PublicKey":                func_MarshalPKCS1PublicKey,
		"NewCertPool":                          func_NewCertPool,
		"SetFallbackRoots":                     func_SetFallbackRoots,
	})
}

func method_ptr_CertPool_AddCert(ctx ixgo.DirectCallContext) {
	(*q.CertPool).AddCert(ixgo.DirectCallArg[*q.CertPool](ctx, 0), ixgo.DirectCallArg[*q.Certificate](ctx, 1))
}

func method_ptr_CertPool_AddCertWithConstraint(ctx ixgo.DirectCallContext) {
	(*q.CertPool).AddCertWithConstraint(ixgo.DirectCallArg[*q.CertPool](ctx, 0), ixgo.DirectCallArg[*q.Certificate](ctx, 1), ixgo.DirectCallArg[func([]*q.Certificate) error](ctx, 2))
}

func method_ptr_CertPool_AppendCertsFromPEM(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.CertPool).AppendCertsFromPEM(ixgo.DirectCallArg[*q.CertPool](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func method_ptr_CertPool_Clone(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.CertPool).Clone(ixgo.DirectCallArg[*q.CertPool](ctx, 0)))
}

func method_ptr_CertPool_Equal(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.CertPool).Equal(ixgo.DirectCallArg[*q.CertPool](ctx, 0), ixgo.DirectCallArg[*q.CertPool](ctx, 1)))
}

func method_ptr_CertPool_Subjects(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.CertPool).Subjects(ixgo.DirectCallArg[*q.CertPool](ctx, 0)))
}

func method_ptr_Certificate_CheckCRLSignature(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Certificate).CheckCRLSignature(ixgo.DirectCallArg[*q.Certificate](ctx, 0), ixgo.DirectCallArg[*pkix.CertificateList](ctx, 1)))
}

func method_ptr_Certificate_CheckSignature(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Certificate).CheckSignature(ixgo.DirectCallArg[*q.Certificate](ctx, 0), ixgo.DirectCallArg[q.SignatureAlgorithm](ctx, 1), ixgo.DirectCallArg[[]byte](ctx, 2), ixgo.DirectCallArg[[]byte](ctx, 3)))
}

func method_ptr_Certificate_CheckSignatureFrom(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Certificate).CheckSignatureFrom(ixgo.DirectCallArg[*q.Certificate](ctx, 0), ixgo.DirectCallArg[*q.Certificate](ctx, 1)))
}

func method_ptr_Certificate_Equal(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Certificate).Equal(ixgo.DirectCallArg[*q.Certificate](ctx, 0), ixgo.DirectCallArg[*q.Certificate](ctx, 1)))
}

func method_ptr_Certificate_VerifyHostname(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Certificate).VerifyHostname(ixgo.DirectCallArg[*q.Certificate](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_CertificateInvalidError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.CertificateInvalidError.Error(ixgo.DirectCallArg[q.CertificateInvalidError](ctx, 0)))
}

func method_ptr_CertificateInvalidError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.CertificateInvalidError).Error(ixgo.DirectCallArg[*q.CertificateInvalidError](ctx, 0)))
}

func method_ptr_CertificateRequest_CheckSignature(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.CertificateRequest).CheckSignature(ixgo.DirectCallArg[*q.CertificateRequest](ctx, 0)))
}

func method_ConstraintViolationError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ConstraintViolationError.Error(ixgo.DirectCallArg[q.ConstraintViolationError](ctx, 0)))
}

func method_ptr_ConstraintViolationError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ConstraintViolationError).Error(ixgo.DirectCallArg[*q.ConstraintViolationError](ctx, 0)))
}

func method_ExtKeyUsage_OID(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ExtKeyUsage.OID(ixgo.DirectCallArg[q.ExtKeyUsage](ctx, 0)))
}

func method_ptr_ExtKeyUsage_OID(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ExtKeyUsage).OID(ixgo.DirectCallArg[*q.ExtKeyUsage](ctx, 0)))
}

func method_ExtKeyUsage_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ExtKeyUsage.String(ixgo.DirectCallArg[q.ExtKeyUsage](ctx, 0)))
}

func method_ptr_ExtKeyUsage_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ExtKeyUsage).String(ixgo.DirectCallArg[*q.ExtKeyUsage](ctx, 0)))
}

func method_HostnameError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.HostnameError.Error(ixgo.DirectCallArg[q.HostnameError](ctx, 0)))
}

func method_ptr_HostnameError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.HostnameError).Error(ixgo.DirectCallArg[*q.HostnameError](ctx, 0)))
}

func method_InsecureAlgorithmError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.InsecureAlgorithmError.Error(ixgo.DirectCallArg[q.InsecureAlgorithmError](ctx, 0)))
}

func method_ptr_InsecureAlgorithmError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.InsecureAlgorithmError).Error(ixgo.DirectCallArg[*q.InsecureAlgorithmError](ctx, 0)))
}

func func_IsEncryptedPEMBlock(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IsEncryptedPEMBlock(ixgo.DirectCallArg[*pem.Block](ctx, 0)))
}

func method_KeyUsage_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.KeyUsage.String(ixgo.DirectCallArg[q.KeyUsage](ctx, 0)))
}

func method_ptr_KeyUsage_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.KeyUsage).String(ixgo.DirectCallArg[*q.KeyUsage](ctx, 0)))
}

func func_MarshalPKCS1PrivateKey(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.MarshalPKCS1PrivateKey(ixgo.DirectCallArg[*rsa.PrivateKey](ctx, 0)))
}

func func_MarshalPKCS1PublicKey(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.MarshalPKCS1PublicKey(ixgo.DirectCallArg[*rsa.PublicKey](ctx, 0)))
}

func func_NewCertPool(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewCertPool())
}

func method_OID_Equal(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.OID.Equal(ixgo.DirectCallArg[q.OID](ctx, 0), ixgo.DirectCallArg[q.OID](ctx, 1)))
}

func method_ptr_OID_Equal(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.OID).Equal(ixgo.DirectCallArg[*q.OID](ctx, 0), ixgo.DirectCallArg[q.OID](ctx, 1)))
}

func method_OID_EqualASN1OID(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.OID.EqualASN1OID(ixgo.DirectCallArg[q.OID](ctx, 0), ixgo.DirectCallArg[asn1.ObjectIdentifier](ctx, 1)))
}

func method_ptr_OID_EqualASN1OID(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.OID).EqualASN1OID(ixgo.DirectCallArg[*q.OID](ctx, 0), ixgo.DirectCallArg[asn1.ObjectIdentifier](ctx, 1)))
}

func method_OID_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.OID.String(ixgo.DirectCallArg[q.OID](ctx, 0)))
}

func method_ptr_OID_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.OID).String(ixgo.DirectCallArg[*q.OID](ctx, 0)))
}

func method_ptr_OID_UnmarshalBinary(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.OID).UnmarshalBinary(ixgo.DirectCallArg[*q.OID](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func method_ptr_OID_UnmarshalText(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.OID).UnmarshalText(ixgo.DirectCallArg[*q.OID](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func method_PublicKeyAlgorithm_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.PublicKeyAlgorithm.String(ixgo.DirectCallArg[q.PublicKeyAlgorithm](ctx, 0)))
}

func method_ptr_PublicKeyAlgorithm_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.PublicKeyAlgorithm).String(ixgo.DirectCallArg[*q.PublicKeyAlgorithm](ctx, 0)))
}

func method_ptr_RevocationList_CheckSignatureFrom(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.RevocationList).CheckSignatureFrom(ixgo.DirectCallArg[*q.RevocationList](ctx, 0), ixgo.DirectCallArg[*q.Certificate](ctx, 1)))
}

func func_SetFallbackRoots(ctx ixgo.DirectCallContext) {
	q.SetFallbackRoots(ixgo.DirectCallArg[*q.CertPool](ctx, 0))
}

func method_SignatureAlgorithm_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.SignatureAlgorithm.String(ixgo.DirectCallArg[q.SignatureAlgorithm](ctx, 0)))
}

func method_ptr_SignatureAlgorithm_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.SignatureAlgorithm).String(ixgo.DirectCallArg[*q.SignatureAlgorithm](ctx, 0)))
}

func method_SystemRootsError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.SystemRootsError.Error(ixgo.DirectCallArg[q.SystemRootsError](ctx, 0)))
}

func method_ptr_SystemRootsError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.SystemRootsError).Error(ixgo.DirectCallArg[*q.SystemRootsError](ctx, 0)))
}

func method_SystemRootsError_Unwrap(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.SystemRootsError.Unwrap(ixgo.DirectCallArg[q.SystemRootsError](ctx, 0)))
}

func method_ptr_SystemRootsError_Unwrap(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.SystemRootsError).Unwrap(ixgo.DirectCallArg[*q.SystemRootsError](ctx, 0)))
}

func method_UnhandledCriticalExtension_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.UnhandledCriticalExtension.Error(ixgo.DirectCallArg[q.UnhandledCriticalExtension](ctx, 0)))
}

func method_ptr_UnhandledCriticalExtension_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.UnhandledCriticalExtension).Error(ixgo.DirectCallArg[*q.UnhandledCriticalExtension](ctx, 0)))
}

func method_UnknownAuthorityError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.UnknownAuthorityError.Error(ixgo.DirectCallArg[q.UnknownAuthorityError](ctx, 0)))
}

func method_ptr_UnknownAuthorityError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.UnknownAuthorityError).Error(ixgo.DirectCallArg[*q.UnknownAuthorityError](ctx, 0)))
}
