package main

import (
	"context"
	"fmt"
	"strings"

	//employee "github.com/inggit_prakasa/HRD/employee"
	"log"
	"net"
	"strconv"

	"HRD/employee"

	"github.com/jung-kurt/gofpdf"
	"google.golang.org/grpc"
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
		Message:            "MASUK",
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
		Message:            "MASUK",
		Role:               "HRD",
	},
}

func (s *server) Login(ctx context.Context, in *employee.DataLogin) (*employee.Employee, error) {
	for i := 0; i < len(employeeList); i++ {
		if employeeList[i].Username == in.Username && employeeList[i].Password == in.Password {
			return employeeList[i], nil
		}
	}

	return &employee.Employee{
		Message: "User tidak ditemukan",
	}, nil
}

func (s *server) Absen(ctx context.Context, in *employee.Employeeid) (*employee.Result, error) {
	for i := 0; i < len(employeeList); i++ {
		if employeeList[i].Id == in.Id {
			employeeList[i].Absen += 1
			return &employee.Result{Response: strconv.FormatInt(employeeList[i].Absen, 10)}, nil
		}
	}

	return &employee.Result{Response: "GAGAL"}, nil
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
		Message:            in.Message,
		Role:               in.Role,
	}, nil
}

func (s *server) ReadAllEmployee(ctx context.Context, in *employee.Empty) (*employee.AllEmployee, error) {
	return &employee.AllEmployee{Employes: employeeList}, nil
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

func (s *server) LaporanByUsername(ctx context.Context, input *employee.Username) (*employee.Employee, error) {
	var EmployeeByUsername *employee.Employee
	flag := true
	for i, empLoop := range employeeList {
		if strings.EqualFold(empLoop.Username, input.Username) {
			flag = false
			tunjanganMakanTemp := 400000 * (empLoop.Absen / 22)
			tunjanganTransportTemp := 500000 * (empLoop.Absen / 22)
			gajiTotalTemp := empLoop.GajiPokok + tunjanganMakanTemp + tunjanganTransportTemp
			EmployeeByUsername = &employee.Employee{
				Id:                 empLoop.Id,
				Absen:              empLoop.Absen,
				Cuti:               empLoop.Cuti,
				Nama:               empLoop.Nama,
				Username:           empLoop.Username,
				Password:           empLoop.Password,
				GajiPokok:          empLoop.GajiPokok,
				TotalGaji:          gajiTotalTemp,
				TunjanganMakan:     tunjanganMakanTemp,
				TunjanganTransport: tunjanganTransportTemp,
				Message:            empLoop.Message,
				Role:               empLoop.Role,
			}
			employeeList[i] = EmployeeByUsername
			writeStaffToPdf(EmployeeByUsername)
			break

		}
	}
	if flag {
		return EmployeeByUsername, nil
	} else {
		return EmployeeByUsername, nil
	}
}

func (s *server) LaporanAll(ctx context.Context, fileName *employee.FileName) (*employee.AllEmployee, error) {
	for i, empLoop := range employeeList {
		tunjanganMakanTemp := 400000 * (empLoop.Absen / 22)
		tunjanganTransportTemp := 500000 * (empLoop.Absen / 22)
		gajiTotalTemp := empLoop.GajiPokok + tunjanganMakanTemp + tunjanganTransportTemp
		EmployeeTemp := &employee.Employee{
			Id:                 empLoop.Id,
			Absen:              empLoop.Absen,
			Cuti:               empLoop.Cuti,
			Nama:               empLoop.Nama,
			Username:           empLoop.Username,
			Password:           empLoop.Password,
			GajiPokok:          empLoop.GajiPokok,
			TotalGaji:          gajiTotalTemp,
			TunjanganMakan:     tunjanganMakanTemp,
			TunjanganTransport: tunjanganTransportTemp,
			Message:            empLoop.Message,
			Role:               empLoop.Role,
		}
		employeeList[i] = EmployeeTemp
	}
	writeAllToPdf(fileName.File)
	return &employee.AllEmployee{Employes: employeeList}, nil
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

func writeStaffToPdf(emp *employee.Employee) {

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 11)
	pdf.Text(20, 10, "============== Laporan"+emp.Nama+"================")
	pdf.Text(20, 20, "Id  \t \t \t :"+strconv.FormatInt(emp.Id, 10))
	pdf.Text(20, 30, "Absen \t \t:"+strconv.FormatInt(emp.Absen, 10))
	pdf.Text(20, 40, "Cuti  \t \t:"+strconv.FormatInt(emp.Cuti, 10))
	pdf.Text(20, 50, "Nama  \t \t:"+emp.Nama)
	pdf.Text(20, 60, "Username  \t \t:"+emp.Username)
	pdf.Text(20, 70, "Gaji Pokok  \t \t:"+fmt.Sprintf("%.3f", emp.GajiPokok))
	pdf.Text(20, 80, "Gaji Total  \t \t:"+fmt.Sprintf("%.3f", emp.TotalGaji))
	pdf.Text(20, 90, "Tunjangan Makan  \t :"+fmt.Sprintf("%.3f", emp.TunjanganMakan))
	pdf.Text(20, 100, "Tunjangan Transport  :"+fmt.Sprintf("%.3f", emp.TunjanganTransport))
	pdf.Text(20, 110, "Message   \t:"+emp.Message)
	pdf.Text(20, 120, "Role  \t \t:"+emp.Role)

	err := pdf.OutputFileAndClose("laporan/" + emp.Nama + ".pdf")
	if err != nil {
		log.Println("ERROR", err.Error())
	}

}

func writeAllToPdf(namaFile string) {
	str1 := "========================================================================================================= \n "
	str2 := "| ID|Nama\t|UserName|Gaji Pokok\t|Tunj Makan\t|Tunj Transport\t|Total Gaji\t|Message\t\t|Role\t\t|"
	str3 := "============================================================================================================ \n "
	pdf := gofpdf.New("L", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 11)
	x := 0
	for _, empLoop := range employeeList {
		if x == 1 {
			pdf.AddPage()
			x = 0
		}
		pdf.Text(20, 10, str1)
		pdf.Text(20, 20, str2)
		pdf.Text(20, 30, str3)
		pdf.Text(20, float64((x+4)*10), "| "+strconv.FormatInt(empLoop.Id, 10)+"|"+empLoop.Nama+"\t|"+empLoop.Username+"\t|"+
			fmt.Sprintf("%.2f", empLoop.GajiPokok)+"\t|"+fmt.Sprintf("%.2f", empLoop.TunjanganMakan)+"\t"+
			"|"+fmt.Sprintf("%.2f", empLoop.TunjanganTransport)+"\t|"+fmt.Sprintf("%.2f", empLoop.TotalGaji)+
			"\t|"+empLoop.Message+"\t\t|"+empLoop.Role+"\t\t|")
		x++

	}

	err := pdf.OutputFileAndClose("laporan/" + namaFile + ".pdf")
	if err != nil {
		log.Println("ERROR", err.Error())
	}
}
