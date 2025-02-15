package types

import (
	bytes "bytes"
	context "context"
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/duration"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	merkle "github.com/okex/exchain/libs/tendermint/crypto/merkle"
	kv "github.com/okex/exchain/libs/tendermint/libs/kv"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type CheckTxType int32

const (
	CheckTxType_New          CheckTxType = 0
	CheckTxType_Recheck      CheckTxType = 1
	CheckTxType_WrappedCheck CheckTxType = 2
)

var CheckTxType_name = map[int32]string{
	0: "New",
	1: "Recheck",
	2: "WrappedCheck",
}

var CheckTxType_value = map[string]int32{
	"New":          0,
	"Recheck":      1,
	"WrappedCheck": 2,
}

func (x CheckTxType) String() string {
	return proto.EnumName(CheckTxType_name, int32(x))
}

func (CheckTxType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9f1eaa49c51fa1ac, []int{0}
}

type Request struct {
	// Types that are valid to be assigned to Value:
	//	*Request_Echo
	//	*Request_Flush
	//	*Request_Info
	//	*Request_SetOption
	//	*Request_InitChain
	//	*Request_Query
	//	*Request_BeginBlock
	//	*Request_CheckTx
	//	*Request_DeliverTx
	//	*Request_EndBlock
	//	*Request_Commit
	Value                isRequest_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f1eaa49c51fa1ac, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Request.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return m.Size()
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

type isRequest_Value interface {
	isRequest_Value()
	Equal(interface{}) bool
	MarshalTo([]byte) (int, error)
	Size() int
}

type Request_Echo struct {
	Echo *RequestEcho `protobuf:"bytes,2,opt,name=echo,proto3,oneof" json:"echo,omitempty"`
}
type Request_Flush struct {
	Flush *RequestFlush `protobuf:"bytes,3,opt,name=flush,proto3,oneof" json:"flush,omitempty"`
}
type Request_Info struct {
	Info *RequestInfo `protobuf:"bytes,4,opt,name=info,proto3,oneof" json:"info,omitempty"`
}
type Request_SetOption struct {
	SetOption *RequestSetOption `protobuf:"bytes,5,opt,name=set_option,json=setOption,proto3,oneof" json:"set_option,omitempty"`
}
type Request_InitChain struct {
	InitChain *RequestInitChain `protobuf:"bytes,6,opt,name=init_chain,json=initChain,proto3,oneof" json:"init_chain,omitempty"`
}
type Request_Query struct {
	Query *RequestQuery `protobuf:"bytes,7,opt,name=query,proto3,oneof" json:"query,omitempty"`
}
type Request_BeginBlock struct {
	BeginBlock *RequestBeginBlock `protobuf:"bytes,8,opt,name=begin_block,json=beginBlock,proto3,oneof" json:"begin_block,omitempty"`
}
type Request_CheckTx struct {
	CheckTx *RequestCheckTx `protobuf:"bytes,9,opt,name=check_tx,json=checkTx,proto3,oneof" json:"check_tx,omitempty"`
}
type Request_DeliverTx struct {
	DeliverTx *RequestDeliverTx `protobuf:"bytes,19,opt,name=deliver_tx,json=deliverTx,proto3,oneof" json:"deliver_tx,omitempty"`
}
type Request_EndBlock struct {
	EndBlock *RequestEndBlock `protobuf:"bytes,11,opt,name=end_block,json=endBlock,proto3,oneof" json:"end_block,omitempty"`
}
type Request_Commit struct {
	Commit *RequestCommit `protobuf:"bytes,12,opt,name=commit,proto3,oneof" json:"commit,omitempty"`
}

func (*Request_Echo) isRequest_Value()       {}
func (*Request_Flush) isRequest_Value()      {}
func (*Request_Info) isRequest_Value()       {}
func (*Request_SetOption) isRequest_Value()  {}
func (*Request_InitChain) isRequest_Value()  {}
func (*Request_Query) isRequest_Value()      {}
func (*Request_BeginBlock) isRequest_Value() {}
func (*Request_CheckTx) isRequest_Value()    {}
func (*Request_DeliverTx) isRequest_Value()  {}
func (*Request_EndBlock) isRequest_Value()   {}
func (*Request_Commit) isRequest_Value()     {}

func (m *Request) GetValue() isRequest_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Request) GetEcho() *RequestEcho {
	if x, ok := m.GetValue().(*Request_Echo); ok {
		return x.Echo
	}
	return nil
}

func (m *Request) GetFlush() *RequestFlush {
	if x, ok := m.GetValue().(*Request_Flush); ok {
		return x.Flush
	}
	return nil
}

func (m *Request) GetInfo() *RequestInfo {
	if x, ok := m.GetValue().(*Request_Info); ok {
		return x.Info
	}
	return nil
}

func (m *Request) GetSetOption() *RequestSetOption {
	if x, ok := m.GetValue().(*Request_SetOption); ok {
		return x.SetOption
	}
	return nil
}

func (m *Request) GetInitChain() *RequestInitChain {
	if x, ok := m.GetValue().(*Request_InitChain); ok {
		return x.InitChain
	}
	return nil
}

func (m *Request) GetQuery() *RequestQuery {
	if x, ok := m.GetValue().(*Request_Query); ok {
		return x.Query
	}
	return nil
}

func (m *Request) GetBeginBlock() *RequestBeginBlock {
	if x, ok := m.GetValue().(*Request_BeginBlock); ok {
		return x.BeginBlock
	}
	return nil
}

func (m *Request) GetCheckTx() *RequestCheckTx {
	if x, ok := m.GetValue().(*Request_CheckTx); ok {
		return x.CheckTx
	}
	return nil
}

func (m *Request) GetDeliverTx() *RequestDeliverTx {
	if x, ok := m.GetValue().(*Request_DeliverTx); ok {
		return x.DeliverTx
	}
	return nil
}

func (m *Request) GetEndBlock() *RequestEndBlock {
	if x, ok := m.GetValue().(*Request_EndBlock); ok {
		return x.EndBlock
	}
	return nil
}

func (m *Request) GetCommit() *RequestCommit {
	if x, ok := m.GetValue().(*Request_Commit); ok {
		return x.Commit
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Request) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Request_Echo)(nil),
		(*Request_Flush)(nil),
		(*Request_Info)(nil),
		(*Request_SetOption)(nil),
		(*Request_InitChain)(nil),
		(*Request_Query)(nil),
		(*Request_BeginBlock)(nil),
		(*Request_CheckTx)(nil),
		(*Request_DeliverTx)(nil),
		(*Request_EndBlock)(nil),
		(*Request_Commit)(nil),
	}
}

type RequestEcho struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestEcho) Reset()         { *m = RequestEcho{} }
func (m *RequestEcho) String() string { return proto.CompactTextString(m) }
func (*RequestEcho) ProtoMessage()    {}
func (*RequestEcho) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f1eaa49c51fa1ac, []int{1}
}
func (m *RequestEcho) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequestEcho) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RequestEcho.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RequestEcho) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestEcho.Merge(m, src)
}
func (m *RequestEcho) XXX_Size() int {
	return m.Size()
}
func (m *RequestEcho) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestEcho.DiscardUnknown(m)
}

var xxx_messageInfo_RequestEcho proto.InternalMessageInfo

func (m *RequestEcho) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type RequestFlush struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestFlush) Reset()         { *m = RequestFlush{} }
func (m *RequestFlush) String() string { return proto.CompactTextString(m) }
func (*RequestFlush) ProtoMessage()    {}
func (*RequestFlush) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f1eaa49c51fa1ac, []int{2}
}
func (m *RequestFlush) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequestFlush) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RequestFlush.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RequestFlush) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestFlush.Merge(m, src)
}
func (m *RequestFlush) XXX_Size() int {
	return m.Size()
}
func (m *RequestFlush) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestFlush.DiscardUnknown(m)
}

var xxx_messageInfo_RequestFlush proto.InternalMessageInfo

type RequestInfo struct {
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	BlockVersion         uint64   `protobuf:"varint,2,opt,name=block_version,json=blockVersion,proto3" json:"block_version,omitempty"`
	P2PVersion           uint64   `protobuf:"varint,3,opt,name=p2p_version,json=p2pVersion,proto3" json:"p2p_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestInfo) Reset()         { *m = RequestInfo{} }
func (m *RequestInfo) String() string { return proto.CompactTextString(m) }
func (*RequestInfo) ProtoMessage()    {}
func (*RequestInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f1eaa49c51fa1ac, []int{3}
}
func (m *RequestInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequestInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RequestInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RequestInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestInfo.Merge(m, src)
}
func (m *RequestInfo) XXX_Size() int {
	return m.Size()
}
func (m *RequestInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RequestInfo proto.InternalMessageInfo

func (m *RequestInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *RequestInfo) GetBlockVersion() uint64 {
	if m != nil {
		return m.BlockVersion
	}
	return 0
}

func (m *RequestInfo) GetP2PVersion() uint64 {
	if m != nil {
		return m.P2PVersion
	}
	return 0
}

// nondeterministic
type RequestSetOption struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestSetOption) Reset()         { *m = RequestSetOption{} }
func (m *RequestSetOption) String() string { return proto.CompactTextString(m) }
func (*RequestSetOption) ProtoMessage()    {}
func (*RequestSetOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f1eaa49c51fa1ac, []int{4}
}
func (m *RequestSetOption) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequestSetOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RequestSetOption.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RequestSetOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestSetOption.Merge(m, src)
}
func (m *RequestSetOption) XXX_Size() int {
	return m.Size()
}
func (m *RequestSetOption) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestSetOption.DiscardUnknown(m)
}

var xxx_messageInfo_RequestSetOption proto.InternalMessageInfo

func (m *RequestSetOption) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *RequestSetOption) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type RequestInitChain struct {
	Time                 time.Time         `protobuf:"bytes,1,opt,name=time,proto3,stdtime" json:"time"`
	ChainId              string            `protobuf:"bytes,2,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	ConsensusParams      *ConsensusParams  `protobuf:"bytes,3,opt,name=consensus_params,json=consensusParams,proto3" json:"consensus_params,omitempty"`
	Validators           []ValidatorUpdate `protobuf:"bytes,4,rep,name=validators,proto3" json:"validators"`
	AppStateBytes        []byte            `protobuf:"bytes,5,opt,name=app_state_bytes,json=appStateBytes,proto3" json:"app_state_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RequestInitChain) Reset()         { *m = RequestInitChain{} }
func (m *RequestInitChain) String() string { return proto.CompactTextString(m) }
func (*RequestInitChain) ProtoMessage()    {}
func (*RequestInitChain) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f1eaa49c51fa1ac, []int{5}
}
func (m *RequestInitChain) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequestInitChain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RequestInitChain.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RequestInitChain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestInitChain.Merge(m, src)
}
func (m *RequestInitChain) XXX_Size() int {
	return m.Size()
}
func (m *RequestInitChain) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestInitChain.DiscardUnknown(m)
}

var xxx_messageInfo_RequestInitChain proto.InternalMessageInfo

func (m *RequestInitChain) GetTime() time.Time {
	if m != nil {
		return m.Time
	}
	return time.Time{}
}

func (m *RequestInitChain) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *RequestInitChain) GetConsensusParams() *ConsensusParams {
	if m != nil {
		return m.ConsensusParams
	}
	return nil
}

func (m *RequestInitChain) GetValidators() []ValidatorUpdate {
	if m != nil {
		return m.Validators
	}
	return nil
}

func (m *RequestInitChain) GetAppStateBytes() []byte {
	if m != nil {
		return m.AppStateBytes
	}
	return nil
}

type RequestQuery struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Path                 string   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Height               int64    `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	Prove                bool     `protobuf:"varint,4,opt,name=prove,proto3" json:"prove,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestQuery) Reset()         { *m = RequestQuery{} }
func (m *RequestQuery) String() string { return proto.CompactTextString(m) }
func (*RequestQuery) ProtoMessage()    {}
func (*RequestQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f1eaa49c51fa1ac, []int{6}
}
func (m *RequestQuery) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequestQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RequestQuery.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RequestQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestQuery.Merge(m, src)
}
func (m *RequestQuery) XXX_Size() int {
	return m.Size()
}
func (m *RequestQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestQuery.DiscardUnknown(m)
}

var xxx_messageInfo_RequestQuery proto.InternalMessageInfo

func (m *RequestQuery) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *RequestQuery) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *RequestQuery) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *RequestQuery) GetProve() bool {
	if m != nil {
		return m.Prove
	}
	return false
}

type RequestBeginBlock struct {
	Hash                 []byte         `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Header               Header         `protobuf:"bytes,2,opt,name=header,proto3" json:"header"`
	LastCommitInfo       LastCommitInfo `protobuf:"bytes,3,opt,name=last_commit_info,json=lastCommitInfo,proto3" json:"last_commit_info"`
	ByzantineValidators  []Evidence     `protobuf:"bytes,4,rep,name=byzantine_validators,json=byzantineValidators,proto3" json:"byzantine_validators"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *RequestBeginBlock) Reset()         { *m = RequestBeginBlock{} }
func (m *RequestBeginBlock) String() string { return proto.CompactTextString(m) }
func (*RequestBeginBlock) ProtoMessage()    {}
func (*RequestBeginBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f1eaa49c51fa1ac, []int{7}
}
func (m *RequestBeginBlock) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequestBeginBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RequestBeginBlock.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RequestBeginBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestBeginBlock.Merge(m, src)
}
func (m *RequestBeginBlock) XXX_Size() int {
	return m.Size()
}
func (m *RequestBeginBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestBeginBlock.DiscardUnknown(m)
}

var xxx_messageInfo_RequestBeginBlock proto.InternalMessageInfo

func (m *RequestBeginBlock) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *RequestBeginBlock) GetHeader() Header {
	if m != nil {
		return m.Header
	}
	return Header{}
}

func (m *RequestBeginBlock) GetLastCommitInfo() LastCommitInfo {
	if m != nil {
		return m.LastCommitInfo
	}
	return LastCommitInfo{}
}

func (m *RequestBeginBlock) GetByzantineValidators() []Evidence {
	if m != nil {
		return m.ByzantineValidators
	}
	return nil
}

type RequestCheckTx struct {
	Tx                   []byte      `protobuf:"bytes,1,opt,name=tx,proto3" json:"tx,omitempty"`
	Type                 CheckTxType `protobuf:"varint,2,opt,name=type,proto3,enum=tendermint.abci.types.CheckTxType" json:"type,omitempty"`
	From                 string      `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *RequestCheckTx) Reset()         { *m = RequestCheckTx{} }
func (m *RequestCheckTx) String() string { return proto.CompactTextString(m) }
func (*RequestCheckTx) ProtoMessage()    {}
func (*RequestCheckTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f1eaa49c51fa1ac, []int{8}
}
func (m *RequestCheckTx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequestCheckTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RequestCheckTx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RequestCheckTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestCheckTx.Merge(m, src)
}
func (m *RequestCheckTx) XXX_Size() int {
	return m.Size()
}
func (m *RequestCheckTx) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestCheckTx.DiscardUnknown(m)
}

var xxx_messageInfo_RequestCheckTx proto.InternalMessageInfo

func (m *RequestCheckTx) GetTx() []byte {
	if m != nil {
		return m.Tx
	}
	return nil
}

func (m *RequestCheckTx) GetType() CheckTxType {
	if m != nil {
		return m.Type
	}
	return CheckTxType_New
}

func (m *RequestCheckTx) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

type RequestDeliverTx struct {
	Tx                   []byte   `protobuf:"bytes,1,opt,name=tx,proto3" json:"tx,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestDeliverTx) Reset()         { *m = RequestDeliverTx{} }
func (m *RequestDeliverTx) String() string { return proto.CompactTextString(m) }
func (*RequestDeliverTx) ProtoMessage()    {}
func (*RequestDeliverTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f1eaa49c51fa1ac, []int{9}
}
func (m *RequestDeliverTx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequestDeliverTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RequestDeliverTx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RequestDeliverTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestDeliverTx.Merge(m, src)
}
func (m *RequestDeliverTx) XXX_Size() int {
	return m.Size()
}
func (m *RequestDeliverTx) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestDeliverTx.DiscardUnknown(m)
}

var xxx_messageInfo_RequestDeliverTx proto.InternalMessageInfo

func (m *RequestDeliverTx) GetTx() []byte {
	if m != nil {
		return m.Tx
	}
	return nil
}

type RequestEndBlock struct {
	Height               int64    `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestEndBlock) Reset()         { *m = RequestEndBlock{} }
func (m *RequestEndBlock) String() string { return proto.CompactTextString(m) }
func (*RequestEndBlock) ProtoMessage()    {}
func (*RequestEndBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f1eaa49c51fa1ac, []int{10}
}
func (m *RequestEndBlock) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequestEndBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RequestEndBlock.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RequestEndBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestEndBlock.Merge(m, src)
}
func (m *RequestEndBlock) XXX_Size() int {
	return m.Size()
}
func (m *RequestEndBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestEndBlock.DiscardUnknown(m)
}

var xxx_messageInfo_RequestEndBlock proto.InternalMessageInfo

func (m *RequestEndBlock) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *RequestCommit) Reset()         { *m = RequestCommit{} }
func (m *RequestCommit) String() string { return proto.CompactTextString(m) }
func (*RequestCommit) ProtoMessage()    {}
func (*RequestCommit) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f1eaa49c51fa1ac, []int{12}
}
func (m *RequestCommit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequestCommit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RequestCommit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RequestCommit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestCommit.Merge(m, src)
}
func (m *RequestCommit) XXX_Size() int {
	return m.Size()
}
func (m *RequestCommit) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestCommit.DiscardUnknown(m)
}

var xxx_messageInfo_RequestCommit proto.InternalMessageInfo

type Response struct {
	// Types that are valid to be assigned to Value:
	//	*Response_Exception
	//	*Response_Echo
	//	*Response_Flush
	//	*Response_Info
	//	*Response_SetOption
	//	*Response_InitChain
	//	*Response_Query
	//	*Response_BeginBlock
	//	*Response_CheckTx
	//	*Response_DeliverTx
	//	*Response_EndBlock
	//	*Response_Commit
	Value                isResponse_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f1eaa49c51fa1ac, []int{13}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Response.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return m.Size()
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

type isResponse_Value interface {
	isResponse_Value()
	Equal(interface{}) bool
	MarshalTo([]byte) (int, error)
	Size() int
}

type Response_Exception struct {
	Exception *ResponseException `protobuf:"bytes,1,opt,name=exception,proto3,oneof" json:"exception,omitempty"`
}
type Response_Echo struct {
	Echo *ResponseEcho `protobuf:"bytes,2,opt,name=echo,proto3,oneof" json:"echo,omitempty"`
}
type Response_Flush struct {
	Flush *ResponseFlush `protobuf:"bytes,3,opt,name=flush,proto3,oneof" json:"flush,omitempty"`
}
type Response_Info struct {
	Info *ResponseInfo `protobuf:"bytes,4,opt,name=info,proto3,oneof" json:"info,omitempty"`
}
type Response_SetOption struct {
	SetOption *ResponseSetOption `protobuf:"bytes,5,opt,name=set_option,json=setOption,proto3,oneof" json:"set_option,omitempty"`
}
type Response_InitChain struct {
	InitChain *ResponseInitChain `protobuf:"bytes,6,opt,name=init_chain,json=initChain,proto3,oneof" json:"init_chain,omitempty"`
}
type Response_Query struct {
	Query *ResponseQuery `protobuf:"bytes,7,opt,name=query,proto3,oneof" json:"query,omitempty"`
}
type Response_BeginBlock struct {
	BeginBlock *ResponseBeginBlock `protobuf:"bytes,8,opt,name=begin_block,json=beginBlock,proto3,oneof" json:"begin_block,omitempty"`
}
type Response_CheckTx struct {
	CheckTx *ResponseCheckTx `protobuf:"bytes,9,opt,name=check_tx,json=checkTx,proto3,oneof" json:"check_tx,omitempty"`
}
type Response_DeliverTx struct {
	DeliverTx *ResponseDeliverTx `protobuf:"bytes,10,opt,name=deliver_tx,json=deliverTx,proto3,oneof" json:"deliver_tx,omitempty"`
}
type Response_EndBlock struct {
	EndBlock *ResponseEndBlock `protobuf:"bytes,11,opt,name=end_block,json=endBlock,proto3,oneof" json:"end_block,omitempty"`
}
type Response_Commit struct {
	Commit *ResponseCommit `protobuf:"bytes,12,opt,name=commit,proto3,oneof" json:"commit,omitempty"`
}

func (*Response_Exception) isResponse_Value()  {}
func (*Response_Echo) isResponse_Value()       {}
func (*Response_Flush) isResponse_Value()      {}
func (*Response_Info) isResponse_Value()       {}
func (*Response_SetOption) isResponse_Value()  {}
func (*Response_InitChain) isResponse_Value()  {}
func (*Response_Query) isResponse_Value()      {}
func (*Response_BeginBlock) isResponse_Value() {}
func (*Response_CheckTx) isResponse_Value()    {}
func (*Response_DeliverTx) isResponse_Value()  {}
func (*Response_EndBlock) isResponse_Value()   {}
func (*Response_Commit) isResponse_Value()     {}

func (m *Response) GetValue() isResponse_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Response) GetException() *ResponseException {
	if x, ok := m.GetValue().(*Response_Exception); ok {
		return x.Exception
	}
	return nil
}

func (m *Response) GetEcho() *ResponseEcho {
	if x, ok := m.GetValue().(*Response_Echo); ok {
		return x.Echo
	}
	return nil
}

func (m *Response) GetFlush() *ResponseFlush {
	if x, ok := m.GetValue().(*Response_Flush); ok {
		return x.Flush
	}
	return nil
}

func (m *Response) GetInfo() *ResponseInfo {
	if x, ok := m.GetValue().(*Response_Info); ok {
		return x.Info
	}
	return nil
}

func (m *Response) GetSetOption() *ResponseSetOption {
	if x, ok := m.GetValue().(*Response_SetOption); ok {
		return x.SetOption
	}
	return nil
}

func (m *Response) GetInitChain() *ResponseInitChain {
	if x, ok := m.GetValue().(*Response_InitChain); ok {
		return x.InitChain
	}
	return nil
}

func (m *Response) GetQuery() *ResponseQuery {
	if x, ok := m.GetValue().(*Response_Query); ok {
		return x.Query
	}
	return nil
}

func (m *Response) GetBeginBlock() *ResponseBeginBlock {
	if x, ok := m.GetValue().(*Response_BeginBlock); ok {
		return x.BeginBlock
	}
	return nil
}

func (m *Response) GetCheckTx() *ResponseCheckTx {
	if x, ok := m.GetValue().(*Response_CheckTx); ok {
		return x.CheckTx
	}
	return nil
}

func (m *Response) GetDeliverTx() *ResponseDeliverTx {
	if x, ok := m.GetValue().(*Response_DeliverTx); ok {
		return x.DeliverTx
	}
	return nil
}

func (m *Response) GetEndBlock() *ResponseEndBlock {
	if x, ok := m.GetValue().(*Response_EndBlock); ok {
		return x.EndBlock
	}
	return nil
}

func (m *Response) GetCommit() *ResponseCommit {
	if x, ok := m.GetValue().(*Response_Commit); ok {
		return x.Commit
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Response) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Response_Exception)(nil),
		(*Response_Echo)(nil),
		(*Response_Flush)(nil),
		(*Response_Info)(nil),
		(*Response_SetOption)(nil),
		(*Response_InitChain)(nil),
		(*Response_Query)(nil),
		(*Response_BeginBlock)(nil),
		(*Response_CheckTx)(nil),
		(*Response_DeliverTx)(nil),
		(*Response_EndBlock)(nil),
		(*Response_Commit)(nil),
	}
}

// nondeterministic
type ResponseException struct {
	Error                string   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseException) Reset()         { *m = ResponseException{} }
func (m *ResponseException) String() string { return proto.CompactTextString(m) }
func (*ResponseException) ProtoMessage()    {}
func (*ResponseException) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f1eaa49c51fa1ac, []int{14}
}
func (m *ResponseException) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResponseException) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResponseException.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResponseException) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseException.Merge(m, src)
}
func (m *ResponseException) XXX_Size() int {
	return m.Size()
}
func (m *ResponseException) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseException.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseException proto.InternalMessageInfo

func (m *ResponseException) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type ResponseEcho struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseEcho) Reset()         { *m = ResponseEcho{} }
func (m *ResponseEcho) String() string { return proto.CompactTextString(m) }
func (*ResponseEcho) ProtoMessage()    {}
func (*ResponseEcho) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f1eaa49c51fa1ac, []int{15}
}
func (m *ResponseEcho) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResponseEcho) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResponseEcho.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResponseEcho) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseEcho.Merge(m, src)
}
func (m *ResponseEcho) XXX_Size() int {
	return m.Size()
}
func (m *ResponseEcho) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseEcho.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseEcho proto.InternalMessageInfo

func (m *ResponseEcho) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ResponseFlush struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseFlush) Reset()         { *m = ResponseFlush{} }
func (m *ResponseFlush) String() string { return proto.CompactTextString(m) }
func (*ResponseFlush) ProtoMessage()    {}
func (*ResponseFlush) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f1eaa49c51fa1ac, []int{16}
}
func (m *ResponseFlush) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResponseFlush) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResponseFlush.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResponseFlush) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseFlush.Merge(m, src)
}
func (m *ResponseFlush) XXX_Size() int {
	return m.Size()
}
func (m *ResponseFlush) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseFlush.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseFlush proto.InternalMessageInfo

type ResponseInfo struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	AppVersion           uint64   `protobuf:"varint,3,opt,name=app_version,json=appVersion,proto3" json:"app_version,omitempty"`
	LastBlockHeight      int64    `protobuf:"varint,4,opt,name=last_block_height,json=lastBlockHeight,proto3" json:"last_block_height,omitempty"`
	LastBlockAppHash     []byte   `protobuf:"bytes,5,opt,name=last_block_app_hash,json=lastBlockAppHash,proto3" json:"last_block_app_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseInfo) Reset()         { *m = ResponseInfo{} }
func (m *ResponseInfo) String() string { return proto.CompactTextString(m) }
func (*ResponseInfo) ProtoMessage()    {}
func (*ResponseInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f1eaa49c51fa1ac, []int{17}
}
func (m *ResponseInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResponseInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResponseInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResponseInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseInfo.Merge(m, src)
}
func (m *ResponseInfo) XXX_Size() int {
	return m.Size()
}
func (m *ResponseInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseInfo proto.InternalMessageInfo

func (m *ResponseInfo) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *ResponseInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ResponseInfo) GetAppVersion() uint64 {
	if m != nil {
		return m.AppVersion
	}
	return 0
}

func (m *ResponseInfo) GetLastBlockHeight() int64 {
	if m != nil {
		return m.LastBlockHeight
	}
	return 0
}

func (m *ResponseInfo) GetLastBlockAppHash() []byte {
	if m != nil {
		return m.LastBlockAppHash
	}
	return nil
}

// nondeterministic
type ResponseSetOption struct {
	Code uint32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// bytes data = 2;
	Log                  string   `protobuf:"bytes,3,opt,name=log,proto3" json:"log,omitempty"`
	Info                 string   `protobuf:"bytes,4,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseSetOption) Reset()         { *m = ResponseSetOption{} }
func (m *ResponseSetOption) String() string { return proto.CompactTextString(m) }
func (*ResponseSetOption) ProtoMessage()    {}
func (*ResponseSetOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f1eaa49c51fa1ac, []int{18}
}
func (m *ResponseSetOption) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResponseSetOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResponseSetOption.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResponseSetOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseSetOption.Merge(m, src)
}
func (m *ResponseSetOption) XXX_Size() int {
	return m.Size()
}
func (m *ResponseSetOption) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseSetOption.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseSetOption proto.InternalMessageInfo

func (m *ResponseSetOption) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ResponseSetOption) GetLog() string {
	if m != nil {
		return m.Log
	}
	return ""
}

func (m *ResponseSetOption) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

type ResponseInitChain struct {
	ConsensusParams      *ConsensusParams  `protobuf:"bytes,1,opt,name=consensus_params,json=consensusParams,proto3" json:"consensus_params,omitempty"`
	Validators           []ValidatorUpdate `protobuf:"bytes,2,rep,name=validators,proto3" json:"validators"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ResponseInitChain) Reset()         { *m = ResponseInitChain{} }
func (m *ResponseInitChain) String() string { return proto.CompactTextString(m) }
func (*ResponseInitChain) ProtoMessage()    {}
func (*ResponseInitChain) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f1eaa49c51fa1ac, []int{19}
}
func (m *ResponseInitChain) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResponseInitChain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResponseInitChain.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResponseInitChain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseInitChain.Merge(m, src)
}
func (m *ResponseInitChain) XXX_Size() int {
	return m.Size()
}
func (m *ResponseInitChain) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseInitChain.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseInitChain proto.InternalMessageInfo

func (m *ResponseInitChain) GetConsensusParams() *ConsensusParams {
	if m != nil {
		return m.ConsensusParams
	}
	return nil
}

func (m *ResponseInitChain) GetValidators() []ValidatorUpdate {
	if m != nil {
		return m.Validators
	}
	return nil
}

type ResponseQuery struct {
	Code uint32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// bytes data = 2; // use "value" instead.
	Log                  string        `protobuf:"bytes,3,opt,name=log,proto3" json:"log,omitempty"`
	Info                 string        `protobuf:"bytes,4,opt,name=info,proto3" json:"info,omitempty"`
	Index                int64         `protobuf:"varint,5,opt,name=index,proto3" json:"index,omitempty"`
	Key                  []byte        `protobuf:"bytes,6,opt,name=key,proto3" json:"key,omitempty"`
	Value                []byte        `protobuf:"bytes,7,opt,name=value,proto3" json:"value,omitempty"`
	Proof                *merkle.Proof `protobuf:"bytes,8,opt,name=proof,proto3" json:"proof,omitempty"`
	Height               int64         `protobuf:"varint,9,opt,name=height,proto3" json:"height,omitempty"`
	Codespace            string        `protobuf:"bytes,10,opt,name=codespace,proto3" json:"codespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ResponseQuery) Reset()         { *m = ResponseQuery{} }
func (m *ResponseQuery) String() string { return proto.CompactTextString(m) }
func (*ResponseQuery) ProtoMessage()    {}
func (*ResponseQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f1eaa49c51fa1ac, []int{20}
}
func (m *ResponseQuery) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResponseQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResponseQuery.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResponseQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseQuery.Merge(m, src)
}
func (m *ResponseQuery) XXX_Size() int {
	return m.Size()
}
func (m *ResponseQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseQuery.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseQuery proto.InternalMessageInfo

func (m *ResponseQuery) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ResponseQuery) GetLog() string {
	if m != nil {
		return m.Log
	}
	return ""
}

func (m *ResponseQuery) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

func (m *ResponseQuery) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *ResponseQuery) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *ResponseQuery) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *ResponseQuery) GetProof() *merkle.Proof {
	if m != nil {
		return m.Proof
	}
	return nil
}

func (m *ResponseQuery) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *ResponseQuery) GetCodespace() string {
	if m != nil {
		return m.Codespace
	}
	return ""
}

type ResponseBeginBlock struct {
	Events               []Event  `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseBeginBlock) Reset()         { *m = ResponseBeginBlock{} }
func (m *ResponseBeginBlock) String() string { return proto.CompactTextString(m) }
func (*ResponseBeginBlock) ProtoMessage()    {}
func (*ResponseBeginBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f1eaa49c51fa1ac, []int{21}
}
func (m *ResponseBeginBlock) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResponseBeginBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResponseBeginBlock.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResponseBeginBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseBeginBlock.Merge(m, src)
}
func (m *ResponseBeginBlock) XXX_Size() int {
	return m.Size()
}
func (m *ResponseBeginBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseBeginBlock.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseBeginBlock proto.InternalMessageInfo

func (m *ResponseBeginBlock) GetEvents() []Event {
	if m != nil {
		return m.Events
	}
	return nil
}

type ResponseCheckTx struct {
	Code                 uint32   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Log                  string   `protobuf:"bytes,3,opt,name=log,proto3" json:"log,omitempty"`
	Info                 string   `protobuf:"bytes,4,opt,name=info,proto3" json:"info,omitempty"`
	GasWanted            int64    `protobuf:"varint,5,opt,name=gas_wanted,json=gasWanted,proto3" json:"gas_wanted,omitempty"`
	GasUsed              int64    `protobuf:"varint,6,opt,name=gas_used,json=gasUsed,proto3" json:"gas_used,omitempty"`
	Events               []Event  `protobuf:"bytes,7,rep,name=events,proto3" json:"events,omitempty"`
	Codespace            string   `protobuf:"bytes,8,opt,name=codespace,proto3" json:"codespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseCheckTx) Reset()         { *m = ResponseCheckTx{} }
func (m *ResponseCheckTx) String() string { return proto.CompactTextString(m) }
func (*ResponseCheckTx) ProtoMessage()    {}
func (*ResponseCheckTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f1eaa49c51fa1ac, []int{22}
}
func (m *ResponseCheckTx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResponseCheckTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResponseCheckTx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResponseCheckTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseCheckTx.Merge(m, src)
}
func (m *ResponseCheckTx) XXX_Size() int {
	return m.Size()
}
func (m *ResponseCheckTx) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseCheckTx.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseCheckTx proto.InternalMessageInfo

func (m *ResponseCheckTx) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ResponseCheckTx) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *ResponseCheckTx) GetLog() string {
	if m != nil {
		return m.Log
	}
	return ""
}

func (m *ResponseCheckTx) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

func (m *ResponseCheckTx) GetGasWanted() int64 {
	if m != nil {
		return m.GasWanted
	}
	return 0
}

func (m *ResponseCheckTx) GetGasUsed() int64 {
	if m != nil {
		return m.GasUsed
	}
	return 0
}

func (m *ResponseCheckTx) GetEvents() []Event {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *ResponseCheckTx) GetCodespace() string {
	if m != nil {
		return m.Codespace
	}
	return ""
}

type ResponseDeliverTx struct {
	Code                 uint32   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Log                  string   `protobuf:"bytes,3,opt,name=log,proto3" json:"log,omitempty"`
	Info                 string   `protobuf:"bytes,4,opt,name=info,proto3" json:"info,omitempty"`
	GasWanted            int64    `protobuf:"varint,5,opt,name=gas_wanted,json=gasWanted,proto3" json:"gas_wanted,omitempty"`
	GasUsed              int64    `protobuf:"varint,6,opt,name=gas_used,json=gasUsed,proto3" json:"gas_used,omitempty"`
	Events               []Event  `protobuf:"bytes,7,rep,name=events,proto3" json:"events,omitempty"`
	Codespace            string   `protobuf:"bytes,8,opt,name=codespace,proto3" json:"codespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseDeliverTx) Reset()         { *m = ResponseDeliverTx{} }
func (m *ResponseDeliverTx) String() string { return proto.CompactTextString(m) }
func (*ResponseDeliverTx) ProtoMessage()    {}
func (*ResponseDeliverTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f1eaa49c51fa1ac, []int{23}
}
func (m *ResponseDeliverTx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResponseDeliverTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResponseDeliverTx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResponseDeliverTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseDeliverTx.Merge(m, src)
}
func (m *ResponseDeliverTx) XXX_Size() int {
	return m.Size()
}
func (m *ResponseDeliverTx) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseDeliverTx.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseDeliverTx proto.InternalMessageInfo

func (m *ResponseDeliverTx) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ResponseDeliverTx) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *ResponseDeliverTx) GetLog() string {
	if m != nil {
		return m.Log
	}
	return ""
}

func (m *ResponseDeliverTx) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

func (m *ResponseDeliverTx) GetGasWanted() int64 {
	if m != nil {
		return m.GasWanted
	}
	return 0
}

func (m *ResponseDeliverTx) GetGasUsed() int64 {
	if m != nil {
		return m.GasUsed
	}
	return 0
}

func (m *ResponseDeliverTx) GetEvents() []Event {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *ResponseDeliverTx) GetCodespace() string {
	if m != nil {
		return m.Codespace
	}
	return ""
}

type ResponseEndBlock struct {
	ValidatorUpdates      []ValidatorUpdate `protobuf:"bytes,1,rep,name=validator_updates,json=validatorUpdates,proto3" json:"validator_updates"`
	ConsensusParamUpdates *ConsensusParams  `protobuf:"bytes,2,opt,name=consensus_param_updates,json=consensusParamUpdates,proto3" json:"consensus_param_updates,omitempty"`
	Events                []Event           `protobuf:"bytes,3,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}          `json:"-"`
	XXX_unrecognized      []byte            `json:"-"`
	XXX_sizecache         int32             `json:"-"`
}

func (m *ResponseEndBlock) Reset()         { *m = ResponseEndBlock{} }
func (m *ResponseEndBlock) String() string { return proto.CompactTextString(m) }
func (*ResponseEndBlock) ProtoMessage()    {}
func (*ResponseEndBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f1eaa49c51fa1ac, []int{24}
}
func (m *ResponseEndBlock) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResponseEndBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResponseEndBlock.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResponseEndBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseEndBlock.Merge(m, src)
}
func (m *ResponseEndBlock) XXX_Size() int {
	return m.Size()
}
func (m *ResponseEndBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseEndBlock.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseEndBlock proto.InternalMessageInfo

func (m *ResponseEndBlock) GetValidatorUpdates() []ValidatorUpdate {
	if m != nil {
		return m.ValidatorUpdates
	}
	return nil
}

func (m *ResponseEndBlock) GetConsensusParamUpdates() *ConsensusParams {
	if m != nil {
		return m.ConsensusParamUpdates
	}
	return nil
}

func (m *ResponseEndBlock) GetEvents() []Event {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *ResponseCommit) Reset()         { *m = ResponseCommit{} }
func (m *ResponseCommit) String() string { return proto.CompactTextString(m) }
func (*ResponseCommit) ProtoMessage()    {}
func (*ResponseCommit) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f1eaa49c51fa1ac, []int{25}
}
func (m *ResponseCommit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResponseCommit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResponseCommit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResponseCommit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseCommit.Merge(m, src)
}
func (m *ResponseCommit) XXX_Size() int {
	return m.Size()
}
func (m *ResponseCommit) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseCommit.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseCommit proto.InternalMessageInfo

func (m *ResponseCommit) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *ResponseCommit) GetRetainHeight() int64 {
	if m != nil {
		return m.RetainHeight
	}
	return 0
}

// ConsensusParams contains all consensus-relevant parameters
// that can be adjusted by the abci app
type ConsensusParams struct {
	Block                *BlockParams     `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
	Evidence             *EvidenceParams  `protobuf:"bytes,2,opt,name=evidence,proto3" json:"evidence,omitempty"`
	Validator            *ValidatorParams `protobuf:"bytes,3,opt,name=validator,proto3" json:"validator,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ConsensusParams) Reset()         { *m = ConsensusParams{} }
func (m *ConsensusParams) String() string { return proto.CompactTextString(m) }
func (*ConsensusParams) ProtoMessage()    {}
func (*ConsensusParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f1eaa49c51fa1ac, []int{26}
}
func (m *ConsensusParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConsensusParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConsensusParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConsensusParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsensusParams.Merge(m, src)
}
func (m *ConsensusParams) XXX_Size() int {
	return m.Size()
}
func (m *ConsensusParams) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsensusParams.DiscardUnknown(m)
}

var xxx_messageInfo_ConsensusParams proto.InternalMessageInfo

func (m *ConsensusParams) GetBlock() *BlockParams {
	if m != nil {
		return m.Block
	}
	return nil
}

func (m *ConsensusParams) GetEvidence() *EvidenceParams {
	if m != nil {
		return m.Evidence
	}
	return nil
}

func (m *ConsensusParams) GetValidator() *ValidatorParams {
	if m != nil {
		return m.Validator
	}
	return nil
}

// BlockParams contains limits on the block size.
type BlockParams struct {
	// Note: must be greater than 0
	MaxBytes int64 `protobuf:"varint,1,opt,name=max_bytes,json=maxBytes,proto3" json:"max_bytes,omitempty"`
	// Note: must be greater or equal to -1
	MaxGas               int64    `protobuf:"varint,2,opt,name=max_gas,json=maxGas,proto3" json:"max_gas,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockParams) Reset()         { *m = BlockParams{} }
func (m *BlockParams) String() string { return proto.CompactTextString(m) }
func (*BlockParams) ProtoMessage()    {}
func (*BlockParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f1eaa49c51fa1ac, []int{27}
}
func (m *BlockParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlockParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockParams.Merge(m, src)
}
func (m *BlockParams) XXX_Size() int {
	return m.Size()
}
func (m *BlockParams) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockParams.DiscardUnknown(m)
}

var xxx_messageInfo_BlockParams proto.InternalMessageInfo

func (m *BlockParams) GetMaxBytes() int64 {
	if m != nil {
		return m.MaxBytes
	}
	return 0
}

func (m *BlockParams) GetMaxGas() int64 {
	if m != nil {
		return m.MaxGas
	}
	return 0
}

type EvidenceParams struct {
	// Note: must be greater than 0
	MaxAgeNumBlocks      int64         `protobuf:"varint,1,opt,name=max_age_num_blocks,json=maxAgeNumBlocks,proto3" json:"max_age_num_blocks,omitempty"`
	MaxAgeDuration       time.Duration `protobuf:"bytes,2,opt,name=max_age_duration,json=maxAgeDuration,proto3,stdduration" json:"max_age_duration"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *EvidenceParams) Reset()         { *m = EvidenceParams{} }
func (m *EvidenceParams) String() string { return proto.CompactTextString(m) }
func (*EvidenceParams) ProtoMessage()    {}
func (*EvidenceParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f1eaa49c51fa1ac, []int{28}
}
func (m *EvidenceParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EvidenceParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EvidenceParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EvidenceParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvidenceParams.Merge(m, src)
}
func (m *EvidenceParams) XXX_Size() int {
	return m.Size()
}
func (m *EvidenceParams) XXX_DiscardUnknown() {
	xxx_messageInfo_EvidenceParams.DiscardUnknown(m)
}

var xxx_messageInfo_EvidenceParams proto.InternalMessageInfo

func (m *EvidenceParams) GetMaxAgeNumBlocks() int64 {
	if m != nil {
		return m.MaxAgeNumBlocks
	}
	return 0
}

func (m *EvidenceParams) GetMaxAgeDuration() time.Duration {
	if m != nil {
		return m.MaxAgeDuration
	}
	return 0
}

// ValidatorParams contains limits on validators.
type ValidatorParams struct {
	PubKeyTypes          []string `protobuf:"bytes,1,rep,name=pub_key_types,json=pubKeyTypes,proto3" json:"pub_key_types,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidatorParams) Reset()         { *m = ValidatorParams{} }
func (m *ValidatorParams) String() string { return proto.CompactTextString(m) }
func (*ValidatorParams) ProtoMessage()    {}
func (*ValidatorParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f1eaa49c51fa1ac, []int{29}
}
func (m *ValidatorParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorParams.Merge(m, src)
}
func (m *ValidatorParams) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorParams) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorParams.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorParams proto.InternalMessageInfo

func (m *ValidatorParams) GetPubKeyTypes() []string {
	if m != nil {
		return m.PubKeyTypes
	}
	return nil
}

type LastCommitInfo struct {
	Round                int32      `protobuf:"varint,1,opt,name=round,proto3" json:"round,omitempty"`
	Votes                []VoteInfo `protobuf:"bytes,2,rep,name=votes,proto3" json:"votes"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *LastCommitInfo) Reset()         { *m = LastCommitInfo{} }
func (m *LastCommitInfo) String() string { return proto.CompactTextString(m) }
func (*LastCommitInfo) ProtoMessage()    {}
func (*LastCommitInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f1eaa49c51fa1ac, []int{30}
}
func (m *LastCommitInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LastCommitInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LastCommitInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LastCommitInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LastCommitInfo.Merge(m, src)
}
func (m *LastCommitInfo) XXX_Size() int {
	return m.Size()
}
func (m *LastCommitInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LastCommitInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LastCommitInfo proto.InternalMessageInfo

func (m *LastCommitInfo) GetRound() int32 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *LastCommitInfo) GetVotes() []VoteInfo {
	if m != nil {
		return m.Votes
	}
	return nil
}

type Event struct {
	Type                 string    `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Attributes           []kv.Pair `protobuf:"bytes,2,rep,name=attributes,proto3" json:"attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f1eaa49c51fa1ac, []int{31}
}
func (m *Event) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Event.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return m.Size()
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Event) GetAttributes() []kv.Pair {
	if m != nil {
		return m.Attributes
	}
	return nil
}

type Header struct {
	// basic block info
	Version Version   `protobuf:"bytes,1,opt,name=version,proto3" json:"version"`
	ChainID string    `protobuf:"bytes,2,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	Height  int64     `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	Time    time.Time `protobuf:"bytes,4,opt,name=time,proto3,stdtime" json:"time"`
	// prev block info
	LastBlockId BlockID `protobuf:"bytes,5,opt,name=last_block_id,json=lastBlockId,proto3" json:"last_block_id"`
	// hashes of block data
	LastCommitHash []byte `protobuf:"bytes,6,opt,name=last_commit_hash,json=lastCommitHash,proto3" json:"last_commit_hash,omitempty"`
	DataHash       []byte `protobuf:"bytes,7,opt,name=data_hash,json=dataHash,proto3" json:"data_hash,omitempty"`
	// hashes from the app output from the prev block
	ValidatorsHash     []byte `protobuf:"bytes,8,opt,name=validators_hash,json=validatorsHash,proto3" json:"validators_hash,omitempty"`
	NextValidatorsHash []byte `protobuf:"bytes,9,opt,name=next_validators_hash,json=nextValidatorsHash,proto3" json:"next_validators_hash,omitempty"`
	ConsensusHash      []byte `protobuf:"bytes,10,opt,name=consensus_hash,json=consensusHash,proto3" json:"consensus_hash,omitempty"`
	AppHash            []byte `protobuf:"bytes,11,opt,name=app_hash,json=appHash,proto3" json:"app_hash,omitempty"`
	LastResultsHash    []byte `protobuf:"bytes,12,opt,name=last_results_hash,json=lastResultsHash,proto3" json:"last_results_hash,omitempty"`
	// consensus info
	EvidenceHash         []byte   `protobuf:"bytes,13,opt,name=evidence_hash,json=evidenceHash,proto3" json:"evidence_hash,omitempty"`
	ProposerAddress      []byte   `protobuf:"bytes,14,opt,name=proposer_address,json=proposerAddress,proto3" json:"proposer_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Header) Reset()         { *m = Header{} }
func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()    {}
func (*Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f1eaa49c51fa1ac, []int{32}
}
func (m *Header) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Header) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Header.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Header) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Header.Merge(m, src)
}
func (m *Header) XXX_Size() int {
	return m.Size()
}
func (m *Header) XXX_DiscardUnknown() {
	xxx_messageInfo_Header.DiscardUnknown(m)
}

var xxx_messageInfo_Header proto.InternalMessageInfo

func (m *Header) GetVersion() Version {
	if m != nil {
		return m.Version
	}
	return Version{}
}

func (m *Header) GetChainID() string {
	if m != nil {
		return m.ChainID
	}
	return ""
}

func (m *Header) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Header) GetTime() time.Time {
	if m != nil {
		return m.Time
	}
	return time.Time{}
}

func (m *Header) GetLastBlockId() BlockID {
	if m != nil {
		return m.LastBlockId
	}
	return BlockID{}
}

func (m *Header) GetLastCommitHash() []byte {
	if m != nil {
		return m.LastCommitHash
	}
	return nil
}

func (m *Header) GetDataHash() []byte {
	if m != nil {
		return m.DataHash
	}
	return nil
}

func (m *Header) GetValidatorsHash() []byte {
	if m != nil {
		return m.ValidatorsHash
	}
	return nil
}

func (m *Header) GetNextValidatorsHash() []byte {
	if m != nil {
		return m.NextValidatorsHash
	}
	return nil
}

func (m *Header) GetConsensusHash() []byte {
	if m != nil {
		return m.ConsensusHash
	}
	return nil
}

func (m *Header) GetAppHash() []byte {
	if m != nil {
		return m.AppHash
	}
	return nil
}

func (m *Header) GetLastResultsHash() []byte {
	if m != nil {
		return m.LastResultsHash
	}
	return nil
}

func (m *Header) GetEvidenceHash() []byte {
	if m != nil {
		return m.EvidenceHash
	}
	return nil
}

func (m *Header) GetProposerAddress() []byte {
	if m != nil {
		return m.ProposerAddress
	}
	return nil
}

type Version struct {
	Block                uint64   `protobuf:"varint,1,opt,name=Block,proto3" json:"Block,omitempty"`
	App                  uint64   `protobuf:"varint,2,opt,name=App,proto3" json:"App,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Version) Reset()         { *m = Version{} }
func (m *Version) String() string { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()    {}
func (*Version) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f1eaa49c51fa1ac, []int{33}
}
func (m *Version) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Version) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Version.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Version) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Version.Merge(m, src)
}
func (m *Version) XXX_Size() int {
	return m.Size()
}
func (m *Version) XXX_DiscardUnknown() {
	xxx_messageInfo_Version.DiscardUnknown(m)
}

var xxx_messageInfo_Version proto.InternalMessageInfo

func (m *Version) GetBlock() uint64 {
	if m != nil {
		return m.Block
	}
	return 0
}

func (m *Version) GetApp() uint64 {
	if m != nil {
		return m.App
	}
	return 0
}

type BlockID struct {
	Hash                 []byte        `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	PartsHeader          PartSetHeader `protobuf:"bytes,2,opt,name=parts_header,json=partsHeader,proto3" json:"parts_header"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *BlockID) Reset()         { *m = BlockID{} }
func (m *BlockID) String() string { return proto.CompactTextString(m) }
func (*BlockID) ProtoMessage()    {}
func (*BlockID) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f1eaa49c51fa1ac, []int{34}
}
func (m *BlockID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlockID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockID.Merge(m, src)
}
func (m *BlockID) XXX_Size() int {
	return m.Size()
}
func (m *BlockID) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockID.DiscardUnknown(m)
}

var xxx_messageInfo_BlockID proto.InternalMessageInfo

func (m *BlockID) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *BlockID) GetPartsHeader() PartSetHeader {
	if m != nil {
		return m.PartsHeader
	}
	return PartSetHeader{}
}

type PartSetHeader struct {
	Total                int32    `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Hash                 []byte   `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PartSetHeader) Reset()         { *m = PartSetHeader{} }
func (m *PartSetHeader) String() string { return proto.CompactTextString(m) }
func (*PartSetHeader) ProtoMessage()    {}
func (*PartSetHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f1eaa49c51fa1ac, []int{35}
}
func (m *PartSetHeader) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PartSetHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PartSetHeader.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PartSetHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PartSetHeader.Merge(m, src)
}
func (m *PartSetHeader) XXX_Size() int {
	return m.Size()
}
func (m *PartSetHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_PartSetHeader.DiscardUnknown(m)
}

var xxx_messageInfo_PartSetHeader proto.InternalMessageInfo

func (m *PartSetHeader) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *PartSetHeader) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

// Validator
type Validator struct {
	Address []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// PubKey pub_key = 2 [(gogoproto.nullable)=false];
	Power                int64    `protobuf:"varint,3,opt,name=power,proto3" json:"power,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Validator) Reset()         { *m = Validator{} }
func (m *Validator) String() string { return proto.CompactTextString(m) }
func (*Validator) ProtoMessage()    {}
func (*Validator) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f1eaa49c51fa1ac, []int{36}
}
func (m *Validator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Validator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Validator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Validator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Validator.Merge(m, src)
}
func (m *Validator) XXX_Size() int {
	return m.Size()
}
func (m *Validator) XXX_DiscardUnknown() {
	xxx_messageInfo_Validator.DiscardUnknown(m)
}

var xxx_messageInfo_Validator proto.InternalMessageInfo

func (m *Validator) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Validator) GetPower() int64 {
	if m != nil {
		return m.Power
	}
	return 0
}

// ValidatorUpdate
type ValidatorUpdate struct {
	PubKey               PubKey   `protobuf:"bytes,1,opt,name=pub_key,json=pubKey,proto3" json:"pub_key"`
	Power                int64    `protobuf:"varint,2,opt,name=power,proto3" json:"power,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidatorUpdate) Reset()         { *m = ValidatorUpdate{} }
func (m *ValidatorUpdate) String() string { return proto.CompactTextString(m) }
func (*ValidatorUpdate) ProtoMessage()    {}
func (*ValidatorUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f1eaa49c51fa1ac, []int{37}
}
func (m *ValidatorUpdate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorUpdate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorUpdate.Merge(m, src)
}
func (m *ValidatorUpdate) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorUpdate proto.InternalMessageInfo

func (m *ValidatorUpdate) GetPubKey() PubKey {
	if m != nil {
		return m.PubKey
	}
	return PubKey{}
}

func (m *ValidatorUpdate) GetPower() int64 {
	if m != nil {
		return m.Power
	}
	return 0
}

// VoteInfo
type VoteInfo struct {
	Validator            Validator `protobuf:"bytes,1,opt,name=validator,proto3" json:"validator"`
	SignedLastBlock      bool      `protobuf:"varint,2,opt,name=signed_last_block,json=signedLastBlock,proto3" json:"signed_last_block,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *VoteInfo) Reset()         { *m = VoteInfo{} }
func (m *VoteInfo) String() string { return proto.CompactTextString(m) }
func (*VoteInfo) ProtoMessage()    {}
func (*VoteInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f1eaa49c51fa1ac, []int{38}
}
func (m *VoteInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VoteInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VoteInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VoteInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoteInfo.Merge(m, src)
}
func (m *VoteInfo) XXX_Size() int {
	return m.Size()
}
func (m *VoteInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_VoteInfo.DiscardUnknown(m)
}

var xxx_messageInfo_VoteInfo proto.InternalMessageInfo

func (m *VoteInfo) GetValidator() Validator {
	if m != nil {
		return m.Validator
	}
	return Validator{}
}

func (m *VoteInfo) GetSignedLastBlock() bool {
	if m != nil {
		return m.SignedLastBlock
	}
	return false
}

type PubKey struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PubKey) Reset()         { *m = PubKey{} }
func (m *PubKey) String() string { return proto.CompactTextString(m) }
func (*PubKey) ProtoMessage()    {}
func (*PubKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f1eaa49c51fa1ac, []int{39}
}
func (m *PubKey) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PubKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PubKey.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PubKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PubKey.Merge(m, src)
}
func (m *PubKey) XXX_Size() int {
	return m.Size()
}
func (m *PubKey) XXX_DiscardUnknown() {
	xxx_messageInfo_PubKey.DiscardUnknown(m)
}

var xxx_messageInfo_PubKey proto.InternalMessageInfo

func (m *PubKey) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *PubKey) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Evidence struct {
	Type                 string    `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Validator            Validator `protobuf:"bytes,2,opt,name=validator,proto3" json:"validator"`
	Height               int64     `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	Time                 time.Time `protobuf:"bytes,4,opt,name=time,proto3,stdtime" json:"time"`
	TotalVotingPower     int64     `protobuf:"varint,5,opt,name=total_voting_power,json=totalVotingPower,proto3" json:"total_voting_power,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Evidence) Reset()         { *m = Evidence{} }
func (m *Evidence) String() string { return proto.CompactTextString(m) }
func (*Evidence) ProtoMessage()    {}
func (*Evidence) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f1eaa49c51fa1ac, []int{40}
}
func (m *Evidence) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Evidence) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Evidence.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Evidence) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Evidence.Merge(m, src)
}
func (m *Evidence) XXX_Size() int {
	return m.Size()
}
func (m *Evidence) XXX_DiscardUnknown() {
	xxx_messageInfo_Evidence.DiscardUnknown(m)
}

var xxx_messageInfo_Evidence proto.InternalMessageInfo

func (m *Evidence) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Evidence) GetValidator() Validator {
	if m != nil {
		return m.Validator
	}
	return Validator{}
}

func (m *Evidence) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Evidence) GetTime() time.Time {
	if m != nil {
		return m.Time
	}
	return time.Time{}
}

func (m *Evidence) GetTotalVotingPower() int64 {
	if m != nil {
		return m.TotalVotingPower
	}
	return 0
}

func init() {
	proto.RegisterEnum("tendermint.abci.types.CheckTxType", CheckTxType_name, CheckTxType_value)
	golang_proto.RegisterEnum("tendermint.abci.types.CheckTxType", CheckTxType_name, CheckTxType_value)
	proto.RegisterType((*Request)(nil), "tendermint.abci.types.Request")
	golang_proto.RegisterType((*Request)(nil), "tendermint.abci.types.Request")
	proto.RegisterType((*RequestEcho)(nil), "tendermint.abci.types.RequestEcho")
	golang_proto.RegisterType((*RequestEcho)(nil), "tendermint.abci.types.RequestEcho")
	proto.RegisterType((*RequestFlush)(nil), "tendermint.abci.types.RequestFlush")
	golang_proto.RegisterType((*RequestFlush)(nil), "tendermint.abci.types.RequestFlush")
	proto.RegisterType((*RequestInfo)(nil), "tendermint.abci.types.RequestInfo")
	golang_proto.RegisterType((*RequestInfo)(nil), "tendermint.abci.types.RequestInfo")
	proto.RegisterType((*RequestSetOption)(nil), "tendermint.abci.types.RequestSetOption")
	golang_proto.RegisterType((*RequestSetOption)(nil), "tendermint.abci.types.RequestSetOption")
	proto.RegisterType((*RequestInitChain)(nil), "tendermint.abci.types.RequestInitChain")
	golang_proto.RegisterType((*RequestInitChain)(nil), "tendermint.abci.types.RequestInitChain")
	proto.RegisterType((*RequestQuery)(nil), "tendermint.abci.types.RequestQuery")
	golang_proto.RegisterType((*RequestQuery)(nil), "tendermint.abci.types.RequestQuery")
	proto.RegisterType((*RequestBeginBlock)(nil), "tendermint.abci.types.RequestBeginBlock")
	golang_proto.RegisterType((*RequestBeginBlock)(nil), "tendermint.abci.types.RequestBeginBlock")
	proto.RegisterType((*RequestCheckTx)(nil), "tendermint.abci.types.RequestCheckTx")
	golang_proto.RegisterType((*RequestCheckTx)(nil), "tendermint.abci.types.RequestCheckTx")
	proto.RegisterType((*RequestDeliverTx)(nil), "tendermint.abci.types.RequestDeliverTx")
	golang_proto.RegisterType((*RequestDeliverTx)(nil), "tendermint.abci.types.RequestDeliverTx")
	proto.RegisterType((*RequestEndBlock)(nil), "tendermint.abci.types.RequestEndBlock")
	golang_proto.RegisterType((*RequestEndBlock)(nil), "tendermint.abci.types.RequestEndBlock")
	proto.RegisterType((*RequestCommit)(nil), "tendermint.abci.types.RequestCommit")
	golang_proto.RegisterType((*RequestCommit)(nil), "tendermint.abci.types.RequestCommit")
	proto.RegisterType((*Response)(nil), "tendermint.abci.types.Response")
	golang_proto.RegisterType((*Response)(nil), "tendermint.abci.types.Response")
	proto.RegisterType((*ResponseException)(nil), "tendermint.abci.types.ResponseException")
	golang_proto.RegisterType((*ResponseException)(nil), "tendermint.abci.types.ResponseException")
	proto.RegisterType((*ResponseEcho)(nil), "tendermint.abci.types.ResponseEcho")
	golang_proto.RegisterType((*ResponseEcho)(nil), "tendermint.abci.types.ResponseEcho")
	proto.RegisterType((*ResponseFlush)(nil), "tendermint.abci.types.ResponseFlush")
	golang_proto.RegisterType((*ResponseFlush)(nil), "tendermint.abci.types.ResponseFlush")
	proto.RegisterType((*ResponseInfo)(nil), "tendermint.abci.types.ResponseInfo")
	golang_proto.RegisterType((*ResponseInfo)(nil), "tendermint.abci.types.ResponseInfo")
	proto.RegisterType((*ResponseSetOption)(nil), "tendermint.abci.types.ResponseSetOption")
	golang_proto.RegisterType((*ResponseSetOption)(nil), "tendermint.abci.types.ResponseSetOption")
	proto.RegisterType((*ResponseInitChain)(nil), "tendermint.abci.types.ResponseInitChain")
	golang_proto.RegisterType((*ResponseInitChain)(nil), "tendermint.abci.types.ResponseInitChain")
	proto.RegisterType((*ResponseQuery)(nil), "tendermint.abci.types.ResponseQuery")
	golang_proto.RegisterType((*ResponseQuery)(nil), "tendermint.abci.types.ResponseQuery")
	proto.RegisterType((*ResponseBeginBlock)(nil), "tendermint.abci.types.ResponseBeginBlock")
	golang_proto.RegisterType((*ResponseBeginBlock)(nil), "tendermint.abci.types.ResponseBeginBlock")
	proto.RegisterType((*ResponseCheckTx)(nil), "tendermint.abci.types.ResponseCheckTx")
	golang_proto.RegisterType((*ResponseCheckTx)(nil), "tendermint.abci.types.ResponseCheckTx")
	proto.RegisterType((*ResponseDeliverTx)(nil), "tendermint.abci.types.ResponseDeliverTx")
	golang_proto.RegisterType((*ResponseDeliverTx)(nil), "tendermint.abci.types.ResponseDeliverTx")
	proto.RegisterType((*ResponseEndBlock)(nil), "tendermint.abci.types.ResponseEndBlock")
	golang_proto.RegisterType((*ResponseEndBlock)(nil), "tendermint.abci.types.ResponseEndBlock")
	proto.RegisterType((*ResponseCommit)(nil), "tendermint.abci.types.ResponseCommit")
	golang_proto.RegisterType((*ResponseCommit)(nil), "tendermint.abci.types.ResponseCommit")
	proto.RegisterType((*ConsensusParams)(nil), "tendermint.abci.types.ConsensusParams")
	golang_proto.RegisterType((*ConsensusParams)(nil), "tendermint.abci.types.ConsensusParams")
	proto.RegisterType((*BlockParams)(nil), "tendermint.abci.types.BlockParams")
	golang_proto.RegisterType((*BlockParams)(nil), "tendermint.abci.types.BlockParams")
	proto.RegisterType((*EvidenceParams)(nil), "tendermint.abci.types.EvidenceParams")
	golang_proto.RegisterType((*EvidenceParams)(nil), "tendermint.abci.types.EvidenceParams")
	proto.RegisterType((*ValidatorParams)(nil), "tendermint.abci.types.ValidatorParams")
	golang_proto.RegisterType((*ValidatorParams)(nil), "tendermint.abci.types.ValidatorParams")
	proto.RegisterType((*LastCommitInfo)(nil), "tendermint.abci.types.LastCommitInfo")
	golang_proto.RegisterType((*LastCommitInfo)(nil), "tendermint.abci.types.LastCommitInfo")
	proto.RegisterType((*Event)(nil), "tendermint.abci.types.Event")
	golang_proto.RegisterType((*Event)(nil), "tendermint.abci.types.Event")
	proto.RegisterType((*Header)(nil), "tendermint.abci.types.Header")
	golang_proto.RegisterType((*Header)(nil), "tendermint.abci.types.Header")
	proto.RegisterType((*Version)(nil), "tendermint.abci.types.Version")
	golang_proto.RegisterType((*Version)(nil), "tendermint.abci.types.Version")
	proto.RegisterType((*BlockID)(nil), "tendermint.abci.types.BlockID")
	golang_proto.RegisterType((*BlockID)(nil), "tendermint.abci.types.BlockID")
	proto.RegisterType((*PartSetHeader)(nil), "tendermint.abci.types.PartSetHeader")
	golang_proto.RegisterType((*PartSetHeader)(nil), "tendermint.abci.types.PartSetHeader")
	proto.RegisterType((*Validator)(nil), "tendermint.abci.types.Validator")
	golang_proto.RegisterType((*Validator)(nil), "tendermint.abci.types.Validator")
	proto.RegisterType((*ValidatorUpdate)(nil), "tendermint.abci.types.ValidatorUpdate")
	golang_proto.RegisterType((*ValidatorUpdate)(nil), "tendermint.abci.types.ValidatorUpdate")
	proto.RegisterType((*VoteInfo)(nil), "tendermint.abci.types.VoteInfo")
	golang_proto.RegisterType((*VoteInfo)(nil), "tendermint.abci.types.VoteInfo")
	proto.RegisterType((*PubKey)(nil), "tendermint.abci.types.PubKey")
	golang_proto.RegisterType((*PubKey)(nil), "tendermint.abci.types.PubKey")
	proto.RegisterType((*Evidence)(nil), "tendermint.abci.types.Evidence")
	golang_proto.RegisterType((*Evidence)(nil), "tendermint.abci.types.Evidence")
}

func init() { proto.RegisterFile("abci/types/types.proto", fileDescriptor_9f1eaa49c51fa1ac) }
func init() { golang_proto.RegisterFile("abci/types/types.proto", fileDescriptor_9f1eaa49c51fa1ac) }

var fileDescriptor_9f1eaa49c51fa1ac = []byte{
	// 2458 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe4, 0x59, 0x3d, 0x70, 0x1b, 0xc7,
	0xf5, 0xe7, 0x81, 0x20, 0x3e, 0x1e, 0xf8, 0x01, 0xad, 0x64, 0x1b, 0xc6, 0x5f, 0x22, 0x35, 0x47,
	0x4b, 0xa2, 0x6c, 0xff, 0x41, 0x87, 0x19, 0x65, 0xac, 0x48, 0xe3, 0x0c, 0x41, 0x4a, 0x21, 0xc7,
	0x92, 0x4c, 0x9f, 0x24, 0x46, 0x49, 0x66, 0x7c, 0xb3, 0xc0, 0xad, 0x80, 0x1b, 0x02, 0x77, 0xe7,
	0xbb, 0x05, 0x44, 0x64, 0x52, 0xa4, 0xcb, 0x64, 0x26, 0x45, 0xca, 0x34, 0xe9, 0x53, 0xa6, 0x70,
	0xe1, 0x32, 0xa5, 0x8b, 0x14, 0x29, 0x52, 0x2b, 0x09, 0x93, 0x2a, 0xe3, 0x32, 0x93, 0x49, 0x99,
	0x79, 0xbb, 0x7b, 0x5f, 0x20, 0x3e, 0x4e, 0x8e, 0xba, 0x34, 0xc0, 0xed, 0xde, 0x7b, 0x6f, 0x77,
	0xdf, 0xbe, 0x8f, 0xdf, 0x7b, 0x07, 0x6f, 0xd2, 0x56, 0xdb, 0xde, 0xe6, 0x23, 0x8f, 0x05, 0xf2,
	0xb7, 0xe1, 0xf9, 0x2e, 0x77, 0xc9, 0x1b, 0x9c, 0x39, 0x16, 0xf3, 0xfb, 0xb6, 0xc3, 0x1b, 0x48,
	0xd2, 0x10, 0x2f, 0xeb, 0xd7, 0x79, 0xd7, 0xf6, 0x2d, 0xd3, 0xa3, 0x3e, 0x1f, 0x6d, 0x0b, 0xca,
	0xed, 0x8e, 0xdb, 0x71, 0xe3, 0x27, 0xc9, 0x5e, 0xaf, 0xb7, 0xfd, 0x91, 0xc7, 0xdd, 0xed, 0x3e,
	0xf3, 0x4f, 0x7a, 0x4c, 0xfd, 0xa9, 0x77, 0x17, 0x7b, 0x76, 0x2b, 0xd8, 0x3e, 0x19, 0x26, 0xd7,
	0xab, 0x6f, 0x74, 0x5c, 0xb7, 0xd3, 0x63, 0x52, 0x66, 0x6b, 0xf0, 0x7c, 0x9b, 0xdb, 0x7d, 0x16,
	0x70, 0xda, 0xf7, 0x14, 0xc1, 0xfa, 0x38, 0x81, 0x35, 0xf0, 0x29, 0xb7, 0x5d, 0x47, 0xbe, 0xd7,
	0xff, 0xb5, 0x04, 0x45, 0x83, 0x7d, 0x3e, 0x60, 0x01, 0x27, 0x1f, 0x42, 0x9e, 0xb5, 0xbb, 0x6e,
	0x2d, 0x77, 0x55, 0xdb, 0xaa, 0xec, 0xe8, 0x8d, 0x89, 0x67, 0x69, 0x28, 0xea, 0x7b, 0xed, 0xae,
	0x7b, 0xb0, 0x60, 0x08, 0x0e, 0x72, 0x07, 0x96, 0x9e, 0xf7, 0x06, 0x41, 0xb7, 0xb6, 0x28, 0x58,
	0x37, 0x67, 0xb3, 0xde, 0x47, 0xd2, 0x83, 0x05, 0x43, 0xf2, 0xe0, 0xb2, 0xb6, 0xf3, 0xdc, 0xad,
	0xe5, 0xb3, 0x2c, 0x7b, 0xe8, 0x3c, 0x17, 0xcb, 0x22, 0x07, 0x39, 0x00, 0x08, 0x18, 0x37, 0x5d,
	0x0f, 0x0f, 0x54, 0x5b, 0x12, 0xfc, 0x37, 0x66, 0xf3, 0x3f, 0x66, 0xfc, 0x13, 0x41, 0x7e, 0xb0,
	0x60, 0x94, 0x83, 0x70, 0x80, 0x92, 0x6c, 0xc7, 0xe6, 0x66, 0xbb, 0x4b, 0x6d, 0xa7, 0x56, 0xc8,
	0x22, 0xe9, 0xd0, 0xb1, 0xf9, 0x1e, 0x92, 0xa3, 0x24, 0x3b, 0x1c, 0xa0, 0x2a, 0x3e, 0x1f, 0x30,
	0x7f, 0x54, 0x2b, 0x66, 0x51, 0xc5, 0xa7, 0x48, 0x8a, 0xaa, 0x10, 0x3c, 0xe4, 0x63, 0xa8, 0xb4,
	0x58, 0xc7, 0x76, 0xcc, 0x56, 0xcf, 0x6d, 0x9f, 0xd4, 0x4a, 0x42, 0xc4, 0xd6, 0x6c, 0x11, 0x4d,
	0x64, 0x68, 0x22, 0xfd, 0xc1, 0x82, 0x01, 0xad, 0x68, 0x44, 0x9a, 0x50, 0x6a, 0x77, 0x59, 0xfb,
	0xc4, 0xe4, 0xa7, 0xb5, 0xb2, 0x90, 0x74, 0x6d, 0xb6, 0xa4, 0x3d, 0xa4, 0x7e, 0x72, 0x7a, 0xb0,
	0x60, 0x14, 0xdb, 0xf2, 0x11, 0xf5, 0x62, 0xb1, 0x9e, 0x3d, 0x64, 0x3e, 0x4a, 0xb9, 0x98, 0x45,
	0x2f, 0xfb, 0x92, 0x5e, 0xc8, 0x29, 0x5b, 0xe1, 0x80, 0xdc, 0x83, 0x32, 0x73, 0x2c, 0x75, 0xb0,
	0x8a, 0x10, 0x74, 0x7d, 0x8e, 0x85, 0x39, 0x56, 0x78, 0xac, 0x12, 0x53, 0xcf, 0xe4, 0x23, 0x28,
	0xb4, 0xdd, 0x7e, 0xdf, 0xe6, 0xb5, 0x65, 0x21, 0xe3, 0x9d, 0x39, 0x47, 0x12, 0xb4, 0x07, 0x0b,
	0x86, 0xe2, 0x6a, 0x16, 0x61, 0x69, 0x48, 0x7b, 0x03, 0xa6, 0xdf, 0x80, 0x4a, 0xc2, 0x92, 0x49,
	0x0d, 0x8a, 0x7d, 0x16, 0x04, 0xb4, 0xc3, 0x6a, 0xda, 0x55, 0x6d, 0xab, 0x6c, 0x84, 0x43, 0x7d,
	0x15, 0x96, 0x93, 0x76, 0xab, 0xf7, 0x23, 0x46, 0xb4, 0x45, 0x64, 0x1c, 0x32, 0x3f, 0x40, 0x03,
	0x54, 0x8c, 0x6a, 0x48, 0x36, 0x61, 0x45, 0x9c, 0xd6, 0x0c, 0xdf, 0xa3, 0x5f, 0xe5, 0x8d, 0x65,
	0x31, 0x79, 0xac, 0x88, 0x36, 0xa0, 0xe2, 0xed, 0x78, 0x11, 0xc9, 0xa2, 0x20, 0x01, 0x6f, 0xc7,
	0x53, 0x04, 0xfa, 0x77, 0xa1, 0x3a, 0x6e, 0xba, 0xa4, 0x0a, 0x8b, 0x27, 0x6c, 0xa4, 0xd6, 0xc3,
	0x47, 0x72, 0x49, 0x1d, 0x4b, 0xac, 0x51, 0x36, 0xd4, 0x19, 0x7f, 0x97, 0x8b, 0x98, 0x23, 0x6b,
	0x45, 0x77, 0xc3, 0x20, 0x21, 0xb8, 0x2b, 0x3b, 0xf5, 0x86, 0x0c, 0x10, 0x8d, 0x30, 0x40, 0x34,
	0x9e, 0x84, 0x11, 0xa4, 0x59, 0xfa, 0xea, 0xe5, 0xc6, 0xc2, 0xaf, 0xfe, 0xbc, 0xa1, 0x19, 0x82,
	0x83, 0xbc, 0x8d, 0x06, 0x45, 0x6d, 0xc7, 0xb4, 0x2d, 0xb5, 0x4e, 0x51, 0x8c, 0x0f, 0x2d, 0xf2,
	0x29, 0x54, 0xdb, 0xae, 0x13, 0x30, 0x27, 0x18, 0x04, 0x18, 0xe6, 0x68, 0x3f, 0x50, 0xb1, 0x60,
	0xda, 0x25, 0xef, 0x85, 0xe4, 0x47, 0x82, 0xda, 0x58, 0x6b, 0xa7, 0x27, 0xc8, 0x03, 0x80, 0x21,
	0xed, 0xd9, 0x16, 0xe5, 0xae, 0x1f, 0xd4, 0xf2, 0x57, 0x17, 0x67, 0x08, 0x3b, 0x0e, 0x09, 0x9f,
	0x7a, 0x16, 0xe5, 0xac, 0x99, 0xc7, 0x9d, 0x1b, 0x09, 0x7e, 0x72, 0x1d, 0xd6, 0xa8, 0xe7, 0x99,
	0x01, 0xa7, 0x9c, 0x99, 0xad, 0x11, 0x67, 0x81, 0x88, 0x17, 0xcb, 0xc6, 0x0a, 0xf5, 0xbc, 0xc7,
	0x38, 0xdb, 0xc4, 0x49, 0xdd, 0x8a, 0x6e, 0x5b, 0xb8, 0x26, 0x21, 0x90, 0xb7, 0x28, 0xa7, 0x42,
	0x5b, 0xcb, 0x86, 0x78, 0xc6, 0x39, 0x8f, 0xf2, 0xae, 0xd2, 0x81, 0x78, 0x26, 0x6f, 0x42, 0xa1,
	0xcb, 0xec, 0x4e, 0x97, 0x8b, 0x63, 0x2f, 0x1a, 0x6a, 0x84, 0x17, 0xe3, 0xf9, 0xee, 0x90, 0x89,
	0xe8, 0x56, 0x32, 0xe4, 0x40, 0xff, 0x22, 0x07, 0x17, 0xce, 0xb9, 0x2f, 0xca, 0xed, 0xd2, 0xa0,
	0x1b, 0xae, 0x85, 0xcf, 0xe4, 0x0e, 0xca, 0xa5, 0x16, 0xf3, 0x55, 0x54, 0xbe, 0x32, 0x45, 0x03,
	0x07, 0x82, 0x48, 0x1d, 0x5c, 0xb1, 0x90, 0xa7, 0x50, 0xed, 0xd1, 0x80, 0x9b, 0xd2, 0xf6, 0x4d,
	0x11, 0x65, 0x17, 0x67, 0x46, 0x82, 0x07, 0x34, 0xf4, 0x19, 0x34, 0x6e, 0x25, 0x6e, 0xb5, 0x97,
	0x9a, 0x25, 0xcf, 0xe0, 0x52, 0x6b, 0xf4, 0x13, 0xea, 0x70, 0xdb, 0x61, 0xe6, 0xb9, 0x3b, 0xda,
	0x98, 0x22, 0xfa, 0xde, 0xd0, 0xb6, 0x98, 0xd3, 0x0e, 0x2f, 0xe7, 0x62, 0x24, 0xe2, 0x38, 0xbe,
	0xa5, 0x2b, 0x00, 0x83, 0x80, 0x99, 0x16, 0xeb, 0x71, 0x2a, 0x2f, 0xa8, 0x64, 0x94, 0x07, 0x01,
	0xdb, 0x17, 0x13, 0xfa, 0x33, 0x58, 0x4d, 0x87, 0x2a, 0xb2, 0x0a, 0x39, 0x7e, 0xaa, 0x14, 0x96,
	0xe3, 0xa7, 0xe4, 0x3b, 0x90, 0xc7, 0xd5, 0x84, 0xb2, 0x56, 0xa7, 0xe6, 0x12, 0xc5, 0xfd, 0x64,
	0xe4, 0x31, 0x43, 0xd0, 0xeb, 0x7a, 0xe4, 0x28, 0x51, 0xf8, 0x1a, 0x97, 0xad, 0xdf, 0x84, 0xb5,
	0xb1, 0xc8, 0x94, 0xb8, 0x75, 0x2d, 0x79, 0xeb, 0xfa, 0x3e, 0x14, 0xe4, 0x96, 0xd1, 0x67, 0x70,
	0x61, 0xd3, 0x0f, 0x3c, 0x25, 0xaa, 0x88, 0x63, 0x23, 0xf0, 0xd0, 0xf5, 0xe5, 0x41, 0x85, 0x3d,
	0x8a, 0x2d, 0x2f, 0x1b, 0x20, 0xa7, 0xd0, 0x18, 0xf5, 0xfb, 0xb0, 0x92, 0x0a, 0x63, 0xe4, 0x16,
	0x14, 0x94, 0x6a, 0xb4, 0x99, 0xc6, 0x20, 0xd7, 0x36, 0x14, 0xb1, 0xfe, 0x87, 0x02, 0x94, 0x0c,
	0x16, 0x78, 0xe8, 0x60, 0xe4, 0x00, 0xca, 0xec, 0xb4, 0xcd, 0x64, 0xca, 0xd4, 0xe6, 0x24, 0x18,
	0xc9, 0x73, 0x2f, 0xa4, 0xc7, 0x88, 0x1e, 0x31, 0x93, 0xdb, 0x29, 0xb8, 0xb0, 0x39, 0x4f, 0x48,
	0x12, 0x2f, 0xdc, 0x4d, 0xe3, 0x85, 0x77, 0xe6, 0xf0, 0x8e, 0x01, 0x86, 0xdb, 0x29, 0xc0, 0x30,
	0x6f, 0xe1, 0x14, 0x62, 0x38, 0x9c, 0x80, 0x18, 0xe6, 0x1d, 0x7f, 0x0a, 0x64, 0x38, 0x9c, 0x00,
	0x19, 0xb6, 0xe6, 0xee, 0x65, 0x22, 0x66, 0xb8, 0x9b, 0xc6, 0x0c, 0xf3, 0xd4, 0x31, 0x06, 0x1a,
	0x1e, 0x4c, 0x02, 0x0d, 0x37, 0xe7, 0xc8, 0x98, 0x8a, 0x1a, 0xf6, 0xce, 0xa1, 0x86, 0xeb, 0x73,
	0x44, 0x4d, 0x80, 0x0d, 0x87, 0x29, 0xd8, 0x00, 0x99, 0x74, 0x33, 0x05, 0x37, 0xdc, 0x3f, 0x8f,
	0x1b, 0x6e, 0xcc, 0x33, 0xb5, 0x49, 0xc0, 0xe1, 0x7b, 0x63, 0xc0, 0xe1, 0xda, 0xbc, 0x53, 0x4d,
	0x45, 0x0e, 0x37, 0x31, 0x76, 0x8f, 0x79, 0x06, 0xc6, 0x79, 0xe6, 0xfb, 0xae, 0xaf, 0x92, 0xb2,
	0x1c, 0xe8, 0x5b, 0x98, 0x4d, 0x62, 0xfb, 0x9f, 0x81, 0x32, 0xd6, 0xd0, 0xd7, 0x13, 0xd6, 0xae,
	0x7f, 0xa9, 0xc5, 0xbc, 0x22, 0xea, 0x26, 0x33, 0x51, 0x59, 0x65, 0xa2, 0x04, 0xf8, 0xc8, 0xa5,
	0xc1, 0xc7, 0x06, 0x54, 0x30, 0xdf, 0x8d, 0xe1, 0x0a, 0xea, 0x85, 0xb8, 0x82, 0xbc, 0x0b, 0x17,
	0x44, 0x6e, 0x90, 0x10, 0x45, 0x45, 0xb1, 0xbc, 0x88, 0x62, 0x6b, 0xf8, 0x42, 0x6a, 0x50, 0x26,
	0xb1, 0xff, 0x87, 0x8b, 0x09, 0x5a, 0x94, 0x2b, 0xf2, 0x94, 0x4c, 0xa0, 0xd5, 0x88, 0x7a, 0xd7,
	0xf3, 0x0e, 0x68, 0xd0, 0xd5, 0x1f, 0xc6, 0x0a, 0x8a, 0x31, 0x0b, 0x81, 0x7c, 0xdb, 0xb5, 0xe4,
	0xb9, 0x57, 0x0c, 0xf1, 0x8c, 0x38, 0xa6, 0xe7, 0x76, 0xc4, 0xe6, 0xca, 0x06, 0x3e, 0x22, 0x55,
	0xe4, 0xda, 0x65, 0xe9, 0xb3, 0xfa, 0x17, 0x5a, 0x2c, 0x2f, 0x86, 0x31, 0x93, 0x10, 0x87, 0xf6,
	0x3a, 0x11, 0x47, 0xee, 0xbf, 0x43, 0x1c, 0xfa, 0x3f, 0xb5, 0xf8, 0x4a, 0x23, 0x2c, 0xf1, 0xcd,
	0x54, 0x80, 0xd6, 0x65, 0x3b, 0x16, 0x3b, 0x15, 0x2a, 0x5f, 0x34, 0xe4, 0x20, 0x84, 0x81, 0x05,
	0x71, 0x0d, 0x69, 0x18, 0x58, 0x14, 0x73, 0x72, 0x40, 0x6e, 0x09, 0x0c, 0xe2, 0x3e, 0x57, 0xa1,
	0x21, 0x95, 0xa0, 0x65, 0xc1, 0xd9, 0x50, 0x95, 0xe6, 0x11, 0x92, 0x19, 0x92, 0x3a, 0x91, 0xdc,
	0xca, 0x29, 0x48, 0x73, 0x19, 0xca, 0xb8, 0xf5, 0xc0, 0xa3, 0x6d, 0x26, 0x7c, 0xbb, 0x6c, 0xc4,
	0x13, 0xba, 0x05, 0xe4, 0x7c, 0x8c, 0x21, 0x8f, 0xa0, 0xc0, 0x86, 0xcc, 0xe1, 0x78, 0x47, 0xa8,
	0xd6, 0xcb, 0x53, 0x41, 0x02, 0x73, 0x78, 0xb3, 0x86, 0xca, 0xfc, 0xc7, 0xcb, 0x8d, 0xaa, 0xe4,
	0x79, 0xdf, 0xed, 0xdb, 0x9c, 0xf5, 0x3d, 0x3e, 0x32, 0x94, 0x14, 0xfd, 0xe7, 0x39, 0x4c, 0xc6,
	0xa9, 0xf8, 0x33, 0x51, 0xbd, 0xa1, 0xd3, 0xe4, 0x12, 0xf0, 0x2d, 0x9b, 0xca, 0xaf, 0x00, 0x74,
	0x68, 0x60, 0xbe, 0xa0, 0x0e, 0x67, 0x96, 0xd2, 0x7b, 0xb9, 0x43, 0x83, 0x1f, 0x88, 0x09, 0xcc,
	0xeb, 0xf8, 0x7a, 0x10, 0x30, 0x4b, 0x5c, 0xc0, 0xa2, 0x51, 0xec, 0xd0, 0xe0, 0x69, 0xc0, 0xac,
	0xc4, 0x59, 0x8b, 0xaf, 0xe3, 0xac, 0x69, 0x7d, 0x97, 0xc6, 0xf5, 0xfd, 0x8b, 0x5c, 0xec, 0x1d,
	0x31, 0x76, 0xf9, 0xdf, 0xd4, 0xc5, 0x6f, 0x44, 0xbd, 0x93, 0x4e, 0x02, 0xe4, 0x87, 0x70, 0x21,
	0xf2, 0x4a, 0x73, 0x20, 0xbc, 0x35, 0xb4, 0xc2, 0x57, 0x73, 0xee, 0xea, 0x30, 0x3d, 0x1d, 0x90,
	0xcf, 0xe0, 0xad, 0xb1, 0x18, 0x14, 0x2d, 0x90, 0x7b, 0xa5, 0x50, 0xf4, 0x46, 0x3a, 0x14, 0x85,
	0xf2, 0x63, 0xed, 0x2d, 0xbe, 0x16, 0xaf, 0xf9, 0x99, 0x86, 0x00, 0x3a, 0x99, 0xdf, 0x26, 0x1a,
	0xc5, 0x26, 0xac, 0xf8, 0x8c, 0x63, 0xa1, 0x97, 0x2a, 0x69, 0x96, 0xe5, 0xa4, 0xca, 0x09, 0x31,
	0x16, 0xcd, 0xbf, 0x0a, 0x16, 0xfd, 0x93, 0x06, 0x6b, 0x63, 0xa7, 0x27, 0x1f, 0xc2, 0x92, 0x4c,
	0xef, 0xda, 0xcc, 0x0e, 0x90, 0xb8, 0x4e, 0xa5, 0x30, 0xc9, 0x40, 0x76, 0xa1, 0xc4, 0x54, 0x59,
	0xa1, 0x34, 0x7e, 0x6d, 0x4e, 0xf5, 0xa1, 0xf8, 0x23, 0x36, 0xb2, 0x0f, 0xe5, 0xe8, 0x5e, 0xe7,
	0x94, 0xac, 0x91, 0x59, 0x28, 0x21, 0x31, 0xa3, 0xbe, 0x07, 0x95, 0xc4, 0xf6, 0xc8, 0xff, 0x41,
	0xb9, 0x4f, 0x4f, 0x55, 0x9d, 0x29, 0x4b, 0x83, 0x52, 0x9f, 0x9e, 0x8a, 0x12, 0x93, 0xbc, 0x05,
	0x45, 0x7c, 0xd9, 0xa1, 0xd2, 0x4a, 0x16, 0x8d, 0x42, 0x9f, 0x9e, 0x7e, 0x9f, 0x06, 0xfa, 0x2f,
	0x35, 0x58, 0x4d, 0xef, 0x93, 0xbc, 0x07, 0x04, 0x69, 0x69, 0x87, 0x99, 0xce, 0xa0, 0x2f, 0x13,
	0x70, 0x28, 0x71, 0xad, 0x4f, 0x4f, 0x77, 0x3b, 0xec, 0xd1, 0xa0, 0x2f, 0x96, 0x0e, 0xc8, 0x43,
	0xa8, 0x86, 0xc4, 0x61, 0x97, 0x4f, 0x69, 0xe5, 0xed, 0x73, 0x55, 0xfe, 0xbe, 0x22, 0x90, 0x45,
	0xfe, 0xaf, 0xb1, 0xc8, 0x5f, 0x95, 0xf2, 0xc2, 0x37, 0xfa, 0x2d, 0x58, 0x1b, 0x3b, 0x31, 0xd1,
	0x61, 0xc5, 0x1b, 0xb4, 0xcc, 0x13, 0x36, 0x32, 0x85, 0x4a, 0x84, 0x1f, 0x95, 0x8d, 0x8a, 0x37,
	0x68, 0x7d, 0xcc, 0x46, 0x58, 0x4f, 0x05, 0x7a, 0x1b, 0x56, 0xd3, 0x55, 0x24, 0x66, 0x25, 0xdf,
	0x1d, 0x38, 0x96, 0xd8, 0xf7, 0x92, 0x21, 0x07, 0xe4, 0x0e, 0x2c, 0x0d, 0x5d, 0xe9, 0x2a, 0xb3,
	0xca, 0xc6, 0x63, 0x97, 0xb3, 0x44, 0x2d, 0x2a, 0x79, 0xf4, 0x00, 0x96, 0x84, 0xd1, 0xa3, 0xfd,
	0x8a, 0x82, 0x4f, 0xa1, 0x22, 0x7c, 0x26, 0xc7, 0x00, 0x94, 0x73, 0xdf, 0x6e, 0x0d, 0x62, 0xf1,
	0xb5, 0xa4, 0xf8, 0x9e, 0xdd, 0x0a, 0x1a, 0x27, 0xc3, 0xc6, 0x11, 0xb5, 0xfd, 0xe6, 0x65, 0xe5,
	0x36, 0x97, 0x62, 0x9e, 0x84, 0xeb, 0x24, 0x24, 0xe9, 0x5f, 0xe7, 0xa1, 0x20, 0xeb, 0x6c, 0xf2,
	0x51, 0xba, 0xeb, 0x53, 0xd9, 0x59, 0x9f, 0xb6, 0x7d, 0x49, 0xa5, 0x76, 0x1f, 0xc1, 0xb3, 0xeb,
	0xe3, 0xad, 0x94, 0x66, 0xe5, 0xec, 0xe5, 0x46, 0x51, 0x40, 0x9b, 0xc3, 0xfd, 0xb8, 0xaf, 0x32,
	0xad, 0xad, 0x10, 0x36, 0x71, 0xf2, 0xaf, 0xdc, 0xc4, 0x39, 0x80, 0x95, 0x04, 0x96, 0xb3, 0x2d,
	0x55, 0x04, 0xad, 0xcf, 0x72, 0xba, 0xc3, 0x7d, 0xb5, 0xff, 0x4a, 0x84, 0xf5, 0x0e, 0x2d, 0xb2,
	0x95, 0xee, 0x2e, 0x08, 0x48, 0x28, 0xb1, 0x48, 0xa2, 0x61, 0x80, 0x80, 0x10, 0xdd, 0x01, 0x03,
	0x8b, 0x24, 0x91, 0xd0, 0xa4, 0x84, 0x13, 0xe2, 0xe5, 0x0d, 0x58, 0x8b, 0x51, 0x93, 0x24, 0x29,
	0x49, 0x29, 0xf1, 0xb4, 0x20, 0xfc, 0x00, 0x2e, 0x39, 0xec, 0x94, 0x9b, 0xe3, 0xd4, 0x65, 0x41,
	0x4d, 0xf0, 0xdd, 0x71, 0x9a, 0xe3, 0x1a, 0xac, 0xc6, 0xf1, 0x59, 0xd0, 0x82, 0xec, 0xf9, 0x44,
	0xb3, 0x82, 0x0c, 0x6b, 0xf4, 0x10, 0xd3, 0x56, 0x54, 0x8d, 0x2e, 0xa1, 0x6c, 0x84, 0x92, 0x7d,
	0x16, 0x0c, 0x7a, 0x5c, 0x09, 0x59, 0x16, 0x34, 0x02, 0x25, 0x1b, 0x72, 0x5e, 0xd0, 0x6e, 0xc2,
	0x4a, 0x18, 0x55, 0x24, 0xdd, 0x8a, 0xa0, 0x5b, 0x0e, 0x27, 0x05, 0xd1, 0x4d, 0xa8, 0x7a, 0xbe,
	0xeb, 0xb9, 0x01, 0xf3, 0x4d, 0x6a, 0x59, 0x3e, 0x0b, 0x82, 0xda, 0xaa, 0x94, 0x17, 0xce, 0xef,
	0xca, 0x69, 0xfd, 0x5b, 0x50, 0x0c, 0xc1, 0xfa, 0x25, 0x58, 0x6a, 0x46, 0x11, 0x32, 0x6f, 0xc8,
	0x01, 0x26, 0xef, 0x5d, 0xcf, 0x53, 0x6d, 0x45, 0x7c, 0xd4, 0x7b, 0x50, 0x54, 0x17, 0x36, 0xb1,
	0x99, 0xf4, 0x10, 0x96, 0x3d, 0xea, 0xe3, 0x31, 0x92, 0x2d, 0xa5, 0x69, 0xe5, 0xe6, 0x11, 0xf5,
	0xf9, 0x63, 0xc6, 0x53, 0x9d, 0xa5, 0x8a, 0xe0, 0x97, 0x53, 0xfa, 0x6d, 0x58, 0x49, 0xd1, 0xe0,
	0x36, 0xb9, 0xcb, 0x69, 0x2f, 0x74, 0x74, 0x31, 0x88, 0x76, 0x92, 0x8b, 0x77, 0xa2, 0xdf, 0x81,
	0x72, 0x74, 0x57, 0x58, 0xc5, 0x84, 0xaa, 0x08, 0x5b, 0x24, 0x72, 0x28, 0xba, 0x67, 0xee, 0x0b,
	0xe6, 0x2b, 0xeb, 0x97, 0x03, 0x9d, 0x25, 0x02, 0x93, 0x4c, 0x95, 0xe4, 0x2e, 0x14, 0x55, 0x60,
	0x9a, 0xd3, 0x1a, 0x39, 0x12, 0x91, 0x2a, 0xec, 0x93, 0xc9, 0xb8, 0x15, 0x2f, 0x93, 0x4b, 0x2e,
	0xf3, 0x53, 0x28, 0x85, 0xc1, 0x27, 0x9d, 0x25, 0xe4, 0x0a, 0x57, 0xe7, 0x65, 0x09, 0xb5, 0x48,
	0xcc, 0x88, 0xd6, 0x14, 0xd8, 0x1d, 0x87, 0x59, 0x66, 0xec, 0x82, 0x62, 0xcd, 0x92, 0xb1, 0x26,
	0x5f, 0x3c, 0x08, 0xfd, 0x4b, 0xff, 0x00, 0x0a, 0x72, 0xaf, 0x13, 0x43, 0xdc, 0x84, 0xb4, 0xad,
	0xff, 0x5d, 0x83, 0x52, 0x98, 0x3e, 0x26, 0x32, 0xa5, 0x0e, 0x91, 0xfb, 0xa6, 0x87, 0x78, 0xfd,
	0x21, 0xe9, 0x7d, 0x20, 0xc2, 0x52, 0xcc, 0xa1, 0xcb, 0x6d, 0xa7, 0x63, 0xca, 0xbb, 0x90, 0x30,
	0xb3, 0x2a, 0xde, 0x1c, 0x8b, 0x17, 0x47, 0x38, 0xff, 0xee, 0x26, 0x54, 0x12, 0xfd, 0x3b, 0x52,
	0x84, 0xc5, 0x47, 0xec, 0x45, 0x75, 0x81, 0x54, 0xa0, 0x68, 0x30, 0xd1, 0x80, 0xa8, 0x6a, 0x3b,
	0x5f, 0x17, 0x61, 0x6d, 0xb7, 0xb9, 0x77, 0xb8, 0xeb, 0x79, 0x3d, 0xbb, 0x2d, 0xf2, 0x19, 0xf9,
	0x04, 0xf2, 0xa2, 0x08, 0xcf, 0xf0, 0x61, 0xab, 0x9e, 0xa5, 0x9b, 0x45, 0x0c, 0x58, 0x12, 0xb5,
	0x3a, 0xc9, 0xf2, 0xbd, 0xab, 0x9e, 0xa9, 0xc9, 0x85, 0x9b, 0x14, 0x06, 0x97, 0xe1, 0x33, 0x58,
	0x3d, 0x4b, 0xe7, 0x8b, 0x7c, 0x06, 0xe5, 0xb8, 0x08, 0xcf, 0xfa, 0x71, 0xac, 0x9e, 0xb9, 0x27,
	0x86, 0xf2, 0xe3, 0xb2, 0x23, 0xeb, 0xa7, 0xa1, 0x7a, 0xe6, 0x66, 0x10, 0x79, 0x06, 0xc5, 0xb0,
	0xc0, 0xcb, 0xf6, 0xf9, 0xaa, 0x9e, 0xb1, 0x5f, 0x85, 0xd7, 0x27, 0xeb, 0xf2, 0x2c, 0xdf, 0xe8,
	0xea, 0x99, 0x9a, 0x72, 0xe4, 0x29, 0x14, 0x14, 0xb0, 0xce, 0xf4, 0x61, 0xaa, 0x9e, 0xad, 0x0b,
	0x85, 0x4a, 0x8e, 0x3b, 0x1f, 0x59, 0xbf, 0x4b, 0xd6, 0x33, 0x77, 0x23, 0x09, 0x05, 0x48, 0x14,
	0xeb, 0x99, 0x3f, 0x38, 0xd6, 0xb3, 0x77, 0x19, 0xc9, 0x8f, 0xa1, 0x14, 0x95, 0x64, 0x19, 0x3f,
	0xfc, 0xd5, 0xb3, 0x36, 0xfa, 0x9a, 0x87, 0xff, 0xfe, 0xeb, 0xba, 0xf6, 0xdb, 0xb3, 0x75, 0xed,
	0xcb, 0xb3, 0x75, 0xed, 0xab, 0xb3, 0x75, 0xed, 0x8f, 0x67, 0xeb, 0xda, 0x5f, 0xce, 0xd6, 0xb5,
	0xdf, 0xff, 0x6d, 0x5d, 0xfb, 0xd1, 0x7b, 0x1d, 0x9b, 0x77, 0x07, 0xad, 0x46, 0xdb, 0xed, 0x6f,
	0xc7, 0x02, 0x93, 0x8f, 0xf1, 0xd7, 0xfc, 0x56, 0x41, 0x04, 0xac, 0x6f, 0xff, 0x27, 0x00, 0x00,
	0xff, 0xff, 0xd6, 0x1b, 0x60, 0x6f, 0xe2, 0x1f, 0x00, 0x00,
}

func (this *Request) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Request)
	if !ok {
		that2, ok := that.(Request)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if that1.Value == nil {
		if this.Value != nil {
			return false
		}
	} else if this.Value == nil {
		return false
	} else if !this.Value.Equal(that1.Value) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Request_Echo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Request_Echo)
	if !ok {
		that2, ok := that.(Request_Echo)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Echo.Equal(that1.Echo) {
		return false
	}
	return true
}
func (this *Request_Flush) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Request_Flush)
	if !ok {
		that2, ok := that.(Request_Flush)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Flush.Equal(that1.Flush) {
		return false
	}
	return true
}
func (this *Request_Info) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Request_Info)
	if !ok {
		that2, ok := that.(Request_Info)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Info.Equal(that1.Info) {
		return false
	}
	return true
}
func (this *Request_SetOption) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Request_SetOption)
	if !ok {
		that2, ok := that.(Request_SetOption)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.SetOption.Equal(that1.SetOption) {
		return false
	}
	return true
}
func (this *Request_InitChain) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Request_InitChain)
	if !ok {
		that2, ok := that.(Request_InitChain)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.InitChain.Equal(that1.InitChain) {
		return false
	}
	return true
}
func (this *Request_Query) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Request_Query)
	if !ok {
		that2, ok := that.(Request_Query)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Query.Equal(that1.Query) {
		return false
	}
	return true
}
func (this *Request_BeginBlock) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Request_BeginBlock)
	if !ok {
		that2, ok := that.(Request_BeginBlock)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.BeginBlock.Equal(that1.BeginBlock) {
		return false
	}
	return true
}
func (this *Request_CheckTx) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Request_CheckTx)
	if !ok {
		that2, ok := that.(Request_CheckTx)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.CheckTx.Equal(that1.CheckTx) {
		return false
	}
	return true
}
func (this *Request_DeliverTx) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Request_DeliverTx)
	if !ok {
		that2, ok := that.(Request_DeliverTx)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.DeliverTx.Equal(that1.DeliverTx) {
		return false
	}
	return true
}
func (this *Request_EndBlock) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Request_EndBlock)
	if !ok {
		that2, ok := that.(Request_EndBlock)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.EndBlock.Equal(that1.EndBlock) {
		return false
	}
	return true
}
func (this *Request_Commit) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Request_Commit)
	if !ok {
		that2, ok := that.(Request_Commit)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Commit.Equal(that1.Commit) {
		return false
	}
	return true
}
func (this *RequestEcho) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RequestEcho)
	if !ok {
		that2, ok := that.(RequestEcho)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Message != that1.Message {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RequestFlush) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RequestFlush)
	if !ok {
		that2, ok := that.(RequestFlush)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RequestInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RequestInfo)
	if !ok {
		that2, ok := that.(RequestInfo)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Version != that1.Version {
		return false
	}
	if this.BlockVersion != that1.BlockVersion {
		return false
	}
	if this.P2PVersion != that1.P2PVersion {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RequestSetOption) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RequestSetOption)
	if !ok {
		that2, ok := that.(RequestSetOption)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Key != that1.Key {
		return false
	}
	if this.Value != that1.Value {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RequestInitChain) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RequestInitChain)
	if !ok {
		that2, ok := that.(RequestInitChain)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Time.Equal(that1.Time) {
		return false
	}
	if this.ChainId != that1.ChainId {
		return false
	}
	if !this.ConsensusParams.Equal(that1.ConsensusParams) {
		return false
	}
	if len(this.Validators) != len(that1.Validators) {
		return false
	}
	for i := range this.Validators {
		if !this.Validators[i].Equal(&that1.Validators[i]) {
			return false
		}
	}
	if !bytes.Equal(this.AppStateBytes, that1.AppStateBytes) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RequestQuery) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RequestQuery)
	if !ok {
		that2, ok := that.(RequestQuery)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.Data, that1.Data) {
		return false
	}
	if this.Path != that1.Path {
		return false
	}
	if this.Height != that1.Height {
		return false
	}
	if this.Prove != that1.Prove {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RequestBeginBlock) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RequestBeginBlock)
	if !ok {
		that2, ok := that.(RequestBeginBlock)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.Hash, that1.Hash) {
		return false
	}
	if !this.Header.Equal(&that1.Header) {
		return false
	}
	if !this.LastCommitInfo.Equal(&that1.LastCommitInfo) {
		return false
	}
	if len(this.ByzantineValidators) != len(that1.ByzantineValidators) {
		return false
	}
	for i := range this.ByzantineValidators {
		if !this.ByzantineValidators[i].Equal(&that1.ByzantineValidators[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RequestCheckTx) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RequestCheckTx)
	if !ok {
		that2, ok := that.(RequestCheckTx)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.Tx, that1.Tx) {
		return false
	}
	if this.Type != that1.Type {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RequestDeliverTx) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RequestDeliverTx)
	if !ok {
		that2, ok := that.(RequestDeliverTx)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.Tx, that1.Tx) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RequestEndBlock) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RequestEndBlock)
	if !ok {
		that2, ok := that.(RequestEndBlock)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Height != that1.Height {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RequestCommit) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RequestCommit)
	if !ok {
		that2, ok := that.(RequestCommit)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Response) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Response)
	if !ok {
		that2, ok := that.(Response)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if that1.Value == nil {
		if this.Value != nil {
			return false
		}
	} else if this.Value == nil {
		return false
	} else if !this.Value.Equal(that1.Value) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Response_Exception) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Response_Exception)
	if !ok {
		that2, ok := that.(Response_Exception)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Exception.Equal(that1.Exception) {
		return false
	}
	return true
}
func (this *Response_Echo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Response_Echo)
	if !ok {
		that2, ok := that.(Response_Echo)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Echo.Equal(that1.Echo) {
		return false
	}
	return true
}
func (this *Response_Flush) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Response_Flush)
	if !ok {
		that2, ok := that.(Response_Flush)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Flush.Equal(that1.Flush) {
		return false
	}
	return true
}
func (this *Response_Info) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Response_Info)
	if !ok {
		that2, ok := that.(Response_Info)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Info.Equal(that1.Info) {
		return false
	}
	return true
}
func (this *Response_SetOption) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Response_SetOption)
	if !ok {
		that2, ok := that.(Response_SetOption)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.SetOption.Equal(that1.SetOption) {
		return false
	}
	return true
}
func (this *Response_InitChain) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Response_InitChain)
	if !ok {
		that2, ok := that.(Response_InitChain)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.InitChain.Equal(that1.InitChain) {
		return false
	}
	return true
}
func (this *Response_Query) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Response_Query)
	if !ok {
		that2, ok := that.(Response_Query)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Query.Equal(that1.Query) {
		return false
	}
	return true
}
func (this *Response_BeginBlock) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Response_BeginBlock)
	if !ok {
		that2, ok := that.(Response_BeginBlock)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.BeginBlock.Equal(that1.BeginBlock) {
		return false
	}
	return true
}
func (this *Response_CheckTx) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Response_CheckTx)
	if !ok {
		that2, ok := that.(Response_CheckTx)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.CheckTx.Equal(that1.CheckTx) {
		return false
	}
	return true
}
func (this *Response_DeliverTx) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Response_DeliverTx)
	if !ok {
		that2, ok := that.(Response_DeliverTx)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.DeliverTx.Equal(that1.DeliverTx) {
		return false
	}
	return true
}
func (this *Response_EndBlock) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Response_EndBlock)
	if !ok {
		that2, ok := that.(Response_EndBlock)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.EndBlock.Equal(that1.EndBlock) {
		return false
	}
	return true
}
func (this *Response_Commit) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Response_Commit)
	if !ok {
		that2, ok := that.(Response_Commit)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Commit.Equal(that1.Commit) {
		return false
	}
	return true
}
func (this *ResponseException) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ResponseException)
	if !ok {
		that2, ok := that.(ResponseException)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Error != that1.Error {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ResponseEcho) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ResponseEcho)
	if !ok {
		that2, ok := that.(ResponseEcho)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Message != that1.Message {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ResponseFlush) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ResponseFlush)
	if !ok {
		that2, ok := that.(ResponseFlush)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ResponseInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ResponseInfo)
	if !ok {
		that2, ok := that.(ResponseInfo)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Data != that1.Data {
		return false
	}
	if this.Version != that1.Version {
		return false
	}
	if this.AppVersion != that1.AppVersion {
		return false
	}
	if this.LastBlockHeight != that1.LastBlockHeight {
		return false
	}
	if !bytes.Equal(this.LastBlockAppHash, that1.LastBlockAppHash) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ResponseSetOption) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ResponseSetOption)
	if !ok {
		that2, ok := that.(ResponseSetOption)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Code != that1.Code {
		return false
	}
	if this.Log != that1.Log {
		return false
	}
	if this.Info != that1.Info {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ResponseInitChain) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ResponseInitChain)
	if !ok {
		that2, ok := that.(ResponseInitChain)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.ConsensusParams.Equal(that1.ConsensusParams) {
		return false
	}
	if len(this.Validators) != len(that1.Validators) {
		return false
	}
	for i := range this.Validators {
		if !this.Validators[i].Equal(&that1.Validators[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ResponseQuery) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ResponseQuery)
	if !ok {
		that2, ok := that.(ResponseQuery)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Code != that1.Code {
		return false
	}
	if this.Log != that1.Log {
		return false
	}
	if this.Info != that1.Info {
		return false
	}
	if this.Index != that1.Index {
		return false
	}
	if !bytes.Equal(this.Key, that1.Key) {
		return false
	}
	if !bytes.Equal(this.Value, that1.Value) {
		return false
	}
	if !this.Proof.Equal(that1.Proof) {
		return false
	}
	if this.Height != that1.Height {
		return false
	}
	if this.Codespace != that1.Codespace {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ResponseBeginBlock) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ResponseBeginBlock)
	if !ok {
		that2, ok := that.(ResponseBeginBlock)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Events) != len(that1.Events) {
		return false
	}
	for i := range this.Events {
		if !this.Events[i].Equal(&that1.Events[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ResponseCheckTx) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ResponseCheckTx)
	if !ok {
		that2, ok := that.(ResponseCheckTx)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Code != that1.Code {
		return false
	}
	if !bytes.Equal(this.Data, that1.Data) {
		return false
	}
	if this.Log != that1.Log {
		return false
	}
	if this.Info != that1.Info {
		return false
	}
	if this.GasWanted != that1.GasWanted {
		return false
	}
	if this.GasUsed != that1.GasUsed {
		return false
	}
	if len(this.Events) != len(that1.Events) {
		return false
	}
	for i := range this.Events {
		if !this.Events[i].Equal(&that1.Events[i]) {
			return false
		}
	}
	if this.Codespace != that1.Codespace {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ResponseDeliverTx) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ResponseDeliverTx)
	if !ok {
		that2, ok := that.(ResponseDeliverTx)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Code != that1.Code {
		return false
	}
	if !bytes.Equal(this.Data, that1.Data) {
		return false
	}
	if this.Log != that1.Log {
		return false
	}
	if this.Info != that1.Info {
		return false
	}
	if this.GasWanted != that1.GasWanted {
		return false
	}
	if this.GasUsed != that1.GasUsed {
		return false
	}
	if len(this.Events) != len(that1.Events) {
		return false
	}
	for i := range this.Events {
		if !this.Events[i].Equal(&that1.Events[i]) {
			return false
		}
	}
	if this.Codespace != that1.Codespace {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ResponseEndBlock) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ResponseEndBlock)
	if !ok {
		that2, ok := that.(ResponseEndBlock)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.ValidatorUpdates) != len(that1.ValidatorUpdates) {
		return false
	}
	for i := range this.ValidatorUpdates {
		if !this.ValidatorUpdates[i].Equal(&that1.ValidatorUpdates[i]) {
			return false
		}
	}
	if !this.ConsensusParamUpdates.Equal(that1.ConsensusParamUpdates) {
		return false
	}
	if len(this.Events) != len(that1.Events) {
		return false
	}
	for i := range this.Events {
		if !this.Events[i].Equal(&that1.Events[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ResponseCommit) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ResponseCommit)
	if !ok {
		that2, ok := that.(ResponseCommit)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.Data, that1.Data) {
		return false
	}
	if this.RetainHeight != that1.RetainHeight {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ConsensusParams) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ConsensusParams)
	if !ok {
		that2, ok := that.(ConsensusParams)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Block.Equal(that1.Block) {
		return false
	}
	if !this.Evidence.Equal(that1.Evidence) {
		return false
	}
	if !this.Validator.Equal(that1.Validator) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *BlockParams) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*BlockParams)
	if !ok {
		that2, ok := that.(BlockParams)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.MaxBytes != that1.MaxBytes {
		return false
	}
	if this.MaxGas != that1.MaxGas {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *EvidenceParams) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*EvidenceParams)
	if !ok {
		that2, ok := that.(EvidenceParams)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.MaxAgeNumBlocks != that1.MaxAgeNumBlocks {
		return false
	}
	if this.MaxAgeDuration != that1.MaxAgeDuration {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ValidatorParams) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ValidatorParams)
	if !ok {
		that2, ok := that.(ValidatorParams)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.PubKeyTypes) != len(that1.PubKeyTypes) {
		return false
	}
	for i := range this.PubKeyTypes {
		if this.PubKeyTypes[i] != that1.PubKeyTypes[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *LastCommitInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LastCommitInfo)
	if !ok {
		that2, ok := that.(LastCommitInfo)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Round != that1.Round {
		return false
	}
	if len(this.Votes) != len(that1.Votes) {
		return false
	}
	for i := range this.Votes {
		if !this.Votes[i].Equal(&that1.Votes[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Event) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Event)
	if !ok {
		that2, ok := that.(Event)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Type != that1.Type {
		return false
	}
	if len(this.Attributes) != len(that1.Attributes) {
		return false
	}
	for i := range this.Attributes {
		if !this.Attributes[i].Equal(&that1.Attributes[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Header) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Header)
	if !ok {
		that2, ok := that.(Header)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Version.Equal(&that1.Version) {
		return false
	}
	if this.ChainID != that1.ChainID {
		return false
	}
	if this.Height != that1.Height {
		return false
	}
	if !this.Time.Equal(that1.Time) {
		return false
	}
	if !this.LastBlockId.Equal(&that1.LastBlockId) {
		return false
	}
	if !bytes.Equal(this.LastCommitHash, that1.LastCommitHash) {
		return false
	}
	if !bytes.Equal(this.DataHash, that1.DataHash) {
		return false
	}
	if !bytes.Equal(this.ValidatorsHash, that1.ValidatorsHash) {
		return false
	}
	if !bytes.Equal(this.NextValidatorsHash, that1.NextValidatorsHash) {
		return false
	}
	if !bytes.Equal(this.ConsensusHash, that1.ConsensusHash) {
		return false
	}
	if !bytes.Equal(this.AppHash, that1.AppHash) {
		return false
	}
	if !bytes.Equal(this.LastResultsHash, that1.LastResultsHash) {
		return false
	}
	if !bytes.Equal(this.EvidenceHash, that1.EvidenceHash) {
		return false
	}
	if !bytes.Equal(this.ProposerAddress, that1.ProposerAddress) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Version) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Version)
	if !ok {
		that2, ok := that.(Version)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Block != that1.Block {
		return false
	}
	if this.App != that1.App {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *BlockID) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*BlockID)
	if !ok {
		that2, ok := that.(BlockID)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.Hash, that1.Hash) {
		return false
	}
	if !this.PartsHeader.Equal(&that1.PartsHeader) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *PartSetHeader) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PartSetHeader)
	if !ok {
		that2, ok := that.(PartSetHeader)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Total != that1.Total {
		return false
	}
	if !bytes.Equal(this.Hash, that1.Hash) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Validator) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Validator)
	if !ok {
		that2, ok := that.(Validator)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.Address, that1.Address) {
		return false
	}
	if this.Power != that1.Power {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ValidatorUpdate) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ValidatorUpdate)
	if !ok {
		that2, ok := that.(ValidatorUpdate)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.PubKey.Equal(&that1.PubKey) {
		return false
	}
	if this.Power != that1.Power {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *VoteInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VoteInfo)
	if !ok {
		that2, ok := that.(VoteInfo)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Validator.Equal(&that1.Validator) {
		return false
	}
	if this.SignedLastBlock != that1.SignedLastBlock {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *PubKey) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PubKey)
	if !ok {
		that2, ok := that.(PubKey)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Type != that1.Type {
		return false
	}
	if !bytes.Equal(this.Data, that1.Data) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Evidence) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Evidence)
	if !ok {
		that2, ok := that.(Evidence)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Type != that1.Type {
		return false
	}
	if !this.Validator.Equal(&that1.Validator) {
		return false
	}
	if this.Height != that1.Height {
		return false
	}
	if !this.Time.Equal(that1.Time) {
		return false
	}
	if this.TotalVotingPower != that1.TotalVotingPower {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ABCIApplicationClient is the client API for ABCIApplication service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ABCIApplicationClient interface {
	Echo(ctx context.Context, in *RequestEcho, opts ...grpc.CallOption) (*ResponseEcho, error)
	Flush(ctx context.Context, in *RequestFlush, opts ...grpc.CallOption) (*ResponseFlush, error)
	Info(ctx context.Context, in *RequestInfo, opts ...grpc.CallOption) (*ResponseInfo, error)
	SetOption(ctx context.Context, in *RequestSetOption, opts ...grpc.CallOption) (*ResponseSetOption, error)
	DeliverTx(ctx context.Context, in *RequestDeliverTx, opts ...grpc.CallOption) (*ResponseDeliverTx, error)
	CheckTx(ctx context.Context, in *RequestCheckTx, opts ...grpc.CallOption) (*ResponseCheckTx, error)
	Query(ctx context.Context, in *RequestQuery, opts ...grpc.CallOption) (*ResponseQuery, error)
	Commit(ctx context.Context, in *RequestCommit, opts ...grpc.CallOption) (*ResponseCommit, error)
	InitChain(ctx context.Context, in *RequestInitChain, opts ...grpc.CallOption) (*ResponseInitChain, error)
	BeginBlock(ctx context.Context, in *RequestBeginBlock, opts ...grpc.CallOption) (*ResponseBeginBlock, error)
	EndBlock(ctx context.Context, in *RequestEndBlock, opts ...grpc.CallOption) (*ResponseEndBlock, error)
}

type aBCIApplicationClient struct {
	cc *grpc.ClientConn
}

func NewABCIApplicationClient(cc *grpc.ClientConn) ABCIApplicationClient {
	return &aBCIApplicationClient{cc}
}

func (c *aBCIApplicationClient) Echo(ctx context.Context, in *RequestEcho, opts ...grpc.CallOption) (*ResponseEcho, error) {
	out := new(ResponseEcho)
	err := c.cc.Invoke(ctx, "/tendermint.abci.types.ABCIApplication/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCIApplicationClient) Flush(ctx context.Context, in *RequestFlush, opts ...grpc.CallOption) (*ResponseFlush, error) {
	out := new(ResponseFlush)
	err := c.cc.Invoke(ctx, "/tendermint.abci.types.ABCIApplication/Flush", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCIApplicationClient) Info(ctx context.Context, in *RequestInfo, opts ...grpc.CallOption) (*ResponseInfo, error) {
	out := new(ResponseInfo)
	err := c.cc.Invoke(ctx, "/tendermint.abci.types.ABCIApplication/Info", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCIApplicationClient) SetOption(ctx context.Context, in *RequestSetOption, opts ...grpc.CallOption) (*ResponseSetOption, error) {
	out := new(ResponseSetOption)
	err := c.cc.Invoke(ctx, "/tendermint.abci.types.ABCIApplication/SetOption", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCIApplicationClient) DeliverTx(ctx context.Context, in *RequestDeliverTx, opts ...grpc.CallOption) (*ResponseDeliverTx, error) {
	out := new(ResponseDeliverTx)
	err := c.cc.Invoke(ctx, "/tendermint.abci.types.ABCIApplication/DeliverTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCIApplicationClient) CheckTx(ctx context.Context, in *RequestCheckTx, opts ...grpc.CallOption) (*ResponseCheckTx, error) {
	out := new(ResponseCheckTx)
	err := c.cc.Invoke(ctx, "/tendermint.abci.types.ABCIApplication/CheckTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCIApplicationClient) Query(ctx context.Context, in *RequestQuery, opts ...grpc.CallOption) (*ResponseQuery, error) {
	out := new(ResponseQuery)
	err := c.cc.Invoke(ctx, "/tendermint.abci.types.ABCIApplication/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCIApplicationClient) Commit(ctx context.Context, in *RequestCommit, opts ...grpc.CallOption) (*ResponseCommit, error) {
	out := new(ResponseCommit)
	err := c.cc.Invoke(ctx, "/tendermint.abci.types.ABCIApplication/Commit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCIApplicationClient) InitChain(ctx context.Context, in *RequestInitChain, opts ...grpc.CallOption) (*ResponseInitChain, error) {
	out := new(ResponseInitChain)
	err := c.cc.Invoke(ctx, "/tendermint.abci.types.ABCIApplication/InitChain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCIApplicationClient) BeginBlock(ctx context.Context, in *RequestBeginBlock, opts ...grpc.CallOption) (*ResponseBeginBlock, error) {
	out := new(ResponseBeginBlock)
	err := c.cc.Invoke(ctx, "/tendermint.abci.types.ABCIApplication/BeginBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCIApplicationClient) EndBlock(ctx context.Context, in *RequestEndBlock, opts ...grpc.CallOption) (*ResponseEndBlock, error) {
	out := new(ResponseEndBlock)
	err := c.cc.Invoke(ctx, "/tendermint.abci.types.ABCIApplication/EndBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ABCIApplicationServer is the server API for ABCIApplication service.
type ABCIApplicationServer interface {
	Echo(context.Context, *RequestEcho) (*ResponseEcho, error)
	Flush(context.Context, *RequestFlush) (*ResponseFlush, error)
	Info(context.Context, *RequestInfo) (*ResponseInfo, error)
	SetOption(context.Context, *RequestSetOption) (*ResponseSetOption, error)
	DeliverTx(context.Context, *RequestDeliverTx) (*ResponseDeliverTx, error)
	CheckTx(context.Context, *RequestCheckTx) (*ResponseCheckTx, error)
	Query(context.Context, *RequestQuery) (*ResponseQuery, error)
	Commit(context.Context, *RequestCommit) (*ResponseCommit, error)
	InitChain(context.Context, *RequestInitChain) (*ResponseInitChain, error)
	BeginBlock(context.Context, *RequestBeginBlock) (*ResponseBeginBlock, error)
	EndBlock(context.Context, *RequestEndBlock) (*ResponseEndBlock, error)
}

// UnimplementedABCIApplicationServer can be embedded to have forward compatible implementations.
type UnimplementedABCIApplicationServer struct {
}

func (*UnimplementedABCIApplicationServer) Echo(ctx context.Context, req *RequestEcho) (*ResponseEcho, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (*UnimplementedABCIApplicationServer) Flush(ctx context.Context, req *RequestFlush) (*ResponseFlush, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Flush not implemented")
}
func (*UnimplementedABCIApplicationServer) Info(ctx context.Context, req *RequestInfo) (*ResponseInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}
func (*UnimplementedABCIApplicationServer) SetOption(ctx context.Context, req *RequestSetOption) (*ResponseSetOption, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetOption not implemented")
}
func (*UnimplementedABCIApplicationServer) DeliverTx(ctx context.Context, req *RequestDeliverTx) (*ResponseDeliverTx, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeliverTx not implemented")
}
func (*UnimplementedABCIApplicationServer) CheckTx(ctx context.Context, req *RequestCheckTx) (*ResponseCheckTx, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckTx not implemented")
}
func (*UnimplementedABCIApplicationServer) Query(ctx context.Context, req *RequestQuery) (*ResponseQuery, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (*UnimplementedABCIApplicationServer) Commit(ctx context.Context, req *RequestCommit) (*ResponseCommit, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Commit not implemented")
}
func (*UnimplementedABCIApplicationServer) InitChain(ctx context.Context, req *RequestInitChain) (*ResponseInitChain, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitChain not implemented")
}
func (*UnimplementedABCIApplicationServer) BeginBlock(ctx context.Context, req *RequestBeginBlock) (*ResponseBeginBlock, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BeginBlock not implemented")
}
func (*UnimplementedABCIApplicationServer) EndBlock(ctx context.Context, req *RequestEndBlock) (*ResponseEndBlock, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EndBlock not implemented")
}

func RegisterABCIApplicationServer(s *grpc.Server, srv ABCIApplicationServer) {
	s.RegisterService(&_ABCIApplication_serviceDesc, srv)
}

func _ABCIApplication_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestEcho)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABCIApplicationServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tendermint.abci.types.ABCIApplication/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABCIApplicationServer).Echo(ctx, req.(*RequestEcho))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABCIApplication_Flush_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestFlush)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABCIApplicationServer).Flush(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tendermint.abci.types.ABCIApplication/Flush",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABCIApplicationServer).Flush(ctx, req.(*RequestFlush))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABCIApplication_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABCIApplicationServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tendermint.abci.types.ABCIApplication/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABCIApplicationServer).Info(ctx, req.(*RequestInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABCIApplication_SetOption_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestSetOption)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABCIApplicationServer).SetOption(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tendermint.abci.types.ABCIApplication/SetOption",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABCIApplicationServer).SetOption(ctx, req.(*RequestSetOption))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABCIApplication_DeliverTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestDeliverTx)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABCIApplicationServer).DeliverTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tendermint.abci.types.ABCIApplication/DeliverTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABCIApplicationServer).DeliverTx(ctx, req.(*RequestDeliverTx))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABCIApplication_CheckTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestCheckTx)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABCIApplicationServer).CheckTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tendermint.abci.types.ABCIApplication/CheckTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABCIApplicationServer).CheckTx(ctx, req.(*RequestCheckTx))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABCIApplication_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABCIApplicationServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tendermint.abci.types.ABCIApplication/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABCIApplicationServer).Query(ctx, req.(*RequestQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABCIApplication_Commit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestCommit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABCIApplicationServer).Commit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tendermint.abci.types.ABCIApplication/Commit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABCIApplicationServer).Commit(ctx, req.(*RequestCommit))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABCIApplication_InitChain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestInitChain)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABCIApplicationServer).InitChain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tendermint.abci.types.ABCIApplication/InitChain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABCIApplicationServer).InitChain(ctx, req.(*RequestInitChain))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABCIApplication_BeginBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestBeginBlock)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABCIApplicationServer).BeginBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tendermint.abci.types.ABCIApplication/BeginBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABCIApplicationServer).BeginBlock(ctx, req.(*RequestBeginBlock))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABCIApplication_EndBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestEndBlock)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABCIApplicationServer).EndBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tendermint.abci.types.ABCIApplication/EndBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABCIApplicationServer).EndBlock(ctx, req.(*RequestEndBlock))
	}
	return interceptor(ctx, in, info, handler)
}

var _ABCIApplication_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tendermint.abci.types.ABCIApplication",
	HandlerType: (*ABCIApplicationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _ABCIApplication_Echo_Handler,
		},
		{
			MethodName: "Flush",
			Handler:    _ABCIApplication_Flush_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _ABCIApplication_Info_Handler,
		},
		{
			MethodName: "SetOption",
			Handler:    _ABCIApplication_SetOption_Handler,
		},
		{
			MethodName: "DeliverTx",
			Handler:    _ABCIApplication_DeliverTx_Handler,
		},
		{
			MethodName: "CheckTx",
			Handler:    _ABCIApplication_CheckTx_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _ABCIApplication_Query_Handler,
		},
		{
			MethodName: "Commit",
			Handler:    _ABCIApplication_Commit_Handler,
		},
		{
			MethodName: "InitChain",
			Handler:    _ABCIApplication_InitChain_Handler,
		},
		{
			MethodName: "BeginBlock",
			Handler:    _ABCIApplication_BeginBlock_Handler,
		},
		{
			MethodName: "EndBlock",
			Handler:    _ABCIApplication_EndBlock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "abci/types/types.proto",
}

func (m *Request) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Request) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Request) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Value != nil {
		{
			size := m.Value.Size()
			i -= size
			if _, err := m.Value.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *Request_Echo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Request_Echo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Echo != nil {
		{
			size, err := m.Echo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *Request_Flush) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Request_Flush) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Flush != nil {
		{
			size, err := m.Flush.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *Request_Info) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Request_Info) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Info != nil {
		{
			size, err := m.Info.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	return len(dAtA) - i, nil
}
func (m *Request_SetOption) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Request_SetOption) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.SetOption != nil {
		{
			size, err := m.SetOption.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	return len(dAtA) - i, nil
}
func (m *Request_InitChain) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Request_InitChain) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.InitChain != nil {
		{
			size, err := m.InitChain.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	return len(dAtA) - i, nil
}
func (m *Request_Query) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Request_Query) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Query != nil {
		{
			size, err := m.Query.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	return len(dAtA) - i, nil
}
func (m *Request_BeginBlock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Request_BeginBlock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.BeginBlock != nil {
		{
			size, err := m.BeginBlock.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	return len(dAtA) - i, nil
}
func (m *Request_CheckTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Request_CheckTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.CheckTx != nil {
		{
			size, err := m.CheckTx.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x4a
	}
	return len(dAtA) - i, nil
}
func (m *Request_EndBlock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Request_EndBlock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.EndBlock != nil {
		{
			size, err := m.EndBlock.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x5a
	}
	return len(dAtA) - i, nil
}
func (m *Request_Commit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Request_Commit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Commit != nil {
		{
			size, err := m.Commit.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x62
	}
	return len(dAtA) - i, nil
}
func (m *Request_DeliverTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Request_DeliverTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.DeliverTx != nil {
		{
			size, err := m.DeliverTx.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x9a
	}
	return len(dAtA) - i, nil
}
func (m *RequestEcho) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestEcho) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestEcho) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RequestFlush) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestFlush) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestFlush) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func (m *RequestInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.P2PVersion != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.P2PVersion))
		i--
		dAtA[i] = 0x18
	}
	if m.BlockVersion != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.BlockVersion))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RequestSetOption) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestSetOption) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestSetOption) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RequestInitChain) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestInitChain) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestInitChain) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.AppStateBytes) > 0 {
		i -= len(m.AppStateBytes)
		copy(dAtA[i:], m.AppStateBytes)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.AppStateBytes)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Validators) > 0 {
		for iNdEx := len(m.Validators) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Validators[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.ConsensusParams != nil {
		{
			size, err := m.ConsensusParams.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0x12
	}
	n13, err13 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Time, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Time):])
	if err13 != nil {
		return 0, err13
	}
	i -= n13
	i = encodeVarintTypes(dAtA, i, uint64(n13))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *RequestQuery) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestQuery) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestQuery) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Prove {
		i--
		if m.Prove {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.Height != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Path) > 0 {
		i -= len(m.Path)
		copy(dAtA[i:], m.Path)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Path)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RequestBeginBlock) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestBeginBlock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestBeginBlock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ByzantineValidators) > 0 {
		for iNdEx := len(m.ByzantineValidators) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ByzantineValidators[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	{
		size, err := m.LastCommitInfo.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RequestCheckTx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestCheckTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestCheckTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Type != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Tx) > 0 {
		i -= len(m.Tx)
		copy(dAtA[i:], m.Tx)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Tx)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RequestDeliverTx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestDeliverTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestDeliverTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Tx) > 0 {
		i -= len(m.Tx)
		copy(dAtA[i:], m.Tx)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Tx)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RequestEndBlock) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestEndBlock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestEndBlock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Height != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *RequestCommit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestCommit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestCommit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func (m *Response) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Response) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Response) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Value != nil {
		{
			size := m.Value.Size()
			i -= size
			if _, err := m.Value.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *Response_Exception) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Response_Exception) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Exception != nil {
		{
			size, err := m.Exception.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *Response_Echo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Response_Echo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Echo != nil {
		{
			size, err := m.Echo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *Response_Flush) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Response_Flush) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Flush != nil {
		{
			size, err := m.Flush.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *Response_Info) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Response_Info) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Info != nil {
		{
			size, err := m.Info.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	return len(dAtA) - i, nil
}
func (m *Response_SetOption) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Response_SetOption) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.SetOption != nil {
		{
			size, err := m.SetOption.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	return len(dAtA) - i, nil
}
func (m *Response_InitChain) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Response_InitChain) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.InitChain != nil {
		{
			size, err := m.InitChain.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	return len(dAtA) - i, nil
}
func (m *Response_Query) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Response_Query) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Query != nil {
		{
			size, err := m.Query.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	return len(dAtA) - i, nil
}
func (m *Response_BeginBlock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Response_BeginBlock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.BeginBlock != nil {
		{
			size, err := m.BeginBlock.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	return len(dAtA) - i, nil
}
func (m *Response_CheckTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Response_CheckTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.CheckTx != nil {
		{
			size, err := m.CheckTx.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x4a
	}
	return len(dAtA) - i, nil
}
func (m *Response_DeliverTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Response_DeliverTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.DeliverTx != nil {
		{
			size, err := m.DeliverTx.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x52
	}
	return len(dAtA) - i, nil
}
func (m *Response_EndBlock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Response_EndBlock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.EndBlock != nil {
		{
			size, err := m.EndBlock.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x5a
	}
	return len(dAtA) - i, nil
}
func (m *Response_Commit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Response_Commit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Commit != nil {
		{
			size, err := m.Commit.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x62
	}
	return len(dAtA) - i, nil
}
func (m *ResponseException) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResponseException) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResponseException) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Error) > 0 {
		i -= len(m.Error)
		copy(dAtA[i:], m.Error)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Error)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ResponseEcho) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResponseEcho) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResponseEcho) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ResponseFlush) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResponseFlush) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResponseFlush) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func (m *ResponseInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResponseInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResponseInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.LastBlockAppHash) > 0 {
		i -= len(m.LastBlockAppHash)
		copy(dAtA[i:], m.LastBlockAppHash)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.LastBlockAppHash)))
		i--
		dAtA[i] = 0x2a
	}
	if m.LastBlockHeight != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.LastBlockHeight))
		i--
		dAtA[i] = 0x20
	}
	if m.AppVersion != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.AppVersion))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ResponseSetOption) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResponseSetOption) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResponseSetOption) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Info) > 0 {
		i -= len(m.Info)
		copy(dAtA[i:], m.Info)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Info)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Log) > 0 {
		i -= len(m.Log)
		copy(dAtA[i:], m.Log)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Log)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Code != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Code))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ResponseInitChain) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResponseInitChain) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResponseInitChain) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Validators) > 0 {
		for iNdEx := len(m.Validators) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Validators[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.ConsensusParams != nil {
		{
			size, err := m.ConsensusParams.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ResponseQuery) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResponseQuery) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResponseQuery) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Codespace) > 0 {
		i -= len(m.Codespace)
		copy(dAtA[i:], m.Codespace)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Codespace)))
		i--
		dAtA[i] = 0x52
	}
	if m.Height != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x48
	}
	if m.Proof != nil {
		{
			size, err := m.Proof.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0x32
	}
	if m.Index != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Info) > 0 {
		i -= len(m.Info)
		copy(dAtA[i:], m.Info)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Info)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Log) > 0 {
		i -= len(m.Log)
		copy(dAtA[i:], m.Log)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Log)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Code != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Code))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ResponseBeginBlock) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResponseBeginBlock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResponseBeginBlock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Events) > 0 {
		for iNdEx := len(m.Events) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Events[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ResponseCheckTx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResponseCheckTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResponseCheckTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Codespace) > 0 {
		i -= len(m.Codespace)
		copy(dAtA[i:], m.Codespace)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Codespace)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Events) > 0 {
		for iNdEx := len(m.Events) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Events[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if m.GasUsed != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.GasUsed))
		i--
		dAtA[i] = 0x30
	}
	if m.GasWanted != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.GasWanted))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Info) > 0 {
		i -= len(m.Info)
		copy(dAtA[i:], m.Info)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Info)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Log) > 0 {
		i -= len(m.Log)
		copy(dAtA[i:], m.Log)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Log)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x12
	}
	if m.Code != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Code))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ResponseDeliverTx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResponseDeliverTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResponseDeliverTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Codespace) > 0 {
		i -= len(m.Codespace)
		copy(dAtA[i:], m.Codespace)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Codespace)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Events) > 0 {
		for iNdEx := len(m.Events) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Events[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if m.GasUsed != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.GasUsed))
		i--
		dAtA[i] = 0x30
	}
	if m.GasWanted != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.GasWanted))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Info) > 0 {
		i -= len(m.Info)
		copy(dAtA[i:], m.Info)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Info)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Log) > 0 {
		i -= len(m.Log)
		copy(dAtA[i:], m.Log)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Log)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x12
	}
	if m.Code != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Code))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ResponseEndBlock) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResponseEndBlock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResponseEndBlock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Events) > 0 {
		for iNdEx := len(m.Events) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Events[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.ConsensusParamUpdates != nil {
		{
			size, err := m.ConsensusParamUpdates.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.ValidatorUpdates) > 0 {
		for iNdEx := len(m.ValidatorUpdates) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ValidatorUpdates[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ResponseCommit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResponseCommit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResponseCommit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.RetainHeight != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.RetainHeight))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}

func (m *ConsensusParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConsensusParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConsensusParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Validator != nil {
		{
			size, err := m.Validator.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Evidence != nil {
		{
			size, err := m.Evidence.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Block != nil {
		{
			size, err := m.Block.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BlockParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlockParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.MaxGas != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.MaxGas))
		i--
		dAtA[i] = 0x10
	}
	if m.MaxBytes != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.MaxBytes))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EvidenceParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EvidenceParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EvidenceParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	n36, err36 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.MaxAgeDuration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.MaxAgeDuration):])
	if err36 != nil {
		return 0, err36
	}
	i -= n36
	i = encodeVarintTypes(dAtA, i, uint64(n36))
	i--
	dAtA[i] = 0x12
	if m.MaxAgeNumBlocks != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.MaxAgeNumBlocks))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ValidatorParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.PubKeyTypes) > 0 {
		for iNdEx := len(m.PubKeyTypes) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.PubKeyTypes[iNdEx])
			copy(dAtA[i:], m.PubKeyTypes[iNdEx])
			i = encodeVarintTypes(dAtA, i, uint64(len(m.PubKeyTypes[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *LastCommitInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LastCommitInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LastCommitInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Votes) > 0 {
		for iNdEx := len(m.Votes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Votes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Round != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Round))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Event) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Event) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Event) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Attributes) > 0 {
		for iNdEx := len(m.Attributes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Attributes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Header) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Header) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Header) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ProposerAddress) > 0 {
		i -= len(m.ProposerAddress)
		copy(dAtA[i:], m.ProposerAddress)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ProposerAddress)))
		i--
		dAtA[i] = 0x72
	}
	if len(m.EvidenceHash) > 0 {
		i -= len(m.EvidenceHash)
		copy(dAtA[i:], m.EvidenceHash)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.EvidenceHash)))
		i--
		dAtA[i] = 0x6a
	}
	if len(m.LastResultsHash) > 0 {
		i -= len(m.LastResultsHash)
		copy(dAtA[i:], m.LastResultsHash)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.LastResultsHash)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.AppHash) > 0 {
		i -= len(m.AppHash)
		copy(dAtA[i:], m.AppHash)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.AppHash)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.ConsensusHash) > 0 {
		i -= len(m.ConsensusHash)
		copy(dAtA[i:], m.ConsensusHash)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ConsensusHash)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.NextValidatorsHash) > 0 {
		i -= len(m.NextValidatorsHash)
		copy(dAtA[i:], m.NextValidatorsHash)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.NextValidatorsHash)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.ValidatorsHash) > 0 {
		i -= len(m.ValidatorsHash)
		copy(dAtA[i:], m.ValidatorsHash)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ValidatorsHash)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.DataHash) > 0 {
		i -= len(m.DataHash)
		copy(dAtA[i:], m.DataHash)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.DataHash)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.LastCommitHash) > 0 {
		i -= len(m.LastCommitHash)
		copy(dAtA[i:], m.LastCommitHash)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.LastCommitHash)))
		i--
		dAtA[i] = 0x32
	}
	{
		size, err := m.LastBlockId.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	n38, err38 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Time, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Time):])
	if err38 != nil {
		return 0, err38
	}
	i -= n38
	i = encodeVarintTypes(dAtA, i, uint64(n38))
	i--
	dAtA[i] = 0x22
	if m.Height != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ChainID) > 0 {
		i -= len(m.ChainID)
		copy(dAtA[i:], m.ChainID)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ChainID)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Version.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Version) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Version) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Version) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.App != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.App))
		i--
		dAtA[i] = 0x10
	}
	if m.Block != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Block))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *BlockID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlockID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	{
		size, err := m.PartsHeader.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PartSetHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PartSetHeader) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PartSetHeader) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0x12
	}
	if m.Total != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Total))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Validator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Validator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Validator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Power != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Power))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ValidatorUpdate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorUpdate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorUpdate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Power != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Power))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.PubKey.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *VoteInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VoteInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VoteInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.SignedLastBlock {
		i--
		if m.SignedLastBlock {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.Validator.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *PubKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PubKey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PubKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Evidence) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Evidence) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Evidence) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.TotalVotingPower != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.TotalVotingPower))
		i--
		dAtA[i] = 0x28
	}
	n43, err43 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Time, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Time):])
	if err43 != nil {
		return 0, err43
	}
	i -= n43
	i = encodeVarintTypes(dAtA, i, uint64(n43))
	i--
	dAtA[i] = 0x22
	if m.Height != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x18
	}
	{
		size, err := m.Validator.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedRequest(r randyTypes, easy bool) *Request {
	this := &Request{}
	oneofNumber_Value := []int32{2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 19}[r.Intn(11)]
	switch oneofNumber_Value {
	case 2:
		this.Value = NewPopulatedRequest_Echo(r, easy)
	case 3:
		this.Value = NewPopulatedRequest_Flush(r, easy)
	case 4:
		this.Value = NewPopulatedRequest_Info(r, easy)
	case 5:
		this.Value = NewPopulatedRequest_SetOption(r, easy)
	case 6:
		this.Value = NewPopulatedRequest_InitChain(r, easy)
	case 7:
		this.Value = NewPopulatedRequest_Query(r, easy)
	case 8:
		this.Value = NewPopulatedRequest_BeginBlock(r, easy)
	case 9:
		this.Value = NewPopulatedRequest_CheckTx(r, easy)
	case 11:
		this.Value = NewPopulatedRequest_EndBlock(r, easy)
	case 12:
		this.Value = NewPopulatedRequest_Commit(r, easy)
	case 19:
		this.Value = NewPopulatedRequest_DeliverTx(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTypes(r, 20)
	}
	return this
}

func NewPopulatedRequest_Echo(r randyTypes, easy bool) *Request_Echo {
	this := &Request_Echo{}
	this.Echo = NewPopulatedRequestEcho(r, easy)
	return this
}
func NewPopulatedRequest_Flush(r randyTypes, easy bool) *Request_Flush {
	this := &Request_Flush{}
	this.Flush = NewPopulatedRequestFlush(r, easy)
	return this
}
func NewPopulatedRequest_Info(r randyTypes, easy bool) *Request_Info {
	this := &Request_Info{}
	this.Info = NewPopulatedRequestInfo(r, easy)
	return this
}
func NewPopulatedRequest_SetOption(r randyTypes, easy bool) *Request_SetOption {
	this := &Request_SetOption{}
	this.SetOption = NewPopulatedRequestSetOption(r, easy)
	return this
}
func NewPopulatedRequest_InitChain(r randyTypes, easy bool) *Request_InitChain {
	this := &Request_InitChain{}
	this.InitChain = NewPopulatedRequestInitChain(r, easy)
	return this
}
func NewPopulatedRequest_Query(r randyTypes, easy bool) *Request_Query {
	this := &Request_Query{}
	this.Query = NewPopulatedRequestQuery(r, easy)
	return this
}
func NewPopulatedRequest_BeginBlock(r randyTypes, easy bool) *Request_BeginBlock {
	this := &Request_BeginBlock{}
	this.BeginBlock = NewPopulatedRequestBeginBlock(r, easy)
	return this
}
func NewPopulatedRequest_CheckTx(r randyTypes, easy bool) *Request_CheckTx {
	this := &Request_CheckTx{}
	this.CheckTx = NewPopulatedRequestCheckTx(r, easy)
	return this
}
func NewPopulatedRequest_EndBlock(r randyTypes, easy bool) *Request_EndBlock {
	this := &Request_EndBlock{}
	this.EndBlock = NewPopulatedRequestEndBlock(r, easy)
	return this
}
func NewPopulatedRequest_Commit(r randyTypes, easy bool) *Request_Commit {
	this := &Request_Commit{}
	this.Commit = NewPopulatedRequestCommit(r, easy)
	return this
}
func NewPopulatedRequest_DeliverTx(r randyTypes, easy bool) *Request_DeliverTx {
	this := &Request_DeliverTx{}
	this.DeliverTx = NewPopulatedRequestDeliverTx(r, easy)
	return this
}
func NewPopulatedRequestEcho(r randyTypes, easy bool) *RequestEcho {
	this := &RequestEcho{}
	this.Message = string(randStringTypes(r))
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTypes(r, 2)
	}
	return this
}

func NewPopulatedRequestFlush(r randyTypes, easy bool) *RequestFlush {
	this := &RequestFlush{}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTypes(r, 1)
	}
	return this
}

func NewPopulatedRequestInfo(r randyTypes, easy bool) *RequestInfo {
	this := &RequestInfo{}
	this.Version = string(randStringTypes(r))
	this.BlockVersion = uint64(uint64(r.Uint32()))
	this.P2PVersion = uint64(uint64(r.Uint32()))
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTypes(r, 4)
	}
	return this
}

func NewPopulatedRequestSetOption(r randyTypes, easy bool) *RequestSetOption {
	this := &RequestSetOption{}
	this.Key = string(randStringTypes(r))
	this.Value = string(randStringTypes(r))
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTypes(r, 3)
	}
	return this
}

func NewPopulatedRequestInitChain(r randyTypes, easy bool) *RequestInitChain {
	this := &RequestInitChain{}
	v1 := github_com_gogo_protobuf_types.NewPopulatedStdTime(r, easy)
	this.Time = *v1
	this.ChainId = string(randStringTypes(r))
	if r.Intn(5) != 0 {
		this.ConsensusParams = NewPopulatedConsensusParams(r, easy)
	}
	if r.Intn(5) != 0 {
		v2 := r.Intn(5)
		this.Validators = make([]ValidatorUpdate, v2)
		for i := 0; i < v2; i++ {
			v3 := NewPopulatedValidatorUpdate(r, easy)
			this.Validators[i] = *v3
		}
	}
	v4 := r.Intn(100)
	this.AppStateBytes = make([]byte, v4)
	for i := 0; i < v4; i++ {
		this.AppStateBytes[i] = byte(r.Intn(256))
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTypes(r, 6)
	}
	return this
}

func NewPopulatedRequestQuery(r randyTypes, easy bool) *RequestQuery {
	this := &RequestQuery{}
	v5 := r.Intn(100)
	this.Data = make([]byte, v5)
	for i := 0; i < v5; i++ {
		this.Data[i] = byte(r.Intn(256))
	}
	this.Path = string(randStringTypes(r))
	this.Height = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Height *= -1
	}
	this.Prove = bool(bool(r.Intn(2) == 0))
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTypes(r, 5)
	}
	return this
}

func NewPopulatedRequestBeginBlock(r randyTypes, easy bool) *RequestBeginBlock {
	this := &RequestBeginBlock{}
	v6 := r.Intn(100)
	this.Hash = make([]byte, v6)
	for i := 0; i < v6; i++ {
		this.Hash[i] = byte(r.Intn(256))
	}
	v7 := NewPopulatedHeader(r, easy)
	this.Header = *v7
	v8 := NewPopulatedLastCommitInfo(r, easy)
	this.LastCommitInfo = *v8
	if r.Intn(5) != 0 {
		v9 := r.Intn(5)
		this.ByzantineValidators = make([]Evidence, v9)
		for i := 0; i < v9; i++ {
			v10 := NewPopulatedEvidence(r, easy)
			this.ByzantineValidators[i] = *v10
		}
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTypes(r, 6)
	}
	return this
}

func NewPopulatedRequestCheckTx(r randyTypes, easy bool) *RequestCheckTx {
	this := &RequestCheckTx{}
	v11 := r.Intn(100)
	this.Tx = make([]byte, v11)
	for i := 0; i < v11; i++ {
		this.Tx[i] = byte(r.Intn(256))
	}
	this.Type = CheckTxType([]int32{0, 1}[r.Intn(2)])
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTypes(r, 3)
	}
	return this
}

func NewPopulatedRequestDeliverTx(r randyTypes, easy bool) *RequestDeliverTx {
	this := &RequestDeliverTx{}
	v12 := r.Intn(100)
	this.Tx = make([]byte, v12)
	for i := 0; i < v12; i++ {
		this.Tx[i] = byte(r.Intn(256))
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTypes(r, 2)
	}
	return this
}

func NewPopulatedRequestEndBlock(r randyTypes, easy bool) *RequestEndBlock {
	this := &RequestEndBlock{}
	this.Height = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Height *= -1
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTypes(r, 2)
	}
	return this
}

func NewPopulatedRequestCommit(r randyTypes, easy bool) *RequestCommit {
	this := &RequestCommit{}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTypes(r, 2)
	}
	return this
}

func NewPopulatedResponse(r randyTypes, easy bool) *Response {
	this := &Response{}
	oneofNumber_Value := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}[r.Intn(12)]
	switch oneofNumber_Value {
	case 1:
		this.Value = NewPopulatedResponse_Exception(r, easy)
	case 2:
		this.Value = NewPopulatedResponse_Echo(r, easy)
	case 3:
		this.Value = NewPopulatedResponse_Flush(r, easy)
	case 4:
		this.Value = NewPopulatedResponse_Info(r, easy)
	case 5:
		this.Value = NewPopulatedResponse_SetOption(r, easy)
	case 6:
		this.Value = NewPopulatedResponse_InitChain(r, easy)
	case 7:
		this.Value = NewPopulatedResponse_Query(r, easy)
	case 8:
		this.Value = NewPopulatedResponse_BeginBlock(r, easy)
	case 9:
		this.Value = NewPopulatedResponse_CheckTx(r, easy)
	case 10:
		this.Value = NewPopulatedResponse_DeliverTx(r, easy)
	case 11:
		this.Value = NewPopulatedResponse_EndBlock(r, easy)
	case 12:
		this.Value = NewPopulatedResponse_Commit(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTypes(r, 13)
	}
	return this
}

func NewPopulatedResponse_Exception(r randyTypes, easy bool) *Response_Exception {
	this := &Response_Exception{}
	this.Exception = NewPopulatedResponseException(r, easy)
	return this
}
func NewPopulatedResponse_Echo(r randyTypes, easy bool) *Response_Echo {
	this := &Response_Echo{}
	this.Echo = NewPopulatedResponseEcho(r, easy)
	return this
}
func NewPopulatedResponse_Flush(r randyTypes, easy bool) *Response_Flush {
	this := &Response_Flush{}
	this.Flush = NewPopulatedResponseFlush(r, easy)
	return this
}
func NewPopulatedResponse_Info(r randyTypes, easy bool) *Response_Info {
	this := &Response_Info{}
	this.Info = NewPopulatedResponseInfo(r, easy)
	return this
}
func NewPopulatedResponse_SetOption(r randyTypes, easy bool) *Response_SetOption {
	this := &Response_SetOption{}
	this.SetOption = NewPopulatedResponseSetOption(r, easy)
	return this
}
func NewPopulatedResponse_InitChain(r randyTypes, easy bool) *Response_InitChain {
	this := &Response_InitChain{}
	this.InitChain = NewPopulatedResponseInitChain(r, easy)
	return this
}
func NewPopulatedResponse_Query(r randyTypes, easy bool) *Response_Query {
	this := &Response_Query{}
	this.Query = NewPopulatedResponseQuery(r, easy)
	return this
}
func NewPopulatedResponse_BeginBlock(r randyTypes, easy bool) *Response_BeginBlock {
	this := &Response_BeginBlock{}
	this.BeginBlock = NewPopulatedResponseBeginBlock(r, easy)
	return this
}
func NewPopulatedResponse_CheckTx(r randyTypes, easy bool) *Response_CheckTx {
	this := &Response_CheckTx{}
	this.CheckTx = NewPopulatedResponseCheckTx(r, easy)
	return this
}
func NewPopulatedResponse_DeliverTx(r randyTypes, easy bool) *Response_DeliverTx {
	this := &Response_DeliverTx{}
	this.DeliverTx = NewPopulatedResponseDeliverTx(r, easy)
	return this
}
func NewPopulatedResponse_EndBlock(r randyTypes, easy bool) *Response_EndBlock {
	this := &Response_EndBlock{}
	this.EndBlock = NewPopulatedResponseEndBlock(r, easy)
	return this
}
func NewPopulatedResponse_Commit(r randyTypes, easy bool) *Response_Commit {
	this := &Response_Commit{}
	this.Commit = NewPopulatedResponseCommit(r, easy)
	return this
}
func NewPopulatedResponseException(r randyTypes, easy bool) *ResponseException {
	this := &ResponseException{}
	this.Error = string(randStringTypes(r))
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTypes(r, 2)
	}
	return this
}

func NewPopulatedResponseEcho(r randyTypes, easy bool) *ResponseEcho {
	this := &ResponseEcho{}
	this.Message = string(randStringTypes(r))
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTypes(r, 2)
	}
	return this
}

func NewPopulatedResponseFlush(r randyTypes, easy bool) *ResponseFlush {
	this := &ResponseFlush{}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTypes(r, 1)
	}
	return this
}

func NewPopulatedResponseInfo(r randyTypes, easy bool) *ResponseInfo {
	this := &ResponseInfo{}
	this.Data = string(randStringTypes(r))
	this.Version = string(randStringTypes(r))
	this.AppVersion = uint64(uint64(r.Uint32()))
	this.LastBlockHeight = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.LastBlockHeight *= -1
	}
	v15 := r.Intn(100)
	this.LastBlockAppHash = make([]byte, v15)
	for i := 0; i < v15; i++ {
		this.LastBlockAppHash[i] = byte(r.Intn(256))
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTypes(r, 6)
	}
	return this
}

func NewPopulatedResponseSetOption(r randyTypes, easy bool) *ResponseSetOption {
	this := &ResponseSetOption{}
	this.Code = uint32(r.Uint32())
	this.Log = string(randStringTypes(r))
	this.Info = string(randStringTypes(r))
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTypes(r, 5)
	}
	return this
}

func NewPopulatedResponseInitChain(r randyTypes, easy bool) *ResponseInitChain {
	this := &ResponseInitChain{}
	if r.Intn(5) != 0 {
		this.ConsensusParams = NewPopulatedConsensusParams(r, easy)
	}
	if r.Intn(5) != 0 {
		v16 := r.Intn(5)
		this.Validators = make([]ValidatorUpdate, v16)
		for i := 0; i < v16; i++ {
			v17 := NewPopulatedValidatorUpdate(r, easy)
			this.Validators[i] = *v17
		}
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTypes(r, 3)
	}
	return this
}

func NewPopulatedResponseQuery(r randyTypes, easy bool) *ResponseQuery {
	this := &ResponseQuery{}
	this.Code = uint32(r.Uint32())
	this.Log = string(randStringTypes(r))
	this.Info = string(randStringTypes(r))
	this.Index = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Index *= -1
	}
	v18 := r.Intn(100)
	this.Key = make([]byte, v18)
	for i := 0; i < v18; i++ {
		this.Key[i] = byte(r.Intn(256))
	}
	v19 := r.Intn(100)
	this.Value = make([]byte, v19)
	for i := 0; i < v19; i++ {
		this.Value[i] = byte(r.Intn(256))
	}
	if r.Intn(5) != 0 {
		this.Proof = merkle.NewPopulatedProof(r, easy)
	}
	this.Height = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Height *= -1
	}
	this.Codespace = string(randStringTypes(r))
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTypes(r, 11)
	}
	return this
}

func NewPopulatedResponseBeginBlock(r randyTypes, easy bool) *ResponseBeginBlock {
	this := &ResponseBeginBlock{}
	if r.Intn(5) != 0 {
		v20 := r.Intn(5)
		this.Events = make([]Event, v20)
		for i := 0; i < v20; i++ {
			v21 := NewPopulatedEvent(r, easy)
			this.Events[i] = *v21
		}
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTypes(r, 2)
	}
	return this
}

func NewPopulatedResponseCheckTx(r randyTypes, easy bool) *ResponseCheckTx {
	this := &ResponseCheckTx{}
	this.Code = uint32(r.Uint32())
	v22 := r.Intn(100)
	this.Data = make([]byte, v22)
	for i := 0; i < v22; i++ {
		this.Data[i] = byte(r.Intn(256))
	}
	this.Log = string(randStringTypes(r))
	this.Info = string(randStringTypes(r))
	this.GasWanted = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.GasWanted *= -1
	}
	this.GasUsed = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.GasUsed *= -1
	}
	if r.Intn(5) != 0 {
		v23 := r.Intn(5)
		this.Events = make([]Event, v23)
		for i := 0; i < v23; i++ {
			v24 := NewPopulatedEvent(r, easy)
			this.Events[i] = *v24
		}
	}
	this.Codespace = string(randStringTypes(r))
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTypes(r, 9)
	}
	return this
}

func NewPopulatedResponseDeliverTx(r randyTypes, easy bool) *ResponseDeliverTx {
	this := &ResponseDeliverTx{}
	this.Code = uint32(r.Uint32())
	v25 := r.Intn(100)
	this.Data = make([]byte, v25)
	for i := 0; i < v25; i++ {
		this.Data[i] = byte(r.Intn(256))
	}
	this.Log = string(randStringTypes(r))
	this.Info = string(randStringTypes(r))
	this.GasWanted = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.GasWanted *= -1
	}
	this.GasUsed = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.GasUsed *= -1
	}
	if r.Intn(5) != 0 {
		v26 := r.Intn(5)
		this.Events = make([]Event, v26)
		for i := 0; i < v26; i++ {
			v27 := NewPopulatedEvent(r, easy)
			this.Events[i] = *v27
		}
	}
	this.Codespace = string(randStringTypes(r))
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTypes(r, 9)
	}
	return this
}

func NewPopulatedResponseEndBlock(r randyTypes, easy bool) *ResponseEndBlock {
	this := &ResponseEndBlock{}
	if r.Intn(5) != 0 {
		v28 := r.Intn(5)
		this.ValidatorUpdates = make([]ValidatorUpdate, v28)
		for i := 0; i < v28; i++ {
			v29 := NewPopulatedValidatorUpdate(r, easy)
			this.ValidatorUpdates[i] = *v29
		}
	}
	if r.Intn(5) != 0 {
		this.ConsensusParamUpdates = NewPopulatedConsensusParams(r, easy)
	}
	if r.Intn(5) != 0 {
		v30 := r.Intn(5)
		this.Events = make([]Event, v30)
		for i := 0; i < v30; i++ {
			v31 := NewPopulatedEvent(r, easy)
			this.Events[i] = *v31
		}
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTypes(r, 4)
	}
	return this
}

func NewPopulatedResponseCommit(r randyTypes, easy bool) *ResponseCommit {
	this := &ResponseCommit{}
	v32 := r.Intn(100)
	this.Data = make([]byte, v32)
	for i := 0; i < v32; i++ {
		this.Data[i] = byte(r.Intn(256))
	}
	this.RetainHeight = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.RetainHeight *= -1
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTypes(r, 5)
	}
	return this
}

func NewPopulatedConsensusParams(r randyTypes, easy bool) *ConsensusParams {
	this := &ConsensusParams{}
	if r.Intn(5) != 0 {
		this.Block = NewPopulatedBlockParams(r, easy)
	}
	if r.Intn(5) != 0 {
		this.Evidence = NewPopulatedEvidenceParams(r, easy)
	}
	if r.Intn(5) != 0 {
		this.Validator = NewPopulatedValidatorParams(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTypes(r, 4)
	}
	return this
}

func NewPopulatedBlockParams(r randyTypes, easy bool) *BlockParams {
	this := &BlockParams{}
	this.MaxBytes = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.MaxBytes *= -1
	}
	this.MaxGas = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.MaxGas *= -1
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTypes(r, 3)
	}
	return this
}

func NewPopulatedEvidenceParams(r randyTypes, easy bool) *EvidenceParams {
	this := &EvidenceParams{}
	this.MaxAgeNumBlocks = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.MaxAgeNumBlocks *= -1
	}
	v33 := github_com_gogo_protobuf_types.NewPopulatedStdDuration(r, easy)
	this.MaxAgeDuration = *v33
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTypes(r, 3)
	}
	return this
}

func NewPopulatedValidatorParams(r randyTypes, easy bool) *ValidatorParams {
	this := &ValidatorParams{}
	v34 := r.Intn(10)
	this.PubKeyTypes = make([]string, v34)
	for i := 0; i < v34; i++ {
		this.PubKeyTypes[i] = string(randStringTypes(r))
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTypes(r, 2)
	}
	return this
}

func NewPopulatedLastCommitInfo(r randyTypes, easy bool) *LastCommitInfo {
	this := &LastCommitInfo{}
	this.Round = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Round *= -1
	}
	if r.Intn(5) != 0 {
		v35 := r.Intn(5)
		this.Votes = make([]VoteInfo, v35)
		for i := 0; i < v35; i++ {
			v36 := NewPopulatedVoteInfo(r, easy)
			this.Votes[i] = *v36
		}
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTypes(r, 3)
	}
	return this
}

func NewPopulatedEvent(r randyTypes, easy bool) *Event {
	this := &Event{}
	this.Type = string(randStringTypes(r))
	if r.Intn(5) != 0 {
		v37 := r.Intn(5)
		this.Attributes = make([]kv.Pair, v37)
		for i := 0; i < v37; i++ {
			v38 := kv.NewPopulatedPair(r, easy)
			this.Attributes[i] = *v38
		}
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTypes(r, 3)
	}
	return this
}

func NewPopulatedHeader(r randyTypes, easy bool) *Header {
	this := &Header{}
	v39 := NewPopulatedVersion(r, easy)
	this.Version = *v39
	this.ChainID = string(randStringTypes(r))
	this.Height = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Height *= -1
	}
	v40 := github_com_gogo_protobuf_types.NewPopulatedStdTime(r, easy)
	this.Time = *v40
	v41 := NewPopulatedBlockID(r, easy)
	this.LastBlockId = *v41
	v42 := r.Intn(100)
	this.LastCommitHash = make([]byte, v42)
	for i := 0; i < v42; i++ {
		this.LastCommitHash[i] = byte(r.Intn(256))
	}
	v43 := r.Intn(100)
	this.DataHash = make([]byte, v43)
	for i := 0; i < v43; i++ {
		this.DataHash[i] = byte(r.Intn(256))
	}
	v44 := r.Intn(100)
	this.ValidatorsHash = make([]byte, v44)
	for i := 0; i < v44; i++ {
		this.ValidatorsHash[i] = byte(r.Intn(256))
	}
	v45 := r.Intn(100)
	this.NextValidatorsHash = make([]byte, v45)
	for i := 0; i < v45; i++ {
		this.NextValidatorsHash[i] = byte(r.Intn(256))
	}
	v46 := r.Intn(100)
	this.ConsensusHash = make([]byte, v46)
	for i := 0; i < v46; i++ {
		this.ConsensusHash[i] = byte(r.Intn(256))
	}
	v47 := r.Intn(100)
	this.AppHash = make([]byte, v47)
	for i := 0; i < v47; i++ {
		this.AppHash[i] = byte(r.Intn(256))
	}
	v48 := r.Intn(100)
	this.LastResultsHash = make([]byte, v48)
	for i := 0; i < v48; i++ {
		this.LastResultsHash[i] = byte(r.Intn(256))
	}
	v49 := r.Intn(100)
	this.EvidenceHash = make([]byte, v49)
	for i := 0; i < v49; i++ {
		this.EvidenceHash[i] = byte(r.Intn(256))
	}
	v50 := r.Intn(100)
	this.ProposerAddress = make([]byte, v50)
	for i := 0; i < v50; i++ {
		this.ProposerAddress[i] = byte(r.Intn(256))
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTypes(r, 15)
	}
	return this
}

func NewPopulatedVersion(r randyTypes, easy bool) *Version {
	this := &Version{}
	this.Block = uint64(uint64(r.Uint32()))
	this.App = uint64(uint64(r.Uint32()))
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTypes(r, 3)
	}
	return this
}

func NewPopulatedBlockID(r randyTypes, easy bool) *BlockID {
	this := &BlockID{}
	v51 := r.Intn(100)
	this.Hash = make([]byte, v51)
	for i := 0; i < v51; i++ {
		this.Hash[i] = byte(r.Intn(256))
	}
	v52 := NewPopulatedPartSetHeader(r, easy)
	this.PartsHeader = *v52
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTypes(r, 3)
	}
	return this
}

func NewPopulatedPartSetHeader(r randyTypes, easy bool) *PartSetHeader {
	this := &PartSetHeader{}
	this.Total = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Total *= -1
	}
	v53 := r.Intn(100)
	this.Hash = make([]byte, v53)
	for i := 0; i < v53; i++ {
		this.Hash[i] = byte(r.Intn(256))
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTypes(r, 3)
	}
	return this
}

func NewPopulatedValidator(r randyTypes, easy bool) *Validator {
	this := &Validator{}
	v54 := r.Intn(100)
	this.Address = make([]byte, v54)
	for i := 0; i < v54; i++ {
		this.Address[i] = byte(r.Intn(256))
	}
	this.Power = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Power *= -1
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTypes(r, 4)
	}
	return this
}

func NewPopulatedValidatorUpdate(r randyTypes, easy bool) *ValidatorUpdate {
	this := &ValidatorUpdate{}
	v55 := NewPopulatedPubKey(r, easy)
	this.PubKey = *v55
	this.Power = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Power *= -1
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTypes(r, 3)
	}
	return this
}

func NewPopulatedVoteInfo(r randyTypes, easy bool) *VoteInfo {
	this := &VoteInfo{}
	v56 := NewPopulatedValidator(r, easy)
	this.Validator = *v56
	this.SignedLastBlock = bool(bool(r.Intn(2) == 0))
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTypes(r, 3)
	}
	return this
}

func NewPopulatedPubKey(r randyTypes, easy bool) *PubKey {
	this := &PubKey{}
	this.Type = string(randStringTypes(r))
	v57 := r.Intn(100)
	this.Data = make([]byte, v57)
	for i := 0; i < v57; i++ {
		this.Data[i] = byte(r.Intn(256))
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTypes(r, 3)
	}
	return this
}

func NewPopulatedEvidence(r randyTypes, easy bool) *Evidence {
	this := &Evidence{}
	this.Type = string(randStringTypes(r))
	v58 := NewPopulatedValidator(r, easy)
	this.Validator = *v58
	this.Height = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Height *= -1
	}
	v59 := github_com_gogo_protobuf_types.NewPopulatedStdTime(r, easy)
	this.Time = *v59
	this.TotalVotingPower = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.TotalVotingPower *= -1
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTypes(r, 6)
	}
	return this
}

type randyTypes interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneTypes(r randyTypes) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringTypes(r randyTypes) string {
	v60 := r.Intn(100)
	tmps := make([]rune, v60)
	for i := 0; i < v60; i++ {
		tmps[i] = randUTF8RuneTypes(r)
	}
	return string(tmps)
}
func randUnrecognizedTypes(r randyTypes, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldTypes(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldTypes(dAtA []byte, r randyTypes, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateTypes(dAtA, uint64(key))
		v61 := r.Int63()
		if r.Intn(2) == 0 {
			v61 *= -1
		}
		dAtA = encodeVarintPopulateTypes(dAtA, uint64(v61))
	case 1:
		dAtA = encodeVarintPopulateTypes(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateTypes(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateTypes(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateTypes(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateTypes(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Request) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value != nil {
		n += m.Value.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Request_Echo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Echo != nil {
		l = m.Echo.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Request_Flush) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Flush != nil {
		l = m.Flush.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Request_Info) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Info != nil {
		l = m.Info.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Request_SetOption) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SetOption != nil {
		l = m.SetOption.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Request_InitChain) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.InitChain != nil {
		l = m.InitChain.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Request_Query) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Query != nil {
		l = m.Query.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Request_BeginBlock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BeginBlock != nil {
		l = m.BeginBlock.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Request_CheckTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CheckTx != nil {
		l = m.CheckTx.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Request_EndBlock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EndBlock != nil {
		l = m.EndBlock.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Request_Commit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Commit != nil {
		l = m.Commit.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Request_DeliverTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DeliverTx != nil {
		l = m.DeliverTx.Size()
		n += 2 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *RequestEcho) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RequestFlush) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RequestInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.BlockVersion != 0 {
		n += 1 + sovTypes(uint64(m.BlockVersion))
	}
	if m.P2PVersion != 0 {
		n += 1 + sovTypes(uint64(m.P2PVersion))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RequestSetOption) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RequestInitChain) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Time)
	n += 1 + l + sovTypes(uint64(l))
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.ConsensusParams != nil {
		l = m.ConsensusParams.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if len(m.Validators) > 0 {
		for _, e := range m.Validators {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	l = len(m.AppStateBytes)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RequestQuery) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Height != 0 {
		n += 1 + sovTypes(uint64(m.Height))
	}
	if m.Prove {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RequestBeginBlock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = m.Header.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = m.LastCommitInfo.Size()
	n += 1 + l + sovTypes(uint64(l))
	if len(m.ByzantineValidators) > 0 {
		for _, e := range m.ByzantineValidators {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RequestCheckTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Tx)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovTypes(uint64(m.Type))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RequestDeliverTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Tx)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RequestEndBlock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovTypes(uint64(m.Height))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RequestCommit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Response) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value != nil {
		n += m.Value.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Response_Exception) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Exception != nil {
		l = m.Exception.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Response_Echo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Echo != nil {
		l = m.Echo.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Response_Flush) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Flush != nil {
		l = m.Flush.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Response_Info) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Info != nil {
		l = m.Info.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Response_SetOption) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SetOption != nil {
		l = m.SetOption.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Response_InitChain) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.InitChain != nil {
		l = m.InitChain.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Response_Query) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Query != nil {
		l = m.Query.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Response_BeginBlock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BeginBlock != nil {
		l = m.BeginBlock.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Response_CheckTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CheckTx != nil {
		l = m.CheckTx.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Response_DeliverTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DeliverTx != nil {
		l = m.DeliverTx.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Response_EndBlock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EndBlock != nil {
		l = m.EndBlock.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Response_Commit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Commit != nil {
		l = m.Commit.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *ResponseException) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Error)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ResponseEcho) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ResponseFlush) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ResponseInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.AppVersion != 0 {
		n += 1 + sovTypes(uint64(m.AppVersion))
	}
	if m.LastBlockHeight != 0 {
		n += 1 + sovTypes(uint64(m.LastBlockHeight))
	}
	l = len(m.LastBlockAppHash)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ResponseSetOption) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovTypes(uint64(m.Code))
	}
	l = len(m.Log)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Info)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ResponseInitChain) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ConsensusParams != nil {
		l = m.ConsensusParams.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if len(m.Validators) > 0 {
		for _, e := range m.Validators {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ResponseQuery) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovTypes(uint64(m.Code))
	}
	l = len(m.Log)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Info)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Index != 0 {
		n += 1 + sovTypes(uint64(m.Index))
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Proof != nil {
		l = m.Proof.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Height != 0 {
		n += 1 + sovTypes(uint64(m.Height))
	}
	l = len(m.Codespace)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ResponseBeginBlock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Events) > 0 {
		for _, e := range m.Events {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ResponseCheckTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovTypes(uint64(m.Code))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Log)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Info)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.GasWanted != 0 {
		n += 1 + sovTypes(uint64(m.GasWanted))
	}
	if m.GasUsed != 0 {
		n += 1 + sovTypes(uint64(m.GasUsed))
	}
	if len(m.Events) > 0 {
		for _, e := range m.Events {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	l = len(m.Codespace)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ResponseDeliverTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovTypes(uint64(m.Code))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Log)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Info)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.GasWanted != 0 {
		n += 1 + sovTypes(uint64(m.GasWanted))
	}
	if m.GasUsed != 0 {
		n += 1 + sovTypes(uint64(m.GasUsed))
	}
	if len(m.Events) > 0 {
		for _, e := range m.Events {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	l = len(m.Codespace)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ResponseEndBlock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ValidatorUpdates) > 0 {
		for _, e := range m.ValidatorUpdates {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if m.ConsensusParamUpdates != nil {
		l = m.ConsensusParamUpdates.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if len(m.Events) > 0 {
		for _, e := range m.Events {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ResponseCommit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.RetainHeight != 0 {
		n += 1 + sovTypes(uint64(m.RetainHeight))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ConsensusParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Block != nil {
		l = m.Block.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Evidence != nil {
		l = m.Evidence.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Validator != nil {
		l = m.Validator.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *BlockParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaxBytes != 0 {
		n += 1 + sovTypes(uint64(m.MaxBytes))
	}
	if m.MaxGas != 0 {
		n += 1 + sovTypes(uint64(m.MaxGas))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *EvidenceParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaxAgeNumBlocks != 0 {
		n += 1 + sovTypes(uint64(m.MaxAgeNumBlocks))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.MaxAgeDuration)
	n += 1 + l + sovTypes(uint64(l))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ValidatorParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.PubKeyTypes) > 0 {
		for _, s := range m.PubKeyTypes {
			l = len(s)
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *LastCommitInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Round != 0 {
		n += 1 + sovTypes(uint64(m.Round))
	}
	if len(m.Votes) > 0 {
		for _, e := range m.Votes {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Event) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if len(m.Attributes) > 0 {
		for _, e := range m.Attributes {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Header) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Version.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = len(m.ChainID)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Height != 0 {
		n += 1 + sovTypes(uint64(m.Height))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Time)
	n += 1 + l + sovTypes(uint64(l))
	l = m.LastBlockId.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = len(m.LastCommitHash)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.DataHash)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.ValidatorsHash)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.NextValidatorsHash)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.ConsensusHash)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.AppHash)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.LastResultsHash)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.EvidenceHash)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.ProposerAddress)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Version) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Block != 0 {
		n += 1 + sovTypes(uint64(m.Block))
	}
	if m.App != 0 {
		n += 1 + sovTypes(uint64(m.App))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *BlockID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = m.PartsHeader.Size()
	n += 1 + l + sovTypes(uint64(l))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PartSetHeader) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Total != 0 {
		n += 1 + sovTypes(uint64(m.Total))
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Validator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Power != 0 {
		n += 1 + sovTypes(uint64(m.Power))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ValidatorUpdate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.PubKey.Size()
	n += 1 + l + sovTypes(uint64(l))
	if m.Power != 0 {
		n += 1 + sovTypes(uint64(m.Power))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *VoteInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Validator.Size()
	n += 1 + l + sovTypes(uint64(l))
	if m.SignedLastBlock {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PubKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Evidence) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = m.Validator.Size()
	n += 1 + l + sovTypes(uint64(l))
	if m.Height != 0 {
		n += 1 + sovTypes(uint64(m.Height))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Time)
	n += 1 + l + sovTypes(uint64(l))
	if m.TotalVotingPower != 0 {
		n += 1 + sovTypes(uint64(m.TotalVotingPower))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Request) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Request: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Echo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &RequestEcho{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &Request_Echo{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Flush", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &RequestFlush{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &Request_Flush{v}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &RequestInfo{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &Request_Info{v}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SetOption", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &RequestSetOption{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &Request_SetOption{v}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitChain", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &RequestInitChain{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &Request_InitChain{v}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Query", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &RequestQuery{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &Request_Query{v}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BeginBlock", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &RequestBeginBlock{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &Request_BeginBlock{v}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CheckTx", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &RequestCheckTx{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &Request_CheckTx{v}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndBlock", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &RequestEndBlock{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &Request_EndBlock{v}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Commit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &RequestCommit{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &Request_Commit{v}
			iNdEx = postIndex
		case 19:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeliverTx", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &RequestDeliverTx{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &Request_DeliverTx{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RequestEcho) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: RequestEcho: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestEcho: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RequestFlush) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: RequestFlush: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestFlush: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RequestInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: RequestInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockVersion", wireType)
			}
			m.BlockVersion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockVersion |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field P2PVersion", wireType)
			}
			m.P2PVersion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.P2PVersion |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RequestSetOption) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: RequestSetOption: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestSetOption: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RequestInitChain) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: RequestInitChain: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestInitChain: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Time, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsensusParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ConsensusParams == nil {
				m.ConsensusParams = &ConsensusParams{}
			}
			if err := m.ConsensusParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validators", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validators = append(m.Validators, ValidatorUpdate{})
			if err := m.Validators[len(m.Validators)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppStateBytes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AppStateBytes = append(m.AppStateBytes[:0], dAtA[iNdEx:postIndex]...)
			if m.AppStateBytes == nil {
				m.AppStateBytes = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RequestQuery) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: RequestQuery: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestQuery: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Prove", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
			m.Prove = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RequestBeginBlock) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: RequestBeginBlock: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestBeginBlock: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = append(m.Hash[:0], dAtA[iNdEx:postIndex]...)
			if m.Hash == nil {
				m.Hash = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastCommitInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LastCommitInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ByzantineValidators", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ByzantineValidators = append(m.ByzantineValidators, Evidence{})
			if err := m.ByzantineValidators[len(m.ByzantineValidators)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RequestCheckTx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: RequestCheckTx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestCheckTx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tx", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tx = append(m.Tx[:0], dAtA[iNdEx:postIndex]...)
			if m.Tx == nil {
				m.Tx = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= CheckTxType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RequestDeliverTx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: RequestDeliverTx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestDeliverTx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tx", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tx = append(m.Tx[:0], dAtA[iNdEx:postIndex]...)
			if m.Tx == nil {
				m.Tx = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RequestEndBlock) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: RequestEndBlock: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestEndBlock: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RequestCommit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: RequestCommit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestCommit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Response) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Response: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Response: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Exception", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ResponseException{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &Response_Exception{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Echo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ResponseEcho{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &Response_Echo{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Flush", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ResponseFlush{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &Response_Flush{v}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ResponseInfo{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &Response_Info{v}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SetOption", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ResponseSetOption{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &Response_SetOption{v}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitChain", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ResponseInitChain{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &Response_InitChain{v}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Query", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ResponseQuery{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &Response_Query{v}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BeginBlock", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ResponseBeginBlock{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &Response_BeginBlock{v}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CheckTx", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ResponseCheckTx{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &Response_CheckTx{v}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeliverTx", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ResponseDeliverTx{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &Response_DeliverTx{v}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndBlock", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ResponseEndBlock{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &Response_EndBlock{v}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Commit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ResponseCommit{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &Response_Commit{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ResponseException) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: ResponseException: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResponseException: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ResponseEcho) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: ResponseEcho: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResponseEcho: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ResponseFlush) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: ResponseFlush: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResponseFlush: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ResponseInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: ResponseInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResponseInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppVersion", wireType)
			}
			m.AppVersion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AppVersion |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastBlockHeight", wireType)
			}
			m.LastBlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastBlockHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastBlockAppHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastBlockAppHash = append(m.LastBlockAppHash[:0], dAtA[iNdEx:postIndex]...)
			if m.LastBlockAppHash == nil {
				m.LastBlockAppHash = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ResponseSetOption) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: ResponseSetOption: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResponseSetOption: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Log", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Log = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Info = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ResponseInitChain) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: ResponseInitChain: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResponseInitChain: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsensusParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ConsensusParams == nil {
				m.ConsensusParams = &ConsensusParams{}
			}
			if err := m.ConsensusParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validators", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validators = append(m.Validators, ValidatorUpdate{})
			if err := m.Validators[len(m.Validators)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ResponseQuery) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: ResponseQuery: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResponseQuery: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Log", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Log = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Info = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proof", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Proof == nil {
				m.Proof = &merkle.Proof{}
			}
			if err := m.Proof.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Codespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Codespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ResponseBeginBlock) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: ResponseBeginBlock: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResponseBeginBlock: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Events", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Events = append(m.Events, Event{})
			if err := m.Events[len(m.Events)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ResponseCheckTx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: ResponseCheckTx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResponseCheckTx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Log", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Log = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Info = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasWanted", wireType)
			}
			m.GasWanted = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasWanted |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasUsed", wireType)
			}
			m.GasUsed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasUsed |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Events", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Events = append(m.Events, Event{})
			if err := m.Events[len(m.Events)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Codespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Codespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ResponseDeliverTx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: ResponseDeliverTx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResponseDeliverTx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Log", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Log = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Info = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasWanted", wireType)
			}
			m.GasWanted = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasWanted |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasUsed", wireType)
			}
			m.GasUsed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasUsed |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Events", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Events = append(m.Events, Event{})
			if err := m.Events[len(m.Events)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Codespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Codespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ResponseEndBlock) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: ResponseEndBlock: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResponseEndBlock: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorUpdates", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorUpdates = append(m.ValidatorUpdates, ValidatorUpdate{})
			if err := m.ValidatorUpdates[len(m.ValidatorUpdates)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsensusParamUpdates", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ConsensusParamUpdates == nil {
				m.ConsensusParamUpdates = &ConsensusParams{}
			}
			if err := m.ConsensusParamUpdates.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Events", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Events = append(m.Events, Event{})
			if err := m.Events[len(m.Events)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ResponseCommit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: ResponseCommit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResponseCommit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RetainHeight", wireType)
			}
			m.RetainHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RetainHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ConsensusParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: ConsensusParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConsensusParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Block == nil {
				m.Block = &BlockParams{}
			}
			if err := m.Block.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Evidence", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Evidence == nil {
				m.Evidence = &EvidenceParams{}
			}
			if err := m.Evidence.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validator", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Validator == nil {
				m.Validator = &ValidatorParams{}
			}
			if err := m.Validator.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BlockParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: BlockParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxBytes", wireType)
			}
			m.MaxBytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxBytes |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxGas", wireType)
			}
			m.MaxGas = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxGas |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *EvidenceParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: EvidenceParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EvidenceParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxAgeNumBlocks", wireType)
			}
			m.MaxAgeNumBlocks = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxAgeNumBlocks |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxAgeDuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.MaxAgeDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ValidatorParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: ValidatorParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKeyTypes", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PubKeyTypes = append(m.PubKeyTypes, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LastCommitInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: LastCommitInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LastCommitInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Round", wireType)
			}
			m.Round = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Round |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Votes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Votes = append(m.Votes, VoteInfo{})
			if err := m.Votes[len(m.Votes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Event) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Event: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Event: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Attributes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Attributes = append(m.Attributes, kv.Pair{})
			if err := m.Attributes[len(m.Attributes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Header) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Header: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Header: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Version.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Time, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastBlockId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LastBlockId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastCommitHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastCommitHash = append(m.LastCommitHash[:0], dAtA[iNdEx:postIndex]...)
			if m.LastCommitHash == nil {
				m.LastCommitHash = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DataHash = append(m.DataHash[:0], dAtA[iNdEx:postIndex]...)
			if m.DataHash == nil {
				m.DataHash = []byte{}
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorsHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorsHash = append(m.ValidatorsHash[:0], dAtA[iNdEx:postIndex]...)
			if m.ValidatorsHash == nil {
				m.ValidatorsHash = []byte{}
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextValidatorsHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NextValidatorsHash = append(m.NextValidatorsHash[:0], dAtA[iNdEx:postIndex]...)
			if m.NextValidatorsHash == nil {
				m.NextValidatorsHash = []byte{}
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsensusHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConsensusHash = append(m.ConsensusHash[:0], dAtA[iNdEx:postIndex]...)
			if m.ConsensusHash == nil {
				m.ConsensusHash = []byte{}
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AppHash = append(m.AppHash[:0], dAtA[iNdEx:postIndex]...)
			if m.AppHash == nil {
				m.AppHash = []byte{}
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastResultsHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastResultsHash = append(m.LastResultsHash[:0], dAtA[iNdEx:postIndex]...)
			if m.LastResultsHash == nil {
				m.LastResultsHash = []byte{}
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EvidenceHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EvidenceHash = append(m.EvidenceHash[:0], dAtA[iNdEx:postIndex]...)
			if m.EvidenceHash == nil {
				m.EvidenceHash = []byte{}
			}
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposerAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProposerAddress = append(m.ProposerAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.ProposerAddress == nil {
				m.ProposerAddress = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Version) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Version: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Version: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
			}
			m.Block = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Block |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field App", wireType)
			}
			m.App = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.App |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BlockID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: BlockID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = append(m.Hash[:0], dAtA[iNdEx:postIndex]...)
			if m.Hash == nil {
				m.Hash = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PartsHeader", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PartsHeader.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PartSetHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: PartSetHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PartSetHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Total", wireType)
			}
			m.Total = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Total |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = append(m.Hash[:0], dAtA[iNdEx:postIndex]...)
			if m.Hash == nil {
				m.Hash = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Validator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Validator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Validator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Power", wireType)
			}
			m.Power = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Power |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ValidatorUpdate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: ValidatorUpdate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorUpdate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PubKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Power", wireType)
			}
			m.Power = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Power |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *VoteInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: VoteInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VoteInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validator", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Validator.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignedLastBlock", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
			m.SignedLastBlock = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PubKey) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: PubKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PubKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Evidence) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Evidence: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Evidence: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validator", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Validator.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Time, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalVotingPower", wireType)
			}
			m.TotalVotingPower = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalVotingPower |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
