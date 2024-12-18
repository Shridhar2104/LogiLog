// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.0
// source: payment.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request to recharge the wallet.
type RechargeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string  `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // The ID of the user.
	Amount float64 `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`             // The amount to recharge.
}

func (x *RechargeRequest) Reset() {
	*x = RechargeRequest{}
	mi := &file_payment_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RechargeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RechargeRequest) ProtoMessage() {}

func (x *RechargeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RechargeRequest.ProtoReflect.Descriptor instead.
func (*RechargeRequest) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{0}
}

func (x *RechargeRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RechargeRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

// Response for wallet recharge.
type RechargeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success    bool    `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`                          // Indicates if the recharge was successful.
	Message    string  `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`                           // Additional message (e.g., "Recharge successful").
	NewBalance float64 `protobuf:"fixed64,3,opt,name=new_balance,json=newBalance,proto3" json:"new_balance,omitempty"` // Updated wallet balance.
}

func (x *RechargeResponse) Reset() {
	*x = RechargeResponse{}
	mi := &file_payment_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RechargeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RechargeResponse) ProtoMessage() {}

func (x *RechargeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RechargeResponse.ProtoReflect.Descriptor instead.
func (*RechargeResponse) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{1}
}

func (x *RechargeResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *RechargeResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *RechargeResponse) GetNewBalance() float64 {
	if x != nil {
		return x.NewBalance
	}
	return 0
}

// Request to deduct balance for shipping.
type DeductionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  string  `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`    // The ID of the user.
	Amount  float64 `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`                // The amount to deduct.
	OrderId string  `protobuf:"bytes,3,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"` // The associated order ID.
}

func (x *DeductionRequest) Reset() {
	*x = DeductionRequest{}
	mi := &file_payment_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeductionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeductionRequest) ProtoMessage() {}

func (x *DeductionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeductionRequest.ProtoReflect.Descriptor instead.
func (*DeductionRequest) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{2}
}

func (x *DeductionRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *DeductionRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *DeductionRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

// Response for balance deduction.
type DeductionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success    bool    `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`                          // Indicates if the deduction was successful.
	Message    string  `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`                           // Additional message (e.g., "Deduction successful").
	NewBalance float64 `protobuf:"fixed64,3,opt,name=new_balance,json=newBalance,proto3" json:"new_balance,omitempty"` // Updated wallet balance.
}

func (x *DeductionResponse) Reset() {
	*x = DeductionResponse{}
	mi := &file_payment_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeductionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeductionResponse) ProtoMessage() {}

func (x *DeductionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeductionResponse.ProtoReflect.Descriptor instead.
func (*DeductionResponse) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{3}
}

func (x *DeductionResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *DeductionResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *DeductionResponse) GetNewBalance() float64 {
	if x != nil {
		return x.NewBalance
	}
	return 0
}

// Request to process COD remittance.
type RemittanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`       // The ID of the user.
	OrderIds []string `protobuf:"bytes,2,rep,name=order_ids,json=orderIds,proto3" json:"order_ids,omitempty"` // List of order IDs eligible for remittance.
}

func (x *RemittanceRequest) Reset() {
	*x = RemittanceRequest{}
	mi := &file_payment_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemittanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemittanceRequest) ProtoMessage() {}

func (x *RemittanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemittanceRequest.ProtoReflect.Descriptor instead.
func (*RemittanceRequest) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{4}
}

func (x *RemittanceRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RemittanceRequest) GetOrderIds() []string {
	if x != nil {
		return x.OrderIds
	}
	return nil
}

// Response for COD remittance processing.
type RemittanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool                `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"` // Indicates if the remittance was successful.
	Message string              `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`  // Additional message (e.g., "Remittance processed").
	Details []*RemittanceDetail `protobuf:"bytes,3,rep,name=details,proto3" json:"details,omitempty"`  // Details of remittance processing.
}

func (x *RemittanceResponse) Reset() {
	*x = RemittanceResponse{}
	mi := &file_payment_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemittanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemittanceResponse) ProtoMessage() {}

func (x *RemittanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemittanceResponse.ProtoReflect.Descriptor instead.
func (*RemittanceResponse) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{5}
}

func (x *RemittanceResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *RemittanceResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *RemittanceResponse) GetDetails() []*RemittanceDetail {
	if x != nil {
		return x.Details
	}
	return nil
}

// Details for each remittance processed.
type RemittanceDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId   string  `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"` // The associated order ID.
	Amount    float64 `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`                // Amount remitted for this order.
	Processed bool    `protobuf:"varint,3,opt,name=processed,proto3" json:"processed,omitempty"`           // Whether the remittance was successful for this order.
}

func (x *RemittanceDetail) Reset() {
	*x = RemittanceDetail{}
	mi := &file_payment_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemittanceDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemittanceDetail) ProtoMessage() {}

func (x *RemittanceDetail) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemittanceDetail.ProtoReflect.Descriptor instead.
func (*RemittanceDetail) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{6}
}

func (x *RemittanceDetail) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *RemittanceDetail) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *RemittanceDetail) GetProcessed() bool {
	if x != nil {
		return x.Processed
	}
	return false
}

// Request to retrieve wallet details.
type WalletDetailsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // The ID of the user.
}

func (x *WalletDetailsRequest) Reset() {
	*x = WalletDetailsRequest{}
	mi := &file_payment_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WalletDetailsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WalletDetailsRequest) ProtoMessage() {}

func (x *WalletDetailsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WalletDetailsRequest.ProtoReflect.Descriptor instead.
func (*WalletDetailsRequest) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{7}
}

func (x *WalletDetailsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

// Response with wallet details and transaction history.
type WalletDetailsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Balance            float64        `protobuf:"fixed64,1,opt,name=balance,proto3" json:"balance,omitempty"`                                               // Current wallet balance.
	TransactionHistory []*Transaction `protobuf:"bytes,2,rep,name=transaction_history,json=transactionHistory,proto3" json:"transaction_history,omitempty"` // List of past transactions.
}

func (x *WalletDetailsResponse) Reset() {
	*x = WalletDetailsResponse{}
	mi := &file_payment_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WalletDetailsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WalletDetailsResponse) ProtoMessage() {}

func (x *WalletDetailsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WalletDetailsResponse.ProtoReflect.Descriptor instead.
func (*WalletDetailsResponse) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{8}
}

func (x *WalletDetailsResponse) GetBalance() float64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *WalletDetailsResponse) GetTransactionHistory() []*Transaction {
	if x != nil {
		return x.TransactionHistory
	}
	return nil
}

// Transaction history details.
type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionId   string  `protobuf:"bytes,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`       // Unique ID for the transaction.
	TransactionType string  `protobuf:"bytes,2,opt,name=transaction_type,json=transactionType,proto3" json:"transaction_type,omitempty"` // Type of transaction (e.g., "recharge", "deduction", "remittance").
	Amount          float64 `protobuf:"fixed64,3,opt,name=amount,proto3" json:"amount,omitempty"`                                        // Amount of the transaction.
	OrderId         string  `protobuf:"bytes,4,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`                         // Associated order ID, if applicable.
	Timestamp       string  `protobuf:"bytes,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`                                    // Timestamp of the transaction.
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	mi := &file_payment_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{9}
}

func (x *Transaction) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *Transaction) GetTransactionType() string {
	if x != nil {
		return x.TransactionType
	}
	return ""
}

func (x *Transaction) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Transaction) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *Transaction) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

var File_payment_proto protoreflect.FileDescriptor

var file_payment_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x42, 0x0a, 0x0f, 0x52, 0x65, 0x63, 0x68,
	0x61, 0x72, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x67, 0x0a, 0x10,
	0x52, 0x65, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x65, 0x77, 0x5f, 0x62, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x6e, 0x65, 0x77, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x5e, 0x0a, 0x10, 0x44, 0x65, 0x64, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x68, 0x0a, 0x11, 0x44, 0x65, 0x64, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x6e, 0x65, 0x77, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x0a, 0x6e, 0x65, 0x77, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22,
	0x49, 0x0a, 0x11, 0x52, 0x65, 0x6d, 0x69, 0x74, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x73, 0x22, 0x7d, 0x0a, 0x12, 0x52, 0x65,
	0x6d, 0x69, 0x74, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x33, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x52, 0x65, 0x6d, 0x69, 0x74, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x63, 0x0a, 0x10, 0x52, 0x65, 0x6d,
	0x69, 0x74, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x19, 0x0a,
	0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x22, 0x2f,
	0x0a, 0x14, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x78, 0x0a, 0x15, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x12, 0x45, 0x0a, 0x13, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x12, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x22, 0xb0, 0x01, 0x0a, 0x0b, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x29, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x32, 0xc0, 0x02, 0x0a,
	0x0e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x45, 0x0a, 0x0e, 0x52, 0x65, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x12, 0x18, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x63, 0x68,
	0x61, 0x72, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x0d, 0x44, 0x65, 0x64, 0x75, 0x63, 0x74,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x19, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x44, 0x65, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x64,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c,
	0x0a, 0x11, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x6d, 0x69, 0x74, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x12, 0x1a, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65,
	0x6d, 0x69, 0x74, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x6d, 0x69, 0x74, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x12, 0x1d, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x57, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x05, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_payment_proto_rawDescOnce sync.Once
	file_payment_proto_rawDescData = file_payment_proto_rawDesc
)

func file_payment_proto_rawDescGZIP() []byte {
	file_payment_proto_rawDescOnce.Do(func() {
		file_payment_proto_rawDescData = protoimpl.X.CompressGZIP(file_payment_proto_rawDescData)
	})
	return file_payment_proto_rawDescData
}

var file_payment_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_payment_proto_goTypes = []any{
	(*RechargeRequest)(nil),       // 0: payment.RechargeRequest
	(*RechargeResponse)(nil),      // 1: payment.RechargeResponse
	(*DeductionRequest)(nil),      // 2: payment.DeductionRequest
	(*DeductionResponse)(nil),     // 3: payment.DeductionResponse
	(*RemittanceRequest)(nil),     // 4: payment.RemittanceRequest
	(*RemittanceResponse)(nil),    // 5: payment.RemittanceResponse
	(*RemittanceDetail)(nil),      // 6: payment.RemittanceDetail
	(*WalletDetailsRequest)(nil),  // 7: payment.WalletDetailsRequest
	(*WalletDetailsResponse)(nil), // 8: payment.WalletDetailsResponse
	(*Transaction)(nil),           // 9: payment.Transaction
}
var file_payment_proto_depIdxs = []int32{
	6, // 0: payment.RemittanceResponse.details:type_name -> payment.RemittanceDetail
	9, // 1: payment.WalletDetailsResponse.transaction_history:type_name -> payment.Transaction
	0, // 2: payment.PaymentService.RechargeWallet:input_type -> payment.RechargeRequest
	2, // 3: payment.PaymentService.DeductBalance:input_type -> payment.DeductionRequest
	4, // 4: payment.PaymentService.ProcessRemittance:input_type -> payment.RemittanceRequest
	7, // 5: payment.PaymentService.GetWalletDetails:input_type -> payment.WalletDetailsRequest
	1, // 6: payment.PaymentService.RechargeWallet:output_type -> payment.RechargeResponse
	3, // 7: payment.PaymentService.DeductBalance:output_type -> payment.DeductionResponse
	5, // 8: payment.PaymentService.ProcessRemittance:output_type -> payment.RemittanceResponse
	8, // 9: payment.PaymentService.GetWalletDetails:output_type -> payment.WalletDetailsResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_payment_proto_init() }
func file_payment_proto_init() {
	if File_payment_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_payment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_payment_proto_goTypes,
		DependencyIndexes: file_payment_proto_depIdxs,
		MessageInfos:      file_payment_proto_msgTypes,
	}.Build()
	File_payment_proto = out.File
	file_payment_proto_rawDesc = nil
	file_payment_proto_goTypes = nil
	file_payment_proto_depIdxs = nil
}
