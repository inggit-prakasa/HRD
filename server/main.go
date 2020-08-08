package main

import (
	"context"
	"log"
	"net"

	employee "github.com/inggit_prakasa/HRD/employee"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type server struct {
	employee.UnimplementedManageEmpServer
}

type Cemployee struct {
	id                 int64
	absen              int64
	cuti               int64
	nama               string
	username           string
	password           string
	gajiPokok          float32
	totalGaji          float32
	tunjanganMakan     float32
	tunjanganTransport float32
	status             string
	role               string
}

var employeeList = []Cemployee{
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

func (s *server) CreateEmployee(ctx context.Context, in *employee.Employee) (*employee.Employee, error) {
	newEmployee := Cemployee{
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

	employeeList = append(employeeList, newEmployee)
	return &employee.Employee{
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
	}, nil
}

func (s *server) LaporanByUsername(ctx context.Context, input *employee.GetEmpByUsername) (*employee.LaporanEmployee, error) {

	EmployeeByUsername := &employee.LaporanEmployee{
		ID:        "string",
		Nama:      "string",
		TotalGaji: "string",
		Status:    "string",
	}
	return EmployeeByUsername, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	employee.RegisterManageEmpServer(s,&server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}


