/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-gogo.
// source: k8s.io/kubernetes/pkg/api/resource/generated.proto
// DO NOT EDIT!

/*
	Package resource is a generated protocol buffer package.

	It is generated from these files:
		k8s.io/kubernetes/pkg/api/resource/generated.proto

	It has these top-level messages:
		Quantity
*/
package resource

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.GoGoProtoPackageIsVersion1

func (m *Quantity) Reset()                    { *m = Quantity{} }
func (*Quantity) ProtoMessage()               {}
func (*Quantity) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func init() {
	proto.RegisterType((*Quantity)(nil), "k8s.io.client-go.1.5.pkg.api.resource.Quantity")
}

var fileDescriptorGenerated = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x32, 0xca, 0xb6, 0x28, 0xd6,
	0xcb, 0xcc, 0xd7, 0xcf, 0x2e, 0x4d, 0x4a, 0x2d, 0xca, 0x4b, 0x2d, 0x49, 0x2d, 0xd6, 0x2f, 0xc8,
	0x4e, 0xd7, 0x4f, 0x2c, 0xc8, 0xd4, 0x2f, 0x4a, 0x2d, 0xce, 0x2f, 0x2d, 0x4a, 0x4e, 0xd5, 0x4f,
	0x4f, 0xcd, 0x4b, 0x2d, 0x4a, 0x2c, 0x49, 0x4d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52,
	0x82, 0xe8, 0xd1, 0x43, 0xe8, 0xd1, 0x03, 0xea, 0xd1, 0x03, 0xea, 0xd1, 0x83, 0xe9, 0x91, 0xd2,
	0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0x4f, 0xcf, 0xd7,
	0x07, 0x6b, 0x4d, 0x2a, 0x4d, 0x03, 0xf3, 0xc0, 0x1c, 0x30, 0x0b, 0x62, 0xa4, 0x94, 0x21, 0x76,
	0x67, 0x94, 0x96, 0x64, 0xe6, 0xe8, 0x67, 0xe6, 0x95, 0x14, 0x97, 0x14, 0xa1, 0xbb, 0x42, 0xc9,
	0x82, 0x8b, 0x23, 0xb0, 0x34, 0x31, 0xaf, 0x24, 0xb3, 0xa4, 0x52, 0x48, 0x8c, 0x8b, 0x0d, 0xa8,
	0x24, 0x33, 0x2f, 0x5d, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xca, 0xb3, 0x12, 0x99, 0xb1,
	0x40, 0x9e, 0xa1, 0x63, 0xa1, 0x3c, 0xc3, 0x04, 0x20, 0x5e, 0x00, 0xc4, 0x0d, 0x77, 0x14, 0x18,
	0x9c, 0xb4, 0x4e, 0x3c, 0x94, 0x63, 0xb8, 0x00, 0xc4, 0x37, 0x80, 0xb8, 0xe1, 0x91, 0x1c, 0xe3,
	0x09, 0x20, 0xbe, 0x00, 0xc4, 0x0f, 0x80, 0x78, 0xc2, 0x63, 0x39, 0x86, 0x28, 0x0e, 0x98, 0x3f,
	0x00, 0x01, 0x00, 0x00, 0xff, 0xff, 0x90, 0x1c, 0x7f, 0xff, 0x20, 0x01, 0x00, 0x00,
}
