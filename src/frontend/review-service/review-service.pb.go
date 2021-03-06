// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: review-service/review-service.proto

package review_service

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

type AddReviewRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId   string `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Rating      int32  `protobuf:"varint,2,opt,name=rating,proto3" json:"rating,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *AddReviewRequest) Reset() {
	*x = AddReviewRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_review_service_review_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddReviewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddReviewRequest) ProtoMessage() {}

func (x *AddReviewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_review_service_review_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddReviewRequest.ProtoReflect.Descriptor instead.
func (*AddReviewRequest) Descriptor() ([]byte, []int) {
	return file_review_service_review_service_proto_rawDescGZIP(), []int{0}
}

func (x *AddReviewRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *AddReviewRequest) GetRating() int32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *AddReviewRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type GetReviewsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId string `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
}

func (x *GetReviewsRequest) Reset() {
	*x = GetReviewsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_review_service_review_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReviewsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReviewsRequest) ProtoMessage() {}

func (x *GetReviewsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_review_service_review_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReviewsRequest.ProtoReflect.Descriptor instead.
func (*GetReviewsRequest) Descriptor() ([]byte, []int) {
	return file_review_service_review_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetReviewsRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

type DeleteReviewRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReviewId string `protobuf:"bytes,1,opt,name=review_id,json=reviewId,proto3" json:"review_id,omitempty"`
}

func (x *DeleteReviewRequest) Reset() {
	*x = DeleteReviewRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_review_service_review_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteReviewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteReviewRequest) ProtoMessage() {}

func (x *DeleteReviewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_review_service_review_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteReviewRequest.ProtoReflect.Descriptor instead.
func (*DeleteReviewRequest) Descriptor() ([]byte, []int) {
	return file_review_service_review_service_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteReviewRequest) GetReviewId() string {
	if x != nil {
		return x.ReviewId
	}
	return ""
}

type Review struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReviewId    string `protobuf:"bytes,1,opt,name=review_id,json=reviewId,proto3" json:"review_id,omitempty"`
	ProductId   string `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Rating      int32  `protobuf:"varint,3,opt,name=rating,proto3" json:"rating,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Review) Reset() {
	*x = Review{}
	if protoimpl.UnsafeEnabled {
		mi := &file_review_service_review_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Review) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Review) ProtoMessage() {}

func (x *Review) ProtoReflect() protoreflect.Message {
	mi := &file_review_service_review_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Review.ProtoReflect.Descriptor instead.
func (*Review) Descriptor() ([]byte, []int) {
	return file_review_service_review_service_proto_rawDescGZIP(), []int{3}
}

func (x *Review) GetReviewId() string {
	if x != nil {
		return x.ReviewId
	}
	return ""
}

func (x *Review) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *Review) GetRating() int32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *Review) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_review_service_review_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_review_service_review_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_review_service_review_service_proto_rawDescGZIP(), []int{4}
}

var File_review_service_review_service_proto protoreflect.FileDescriptor

var file_review_service_review_service_proto_rawDesc = []byte{
	0x0a, 0x23, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6b, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x52, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x32, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x22, 0x32, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x49, 0x64, 0x22, 0x7e, 0x0a, 0x06, 0x52, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x32, 0x98, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x52, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x12, 0x11, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12,
	0x2d, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x12, 0x12, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x22, 0x00, 0x30, 0x01, 0x12, 0x2e,
	0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x14,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x12,
	0x5a, 0x10, 0x2e, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_review_service_review_service_proto_rawDescOnce sync.Once
	file_review_service_review_service_proto_rawDescData = file_review_service_review_service_proto_rawDesc
)

func file_review_service_review_service_proto_rawDescGZIP() []byte {
	file_review_service_review_service_proto_rawDescOnce.Do(func() {
		file_review_service_review_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_review_service_review_service_proto_rawDescData)
	})
	return file_review_service_review_service_proto_rawDescData
}

var file_review_service_review_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_review_service_review_service_proto_goTypes = []interface{}{
	(*AddReviewRequest)(nil),    // 0: AddReviewRequest
	(*GetReviewsRequest)(nil),   // 1: GetReviewsRequest
	(*DeleteReviewRequest)(nil), // 2: DeleteReviewRequest
	(*Review)(nil),              // 3: Review
	(*Empty)(nil),               // 4: Empty
}
var file_review_service_review_service_proto_depIdxs = []int32{
	0, // 0: ReviewService.AddReview:input_type -> AddReviewRequest
	1, // 1: ReviewService.GetReviews:input_type -> GetReviewsRequest
	2, // 2: ReviewService.DeleteReview:input_type -> DeleteReviewRequest
	4, // 3: ReviewService.AddReview:output_type -> Empty
	3, // 4: ReviewService.GetReviews:output_type -> Review
	4, // 5: ReviewService.DeleteReview:output_type -> Empty
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_review_service_review_service_proto_init() }
func file_review_service_review_service_proto_init() {
	if File_review_service_review_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_review_service_review_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddReviewRequest); i {
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
		file_review_service_review_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReviewsRequest); i {
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
		file_review_service_review_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteReviewRequest); i {
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
		file_review_service_review_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Review); i {
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
		file_review_service_review_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_review_service_review_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_review_service_review_service_proto_goTypes,
		DependencyIndexes: file_review_service_review_service_proto_depIdxs,
		MessageInfos:      file_review_service_review_service_proto_msgTypes,
	}.Build()
	File_review_service_review_service_proto = out.File
	file_review_service_review_service_proto_rawDesc = nil
	file_review_service_review_service_proto_goTypes = nil
	file_review_service_review_service_proto_depIdxs = nil
}
