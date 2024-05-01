// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zgc/committee/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// VoteType enumerates the valid types of a vote.
type VoteType int32

const (
	// VOTE_TYPE_UNSPECIFIED defines a no-op vote option.
	VOTE_TYPE_UNSPECIFIED VoteType = 0
	// VOTE_TYPE_YES defines a yes vote option.
	VOTE_TYPE_YES VoteType = 1
	// VOTE_TYPE_NO defines a no vote option.
	VOTE_TYPE_NO VoteType = 2
	// VOTE_TYPE_ABSTAIN defines an abstain vote option.
	VOTE_TYPE_ABSTAIN VoteType = 3
)

var VoteType_name = map[int32]string{
	0: "VOTE_TYPE_UNSPECIFIED",
	1: "VOTE_TYPE_YES",
	2: "VOTE_TYPE_NO",
	3: "VOTE_TYPE_ABSTAIN",
}

var VoteType_value = map[string]int32{
	"VOTE_TYPE_UNSPECIFIED": 0,
	"VOTE_TYPE_YES":         1,
	"VOTE_TYPE_NO":          2,
	"VOTE_TYPE_ABSTAIN":     3,
}

func (x VoteType) String() string {
	return proto.EnumName(VoteType_name, int32(x))
}

func (VoteType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dc916f377aadb716, []int{0}
}

// GenesisState defines the committee module's genesis state.
type GenesisState struct {
	NextProposalID uint64       `protobuf:"varint,1,opt,name=next_proposal_id,json=nextProposalId,proto3" json:"next_proposal_id,omitempty"`
	Committees     []*types.Any `protobuf:"bytes,2,rep,name=committees,proto3" json:"committees,omitempty"`
	Proposals      Proposals    `protobuf:"bytes,3,rep,name=proposals,proto3,castrepeated=Proposals" json:"proposals"`
	Votes          []Vote       `protobuf:"bytes,4,rep,name=votes,proto3" json:"votes"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc916f377aadb716, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

// Proposal is an internal record of a governance proposal submitted to a committee.
type Proposal struct {
	Content     *types.Any `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	ID          uint64     `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	CommitteeID uint64     `protobuf:"varint,3,opt,name=committee_id,json=committeeId,proto3" json:"committee_id,omitempty"`
	Deadline    time.Time  `protobuf:"bytes,4,opt,name=deadline,proto3,stdtime" json:"deadline"`
}

func (m *Proposal) Reset()      { *m = Proposal{} }
func (*Proposal) ProtoMessage() {}
func (*Proposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc916f377aadb716, []int{1}
}
func (m *Proposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Proposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Proposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Proposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proposal.Merge(m, src)
}
func (m *Proposal) XXX_Size() int {
	return m.Size()
}
func (m *Proposal) XXX_DiscardUnknown() {
	xxx_messageInfo_Proposal.DiscardUnknown(m)
}

var xxx_messageInfo_Proposal proto.InternalMessageInfo

// Vote is an internal record of a single governance vote.
type Vote struct {
	ProposalID uint64                                        `protobuf:"varint,1,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id,omitempty"`
	Voter      github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=voter,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"voter,omitempty"`
	VoteType   VoteType                                      `protobuf:"varint,3,opt,name=vote_type,json=voteType,proto3,enum=zgc.committee.v1beta1.VoteType" json:"vote_type,omitempty"`
}

func (m *Vote) Reset()         { *m = Vote{} }
func (m *Vote) String() string { return proto.CompactTextString(m) }
func (*Vote) ProtoMessage()    {}
func (*Vote) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc916f377aadb716, []int{2}
}
func (m *Vote) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Vote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Vote.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Vote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vote.Merge(m, src)
}
func (m *Vote) XXX_Size() int {
	return m.Size()
}
func (m *Vote) XXX_DiscardUnknown() {
	xxx_messageInfo_Vote.DiscardUnknown(m)
}

var xxx_messageInfo_Vote proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("zgc.committee.v1beta1.VoteType", VoteType_name, VoteType_value)
	proto.RegisterType((*GenesisState)(nil), "zgc.committee.v1beta1.GenesisState")
	proto.RegisterType((*Proposal)(nil), "zgc.committee.v1beta1.Proposal")
	proto.RegisterType((*Vote)(nil), "zgc.committee.v1beta1.Vote")
}

func init() {
	proto.RegisterFile("zgc/committee/v1beta1/genesis.proto", fileDescriptor_dc916f377aadb716)
}

var fileDescriptor_dc916f377aadb716 = []byte{
	// 655 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0x4f, 0x4f, 0xdb, 0x4e,
	0x10, 0xb5, 0x1d, 0xff, 0xf8, 0x25, 0x9b, 0x90, 0x86, 0x2d, 0x54, 0x21, 0x95, 0x6c, 0x44, 0x2f,
	0xa8, 0x6a, 0x6c, 0xa0, 0x87, 0x4a, 0x88, 0x43, 0xe3, 0x24, 0xb4, 0xbe, 0x84, 0xc8, 0x49, 0x91,
	0xe8, 0xa1, 0x91, 0x63, 0x6f, 0x8d, 0xd5, 0xc4, 0x1b, 0x65, 0x97, 0x88, 0xf0, 0x09, 0x38, 0x72,
	0xec, 0xb1, 0x6a, 0x6f, 0x3d, 0xf3, 0x21, 0x10, 0x27, 0xd4, 0x53, 0xa5, 0x4a, 0xa6, 0x32, 0xdf,
	0xa0, 0xc7, 0x9e, 0x2a, 0xaf, 0xff, 0x24, 0x2a, 0xe5, 0xe4, 0xdd, 0x99, 0x37, 0x6f, 0xe7, 0xbd,
	0x19, 0x19, 0x3c, 0x39, 0x75, 0x2c, 0xd5, 0xc2, 0xc3, 0xa1, 0x4b, 0x29, 0x42, 0xea, 0x64, 0xab,
	0x8f, 0xa8, 0xb9, 0xa5, 0x3a, 0xc8, 0x43, 0xc4, 0x25, 0xca, 0x68, 0x8c, 0x29, 0x86, 0x2b, 0xa7,
	0x8e, 0xa5, 0xa4, 0x20, 0x25, 0x06, 0x55, 0x56, 0x2d, 0x4c, 0x86, 0x98, 0xf4, 0x18, 0x48, 0x8d,
	0x2e, 0x51, 0x45, 0x65, 0xd9, 0xc1, 0x0e, 0x8e, 0xe2, 0xe1, 0x29, 0x8e, 0xae, 0x3a, 0x18, 0x3b,
	0x03, 0xa4, 0xb2, 0x5b, 0xff, 0xf8, 0xbd, 0x6a, 0x7a, 0xd3, 0x38, 0x25, 0xff, 0x9d, 0xa2, 0xee,
	0x10, 0x11, 0x6a, 0x0e, 0x47, 0x11, 0x60, 0xfd, 0xb3, 0x00, 0x0a, 0xaf, 0xa2, 0xae, 0x3a, 0xd4,
	0xa4, 0x08, 0xee, 0x82, 0x92, 0x87, 0x4e, 0x68, 0xf8, 0xfa, 0x08, 0x13, 0x73, 0xd0, 0x73, 0xed,
	0x32, 0xbf, 0xc6, 0x6f, 0x88, 0x1a, 0x0c, 0x7c, 0xb9, 0xd8, 0x42, 0x27, 0xb4, 0x1d, 0xa7, 0xf4,
	0x86, 0x51, 0xf4, 0xe6, 0xef, 0x36, 0xac, 0x03, 0x90, 0x0a, 0x22, 0x65, 0x61, 0x2d, 0xb3, 0x91,
	0xdf, 0x5e, 0x56, 0xa2, 0x26, 0x94, 0xa4, 0x09, 0xa5, 0xe6, 0x4d, 0xb5, 0xc5, 0xab, 0x8b, 0x6a,
	0xae, 0x9e, 0x60, 0x8d, 0xb9, 0x32, 0xd8, 0x06, 0xb9, 0xe4, 0x75, 0x52, 0xce, 0x30, 0x0e, 0x59,
	0xf9, 0xa7, 0x57, 0x4a, 0xf2, 0xb4, 0xb6, 0x74, 0xe9, 0xcb, 0xdc, 0xd7, 0x1b, 0x39, 0x97, 0x44,
	0x88, 0x31, 0x23, 0x81, 0x2f, 0xc0, 0x7f, 0x13, 0x4c, 0x11, 0x29, 0x8b, 0x8c, 0xed, 0xf1, 0x3d,
	0x6c, 0x07, 0x98, 0x22, 0x4d, 0x0c, 0x99, 0x8c, 0x08, 0xbf, 0x23, 0x9e, 0x7d, 0x92, 0xb9, 0xf5,
	0x5f, 0x3c, 0xc8, 0x26, 0xbc, 0xb0, 0x05, 0xfe, 0xb7, 0xb0, 0x47, 0x91, 0x47, 0x99, 0x2f, 0xf7,
	0xe9, 0x93, 0xae, 0x2e, 0xaa, 0x95, 0x78, 0x78, 0x0e, 0x9e, 0xa4, 0x6f, 0xd4, 0xa3, 0x5a, 0x23,
	0x21, 0x81, 0x8f, 0x80, 0xe0, 0xda, 0x65, 0x81, 0x59, 0xbc, 0x10, 0xf8, 0xb2, 0xa0, 0x37, 0x0c,
	0xc1, 0xb5, 0xe1, 0x36, 0x28, 0xa4, 0x1d, 0x86, 0x43, 0xc8, 0x30, 0xc4, 0x83, 0xc0, 0x97, 0xf3,
	0xa9, 0x6d, 0x7a, 0xc3, 0xc8, 0xa7, 0x20, 0xdd, 0x86, 0x2f, 0x41, 0xd6, 0x46, 0xa6, 0x3d, 0x70,
	0x3d, 0x54, 0x16, 0x59, 0x73, 0x95, 0x3b, 0xcd, 0x75, 0x93, 0x0d, 0xd0, 0xb2, 0xa1, 0xd2, 0xf3,
	0x1b, 0x99, 0x37, 0xd2, 0xaa, 0x9d, 0x6c, 0x28, 0xf8, 0x63, 0x28, 0xfa, 0x07, 0x0f, 0xc4, 0xd0,
	0x10, 0xa8, 0x82, 0xfc, 0xdd, 0x65, 0x28, 0x06, 0xbe, 0x0c, 0xe6, 0x16, 0x01, 0x8c, 0x66, 0x4b,
	0xf0, 0x2e, 0x72, 0x7b, 0xcc, 0x44, 0x15, 0xb4, 0xd7, 0xbf, 0x7d, 0xb9, 0xea, 0xb8, 0xf4, 0xe8,
	0xb8, 0x1f, 0x7a, 0x1e, 0x6f, 0x74, 0xfc, 0xa9, 0x12, 0xfb, 0x83, 0x4a, 0xa7, 0x23, 0x44, 0x94,
	0x9a, 0x65, 0xd5, 0x6c, 0x7b, 0x8c, 0x08, 0xf9, 0x76, 0x51, 0x7d, 0x18, 0x5b, 0x17, 0x47, 0xb4,
	0x29, 0x45, 0x24, 0x1a, 0xca, 0x18, 0xee, 0x82, 0x5c, 0x78, 0xe8, 0x85, 0x65, 0xcc, 0x96, 0xe2,
	0xbd, 0xfb, 0x11, 0x0a, 0xe8, 0x4e, 0x47, 0xc8, 0xc8, 0x4e, 0xe2, 0x53, 0x34, 0xd2, 0xa7, 0x0e,
	0xc8, 0x26, 0x39, 0xb8, 0x0a, 0x56, 0x0e, 0xf6, 0xbb, 0xcd, 0x5e, 0xf7, 0xb0, 0xdd, 0xec, 0xbd,
	0x69, 0x75, 0xda, 0xcd, 0xba, 0xbe, 0xa7, 0x37, 0x1b, 0x25, 0x0e, 0x2e, 0x81, 0xc5, 0x59, 0xea,
	0xb0, 0xd9, 0x29, 0xf1, 0xb0, 0x04, 0x0a, 0xb3, 0x50, 0x6b, 0xbf, 0x24, 0xc0, 0x15, 0xb0, 0x34,
	0x8b, 0xd4, 0xb4, 0x4e, 0xb7, 0xa6, 0xb7, 0x4a, 0x99, 0x8a, 0x78, 0xf6, 0x45, 0xe2, 0xb4, 0xbd,
	0xcb, 0x40, 0xe2, 0xaf, 0x03, 0x89, 0xff, 0x19, 0x48, 0xfc, 0xf9, 0xad, 0xc4, 0x5d, 0xdf, 0x4a,
	0xdc, 0xf7, 0x5b, 0x89, 0x7b, 0xfb, 0x6c, 0xce, 0x93, 0x4d, 0x67, 0x60, 0xf6, 0x89, 0xba, 0xe9,
	0x54, 0xad, 0x23, 0xd3, 0xf5, 0xd4, 0x93, 0xb9, 0x9f, 0x07, 0x73, 0xa7, 0xbf, 0xc0, 0x06, 0xf8,
	0xfc, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfb, 0xa9, 0xa7, 0x8e, 0x5a, 0x04, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Votes) > 0 {
		for iNdEx := len(m.Votes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Votes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Proposals) > 0 {
		for iNdEx := len(m.Proposals) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Proposals[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Committees) > 0 {
		for iNdEx := len(m.Committees) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Committees[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.NextProposalID != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.NextProposalID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Proposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Proposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Proposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Deadline, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Deadline):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintGenesis(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	if m.CommitteeID != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.CommitteeID))
		i--
		dAtA[i] = 0x18
	}
	if m.ID != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x10
	}
	if m.Content != nil {
		{
			size, err := m.Content.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Vote) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Vote) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Vote) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.VoteType != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.VoteType))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Voter) > 0 {
		i -= len(m.Voter)
		copy(dAtA[i:], m.Voter)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Voter)))
		i--
		dAtA[i] = 0x12
	}
	if m.ProposalID != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.ProposalID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NextProposalID != 0 {
		n += 1 + sovGenesis(uint64(m.NextProposalID))
	}
	if len(m.Committees) > 0 {
		for _, e := range m.Committees {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Proposals) > 0 {
		for _, e := range m.Proposals {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Votes) > 0 {
		for _, e := range m.Votes {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *Proposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Content != nil {
		l = m.Content.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.ID != 0 {
		n += 1 + sovGenesis(uint64(m.ID))
	}
	if m.CommitteeID != 0 {
		n += 1 + sovGenesis(uint64(m.CommitteeID))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Deadline)
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *Vote) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProposalID != 0 {
		n += 1 + sovGenesis(uint64(m.ProposalID))
	}
	l = len(m.Voter)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.VoteType != 0 {
		n += 1 + sovGenesis(uint64(m.VoteType))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextProposalID", wireType)
			}
			m.NextProposalID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NextProposalID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Committees", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Committees = append(m.Committees, &types.Any{})
			if err := m.Committees[len(m.Committees)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proposals", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proposals = append(m.Proposals, Proposal{})
			if err := m.Proposals[len(m.Proposals)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Votes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Votes = append(m.Votes, Vote{})
			if err := m.Votes[len(m.Votes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *Proposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: Proposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Proposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Content == nil {
				m.Content = &types.Any{}
			}
			if err := m.Content.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommitteeID", wireType)
			}
			m.CommitteeID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CommitteeID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deadline", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Deadline, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *Vote) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: Vote: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Vote: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalID", wireType)
			}
			m.ProposalID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProposalID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Voter", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Voter = append(m.Voter[:0], dAtA[iNdEx:postIndex]...)
			if m.Voter == nil {
				m.Voter = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoteType", wireType)
			}
			m.VoteType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VoteType |= VoteType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
