// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package blog

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_TimeoutPost         protoreflect.MessageDescriptor
	fd_TimeoutPost_id      protoreflect.FieldDescriptor
	fd_TimeoutPost_title   protoreflect.FieldDescriptor
	fd_TimeoutPost_chain   protoreflect.FieldDescriptor
	fd_TimeoutPost_creator protoreflect.FieldDescriptor
)

func init() {
	file_plante_blog_timeout_post_proto_init()
	md_TimeoutPost = File_plante_blog_timeout_post_proto.Messages().ByName("TimeoutPost")
	fd_TimeoutPost_id = md_TimeoutPost.Fields().ByName("id")
	fd_TimeoutPost_title = md_TimeoutPost.Fields().ByName("title")
	fd_TimeoutPost_chain = md_TimeoutPost.Fields().ByName("chain")
	fd_TimeoutPost_creator = md_TimeoutPost.Fields().ByName("creator")
}

var _ protoreflect.Message = (*fastReflection_TimeoutPost)(nil)

type fastReflection_TimeoutPost TimeoutPost

func (x *TimeoutPost) ProtoReflect() protoreflect.Message {
	return (*fastReflection_TimeoutPost)(x)
}

func (x *TimeoutPost) slowProtoReflect() protoreflect.Message {
	mi := &file_plante_blog_timeout_post_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_TimeoutPost_messageType fastReflection_TimeoutPost_messageType
var _ protoreflect.MessageType = fastReflection_TimeoutPost_messageType{}

type fastReflection_TimeoutPost_messageType struct{}

func (x fastReflection_TimeoutPost_messageType) Zero() protoreflect.Message {
	return (*fastReflection_TimeoutPost)(nil)
}
func (x fastReflection_TimeoutPost_messageType) New() protoreflect.Message {
	return new(fastReflection_TimeoutPost)
}
func (x fastReflection_TimeoutPost_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_TimeoutPost
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_TimeoutPost) Descriptor() protoreflect.MessageDescriptor {
	return md_TimeoutPost
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_TimeoutPost) Type() protoreflect.MessageType {
	return _fastReflection_TimeoutPost_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_TimeoutPost) New() protoreflect.Message {
	return new(fastReflection_TimeoutPost)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_TimeoutPost) Interface() protoreflect.ProtoMessage {
	return (*TimeoutPost)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_TimeoutPost) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Id != uint64(0) {
		value := protoreflect.ValueOfUint64(x.Id)
		if !f(fd_TimeoutPost_id, value) {
			return
		}
	}
	if x.Title != "" {
		value := protoreflect.ValueOfString(x.Title)
		if !f(fd_TimeoutPost_title, value) {
			return
		}
	}
	if x.Chain != "" {
		value := protoreflect.ValueOfString(x.Chain)
		if !f(fd_TimeoutPost_chain, value) {
			return
		}
	}
	if x.Creator != "" {
		value := protoreflect.ValueOfString(x.Creator)
		if !f(fd_TimeoutPost_creator, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_TimeoutPost) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "plante.blog.TimeoutPost.id":
		return x.Id != uint64(0)
	case "plante.blog.TimeoutPost.title":
		return x.Title != ""
	case "plante.blog.TimeoutPost.chain":
		return x.Chain != ""
	case "plante.blog.TimeoutPost.creator":
		return x.Creator != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: plante.blog.TimeoutPost"))
		}
		panic(fmt.Errorf("message plante.blog.TimeoutPost does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_TimeoutPost) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "plante.blog.TimeoutPost.id":
		x.Id = uint64(0)
	case "plante.blog.TimeoutPost.title":
		x.Title = ""
	case "plante.blog.TimeoutPost.chain":
		x.Chain = ""
	case "plante.blog.TimeoutPost.creator":
		x.Creator = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: plante.blog.TimeoutPost"))
		}
		panic(fmt.Errorf("message plante.blog.TimeoutPost does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_TimeoutPost) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "plante.blog.TimeoutPost.id":
		value := x.Id
		return protoreflect.ValueOfUint64(value)
	case "plante.blog.TimeoutPost.title":
		value := x.Title
		return protoreflect.ValueOfString(value)
	case "plante.blog.TimeoutPost.chain":
		value := x.Chain
		return protoreflect.ValueOfString(value)
	case "plante.blog.TimeoutPost.creator":
		value := x.Creator
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: plante.blog.TimeoutPost"))
		}
		panic(fmt.Errorf("message plante.blog.TimeoutPost does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_TimeoutPost) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "plante.blog.TimeoutPost.id":
		x.Id = value.Uint()
	case "plante.blog.TimeoutPost.title":
		x.Title = value.Interface().(string)
	case "plante.blog.TimeoutPost.chain":
		x.Chain = value.Interface().(string)
	case "plante.blog.TimeoutPost.creator":
		x.Creator = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: plante.blog.TimeoutPost"))
		}
		panic(fmt.Errorf("message plante.blog.TimeoutPost does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_TimeoutPost) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "plante.blog.TimeoutPost.id":
		panic(fmt.Errorf("field id of message plante.blog.TimeoutPost is not mutable"))
	case "plante.blog.TimeoutPost.title":
		panic(fmt.Errorf("field title of message plante.blog.TimeoutPost is not mutable"))
	case "plante.blog.TimeoutPost.chain":
		panic(fmt.Errorf("field chain of message plante.blog.TimeoutPost is not mutable"))
	case "plante.blog.TimeoutPost.creator":
		panic(fmt.Errorf("field creator of message plante.blog.TimeoutPost is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: plante.blog.TimeoutPost"))
		}
		panic(fmt.Errorf("message plante.blog.TimeoutPost does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_TimeoutPost) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "plante.blog.TimeoutPost.id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "plante.blog.TimeoutPost.title":
		return protoreflect.ValueOfString("")
	case "plante.blog.TimeoutPost.chain":
		return protoreflect.ValueOfString("")
	case "plante.blog.TimeoutPost.creator":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: plante.blog.TimeoutPost"))
		}
		panic(fmt.Errorf("message plante.blog.TimeoutPost does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_TimeoutPost) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in plante.blog.TimeoutPost", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_TimeoutPost) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_TimeoutPost) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_TimeoutPost) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_TimeoutPost) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*TimeoutPost)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.Id != 0 {
			n += 1 + runtime.Sov(uint64(x.Id))
		}
		l = len(x.Title)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Chain)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Creator)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*TimeoutPost)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Creator) > 0 {
			i -= len(x.Creator)
			copy(dAtA[i:], x.Creator)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Creator)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.Chain) > 0 {
			i -= len(x.Chain)
			copy(dAtA[i:], x.Chain)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Chain)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.Title) > 0 {
			i -= len(x.Title)
			copy(dAtA[i:], x.Title)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Title)))
			i--
			dAtA[i] = 0x12
		}
		if x.Id != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Id))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*TimeoutPost)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: TimeoutPost: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: TimeoutPost: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
				}
				x.Id = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Id |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Title = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Chain = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Creator = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: plante/blog/timeout_post.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TimeoutPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title   string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Chain   string `protobuf:"bytes,3,opt,name=chain,proto3" json:"chain,omitempty"`
	Creator string `protobuf:"bytes,4,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (x *TimeoutPost) Reset() {
	*x = TimeoutPost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plante_blog_timeout_post_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeoutPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeoutPost) ProtoMessage() {}

// Deprecated: Use TimeoutPost.ProtoReflect.Descriptor instead.
func (*TimeoutPost) Descriptor() ([]byte, []int) {
	return file_plante_blog_timeout_post_proto_rawDescGZIP(), []int{0}
}

func (x *TimeoutPost) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TimeoutPost) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *TimeoutPost) GetChain() string {
	if x != nil {
		return x.Chain
	}
	return ""
}

func (x *TimeoutPost) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

var File_plante_blog_timeout_post_proto protoreflect.FileDescriptor

var file_plante_blog_timeout_post_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x65, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0b, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x65, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x22, 0x63, 0x0a,
	0x0b, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x6f, 0x72, 0x42, 0x88, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x65, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x42, 0x10, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x50,
	0x6f, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x16, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x65, 0x2f, 0x62, 0x6c,
	0x6f, 0x67, 0xa2, 0x02, 0x03, 0x50, 0x42, 0x58, 0xaa, 0x02, 0x0b, 0x50, 0x6c, 0x61, 0x6e, 0x74,
	0x65, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0xca, 0x02, 0x0b, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x65, 0x5c,
	0x42, 0x6c, 0x6f, 0x67, 0xe2, 0x02, 0x17, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x65, 0x5c, 0x42, 0x6c,
	0x6f, 0x67, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x0c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x65, 0x3a, 0x3a, 0x42, 0x6c, 0x6f, 0x67, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_plante_blog_timeout_post_proto_rawDescOnce sync.Once
	file_plante_blog_timeout_post_proto_rawDescData = file_plante_blog_timeout_post_proto_rawDesc
)

func file_plante_blog_timeout_post_proto_rawDescGZIP() []byte {
	file_plante_blog_timeout_post_proto_rawDescOnce.Do(func() {
		file_plante_blog_timeout_post_proto_rawDescData = protoimpl.X.CompressGZIP(file_plante_blog_timeout_post_proto_rawDescData)
	})
	return file_plante_blog_timeout_post_proto_rawDescData
}

var file_plante_blog_timeout_post_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_plante_blog_timeout_post_proto_goTypes = []interface{}{
	(*TimeoutPost)(nil), // 0: plante.blog.TimeoutPost
}
var file_plante_blog_timeout_post_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_plante_blog_timeout_post_proto_init() }
func file_plante_blog_timeout_post_proto_init() {
	if File_plante_blog_timeout_post_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_plante_blog_timeout_post_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeoutPost); i {
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
			RawDescriptor: file_plante_blog_timeout_post_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_plante_blog_timeout_post_proto_goTypes,
		DependencyIndexes: file_plante_blog_timeout_post_proto_depIdxs,
		MessageInfos:      file_plante_blog_timeout_post_proto_msgTypes,
	}.Build()
	File_plante_blog_timeout_post_proto = out.File
	file_plante_blog_timeout_post_proto_rawDesc = nil
	file_plante_blog_timeout_post_proto_goTypes = nil
	file_plante_blog_timeout_post_proto_depIdxs = nil
}
