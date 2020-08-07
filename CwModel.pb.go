// Code generated by protoc-gen-go. DO NOT EDIT.
// source: CwModel.proto

package jwParser

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Monetize int32

const (
	Monetize_Free         Monetize = 0
	Monetize_Paid         Monetize = 1
	Monetize_Subscription Monetize = 2
	Monetize_Rent         Monetize = 3
)

var Monetize_name = map[int32]string{
	0: "Free",
	1: "Paid",
	2: "Subscription",
	3: "Rent",
}

var Monetize_value = map[string]int32{
	"Free":         0,
	"Paid":         1,
	"Subscription": 2,
	"Rent":         3,
}

func (x Monetize) String() string {
	return proto.EnumName(Monetize_name, int32(x))
}

func (Monetize) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9b255d0adab92fe1, []int{0}
}

type TileType int32

const (
	TileType_ImageTile     TileType = 0
	TileType_VideoTile     TileType = 1
	TileType_FeatureTile   TileType = 2
	TileType_AdvertiseTile TileType = 3
	TileType_CarouselTile  TileType = 4
)

var TileType_name = map[int32]string{
	0: "ImageTile",
	1: "VideoTile",
	2: "FeatureTile",
	3: "AdvertiseTile",
	4: "CarouselTile",
}

var TileType_value = map[string]int32{
	"ImageTile":     0,
	"VideoTile":     1,
	"FeatureTile":   2,
	"AdvertiseTile": 3,
	"CarouselTile":  4,
}

func (x TileType) String() string {
	return proto.EnumName(TileType_name, int32(x))
}

func (TileType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9b255d0adab92fe1, []int{1}
}

type ContentAvailable struct {
	Monetize             Monetize `protobuf:"varint,6,opt,name=monetize,proto3,enum=jwParser.Monetize" json:"monetize,omitempty"`
	TargetId             string   `protobuf:"bytes,1,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`
	Source               string   `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	Package              string   `protobuf:"bytes,3,opt,name=package,proto3" json:"package,omitempty"`
	Type                 string   `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Target               string   `protobuf:"bytes,5,opt,name=target,proto3" json:"target,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContentAvailable) Reset()         { *m = ContentAvailable{} }
func (m *ContentAvailable) String() string { return proto.CompactTextString(m) }
func (*ContentAvailable) ProtoMessage()    {}
func (*ContentAvailable) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b255d0adab92fe1, []int{0}
}

func (m *ContentAvailable) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContentAvailable.Unmarshal(m, b)
}
func (m *ContentAvailable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContentAvailable.Marshal(b, m, deterministic)
}
func (m *ContentAvailable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentAvailable.Merge(m, src)
}
func (m *ContentAvailable) XXX_Size() int {
	return xxx_messageInfo_ContentAvailable.Size(m)
}
func (m *ContentAvailable) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentAvailable.DiscardUnknown(m)
}

var xxx_messageInfo_ContentAvailable proto.InternalMessageInfo

func (m *ContentAvailable) GetMonetize() Monetize {
	if m != nil {
		return m.Monetize
	}
	return Monetize_Free
}

func (m *ContentAvailable) GetTargetId() string {
	if m != nil {
		return m.TargetId
	}
	return ""
}

func (m *ContentAvailable) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *ContentAvailable) GetPackage() string {
	if m != nil {
		return m.Package
	}
	return ""
}

func (m *ContentAvailable) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ContentAvailable) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

type Media struct {
	Landscape            []string `protobuf:"bytes,5,rep,name=landscape,proto3" json:"landscape,omitempty"`
	Portrait             []string `protobuf:"bytes,1,rep,name=portrait,proto3" json:"portrait,omitempty"`
	Backdrop             []string `protobuf:"bytes,2,rep,name=backdrop,proto3" json:"backdrop,omitempty"`
	Banner               []string `protobuf:"bytes,3,rep,name=banner,proto3" json:"banner,omitempty"`
	Video                []string `protobuf:"bytes,4,rep,name=video,proto3" json:"video,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Media) Reset()         { *m = Media{} }
func (m *Media) String() string { return proto.CompactTextString(m) }
func (*Media) ProtoMessage()    {}
func (*Media) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b255d0adab92fe1, []int{1}
}

func (m *Media) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Media.Unmarshal(m, b)
}
func (m *Media) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Media.Marshal(b, m, deterministic)
}
func (m *Media) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Media.Merge(m, src)
}
func (m *Media) XXX_Size() int {
	return xxx_messageInfo_Media.Size(m)
}
func (m *Media) XXX_DiscardUnknown() {
	xxx_messageInfo_Media.DiscardUnknown(m)
}

var xxx_messageInfo_Media proto.InternalMessageInfo

func (m *Media) GetLandscape() []string {
	if m != nil {
		return m.Landscape
	}
	return nil
}

func (m *Media) GetPortrait() []string {
	if m != nil {
		return m.Portrait
	}
	return nil
}

func (m *Media) GetBackdrop() []string {
	if m != nil {
		return m.Backdrop
	}
	return nil
}

func (m *Media) GetBanner() []string {
	if m != nil {
		return m.Banner
	}
	return nil
}

func (m *Media) GetVideo() []string {
	if m != nil {
		return m.Video
	}
	return nil
}

type Content struct {
	PublishState         bool     `protobuf:"varint,3,opt,name=publish_state,json=publishState,proto3" json:"publish_state,omitempty"`
	DetailPage           bool     `protobuf:"varint,1,opt,name=detail_page,json=detailPage,proto3" json:"detail_page,omitempty"`
	Sources              []string `protobuf:"bytes,2,rep,name=sources,proto3" json:"sources,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Content) Reset()         { *m = Content{} }
func (m *Content) String() string { return proto.CompactTextString(m) }
func (*Content) ProtoMessage()    {}
func (*Content) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b255d0adab92fe1, []int{2}
}

func (m *Content) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Content.Unmarshal(m, b)
}
func (m *Content) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Content.Marshal(b, m, deterministic)
}
func (m *Content) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Content.Merge(m, src)
}
func (m *Content) XXX_Size() int {
	return xxx_messageInfo_Content.Size(m)
}
func (m *Content) XXX_DiscardUnknown() {
	xxx_messageInfo_Content.DiscardUnknown(m)
}

var xxx_messageInfo_Content proto.InternalMessageInfo

func (m *Content) GetPublishState() bool {
	if m != nil {
		return m.PublishState
	}
	return false
}

func (m *Content) GetDetailPage() bool {
	if m != nil {
		return m.DetailPage
	}
	return false
}

func (m *Content) GetSources() []string {
	if m != nil {
		return m.Sources
	}
	return nil
}

type Metadata struct {
	Title                string   `protobuf:"bytes,20,opt,name=title,proto3" json:"title,omitempty"`
	ImdbId               string   `protobuf:"bytes,1,opt,name=imdb_id,json=imdbId,proto3" json:"imdb_id,omitempty"`
	Synopsis             string   `protobuf:"bytes,2,opt,name=synopsis,proto3" json:"synopsis,omitempty"`
	Country              []string `protobuf:"bytes,3,rep,name=country,proto3" json:"country,omitempty"`
	Runtime              string   `protobuf:"bytes,4,opt,name=runtime,proto3" json:"runtime,omitempty"`
	Rating               float64  `protobuf:"fixed64,5,opt,name=rating,proto3" json:"rating,omitempty"`
	ReleaseDate          string   `protobuf:"bytes,6,opt,name=release_date,json=releaseDate,proto3" json:"release_date,omitempty"`
	Tags                 []string `protobuf:"bytes,7,rep,name=tags,proto3" json:"tags,omitempty"`
	Year                 int32    `protobuf:"varint,8,opt,name=year,proto3" json:"year,omitempty"`
	Cast                 []string `protobuf:"bytes,9,rep,name=cast,proto3" json:"cast,omitempty"`
	Directors            []string `protobuf:"bytes,10,rep,name=directors,proto3" json:"directors,omitempty"`
	Genre                []string `protobuf:"bytes,11,rep,name=genre,proto3" json:"genre,omitempty"`
	Categories           []string `protobuf:"bytes,12,rep,name=categories,proto3" json:"categories,omitempty"`
	Languages            []string `protobuf:"bytes,13,rep,name=languages,proto3" json:"languages,omitempty"`
	KidsSafe             bool     `protobuf:"varint,14,opt,name=kids_safe,json=kidsSafe,proto3" json:"kids_safe,omitempty"`
	ViewCount            float64  `protobuf:"fixed64,15,opt,name=view_count,json=viewCount,proto3" json:"view_count,omitempty"`
	Season               int32    `protobuf:"varint,16,opt,name=season,proto3" json:"season,omitempty"`
	Episode              int32    `protobuf:"varint,17,opt,name=episode,proto3" json:"episode,omitempty"`
	Part                 int32    `protobuf:"varint,18,opt,name=part,proto3" json:"part,omitempty"`
	Mood                 []int32  `protobuf:"varint,19,rep,packed,name=mood,proto3" json:"mood,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b255d0adab92fe1, []int{3}
}

func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metadata.Unmarshal(m, b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
}
func (m *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(m, src)
}
func (m *Metadata) XXX_Size() int {
	return xxx_messageInfo_Metadata.Size(m)
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Metadata) GetImdbId() string {
	if m != nil {
		return m.ImdbId
	}
	return ""
}

func (m *Metadata) GetSynopsis() string {
	if m != nil {
		return m.Synopsis
	}
	return ""
}

func (m *Metadata) GetCountry() []string {
	if m != nil {
		return m.Country
	}
	return nil
}

func (m *Metadata) GetRuntime() string {
	if m != nil {
		return m.Runtime
	}
	return ""
}

func (m *Metadata) GetRating() float64 {
	if m != nil {
		return m.Rating
	}
	return 0
}

func (m *Metadata) GetReleaseDate() string {
	if m != nil {
		return m.ReleaseDate
	}
	return ""
}

func (m *Metadata) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Metadata) GetYear() int32 {
	if m != nil {
		return m.Year
	}
	return 0
}

func (m *Metadata) GetCast() []string {
	if m != nil {
		return m.Cast
	}
	return nil
}

func (m *Metadata) GetDirectors() []string {
	if m != nil {
		return m.Directors
	}
	return nil
}

func (m *Metadata) GetGenre() []string {
	if m != nil {
		return m.Genre
	}
	return nil
}

func (m *Metadata) GetCategories() []string {
	if m != nil {
		return m.Categories
	}
	return nil
}

func (m *Metadata) GetLanguages() []string {
	if m != nil {
		return m.Languages
	}
	return nil
}

func (m *Metadata) GetKidsSafe() bool {
	if m != nil {
		return m.KidsSafe
	}
	return false
}

func (m *Metadata) GetViewCount() float64 {
	if m != nil {
		return m.ViewCount
	}
	return 0
}

func (m *Metadata) GetSeason() int32 {
	if m != nil {
		return m.Season
	}
	return 0
}

func (m *Metadata) GetEpisode() int32 {
	if m != nil {
		return m.Episode
	}
	return 0
}

func (m *Metadata) GetPart() int32 {
	if m != nil {
		return m.Part
	}
	return 0
}

func (m *Metadata) GetMood() []int32 {
	if m != nil {
		return m.Mood
	}
	return nil
}

type Optimus struct {
	RefId                string              `protobuf:"bytes,1,opt,name=ref_id,json=refId,proto3" json:"ref_id,omitempty"`
	TileType             TileType            `protobuf:"varint,2,opt,name=tile_type,json=tileType,proto3,enum=jwParser.TileType" json:"tile_type,omitempty"`
	Content              *Content            `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Metadata             *Metadata           `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
	ContentAvailable     []*ContentAvailable `protobuf:"bytes,5,rep,name=content_available,json=contentAvailable,proto3" json:"content_available,omitempty"`
	Media                *Media              `protobuf:"bytes,6,opt,name=media,proto3" json:"media,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Optimus) Reset()         { *m = Optimus{} }
func (m *Optimus) String() string { return proto.CompactTextString(m) }
func (*Optimus) ProtoMessage()    {}
func (*Optimus) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b255d0adab92fe1, []int{4}
}

func (m *Optimus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Optimus.Unmarshal(m, b)
}
func (m *Optimus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Optimus.Marshal(b, m, deterministic)
}
func (m *Optimus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Optimus.Merge(m, src)
}
func (m *Optimus) XXX_Size() int {
	return xxx_messageInfo_Optimus.Size(m)
}
func (m *Optimus) XXX_DiscardUnknown() {
	xxx_messageInfo_Optimus.DiscardUnknown(m)
}

var xxx_messageInfo_Optimus proto.InternalMessageInfo

func (m *Optimus) GetRefId() string {
	if m != nil {
		return m.RefId
	}
	return ""
}

func (m *Optimus) GetTileType() TileType {
	if m != nil {
		return m.TileType
	}
	return TileType_ImageTile
}

func (m *Optimus) GetContent() *Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *Optimus) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Optimus) GetContentAvailable() []*ContentAvailable {
	if m != nil {
		return m.ContentAvailable
	}
	return nil
}

func (m *Optimus) GetMedia() *Media {
	if m != nil {
		return m.Media
	}
	return nil
}

func init() {
	proto.RegisterEnum("jwParser.Monetize", Monetize_name, Monetize_value)
	proto.RegisterEnum("jwParser.TileType", TileType_name, TileType_value)
	proto.RegisterType((*ContentAvailable)(nil), "jwParser.ContentAvailable")
	proto.RegisterType((*Media)(nil), "jwParser.Media")
	proto.RegisterType((*Content)(nil), "jwParser.Content")
	proto.RegisterType((*Metadata)(nil), "jwParser.Metadata")
	proto.RegisterType((*Optimus)(nil), "jwParser.Optimus")
}

func init() {
	proto.RegisterFile("CwModel.proto", fileDescriptor_9b255d0adab92fe1)
}

var fileDescriptor_9b255d0adab92fe1 = []byte{
	// 780 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x54, 0xdb, 0x6e, 0x24, 0x35,
	0x10, 0xdd, 0x9e, 0x6b, 0x4f, 0x4d, 0xb2, 0xe9, 0x98, 0x05, 0xac, 0xe5, 0x36, 0x0c, 0x42, 0x8a,
	0x16, 0x29, 0x48, 0xe1, 0x8d, 0xb7, 0x55, 0xd0, 0xa2, 0x3c, 0x44, 0x44, 0x9d, 0x15, 0x8f, 0x8c,
	0x6a, 0xda, 0x95, 0xc6, 0xa4, 0xa7, 0xdd, 0xb2, 0xdd, 0x89, 0x86, 0x2f, 0x40, 0xe2, 0x1f, 0xf8,
	0x13, 0xfe, 0x0d, 0x95, 0xed, 0x9e, 0x2c, 0xbb, 0x6f, 0x75, 0xce, 0x71, 0xbb, 0xda, 0xa7, 0x2e,
	0x70, 0x7c, 0xf9, 0x78, 0x6d, 0x14, 0x35, 0xe7, 0x9d, 0x35, 0xde, 0x88, 0xfc, 0x8f, 0xc7, 0x1b,
	0xb4, 0x8e, 0xec, 0xfa, 0xdf, 0x0c, 0x8a, 0x4b, 0xd3, 0x7a, 0x6a, 0xfd, 0xeb, 0x07, 0xd4, 0x0d,
	0x6e, 0x1b, 0x12, 0xe7, 0x90, 0xef, 0x4c, 0x4b, 0x5e, 0xff, 0x49, 0x72, 0xb6, 0xca, 0xce, 0x9e,
	0x5f, 0x88, 0xf3, 0xe1, 0x8b, 0xf3, 0xeb, 0xa4, 0x94, 0x87, 0x33, 0xe2, 0x33, 0x58, 0x78, 0xb4,
	0x35, 0xf9, 0x8d, 0x56, 0x32, 0x5b, 0x65, 0x67, 0x8b, 0x32, 0x8f, 0xc4, 0x95, 0x12, 0x9f, 0xc0,
	0xcc, 0x99, 0xde, 0x56, 0x24, 0x47, 0x41, 0x49, 0x48, 0x48, 0x98, 0x77, 0x58, 0xdd, 0x63, 0x4d,
	0x72, 0x1c, 0x84, 0x01, 0x0a, 0x01, 0x13, 0xbf, 0xef, 0x48, 0x4e, 0x02, 0x1d, 0x62, 0xbe, 0x25,
	0xde, 0x28, 0xa7, 0xf1, 0x96, 0x88, 0xd6, 0x7f, 0x67, 0x30, 0xbd, 0x26, 0xa5, 0x51, 0x7c, 0x0e,
	0x8b, 0x06, 0x5b, 0xe5, 0x2a, 0xec, 0x48, 0x4e, 0x57, 0xe3, 0xb3, 0x45, 0xf9, 0x44, 0x88, 0x97,
	0x90, 0x77, 0xc6, 0x7a, 0x8b, 0xda, 0xcb, 0x2c, 0x88, 0x07, 0xcc, 0xda, 0x16, 0xab, 0x7b, 0x65,
	0x4d, 0x27, 0x47, 0x51, 0x1b, 0x30, 0xe7, 0xdd, 0x62, 0xdb, 0x92, 0x95, 0xe3, 0xa0, 0x24, 0x24,
	0x5e, 0xc0, 0xf4, 0x41, 0x2b, 0x32, 0x72, 0x12, 0xe8, 0x08, 0xd6, 0x1a, 0xe6, 0xc9, 0x4c, 0xf1,
	0x0d, 0x1c, 0x77, 0xfd, 0xb6, 0xd1, 0xee, 0xf7, 0x8d, 0xf3, 0xe8, 0xe3, 0x23, 0xf3, 0xf2, 0x28,
	0x91, 0xb7, 0xcc, 0x89, 0xaf, 0x60, 0xa9, 0xc8, 0xa3, 0x6e, 0x36, 0x1d, 0xfb, 0x90, 0x85, 0x23,
	0x10, 0xa9, 0x1b, 0xb6, 0x42, 0xc2, 0x3c, 0xda, 0xe5, 0xd2, 0x9f, 0x0d, 0x70, 0xfd, 0xd7, 0x04,
	0xf2, 0x6b, 0xf2, 0xa8, 0xd0, 0x23, 0xff, 0x8d, 0xd7, 0xbe, 0x21, 0xf9, 0x22, 0x98, 0x13, 0x81,
	0xf8, 0x14, 0xe6, 0x7a, 0xa7, 0xb6, 0x4f, 0x45, 0x99, 0x31, 0xbc, 0x52, 0xfc, 0x60, 0xb7, 0x6f,
	0x4d, 0xe7, 0xb4, 0x4b, 0x45, 0x39, 0x60, 0xce, 0x58, 0x99, 0xbe, 0xf5, 0x76, 0x9f, 0x5e, 0x3c,
	0x40, 0x56, 0x6c, 0xdf, 0x7a, 0xbd, 0x1b, 0x2a, 0x33, 0x40, 0x36, 0xc9, 0xa2, 0xd7, 0x6d, 0x1d,
	0x8a, 0x93, 0x95, 0x09, 0x89, 0xaf, 0xe1, 0xc8, 0x52, 0x43, 0xe8, 0x68, 0xa3, 0xd8, 0x82, 0x59,
	0xf8, 0x6c, 0x99, 0xb8, 0x9f, 0xd8, 0x01, 0xae, 0x35, 0xd6, 0x4e, 0xce, 0x43, 0xae, 0x10, 0x33,
	0xb7, 0x27, 0xb4, 0x32, 0x5f, 0x65, 0x67, 0xd3, 0x32, 0xc4, 0xcc, 0x55, 0xe8, 0xbc, 0x5c, 0xc4,
	0x73, 0x1c, 0x73, 0xc5, 0x95, 0xb6, 0x54, 0x79, 0x63, 0x9d, 0x84, 0x58, 0xf1, 0x03, 0xc1, 0x9e,
	0xd4, 0xd4, 0x5a, 0x92, 0xcb, 0x58, 0xa1, 0x00, 0xc4, 0x97, 0x00, 0x15, 0x7a, 0xaa, 0x8d, 0xd5,
	0xe4, 0xe4, 0x51, 0x90, 0xde, 0x61, 0x52, 0x17, 0xd5, 0x3d, 0xd6, 0xe4, 0xe4, 0xf1, 0xa1, 0x8b,
	0x22, 0xc1, 0x8d, 0x7e, 0xaf, 0x95, 0xdb, 0x38, 0xbc, 0x23, 0xf9, 0x3c, 0x54, 0x2b, 0x67, 0xe2,
	0x16, 0xef, 0x48, 0x7c, 0x01, 0xf0, 0xa0, 0xe9, 0x71, 0x13, 0xfc, 0x92, 0x27, 0xc1, 0x89, 0x05,
	0x33, 0x97, 0x4c, 0x84, 0x39, 0x20, 0x74, 0xa6, 0x95, 0x45, 0x78, 0x57, 0x42, 0x6c, 0x2b, 0x75,
	0xda, 0x19, 0x45, 0xf2, 0x34, 0x08, 0x03, 0xe4, 0x37, 0x77, 0x68, 0xbd, 0x14, 0xd1, 0x07, 0x8e,
	0x99, 0xdb, 0x19, 0xa3, 0xe4, 0x47, 0xab, 0x31, 0x73, 0x1c, 0xaf, 0xff, 0x19, 0xc1, 0xfc, 0x97,
	0xce, 0xeb, 0x5d, 0xef, 0xc4, 0xc7, 0x30, 0xb3, 0x74, 0xf7, 0x54, 0xf2, 0xa9, 0xa5, 0xbb, 0x2b,
	0x25, 0xbe, 0x87, 0x85, 0xd7, 0x0d, 0x6d, 0xc2, 0x5c, 0x8d, 0xde, 0x1f, 0xe9, 0xb7, 0xba, 0xa1,
	0xb7, 0xfb, 0x8e, 0xca, 0xdc, 0xa7, 0x48, 0x7c, 0xc7, 0x6d, 0x10, 0x3a, 0x39, 0x34, 0xee, 0xf2,
	0xe2, 0xf4, 0xe9, 0x78, 0x6a, 0xf1, 0x72, 0x38, 0x11, 0xf6, 0x45, 0x6a, 0xc5, 0xd0, 0x1a, 0xcb,
	0xff, 0xed, 0x8b, 0xa4, 0x94, 0x87, 0x33, 0xe2, 0x67, 0x38, 0x4d, 0x9f, 0x6e, 0x70, 0x58, 0x3a,
	0x61, 0x64, 0x97, 0x17, 0x2f, 0x3f, 0x48, 0x73, 0x58, 0x4b, 0x65, 0x51, 0xbd, 0xbf, 0xa8, 0xbe,
	0x85, 0xe9, 0x8e, 0x87, 0x3f, 0x74, 0xd6, 0xf2, 0xe2, 0xe4, 0xdd, 0xac, 0x4a, 0x63, 0x19, 0xd5,
	0x57, 0x3f, 0x42, 0x3e, 0x6c, 0x2d, 0x91, 0xc3, 0xe4, 0x8d, 0x25, 0x2a, 0x9e, 0x71, 0x74, 0x83,
	0x5a, 0x15, 0x99, 0x28, 0xe0, 0xe8, 0xb6, 0xdf, 0xba, 0xca, 0xea, 0xce, 0x6b, 0xd3, 0x16, 0x23,
	0xd6, 0x4a, 0x6a, 0x7d, 0x31, 0x7e, 0xf5, 0x1b, 0xe4, 0x83, 0x3d, 0xe2, 0x18, 0x16, 0x57, 0x3b,
	0xac, 0x89, 0x89, 0xe2, 0x19, 0xc3, 0x5f, 0x79, 0xec, 0x03, 0xcc, 0xc4, 0x09, 0x2c, 0xdf, 0x10,
	0xfa, 0xde, 0x46, 0x7d, 0x24, 0x4e, 0xe1, 0xf8, 0xb5, 0x7a, 0x20, 0xeb, 0xb5, 0x8b, 0xd4, 0x98,
	0x33, 0x5d, 0xa2, 0x35, 0xbd, 0xa3, 0x26, 0x30, 0x93, 0xed, 0x2c, 0x6c, 0xe4, 0x1f, 0xfe, 0x0b,
	0x00, 0x00, 0xff, 0xff, 0xc0, 0x19, 0x5b, 0x0c, 0xa2, 0x05, 0x00, 0x00,
}