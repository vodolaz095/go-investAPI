// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.0
// source: common.proto

package investapi

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

//Режим торгов инструмента
type SecurityTradingStatus int32

const (
	SecurityTradingStatus_SECURITY_TRADING_STATUS_UNSPECIFIED                      SecurityTradingStatus = 0  //Торговый статус не определён
	SecurityTradingStatus_SECURITY_TRADING_STATUS_NOT_AVAILABLE_FOR_TRADING        SecurityTradingStatus = 1  //Недоступен для торгов
	SecurityTradingStatus_SECURITY_TRADING_STATUS_OPENING_PERIOD                   SecurityTradingStatus = 2  //Период открытия торгов
	SecurityTradingStatus_SECURITY_TRADING_STATUS_CLOSING_PERIOD                   SecurityTradingStatus = 3  //Период закрытия торгов
	SecurityTradingStatus_SECURITY_TRADING_STATUS_BREAK_IN_TRADING                 SecurityTradingStatus = 4  //Перерыв в торговле
	SecurityTradingStatus_SECURITY_TRADING_STATUS_NORMAL_TRADING                   SecurityTradingStatus = 5  //Нормальная торговля
	SecurityTradingStatus_SECURITY_TRADING_STATUS_CLOSING_AUCTION                  SecurityTradingStatus = 6  //Аукцион закрытия
	SecurityTradingStatus_SECURITY_TRADING_STATUS_DARK_POOL_AUCTION                SecurityTradingStatus = 7  //Аукцион крупных пакетов
	SecurityTradingStatus_SECURITY_TRADING_STATUS_DISCRETE_AUCTION                 SecurityTradingStatus = 8  //Дискретный аукцион
	SecurityTradingStatus_SECURITY_TRADING_STATUS_OPENING_AUCTION_PERIOD           SecurityTradingStatus = 9  //Аукцион открытия
	SecurityTradingStatus_SECURITY_TRADING_STATUS_TRADING_AT_CLOSING_AUCTION_PRICE SecurityTradingStatus = 10 //Период торгов по цене аукциона закрытия
	SecurityTradingStatus_SECURITY_TRADING_STATUS_SESSION_ASSIGNED                 SecurityTradingStatus = 11 //Сессия назначена
	SecurityTradingStatus_SECURITY_TRADING_STATUS_SESSION_CLOSE                    SecurityTradingStatus = 12 //Сессия закрыта
	SecurityTradingStatus_SECURITY_TRADING_STATUS_SESSION_OPEN                     SecurityTradingStatus = 13 //Сессия открыта
	SecurityTradingStatus_SECURITY_TRADING_STATUS_DEALER_NORMAL_TRADING            SecurityTradingStatus = 14 //Доступна торговля в режиме внутренней ликвидности брокера
	SecurityTradingStatus_SECURITY_TRADING_STATUS_DEALER_BREAK_IN_TRADING          SecurityTradingStatus = 15 //Перерыв торговли в режиме внутренней ликвидности брокера
	SecurityTradingStatus_SECURITY_TRADING_STATUS_DEALER_NOT_AVAILABLE_FOR_TRADING SecurityTradingStatus = 16 //Недоступна торговля в режиме внутренней ликвидности брокера
)

// Enum value maps for SecurityTradingStatus.
var (
	SecurityTradingStatus_name = map[int32]string{
		0:  "SECURITY_TRADING_STATUS_UNSPECIFIED",
		1:  "SECURITY_TRADING_STATUS_NOT_AVAILABLE_FOR_TRADING",
		2:  "SECURITY_TRADING_STATUS_OPENING_PERIOD",
		3:  "SECURITY_TRADING_STATUS_CLOSING_PERIOD",
		4:  "SECURITY_TRADING_STATUS_BREAK_IN_TRADING",
		5:  "SECURITY_TRADING_STATUS_NORMAL_TRADING",
		6:  "SECURITY_TRADING_STATUS_CLOSING_AUCTION",
		7:  "SECURITY_TRADING_STATUS_DARK_POOL_AUCTION",
		8:  "SECURITY_TRADING_STATUS_DISCRETE_AUCTION",
		9:  "SECURITY_TRADING_STATUS_OPENING_AUCTION_PERIOD",
		10: "SECURITY_TRADING_STATUS_TRADING_AT_CLOSING_AUCTION_PRICE",
		11: "SECURITY_TRADING_STATUS_SESSION_ASSIGNED",
		12: "SECURITY_TRADING_STATUS_SESSION_CLOSE",
		13: "SECURITY_TRADING_STATUS_SESSION_OPEN",
		14: "SECURITY_TRADING_STATUS_DEALER_NORMAL_TRADING",
		15: "SECURITY_TRADING_STATUS_DEALER_BREAK_IN_TRADING",
		16: "SECURITY_TRADING_STATUS_DEALER_NOT_AVAILABLE_FOR_TRADING",
	}
	SecurityTradingStatus_value = map[string]int32{
		"SECURITY_TRADING_STATUS_UNSPECIFIED":                      0,
		"SECURITY_TRADING_STATUS_NOT_AVAILABLE_FOR_TRADING":        1,
		"SECURITY_TRADING_STATUS_OPENING_PERIOD":                   2,
		"SECURITY_TRADING_STATUS_CLOSING_PERIOD":                   3,
		"SECURITY_TRADING_STATUS_BREAK_IN_TRADING":                 4,
		"SECURITY_TRADING_STATUS_NORMAL_TRADING":                   5,
		"SECURITY_TRADING_STATUS_CLOSING_AUCTION":                  6,
		"SECURITY_TRADING_STATUS_DARK_POOL_AUCTION":                7,
		"SECURITY_TRADING_STATUS_DISCRETE_AUCTION":                 8,
		"SECURITY_TRADING_STATUS_OPENING_AUCTION_PERIOD":           9,
		"SECURITY_TRADING_STATUS_TRADING_AT_CLOSING_AUCTION_PRICE": 10,
		"SECURITY_TRADING_STATUS_SESSION_ASSIGNED":                 11,
		"SECURITY_TRADING_STATUS_SESSION_CLOSE":                    12,
		"SECURITY_TRADING_STATUS_SESSION_OPEN":                     13,
		"SECURITY_TRADING_STATUS_DEALER_NORMAL_TRADING":            14,
		"SECURITY_TRADING_STATUS_DEALER_BREAK_IN_TRADING":          15,
		"SECURITY_TRADING_STATUS_DEALER_NOT_AVAILABLE_FOR_TRADING": 16,
	}
)

func (x SecurityTradingStatus) Enum() *SecurityTradingStatus {
	p := new(SecurityTradingStatus)
	*p = x
	return p
}

func (x SecurityTradingStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SecurityTradingStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[0].Descriptor()
}

func (SecurityTradingStatus) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[0]
}

func (x SecurityTradingStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SecurityTradingStatus.Descriptor instead.
func (SecurityTradingStatus) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

//Денежная сумма в определенной валюте
type MoneyValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// строковый ISO-код валюты
	Currency string `protobuf:"bytes,1,opt,name=currency,proto3" json:"currency,omitempty"`
	// целая часть суммы, может быть отрицательным числом
	Units int64 `protobuf:"varint,2,opt,name=units,proto3" json:"units,omitempty"`
	// дробная часть суммы, может быть отрицательным числом
	Nano int32 `protobuf:"varint,3,opt,name=nano,proto3" json:"nano,omitempty"`
}

func (x *MoneyValue) Reset() {
	*x = MoneyValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MoneyValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoneyValue) ProtoMessage() {}

func (x *MoneyValue) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoneyValue.ProtoReflect.Descriptor instead.
func (*MoneyValue) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

func (x *MoneyValue) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *MoneyValue) GetUnits() int64 {
	if x != nil {
		return x.Units
	}
	return 0
}

func (x *MoneyValue) GetNano() int32 {
	if x != nil {
		return x.Nano
	}
	return 0
}

//Котировка - денежная сумма без указания валюты
type Quotation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// целая часть суммы, может быть отрицательным числом
	Units int64 `protobuf:"varint,1,opt,name=units,proto3" json:"units,omitempty"`
	// дробная часть суммы, может быть отрицательным числом
	Nano int32 `protobuf:"varint,2,opt,name=nano,proto3" json:"nano,omitempty"`
}

func (x *Quotation) Reset() {
	*x = Quotation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Quotation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Quotation) ProtoMessage() {}

func (x *Quotation) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Quotation.ProtoReflect.Descriptor instead.
func (*Quotation) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1}
}

func (x *Quotation) GetUnits() int64 {
	if x != nil {
		return x.Units
	}
	return 0
}

func (x *Quotation) GetNano() int32 {
	if x != nil {
		return x.Nano
	}
	return 0
}

//Проверка активности стрима.
type Ping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//Время проверки.
	Time *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *Ping) Reset() {
	*x = Ping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ping) ProtoMessage() {}

func (x *Ping) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ping.ProtoReflect.Descriptor instead.
func (*Ping) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{2}
}

func (x *Ping) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

var File_common_proto protoreflect.FileDescriptor

var file_common_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x25,
	0x74, 0x69, 0x6e, 0x6b, 0x6f, 0x66, 0x66, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x69,
	0x6e, 0x76, 0x65, 0x73, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x52, 0x0a, 0x0a, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6e, 0x6f, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6e, 0x61, 0x6e, 0x6f, 0x22, 0x35, 0x0a, 0x09, 0x51, 0x75,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x6e, 0x69, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6e, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6e, 0x61, 0x6e,
	0x6f, 0x22, 0x36, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x2a, 0xce, 0x06, 0x0a, 0x15, 0x53, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x54, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x27, 0x0a, 0x23, 0x53, 0x45, 0x43, 0x55, 0x52, 0x49, 0x54, 0x59, 0x5f,
	0x54, 0x52, 0x41, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x35, 0x0a, 0x31,
	0x53, 0x45, 0x43, 0x55, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x52, 0x41, 0x44, 0x49, 0x4e, 0x47,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x56, 0x41, 0x49,
	0x4c, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x54, 0x52, 0x41, 0x44, 0x49, 0x4e,
	0x47, 0x10, 0x01, 0x12, 0x2a, 0x0a, 0x26, 0x53, 0x45, 0x43, 0x55, 0x52, 0x49, 0x54, 0x59, 0x5f,
	0x54, 0x52, 0x41, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4f,
	0x50, 0x45, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x50, 0x45, 0x52, 0x49, 0x4f, 0x44, 0x10, 0x02, 0x12,
	0x2a, 0x0a, 0x26, 0x53, 0x45, 0x43, 0x55, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x52, 0x41, 0x44,
	0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x49,
	0x4e, 0x47, 0x5f, 0x50, 0x45, 0x52, 0x49, 0x4f, 0x44, 0x10, 0x03, 0x12, 0x2c, 0x0a, 0x28, 0x53,
	0x45, 0x43, 0x55, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x52, 0x41, 0x44, 0x49, 0x4e, 0x47, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x42, 0x52, 0x45, 0x41, 0x4b, 0x5f, 0x49, 0x4e, 0x5f,
	0x54, 0x52, 0x41, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x12, 0x2a, 0x0a, 0x26, 0x53, 0x45, 0x43,
	0x55, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x52, 0x41, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x5f, 0x54, 0x52, 0x41, 0x44,
	0x49, 0x4e, 0x47, 0x10, 0x05, 0x12, 0x2b, 0x0a, 0x27, 0x53, 0x45, 0x43, 0x55, 0x52, 0x49, 0x54,
	0x59, 0x5f, 0x54, 0x52, 0x41, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x49, 0x4e, 0x47, 0x5f, 0x41, 0x55, 0x43, 0x54, 0x49, 0x4f, 0x4e,
	0x10, 0x06, 0x12, 0x2d, 0x0a, 0x29, 0x53, 0x45, 0x43, 0x55, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x54,
	0x52, 0x41, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x41,
	0x52, 0x4b, 0x5f, 0x50, 0x4f, 0x4f, 0x4c, 0x5f, 0x41, 0x55, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10,
	0x07, 0x12, 0x2c, 0x0a, 0x28, 0x53, 0x45, 0x43, 0x55, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x52,
	0x41, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x49, 0x53,
	0x43, 0x52, 0x45, 0x54, 0x45, 0x5f, 0x41, 0x55, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x08, 0x12,
	0x32, 0x0a, 0x2e, 0x53, 0x45, 0x43, 0x55, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x52, 0x41, 0x44,
	0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4f, 0x50, 0x45, 0x4e, 0x49,
	0x4e, 0x47, 0x5f, 0x41, 0x55, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x50, 0x45, 0x52, 0x49, 0x4f,
	0x44, 0x10, 0x09, 0x12, 0x3c, 0x0a, 0x38, 0x53, 0x45, 0x43, 0x55, 0x52, 0x49, 0x54, 0x59, 0x5f,
	0x54, 0x52, 0x41, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x54,
	0x52, 0x41, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x41, 0x54, 0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x49, 0x4e,
	0x47, 0x5f, 0x41, 0x55, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x50, 0x52, 0x49, 0x43, 0x45, 0x10,
	0x0a, 0x12, 0x2c, 0x0a, 0x28, 0x53, 0x45, 0x43, 0x55, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x52,
	0x41, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x45, 0x53,
	0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x53, 0x53, 0x49, 0x47, 0x4e, 0x45, 0x44, 0x10, 0x0b, 0x12,
	0x29, 0x0a, 0x25, 0x53, 0x45, 0x43, 0x55, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x52, 0x41, 0x44,
	0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x45, 0x53, 0x53, 0x49,
	0x4f, 0x4e, 0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x10, 0x0c, 0x12, 0x28, 0x0a, 0x24, 0x53, 0x45,
	0x43, 0x55, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x52, 0x41, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x4f, 0x50,
	0x45, 0x4e, 0x10, 0x0d, 0x12, 0x31, 0x0a, 0x2d, 0x53, 0x45, 0x43, 0x55, 0x52, 0x49, 0x54, 0x59,
	0x5f, 0x54, 0x52, 0x41, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x44, 0x45, 0x41, 0x4c, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x5f, 0x54, 0x52,
	0x41, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x0e, 0x12, 0x33, 0x0a, 0x2f, 0x53, 0x45, 0x43, 0x55, 0x52,
	0x49, 0x54, 0x59, 0x5f, 0x54, 0x52, 0x41, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x44, 0x45, 0x41, 0x4c, 0x45, 0x52, 0x5f, 0x42, 0x52, 0x45, 0x41, 0x4b, 0x5f,
	0x49, 0x4e, 0x5f, 0x54, 0x52, 0x41, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x0f, 0x12, 0x3c, 0x0a, 0x38,
	0x53, 0x45, 0x43, 0x55, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x52, 0x41, 0x44, 0x49, 0x4e, 0x47,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x45, 0x41, 0x4c, 0x45, 0x52, 0x5f, 0x4e,
	0x4f, 0x54, 0x5f, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x46, 0x4f, 0x52,
	0x5f, 0x54, 0x52, 0x41, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x10, 0x42, 0x61, 0x0a, 0x1c, 0x72, 0x75,
	0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x6f, 0x66, 0x66, 0x2e, 0x70, 0x69, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x0c, 0x2e, 0x2f,
	0x3b, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x61, 0x70, 0x69, 0xa2, 0x02, 0x05, 0x54, 0x49, 0x41,
	0x50, 0x49, 0xaa, 0x02, 0x14, 0x54, 0x69, 0x6e, 0x6b, 0x6f, 0x66, 0x66, 0x2e, 0x49, 0x6e, 0x76,
	0x65, 0x73, 0x74, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x11, 0x54, 0x69, 0x6e, 0x6b,
	0x6f, 0x66, 0x66, 0x5c, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x5c, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_proto_rawDescOnce sync.Once
	file_common_proto_rawDescData = file_common_proto_rawDesc
)

func file_common_proto_rawDescGZIP() []byte {
	file_common_proto_rawDescOnce.Do(func() {
		file_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_rawDescData)
	})
	return file_common_proto_rawDescData
}

var file_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_common_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_common_proto_goTypes = []interface{}{
	(SecurityTradingStatus)(0),    // 0: tinkoff.public.invest.api.contract.v1.SecurityTradingStatus
	(*MoneyValue)(nil),            // 1: tinkoff.public.invest.api.contract.v1.MoneyValue
	(*Quotation)(nil),             // 2: tinkoff.public.invest.api.contract.v1.Quotation
	(*Ping)(nil),                  // 3: tinkoff.public.invest.api.contract.v1.Ping
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_common_proto_depIdxs = []int32{
	4, // 0: tinkoff.public.invest.api.contract.v1.Ping.time:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_common_proto_init() }
func file_common_proto_init() {
	if File_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MoneyValue); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Quotation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ping); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_proto_goTypes,
		DependencyIndexes: file_common_proto_depIdxs,
		EnumInfos:         file_common_proto_enumTypes,
		MessageInfos:      file_common_proto_msgTypes,
	}.Build()
	File_common_proto = out.File
	file_common_proto_rawDesc = nil
	file_common_proto_goTypes = nil
	file_common_proto_depIdxs = nil
}
