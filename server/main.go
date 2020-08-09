package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

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

	idCount := int64(len(employeeList) + 1)
	//error username sama
	newEmployee := Cemployee{
		id:                 int64(len(employeeList) + 1),
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
		Id:                 idCount,
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

func (s *server) ReadEmployee(ctx context.Context, in *employee.NameId) (*employee.Employee, error) {
	flag := false
	index := 0
	for i := 0; i < len(employeeList); i++ {
		if employeeList[i].nama == in.Name {
			index = i
			flag = true
			break
		}
	}

	if flag {
		return &employee.Employee{
			Id:                 employeeList[index].id,
			Absen:              employeeList[index].absen,
			Cuti:               employeeList[index].cuti,
			Nama:               employeeList[index].nama,
			Username:           employeeList[index].username,
			Password:           employeeList[index].password,
			GajiPokok:          employeeList[index].gajiPokok,
			TotalGaji:          employeeList[index].totalGaji,
			TunjanganMakan:     employeeList[index].tunjanganMakan,
			TunjanganTransport: employeeList[index].tunjanganTransport,
			Status:             employeeList[index].status,
			Role:               employeeList[index].role,
		}, nil
	} else {
		return nil, nil
	}
}

func (s *server) UpdateEmployee(ctx context.Context, in *employee.Employee) (*employee.Employee, error) {
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

func (s *server) DeleteEmployee(ctx context.Context, in *employee.NameId) (*employee.Result, error) {
	flag := false

	for i := 0; i < len(employeeList); i++ {
		if employeeList[i].nama == in.Name {
			employeeList = append(employeeList[:i], employeeList[i+1:]...)
			flag = true
			break
		}
	}

	if flag {
		return &employee.Result{Response: "BERHASIl"}, nil
	} else {
		return &employee.Result{Response: "GAGAL"}, nil
	}
}

//FUNC ABSEN(ctx, in *user) (*employee.Employee) {
// for _,emps {
//		if username=emps.id{
//			returns ---
//		}
//}}
//

func (s *server) LaporanByUsername(ctx context.Context, input *employee.GetEmpByUsername) (*employee.Employee, error) {
	var EmployeeByUsername *employee.Employee
	flag := true
	for _, Cemployee := range employeeList {
		if strings.EqualFold(Cemployee.username, input.Username) {
			flag = false
			tunjanganMakanTemp := Cemployee.tunjanganMakan * (float32(Cemployee.absen) / 22.0)
			tunjanganTransportTemp := Cemployee.tunjanganTransport * (float32(Cemployee.absen) / 22.0)
			gajiTotalTemp := Cemployee.gajiPokok + tunjanganMakanTemp + tunjanganTransportTemp
			EmployeeByUsername = &employee.Employee{
				Id:                 Cemployee.id,
				Absen:              Cemployee.absen,
				Cuti:               Cemployee.cuti,
				Nama:               Cemployee.nama,
				Username:           Cemployee.username,
				Password:           Cemployee.password,
				GajiPokok:          Cemployee.gajiPokok,
				TotalGaji:          gajiTotalTemp,
				TunjanganMakan:     tunjanganMakanTemp,
				TunjanganTransport: tunjanganTransportTemp,
				Status:             Cemployee.status,
				Role:               Cemployee.role,
			}
			writeStringToFile(EmployeeByUsername)
			break

		}
	}
	if flag {
		return EmployeeByUsername, nil
	} else {
		return EmployeeByUsername, nil
	}
}

func main() {
	lis, err := net.Listen("tcp", port)
	log.Println("Serv...")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	employee.RegisterManageEmpServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func writeStringToFile(emp *employee.Employee) {
	var str string
	str = "============== Laporan" + emp.Nama + "================" + "\n" +
		"Id  \t \t :" + strconv.FormatInt(emp.Id, 10) + "\n" +
		"Absen \t \t:" + strconv.FormatInt(emp.Absen, 10) + "\n" +
		"Cuti  \t \t:" + strconv.FormatInt(emp.Absen, 10) + "\n" +
		"Nama  \t \t:" + emp.Nama + "\n" +
		"Username  \t \t:" + emp.Username + "\n" +
		"Gaji Pokok  \t \t:" + fmt.Sprintf("%.3f", emp.GajiPokok) + "\n" +
		"Gaji Total  \t \t:" + fmt.Sprintf("%.3f", emp.TotalGaji) + "\n" +
		"Tunjangan Makan  \t :" + fmt.Sprintf("%.3f", emp.TunjanganMakan) + "\n" +
		"Tunjangan Transport  :" + fmt.Sprintf("%.3f", emp.TunjanganTransport) + "\n" +
		"Status   \t:" + emp.Status + "\n" +
		"Role  \t \t:" + emp.Role + "\n" + ""

	f, err := os.Create("laporan/" + emp.Nama + ".pdf")
	if err != nil {
		fmt.Println(err)
		return
	}

	l, err := f.WriteString(str)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
