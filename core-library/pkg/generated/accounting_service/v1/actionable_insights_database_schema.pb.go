// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: accounting_service/v1/actionable_insights_database_schema.proto

package accounting_servicev1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
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

type BusinessActionableInsight struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique identifier for the insight
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Type of the actionable insight
	Type BusinessActionableInsightType `protobuf:"varint,2,opt,name=type,proto3,enum=accounting_service.v1.BusinessActionableInsightType" json:"type,omitempty"`
	// Detailed description of the insight
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Key takeaway or summary of the insight
	Takeaway string `protobuf:"bytes,4,opt,name=takeaway,proto3" json:"takeaway,omitempty"`
	// Suggested action based on the insight
	Action string `protobuf:"bytes,5,opt,name=action,proto3" json:"action,omitempty"`
	// Expected benefit or outcome of taking the suggested action
	ExpectedBenefit string `protobuf:"bytes,6,opt,name=expected_benefit,json=expectedBenefit,proto3" json:"expected_benefit,omitempty"`
	// Tags associated with the insight for categorization
	Tags []string `protobuf:"bytes,7,rep,name=tags,proto3" json:"tags,omitempty"`
	// Time when the insight was generated
	GeneratedTime *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=generated_time,json=generatedTime,proto3" json:"generated_time,omitempty"`
	// Specific business metrics that this insight aims to optimize
	MetricsToOptimizeFor []string `protobuf:"bytes,9,rep,name=metrics_to_optimize_for,json=metricsToOptimizeFor,proto3" json:"metrics_to_optimize_for,omitempty"`
}

func (x *BusinessActionableInsight) Reset() {
	*x = BusinessActionableInsight{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounting_service_v1_actionable_insights_database_schema_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BusinessActionableInsight) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BusinessActionableInsight) ProtoMessage() {}

func (x *BusinessActionableInsight) ProtoReflect() protoreflect.Message {
	mi := &file_accounting_service_v1_actionable_insights_database_schema_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BusinessActionableInsight.ProtoReflect.Descriptor instead.
func (*BusinessActionableInsight) Descriptor() ([]byte, []int) {
	return file_accounting_service_v1_actionable_insights_database_schema_proto_rawDescGZIP(), []int{0}
}

func (x *BusinessActionableInsight) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BusinessActionableInsight) GetType() BusinessActionableInsightType {
	if x != nil {
		return x.Type
	}
	return BusinessActionableInsightType_BUSINESS_ACTIONABLE_INSIGHT_TYPE_UNSPECIFIED
}

func (x *BusinessActionableInsight) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *BusinessActionableInsight) GetTakeaway() string {
	if x != nil {
		return x.Takeaway
	}
	return ""
}

func (x *BusinessActionableInsight) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *BusinessActionableInsight) GetExpectedBenefit() string {
	if x != nil {
		return x.ExpectedBenefit
	}
	return ""
}

func (x *BusinessActionableInsight) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *BusinessActionableInsight) GetGeneratedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.GeneratedTime
	}
	return nil
}

func (x *BusinessActionableInsight) GetMetricsToOptimizeFor() []string {
	if x != nil {
		return x.MetricsToOptimizeFor
	}
	return nil
}

var File_accounting_service_v1_actionable_insights_database_schema_proto protoreflect.FileDescriptor

var file_accounting_service_v1_actionable_insights_database_schema_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x15, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x72, 0x6d, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x67, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x9b, 0x06, 0x0a, 0x19, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x12, 0x47,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x37, 0x92, 0x41, 0x0e, 0x4a,
	0x0c, 0x22, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x30, 0x22, 0xba, 0xb9, 0x19,
	0x22, 0x0a, 0x20, 0x5a, 0x1e, 0x69, 0x64, 0x78, 0x5f, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x5f, 0x62,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x5f, 0x69, 0x64, 0x52, 0x02, 0x69, 0x64, 0x12, 0x48, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x34, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69,
	0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75,
	0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x5b, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x39, 0x92, 0x41, 0x36, 0x4a, 0x34, 0x22, 0x52, 0x65,
	0x64, 0x75, 0x63, 0x65, 0x20, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x20, 0x63,
	0x6f, 0x73, 0x74, 0x73, 0x20, 0x62, 0x79, 0x20, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a, 0x69,
	0x6e, 0x67, 0x20, 0x65, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x20, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x22, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x60,
	0x0a, 0x08, 0x74, 0x61, 0x6b, 0x65, 0x61, 0x77, 0x61, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x44, 0x92, 0x41, 0x41, 0x4a, 0x3f, 0x22, 0x4f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a, 0x69,
	0x6e, 0x67, 0x20, 0x65, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x20, 0x75, 0x73, 0x61, 0x67, 0x65, 0x20,
	0x63, 0x61, 0x6e, 0x20, 0x6c, 0x65, 0x61, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x73, 0x69, 0x67, 0x6e,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x6e, 0x74, 0x20, 0x63, 0x6f, 0x73, 0x74, 0x20, 0x73, 0x61, 0x76,
	0x69, 0x6e, 0x67, 0x73, 0x2e, 0x22, 0x52, 0x08, 0x74, 0x61, 0x6b, 0x65, 0x61, 0x77, 0x61, 0x79,
	0x12, 0x51, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x39, 0x92, 0x41, 0x36, 0x4a, 0x34, 0x22, 0x49, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x20, 0x65, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x2d, 0x65, 0x66, 0x66, 0x69, 0x63, 0x69, 0x65,
	0x6e, 0x74, 0x20, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x69, 0x6e, 0x67, 0x20, 0x61, 0x6e, 0x64, 0x20,
	0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2e, 0x22, 0x52, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x68, 0x0a, 0x10, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f,
	0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3d, 0x92,
	0x41, 0x3a, 0x4a, 0x38, 0x22, 0x45, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x20, 0x74, 0x6f,
	0x20, 0x72, 0x65, 0x64, 0x75, 0x63, 0x65, 0x20, 0x65, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x20, 0x63,
	0x6f, 0x73, 0x74, 0x73, 0x20, 0x62, 0x79, 0x20, 0x75, 0x70, 0x20, 0x74, 0x6f, 0x20, 0x32, 0x30,
	0x25, 0x20, 0x61, 0x6e, 0x6e, 0x75, 0x61, 0x6c, 0x6c, 0x79, 0x2e, 0x22, 0x52, 0x0f, 0x65, 0x78,
	0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x12, 0x3d, 0x0a,
	0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x42, 0x29, 0x92, 0x41, 0x26,
	0x4a, 0x24, 0x5b, 0x22, 0x63, 0x6f, 0x73, 0x74, 0x2d, 0x73, 0x61, 0x76, 0x69, 0x6e, 0x67, 0x22,
	0x2c, 0x20, 0x22, 0x65, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x2d, 0x65, 0x66, 0x66, 0x69, 0x63, 0x69,
	0x65, 0x6e, 0x63, 0x79, 0x22, 0x5d, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x41, 0x0a, 0x0e,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0d, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x65, 0x0a, 0x17, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x5f, 0x74, 0x6f, 0x5f, 0x6f, 0x70,
	0x74, 0x69, 0x6d, 0x69, 0x7a, 0x65, 0x5f, 0x66, 0x6f, 0x72, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09,
	0x42, 0x2e, 0x92, 0x41, 0x2b, 0x4a, 0x29, 0x5b, 0x22, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x20, 0x63, 0x6f, 0x73, 0x74, 0x73, 0x22, 0x2c, 0x20, 0x22, 0x65, 0x6e, 0x65, 0x72,
	0x67, 0x79, 0x20, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x5d,
	0x52, 0x14, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x54, 0x6f, 0x4f, 0x70, 0x74, 0x69, 0x6d,
	0x69, 0x7a, 0x65, 0x46, 0x6f, 0x72, 0x3a, 0x06, 0xba, 0xb9, 0x19, 0x02, 0x08, 0x01, 0x42, 0xb6,
	0x02, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e,
	0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x25, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x73,
	0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x80, 0x01, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x53, 0x6f, 0x6c, 0x6f, 0x6d, 0x6f, 0x6e, 0x41, 0x49, 0x45, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2d, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69,
	0x6e, 0x67, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x76, 0x31, 0x3b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x58, 0x58, 0xaa, 0x02,
	0x14, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x14, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69,
	0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x20, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x15, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_accounting_service_v1_actionable_insights_database_schema_proto_rawDescOnce sync.Once
	file_accounting_service_v1_actionable_insights_database_schema_proto_rawDescData = file_accounting_service_v1_actionable_insights_database_schema_proto_rawDesc
)

func file_accounting_service_v1_actionable_insights_database_schema_proto_rawDescGZIP() []byte {
	file_accounting_service_v1_actionable_insights_database_schema_proto_rawDescOnce.Do(func() {
		file_accounting_service_v1_actionable_insights_database_schema_proto_rawDescData = protoimpl.X.CompressGZIP(file_accounting_service_v1_actionable_insights_database_schema_proto_rawDescData)
	})
	return file_accounting_service_v1_actionable_insights_database_schema_proto_rawDescData
}

var file_accounting_service_v1_actionable_insights_database_schema_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_accounting_service_v1_actionable_insights_database_schema_proto_goTypes = []interface{}{
	(*BusinessActionableInsight)(nil),  // 0: accounting_service.v1.BusinessActionableInsight
	(BusinessActionableInsightType)(0), // 1: accounting_service.v1.BusinessActionableInsightType
	(*timestamppb.Timestamp)(nil),      // 2: google.protobuf.Timestamp
}
var file_accounting_service_v1_actionable_insights_database_schema_proto_depIdxs = []int32{
	1, // 0: accounting_service.v1.BusinessActionableInsight.type:type_name -> accounting_service.v1.BusinessActionableInsightType
	2, // 1: accounting_service.v1.BusinessActionableInsight.generated_time:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_accounting_service_v1_actionable_insights_database_schema_proto_init() }
func file_accounting_service_v1_actionable_insights_database_schema_proto_init() {
	if File_accounting_service_v1_actionable_insights_database_schema_proto != nil {
		return
	}
	file_accounting_service_v1_enums_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_accounting_service_v1_actionable_insights_database_schema_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BusinessActionableInsight); i {
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
			RawDescriptor: file_accounting_service_v1_actionable_insights_database_schema_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_accounting_service_v1_actionable_insights_database_schema_proto_goTypes,
		DependencyIndexes: file_accounting_service_v1_actionable_insights_database_schema_proto_depIdxs,
		MessageInfos:      file_accounting_service_v1_actionable_insights_database_schema_proto_msgTypes,
	}.Build()
	File_accounting_service_v1_actionable_insights_database_schema_proto = out.File
	file_accounting_service_v1_actionable_insights_database_schema_proto_rawDesc = nil
	file_accounting_service_v1_actionable_insights_database_schema_proto_goTypes = nil
	file_accounting_service_v1_actionable_insights_database_schema_proto_depIdxs = nil
}