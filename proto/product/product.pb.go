// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/product/product.proto

/*
Package go_micro_service_product is a generated protocol buffer package.

It is generated from these files:
	proto/product/product.proto

It has these top-level messages:
	ProductInfo
	ProductImage
	ProductSize
	ProductSeo
	ResponseProduct
	RequestID
	Response
	RequestAll
	AllProduct
*/
package go_micro_service_product

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ProductInfo struct {
	Id                 int64           `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	ProductName        string          `protobuf:"bytes,2,opt,name=product_name,json=productName" json:"product_name,omitempty"`
	ProductSku         string          `protobuf:"bytes,3,opt,name=product_sku,json=productSku" json:"product_sku,omitempty"`
	ProductPrice       float64         `protobuf:"fixed64,4,opt,name=product_price,json=productPrice" json:"product_price,omitempty"`
	ProductDescription string          `protobuf:"bytes,5,opt,name=product_description,json=productDescription" json:"product_description,omitempty"`
	ProductCategoryId  int64           `protobuf:"varint,6,opt,name=product_category_id,json=productCategoryId" json:"product_category_id,omitempty"`
	ProductImage       []*ProductImage `protobuf:"bytes,7,rep,name=product_image,json=productImage" json:"product_image,omitempty"`
	ProductSize        []*ProductSize  `protobuf:"bytes,8,rep,name=product_size,json=productSize" json:"product_size,omitempty"`
	ProductSeo         []*ProductSeo   `protobuf:"bytes,9,rep,name=product_seo,json=productSeo" json:"product_seo,omitempty"`
}

func (m *ProductInfo) Reset()                    { *m = ProductInfo{} }
func (m *ProductInfo) String() string            { return proto.CompactTextString(m) }
func (*ProductInfo) ProtoMessage()               {}
func (*ProductInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ProductInfo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ProductInfo) GetProductName() string {
	if m != nil {
		return m.ProductName
	}
	return ""
}

func (m *ProductInfo) GetProductSku() string {
	if m != nil {
		return m.ProductSku
	}
	return ""
}

func (m *ProductInfo) GetProductPrice() float64 {
	if m != nil {
		return m.ProductPrice
	}
	return 0
}

func (m *ProductInfo) GetProductDescription() string {
	if m != nil {
		return m.ProductDescription
	}
	return ""
}

func (m *ProductInfo) GetProductCategoryId() int64 {
	if m != nil {
		return m.ProductCategoryId
	}
	return 0
}

func (m *ProductInfo) GetProductImage() []*ProductImage {
	if m != nil {
		return m.ProductImage
	}
	return nil
}

func (m *ProductInfo) GetProductSize() []*ProductSize {
	if m != nil {
		return m.ProductSize
	}
	return nil
}

func (m *ProductInfo) GetProductSeo() []*ProductSeo {
	if m != nil {
		return m.ProductSeo
	}
	return nil
}

type ProductImage struct {
	Id        int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	ImageName string `protobuf:"bytes,2,opt,name=image_name,json=imageName" json:"image_name,omitempty"`
	ImageCode string `protobuf:"bytes,3,opt,name=image_code,json=imageCode" json:"image_code,omitempty"`
	ImageUrl  string `protobuf:"bytes,4,opt,name=image_url,json=imageUrl" json:"image_url,omitempty"`
}

func (m *ProductImage) Reset()                    { *m = ProductImage{} }
func (m *ProductImage) String() string            { return proto.CompactTextString(m) }
func (*ProductImage) ProtoMessage()               {}
func (*ProductImage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ProductImage) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ProductImage) GetImageName() string {
	if m != nil {
		return m.ImageName
	}
	return ""
}

func (m *ProductImage) GetImageCode() string {
	if m != nil {
		return m.ImageCode
	}
	return ""
}

func (m *ProductImage) GetImageUrl() string {
	if m != nil {
		return m.ImageUrl
	}
	return ""
}

type ProductSize struct {
	Id       int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	SizeName string `protobuf:"bytes,2,opt,name=size_name,json=sizeName" json:"size_name,omitempty"`
	SizeCode string `protobuf:"bytes,3,opt,name=size_code,json=sizeCode" json:"size_code,omitempty"`
}

func (m *ProductSize) Reset()                    { *m = ProductSize{} }
func (m *ProductSize) String() string            { return proto.CompactTextString(m) }
func (*ProductSize) ProtoMessage()               {}
func (*ProductSize) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ProductSize) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ProductSize) GetSizeName() string {
	if m != nil {
		return m.SizeName
	}
	return ""
}

func (m *ProductSize) GetSizeCode() string {
	if m != nil {
		return m.SizeCode
	}
	return ""
}

type ProductSeo struct {
	Id             int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	SeoTitle       string `protobuf:"bytes,2,opt,name=seo_title,json=seoTitle" json:"seo_title,omitempty"`
	SeoKeywords    string `protobuf:"bytes,3,opt,name=seo_keywords,json=seoKeywords" json:"seo_keywords,omitempty"`
	SeoDescription string `protobuf:"bytes,4,opt,name=seo_description,json=seoDescription" json:"seo_description,omitempty"`
	SeoCode        string `protobuf:"bytes,6,opt,name=seo_code,json=seoCode" json:"seo_code,omitempty"`
}

func (m *ProductSeo) Reset()                    { *m = ProductSeo{} }
func (m *ProductSeo) String() string            { return proto.CompactTextString(m) }
func (*ProductSeo) ProtoMessage()               {}
func (*ProductSeo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ProductSeo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ProductSeo) GetSeoTitle() string {
	if m != nil {
		return m.SeoTitle
	}
	return ""
}

func (m *ProductSeo) GetSeoKeywords() string {
	if m != nil {
		return m.SeoKeywords
	}
	return ""
}

func (m *ProductSeo) GetSeoDescription() string {
	if m != nil {
		return m.SeoDescription
	}
	return ""
}

func (m *ProductSeo) GetSeoCode() string {
	if m != nil {
		return m.SeoCode
	}
	return ""
}

type ResponseProduct struct {
	ProductId int64 `protobuf:"varint,1,opt,name=product_id,json=productId" json:"product_id,omitempty"`
}

func (m *ResponseProduct) Reset()                    { *m = ResponseProduct{} }
func (m *ResponseProduct) String() string            { return proto.CompactTextString(m) }
func (*ResponseProduct) ProtoMessage()               {}
func (*ResponseProduct) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ResponseProduct) GetProductId() int64 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

type RequestID struct {
	ProductId int64 `protobuf:"varint,1,opt,name=product_id,json=productId" json:"product_id,omitempty"`
}

func (m *RequestID) Reset()                    { *m = RequestID{} }
func (m *RequestID) String() string            { return proto.CompactTextString(m) }
func (*RequestID) ProtoMessage()               {}
func (*RequestID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RequestID) GetProductId() int64 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

type Response struct {
	Msg string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type RequestAll struct {
}

func (m *RequestAll) Reset()                    { *m = RequestAll{} }
func (m *RequestAll) String() string            { return proto.CompactTextString(m) }
func (*RequestAll) ProtoMessage()               {}
func (*RequestAll) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type AllProduct struct {
	ProductInfo []*ProductInfo `protobuf:"bytes,1,rep,name=product_info,json=productInfo" json:"product_info,omitempty"`
}

func (m *AllProduct) Reset()                    { *m = AllProduct{} }
func (m *AllProduct) String() string            { return proto.CompactTextString(m) }
func (*AllProduct) ProtoMessage()               {}
func (*AllProduct) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *AllProduct) GetProductInfo() []*ProductInfo {
	if m != nil {
		return m.ProductInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*ProductInfo)(nil), "go.micro.service.product.ProductInfo")
	proto.RegisterType((*ProductImage)(nil), "go.micro.service.product.ProductImage")
	proto.RegisterType((*ProductSize)(nil), "go.micro.service.product.ProductSize")
	proto.RegisterType((*ProductSeo)(nil), "go.micro.service.product.ProductSeo")
	proto.RegisterType((*ResponseProduct)(nil), "go.micro.service.product.ResponseProduct")
	proto.RegisterType((*RequestID)(nil), "go.micro.service.product.RequestID")
	proto.RegisterType((*Response)(nil), "go.micro.service.product.Response")
	proto.RegisterType((*RequestAll)(nil), "go.micro.service.product.RequestAll")
	proto.RegisterType((*AllProduct)(nil), "go.micro.service.product.AllProduct")
}

func init() { proto.RegisterFile("proto/product/product.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 596 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0x51, 0x6e, 0xd3, 0x40,
	0x10, 0x8d, 0x6b, 0x68, 0xe3, 0x49, 0xdb, 0xd0, 0xe5, 0xc7, 0x24, 0x20, 0xc2, 0xb6, 0x40, 0xe0,
	0xc3, 0x45, 0xe5, 0x04, 0xa1, 0x01, 0x11, 0x55, 0x42, 0x95, 0x4b, 0xe1, 0x07, 0x61, 0x82, 0x77,
	0x1a, 0xad, 0xe2, 0x78, 0x8d, 0xd7, 0x01, 0xa5, 0xa7, 0xe1, 0x32, 0x5c, 0x86, 0x53, 0xa0, 0xdd,
	0xec, 0xda, 0x56, 0xa1, 0x4d, 0xf8, 0x8a, 0x67, 0xe6, 0xf9, 0xcd, 0x9b, 0x37, 0xa3, 0x18, 0xba,
	0x59, 0x2e, 0x0a, 0x71, 0x98, 0xe5, 0x82, 0xcd, 0xe3, 0xc2, 0xfe, 0x06, 0x3a, 0x4b, 0xfc, 0x89,
	0x08, 0x66, 0x3c, 0xce, 0x45, 0x20, 0x31, 0xff, 0xce, 0x63, 0x0c, 0x4c, 0x9d, 0xfe, 0x72, 0xa1,
	0x75, 0xba, 0x7c, 0x1e, 0xa5, 0x17, 0x82, 0xec, 0xc2, 0x06, 0x67, 0xbe, 0xd3, 0x73, 0xfa, 0x6e,
	0xb8, 0xc1, 0x19, 0x79, 0x04, 0xdb, 0x06, 0x1a, 0xa5, 0xe3, 0x19, 0xfa, 0x1b, 0x3d, 0xa7, 0xef,
	0x85, 0x2d, 0x93, 0x7b, 0x37, 0x9e, 0x21, 0x79, 0x08, 0x36, 0x8c, 0xe4, 0x74, 0xee, 0xbb, 0x1a,
	0x01, 0x26, 0x75, 0x36, 0x9d, 0x93, 0x7d, 0xd8, 0xb1, 0x80, 0x2c, 0xe7, 0x31, 0xfa, 0xb7, 0x7a,
	0x4e, 0xdf, 0x09, 0x2d, 0xf1, 0xa9, 0xca, 0x91, 0x43, 0xb8, 0x6b, 0x41, 0x0c, 0x65, 0x9c, 0xf3,
	0xac, 0xe0, 0x22, 0xf5, 0x6f, 0x6b, 0x36, 0x62, 0x4a, 0xc3, 0xaa, 0x42, 0x82, 0xea, 0x85, 0x78,
	0x5c, 0xe0, 0x44, 0xe4, 0x8b, 0x88, 0x33, 0x7f, 0x53, 0x4b, 0xdf, 0x33, 0xa5, 0x63, 0x53, 0x19,
	0x31, 0x72, 0x52, 0xa9, 0xe0, 0xb3, 0xf1, 0x04, 0xfd, 0xad, 0x9e, 0xdb, 0x6f, 0x1d, 0x3d, 0x09,
	0xae, 0xf3, 0x26, 0xb0, 0xbe, 0x28, 0x74, 0xa9, 0x56, 0x47, 0xe4, 0x6d, 0x65, 0x8b, 0xe4, 0x97,
	0xe8, 0x37, 0x35, 0xd7, 0xe3, 0x95, 0x5c, 0x67, 0xfc, 0x12, 0x4b, 0xf7, 0x54, 0x40, 0x5e, 0xd7,
	0xdc, 0x43, 0xe1, 0x7b, 0x9a, 0xe8, 0x60, 0x35, 0x11, 0x8a, 0xca, 0x63, 0x14, 0x74, 0x01, 0xdb,
	0x75, 0xb9, 0x7f, 0xed, 0xf1, 0x01, 0x80, 0x9e, 0xba, 0xbe, 0x45, 0x4f, 0x67, 0xf4, 0x0e, 0xcb,
	0x72, 0x2c, 0x18, 0x9a, 0x15, 0x2e, 0xcb, 0xc7, 0x82, 0x21, 0xe9, 0xc2, 0x32, 0x88, 0xe6, 0x79,
	0xa2, 0xb7, 0xe7, 0x85, 0x4d, 0x9d, 0x38, 0xcf, 0x13, 0xfa, 0xb1, 0xbc, 0x20, 0x3d, 0xd0, 0xd5,
	0xce, 0x5d, 0xf0, 0x94, 0x45, 0xf5, 0xc6, 0x4d, 0x95, 0xd0, 0x7d, 0x6d, 0xb1, 0xd6, 0x56, 0x17,
	0x55, 0x57, 0xfa, 0xd3, 0x01, 0xa8, 0xc6, 0xfd, 0x27, 0x31, 0x8a, 0xa8, 0xe0, 0x45, 0x52, 0x11,
	0xa3, 0x78, 0xaf, 0x62, 0x75, 0xb7, 0xaa, 0x38, 0xc5, 0xc5, 0x0f, 0x91, 0x33, 0x69, 0xb8, 0x5b,
	0x12, 0xc5, 0x89, 0x49, 0x91, 0xa7, 0xd0, 0x56, 0x90, 0xfa, 0xb5, 0x2d, 0x47, 0xdb, 0x95, 0x28,
	0xea, 0x97, 0x76, 0x0f, 0x14, 0xef, 0x52, 0xe3, 0xa6, 0x46, 0x6c, 0x49, 0x14, 0x5a, 0xe2, 0x0b,
	0x68, 0x87, 0x28, 0x33, 0x91, 0x4a, 0x34, 0x4a, 0x95, 0x95, 0xe5, 0x9d, 0x59, 0xb9, 0x9e, 0x3d,
	0x1e, 0x46, 0x9f, 0x83, 0x17, 0xe2, 0xb7, 0x39, 0xca, 0x62, 0x34, 0x5c, 0x85, 0xbd, 0x0f, 0x4d,
	0xcb, 0x4e, 0xee, 0x80, 0x3b, 0x93, 0x13, 0x8d, 0xf1, 0x42, 0xf5, 0x48, 0xb7, 0x01, 0x0c, 0xd3,
	0x20, 0x49, 0xe8, 0x07, 0x80, 0x41, 0x92, 0x58, 0x11, 0xb5, 0xfb, 0xe4, 0xe9, 0x85, 0xf0, 0x9d,
	0x35, 0xef, 0x53, 0xfd, 0x07, 0x94, 0xf7, 0xa9, 0x82, 0xa3, 0xdf, 0x2e, 0x6c, 0x59, 0xd6, 0x2f,
	0x00, 0x03, 0xc6, 0x6c, 0xb4, 0x1e, 0x5b, 0xe7, 0xd9, 0xf5, 0xb0, 0x2b, 0xd6, 0xd1, 0x06, 0x89,
	0xa0, 0xfd, 0x86, 0xa7, 0xb6, 0xc5, 0xab, 0xc5, 0x68, 0x48, 0xf6, 0x6f, 0x7a, 0xdf, 0x18, 0xd9,
	0x59, 0x4f, 0x0b, 0x6d, 0x90, 0x4f, 0xb0, 0x73, 0x9e, 0xb1, 0x71, 0x81, 0xff, 0x39, 0x05, 0x5d,
	0x3d, 0x05, 0x6d, 0x90, 0xcf, 0xb0, 0x37, 0xc4, 0x04, 0x4b, 0xf6, 0xf5, 0x07, 0x58, 0x97, 0x7f,
	0x57, 0xd9, 0x53, 0x5b, 0xf4, 0xc1, 0x4a, 0xf2, 0x41, 0x92, 0x74, 0x6e, 0x40, 0x55, 0x5c, 0xb4,
	0xf1, 0x75, 0x53, 0x7f, 0x2e, 0x5e, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0x92, 0x43, 0x2a, 0xa7,
	0x4d, 0x06, 0x00, 0x00,
}
