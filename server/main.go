package main

import (
	"context"
	"fmt"
	employee "github.com/inggit_prakasa/HRD/employee"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"strconv"
)

const (
	port = ":8080"
)

type server struct {
	employee.UnimplementedManageEmpServer
}

var employeeList = []*employee.Employee{
	{
		Id:                 1,
		Absen:              20,
		Cuti:               2,
		Nama:               "Saepul",
		Username:           "staff",
		Password:           "1234",
		GajiPokok:          3000000,
		TotalGaji:          0,
		TunjanganMakan:     400000,
		TunjanganTransport: 500000,
		Status:             "MASUK",
		Role:               "STAFF",
	},
	{
		Id:                 2,
		Absen:              20,
		Cuti:               2,
		Nama:               "Saepul",
		Username:           "hrd",
		Password:           "1234",
		GajiPokok:          3000000,
		TotalGaji:          0,
		TunjanganMakan:     400000,
		TunjanganTransport: 500000,
		Status:             "MASUK",
		Role:               "HRD",
	},
}

func (s *server) Login(ctx context.Context, in *employee.DataLogin) (*employee.Employee, error) {
	for i:= 0; i < len(employeeList); i++ {
		if employeeList[i].Username == in.Username && employeeList[i].Password == in.Password {
			return employeeList[i],nil
		}
	}

	return &employee.Employee{
		Status: "User tidak ditemukan",
	},nil
}

func (s *server) Absen(ctx context.Context, in *employee.Employeeid) (*employee.Result, error) {
	for i:= 0; i < len(employeeList); i++ {
		if employeeList[i].Id == in.Id {
			employeeList[i].Absen += 1
			return &employee.Result{Response: strconv.FormatInt(employeeList[i].Absen, 10)},nil
		}
	}

	return &employee.Result{Response: "GAGAL"},nil
}

func (s *server) CreateEmployee(ctx context.Context, in *employee.Employee) (*employee.Employee, error) {
	idCount := int64(len(employeeList) + 1)

	//error username sama
	employeeList = append(employeeList, in)
	employeeList[len(employeeList)-1].Id = idCount
	employeeList[len(employeeList)-1].Absen = 20
	employeeList[len(employeeList)-1].TunjanganMakan = 400000
	employeeList[len(employeeList)-1].TunjanganTransport = 800000
	employeeList[len(employeeList)-1].GajiPokok = 5000000

	return employeeList[len(employeeList)-1], nil
}

func (s *server) ReadEmployee(ctx context.Context, in *employee.NameId) (*employee.Employee, error) {
	flag := false
	index := 0
	for i := 0; i < len(employeeList); i++ {
		if employeeList[i].Nama == in.Name {
			index = i
			flag = true
			break
		}
	}

	if flag {
		return employeeList[index], nil
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

func (s *server) ReadAllEmployee(ctx context.Context, in *employee.Empty) (*employee.AllEmployee, error) {
	return &employee.AllEmployee{Employes: employeeList},nil
}

func (s *server) DeleteEmployee(ctx context.Context, in *employee.NameId) (*employee.Result, error) {
	flag := false

	for i := 0; i < len(employeeList); i++ {
		if employeeList[i].Nama == in.Name {
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

//func (s *server) LaporanByUsername(ctx context.Context, input *employee.GetEmpByUsername) (*employee.Employee, error) {
//	var EmployeeByUsername *employee.Employee
//	flag := true
//	for _, Cemployee := range employeeList {
//		if strings.EqualFold(Cemployee.username, input.Username) {
//			flag = false
//			tunjanganMakanTemp := Cemployee.tunjanganMakan * (float32(Cemployee.absen) / 22.0)
//			tunjanganTransportTemp := Cemployee.tunjanganTransport * (float32(Cemployee.absen) / 22.0)
//			gajiTotalTemp := Cemployee.gajiPokok + tunjanganMakanTemp + tunjanganTransportTemp
//			EmployeeByUsername = &employee.Employee{
//				Id:                 Cemployee.id,
//				Absen:              Cemployee.absen,
//				Cuti:               Cemployee.cuti,
//				Nama:               Cemployee.nama,
//				Username:           Cemployee.username,
//				Password:           Cemployee.password,
//				GajiPokok:          Cemployee.gajiPokok,
//				TotalGaji:          gajiTotalTemp,
//				TunjanganMakan:     tunjanganMakanTemp,
//				TunjanganTransport: tunjanganTransportTemp,
//				Status:             Cemployee.status,
//				Role:               Cemployee.role,
//			}
//			writeStringToFile(EmployeeByUsername)
//			break
//
//		}
//	}
//	if flag {
//		return EmployeeByUsername, nil
//	} else {
//		return EmployeeByUsername, nil
//	}
//}

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
