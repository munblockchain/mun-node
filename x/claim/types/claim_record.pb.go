// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: claim/v1beta1/claim_record.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Action int32

const (
	ActionInitialClaim  Action = 0
	ActionDelegateStake Action = 1
	ActionVote          Action = 2
	ActionSwap          Action = 3
)

var Action_name = map[int32]string{
	0: "ActionInitialClaim",
	1: "ActionDelegateStake",
	2: "ActionVote",
	3: "ActionSwap",
}

var Action_value = map[string]int32{
	"ActionInitialClaim":  0,
	"ActionDelegateStake": 1,
	"ActionVote":          2,
	"ActionSwap":          3,
}

func (x Action) String() string {
	return proto.EnumName(Action_name, int32(x))
}

func (Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_39a8735094063ef9, []int{0}
}

type ClaimRecord struct {
	// address of claim user
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
	// total initial claimable amount for the user
	InitialClaimableAmount github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=initial_claimable_amount,json=initialClaimableAmount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"initial_claimable_amount" yaml:"initial_claimable_amount"`
	// true if action is completed
	// index of bool in array refers to action enum #
	ActionCompleted []bool `protobuf:"varint,3,rep,packed,name=action_completed,json=actionCompleted,proto3" json:"action_completed,omitempty" yaml:"action_completed"`
	// true if action is ready to claim
	// index of bool in array refers to action enum #
	ActionReady []bool `protobuf:"varint,4,rep,packed,name=action_ready,json=actionReady,proto3" json:"action_ready,omitempty" yaml:"action_ready"`
}

func (m *ClaimRecord) Reset()         { *m = ClaimRecord{} }
func (m *ClaimRecord) String() string { return proto.CompactTextString(m) }
func (*ClaimRecord) ProtoMessage()    {}
func (*ClaimRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_39a8735094063ef9, []int{0}
}
func (m *ClaimRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClaimRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClaimRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClaimRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClaimRecord.Merge(m, src)
}
func (m *ClaimRecord) XXX_Size() int {
	return m.Size()
}
func (m *ClaimRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_ClaimRecord.DiscardUnknown(m)
}

var xxx_messageInfo_ClaimRecord proto.InternalMessageInfo

func (m *ClaimRecord) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ClaimRecord) GetInitialClaimableAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.InitialClaimableAmount
	}
	return nil
}

func (m *ClaimRecord) GetActionCompleted() []bool {
	if m != nil {
		return m.ActionCompleted
	}
	return nil
}

func (m *ClaimRecord) GetActionReady() []bool {
	if m != nil {
		return m.ActionReady
	}
	return nil
}

func init() {
	proto.RegisterEnum("mun.claim.v1beta1.Action", Action_name, Action_value)
	proto.RegisterType((*ClaimRecord)(nil), "mun.claim.v1beta1.ClaimRecord")
}

func init() { proto.RegisterFile("claim/v1beta1/claim_record.proto", fileDescriptor_39a8735094063ef9) }

var fileDescriptor_39a8735094063ef9 = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xc1, 0x6e, 0xd3, 0x30,
	0x1c, 0xc6, 0x93, 0x76, 0x1a, 0xe0, 0xa2, 0x11, 0x3c, 0xb4, 0x85, 0x22, 0x25, 0x55, 0x4e, 0x15,
	0x62, 0x89, 0x06, 0x37, 0x0e, 0x48, 0x6b, 0xb9, 0x8c, 0x63, 0x2a, 0x71, 0xe0, 0x12, 0x39, 0x8e,
	0xd5, 0x59, 0x8d, 0xfd, 0xaf, 0x62, 0x17, 0xe8, 0x1b, 0x70, 0xe4, 0x1d, 0xb8, 0x20, 0x9e, 0x64,
	0xc7, 0x1d, 0x39, 0x15, 0xd4, 0xbe, 0x41, 0x9f, 0x00, 0xc5, 0x76, 0xb4, 0x81, 0xc4, 0xc9, 0xfe,
	0xfe, 0xdf, 0xa7, 0x9f, 0x3e, 0x5b, 0x7f, 0x34, 0xa2, 0x35, 0xe1, 0x22, 0xfb, 0x78, 0x5e, 0x32,
	0x4d, 0xce, 0x33, 0xa3, 0x8a, 0x86, 0x51, 0x68, 0xaa, 0x74, 0xd9, 0x80, 0x06, 0xfc, 0x58, 0xac,
	0x64, 0x6a, 0xe6, 0xa9, 0x4b, 0x0d, 0x9f, 0xcc, 0x61, 0x0e, 0xc6, 0xcd, 0xda, 0x9b, 0x0d, 0x0e,
	0x23, 0x0a, 0x4a, 0x80, 0xca, 0x4a, 0xa2, 0xd8, 0x2d, 0x10, 0xb8, 0xb4, 0x7e, 0xb2, 0xef, 0xa1,
	0xc1, 0xb4, 0xe5, 0xe4, 0x06, 0x8f, 0x5f, 0xa0, 0x7b, 0xa4, 0xaa, 0x1a, 0xa6, 0x54, 0xe8, 0x8f,
	0xfc, 0xf1, 0x83, 0x09, 0xde, 0x6f, 0xe2, 0xa3, 0x35, 0x11, 0xf5, 0xeb, 0xc4, 0x19, 0x49, 0xde,
	0x45, 0xf0, 0x77, 0x1f, 0x85, 0x5c, 0x72, 0xcd, 0x49, 0x5d, 0x98, 0x36, 0xa4, 0xac, 0x59, 0x41,
	0x04, 0xac, 0xa4, 0x0e, 0x7b, 0xa3, 0xfe, 0x78, 0xf0, 0xf2, 0x69, 0x6a, 0x1b, 0xa4, 0x6d, 0x83,
	0xae, 0x6c, 0x3a, 0x05, 0x2e, 0x27, 0xb3, 0xeb, 0x4d, 0xec, 0xed, 0x37, 0x71, 0x6c, 0xf1, 0xff,
	0x03, 0x25, 0x3f, 0x7e, 0xc5, 0xe3, 0x39, 0xd7, 0x57, 0xab, 0x32, 0xa5, 0x20, 0x32, 0xf7, 0x22,
	0x7b, 0x9c, 0xa9, 0x6a, 0x91, 0xe9, 0xf5, 0x92, 0x29, 0xc3, 0x54, 0xf9, 0x89, 0xc3, 0x4c, 0x3b,
	0xca, 0x85, 0x81, 0xe0, 0x77, 0x28, 0x20, 0x54, 0x73, 0x90, 0x05, 0x05, 0xb1, 0xac, 0x99, 0x66,
	0x55, 0xd8, 0x1f, 0xf5, 0xc7, 0xf7, 0x27, 0xb1, 0xab, 0x71, 0xea, 0x5e, 0xf9, 0x4f, 0x2a, 0xc9,
	0x1f, 0xd9, 0xd1, 0xb4, 0x9b, 0xe0, 0x37, 0xe8, 0xa1, 0x4b, 0x35, 0x8c, 0x54, 0xeb, 0xf0, 0xc0,
	0x70, 0x9e, 0x39, 0xce, 0xf1, 0x5f, 0x1c, 0x93, 0x48, 0xf2, 0x81, 0x95, 0x79, 0xab, 0x9e, 0x17,
	0xe8, 0xf0, 0xc2, 0x48, 0x7c, 0x82, 0xb0, 0xbd, 0x5d, 0xde, 0x69, 0x1d, 0x78, 0xf8, 0x14, 0x1d,
	0xdb, 0xf9, 0x5b, 0x56, 0xb3, 0x39, 0xd1, 0x6c, 0xa6, 0xc9, 0x82, 0x05, 0x3e, 0x3e, 0x42, 0xc8,
	0x1a, 0xef, 0x41, 0xb3, 0xa0, 0x77, 0xab, 0x67, 0x9f, 0xc8, 0x32, 0xe8, 0x0f, 0x0f, 0xbe, 0x7c,
	0x8b, 0xbc, 0xc9, 0xe5, 0xf5, 0x36, 0xf2, 0x6f, 0xb6, 0x91, 0xff, 0x7b, 0x1b, 0xf9, 0x5f, 0x77,
	0x91, 0x77, 0xb3, 0x8b, 0xbc, 0x9f, 0xbb, 0xc8, 0xfb, 0x90, 0xdd, 0xf9, 0x48, 0xb1, 0x92, 0x65,
	0x0d, 0x74, 0x41, 0xaf, 0x08, 0x97, 0xad, 0x3a, 0x93, 0x50, 0xb1, 0xec, 0xb3, 0x5d, 0x38, 0xfb,
	0xab, 0xe5, 0xa1, 0xd9, 0x93, 0x57, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xb7, 0x33, 0x7f, 0x3d,
	0x94, 0x02, 0x00, 0x00,
}

func (m *ClaimRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClaimRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClaimRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ActionReady) > 0 {
		for iNdEx := len(m.ActionReady) - 1; iNdEx >= 0; iNdEx-- {
			i--
			if m.ActionReady[iNdEx] {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
		}
		i = encodeVarintClaimRecord(dAtA, i, uint64(len(m.ActionReady)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ActionCompleted) > 0 {
		for iNdEx := len(m.ActionCompleted) - 1; iNdEx >= 0; iNdEx-- {
			i--
			if m.ActionCompleted[iNdEx] {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
		}
		i = encodeVarintClaimRecord(dAtA, i, uint64(len(m.ActionCompleted)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.InitialClaimableAmount) > 0 {
		for iNdEx := len(m.InitialClaimableAmount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.InitialClaimableAmount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintClaimRecord(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintClaimRecord(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintClaimRecord(dAtA []byte, offset int, v uint64) int {
	offset -= sovClaimRecord(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ClaimRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovClaimRecord(uint64(l))
	}
	if len(m.InitialClaimableAmount) > 0 {
		for _, e := range m.InitialClaimableAmount {
			l = e.Size()
			n += 1 + l + sovClaimRecord(uint64(l))
		}
	}
	if len(m.ActionCompleted) > 0 {
		n += 1 + sovClaimRecord(uint64(len(m.ActionCompleted))) + len(m.ActionCompleted)*1
	}
	if len(m.ActionReady) > 0 {
		n += 1 + sovClaimRecord(uint64(len(m.ActionReady))) + len(m.ActionReady)*1
	}
	return n
}

func sovClaimRecord(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozClaimRecord(x uint64) (n int) {
	return sovClaimRecord(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClaimRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClaimRecord
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ClaimRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClaimRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaimRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClaimRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaimRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialClaimableAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaimRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthClaimRecord
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClaimRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InitialClaimableAmount = append(m.InitialClaimableAmount, types.Coin{})
			if err := m.InitialClaimableAmount[len(m.InitialClaimableAmount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType == 0 {
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClaimRecord
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.ActionCompleted = append(m.ActionCompleted, bool(v != 0))
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClaimRecord
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthClaimRecord
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthClaimRecord
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				elementCount = packedLen
				if elementCount != 0 && len(m.ActionCompleted) == 0 {
					m.ActionCompleted = make([]bool, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowClaimRecord
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.ActionCompleted = append(m.ActionCompleted, bool(v != 0))
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field ActionCompleted", wireType)
			}
		case 4:
			if wireType == 0 {
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClaimRecord
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.ActionReady = append(m.ActionReady, bool(v != 0))
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClaimRecord
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthClaimRecord
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthClaimRecord
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				elementCount = packedLen
				if elementCount != 0 && len(m.ActionReady) == 0 {
					m.ActionReady = make([]bool, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowClaimRecord
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.ActionReady = append(m.ActionReady, bool(v != 0))
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field ActionReady", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipClaimRecord(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClaimRecord
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
func skipClaimRecord(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClaimRecord
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
					return 0, ErrIntOverflowClaimRecord
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowClaimRecord
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
			if length < 0 {
				return 0, ErrInvalidLengthClaimRecord
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupClaimRecord
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthClaimRecord
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthClaimRecord        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClaimRecord          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupClaimRecord = fmt.Errorf("proto: unexpected end of group")
)