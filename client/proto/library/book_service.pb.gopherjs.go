// Code generated by protoc-gen-gopherjs. DO NOT EDIT.
// source: proto/library/book_service.proto

/*
	Package library is a generated protocol buffer package.

	Package library exposes a list of books
	over a gRPC API.

	It is generated from these files:
		proto/library/book_service.proto

	It has these top-level messages:
		Publisher
		Book
		GetBookRequest
		QueryBooksRequest
*/
package library

import js "github.com/gopherjs/gopherjs/js"
import grpcweb "github.com/johanbrandhorst/gopherjs-improbable-grpc-web"

import (
	context "context"

	grpcweb1 "github.com/johanbrandhorst/gopherjs-improbable-grpc-web"
)

// BookType describes the different types
// a book in the library can be.
type BookType int

const (
	// Hardcover is a book with a hard back.
	BookType_HARDCOVER BookType = 0
	// Paperback is a book with a soft back.
	BookType_PAPERBACK BookType = 1
	// Audiobook is an audio recording of the book.
	BookType_AUDIOBOOK BookType = 2
)

var BookType_name = map[int]string{
	0: "HARDCOVER",
	1: "PAPERBACK",
	2: "AUDIOBOOK",
}
var BookType_value = map[string]int{
	"HARDCOVER": 0,
	"PAPERBACK": 1,
	"AUDIOBOOK": 2,
}

func (x BookType) String() string {
	return BookType_name[int(x)]
}

// Publisher describes a Book Publisher.
type Publisher struct {
	*js.Object
}

// NewPublisher creates a new Publisher.
// Name is the name of the Publisher.
func NewPublisher(name string) *Publisher {
	m := &Publisher{
		Object: js.Global.Get("proto").Get("library").Get("Publisher").New([]interface{}{
			name,
		}),
	}

	return m
}

// GetName gets the Name of the Publisher.
// Name is the name of the Publisher.
func (m *Publisher) GetName() string {
	return m.Call("getName").String()
}

// SetName sets the Name of the Publisher.
// Name is the name of the Publisher.
func (m *Publisher) SetName(v string) {
	m.Call("setName", v)
}

func (m *Publisher) serialize() (rawBytes []byte, err error) {
	return grpcweb.Serialize(m)
}

func deserializePublisher(rawBytes []byte) (*Publisher, error) {
	obj, err := grpcweb.Deserialize(js.Global.Get("proto").Get("library").Get("Publisher"), rawBytes)
	if err != nil {
		return nil, err
	}

	return &Publisher{
		Object: obj,
	}, nil
}

// Book represents a book in the library.
type Book struct {
	*js.Object
}

// NewBook creates a new Book.
// Isbn is the ISBN number of the book.
// Title is the title of the book.
// Author is the author of the book.
// BookType is the type of the book.
func NewBook(isbn int64, title string, author string, bookType BookType) *Book {
	m := &Book{
		Object: js.Global.Get("proto").Get("library").Get("Book").New([]interface{}{
			isbn,
			title,
			author,
			bookType,
			js.Undefined,
			js.Undefined,
		}),
	}

	return m
}

// PublishingMethod is the publishing method
// used for this Book.
//
// Types that are valid to be assigned to PublishingMethod:
//	*Book_SelfPublished
//	*Book_Publisher
type isBook_PublishingMethod interface {
	isBook_PublishingMethod()
}

type Book_SelfPublished struct {
	// SelfPublished means this book was
	// self published.
	SelfPublished bool
}
type Book_Publisher struct {
	// Publisher means this book was published
	// through a Publisher.
	Publisher *Publisher
}

func (*Book_SelfPublished) isBook_PublishingMethod() {}
func (*Book_Publisher) isBook_PublishingMethod()     {}

func (m *Book) GetPublishingMethod() (x isBook_PublishingMethod) {
	switch m.Call("getPublishingMethodCase").Int() {
	case 5:
		x = &Book_SelfPublished{
			SelfPublished: m.GetSelfPublished(),
		}
	case 6:
		x = &Book_Publisher{
			Publisher: m.GetPublisher(),
		}
	}
	return x
}

// GetIsbn gets the Isbn of the Book.
// Isbn is the ISBN number of the book.
func (m *Book) GetIsbn() int64 {
	return m.Call("getIsbn").Int64()
}

// SetIsbn sets the Isbn of the Book.
// Isbn is the ISBN number of the book.
func (m *Book) SetIsbn(v int64) {
	m.Call("setIsbn", v)
}

// GetTitle gets the Title of the Book.
// Title is the title of the book.
func (m *Book) GetTitle() string {
	return m.Call("getTitle").String()
}

// SetTitle sets the Title of the Book.
// Title is the title of the book.
func (m *Book) SetTitle(v string) {
	m.Call("setTitle", v)
}

// GetAuthor gets the Author of the Book.
// Author is the author of the book.
func (m *Book) GetAuthor() string {
	return m.Call("getAuthor").String()
}

// SetAuthor sets the Author of the Book.
// Author is the author of the book.
func (m *Book) SetAuthor(v string) {
	m.Call("setAuthor", v)
}

// GetBookType gets the BookType of the Book.
// BookType is the type of the book.
func (m *Book) GetBookType() BookType {
	return BookType(m.Call("getBookType").Int())
}

// SetBookType sets the BookType of the Book.
// BookType is the type of the book.
func (m *Book) SetBookType(v BookType) {
	m.Call("setBookType", v)
}

// GetSelfPublished gets the SelfPublished of the Book.
// SelfPublished means this book was
// self published.
func (m *Book) GetSelfPublished() bool {
	return m.Call("getSelfPublished").Bool()
}

// SetSelfPublished sets the SelfPublished of the Book.
// SelfPublished means this book was
// self published.
func (m *Book) SetSelfPublished(v bool) {
	m.Call("setSelfPublished", v)
}

// GetPublisher gets the Publisher of the Book.
// Publisher means this book was published
// through a Publisher.
func (m *Book) GetPublisher() *Publisher {
	return &Publisher{Object: m.Call("getPublisher")}
}

// SetPublisher sets the Publisher of the Book.
// Publisher means this book was published
// through a Publisher.
func (m *Book) SetPublisher(v *Publisher) {
	m.Call("setPublisher", v)
}

func (m *Book) serialize() (rawBytes []byte, err error) {
	return grpcweb.Serialize(m)
}

func deserializeBook(rawBytes []byte) (*Book, error) {
	obj, err := grpcweb.Deserialize(js.Global.Get("proto").Get("library").Get("Book"), rawBytes)
	if err != nil {
		return nil, err
	}

	return &Book{
		Object: obj,
	}, nil
}

// GetBookRequest is the input to the GetBook method.
type GetBookRequest struct {
	*js.Object
}

// NewGetBookRequest creates a new GetBookRequest.
// Isbn is the ISBN with which
// to match against the ISBN of a book in the library.
func NewGetBookRequest(isbn int64) *GetBookRequest {
	m := &GetBookRequest{
		Object: js.Global.Get("proto").Get("library").Get("GetBookRequest").New([]interface{}{
			isbn,
		}),
	}

	return m
}

// GetIsbn gets the Isbn of the GetBookRequest.
// Isbn is the ISBN with which
// to match against the ISBN of a book in the library.
func (m *GetBookRequest) GetIsbn() int64 {
	return m.Call("getIsbn").Int64()
}

// SetIsbn sets the Isbn of the GetBookRequest.
// Isbn is the ISBN with which
// to match against the ISBN of a book in the library.
func (m *GetBookRequest) SetIsbn(v int64) {
	m.Call("setIsbn", v)
}

func (m *GetBookRequest) serialize() (rawBytes []byte, err error) {
	return grpcweb.Serialize(m)
}

func deserializeGetBookRequest(rawBytes []byte) (*GetBookRequest, error) {
	obj, err := grpcweb.Deserialize(js.Global.Get("proto").Get("library").Get("GetBookRequest"), rawBytes)
	if err != nil {
		return nil, err
	}

	return &GetBookRequest{
		Object: obj,
	}, nil
}

// QueryBooksRequest is the input to the QueryBooks method.
type QueryBooksRequest struct {
	*js.Object
}

// NewQueryBooksRequest creates a new QueryBooksRequest.
// AuthorPrefix is the prefix with which
// to match against the author of a book in the library.
func NewQueryBooksRequest(authorPrefix string) *QueryBooksRequest {
	m := &QueryBooksRequest{
		Object: js.Global.Get("proto").Get("library").Get("QueryBooksRequest").New([]interface{}{
			authorPrefix,
		}),
	}

	return m
}

// GetAuthorPrefix gets the AuthorPrefix of the QueryBooksRequest.
// AuthorPrefix is the prefix with which
// to match against the author of a book in the library.
func (m *QueryBooksRequest) GetAuthorPrefix() string {
	return m.Call("getAuthorPrefix").String()
}

// SetAuthorPrefix sets the AuthorPrefix of the QueryBooksRequest.
// AuthorPrefix is the prefix with which
// to match against the author of a book in the library.
func (m *QueryBooksRequest) SetAuthorPrefix(v string) {
	m.Call("setAuthorPrefix", v)
}

func (m *QueryBooksRequest) serialize() (rawBytes []byte, err error) {
	return grpcweb.Serialize(m)
}

func deserializeQueryBooksRequest(rawBytes []byte) (*QueryBooksRequest, error) {
	obj, err := grpcweb.Deserialize(js.Global.Get("proto").Get("library").Get("QueryBooksRequest"), rawBytes)
	if err != nil {
		return nil, err
	}

	return &QueryBooksRequest{
		Object: obj,
	}, nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpcweb1.Client

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpcweb package it is being compiled against.
const _ = grpcweb1.GrpcWebPackageIsVersion1

// Client API for BookService service

type BookServiceClient interface {
	// GetBook returns a Book from the library
	// that matches the ISBN provided, if found.
	// Otherwise it returns a NotFound error.
	GetBook(ctx context.Context, in *GetBookRequest, opts ...grpcweb1.CallOption) (*Book, error)
	// QueryBooks returns all Books whos author
	// matches the author prefix provided, as a stream
	// of Books.
	QueryBooks(ctx context.Context, in *QueryBooksRequest, opts ...grpcweb1.CallOption) (BookService_QueryBooksClient, error)
}

type bookServiceClient struct {
	client *grpcweb1.Client
}

func NewBookServiceClient(hostname string, opts ...grpcweb.DialOption) BookServiceClient {
	return &bookServiceClient{
		client: grpcweb1.NewClient(hostname, "library.BookService", opts...),
	}
}

func (c *bookServiceClient) GetBook(ctx context.Context, in *GetBookRequest, opts ...grpcweb1.CallOption) (*Book, error) {
	req, err := in.serialize()
	if err != nil {
		return nil, err
	}
	resp, err := c.client.RPCCall(ctx, "GetBook", req, opts...)
	if err != nil {
		return nil, err
	}
	return deserializeBook(resp)
}

func (c *bookServiceClient) QueryBooks(ctx context.Context, in *QueryBooksRequest, opts ...grpcweb1.CallOption) (BookService_QueryBooksClient, error) {
	req, err := in.serialize()
	if err != nil {
		return nil, err
	}
	srv, err := c.client.Stream(ctx, "QueryBooks", req, opts...)
	if err != nil {
		return nil, err
	}
	return &bookServiceQueryBooksClient{
		stream: srv,
	}, nil
}

type BookService_QueryBooksClient interface {
	Recv() (*Book, error)
}

type bookServiceQueryBooksClient struct {
	stream *grpcweb1.StreamClient
}

func (x *bookServiceQueryBooksClient) Recv() (*Book, error) {
	resp, err := x.stream.Recv()
	if err != nil {
		return nil, err
	}
	return deserializeBook(resp)
}