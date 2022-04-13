// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/lightclients/beefy/v1/beefy.proto

package types

import (
	fmt "fmt"
	math "math"
	time "time"

	_ "github.com/confio/ics23/go"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// ClientState from Tendermint tracks the current validator set, latest height,
// and a possible frozen height.
type ClientState struct {
	// Latest mmr root hash
	MmrRootHash []byte `protobuf:"bytes,1,opt,name=mmr_root_hash,json=mmrRootHash,proto3" json:"mmr_root_hash,omitempty"`
	// block number for the latest mmr_root_hash
	LatestBeefyHeight uint32 `protobuf:"varint,2,opt,name=latest_beefy_height,json=latestBeefyHeight,proto3" json:"latest_beefy_height,omitempty"`
	// Block height when the client was frozen due to a misbehaviour
	FrozenHeight uint64 `protobuf:"varint,3,opt,name=frozen_height,json=frozenHeight,proto3" json:"frozen_height,omitempty"`
	// block number that the beefy protocol was activated on the relay chain.
	// This shoould be the first block in the merkle-mountain-range tree.
	BeefyActivationBlock uint32 `protobuf:"varint,4,opt,name=beefy_activation_block,json=beefyActivationBlock,proto3" json:"beefy_activation_block,omitempty"`
	// authorities for the current round
	Authority *BeefyAuthoritySet `protobuf:"bytes,5,opt,name=authority,proto3" json:"authority,omitempty"`
	// authorities for the next round
	NextAuthoritySet *BeefyAuthoritySet `protobuf:"bytes,6,opt,name=next_authority_set,json=nextAuthoritySet,proto3" json:"next_authority_set,omitempty"`
}

func (cs *ClientState) Reset()         { *cs = ClientState{} }
func (cs *ClientState) String() string { return proto.CompactTextString(cs) }
func (*ClientState) ProtoMessage()     {}
func (*ClientState) Descriptor() ([]byte, []int) {
	return fileDescriptor_43205c4bfbe9a422, []int{0}
}
func (cs *ClientState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientState.Unmarshal(cs, b)
}
func (cs *ClientState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientState.Marshal(b, cs, deterministic)
}
func (cs *ClientState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientState.Merge(cs, src)
}
func (cs *ClientState) XXX_Size() int {
	return xxx_messageInfo_ClientState.Size(cs)
}
func (cs *ClientState) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientState.DiscardUnknown(cs)
}

var xxx_messageInfo_ClientState proto.InternalMessageInfo

func (*ClientState) XXX_MessageName() string {
	return "ibc.lightclients.beefy.v1.ClientState"
}

// Actual payload items
type PayloadItem struct {
	// 2-byte payload id
	PayloadId *[2]byte `protobuf:"bytes,1,opt,name=payload_id,json=payloadId,proto3,customtype=[2]byte" json:"payload_id,omitempty"`
	// arbitrary length payload data., eg mmr_root_hash
	PayloadData []byte `protobuf:"bytes,2,opt,name=payload_data,json=payloadData,proto3" json:"payload_data,omitempty"`
}

func (m *PayloadItem) Reset()         { *m = PayloadItem{} }
func (m *PayloadItem) String() string { return proto.CompactTextString(m) }
func (*PayloadItem) ProtoMessage()    {}
func (*PayloadItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_43205c4bfbe9a422, []int{1}
}
func (m *PayloadItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PayloadItem.Unmarshal(m, b)
}
func (m *PayloadItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PayloadItem.Marshal(b, m, deterministic)
}
func (m *PayloadItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PayloadItem.Merge(m, src)
}
func (m *PayloadItem) XXX_Size() int {
	return xxx_messageInfo_PayloadItem.Size(m)
}
func (m *PayloadItem) XXX_DiscardUnknown() {
	xxx_messageInfo_PayloadItem.DiscardUnknown(m)
}

var xxx_messageInfo_PayloadItem proto.InternalMessageInfo

func (*PayloadItem) XXX_MessageName() string {
	return "ibc.lightclients.beefy.v1.PayloadItem"
}

// Commitment message signed by beefy validators
type Commitment struct {
	// array of payload items signed by Beefy validators
	Payload []*PayloadItem `protobuf:"bytes,1,rep,name=payload,proto3" json:"payload,omitempty"`
	// block number for this commitment
	BlockNumer uint32 `protobuf:"varint,2,opt,name=block_numer,json=blockNumer,proto3" json:"block_numer,omitempty"`
	// validator set that signed this commitment
	ValidatorSetId uint64 `protobuf:"varint,3,opt,name=validator_set_id,json=validatorSetId,proto3" json:"validator_set_id,omitempty"`
}

func (m *Commitment) Reset()         { *m = Commitment{} }
func (m *Commitment) String() string { return proto.CompactTextString(m) }
func (*Commitment) ProtoMessage()    {}
func (*Commitment) Descriptor() ([]byte, []int) {
	return fileDescriptor_43205c4bfbe9a422, []int{2}
}
func (m *Commitment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Commitment.Unmarshal(m, b)
}
func (m *Commitment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Commitment.Marshal(b, m, deterministic)
}
func (m *Commitment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Commitment.Merge(m, src)
}
func (m *Commitment) XXX_Size() int {
	return xxx_messageInfo_Commitment.Size(m)
}
func (m *Commitment) XXX_DiscardUnknown() {
	xxx_messageInfo_Commitment.DiscardUnknown(m)
}

var xxx_messageInfo_Commitment proto.InternalMessageInfo

func (*Commitment) XXX_MessageName() string {
	return "ibc.lightclients.beefy.v1.Commitment"
}

// Signature belonging to a single validator
type CommitmentSignature struct {
	// actual signature bytes
	Signature []byte `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	// authority leaf index in the merkle tree.
	AuthorityIndex uint32 `protobuf:"varint,2,opt,name=authority_index,json=authorityIndex,proto3" json:"authority_index,omitempty"`
}

func (m *CommitmentSignature) Reset()         { *m = CommitmentSignature{} }
func (m *CommitmentSignature) String() string { return proto.CompactTextString(m) }
func (*CommitmentSignature) ProtoMessage()    {}
func (*CommitmentSignature) Descriptor() ([]byte, []int) {
	return fileDescriptor_43205c4bfbe9a422, []int{3}
}
func (m *CommitmentSignature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommitmentSignature.Unmarshal(m, b)
}
func (m *CommitmentSignature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommitmentSignature.Marshal(b, m, deterministic)
}
func (m *CommitmentSignature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommitmentSignature.Merge(m, src)
}
func (m *CommitmentSignature) XXX_Size() int {
	return xxx_messageInfo_CommitmentSignature.Size(m)
}
func (m *CommitmentSignature) XXX_DiscardUnknown() {
	xxx_messageInfo_CommitmentSignature.DiscardUnknown(m)
}

var xxx_messageInfo_CommitmentSignature proto.InternalMessageInfo

func (*CommitmentSignature) XXX_MessageName() string {
	return "ibc.lightclients.beefy.v1.CommitmentSignature"
}

// signed commitment data
type SignedCommitment struct {
	// commitment data being signed
	Commitment *Commitment `protobuf:"bytes,1,opt,name=commitment,proto3" json:"commitment,omitempty"`
	// gotten from rpc subscription
	Signatures []*CommitmentSignature `protobuf:"bytes,2,rep,name=signatures,proto3" json:"signatures,omitempty"`
}

func (m *SignedCommitment) Reset()         { *m = SignedCommitment{} }
func (m *SignedCommitment) String() string { return proto.CompactTextString(m) }
func (*SignedCommitment) ProtoMessage()    {}
func (*SignedCommitment) Descriptor() ([]byte, []int) {
	return fileDescriptor_43205c4bfbe9a422, []int{4}
}
func (m *SignedCommitment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignedCommitment.Unmarshal(m, b)
}
func (m *SignedCommitment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignedCommitment.Marshal(b, m, deterministic)
}
func (m *SignedCommitment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignedCommitment.Merge(m, src)
}
func (m *SignedCommitment) XXX_Size() int {
	return xxx_messageInfo_SignedCommitment.Size(m)
}
func (m *SignedCommitment) XXX_DiscardUnknown() {
	xxx_messageInfo_SignedCommitment.DiscardUnknown(m)
}

var xxx_messageInfo_SignedCommitment proto.InternalMessageInfo

func (*SignedCommitment) XXX_MessageName() string {
	return "ibc.lightclients.beefy.v1.SignedCommitment"
}

// data needed to update the client
type MmrUpdateProof struct {
	// the new mmr leaf SCALE encoded.
	MmrLeaf *BeefyMmrLeaf `protobuf:"bytes,1,opt,name=mmr_leaf,json=mmrLeaf,proto3" json:"mmr_leaf,omitempty"`
	// leaf index for the mmr_leaf
	MmrLeafIndex uint64 `protobuf:"varint,2,opt,name=mmr_leaf_index,json=mmrLeafIndex,proto3" json:"mmr_leaf_index,omitempty"`
	// proof that this mmr_leaf index is valid.
	MmrProof [][]byte `protobuf:"bytes,3,rep,name=mmr_proof,json=mmrProof,proto3" json:"mmr_proof,omitempty"`
	// signed commitment data
	SignedCommitment *SignedCommitment `protobuf:"bytes,4,opt,name=signed_commitment,json=signedCommitment,proto3" json:"signed_commitment,omitempty"`
	// generated using full authority list from runtime
	AuthoritiesProof [][]byte `protobuf:"bytes,5,rep,name=authorities_proof,json=authoritiesProof,proto3" json:"authorities_proof,omitempty"`
}

func (m *MmrUpdateProof) Reset()         { *m = MmrUpdateProof{} }
func (m *MmrUpdateProof) String() string { return proto.CompactTextString(m) }
func (*MmrUpdateProof) ProtoMessage()    {}
func (*MmrUpdateProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_43205c4bfbe9a422, []int{5}
}
func (m *MmrUpdateProof) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MmrUpdateProof.Unmarshal(m, b)
}
func (m *MmrUpdateProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MmrUpdateProof.Marshal(b, m, deterministic)
}
func (m *MmrUpdateProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MmrUpdateProof.Merge(m, src)
}
func (m *MmrUpdateProof) XXX_Size() int {
	return xxx_messageInfo_MmrUpdateProof.Size(m)
}
func (m *MmrUpdateProof) XXX_DiscardUnknown() {
	xxx_messageInfo_MmrUpdateProof.DiscardUnknown(m)
}

var xxx_messageInfo_MmrUpdateProof proto.InternalMessageInfo

func (*MmrUpdateProof) XXX_MessageName() string {
	return "ibc.lightclients.beefy.v1.MmrUpdateProof"
}

// ConsensusState defines the consensus state from Tendermint.
type ConsensusState struct {
	// timestamp that corresponds to the block height in which the ConsensusState
	// was stored.
	Timestamp time.Time `protobuf:"bytes,1,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
	// packet commitment root
	Root []byte `protobuf:"bytes,2,opt,name=root,proto3" json:"root,omitempty"`
	// proof of inclusion for this parachain header in the Mmr.
	ParachainHeader ParachainHeader `protobuf:"bytes,4,opt,name=parachain_header,json=parachainHeader,proto3" json:"parachain_header"`
}

func (cs *ConsensusState) Reset()         { *cs = ConsensusState{} }
func (cs *ConsensusState) String() string { return proto.CompactTextString(cs) }
func (*ConsensusState) ProtoMessage()     {}
func (*ConsensusState) Descriptor() ([]byte, []int) {
	return fileDescriptor_43205c4bfbe9a422, []int{6}
}
func (cs *ConsensusState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsensusState.Unmarshal(cs, b)
}
func (cs *ConsensusState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsensusState.Marshal(b, cs, deterministic)
}
func (cs *ConsensusState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsensusState.Merge(cs, src)
}
func (cs *ConsensusState) XXX_Size() int {
	return xxx_messageInfo_ConsensusState.Size(cs)
}
func (cs *ConsensusState) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsensusState.DiscardUnknown(cs)
}

var xxx_messageInfo_ConsensusState proto.InternalMessageInfo

func (*ConsensusState) XXX_MessageName() string {
	return "ibc.lightclients.beefy.v1.ConsensusState"
}

// Misbehaviour is a wrapper over two conflicting Headers
// that implements Misbehaviour interface expected by ICS-02
type Misbehaviour struct {
	Header1 *Header `protobuf:"bytes,2,opt,name=header_1,json=header1,proto3" json:"header_1,omitempty" yaml:"header_1"`
	Header2 *Header `protobuf:"bytes,3,opt,name=header_2,json=header2,proto3" json:"header_2,omitempty" yaml:"header_2"`
}

func (m *Misbehaviour) Reset()         { *m = Misbehaviour{} }
func (m *Misbehaviour) String() string { return proto.CompactTextString(m) }
func (*Misbehaviour) ProtoMessage()    {}
func (*Misbehaviour) Descriptor() ([]byte, []int) {
	return fileDescriptor_43205c4bfbe9a422, []int{7}
}
func (m *Misbehaviour) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Misbehaviour.Unmarshal(m, b)
}
func (m *Misbehaviour) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Misbehaviour.Marshal(b, m, deterministic)
}
func (m *Misbehaviour) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Misbehaviour.Merge(m, src)
}
func (m *Misbehaviour) XXX_Size() int {
	return xxx_messageInfo_Misbehaviour.Size(m)
}
func (m *Misbehaviour) XXX_DiscardUnknown() {
	xxx_messageInfo_Misbehaviour.DiscardUnknown(m)
}

var xxx_messageInfo_Misbehaviour proto.InternalMessageInfo

func (*Misbehaviour) XXX_MessageName() string {
	return "ibc.lightclients.beefy.v1.Misbehaviour"
}

// Header contains the neccessary data to proove finality about IBC commitments
type Header struct {
	// parachain headers needed for proofs and ConsensusState
	ParachainHeaders []*ParachainHeader `protobuf:"bytes,1,rep,name=parachain_headers,json=parachainHeaders,proto3" json:"parachain_headers,omitempty"`
	// mmr proofs for the headers gotten from rpc "mmr_generateProofs"
	MmrProofs [][]byte `protobuf:"bytes,2,rep,name=mmr_proofs,json=mmrProofs,proto3" json:"mmr_proofs,omitempty"`
	// size of the mmr for the given proof
	MmrSize uint64 `protobuf:"varint,3,opt,name=mmr_size,json=mmrSize,proto3" json:"mmr_size,omitempty"`
	// optional payload to update the mmr root hash.
	MmrUpdateProof *MmrUpdateProof `protobuf:"bytes,4,opt,name=mmr_update_proof,json=mmrUpdateProof,proto3" json:"mmr_update_proof,omitempty"`
}

func (h *Header) Reset()         { *h = Header{} }
func (h *Header) String() string { return proto.CompactTextString(h) }
func (*Header) ProtoMessage()    {}
func (*Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_43205c4bfbe9a422, []int{8}
}
func (h *Header) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Header.Unmarshal(h, b)
}
func (h *Header) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Header.Marshal(b, h, deterministic)
}
func (h *Header) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Header.Merge(h, src)
}
func (h *Header) XXX_Size() int {
	return xxx_messageInfo_Header.Size(h)
}
func (h *Header) XXX_DiscardUnknown() {
	xxx_messageInfo_Header.DiscardUnknown(h)
}

var xxx_messageInfo_Header proto.InternalMessageInfo

func (*Header) XXX_MessageName() string {
	return "ibc.lightclients.beefy.v1.Header"
}

// data needed to prove parachain header inclusion in mmr.
type ParachainHeader struct {
	// scale-encoded parachain header bytes
	ParachainHeader []byte `protobuf:"bytes,1,opt,name=parachain_header,json=parachainHeader,proto3" json:"parachain_header,omitempty"`
	// reconstructed MmrLeaf, see beefy-go spec
	MmrLeafPartial *BeefyMmrLeafPartial `protobuf:"bytes,2,opt,name=mmr_leaf_partial,json=mmrLeafPartial,proto3" json:"mmr_leaf_partial,omitempty"`
	// para_id of the header.
	ParaId uint32 `protobuf:"varint,3,opt,name=para_id,json=paraId,proto3" json:"para_id,omitempty"`
	// proofs for our header in the parachain heads root
	ParachainHeadsProof [][]byte `protobuf:"bytes,4,rep,name=parachain_heads_proof,json=parachainHeadsProof,proto3" json:"parachain_heads_proof,omitempty"`
	// leaf index for parachain heads proof
	HeadsLeafIndex uint32 `protobuf:"varint,5,opt,name=heads_leaf_index,json=headsLeafIndex,proto3" json:"heads_leaf_index,omitempty"`
	// total number of para heads in parachain_heads_root
	HeadsTotalCount uint32 `protobuf:"varint,6,opt,name=heads_total_count,json=headsTotalCount,proto3" json:"heads_total_count,omitempty"`
	// trie merkle proof of inclusion in header.extrinsic_root
	// this already encodes the actual extrinsic
	ExtrinsicProof [][]byte `protobuf:"bytes,7,rep,name=extrinsic_proof,json=extrinsicProof,proto3" json:"extrinsic_proof,omitempty"`
}

func (m *ParachainHeader) Reset()         { *m = ParachainHeader{} }
func (m *ParachainHeader) String() string { return proto.CompactTextString(m) }
func (*ParachainHeader) ProtoMessage()    {}
func (*ParachainHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_43205c4bfbe9a422, []int{9}
}
func (m *ParachainHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParachainHeader.Unmarshal(m, b)
}
func (m *ParachainHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParachainHeader.Marshal(b, m, deterministic)
}
func (m *ParachainHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParachainHeader.Merge(m, src)
}
func (m *ParachainHeader) XXX_Size() int {
	return xxx_messageInfo_ParachainHeader.Size(m)
}
func (m *ParachainHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_ParachainHeader.DiscardUnknown(m)
}

var xxx_messageInfo_ParachainHeader proto.InternalMessageInfo

func (*ParachainHeader) XXX_MessageName() string {
	return "ibc.lightclients.beefy.v1.ParachainHeader"
}

// Partial data for MmrLeaf
type BeefyMmrLeafPartial struct {
	// leaf version
	Version uint8 `protobuf:"varint,1,opt,name=version,proto3,customtype=uint8" json:"version"`
	// parent block for this leaf
	ParentNumber uint32 `protobuf:"varint,2,opt,name=parent_number,json=parentNumber,proto3" json:"parent_number,omitempty"`
	// parent hash for this leaf
	ParentHash *[32]byte `protobuf:"bytes,3,opt,name=parent_hash,json=parentHash,proto3,customtype=[32]byte" json:"parent_hash,omitempty"`
	// next authority set.
	BeefyNextAuthoritySet BeefyAuthoritySet `protobuf:"bytes,4,opt,name=beefy_next_authority_set,json=beefyNextAuthoritySet,proto3" json:"beefy_next_authority_set"`
}

func (m *BeefyMmrLeafPartial) Reset()         { *m = BeefyMmrLeafPartial{} }
func (m *BeefyMmrLeafPartial) String() string { return proto.CompactTextString(m) }
func (*BeefyMmrLeafPartial) ProtoMessage()    {}
func (*BeefyMmrLeafPartial) Descriptor() ([]byte, []int) {
	return fileDescriptor_43205c4bfbe9a422, []int{10}
}
func (m *BeefyMmrLeafPartial) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BeefyMmrLeafPartial.Unmarshal(m, b)
}
func (m *BeefyMmrLeafPartial) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BeefyMmrLeafPartial.Marshal(b, m, deterministic)
}
func (m *BeefyMmrLeafPartial) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BeefyMmrLeafPartial.Merge(m, src)
}
func (m *BeefyMmrLeafPartial) XXX_Size() int {
	return xxx_messageInfo_BeefyMmrLeafPartial.Size(m)
}
func (m *BeefyMmrLeafPartial) XXX_DiscardUnknown() {
	xxx_messageInfo_BeefyMmrLeafPartial.DiscardUnknown(m)
}

var xxx_messageInfo_BeefyMmrLeafPartial proto.InternalMessageInfo

func (*BeefyMmrLeafPartial) XXX_MessageName() string {
	return "ibc.lightclients.beefy.v1.BeefyMmrLeafPartial"
}

// Beefy Authority Info
type BeefyAuthoritySet struct {
	// Id of the authority set, it should be strictly increasing
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// size of the authority set
	Len uint32 `protobuf:"varint,2,opt,name=len,proto3" json:"len,omitempty"`
	// merkle root of the sorted authority public keys.
	AuthorityRoot *[32]byte `protobuf:"bytes,3,opt,name=authority_root,json=authorityRoot,proto3,customtype=[32]byte" json:"authority_root,omitempty"`
}

func (m *BeefyAuthoritySet) Reset()         { *m = BeefyAuthoritySet{} }
func (m *BeefyAuthoritySet) String() string { return proto.CompactTextString(m) }
func (*BeefyAuthoritySet) ProtoMessage()    {}
func (*BeefyAuthoritySet) Descriptor() ([]byte, []int) {
	return fileDescriptor_43205c4bfbe9a422, []int{11}
}
func (m *BeefyAuthoritySet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BeefyAuthoritySet.Unmarshal(m, b)
}
func (m *BeefyAuthoritySet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BeefyAuthoritySet.Marshal(b, m, deterministic)
}
func (m *BeefyAuthoritySet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BeefyAuthoritySet.Merge(m, src)
}
func (m *BeefyAuthoritySet) XXX_Size() int {
	return xxx_messageInfo_BeefyAuthoritySet.Size(m)
}
func (m *BeefyAuthoritySet) XXX_DiscardUnknown() {
	xxx_messageInfo_BeefyAuthoritySet.DiscardUnknown(m)
}

var xxx_messageInfo_BeefyAuthoritySet proto.InternalMessageInfo

func (*BeefyAuthoritySet) XXX_MessageName() string {
	return "ibc.lightclients.beefy.v1.BeefyAuthoritySet"
}

type BeefyMmrLeaf struct {
	// leaf version
	Version uint8 `protobuf:"varint,1,opt,name=version,proto3,customtype=uint8" json:"version"`
	// parent block for this leaf
	ParentNumber uint32 `protobuf:"varint,2,opt,name=parent_number,json=parentNumber,proto3" json:"parent_number,omitempty"`
	// parent hash for this leaf
	ParentHash *[32]byte `protobuf:"bytes,3,opt,name=parent_hash,json=parentHash,proto3,customtype=[32]byte" json:"parent_hash,omitempty"`
	// beefy next authority set.
	BeefyNextAuthoritySet BeefyAuthoritySet `protobuf:"bytes,4,opt,name=beefy_next_authority_set,json=beefyNextAuthoritySet,proto3" json:"beefy_next_authority_set"`
	// merkle root hash of parachain heads included in the leaf.
	ParachainHeads *[32]byte `protobuf:"bytes,5,opt,name=parachain_heads,json=parachainHeads,proto3,customtype=[32]byte" json:"parachain_heads,omitempty"`
}

func (m *BeefyMmrLeaf) Reset()         { *m = BeefyMmrLeaf{} }
func (m *BeefyMmrLeaf) String() string { return proto.CompactTextString(m) }
func (*BeefyMmrLeaf) ProtoMessage()    {}
func (*BeefyMmrLeaf) Descriptor() ([]byte, []int) {
	return fileDescriptor_43205c4bfbe9a422, []int{12}
}
func (m *BeefyMmrLeaf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BeefyMmrLeaf.Unmarshal(m, b)
}
func (m *BeefyMmrLeaf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BeefyMmrLeaf.Marshal(b, m, deterministic)
}
func (m *BeefyMmrLeaf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BeefyMmrLeaf.Merge(m, src)
}
func (m *BeefyMmrLeaf) XXX_Size() int {
	return xxx_messageInfo_BeefyMmrLeaf.Size(m)
}
func (m *BeefyMmrLeaf) XXX_DiscardUnknown() {
	xxx_messageInfo_BeefyMmrLeaf.DiscardUnknown(m)
}

var xxx_messageInfo_BeefyMmrLeaf proto.InternalMessageInfo

func (*BeefyMmrLeaf) XXX_MessageName() string {
	return "ibc.lightclients.beefy.v1.BeefyMmrLeaf"
}
func init() {
	proto.RegisterType((*ClientState)(nil), "ibc.lightclients.beefy.v1.ClientState")
	golang_proto.RegisterType((*ClientState)(nil), "ibc.lightclients.beefy.v1.ClientState")
	proto.RegisterType((*PayloadItem)(nil), "ibc.lightclients.beefy.v1.PayloadItem")
	golang_proto.RegisterType((*PayloadItem)(nil), "ibc.lightclients.beefy.v1.PayloadItem")
	proto.RegisterType((*Commitment)(nil), "ibc.lightclients.beefy.v1.Commitment")
	golang_proto.RegisterType((*Commitment)(nil), "ibc.lightclients.beefy.v1.Commitment")
	proto.RegisterType((*CommitmentSignature)(nil), "ibc.lightclients.beefy.v1.CommitmentSignature")
	golang_proto.RegisterType((*CommitmentSignature)(nil), "ibc.lightclients.beefy.v1.CommitmentSignature")
	proto.RegisterType((*SignedCommitment)(nil), "ibc.lightclients.beefy.v1.SignedCommitment")
	golang_proto.RegisterType((*SignedCommitment)(nil), "ibc.lightclients.beefy.v1.SignedCommitment")
	proto.RegisterType((*MmrUpdateProof)(nil), "ibc.lightclients.beefy.v1.MmrUpdateProof")
	golang_proto.RegisterType((*MmrUpdateProof)(nil), "ibc.lightclients.beefy.v1.MmrUpdateProof")
	proto.RegisterType((*ConsensusState)(nil), "ibc.lightclients.beefy.v1.ConsensusState")
	golang_proto.RegisterType((*ConsensusState)(nil), "ibc.lightclients.beefy.v1.ConsensusState")
	proto.RegisterType((*Misbehaviour)(nil), "ibc.lightclients.beefy.v1.Misbehaviour")
	golang_proto.RegisterType((*Misbehaviour)(nil), "ibc.lightclients.beefy.v1.Misbehaviour")
	proto.RegisterType((*Header)(nil), "ibc.lightclients.beefy.v1.Header")
	golang_proto.RegisterType((*Header)(nil), "ibc.lightclients.beefy.v1.Header")
	proto.RegisterType((*ParachainHeader)(nil), "ibc.lightclients.beefy.v1.ParachainHeader")
	golang_proto.RegisterType((*ParachainHeader)(nil), "ibc.lightclients.beefy.v1.ParachainHeader")
	proto.RegisterType((*BeefyMmrLeafPartial)(nil), "ibc.lightclients.beefy.v1.BeefyMmrLeafPartial")
	golang_proto.RegisterType((*BeefyMmrLeafPartial)(nil), "ibc.lightclients.beefy.v1.BeefyMmrLeafPartial")
	proto.RegisterType((*BeefyAuthoritySet)(nil), "ibc.lightclients.beefy.v1.BeefyAuthoritySet")
	golang_proto.RegisterType((*BeefyAuthoritySet)(nil), "ibc.lightclients.beefy.v1.BeefyAuthoritySet")
	proto.RegisterType((*BeefyMmrLeaf)(nil), "ibc.lightclients.beefy.v1.BeefyMmrLeaf")
	golang_proto.RegisterType((*BeefyMmrLeaf)(nil), "ibc.lightclients.beefy.v1.BeefyMmrLeaf")
}

func init() {
	proto.RegisterFile("ibc/lightclients/beefy/v1/beefy.proto", fileDescriptor_43205c4bfbe9a422)
}
func init() {
	golang_proto.RegisterFile("ibc/lightclients/beefy/v1/beefy.proto", fileDescriptor_43205c4bfbe9a422)
}

var fileDescriptor_43205c4bfbe9a422 = []byte{
	// 1279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe4, 0x57, 0xcf, 0x6f, 0x1b, 0xc5,
	0x17, 0xf7, 0x3a, 0x4e, 0x9c, 0x3c, 0xff, 0x88, 0x33, 0x69, 0xbf, 0x5f, 0xb7, 0x80, 0x9d, 0x06,
	0x4a, 0xdd, 0x96, 0xd8, 0xb2, 0x03, 0x12, 0xea, 0xa9, 0x75, 0x40, 0x6a, 0x10, 0x8d, 0xaa, 0x49,
	0x11, 0xa5, 0x05, 0x59, 0x63, 0xef, 0xd8, 0x1e, 0x75, 0x77, 0xc7, 0xda, 0x19, 0x5b, 0x4d, 0xef,
	0x48, 0x1c, 0xfb, 0x1f, 0xd0, 0x03, 0x77, 0x4e, 0xdc, 0x39, 0x56, 0xe2, 0xd2, 0x23, 0xea, 0x21,
	0xa0, 0xe4, 0x82, 0x38, 0x72, 0xe6, 0x80, 0xe6, 0xc7, 0xae, 0x37, 0x6e, 0x30, 0x0d, 0x57, 0x6e,
	0xb3, 0xef, 0xe7, 0xe7, 0x7d, 0x66, 0xde, 0x9b, 0x59, 0xb8, 0xcc, 0xba, 0xbd, 0x86, 0xc7, 0x06,
	0x43, 0xd9, 0xf3, 0x18, 0x0d, 0xa4, 0x68, 0x74, 0x29, 0xed, 0x1f, 0x34, 0x26, 0x4d, 0xb3, 0xa8,
	0x8f, 0x42, 0x2e, 0x39, 0xba, 0xc0, 0xba, 0xbd, 0x7a, 0xd2, 0xac, 0x6e, 0xb4, 0x93, 0xe6, 0xc5,
	0xca, 0x80, 0xf3, 0x81, 0x47, 0x1b, 0xda, 0xb0, 0x3b, 0xee, 0x37, 0xdc, 0x71, 0x48, 0x24, 0xe3,
	0x81, 0x71, 0xbd, 0x58, 0x9d, 0xd5, 0x4b, 0xe6, 0x53, 0x21, 0x89, 0x3f, 0xb2, 0x06, 0xe7, 0x06,
	0x7c, 0xc0, 0xf5, 0xb2, 0xa1, 0x56, 0x56, 0x9a, 0x1f, 0x85, 0x9c, 0xf7, 0x85, 0xf9, 0xda, 0xfc,
	0x3d, 0x0d, 0xb9, 0x1d, 0x9d, 0x79, 0x5f, 0x12, 0x49, 0xd1, 0x26, 0x14, 0x7c, 0x3f, 0xec, 0x84,
	0x9c, 0xcb, 0xce, 0x90, 0x88, 0x61, 0xd9, 0xd9, 0x70, 0x6a, 0x79, 0x9c, 0xf3, 0xfd, 0x10, 0x73,
	0x2e, 0x6f, 0x13, 0x31, 0x44, 0x75, 0x58, 0xf7, 0x88, 0xa4, 0x42, 0x76, 0x34, 0xd6, 0xce, 0x90,
	0x2a, 0xfc, 0xe5, 0xf4, 0x86, 0x53, 0x2b, 0xe0, 0x35, 0xa3, 0x6a, 0x2b, 0xcd, 0x6d, 0xad, 0x40,
	0x6f, 0x43, 0xa1, 0x1f, 0xf2, 0x27, 0x34, 0x88, 0x2c, 0x17, 0x36, 0x9c, 0x5a, 0x06, 0xe7, 0x8d,
	0xd0, 0x1a, 0xbd, 0x0f, 0xff, 0x33, 0xd1, 0x48, 0x4f, 0xb2, 0x89, 0xae, 0xb3, 0xd3, 0xf5, 0x78,
	0xef, 0x51, 0x39, 0xa3, 0xe3, 0x9e, 0xd3, 0xda, 0x5b, 0xb1, 0xb2, 0xad, 0x74, 0xe8, 0x13, 0x58,
	0x21, 0x63, 0x39, 0xe4, 0x21, 0x93, 0x07, 0xe5, 0xc5, 0x0d, 0xa7, 0x96, 0x6b, 0xbd, 0x57, 0xff,
	0x5b, 0x4a, 0xeb, 0x1a, 0xd5, 0xad, 0xc8, 0x61, 0x9f, 0x4a, 0x3c, 0x75, 0x47, 0x0f, 0x00, 0x05,
	0xf4, 0xb1, 0xec, 0xc4, 0x92, 0x8e, 0xa0, 0xb2, 0xbc, 0xf4, 0x2f, 0x82, 0x96, 0x54, 0x9c, 0xa4,
	0xe4, 0x46, 0xe6, 0x9b, 0x67, 0xd5, 0xd4, 0x66, 0x17, 0x72, 0x77, 0xc9, 0x81, 0xc7, 0x89, 0xbb,
	0x2b, 0xa9, 0x8f, 0xae, 0x01, 0x8c, 0xcc, 0x67, 0x87, 0xb9, 0x86, 0xe8, 0x76, 0xee, 0xe5, 0x61,
	0x35, 0xfb, 0xb0, 0xf5, 0x55, 0xf7, 0x40, 0x52, 0xbc, 0x62, 0xd5, 0xbb, 0x2e, 0xba, 0x04, 0xf9,
	0xc8, 0xd6, 0x25, 0x92, 0x68, 0xb2, 0xf3, 0x38, 0x67, 0x65, 0x1f, 0x11, 0x49, 0x6c, 0x8e, 0x6f,
	0x1d, 0x80, 0x1d, 0xee, 0xfb, 0x4c, 0xfa, 0x34, 0x90, 0xe8, 0x26, 0x64, 0xad, 0x4d, 0xd9, 0xd9,
	0x58, 0xa8, 0xe5, 0x5a, 0xef, 0xce, 0xa9, 0x24, 0x01, 0x0e, 0x47, 0x6e, 0xa8, 0x0a, 0x39, 0xbd,
	0x0f, 0x9d, 0x60, 0xec, 0xd3, 0xd0, 0xee, 0x32, 0x68, 0xd1, 0x9e, 0x92, 0xa0, 0x1a, 0x94, 0x26,
	0xc4, 0x63, 0x2e, 0x91, 0x3c, 0x54, 0x94, 0xa9, 0x62, 0xcc, 0x0e, 0x17, 0x63, 0xf9, 0x3e, 0x95,
	0xbb, 0x6e, 0xcc, 0xc2, 0xfa, 0x14, 0xe0, 0x3e, 0x1b, 0x04, 0x44, 0x8e, 0x43, 0x8a, 0xde, 0x84,
	0x15, 0x11, 0x7d, 0xd8, 0x53, 0x37, 0x15, 0xa0, 0x2b, 0xb0, 0x3a, 0xdd, 0x17, 0x16, 0xb8, 0xf4,
	0xb1, 0x45, 0x52, 0x8c, 0xc5, 0xbb, 0x4a, 0x6a, 0x73, 0x7c, 0xef, 0x40, 0x49, 0x85, 0xa6, 0x6e,
	0x82, 0x8b, 0x8f, 0x01, 0x7a, 0xf1, 0x97, 0x4e, 0x91, 0x6b, 0x5d, 0x9e, 0x43, 0xc7, 0xd4, 0x15,
	0x27, 0x1c, 0xd1, 0x1e, 0x40, 0x8c, 0x4b, 0x94, 0xd3, 0x9a, 0xd5, 0xfa, 0x6b, 0x85, 0x89, 0x8b,
	0xc5, 0x89, 0x08, 0x16, 0xf1, 0x77, 0x69, 0x28, 0xde, 0xf1, 0xc3, 0xcf, 0x46, 0x2e, 0x91, 0xf4,
	0xae, 0x6a, 0x51, 0xd4, 0x86, 0x65, 0xd5, 0x8b, 0x1e, 0x25, 0x7d, 0x8b, 0xf6, 0xca, 0x3f, 0x1d,
	0xc3, 0x3b, 0x7e, 0xf8, 0x29, 0x25, 0x7d, 0x9c, 0xf5, 0xcd, 0x02, 0xbd, 0x03, 0xc5, 0x28, 0x46,
	0x82, 0xb6, 0x0c, 0xce, 0x5b, 0x03, 0x4d, 0x1a, 0x7a, 0x03, 0x56, 0x94, 0x95, 0x9e, 0x0c, 0xe5,
	0x85, 0x8d, 0x85, 0x5a, 0x1e, 0xab, 0xd4, 0x06, 0xc6, 0x7d, 0x58, 0x13, 0x9a, 0xca, 0x4e, 0x82,
	0xbd, 0x8c, 0xc6, 0x73, 0x7d, 0x0e, 0x9e, 0x59, 0xfa, 0x71, 0x49, 0xcc, 0x6e, 0xc8, 0x75, 0x58,
	0x8b, 0x76, 0x8f, 0x51, 0x61, 0xd3, 0x2f, 0xea, 0xf4, 0xa5, 0x84, 0x42, 0xc3, 0xb0, 0x34, 0xfd,
	0xe4, 0x40, 0x71, 0x87, 0x07, 0x82, 0x06, 0x62, 0x2c, 0xcc, 0xc8, 0x6a, 0xc3, 0x4a, 0x3c, 0xf9,
	0x2c, 0x4f, 0x17, 0xeb, 0x66, 0x36, 0xd6, 0xa3, 0xd9, 0x58, 0xbf, 0x17, 0x59, 0xb4, 0x97, 0x9f,
	0x1f, 0x56, 0x53, 0x4f, 0x7f, 0xa9, 0x3a, 0x78, 0xea, 0x86, 0x10, 0x64, 0xd4, 0xc8, 0xb3, 0x6d,
	0xa5, 0xd7, 0xe8, 0x21, 0x94, 0x46, 0x24, 0x24, 0xbd, 0x21, 0x61, 0x6a, 0x72, 0x11, 0x97, 0x86,
	0xb6, 0xec, 0x6b, 0x73, 0x7b, 0xc8, 0xba, 0xdc, 0xd6, 0x1e, 0xed, 0x8c, 0x4a, 0x87, 0x57, 0x47,
	0x27, 0xc5, 0xb6, 0x9a, 0x17, 0x0e, 0xe4, 0xef, 0x30, 0xd1, 0xa5, 0x43, 0x32, 0x61, 0x7c, 0x1c,
	0xa2, 0x2f, 0x61, 0xd9, 0x64, 0xea, 0x34, 0x35, 0x96, 0x5c, 0xeb, 0xd2, 0x9c, 0x5c, 0x36, 0x45,
	0xe5, 0xe8, 0xb0, 0x9a, 0x35, 0xeb, 0xe6, 0x1f, 0x87, 0xd5, 0xd5, 0x03, 0xe2, 0x7b, 0x37, 0x36,
	0xa3, 0x38, 0x9b, 0x38, 0x6b, 0x96, 0xcd, 0x44, 0xf4, 0x96, 0xee, 0xd0, 0xb3, 0x46, 0x6f, 0xbd,
	0x12, 0xbd, 0x15, 0x47, 0x6f, 0xd9, 0x92, 0xfe, 0x74, 0x60, 0xc9, 0x58, 0xa3, 0xcf, 0x61, 0x6d,
	0x96, 0x40, 0x61, 0xa7, 0xd0, 0x19, 0x18, 0xc4, 0xa5, 0x19, 0xee, 0x04, 0x7a, 0x0b, 0x20, 0x3e,
	0xae, 0xa6, 0x03, 0xf3, 0x78, 0x25, 0x3a, 0xaf, 0x02, 0x5d, 0x30, 0x7d, 0x23, 0xd8, 0x13, 0x6a,
	0x07, 0x91, 0x6a, 0x87, 0x7d, 0xf6, 0x84, 0xa2, 0x2f, 0xa0, 0xa4, 0x54, 0x63, 0xdd, 0x65, 0xf6,
	0xc0, 0x99, 0x3d, 0xbd, 0x3a, 0x07, 0xd1, 0xc9, 0xbe, 0xd4, 0x5b, 0xea, 0x60, 0xd5, 0x57, 0x09,
	0xa9, 0x2d, 0xff, 0xb7, 0x34, 0xac, 0xce, 0x14, 0x80, 0xae, 0x9e, 0x72, 0x90, 0xcc, 0x80, 0x9b,
	0x3d, 0x16, 0xe8, 0xbe, 0xc1, 0xa7, 0xdb, 0x75, 0x44, 0x42, 0xc9, 0x88, 0x67, 0xcf, 0x41, 0xfd,
	0x35, 0x5b, 0xff, 0xae, 0xf1, 0xd2, 0xf0, 0x12, 0xdf, 0xe8, 0xff, 0xea, 0x22, 0x08, 0x49, 0x34,
	0x9c, 0x0b, 0x78, 0x49, 0x7d, 0xee, 0xba, 0xa8, 0x05, 0xe7, 0x4f, 0xa2, 0x13, 0x31, 0x2f, 0x8a,
	0xd7, 0xf5, 0x13, 0x10, 0x4d, 0x2f, 0xaa, 0x91, 0x6f, 0x2c, 0x13, 0x73, 0x65, 0xd1, 0x8c, 0x63,
	0x2d, 0x9f, 0x4e, 0x96, 0x6b, 0xb0, 0x66, 0x2c, 0x25, 0x97, 0xc4, 0xeb, 0xf4, 0xf8, 0x38, 0x30,
	0x77, 0x6a, 0x01, 0xaf, 0x6a, 0xc5, 0x3d, 0x25, 0xdf, 0x51, 0x62, 0x35, 0xe3, 0xe9, 0x63, 0x19,
	0xb2, 0x40, 0xb0, 0x9e, 0xc5, 0x90, 0xd5, 0x18, 0x8a, 0xb1, 0x38, 0x49, 0xf5, 0xd7, 0x69, 0x58,
	0x3f, 0xa5, 0x72, 0x74, 0x05, 0xb2, 0x13, 0x1a, 0x0a, 0xc6, 0x03, 0xcd, 0x72, 0xa1, 0x5d, 0x50,
	0x2d, 0xf8, 0xf2, 0xb0, 0xba, 0x38, 0x66, 0x81, 0xfc, 0x10, 0x47, 0x5a, 0xf5, 0x2e, 0x19, 0x91,
	0x90, 0x06, 0x52, 0x5d, 0x6d, 0xdd, 0xf8, 0x6e, 0xcb, 0x1b, 0xe1, 0x9e, 0x96, 0xa1, 0x2d, 0xc8,
	0x59, 0x23, 0xfd, 0x1c, 0x5a, 0xd0, 0xb7, 0x74, 0xfe, 0xe5, 0x61, 0x75, 0xf9, 0xe1, 0xb6, 0xbd,
	0xa6, 0xc1, 0x18, 0xe8, 0xb7, 0xd1, 0x23, 0x28, 0x9b, 0x67, 0xcc, 0x29, 0x4f, 0x89, 0xcc, 0xd9,
	0x9f, 0x12, 0x76, 0x7c, 0x9c, 0xd7, 0x16, 0x7b, 0xa7, 0xbf, 0x2a, 0x46, 0xb0, 0xf6, 0x8a, 0x1f,
	0x2a, 0x42, 0xda, 0xbe, 0x29, 0x32, 0x38, 0xcd, 0x5c, 0x54, 0x82, 0x05, 0x8f, 0x06, 0xb6, 0x42,
	0xb5, 0x44, 0xdb, 0x30, 0xbd, 0x3a, 0xf5, 0x7b, 0xef, 0xd4, 0xda, 0x0a, 0xb1, 0x8d, 0x7a, 0xfe,
	0xd9, 0x8c, 0x3f, 0xa4, 0x21, 0x9f, 0x64, 0xfe, 0xbf, 0x43, 0x39, 0xfa, 0x00, 0x56, 0x67, 0xba,
	0x45, 0x1f, 0xfc, 0x59, 0x7c, 0xc5, 0x93, 0x5d, 0x63, 0x78, 0x6b, 0xf7, 0x9f, 0x1f, 0x55, 0x52,
	0x2f, 0x8e, 0x2a, 0xa9, 0x5f, 0x8f, 0x2a, 0xa9, 0xa7, 0xc7, 0x95, 0xd4, 0xb3, 0xe3, 0x4a, 0xea,
	0xc7, 0xe3, 0x8a, 0xf3, 0xfc, 0xb8, 0xe2, 0xbc, 0x38, 0xae, 0xa4, 0x7e, 0x3e, 0xae, 0xa4, 0x1e,
	0xdc, 0x1c, 0x30, 0x39, 0x1c, 0x77, 0xeb, 0x3d, 0xee, 0x37, 0x7a, 0x5c, 0xf8, 0x5c, 0x34, 0x58,
	0xb7, 0xb7, 0x35, 0xe0, 0x8d, 0xc9, 0x76, 0xc3, 0xe7, 0xee, 0xd8, 0xa3, 0xc2, 0xfc, 0x59, 0x6c,
	0x45, 0xbf, 0x16, 0xcd, 0xe6, 0x96, 0xf9, 0xbb, 0x90, 0x07, 0x23, 0x2a, 0xba, 0x4b, 0xfa, 0xda,
	0xdb, 0xfe, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x7b, 0x6f, 0x9c, 0xf3, 0x84, 0x0c, 0x00, 0x00,
}
