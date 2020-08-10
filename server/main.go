package main

import (
	"context"
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
			tunjanganMakanTemp := float64(empLoop.TunjanganMakan) * (float64(empLoop.Absen) / float64(22))
			tunjanganTransportTemp := float64(empLoop.TunjanganTransport) * (float64(empLoop.Absen) / float64(22))
			gajiTotalTemp := empLoop.GajiPokok + int64(tunjanganMakanTemp) + int64(tunjanganTransportTemp)
			EmployeeByUsername = &employee.Employee{
				Id:                 empLoop.Id,
				Absen:              empLoop.Absen,
				Cuti:               empLoop.Cuti,
				Nama:               empLoop.Nama,
				Username:           empLoop.Username,
				Password:           empLoop.Password,
				GajiPokok:          empLoop.GajiPokok,
				TotalGaji:          gajiTotalTemp,
				TunjanganMakan:     empLoop.TunjanganMakan,
				TunjanganTransport: empLoop.TunjanganTransport,
				Message:            empLoop.Message,
				Role:               empLoop.Role,
			}
			employeeList[i] = EmployeeByUsername
			writeStaffToPdf(EmployeeByUsername)
		}
	}
	if flag {
		EmployeeByUsername.Message = "Username Not Found, Please Insert Your Username"
		return EmployeeByUsername, nil
	} else {
		return EmployeeByUsername, nil
	}
}

func (s *server) LaporanAll(ctx context.Context, fileName *employee.FileName) (*employee.AllEmployee, error) {
	for i, empLoop := range employeeList {
		tunjanganMakanTemp := float64(empLoop.TunjanganMakan) * (float64(empLoop.Absen) / float64(22))
		tunjanganTransportTemp := float64(empLoop.TunjanganTransport) * (float64(empLoop.Absen) / float64(22))
		gajiTotalTemp := empLoop.GajiPokok + int64(tunjanganMakanTemp) + int64(tunjanganTransportTemp)
		EmployeeTemp := &employee.Employee{
			Id:                 empLoop.Id,
			Absen:              empLoop.Absen,
			Cuti:               empLoop.Cuti,
			Nama:               empLoop.Nama,
			Username:           empLoop.Username,
			Password:           empLoop.Password,
			GajiPokok:          empLoop.GajiPokok,
			TotalGaji:          gajiTotalTemp,
			TunjanganMakan:     empLoop.TunjanganMakan,
			TunjanganTransport: empLoop.TunjanganTransport,
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
	pdf.Text(20, 10, "============== Laporan"+emp.Username+"================")
	pdf.Text(20, 20, ":"+"Id")
	pdf.Text(60, 20, ":"+strconv.FormatInt(emp.Id, 10))
	pdf.Text(20, 30, "Absen")
	pdf.Text(60, 30, ":"+strconv.FormatInt(emp.Absen, 10))
	pdf.Text(20, 40, "Cuti")
	pdf.Text(60, 40, ":"+strconv.FormatInt(emp.Cuti, 10))
	pdf.Text(20, 50, "Nama")
	pdf.Text(60, 50, ":"+emp.Nama)
	pdf.Text(20, 60, "Username")
	pdf.Text(60, 60, ":"+emp.Username)
	pdf.Text(20, 70, "Gaji Pokok")
	pdf.Text(60, 70, ":"+strconv.FormatInt(emp.GajiPokok, 10))
	pdf.Text(20, 80, "Gaji Total")
	pdf.Text(60, 80, ":"+strconv.FormatInt(emp.TotalGaji, 10))
	pdf.Text(20, 90, "Tunjangan Makan")
	pdf.Text(60, 90, ":"+strconv.FormatInt(emp.TunjanganMakan, 10))
	pdf.Text(20, 100, "Tunjangan Transport")
	pdf.Text(60, 100, ":"+strconv.FormatInt(emp.TunjanganTransport, 10))
	pdf.Text(20, 110, "Role")
	pdf.Text(60, 110, ":"+emp.Role)

	err := pdf.OutputFileAndClose(emp.Username + ".pdf")
	if err != nil {
		log.Println("ERROR", err.Error())
	}

}

func writeAllToPdf(namaFile string) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetFont("Arial", "", 11)
	flag := 120
	for i, empLoop := range employeeList {
		if i%2 == 0 {
			pdf.AddPage()
			flag = 0
		}
		pdf.Text(20, float64(flag+10), "============== Laporan"+empLoop.Nama+"================")
		pdf.Text(20, float64(flag+20), "Id")
		pdf.Text(60, float64(flag+20), ":"+strconv.FormatInt(empLoop.Id, 10))
		pdf.Text(20, float64(flag+30), "Absen")
		pdf.Text(60, float64(flag+30), ":"+strconv.FormatInt(empLoop.Absen, 10))
		pdf.Text(20, float64(flag+40), "Cuti")
		pdf.Text(60, float64(flag+40), ":"+strconv.FormatInt(empLoop.Cuti, 10))
		pdf.Text(20, float64(flag+50), "Nama")
		pdf.Text(60, float64(flag+50), ":"+empLoop.Nama)
		pdf.Text(20, float64(flag+60), "Username")
		pdf.Text(60, float64(flag+60), ":"+empLoop.Username)
		pdf.Text(20, float64(flag+70), "Gaji Pokok")
		pdf.Text(60, float64(flag+70), ":"+strconv.FormatInt(empLoop.GajiPokok, 10))
		pdf.Text(20, float64(flag+80), "Gaji Total")
		pdf.Text(60, float64(flag+80), ":"+strconv.FormatInt(empLoop.TotalGaji, 10))
		pdf.Text(20, float64(flag+90), "Tunjangan Makan")
		pdf.Text(60, float64(flag+90), ":"+strconv.FormatInt(empLoop.TunjanganMakan, 10))
		pdf.Text(20, float64(flag+100), "Tunjangan Transport")
		pdf.Text(60, float64(flag+100), ":"+strconv.FormatInt(empLoop.TunjanganTransport, 10))
		pdf.Text(20, float64(flag+110), "Role")
		pdf.Text(60, float64(flag+110), ":"+empLoop.Role)
		flag = 120

	}

	err := pdf.OutputFileAndClose(namaFile + ".pdf")
	if err != nil {
		log.Println("ERROR", err.Error())
	}
}
