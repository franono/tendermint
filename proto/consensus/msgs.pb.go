// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/consensus/msgs.proto

package consensus

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	bits "github.com/franono/tendermint/proto/libs/bits"
	types "github.com/franono/tendermint/proto/types"
	math "math"
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

// NewRoundStepMessage is sent for every step taken in the ConsensusState.
// For every height/round/step transition
type NewRoundStep struct {
	Height                int64    `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Round                 int32    `protobuf:"varint,2,opt,name=round,proto3" json:"round,omitempty"`
	Step                  uint32   `protobuf:"varint,3,opt,name=step,proto3" json:"step,omitempty"`
	SecondsSinceStartTime int64    `protobuf:"varint,4,opt,name=seconds_since_start_time,json=secondsSinceStartTime,proto3" json:"seconds_since_start_time,omitempty"`
	LastCommitRound       int32    `protobuf:"varint,5,opt,name=last_commit_round,json=lastCommitRound,proto3" json:"last_commit_round,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *NewRoundStep) Reset()         { *m = NewRoundStep{} }
func (m *NewRoundStep) String() string { return proto.CompactTextString(m) }
func (*NewRoundStep) ProtoMessage()    {}
func (*NewRoundStep) Descriptor() ([]byte, []int) {
	return fileDescriptor_9de64017f8b3fc88, []int{0}
}
func (m *NewRoundStep) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewRoundStep.Unmarshal(m, b)
}
func (m *NewRoundStep) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewRoundStep.Marshal(b, m, deterministic)
}
func (m *NewRoundStep) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewRoundStep.Merge(m, src)
}
func (m *NewRoundStep) XXX_Size() int {
	return xxx_messageInfo_NewRoundStep.Size(m)
}
func (m *NewRoundStep) XXX_DiscardUnknown() {
	xxx_messageInfo_NewRoundStep.DiscardUnknown(m)
}

var xxx_messageInfo_NewRoundStep proto.InternalMessageInfo

func (m *NewRoundStep) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *NewRoundStep) GetRound() int32 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *NewRoundStep) GetStep() uint32 {
	if m != nil {
		return m.Step
	}
	return 0
}

func (m *NewRoundStep) GetSecondsSinceStartTime() int64 {
	if m != nil {
		return m.SecondsSinceStartTime
	}
	return 0
}

func (m *NewRoundStep) GetLastCommitRound() int32 {
	if m != nil {
		return m.LastCommitRound
	}
	return 0
}

// NewValidBlockMessage is sent when a validator observes a valid block B in some round r,
//i.e., there is a Proposal for block B and 2/3+ prevotes for the block B in the round r.
// In case the block is also committed, then IsCommit flag is set to true.
type NewValidBlock struct {
	Height               int64               `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Round                int32               `protobuf:"varint,2,opt,name=round,proto3" json:"round,omitempty"`
	BlockPartsHeader     types.PartSetHeader `protobuf:"bytes,3,opt,name=block_parts_header,json=blockPartsHeader,proto3" json:"block_parts_header"`
	BlockParts           *bits.BitArray      `protobuf:"bytes,4,opt,name=block_parts,json=blockParts,proto3" json:"block_parts,omitempty"`
	IsCommit             bool                `protobuf:"varint,5,opt,name=is_commit,json=isCommit,proto3" json:"is_commit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *NewValidBlock) Reset()         { *m = NewValidBlock{} }
func (m *NewValidBlock) String() string { return proto.CompactTextString(m) }
func (*NewValidBlock) ProtoMessage()    {}
func (*NewValidBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_9de64017f8b3fc88, []int{1}
}
func (m *NewValidBlock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewValidBlock.Unmarshal(m, b)
}
func (m *NewValidBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewValidBlock.Marshal(b, m, deterministic)
}
func (m *NewValidBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewValidBlock.Merge(m, src)
}
func (m *NewValidBlock) XXX_Size() int {
	return xxx_messageInfo_NewValidBlock.Size(m)
}
func (m *NewValidBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_NewValidBlock.DiscardUnknown(m)
}

var xxx_messageInfo_NewValidBlock proto.InternalMessageInfo

func (m *NewValidBlock) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *NewValidBlock) GetRound() int32 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *NewValidBlock) GetBlockPartsHeader() types.PartSetHeader {
	if m != nil {
		return m.BlockPartsHeader
	}
	return types.PartSetHeader{}
}

func (m *NewValidBlock) GetBlockParts() *bits.BitArray {
	if m != nil {
		return m.BlockParts
	}
	return nil
}

func (m *NewValidBlock) GetIsCommit() bool {
	if m != nil {
		return m.IsCommit
	}
	return false
}

// ProposalMessage is sent when a new block is proposed.
type Proposal struct {
	Proposal             types.Proposal `protobuf:"bytes,1,opt,name=proposal,proto3" json:"proposal"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Proposal) Reset()         { *m = Proposal{} }
func (m *Proposal) String() string { return proto.CompactTextString(m) }
func (*Proposal) ProtoMessage()    {}
func (*Proposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_9de64017f8b3fc88, []int{2}
}
func (m *Proposal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Proposal.Unmarshal(m, b)
}
func (m *Proposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Proposal.Marshal(b, m, deterministic)
}
func (m *Proposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proposal.Merge(m, src)
}
func (m *Proposal) XXX_Size() int {
	return xxx_messageInfo_Proposal.Size(m)
}
func (m *Proposal) XXX_DiscardUnknown() {
	xxx_messageInfo_Proposal.DiscardUnknown(m)
}

var xxx_messageInfo_Proposal proto.InternalMessageInfo

func (m *Proposal) GetProposal() types.Proposal {
	if m != nil {
		return m.Proposal
	}
	return types.Proposal{}
}

// ProposalPOLMessage is sent when a previous proposal is re-proposed.
type ProposalPOL struct {
	Height               int64         `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	ProposalPolRound     int32         `protobuf:"varint,2,opt,name=proposal_pol_round,json=proposalPolRound,proto3" json:"proposal_pol_round,omitempty"`
	ProposalPol          bits.BitArray `protobuf:"bytes,3,opt,name=proposal_pol,json=proposalPol,proto3" json:"proposal_pol"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ProposalPOL) Reset()         { *m = ProposalPOL{} }
func (m *ProposalPOL) String() string { return proto.CompactTextString(m) }
func (*ProposalPOL) ProtoMessage()    {}
func (*ProposalPOL) Descriptor() ([]byte, []int) {
	return fileDescriptor_9de64017f8b3fc88, []int{3}
}
func (m *ProposalPOL) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProposalPOL.Unmarshal(m, b)
}
func (m *ProposalPOL) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProposalPOL.Marshal(b, m, deterministic)
}
func (m *ProposalPOL) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProposalPOL.Merge(m, src)
}
func (m *ProposalPOL) XXX_Size() int {
	return xxx_messageInfo_ProposalPOL.Size(m)
}
func (m *ProposalPOL) XXX_DiscardUnknown() {
	xxx_messageInfo_ProposalPOL.DiscardUnknown(m)
}

var xxx_messageInfo_ProposalPOL proto.InternalMessageInfo

func (m *ProposalPOL) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *ProposalPOL) GetProposalPolRound() int32 {
	if m != nil {
		return m.ProposalPolRound
	}
	return 0
}

func (m *ProposalPOL) GetProposalPol() bits.BitArray {
	if m != nil {
		return m.ProposalPol
	}
	return bits.BitArray{}
}

// BlockPartMessage is sent when gossipping a piece of the proposed block.
type BlockPart struct {
	Height               int64      `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Round                int32      `protobuf:"varint,2,opt,name=round,proto3" json:"round,omitempty"`
	Part                 types.Part `protobuf:"bytes,3,opt,name=part,proto3" json:"part"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *BlockPart) Reset()         { *m = BlockPart{} }
func (m *BlockPart) String() string { return proto.CompactTextString(m) }
func (*BlockPart) ProtoMessage()    {}
func (*BlockPart) Descriptor() ([]byte, []int) {
	return fileDescriptor_9de64017f8b3fc88, []int{4}
}
func (m *BlockPart) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockPart.Unmarshal(m, b)
}
func (m *BlockPart) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockPart.Marshal(b, m, deterministic)
}
func (m *BlockPart) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockPart.Merge(m, src)
}
func (m *BlockPart) XXX_Size() int {
	return xxx_messageInfo_BlockPart.Size(m)
}
func (m *BlockPart) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockPart.DiscardUnknown(m)
}

var xxx_messageInfo_BlockPart proto.InternalMessageInfo

func (m *BlockPart) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *BlockPart) GetRound() int32 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *BlockPart) GetPart() types.Part {
	if m != nil {
		return m.Part
	}
	return types.Part{}
}

// VoteMessage is sent when voting for a proposal (or lack thereof).
type Vote struct {
	Vote                 *types.Vote `protobuf:"bytes,1,opt,name=vote,proto3" json:"vote,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Vote) Reset()         { *m = Vote{} }
func (m *Vote) String() string { return proto.CompactTextString(m) }
func (*Vote) ProtoMessage()    {}
func (*Vote) Descriptor() ([]byte, []int) {
	return fileDescriptor_9de64017f8b3fc88, []int{5}
}
func (m *Vote) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vote.Unmarshal(m, b)
}
func (m *Vote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vote.Marshal(b, m, deterministic)
}
func (m *Vote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vote.Merge(m, src)
}
func (m *Vote) XXX_Size() int {
	return xxx_messageInfo_Vote.Size(m)
}
func (m *Vote) XXX_DiscardUnknown() {
	xxx_messageInfo_Vote.DiscardUnknown(m)
}

var xxx_messageInfo_Vote proto.InternalMessageInfo

func (m *Vote) GetVote() *types.Vote {
	if m != nil {
		return m.Vote
	}
	return nil
}

// HasVoteMessage is sent to indicate that a particular vote has been received.
type HasVote struct {
	Height               int64               `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Round                int32               `protobuf:"varint,2,opt,name=round,proto3" json:"round,omitempty"`
	Type                 types.SignedMsgType `protobuf:"varint,3,opt,name=type,proto3,enum=tendermint.proto.types.SignedMsgType" json:"type,omitempty"`
	Index                uint32              `protobuf:"varint,4,opt,name=index,proto3" json:"index,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *HasVote) Reset()         { *m = HasVote{} }
func (m *HasVote) String() string { return proto.CompactTextString(m) }
func (*HasVote) ProtoMessage()    {}
func (*HasVote) Descriptor() ([]byte, []int) {
	return fileDescriptor_9de64017f8b3fc88, []int{6}
}
func (m *HasVote) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HasVote.Unmarshal(m, b)
}
func (m *HasVote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HasVote.Marshal(b, m, deterministic)
}
func (m *HasVote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HasVote.Merge(m, src)
}
func (m *HasVote) XXX_Size() int {
	return xxx_messageInfo_HasVote.Size(m)
}
func (m *HasVote) XXX_DiscardUnknown() {
	xxx_messageInfo_HasVote.DiscardUnknown(m)
}

var xxx_messageInfo_HasVote proto.InternalMessageInfo

func (m *HasVote) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *HasVote) GetRound() int32 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *HasVote) GetType() types.SignedMsgType {
	if m != nil {
		return m.Type
	}
	return types.SIGNED_MSG_TYPE_UNKNOWN
}

func (m *HasVote) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

// VoteSetMaj23Message is sent to indicate that a given BlockID has seen +2/3 votes.
type VoteSetMaj23 struct {
	Height               int64               `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Round                int32               `protobuf:"varint,2,opt,name=round,proto3" json:"round,omitempty"`
	Type                 types.SignedMsgType `protobuf:"varint,3,opt,name=type,proto3,enum=tendermint.proto.types.SignedMsgType" json:"type,omitempty"`
	BlockID              types.BlockID       `protobuf:"bytes,4,opt,name=block_id,json=blockId,proto3" json:"block_id"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *VoteSetMaj23) Reset()         { *m = VoteSetMaj23{} }
func (m *VoteSetMaj23) String() string { return proto.CompactTextString(m) }
func (*VoteSetMaj23) ProtoMessage()    {}
func (*VoteSetMaj23) Descriptor() ([]byte, []int) {
	return fileDescriptor_9de64017f8b3fc88, []int{7}
}
func (m *VoteSetMaj23) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VoteSetMaj23.Unmarshal(m, b)
}
func (m *VoteSetMaj23) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VoteSetMaj23.Marshal(b, m, deterministic)
}
func (m *VoteSetMaj23) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoteSetMaj23.Merge(m, src)
}
func (m *VoteSetMaj23) XXX_Size() int {
	return xxx_messageInfo_VoteSetMaj23.Size(m)
}
func (m *VoteSetMaj23) XXX_DiscardUnknown() {
	xxx_messageInfo_VoteSetMaj23.DiscardUnknown(m)
}

var xxx_messageInfo_VoteSetMaj23 proto.InternalMessageInfo

func (m *VoteSetMaj23) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *VoteSetMaj23) GetRound() int32 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *VoteSetMaj23) GetType() types.SignedMsgType {
	if m != nil {
		return m.Type
	}
	return types.SIGNED_MSG_TYPE_UNKNOWN
}

func (m *VoteSetMaj23) GetBlockID() types.BlockID {
	if m != nil {
		return m.BlockID
	}
	return types.BlockID{}
}

// VoteSetBitsMessage is sent to communicate the bit-array of votes seen for the BlockID.
type VoteSetBits struct {
	Height               int64               `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Round                int32               `protobuf:"varint,2,opt,name=round,proto3" json:"round,omitempty"`
	Type                 types.SignedMsgType `protobuf:"varint,3,opt,name=type,proto3,enum=tendermint.proto.types.SignedMsgType" json:"type,omitempty"`
	BlockID              types.BlockID       `protobuf:"bytes,4,opt,name=block_id,json=blockId,proto3" json:"block_id"`
	Votes                bits.BitArray       `protobuf:"bytes,5,opt,name=votes,proto3" json:"votes"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *VoteSetBits) Reset()         { *m = VoteSetBits{} }
func (m *VoteSetBits) String() string { return proto.CompactTextString(m) }
func (*VoteSetBits) ProtoMessage()    {}
func (*VoteSetBits) Descriptor() ([]byte, []int) {
	return fileDescriptor_9de64017f8b3fc88, []int{8}
}
func (m *VoteSetBits) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VoteSetBits.Unmarshal(m, b)
}
func (m *VoteSetBits) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VoteSetBits.Marshal(b, m, deterministic)
}
func (m *VoteSetBits) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoteSetBits.Merge(m, src)
}
func (m *VoteSetBits) XXX_Size() int {
	return xxx_messageInfo_VoteSetBits.Size(m)
}
func (m *VoteSetBits) XXX_DiscardUnknown() {
	xxx_messageInfo_VoteSetBits.DiscardUnknown(m)
}

var xxx_messageInfo_VoteSetBits proto.InternalMessageInfo

func (m *VoteSetBits) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *VoteSetBits) GetRound() int32 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *VoteSetBits) GetType() types.SignedMsgType {
	if m != nil {
		return m.Type
	}
	return types.SIGNED_MSG_TYPE_UNKNOWN
}

func (m *VoteSetBits) GetBlockID() types.BlockID {
	if m != nil {
		return m.BlockID
	}
	return types.BlockID{}
}

func (m *VoteSetBits) GetVotes() bits.BitArray {
	if m != nil {
		return m.Votes
	}
	return bits.BitArray{}
}

type Message struct {
	// Types that are valid to be assigned to Sum:
	//	*Message_NewRoundStep
	//	*Message_NewValidBlock
	//	*Message_Proposal
	//	*Message_ProposalPol
	//	*Message_BlockPart
	//	*Message_Vote
	//	*Message_HasVote
	//	*Message_VoteSetMaj23
	//	*Message_VoteSetBits
	Sum                  isMessage_Sum `protobuf_oneof:"sum"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_9de64017f8b3fc88, []int{9}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

type isMessage_Sum interface {
	isMessage_Sum()
}

type Message_NewRoundStep struct {
	NewRoundStep *NewRoundStep `protobuf:"bytes,1,opt,name=new_round_step,json=newRoundStep,proto3,oneof" json:"new_round_step,omitempty"`
}
type Message_NewValidBlock struct {
	NewValidBlock *NewValidBlock `protobuf:"bytes,2,opt,name=new_valid_block,json=newValidBlock,proto3,oneof" json:"new_valid_block,omitempty"`
}
type Message_Proposal struct {
	Proposal *Proposal `protobuf:"bytes,3,opt,name=proposal,proto3,oneof" json:"proposal,omitempty"`
}
type Message_ProposalPol struct {
	ProposalPol *ProposalPOL `protobuf:"bytes,4,opt,name=proposal_pol,json=proposalPol,proto3,oneof" json:"proposal_pol,omitempty"`
}
type Message_BlockPart struct {
	BlockPart *BlockPart `protobuf:"bytes,5,opt,name=block_part,json=blockPart,proto3,oneof" json:"block_part,omitempty"`
}
type Message_Vote struct {
	Vote *Vote `protobuf:"bytes,6,opt,name=vote,proto3,oneof" json:"vote,omitempty"`
}
type Message_HasVote struct {
	HasVote *HasVote `protobuf:"bytes,7,opt,name=has_vote,json=hasVote,proto3,oneof" json:"has_vote,omitempty"`
}
type Message_VoteSetMaj23 struct {
	VoteSetMaj23 *VoteSetMaj23 `protobuf:"bytes,8,opt,name=vote_set_maj23,json=voteSetMaj23,proto3,oneof" json:"vote_set_maj23,omitempty"`
}
type Message_VoteSetBits struct {
	VoteSetBits *VoteSetBits `protobuf:"bytes,9,opt,name=vote_set_bits,json=voteSetBits,proto3,oneof" json:"vote_set_bits,omitempty"`
}

func (*Message_NewRoundStep) isMessage_Sum()  {}
func (*Message_NewValidBlock) isMessage_Sum() {}
func (*Message_Proposal) isMessage_Sum()      {}
func (*Message_ProposalPol) isMessage_Sum()   {}
func (*Message_BlockPart) isMessage_Sum()     {}
func (*Message_Vote) isMessage_Sum()          {}
func (*Message_HasVote) isMessage_Sum()       {}
func (*Message_VoteSetMaj23) isMessage_Sum()  {}
func (*Message_VoteSetBits) isMessage_Sum()   {}

func (m *Message) GetSum() isMessage_Sum {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *Message) GetNewRoundStep() *NewRoundStep {
	if x, ok := m.GetSum().(*Message_NewRoundStep); ok {
		return x.NewRoundStep
	}
	return nil
}

func (m *Message) GetNewValidBlock() *NewValidBlock {
	if x, ok := m.GetSum().(*Message_NewValidBlock); ok {
		return x.NewValidBlock
	}
	return nil
}

func (m *Message) GetProposal() *Proposal {
	if x, ok := m.GetSum().(*Message_Proposal); ok {
		return x.Proposal
	}
	return nil
}

func (m *Message) GetProposalPol() *ProposalPOL {
	if x, ok := m.GetSum().(*Message_ProposalPol); ok {
		return x.ProposalPol
	}
	return nil
}

func (m *Message) GetBlockPart() *BlockPart {
	if x, ok := m.GetSum().(*Message_BlockPart); ok {
		return x.BlockPart
	}
	return nil
}

func (m *Message) GetVote() *Vote {
	if x, ok := m.GetSum().(*Message_Vote); ok {
		return x.Vote
	}
	return nil
}

func (m *Message) GetHasVote() *HasVote {
	if x, ok := m.GetSum().(*Message_HasVote); ok {
		return x.HasVote
	}
	return nil
}

func (m *Message) GetVoteSetMaj23() *VoteSetMaj23 {
	if x, ok := m.GetSum().(*Message_VoteSetMaj23); ok {
		return x.VoteSetMaj23
	}
	return nil
}

func (m *Message) GetVoteSetBits() *VoteSetBits {
	if x, ok := m.GetSum().(*Message_VoteSetBits); ok {
		return x.VoteSetBits
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Message) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Message_NewRoundStep)(nil),
		(*Message_NewValidBlock)(nil),
		(*Message_Proposal)(nil),
		(*Message_ProposalPol)(nil),
		(*Message_BlockPart)(nil),
		(*Message_Vote)(nil),
		(*Message_HasVote)(nil),
		(*Message_VoteSetMaj23)(nil),
		(*Message_VoteSetBits)(nil),
	}
}

func init() {
	proto.RegisterType((*NewRoundStep)(nil), "tendermint.proto.consensus.NewRoundStep")
	proto.RegisterType((*NewValidBlock)(nil), "tendermint.proto.consensus.NewValidBlock")
	proto.RegisterType((*Proposal)(nil), "tendermint.proto.consensus.Proposal")
	proto.RegisterType((*ProposalPOL)(nil), "tendermint.proto.consensus.ProposalPOL")
	proto.RegisterType((*BlockPart)(nil), "tendermint.proto.consensus.BlockPart")
	proto.RegisterType((*Vote)(nil), "tendermint.proto.consensus.Vote")
	proto.RegisterType((*HasVote)(nil), "tendermint.proto.consensus.HasVote")
	proto.RegisterType((*VoteSetMaj23)(nil), "tendermint.proto.consensus.VoteSetMaj23")
	proto.RegisterType((*VoteSetBits)(nil), "tendermint.proto.consensus.VoteSetBits")
	proto.RegisterType((*Message)(nil), "tendermint.proto.consensus.Message")
}

func init() { proto.RegisterFile("proto/consensus/msgs.proto", fileDescriptor_9de64017f8b3fc88) }

var fileDescriptor_9de64017f8b3fc88 = []byte{
	// 833 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x56, 0xcd, 0x6e, 0xeb, 0x44,
	0x14, 0x8e, 0x6f, 0x93, 0x26, 0x39, 0x4e, 0xda, 0xcb, 0x88, 0x9f, 0x28, 0x17, 0xa9, 0x91, 0xe1,
	0x42, 0x40, 0xc8, 0xb9, 0x4a, 0x25, 0x7e, 0x76, 0xc5, 0xfc, 0xc8, 0x15, 0x4d, 0x1b, 0x39, 0x55,
	0x25, 0xd8, 0x58, 0x4e, 0x3c, 0x72, 0x06, 0x62, 0x8f, 0xf1, 0x4c, 0x52, 0xf2, 0x00, 0x48, 0x3c,
	0x07, 0x6b, 0xb6, 0xbc, 0x01, 0x0b, 0x9e, 0xa2, 0x0b, 0x9e, 0x83, 0x05, 0x9a, 0x9f, 0xc4, 0x6e,
	0x2b, 0xb7, 0xcd, 0x06, 0xe9, 0x6e, 0xaa, 0x99, 0x39, 0xe7, 0xfb, 0x7c, 0xe6, 0x3b, 0x73, 0xbe,
	0x06, 0xba, 0x69, 0x46, 0x39, 0x1d, 0xcc, 0x68, 0xc2, 0x70, 0xc2, 0x96, 0x6c, 0x10, 0xb3, 0x88,
	0xd9, 0xf2, 0x10, 0x75, 0x39, 0x4e, 0x42, 0x9c, 0xc5, 0x24, 0xe1, 0xea, 0xc4, 0xde, 0xa6, 0x75,
	0x3f, 0xe0, 0x73, 0x92, 0x85, 0x7e, 0x1a, 0x64, 0x7c, 0x3d, 0x50, 0x1c, 0x11, 0x8d, 0x68, 0xbe,
	0x52, 0x88, 0xee, 0x3b, 0xea, 0x84, 0xaf, 0x53, 0xcc, 0xd4, 0x5f, 0x1d, 0x78, 0xa1, 0x02, 0x0b,
	0x32, 0x65, 0x83, 0x29, 0xe1, 0xb7, 0x82, 0xd6, 0x9f, 0x06, 0xb4, 0xce, 0xf1, 0xb5, 0x47, 0x97,
	0x49, 0x38, 0xe1, 0x38, 0x45, 0x6f, 0xc3, 0xfe, 0x1c, 0x93, 0x68, 0xce, 0x3b, 0x46, 0xcf, 0xe8,
	0xef, 0x79, 0x7a, 0x87, 0xde, 0x84, 0x5a, 0x26, 0x92, 0x3a, 0xcf, 0x7a, 0x46, 0xbf, 0xe6, 0xa9,
	0x0d, 0x42, 0x50, 0x65, 0x1c, 0xa7, 0x9d, 0xbd, 0x9e, 0xd1, 0x6f, 0x7b, 0x72, 0x8d, 0x3e, 0x83,
	0x0e, 0xc3, 0x33, 0x9a, 0x84, 0xcc, 0x67, 0x24, 0x99, 0x61, 0x9f, 0xf1, 0x20, 0xe3, 0x3e, 0x27,
	0x31, 0xee, 0x54, 0x25, 0xe7, 0x5b, 0x3a, 0x3e, 0x11, 0xe1, 0x89, 0x88, 0x5e, 0x92, 0x18, 0xa3,
	0x8f, 0xe1, 0x8d, 0x45, 0xc0, 0xb8, 0x3f, 0xa3, 0x71, 0x4c, 0xb8, 0xaf, 0x3e, 0x57, 0x93, 0x9f,
	0x3b, 0x14, 0x81, 0xaf, 0xe4, 0xb9, 0x2c, 0xd5, 0xfa, 0xd7, 0x80, 0xf6, 0x39, 0xbe, 0xbe, 0x0a,
	0x16, 0x24, 0x74, 0x16, 0x74, 0xf6, 0xd3, 0x8e, 0x85, 0x7f, 0x0f, 0x68, 0x2a, 0x60, 0x52, 0x57,
	0xe6, 0xcf, 0x71, 0x10, 0xe2, 0x4c, 0x5e, 0xc3, 0x1c, 0xbe, 0xb4, 0xef, 0xb5, 0x43, 0x49, 0x36,
	0x0e, 0x32, 0x3e, 0xc1, 0xdc, 0x95, 0xc9, 0x4e, 0xf5, 0xef, 0x9b, 0xa3, 0x8a, 0xf7, 0x5c, 0xd2,
	0x88, 0x08, 0x53, 0xe7, 0xe8, 0x1b, 0x30, 0x0b, 0xd4, 0xf2, 0xca, 0xe6, 0xf0, 0xfd, 0xfb, 0x9c,
	0xa2, 0x21, 0xb6, 0x68, 0x88, 0xed, 0x10, 0xfe, 0x65, 0x96, 0x05, 0x6b, 0x0f, 0x72, 0x32, 0xf4,
	0x02, 0x9a, 0x84, 0x69, 0x2d, 0xa4, 0x0a, 0x0d, 0xaf, 0x41, 0x98, 0xd2, 0xc0, 0x3a, 0x87, 0xc6,
	0x38, 0xa3, 0x29, 0x65, 0xc1, 0x02, 0x39, 0xd0, 0x48, 0xf5, 0x5a, 0x5e, 0xdd, 0x1c, 0xf6, 0x4a,
	0x2f, 0xa0, 0xf3, 0x74, 0xed, 0x5b, 0x9c, 0xf5, 0xbb, 0x01, 0xe6, 0x26, 0x38, 0xbe, 0x38, 0x2b,
	0x15, 0xf3, 0x13, 0x40, 0x1b, 0x8c, 0x9f, 0xd2, 0x85, 0x5f, 0x54, 0xf6, 0xf9, 0x26, 0x32, 0xa6,
	0x0b, 0xd9, 0x24, 0x34, 0x82, 0x56, 0x31, 0x5b, 0xcb, 0xfb, 0x24, 0x29, 0x74, 0x85, 0x66, 0x81,
	0xd3, 0xfa, 0x19, 0x9a, 0xce, 0x46, 0x9f, 0x1d, 0xdb, 0xfd, 0x29, 0x54, 0x45, 0x37, 0x74, 0x05,
	0xef, 0x3e, 0xd4, 0x60, 0xfd, 0x65, 0x99, 0x6f, 0x7d, 0x0e, 0xd5, 0x2b, 0xca, 0x31, 0x7a, 0x05,
	0xd5, 0x15, 0xe5, 0x58, 0xeb, 0x5b, 0x8a, 0x17, 0xb9, 0x9e, 0xcc, 0xb4, 0x7e, 0x33, 0xa0, 0xee,
	0x06, 0x4c, 0xa2, 0x77, 0xab, 0xf5, 0x0b, 0xa8, 0x0a, 0x36, 0x59, 0xeb, 0x41, 0xf9, 0x63, 0x9c,
	0x90, 0x28, 0xc1, 0xe1, 0x88, 0x45, 0x97, 0xeb, 0x14, 0x7b, 0x12, 0x22, 0x08, 0x49, 0x12, 0xe2,
	0x5f, 0xe4, 0xa3, 0x6b, 0x7b, 0x6a, 0x63, 0xfd, 0x65, 0x40, 0x4b, 0xd4, 0x31, 0xc1, 0x7c, 0x14,
	0xfc, 0x38, 0x3c, 0xfe, 0xff, 0xea, 0xf9, 0x0e, 0x1a, 0x6a, 0x14, 0x48, 0xa8, 0xe7, 0xe0, 0xa8,
	0x0c, 0x2e, 0x3b, 0x7b, 0xfa, 0xb5, 0x73, 0x28, 0xd4, 0xff, 0xe7, 0xe6, 0xa8, 0xae, 0x0f, 0xbc,
	0xba, 0x64, 0x38, 0x0d, 0xad, 0x5f, 0x9f, 0x81, 0xa9, 0xaf, 0xe1, 0x10, 0xce, 0x5e, 0xcf, 0x5b,
	0xa0, 0x13, 0xa8, 0x89, 0xf7, 0xc1, 0xe4, 0x48, 0xef, 0x36, 0x0c, 0x0a, 0x68, 0xfd, 0x51, 0x83,
	0xfa, 0x08, 0x33, 0x16, 0x44, 0x18, 0x8d, 0xe1, 0x20, 0xc1, 0xd7, 0x6a, 0x0c, 0x7d, 0xe9, 0xc4,
	0xea, 0x85, 0xf6, 0xed, 0xf2, 0xff, 0x28, 0x76, 0xd1, 0xef, 0xdd, 0x8a, 0xd7, 0x4a, 0x8a, 0xfe,
	0x3f, 0x81, 0x43, 0xc1, 0xb8, 0x12, 0xc6, 0xea, 0xcb, 0xa2, 0xa5, 0x8e, 0xe6, 0xf0, 0xa3, 0x47,
	0x28, 0x73, 0x2b, 0x76, 0x2b, 0x5e, 0x3b, 0xb9, 0xe5, 0xcd, 0x45, 0x8b, 0x2a, 0x35, 0x81, 0x9c,
	0x6d, 0xe3, 0x44, 0x6e, 0xc1, 0xa2, 0xd0, 0xd9, 0x1d, 0x33, 0x51, 0x9d, 0xf8, 0xf0, 0x29, 0x3c,
	0xe3, 0x8b, 0x33, 0xf7, 0xb6, 0x97, 0xa0, 0x6f, 0x01, 0x72, 0x93, 0xd6, 0xbd, 0x78, 0xf9, 0x10,
	0xd7, 0xd6, 0x79, 0xdc, 0x8a, 0xd7, 0xdc, 0xda, 0xb4, 0x30, 0x16, 0x69, 0x0c, 0xfb, 0x65, 0xc6,
	0x9b, 0x33, 0x88, 0xb7, 0xeb, 0x56, 0x94, 0x3d, 0xa0, 0x13, 0x68, 0xcc, 0x03, 0xe6, 0x4b, 0x6c,
	0x5d, 0x62, 0xdf, 0x7b, 0x08, 0xab, 0x9d, 0xc4, 0xad, 0x78, 0xf5, 0xb9, 0x36, 0x95, 0x31, 0x1c,
	0x08, 0xb4, 0xcf, 0x30, 0xf7, 0x63, 0x31, 0xd6, 0x9d, 0xc6, 0xe3, 0xad, 0x2f, 0xda, 0x80, 0x68,
	0xfd, 0xaa, 0x68, 0x0b, 0x23, 0x68, 0x6f, 0x19, 0xc5, 0xfb, 0xeb, 0x34, 0x1f, 0x97, 0xb8, 0x30,
	0x90, 0x42, 0xe2, 0x55, 0xbe, 0x75, 0x6a, 0xb0, 0xc7, 0x96, 0xb1, 0x33, 0xfc, 0xe1, 0x55, 0x44,
	0xf8, 0x7c, 0x39, 0xb5, 0x67, 0x34, 0x1e, 0xe4, 0x54, 0xc5, 0xe5, 0x9d, 0x9f, 0x46, 0xd3, 0x7d,
	0x79, 0x70, 0xfc, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xae, 0x4b, 0x4c, 0x2a, 0x34, 0x09, 0x00,
	0x00,
}
