// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: adapter/config/config.proto

/*
	Package config is a generated protocol buffer package.

	$title: Apigee istio-mixer-adapter
	$description: This adapter for Istio's Mixer brings Apigee's distributed policy checks and analytics to Istio.
	$location: https://github.com/apigee/istio-mixer-adapter

	The Apigee istio-mixer-adapter provides Apigee's distributed authentication and quota policy checks
	as well as the ingestion of Istio telemetry for analysis and reporting.
	For additional information or support please contact anchor-prega-support@google.com.

	This adapter supports the [authorization template](https://istio.io/docs/reference/config/policy-and-telemetry/templates/authorization/).
	and Apigee's [analytics template](https://istio.io/docs/reference/config/policy-and-telemetry/templates/authorization/).

	Example config:

	```yaml
	apiVersion: config.istio.io/v1alpha2
	kind: apigee
	metadata:
	 name: apigee-handler
	 namespace: istio-system
	spec:
	 apigee_base: https://istioservices.apigee.net/edgemicro
	 customer_base: https://myorg-test.apigee.net/istio-auth
	 org_name: myorg
	 env_name: test
	 key: 5f1132b7ff037fa187463c324d029ca26de28b7279df0ea161
	 secret: fa147e8afc35219b7e1db688c609196923f663b5e835975
	 temp_dir: "/tmp/apigee-istio"
	 server_timeout_secs: 60
	 products:
	   refresh_rate_mins: 1
	 analytics:
	   legacy_endpoint: false
	   file_limit: 1024
	 api_key_claim:
	```

	It is generated from these files:
		adapter/config/config.proto

	It has these top-level messages:
		Params
*/
package config

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// The Configuration for the Apigee adapter provides information on how the adapter should contact
// the Apigee proxies and how it should operate. Running the `apigee-istio provision` CLI command
// will ensure that all proxies are installed into your Apigee environment and generate this file
// with all required settings for you.
// For additional information on this adapter or support please contact anchor-prega-support@google.com.
type Params struct {
	// Apigee Base is the URI for a shared proxy on Apigee.
	// Required.
	ApigeeBase string `protobuf:"bytes,1,opt,name=apigee_base,json=apigeeBase,proto3" json:"apigee_base,omitempty"`
	// Customer Base is the URI for an organization-specific proxy on Apigee.
	// Required.
	CustomerBase string `protobuf:"bytes,2,opt,name=customer_base,json=customerBase,proto3" json:"customer_base,omitempty"`
	// Org Name is the name of the organization on Apigee.
	// Required.
	OrgName string `protobuf:"bytes,3,opt,name=org_name,json=orgName,proto3" json:"org_name,omitempty"`
	// Env Name is the name of the environment on Apigee.
	// Required.
	EnvName string `protobuf:"bytes,4,opt,name=env_name,json=envName,proto3" json:"env_name,omitempty"`
	// Key is used to authenticate to the Apigee proxy endpoints, generated during provisioning.
	// Required.
	Key string `protobuf:"bytes,5,opt,name=key,proto3" json:"key,omitempty"`
	// Secret is used to authenticate to the Apigee proxy endpoints, generated during provisioning.
	// Required.
	Secret string `protobuf:"bytes,6,opt,name=secret,proto3" json:"secret,omitempty"`
	// The local directory to be used by the adapter for temporary files.
	// Optional. Default: "/tmp/apigee-istio".
	TempDir string `protobuf:"bytes,7,opt,name=temp_dir,json=tempDir,proto3" json:"temp_dir,omitempty"`
	// The http timeout to be used for adapter connections to services.
	// Optional. Default: 60.
	ServerTimeoutSecs int64 `protobuf:"varint,8,opt,name=server_timeout_secs,json=serverTimeoutSecs,proto3" json:"server_timeout_secs,omitempty"`
	// The name of a JWT claim from which to look for an api_key.
	// Optional. Default: none.
	ApiKeyClaim string `protobuf:"bytes,9,opt,name=api_key_claim,json=apiKeyClaim,proto3" json:"api_key_claim,omitempty"`
	// Options specific to to products handling.
	Products *ParamsProductOptions `protobuf:"bytes,15,opt,name=products" json:"products,omitempty"`
	// Options specific to to analytics handling.
	Analytics *ParamsAnalyticsOptions `protobuf:"bytes,16,opt,name=analytics" json:"analytics,omitempty"`
}

func (m *Params) Reset()                    { *m = Params{} }
func (*Params) ProtoMessage()               {}
func (*Params) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{0} }

// Options specific to to products handling.
type ParamsProductOptions struct {
	// The rate at which the list of products is refreshed from Apigee in minutes.
	// Optional. Default: 1.
	RefreshRateMins int64 `protobuf:"varint,1,opt,name=refresh_rate_mins,json=refreshRateMins,proto3" json:"refresh_rate_mins,omitempty"`
}

func (m *ParamsProductOptions) Reset()                    { *m = ParamsProductOptions{} }
func (*ParamsProductOptions) ProtoMessage()               {}
func (*ParamsProductOptions) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{0, 0} }

// Options specific to to analytics handling.
type ParamsAnalyticsOptions struct {
	// If true, use legacy direct communication analytics protocol instead of buffering. Must be true for OPDK.
	// Optional. Default: false.
	LegacyEndpoint bool `protobuf:"varint,1,opt,name=legacy_endpoint,json=legacyEndpoint,proto3" json:"legacy_endpoint,omitempty"`
	// The number of analytics files that can be buffered before oldest files are dropped.
	// Optional. Default: 1024.
	FileLimit int64 `protobuf:"varint,2,opt,name=file_limit,json=fileLimit,proto3" json:"file_limit,omitempty"`
}

func (m *ParamsAnalyticsOptions) Reset()                    { *m = ParamsAnalyticsOptions{} }
func (*ParamsAnalyticsOptions) ProtoMessage()               {}
func (*ParamsAnalyticsOptions) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{0, 1} }

func init() {
	proto.RegisterType((*Params)(nil), "config.Params")
	proto.RegisterType((*ParamsProductOptions)(nil), "config.Params.product_options")
	proto.RegisterType((*ParamsAnalyticsOptions)(nil), "config.Params.analytics_options")
}
func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ApigeeBase) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.ApigeeBase)))
		i += copy(dAtA[i:], m.ApigeeBase)
	}
	if len(m.CustomerBase) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.CustomerBase)))
		i += copy(dAtA[i:], m.CustomerBase)
	}
	if len(m.OrgName) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.OrgName)))
		i += copy(dAtA[i:], m.OrgName)
	}
	if len(m.EnvName) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.EnvName)))
		i += copy(dAtA[i:], m.EnvName)
	}
	if len(m.Key) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.Key)))
		i += copy(dAtA[i:], m.Key)
	}
	if len(m.Secret) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.Secret)))
		i += copy(dAtA[i:], m.Secret)
	}
	if len(m.TempDir) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.TempDir)))
		i += copy(dAtA[i:], m.TempDir)
	}
	if m.ServerTimeoutSecs != 0 {
		dAtA[i] = 0x40
		i++
		i = encodeVarintConfig(dAtA, i, uint64(m.ServerTimeoutSecs))
	}
	if len(m.ApiKeyClaim) > 0 {
		dAtA[i] = 0x4a
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.ApiKeyClaim)))
		i += copy(dAtA[i:], m.ApiKeyClaim)
	}
	if m.Products != nil {
		dAtA[i] = 0x7a
		i++
		i = encodeVarintConfig(dAtA, i, uint64(m.Products.Size()))
		n1, err := m.Products.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Analytics != nil {
		dAtA[i] = 0x82
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintConfig(dAtA, i, uint64(m.Analytics.Size()))
		n2, err := m.Analytics.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *ParamsProductOptions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ParamsProductOptions) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.RefreshRateMins != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintConfig(dAtA, i, uint64(m.RefreshRateMins))
	}
	return i, nil
}

func (m *ParamsAnalyticsOptions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ParamsAnalyticsOptions) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.LegacyEndpoint {
		dAtA[i] = 0x8
		i++
		if m.LegacyEndpoint {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.FileLimit != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintConfig(dAtA, i, uint64(m.FileLimit))
	}
	return i, nil
}

func encodeVarintConfig(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Params) Size() (n int) {
	var l int
	_ = l
	l = len(m.ApigeeBase)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.CustomerBase)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.OrgName)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.EnvName)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.Secret)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.TempDir)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.ServerTimeoutSecs != 0 {
		n += 1 + sovConfig(uint64(m.ServerTimeoutSecs))
	}
	l = len(m.ApiKeyClaim)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.Products != nil {
		l = m.Products.Size()
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.Analytics != nil {
		l = m.Analytics.Size()
		n += 2 + l + sovConfig(uint64(l))
	}
	return n
}

func (m *ParamsProductOptions) Size() (n int) {
	var l int
	_ = l
	if m.RefreshRateMins != 0 {
		n += 1 + sovConfig(uint64(m.RefreshRateMins))
	}
	return n
}

func (m *ParamsAnalyticsOptions) Size() (n int) {
	var l int
	_ = l
	if m.LegacyEndpoint {
		n += 2
	}
	if m.FileLimit != 0 {
		n += 1 + sovConfig(uint64(m.FileLimit))
	}
	return n
}

func sovConfig(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozConfig(x uint64) (n int) {
	return sovConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Params) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Params{`,
		`ApigeeBase:` + fmt.Sprintf("%v", this.ApigeeBase) + `,`,
		`CustomerBase:` + fmt.Sprintf("%v", this.CustomerBase) + `,`,
		`OrgName:` + fmt.Sprintf("%v", this.OrgName) + `,`,
		`EnvName:` + fmt.Sprintf("%v", this.EnvName) + `,`,
		`Key:` + fmt.Sprintf("%v", this.Key) + `,`,
		`Secret:` + fmt.Sprintf("%v", this.Secret) + `,`,
		`TempDir:` + fmt.Sprintf("%v", this.TempDir) + `,`,
		`ServerTimeoutSecs:` + fmt.Sprintf("%v", this.ServerTimeoutSecs) + `,`,
		`ApiKeyClaim:` + fmt.Sprintf("%v", this.ApiKeyClaim) + `,`,
		`Products:` + strings.Replace(fmt.Sprintf("%v", this.Products), "ParamsProductOptions", "ParamsProductOptions", 1) + `,`,
		`Analytics:` + strings.Replace(fmt.Sprintf("%v", this.Analytics), "ParamsAnalyticsOptions", "ParamsAnalyticsOptions", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ParamsProductOptions) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ParamsProductOptions{`,
		`RefreshRateMins:` + fmt.Sprintf("%v", this.RefreshRateMins) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ParamsAnalyticsOptions) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ParamsAnalyticsOptions{`,
		`LegacyEndpoint:` + fmt.Sprintf("%v", this.LegacyEndpoint) + `,`,
		`FileLimit:` + fmt.Sprintf("%v", this.FileLimit) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringConfig(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApigeeBase", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ApigeeBase = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CustomerBase", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CustomerBase = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrgName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OrgName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnvName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EnvName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Secret", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Secret = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TempDir", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TempDir = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServerTimeoutSecs", wireType)
			}
			m.ServerTimeoutSecs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ServerTimeoutSecs |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApiKeyClaim", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ApiKeyClaim = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Products", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Products == nil {
				m.Products = &ParamsProductOptions{}
			}
			if err := m.Products.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 16:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Analytics", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Analytics == nil {
				m.Analytics = &ParamsAnalyticsOptions{}
			}
			if err := m.Analytics.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ParamsProductOptions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: product_options: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: product_options: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RefreshRateMins", wireType)
			}
			m.RefreshRateMins = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RefreshRateMins |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ParamsAnalyticsOptions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: analytics_options: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: analytics_options: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LegacyEndpoint", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.LegacyEndpoint = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FileLimit", wireType)
			}
			m.FileLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FileLimit |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthConfig
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowConfig
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipConfig(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthConfig = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConfig   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("adapter/config/config.proto", fileDescriptorConfig) }

var fileDescriptorConfig = []byte{
	// 467 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xbf, 0x6e, 0x14, 0x3f,
	0x10, 0xc7, 0x77, 0x7f, 0x9b, 0xdf, 0xe6, 0x6e, 0x42, 0xb8, 0x9c, 0x41, 0x68, 0x39, 0x84, 0x39,
	0x85, 0x82, 0x13, 0xc5, 0x45, 0x82, 0x06, 0x21, 0x41, 0x11, 0xa0, 0xe2, 0x8f, 0xd0, 0x42, 0x05,
	0x85, 0xe5, 0xec, 0xcd, 0x2d, 0x56, 0xd6, 0xf6, 0xca, 0xf6, 0x9d, 0xb4, 0x1d, 0x8f, 0xc0, 0x63,
	0xf0, 0x02, 0xbc, 0x43, 0xca, 0x94, 0x94, 0xdc, 0xd2, 0x50, 0xe6, 0x11, 0xd0, 0xda, 0xcb, 0x21,
	0x85, 0xca, 0x9e, 0xcf, 0x67, 0xe6, 0x5b, 0x8c, 0x0d, 0xb7, 0xf8, 0x82, 0xd7, 0x0e, 0xcd, 0x51,
	0xa1, 0xd5, 0x52, 0x94, 0xfd, 0x31, 0xaf, 0x8d, 0x76, 0x9a, 0xa4, 0xa1, 0x9a, 0x5c, 0x2f, 0x75,
	0xa9, 0x3d, 0x3a, 0xea, 0x6e, 0xc1, 0x1e, 0x7e, 0xdb, 0x81, 0xf4, 0x2d, 0x37, 0x5c, 0x5a, 0x72,
	0x07, 0xf6, 0x78, 0x2d, 0x4a, 0x44, 0x76, 0xc2, 0x2d, 0x66, 0xf1, 0x34, 0x9e, 0x0d, 0x73, 0x08,
	0xe8, 0x98, 0x5b, 0x24, 0x77, 0x61, 0xbf, 0x58, 0x59, 0xa7, 0x25, 0x9a, 0xd0, 0xf2, 0x9f, 0x6f,
	0xb9, 0xf2, 0x07, 0xfa, 0xa6, 0x9b, 0x30, 0xd0, 0xa6, 0x64, 0x8a, 0x4b, 0xcc, 0x12, 0xef, 0x77,
	0xb5, 0x29, 0xdf, 0x70, 0xe9, 0x15, 0xaa, 0x75, 0x50, 0x3b, 0x41, 0xa1, 0x5a, 0x7b, 0x75, 0x00,
	0xc9, 0x29, 0x36, 0xd9, 0xff, 0x9e, 0x76, 0x57, 0x72, 0x03, 0x52, 0x8b, 0x85, 0x41, 0x97, 0xa5,
	0x1e, 0xf6, 0x55, 0x17, 0xe2, 0x50, 0xd6, 0x6c, 0x21, 0x4c, 0xb6, 0x1b, 0x42, 0xba, 0xfa, 0xb9,
	0x30, 0x64, 0x0e, 0xd7, 0x2c, 0x9a, 0x35, 0x1a, 0xe6, 0x84, 0x44, 0xbd, 0x72, 0xcc, 0x62, 0x61,
	0xb3, 0xc1, 0x34, 0x9e, 0x25, 0xf9, 0x38, 0xa8, 0xf7, 0xc1, 0xbc, 0xc3, 0xc2, 0x92, 0x43, 0xd8,
	0xe7, 0xb5, 0x60, 0xa7, 0xd8, 0xb0, 0xa2, 0xe2, 0x42, 0x66, 0x43, 0x9f, 0xd7, 0x6d, 0xe1, 0x25,
	0x36, 0xcf, 0x3a, 0x44, 0x1e, 0xc3, 0xa0, 0x36, 0x7a, 0xb1, 0x2a, 0x9c, 0xcd, 0x46, 0xd3, 0x78,
	0xb6, 0xf7, 0x80, 0xce, 0xfb, 0xf5, 0x86, 0xb5, 0xcd, 0x7b, 0xcd, 0x74, 0xed, 0x84, 0x56, 0x36,
	0xdf, 0xf6, 0x93, 0xa7, 0x30, 0xe4, 0x8a, 0x57, 0x8d, 0x13, 0x85, 0xcd, 0x0e, 0xfc, 0xf0, 0xf4,
	0xd2, 0xf0, 0xd6, 0x6f, 0xc7, 0xff, 0x8e, 0x4c, 0x9e, 0xc0, 0xe8, 0x52, 0x38, 0xb9, 0x0f, 0x63,
	0x83, 0x4b, 0x83, 0xf6, 0x13, 0x33, 0xdc, 0x21, 0x93, 0x42, 0x59, 0xff, 0x52, 0x49, 0x3e, 0xea,
	0x45, 0xce, 0x1d, 0xbe, 0x16, 0xca, 0x4e, 0x3e, 0xc2, 0xf8, 0x9f, 0x78, 0x72, 0x0f, 0x46, 0x15,
	0x96, 0xbc, 0x68, 0x18, 0xaa, 0x45, 0xad, 0x85, 0x72, 0x7e, 0x7c, 0x90, 0x5f, 0x0d, 0xf8, 0x45,
	0x4f, 0xc9, 0x6d, 0x80, 0xa5, 0xa8, 0x90, 0x55, 0x42, 0x0a, 0xe7, 0x5f, 0x3a, 0xc9, 0x87, 0x1d,
	0x79, 0xd5, 0x81, 0xe3, 0x47, 0x67, 0x1b, 0x1a, 0x9d, 0x6f, 0x68, 0xf4, 0x7d, 0x43, 0xa3, 0x8b,
	0x0d, 0x8d, 0x3e, 0xb7, 0x34, 0xfe, 0xda, 0xd2, 0xe8, 0xac, 0xa5, 0xf1, 0x79, 0x4b, 0xe3, 0x1f,
	0x2d, 0x8d, 0x7f, 0xb5, 0x34, 0xba, 0x68, 0x69, 0xfc, 0xe5, 0x27, 0x8d, 0x3e, 0xf4, 0xff, 0xf0,
	0x24, 0xf5, 0x1f, 0xef, 0xe1, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x41, 0xf0, 0xb2, 0x7b, 0xb5,
	0x02, 0x00, 0x00,
}
