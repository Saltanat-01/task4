package main

import (
	"context"
	"log"

	pb "asg4/protofile"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Не удалось подключиться: %v", err)
	}
	defer conn.Close()
	client := pb.NewUserServiceClient(conn)

	// Добавление пользователя с измененными именем и электронным адресом
	addUserResponse, err := client.AddUser(context.Background(), &pb.User{
		Id:    1,
		Name:  "Боб",
		Email: "bob@example.com",
	})
	if err != nil {
		log.Fatalf("Не удалось добавить пользователя: %v", err)
	}
	log.Printf("Пользователь добавлен с ID: %d", addUserResponse.GetId())

	// Получение пользователя по ID
	getUserResponse, err := client.GetUser(context.Background(), &pb.UserId{Id: 1})
	if err != nil {
		log.Fatalf("Не удалось получить пользователя: %v", err)
	}
	log.Printf("Получен пользователь: ID: %d, Имя: %s, Email: %s", getUserResponse.GetId(), getUserResponse.GetName(), getUserResponse.GetEmail())

	// Получение списка пользователей
	listUsersResponse, err := client.ListUsers(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatalf("Не удалось получить список пользователей: %v", err)
	}
	log.Println("Список пользователей:")
	for _, user := range listUsersResponse.GetUsers() {
		log.Printf("ID: %d, Имя: %s, Email: %s", user.GetId(), user.GetName(), user.GetEmail())
	}
}