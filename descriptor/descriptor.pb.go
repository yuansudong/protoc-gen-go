package descriptor

import (
	reflect "reflect"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
)

// Symbols defined in public import of google/protobuf/descriptor.proto.

type FieldDescriptorProto_Type = descriptorpb.FieldDescriptorProto_Type

const FieldDescriptorProto_TYPE_DOUBLE = descriptorpb.FieldDescriptorProto_TYPE_DOUBLE
const FieldDescriptorProto_TYPE_FLOAT = descriptorpb.FieldDescriptorProto_TYPE_FLOAT
const FieldDescriptorProto_TYPE_INT64 = descriptorpb.FieldDescriptorProto_TYPE_INT64
const FieldDescriptorProto_TYPE_UINT64 = descriptorpb.FieldDescriptorProto_TYPE_UINT64
const FieldDescriptorProto_TYPE_INT32 = descriptorpb.FieldDescriptorProto_TYPE_INT32
const FieldDescriptorProto_TYPE_FIXED64 = descriptorpb.FieldDescriptorProto_TYPE_FIXED64
const FieldDescriptorProto_TYPE_FIXED32 = descriptorpb.FieldDescriptorProto_TYPE_FIXED32
const FieldDescriptorProto_TYPE_BOOL = descriptorpb.FieldDescriptorProto_TYPE_BOOL
const FieldDescriptorProto_TYPE_STRING = descriptorpb.FieldDescriptorProto_TYPE_STRING
const FieldDescriptorProto_TYPE_GROUP = descriptorpb.FieldDescriptorProto_TYPE_GROUP
const FieldDescriptorProto_TYPE_MESSAGE = descriptorpb.FieldDescriptorProto_TYPE_MESSAGE
const FieldDescriptorProto_TYPE_BYTES = descriptorpb.FieldDescriptorProto_TYPE_BYTES
const FieldDescriptorProto_TYPE_UINT32 = descriptorpb.FieldDescriptorProto_TYPE_UINT32
const FieldDescriptorProto_TYPE_ENUM = descriptorpb.FieldDescriptorProto_TYPE_ENUM
const FieldDescriptorProto_TYPE_SFIXED32 = descriptorpb.FieldDescriptorProto_TYPE_SFIXED32
const FieldDescriptorProto_TYPE_SFIXED64 = descriptorpb.FieldDescriptorProto_TYPE_SFIXED64
const FieldDescriptorProto_TYPE_SINT32 = descriptorpb.FieldDescriptorProto_TYPE_SINT32
const FieldDescriptorProto_TYPE_SINT64 = descriptorpb.FieldDescriptorProto_TYPE_SINT64

var FieldDescriptorProto_Type_name = descriptorpb.FieldDescriptorProto_Type_name
var FieldDescriptorProto_Type_value = descriptorpb.FieldDescriptorProto_Type_value

type FieldDescriptorProto_Label = descriptorpb.FieldDescriptorProto_Label

const FieldDescriptorProto_LABEL_OPTIONAL = descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL
const FieldDescriptorProto_LABEL_REQUIRED = descriptorpb.FieldDescriptorProto_LABEL_REQUIRED
const FieldDescriptorProto_LABEL_REPEATED = descriptorpb.FieldDescriptorProto_LABEL_REPEATED

var FieldDescriptorProto_Label_name = descriptorpb.FieldDescriptorProto_Label_name
var FieldDescriptorProto_Label_value = descriptorpb.FieldDescriptorProto_Label_value

type FileOptions_OptimizeMode = descriptorpb.FileOptions_OptimizeMode

const FileOptions_SPEED = descriptorpb.FileOptions_SPEED
const FileOptions_CODE_SIZE = descriptorpb.FileOptions_CODE_SIZE
const FileOptions_LITE_RUNTIME = descriptorpb.FileOptions_LITE_RUNTIME

var FileOptions_OptimizeMode_name = descriptorpb.FileOptions_OptimizeMode_name
var FileOptions_OptimizeMode_value = descriptorpb.FileOptions_OptimizeMode_value

type FieldOptions_CType = descriptorpb.FieldOptions_CType

const FieldOptions_STRING = descriptorpb.FieldOptions_STRING
const FieldOptions_CORD = descriptorpb.FieldOptions_CORD
const FieldOptions_STRING_PIECE = descriptorpb.FieldOptions_STRING_PIECE

var FieldOptions_CType_name = descriptorpb.FieldOptions_CType_name
var FieldOptions_CType_value = descriptorpb.FieldOptions_CType_value

type FieldOptions_JSType = descriptorpb.FieldOptions_JSType

const FieldOptions_JS_NORMAL = descriptorpb.FieldOptions_JS_NORMAL
const FieldOptions_JS_STRING = descriptorpb.FieldOptions_JS_STRING
const FieldOptions_JS_NUMBER = descriptorpb.FieldOptions_JS_NUMBER

var FieldOptions_JSType_name = descriptorpb.FieldOptions_JSType_name
var FieldOptions_JSType_value = descriptorpb.FieldOptions_JSType_value

type MethodOptions_IdempotencyLevel = descriptorpb.MethodOptions_IdempotencyLevel

const MethodOptions_IDEMPOTENCY_UNKNOWN = descriptorpb.MethodOptions_IDEMPOTENCY_UNKNOWN
const MethodOptions_NO_SIDE_EFFECTS = descriptorpb.MethodOptions_NO_SIDE_EFFECTS
const MethodOptions_IDEMPOTENT = descriptorpb.MethodOptions_IDEMPOTENT

var MethodOptions_IdempotencyLevel_name = descriptorpb.MethodOptions_IdempotencyLevel_name
var MethodOptions_IdempotencyLevel_value = descriptorpb.MethodOptions_IdempotencyLevel_value

type FileDescriptorSet = descriptorpb.FileDescriptorSet
type FileDescriptorProto = descriptorpb.FileDescriptorProto
type DescriptorProto = descriptorpb.DescriptorProto
type ExtensionRangeOptions = descriptorpb.ExtensionRangeOptions
type FieldDescriptorProto = descriptorpb.FieldDescriptorProto
type OneofDescriptorProto = descriptorpb.OneofDescriptorProto
type EnumDescriptorProto = descriptorpb.EnumDescriptorProto
type EnumValueDescriptorProto = descriptorpb.EnumValueDescriptorProto
type ServiceDescriptorProto = descriptorpb.ServiceDescriptorProto
type MethodDescriptorProto = descriptorpb.MethodDescriptorProto

const Default_MethodDescriptorProto_ClientStreaming = descriptorpb.Default_MethodDescriptorProto_ClientStreaming
const Default_MethodDescriptorProto_ServerStreaming = descriptorpb.Default_MethodDescriptorProto_ServerStreaming

type FileOptions = descriptorpb.FileOptions

const Default_FileOptions_JavaMultipleFiles = descriptorpb.Default_FileOptions_JavaMultipleFiles
const Default_FileOptions_JavaStringCheckUtf8 = descriptorpb.Default_FileOptions_JavaStringCheckUtf8
const Default_FileOptions_OptimizeFor = descriptorpb.Default_FileOptions_OptimizeFor
const Default_FileOptions_CcGenericServices = descriptorpb.Default_FileOptions_CcGenericServices
const Default_FileOptions_JavaGenericServices = descriptorpb.Default_FileOptions_JavaGenericServices
const Default_FileOptions_PyGenericServices = descriptorpb.Default_FileOptions_PyGenericServices
const Default_FileOptions_PhpGenericServices = descriptorpb.Default_FileOptions_PhpGenericServices
const Default_FileOptions_Deprecated = descriptorpb.Default_FileOptions_Deprecated
const Default_FileOptions_CcEnableArenas = descriptorpb.Default_FileOptions_CcEnableArenas

type MessageOptions = descriptorpb.MessageOptions

const Default_MessageOptions_MessageSetWireFormat = descriptorpb.Default_MessageOptions_MessageSetWireFormat
const Default_MessageOptions_NoStandardDescriptorAccessor = descriptorpb.Default_MessageOptions_NoStandardDescriptorAccessor
const Default_MessageOptions_Deprecated = descriptorpb.Default_MessageOptions_Deprecated

type FieldOptions = descriptorpb.FieldOptions

const Default_FieldOptions_Ctype = descriptorpb.Default_FieldOptions_Ctype
const Default_FieldOptions_Jstype = descriptorpb.Default_FieldOptions_Jstype
const Default_FieldOptions_Lazy = descriptorpb.Default_FieldOptions_Lazy
const Default_FieldOptions_Deprecated = descriptorpb.Default_FieldOptions_Deprecated
const Default_FieldOptions_Weak = descriptorpb.Default_FieldOptions_Weak

type OneofOptions = descriptorpb.OneofOptions
type EnumOptions = descriptorpb.EnumOptions

const Default_EnumOptions_Deprecated = descriptorpb.Default_EnumOptions_Deprecated

type EnumValueOptions = descriptorpb.EnumValueOptions

const Default_EnumValueOptions_Deprecated = descriptorpb.Default_EnumValueOptions_Deprecated

type ServiceOptions = descriptorpb.ServiceOptions

const Default_ServiceOptions_Deprecated = descriptorpb.Default_ServiceOptions_Deprecated

type MethodOptions = descriptorpb.MethodOptions

const Default_MethodOptions_Deprecated = descriptorpb.Default_MethodOptions_Deprecated
const Default_MethodOptions_IdempotencyLevel = descriptorpb.Default_MethodOptions_IdempotencyLevel

type UninterpretedOption = descriptorpb.UninterpretedOption
type SourceCodeInfo = descriptorpb.SourceCodeInfo
type GeneratedCodeInfo = descriptorpb.GeneratedCodeInfo
type DescriptorProto_ExtensionRange = descriptorpb.DescriptorProto_ExtensionRange
type DescriptorProto_ReservedRange = descriptorpb.DescriptorProto_ReservedRange
type EnumDescriptorProto_EnumReservedRange = descriptorpb.EnumDescriptorProto_EnumReservedRange
type UninterpretedOption_NamePart = descriptorpb.UninterpretedOption_NamePart
type SourceCodeInfo_Location = descriptorpb.SourceCodeInfo_Location
type GeneratedCodeInfo_Annotation = descriptorpb.GeneratedCodeInfo_Annotation

var File_github_com_golang_protobuf_protoc_gen_go_descriptor_descriptor_proto protoreflect.FileDescriptor

var file_github_com_golang_protobuf_protoc_gen_go_descriptor_descriptor_proto_rawDesc = []byte{
	0x0a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x3b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x50, 0x00, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x32,
}

var file_github_com_golang_protobuf_protoc_gen_go_descriptor_descriptor_proto_goTypes = []interface{}{}
var file_github_com_golang_protobuf_protoc_gen_go_descriptor_descriptor_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_github_com_golang_protobuf_protoc_gen_go_descriptor_descriptor_proto_init() }
func file_github_com_golang_protobuf_protoc_gen_go_descriptor_descriptor_proto_init() {
	if File_github_com_golang_protobuf_protoc_gen_go_descriptor_descriptor_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_golang_protobuf_protoc_gen_go_descriptor_descriptor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_golang_protobuf_protoc_gen_go_descriptor_descriptor_proto_goTypes,
		DependencyIndexes: file_github_com_golang_protobuf_protoc_gen_go_descriptor_descriptor_proto_depIdxs,
	}.Build()
	File_github_com_golang_protobuf_protoc_gen_go_descriptor_descriptor_proto = out.File
	file_github_com_golang_protobuf_protoc_gen_go_descriptor_descriptor_proto_rawDesc = nil
	file_github_com_golang_protobuf_protoc_gen_go_descriptor_descriptor_proto_goTypes = nil
	file_github_com_golang_protobuf_protoc_gen_go_descriptor_descriptor_proto_depIdxs = nil
}
