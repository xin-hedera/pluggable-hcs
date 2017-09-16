// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msp/msp_config.proto

package msp

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MSPConfig collects all the configuration information for
// an MSP. The Config field should be unmarshalled in a way
// that depends on the Type
type MSPConfig struct {
	// Type holds the type of the MSP; the default one would
	// be of type FABRIC implementing an X.509 based provider
	Type int32 `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`
	// Config is MSP dependent configuration info
	Config []byte `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
}

func (m *MSPConfig) Reset()                    { *m = MSPConfig{} }
func (m *MSPConfig) String() string            { return proto.CompactTextString(m) }
func (*MSPConfig) ProtoMessage()               {}
func (*MSPConfig) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *MSPConfig) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *MSPConfig) GetConfig() []byte {
	if m != nil {
		return m.Config
	}
	return nil
}

// FabricMSPConfig collects all the configuration information for
// a Fabric MSP.
// Here we assume a default certificate validation policy, where
// any certificate signed by any of the listed rootCA certs would
// be considered as valid under this MSP.
// This MSP may or may not come with a signing identity. If it does,
// it can also issue signing identities. If it does not, it can only
// be used to validate and verify certificates.
type FabricMSPConfig struct {
	// Name holds the identifier of the MSP; MSP identifier
	// is chosen by the application that governs this MSP.
	// For example, and assuming the default implementation of MSP,
	// that is X.509-based and considers a single Issuer,
	// this can refer to the Subject OU field or the Issuer OU field.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// List of root certificates trusted by this MSP
	// they are used upon certificate validation (see
	// comment for IntermediateCerts below)
	RootCerts [][]byte `protobuf:"bytes,2,rep,name=root_certs,json=rootCerts,proto3" json:"root_certs,omitempty"`
	// List of intermediate certificates trusted by this MSP;
	// they are used upon certificate validation as follows:
	// validation attempts to build a path from the certificate
	// to be validated (which is at one end of the path) and
	// one of the certs in the RootCerts field (which is at
	// the other end of the path). If the path is longer than
	// 2, certificates in the middle are searched within the
	// IntermediateCerts pool
	IntermediateCerts [][]byte `protobuf:"bytes,3,rep,name=intermediate_certs,json=intermediateCerts,proto3" json:"intermediate_certs,omitempty"`
	// Identity denoting the administrator of this MSP
	Admins [][]byte `protobuf:"bytes,4,rep,name=admins,proto3" json:"admins,omitempty"`
	// Identity revocation list
	RevocationList [][]byte `protobuf:"bytes,5,rep,name=revocation_list,json=revocationList,proto3" json:"revocation_list,omitempty"`
	// SigningIdentity holds information on the signing identity
	// this peer is to use, and which is to be imported by the
	// MSP defined before
	SigningIdentity *SigningIdentityInfo `protobuf:"bytes,6,opt,name=signing_identity,json=signingIdentity" json:"signing_identity,omitempty"`
	// OrganizationalUnitIdentifiers holds one or more
	// fabric organizational unit identifiers that belong to
	// this MSP configuration
	OrganizationalUnitIdentifiers []*FabricOUIdentifier `protobuf:"bytes,7,rep,name=organizational_unit_identifiers,json=organizationalUnitIdentifiers" json:"organizational_unit_identifiers,omitempty"`
	// FabricCryptoConfig contains the configuration parameters
	// for the cryptographic algorithms used by this MSP
	CryptoConfig *FabricCryptoConfig `protobuf:"bytes,8,opt,name=crypto_config,json=cryptoConfig" json:"crypto_config,omitempty"`
	// List of TLS root certificates trusted by this MSP.
	// They are returned by GetTLSRootCerts.
	TlsRootCerts [][]byte `protobuf:"bytes,9,rep,name=tls_root_certs,json=tlsRootCerts,proto3" json:"tls_root_certs,omitempty"`
	// List of TLS intermediate certificates trusted by this MSP;
	// They are returned by GetTLSIntermediateCerts.
	TlsIntermediateCerts [][]byte `protobuf:"bytes,10,rep,name=tls_intermediate_certs,json=tlsIntermediateCerts,proto3" json:"tls_intermediate_certs,omitempty"`
}

func (m *FabricMSPConfig) Reset()                    { *m = FabricMSPConfig{} }
func (m *FabricMSPConfig) String() string            { return proto.CompactTextString(m) }
func (*FabricMSPConfig) ProtoMessage()               {}
func (*FabricMSPConfig) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *FabricMSPConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FabricMSPConfig) GetRootCerts() [][]byte {
	if m != nil {
		return m.RootCerts
	}
	return nil
}

func (m *FabricMSPConfig) GetIntermediateCerts() [][]byte {
	if m != nil {
		return m.IntermediateCerts
	}
	return nil
}

func (m *FabricMSPConfig) GetAdmins() [][]byte {
	if m != nil {
		return m.Admins
	}
	return nil
}

func (m *FabricMSPConfig) GetRevocationList() [][]byte {
	if m != nil {
		return m.RevocationList
	}
	return nil
}

func (m *FabricMSPConfig) GetSigningIdentity() *SigningIdentityInfo {
	if m != nil {
		return m.SigningIdentity
	}
	return nil
}

func (m *FabricMSPConfig) GetOrganizationalUnitIdentifiers() []*FabricOUIdentifier {
	if m != nil {
		return m.OrganizationalUnitIdentifiers
	}
	return nil
}

func (m *FabricMSPConfig) GetCryptoConfig() *FabricCryptoConfig {
	if m != nil {
		return m.CryptoConfig
	}
	return nil
}

func (m *FabricMSPConfig) GetTlsRootCerts() [][]byte {
	if m != nil {
		return m.TlsRootCerts
	}
	return nil
}

func (m *FabricMSPConfig) GetTlsIntermediateCerts() [][]byte {
	if m != nil {
		return m.TlsIntermediateCerts
	}
	return nil
}

// FabricCryptoConfig contains configuration parameters
// for the cryptographic algorithms used by the MSP
// this configuration refers to
type FabricCryptoConfig struct {
	// SignatureHashFamily is a string representing the hash family to be used
	// during sign and verify operations.
	// Allowed values are "SHA2" and "SHA3".
	SignatureHashFamily string `protobuf:"bytes,1,opt,name=signature_hash_family,json=signatureHashFamily" json:"signature_hash_family,omitempty"`
	// IdentityIdentifierHashFunction is a string representing the hash function
	// to be used during the computation of the identity identifier of an MSP identity.
	// Allowed values are "SHA256", "SHA384" and "SHA3_256", "SHA3_384".
	IdentityIdentifierHashFunction string `protobuf:"bytes,2,opt,name=identity_identifier_hash_function,json=identityIdentifierHashFunction" json:"identity_identifier_hash_function,omitempty"`
}

func (m *FabricCryptoConfig) Reset()                    { *m = FabricCryptoConfig{} }
func (m *FabricCryptoConfig) String() string            { return proto.CompactTextString(m) }
func (*FabricCryptoConfig) ProtoMessage()               {}
func (*FabricCryptoConfig) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *FabricCryptoConfig) GetSignatureHashFamily() string {
	if m != nil {
		return m.SignatureHashFamily
	}
	return ""
}

func (m *FabricCryptoConfig) GetIdentityIdentifierHashFunction() string {
	if m != nil {
		return m.IdentityIdentifierHashFunction
	}
	return ""
}

// IdemixMSPConfig collects all the configuration information for
// an Idemix MSP.
type IdemixMSPConfig struct {
	// Name holds the identifier of the MSP
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// IPk represents the (serialized) issuer public key
	IPk []byte `protobuf:"bytes,2,opt,name=IPk,proto3" json:"IPk,omitempty"`
	// signer may contain crypto material to configure a default signer
	Signer *IdemixMSPSignerConfig `protobuf:"bytes,3,opt,name=signer" json:"signer,omitempty"`
}

func (m *IdemixMSPConfig) Reset()                    { *m = IdemixMSPConfig{} }
func (m *IdemixMSPConfig) String() string            { return proto.CompactTextString(m) }
func (*IdemixMSPConfig) ProtoMessage()               {}
func (*IdemixMSPConfig) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *IdemixMSPConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *IdemixMSPConfig) GetIPk() []byte {
	if m != nil {
		return m.IPk
	}
	return nil
}

func (m *IdemixMSPConfig) GetSigner() *IdemixMSPSignerConfig {
	if m != nil {
		return m.Signer
	}
	return nil
}

// IdemixMSPSIgnerConfig contains the crypto material to set up an idemix signing identity
type IdemixMSPSignerConfig struct {
	// Cred represents the serialized idemix credential of the default signer
	Cred []byte `protobuf:"bytes,1,opt,name=Cred,proto3" json:"Cred,omitempty"`
	// Sk is the secret key of the default signer, corresponding to credential Cred
	Sk []byte `protobuf:"bytes,2,opt,name=Sk,proto3" json:"Sk,omitempty"`
	// organizational_unit_identifier defines the organizational unit the default signer is in
	OrganizationalUnitIdentifier string `protobuf:"bytes,3,opt,name=organizational_unit_identifier,json=organizationalUnitIdentifier" json:"organizational_unit_identifier,omitempty"`
	// is_admin defines whether the default signer is admin or not
	IsAdmin bool `protobuf:"varint,4,opt,name=is_admin,json=isAdmin" json:"is_admin,omitempty"`
}

func (m *IdemixMSPSignerConfig) Reset()                    { *m = IdemixMSPSignerConfig{} }
func (m *IdemixMSPSignerConfig) String() string            { return proto.CompactTextString(m) }
func (*IdemixMSPSignerConfig) ProtoMessage()               {}
func (*IdemixMSPSignerConfig) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *IdemixMSPSignerConfig) GetCred() []byte {
	if m != nil {
		return m.Cred
	}
	return nil
}

func (m *IdemixMSPSignerConfig) GetSk() []byte {
	if m != nil {
		return m.Sk
	}
	return nil
}

func (m *IdemixMSPSignerConfig) GetOrganizationalUnitIdentifier() string {
	if m != nil {
		return m.OrganizationalUnitIdentifier
	}
	return ""
}

func (m *IdemixMSPSignerConfig) GetIsAdmin() bool {
	if m != nil {
		return m.IsAdmin
	}
	return false
}

// SigningIdentityInfo represents the configuration information
// related to the signing identity the peer is to use for generating
// endorsements
type SigningIdentityInfo struct {
	// PublicSigner carries the public information of the signing
	// identity. For an X.509 provider this would be represented by
	// an X.509 certificate
	PublicSigner []byte `protobuf:"bytes,1,opt,name=public_signer,json=publicSigner,proto3" json:"public_signer,omitempty"`
	// PrivateSigner denotes a reference to the private key of the
	// peer's signing identity
	PrivateSigner *KeyInfo `protobuf:"bytes,2,opt,name=private_signer,json=privateSigner" json:"private_signer,omitempty"`
}

func (m *SigningIdentityInfo) Reset()                    { *m = SigningIdentityInfo{} }
func (m *SigningIdentityInfo) String() string            { return proto.CompactTextString(m) }
func (*SigningIdentityInfo) ProtoMessage()               {}
func (*SigningIdentityInfo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *SigningIdentityInfo) GetPublicSigner() []byte {
	if m != nil {
		return m.PublicSigner
	}
	return nil
}

func (m *SigningIdentityInfo) GetPrivateSigner() *KeyInfo {
	if m != nil {
		return m.PrivateSigner
	}
	return nil
}

// KeyInfo represents a (secret) key that is either already stored
// in the bccsp/keystore or key material to be imported to the
// bccsp key-store. In later versions it may contain also a
// keystore identifier
type KeyInfo struct {
	// Identifier of the key inside the default keystore; this for
	// the case of Software BCCSP as well as the HSM BCCSP would be
	// the SKI of the key
	KeyIdentifier string `protobuf:"bytes,1,opt,name=key_identifier,json=keyIdentifier" json:"key_identifier,omitempty"`
	// KeyMaterial (optional) for the key to be imported; this is
	// properly encoded key bytes, prefixed by the type of the key
	KeyMaterial []byte `protobuf:"bytes,2,opt,name=key_material,json=keyMaterial,proto3" json:"key_material,omitempty"`
}

func (m *KeyInfo) Reset()                    { *m = KeyInfo{} }
func (m *KeyInfo) String() string            { return proto.CompactTextString(m) }
func (*KeyInfo) ProtoMessage()               {}
func (*KeyInfo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *KeyInfo) GetKeyIdentifier() string {
	if m != nil {
		return m.KeyIdentifier
	}
	return ""
}

func (m *KeyInfo) GetKeyMaterial() []byte {
	if m != nil {
		return m.KeyMaterial
	}
	return nil
}

// FabricOUIdentifier represents an organizational unit and
// its related chain of trust identifier.
type FabricOUIdentifier struct {
	// Certificate represents the second certificate in a certification chain.
	// (Notice that the first certificate in a certification chain is supposed
	// to be the certificate of an identity).
	// It must correspond to the certificate of root or intermediate CA
	// recognized by the MSP this message belongs to.
	// Starting from this certificate, a certification chain is computed
	// and boud to the OrganizationUnitIdentifier specified
	Certificate []byte `protobuf:"bytes,1,opt,name=certificate,proto3" json:"certificate,omitempty"`
	// OrganizationUnitIdentifier defines the organizational unit under the
	// MSP identified with MSPIdentifier
	OrganizationalUnitIdentifier string `protobuf:"bytes,2,opt,name=organizational_unit_identifier,json=organizationalUnitIdentifier" json:"organizational_unit_identifier,omitempty"`
}

func (m *FabricOUIdentifier) Reset()                    { *m = FabricOUIdentifier{} }
func (m *FabricOUIdentifier) String() string            { return proto.CompactTextString(m) }
func (*FabricOUIdentifier) ProtoMessage()               {}
func (*FabricOUIdentifier) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *FabricOUIdentifier) GetCertificate() []byte {
	if m != nil {
		return m.Certificate
	}
	return nil
}

func (m *FabricOUIdentifier) GetOrganizationalUnitIdentifier() string {
	if m != nil {
		return m.OrganizationalUnitIdentifier
	}
	return ""
}

func init() {
	proto.RegisterType((*MSPConfig)(nil), "msp.MSPConfig")
	proto.RegisterType((*FabricMSPConfig)(nil), "msp.FabricMSPConfig")
	proto.RegisterType((*FabricCryptoConfig)(nil), "msp.FabricCryptoConfig")
	proto.RegisterType((*IdemixMSPConfig)(nil), "msp.IdemixMSPConfig")
	proto.RegisterType((*IdemixMSPSignerConfig)(nil), "msp.IdemixMSPSignerConfig")
	proto.RegisterType((*SigningIdentityInfo)(nil), "msp.SigningIdentityInfo")
	proto.RegisterType((*KeyInfo)(nil), "msp.KeyInfo")
	proto.RegisterType((*FabricOUIdentifier)(nil), "msp.FabricOUIdentifier")
}

func init() { proto.RegisterFile("msp/msp_config.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 691 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4d, 0x6f, 0xd3, 0x4a,
	0x14, 0x55, 0xe2, 0x36, 0x6d, 0x6e, 0x9c, 0xa4, 0x6f, 0xfa, 0xf1, 0xfc, 0x9e, 0x68, 0x71, 0x0d,
	0x08, 0x6f, 0x48, 0xa4, 0x14, 0x89, 0x0d, 0x1b, 0x1a, 0x54, 0x61, 0x41, 0x45, 0xe5, 0xa8, 0x1b,
	0x36, 0xd6, 0xc4, 0x99, 0x38, 0x23, 0x7f, 0x6a, 0x66, 0x52, 0x11, 0xc4, 0xbf, 0x60, 0xcb, 0xcf,
	0xe4, 0x07, 0xa0, 0xf9, 0x48, 0xea, 0x42, 0x55, 0xd8, 0xdd, 0xb9, 0xe7, 0xdc, 0x3b, 0x77, 0xce,
	0x3d, 0x36, 0x1c, 0xe4, 0xbc, 0x1a, 0xe6, 0xbc, 0x8a, 0xe2, 0xb2, 0x98, 0xd3, 0x64, 0x50, 0xb1,
	0x52, 0x94, 0xc8, 0xca, 0x79, 0xe5, 0xbd, 0x82, 0xf6, 0xe5, 0xe4, 0x6a, 0xac, 0xf2, 0x08, 0xc1,
	0x96, 0x58, 0x55, 0xc4, 0x69, 0xb8, 0x0d, 0x7f, 0x3b, 0x54, 0x31, 0x3a, 0x82, 0x96, 0xae, 0x72,
	0x9a, 0x6e, 0xc3, 0xb7, 0x43, 0x73, 0xf2, 0x7e, 0x58, 0xd0, 0xbf, 0xc0, 0x53, 0x46, 0xe3, 0x3b,
	0xf5, 0x05, 0xce, 0x75, 0x7d, 0x3b, 0x54, 0x31, 0x3a, 0x06, 0x60, 0x65, 0x29, 0xa2, 0x98, 0x30,
	0xc1, 0x9d, 0xa6, 0x6b, 0xf9, 0x76, 0xd8, 0x96, 0x99, 0xb1, 0x4c, 0xa0, 0x17, 0x80, 0x68, 0x21,
	0x08, 0xcb, 0xc9, 0x8c, 0x62, 0x41, 0x0c, 0xcd, 0x52, 0xb4, 0x7f, 0xea, 0x88, 0xa6, 0x1f, 0x41,
	0x0b, 0xcf, 0x72, 0x5a, 0x70, 0x67, 0x4b, 0x51, 0xcc, 0x09, 0x3d, 0x87, 0x3e, 0x23, 0x37, 0x65,
	0x8c, 0x05, 0x2d, 0x8b, 0x28, 0xa3, 0x5c, 0x38, 0xdb, 0x8a, 0xd0, 0xbb, 0x4d, 0x7f, 0xa0, 0x5c,
	0xa0, 0x31, 0xec, 0x71, 0x9a, 0x14, 0xb4, 0x48, 0x22, 0x3a, 0x23, 0x85, 0xa0, 0x62, 0xe5, 0xb4,
	0xdc, 0x86, 0xdf, 0x19, 0x39, 0x83, 0x9c, 0x57, 0x83, 0x89, 0x06, 0x03, 0x83, 0x05, 0xc5, 0xbc,
	0x0c, 0xfb, 0xfc, 0x6e, 0x12, 0x45, 0xf0, 0xb8, 0x64, 0x09, 0x2e, 0xe8, 0x17, 0xd5, 0x18, 0x67,
	0xd1, 0xb2, 0xa0, 0xc2, 0x34, 0x9c, 0x53, 0xc2, 0xb8, 0xb3, 0xe3, 0x5a, 0x7e, 0x67, 0xf4, 0xaf,
	0xea, 0xa9, 0x65, 0xfa, 0x78, 0x1d, 0x6c, 0xf0, 0xf0, 0xf8, 0x6e, 0xfd, 0x75, 0x41, 0xc5, 0x2d,
	0xca, 0xd1, 0x6b, 0xe8, 0xc6, 0x6c, 0x55, 0x89, 0xd2, 0x6c, 0xcc, 0xd9, 0x55, 0x23, 0xd6, 0xdb,
	0x8d, 0x15, 0xae, 0x85, 0x0f, 0xed, 0xb8, 0x76, 0x42, 0x4f, 0xa1, 0x27, 0x32, 0x1e, 0xd5, 0x64,
	0x6f, 0x2b, 0x2d, 0x6c, 0x91, 0xf1, 0x70, 0xa3, 0xfc, 0x4b, 0x38, 0x92, 0xac, 0x7b, 0xd4, 0x07,
	0xc5, 0x3e, 0x10, 0x19, 0x0f, 0x7e, 0x5d, 0x80, 0xf7, 0xad, 0x01, 0xe8, 0xf7, 0x01, 0xd0, 0x08,
	0x0e, 0xa5, 0x48, 0x58, 0x2c, 0x19, 0x89, 0x16, 0x98, 0x2f, 0xa2, 0x39, 0xce, 0x69, 0xb6, 0x32,
	0x56, 0xd8, 0xdf, 0x80, 0xef, 0x30, 0x5f, 0x5c, 0x28, 0x08, 0x05, 0x70, 0xba, 0x5e, 0x41, 0x4d,
	0x3a, 0x53, 0xbd, 0x2c, 0x62, 0x29, 0x8d, 0x32, 0x5d, 0x3b, 0x3c, 0x59, 0x13, 0x6f, 0x45, 0x52,
	0x8d, 0x0c, 0xcb, 0x4b, 0xa1, 0x1f, 0xcc, 0x48, 0x4e, 0x3f, 0x3f, 0xec, 0xc5, 0x3d, 0xb0, 0x82,
	0xab, 0xd4, 0x18, 0x59, 0x86, 0x68, 0x04, 0x2d, 0x39, 0x1a, 0x61, 0x8e, 0xa5, 0x14, 0xfe, 0x5f,
	0x29, 0xbc, 0xe9, 0x35, 0x51, 0x98, 0x11, 0xd9, 0x30, 0xbd, 0xef, 0x0d, 0x38, 0xbc, 0x97, 0x21,
	0xef, 0x1c, 0x33, 0x32, 0x53, 0x77, 0xda, 0xa1, 0x8a, 0x51, 0x0f, 0x9a, 0x93, 0xf5, 0x95, 0xcd,
	0x49, 0x8a, 0xde, 0xc2, 0xc9, 0xc3, 0xde, 0x51, 0x93, 0xb4, 0xc3, 0x47, 0x0f, 0x39, 0x04, 0xfd,
	0x07, 0xbb, 0x94, 0x47, 0xca, 0xfc, 0xce, 0x96, 0xdb, 0xf0, 0x77, 0xc3, 0x1d, 0xca, 0xdf, 0xc8,
	0xa3, 0x57, 0xc2, 0xfe, 0x3d, 0x26, 0x46, 0x4f, 0xa0, 0x5b, 0x2d, 0xa7, 0x19, 0x8d, 0x23, 0xf3,
	0x60, 0x3d, 0xa4, 0xad, 0x93, 0xfa, 0x19, 0xe8, 0x0c, 0x7a, 0x15, 0xa3, 0x37, 0xd2, 0x0a, 0x86,
	0xd5, 0x54, 0xb2, 0xd8, 0x4a, 0x96, 0xf7, 0x44, 0x7f, 0x0f, 0x5d, 0xc3, 0xd1, 0x45, 0xde, 0x04,
	0x76, 0x0c, 0x82, 0x9e, 0x41, 0x2f, 0x25, 0xf5, 0x6d, 0x1a, 0xf9, 0xbb, 0x29, 0xa9, 0xad, 0x0e,
	0x9d, 0x82, 0x2d, 0x69, 0x39, 0x16, 0x84, 0x51, 0x9c, 0x19, 0x75, 0x3a, 0x29, 0x59, 0x5d, 0x9a,
	0x94, 0xf7, 0x75, 0x6d, 0xb3, 0xfa, 0x67, 0x83, 0x5c, 0xe8, 0x48, 0x8b, 0xd2, 0x39, 0x8d, 0xb1,
	0x20, 0xe6, 0x09, 0xf5, 0xd4, 0x5f, 0xc8, 0xdb, 0xfc, 0xb3, 0xbc, 0xe7, 0x11, 0x9c, 0x96, 0x2c,
	0x19, 0x2c, 0x56, 0x15, 0x61, 0x19, 0x99, 0x25, 0x84, 0x0d, 0xe6, 0x6a, 0x1a, 0xfd, 0xeb, 0xe4,
	0x52, 0x8e, 0xf3, 0xbd, 0x4b, 0x5e, 0xe9, 0xc5, 0x5f, 0xe1, 0x38, 0xc5, 0x09, 0xf9, 0xe4, 0x27,
	0x54, 0x2c, 0x96, 0xd3, 0x41, 0x5c, 0xe6, 0xc3, 0x5a, 0xed, 0x50, 0xd7, 0x0e, 0x75, 0xad, 0xfc,
	0x11, 0x4f, 0x5b, 0x2a, 0x3e, 0xfb, 0x19, 0x00, 0x00, 0xff, 0xff, 0x38, 0x5b, 0xfe, 0x15, 0x9a,
	0x05, 0x00, 0x00,
}
