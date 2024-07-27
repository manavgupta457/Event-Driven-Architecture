package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/jackc/pgx/v4"

	pb "example.com/go-usermgmt-grpc/usermgmt"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func NewUserManagementServer() *UserManagementServer {
	return &UserManagementServer{
		// user_list: &pb.UserList{},
	}
}

type UserManagementServer struct {
	conn *pgx.Conn
	pb.UnimplementedUserManagementServer
}

func (server *UserManagementServer) Run() error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserManagementServer(s, server)
	log.Printf("server listening at %v", lis.Addr())

	return s.Serve(lis)
}

func (s *UserManagementServer) CreateNewUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	createSql := `
	create table if not exists(
	id: SERIAL PRIMARY KEY,
	name text,
	age int
	);
	`
	_, err := server.conn.Exec(context.Background(), createSql)

	if err != nil {
		fmt.Printf(os.Stderr, "Table creation failed: %v\n", err)
		os.Exit(1)
	}

	created_user := &pb.User{Name: in.GetName(), Age: in.GetAge()}
	tx, err := server.conn.Begin(context.Background())
	if err != nil {
		log.Fatalf("conn.begin failed: %v", err)
	}

	_, err = tx.Exec(context.Background(), "insert into users(name, age) values ($1, $2)", created_user.Name, created_user.Age)
	if err != nil {
		log.Fatalf("tx.Exec failed: %v", err)
	}
	tx.Commit(context.Background())
	return created_user, nil
}

func (s *UserManagementServer) GetUsers(ctx context.Context, in *pb.GetUsersParams) (*pb.UserList, error) {
	var user_list *pb.UserList = &pb.UserList{}
	rows, err := server.conn.Query(context.Background(), "select * from users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		user := pb.User{}
		err = rows.Scan(&user.Id, &user.Name, &user.Age)
		user_list.Users = append(user_list.Users, &user)
	}
	return user_list, nil
}

func main() {
	var user_mgmt_server *UserManagementServer = NewUserManagementServer()
	if err := user_mgmt_server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
