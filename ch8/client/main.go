package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	pb "ch8proto/todo/v2"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	//"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/encoding/gzip"
	"google.golang.org/grpc/credentials"

	//"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func addTask(c pb.TodoServiceClient, description string, dueDate time.Time) uint64 {
	ctx := context.Background()
	//ctx = metadata.AppendToOutgoingContext(ctx, "auth_token", "authd")
	
	req := &pb.AddTaskRequest{
		Description: description,
		DueDate: timestamppb.New(dueDate),
	}
	res, err := c.AddTask(ctx, req)
	if err != nil {
		if s, ok := status.FromError(err); ok {
			switch s.Code() {
			case codes.InvalidArgument, codes.Internal:
				log.Fatalf("%s: %s", s.Code(), s.Message())
	
			default:
				log.Fatal(s)
			}
		}else {
			panic(err)
		}
	}

	fmt.Printf("added task: %d\n", res.Id)
	return res.Id

}

func printTasks(c pb.TodoServiceClient) {
	req := &pb.ListTasksRequest{}
	//ctx, cancel := context.WithCancel(context.Background())
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	//ctx = metadata.AppendToOutgoingContext(ctx, "auth_token", "authd")
	defer cancel()
	stream, err := c.ListTasks(ctx, req)
	if err != nil {
		log.Fatalf("unexpected error: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}else if err != nil {
			log.Fatalf("unexpected error: %v", err)
		}
		if res.Overdue {
			cancel()
		}
		fmt.Println(res.Task.String(), "overdue: ", res.Overdue)
	}
}

func updateTasks(c pb.TodoServiceClient, reqs ...*pb.UpdateTasksRequest) {
	ctx := context.Background()
	//ctx = metadata.AppendToOutgoingContext(ctx, "auth_token", "authd")
	stream, err := c.UpdateTasks(ctx)
	
	if err != nil {
		log.Fatalf("unexpected error: %v", err)
	}

	for _, req := range reqs {
		err := stream.Send(req)
		if err != nil {
			log.Fatalf("unexpected error: %v", err)
		}

		if req.Task != nil {
			fmt.Printf("updated task with id: %d\n", req.Task.Id)
		}
	}
	if _, err = stream.CloseAndRecv(); err != nil {
		log.Fatalf("unexpected error: %v", err)
	}
}

func deleteTasks(c pb.TodoServiceClient, reqs ...*pb.DeleteTasksRequest) {
	ctx := context.Background()
	stream , err := c.DeleteTasks(ctx)
	if err != nil {
		log.Fatalf("unexpected error: %v", err)
	}
	waitc := make(chan struct{})

	go func() {
		for {
			_, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				break
			}
			if err != nil {
				log.Fatalf("unexpected error: %v", err)
			}
			log.Println("deleted tasks")
		}
	}()

	for _, req := range reqs {
		if err := stream.Send(req); err != nil {
			return
		}
	}
	if err := stream.CloseSend(); err != nil {
		return
	}
	<-waitc
}



func main() {
	args := os.Args[1:]
	
	if len(args) == 0 {
		log.Fatalln("usage: client [IP_ADDR]")
	}

	creds, err := credentials.NewClientTLSFromFile("./../certs/ca_cert.pem", "x.test.example.com")
	if err != nil {
		log.Fatalf("failed to create credential: %v", err)
	}

	addr := args[0]
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
		//grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(unaryAuthInterceptor),
		grpc.WithStreamInterceptor(streamAuthInterceptor),
		grpc.WithDefaultCallOptions(grpc.UseCompressor(gzip.Name)),
		grpc.WithDefaultServiceConfig(`{"loadBalancingConfig":[{"round_robin":{}}]}`),
	}

	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	c := pb.NewTodoServiceClient(conn)
	fmt.Println("------ADD-----")
	dueDate := time.Now().Add(5 * time.Second)
	id1 := addTask(c, "This is a task", dueDate)
	id2 := addTask(c, "This is another task", dueDate)
	id3 := addTask(c, "And yet another task", dueDate)
	//addTask(c, "", dueDate)
	fmt.Println("--------------")

	fmt.Println("-----LIST-----")
	printTasks(c)
	fmt.Println("-------------")

	fmt.Println("----UPDATE-----")
	updateTasks(c, []*pb.UpdateTasksRequest{
		{Task: &pb.Task{Id: id1, Description: "A better name for the task"}},
		{Task: &pb.Task{Id: id2, DueDate: timestamppb.New(dueDate.Add(5 * time.Hour))}},
		{Task: &pb.Task{Id: id3, Description: "Better solution", Done: true}},
	}...)
	printTasks(c)
	fmt.Println("-------------")

	fmt.Println("----DELETE---")
	deleteTasks(c, []*pb.DeleteTasksRequest{
		{Id: id1},
		{Id: id2},
		{Id: id3},
	}...)
	fmt.Println("-------------")


	defer func(conn *grpc.ClientConn) {
		if err := conn.Close(); err != nil {
			log.Fatalf("unexpected error: %v", err)
		}
	}(conn)



}