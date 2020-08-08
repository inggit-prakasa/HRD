package main

import (
	"context"
	hrd "github.com/inggit_prakasa/HRD/HRD/employee"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"
)

type server struct {
	hrd.UnimplementedManageEmpServer
}

type employee struct {
	id int64
	absen int64
	cuti int64
	nama string
	username string
	password string
	gajiPokok float32
	totalGaji float32
	tunjanganMakan float32
	tunjanganTransport float32
	status string
	role string
}

var employeeList = []employee {
	{
		id:                 1,
		absen:              20,
		cuti:               2,
		nama:               "Saepul",
		username:           "staff",
		password:           "1234",
		gajiPokok:          3000000,
		totalGaji:          0,
		tunjanganMakan:     400000,
		tunjanganTransport: 500000,
		status:             "MASUK",
		role:               "STAFF",
	},
}

func (s *server) CreateEmployee(ctx context.Context, in *hrd.Employee) (*hrd.Employee, error) {
	newEmployee := employee{
		id:                 in.Id,
		absen:              in.Absen,
		cuti:               in.Cuti,
		nama:               in.Nama,
		username:           in.Username,
		password:           in.Password,
		gajiPokok:          in.GajiPokok,
		totalGaji:          in.TotalGaji,
		tunjanganMakan:     in.TunjanganMakan,
		tunjanganTransport: in.TunjanganTransport,
		status:             in.Status,
		role:               in.Role,
	}

	employeeList = append(employeeList,newEmployee)
	return &hrd.Employee{
		Id:                 in.Id,
		Absen:              in.Absen,
		Cuti:               in.Cuti,
		Nama:               in.Nama,
		Username:           in.Username,
		Password:           in.Password,
		GajiPokok:          in.GajiPokok,
		TotalGaji:          in.TotalGaji,
		TunjanganMakan:     in.TunjanganMakan,
		TunjanganTransport: in.TunjanganTransport,
		Status:             in.Status,
		Role:               in.Role,
	},nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	hrd.RegisterManageEmpServer(s,&server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}