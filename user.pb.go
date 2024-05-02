package protofile

import (
	context "context"
	grpc "google.golang.org/grpc"
)

// User структура представляет пользователя.
type User struct {
	Id    int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

// UserServiceServer описывает интерфейс для взаимодействия с сервисом пользователей.
type UserServiceServer interface {
	AddUser(context.Context, *User) (*User, error)
	GetUser(context.Context, *UserID) (*User, error)
	ListUsers(context.Context, *Empty) (*UserList, error)
}

// RegisterUserServiceServer регистрирует UserServiceServer в gRPC сервере.
func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

// Конкретные методы обработчиков для каждого gRPC метода.
func _UserService_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	// Обработка запроса на добавление пользователя
}

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	// Обработка запроса на получение пользователя
}

func _UserService_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	// Обработка запроса на получение списка пользователей
}

// UserList структура представляет список пользователей.
type UserList struct {
	Users []*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

// UserID структура представляет идентификатор пользователя.
type UserID struct {
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

// Empty структура представляет пустой запрос.
type Empty struct{}


