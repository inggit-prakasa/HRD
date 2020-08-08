package main

import (
	"context"
	"log"
	"net"

	"HRD/employee"

	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type HRD struct{}

func (HRD) LaporanByUsername(ctx context.Context, input *employee.GetEmpByUsername) (*employee.LaporanEmployee, error) {

	EmployeeByUsername := &employee.LaporanEmployee{
		ID:        "string",
		Nama:      "string",
		TotalGaji: "string",
		Status:    "string",
	}
	return EmployeeByUsername, nil
}

func main() {
	s := grpc.NewServer()
	log.Println("Server Running --> PORT", port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failedd to listen: %v", err)
	}

	var hrd HRD
	employee.RegisterHRDServer(s, hrd)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	log.Print(s.Serve(lis))
}
