package main

import (
	"context"
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"strings"

	employee "github.com/inggit_prakasa/HRD/employee"
	"log"
	"net"
	"strconv"

	//"HRD/employee"

	"github.com/jung-kurt/gofpdf"
	"google.golang.org/grpc"
	_ "github.com/go-sql-driver/mysql"
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

func conn() *sql.DB {
	db, err := sql.Open("mysql","root:12345678@tcp(127.0.0.1:3306)/bootcamp")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func (s *server) Login(ctx context.Context, in *employee.DataLogin) (*employee.Employee, error) {
	db := conn()
	defer db.Close()

	results, err := db.Query("SELECT * FROM employee")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		empl := &employee.Employee{}
		err = results.Scan(&empl.Id,&empl.Nama,&empl.Absen,&empl.Cuti,&empl.Username,&empl.Password,&empl.GajiPokok,&empl.TotalGaji,&empl.TunjanganMakan,&empl.TunjanganTransport,&empl.Role)
		if err != nil {
			panic(err.Error())
		}

		if (empl.Username == in.Username) && empl.Password == in.Password {
			return empl,nil
		}
	}

	return &employee.Employee{},nil
}

func (s *server) Absen(ctx context.Context, in *employee.Employeeid) (*employee.Result, error) {
	db := conn()
	defer db.Close()

	results, err := db.Query("SELECT id,absen FROM employee")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		empl := &employee.Employee{}
		err = results.Scan(&empl.Id,&empl.Absen)

		if in.Id == empl.Id {
			_, err := db.Exec("UPDATE employee SET absen = ? WHERE id = ?",empl.Absen+1,empl.Id)
			if err != nil {
				return &employee.Result{Response: strconv.FormatInt(empl.Absen+1, 10)},err
			}
			return &employee.Result{Response: strconv.FormatInt(empl.Absen+1, 10)},nil
		}
	}

	return &employee.Result{Response: "GAGAL"},nil
}


func (s *server) CreateEmployee(ctx context.Context, in *employee.Employee) (*employee.Employee, error) {
	db := conn()
	defer db.Close()

	if handUsername(in.Username) {
		return &employee.Employee{},nil
	}

	//pass, _ := HashPassword(in.Password)
	password1 := []byte(in.Password)

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password1, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(hashedPassword))

	stmt, err := db.Prepare("INSERT INTO employee (nama,username,password,role,absen,tunjanganMakan,tunjanganTransport,gajiPokok) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?)")

	if err != nil {
		return &employee.Employee{},err
	}

	absen := 20
	tunjMakan := 400000
	tunjTransport := 800000
	gajiPokok := 5000000

	_, err = stmt.Exec(in.Nama, in.Username, password1,in.Role,absen,tunjMakan,tunjTransport,gajiPokok)

	if err != nil {
		return &employee.Employee{},err
	}

	return &employee.Employee{
		Id:					1,
		Absen:              int64(absen),
		Nama:               in.Nama,
		Username:           in.Username,
		Password:           in.Password,
		GajiPokok:          int64(gajiPokok),
		TunjanganMakan:     int64(tunjMakan),
		TunjanganTransport: int64(tunjTransport),
		Role:               in.Role,
	}, nil
}

func (s *server) ReadEmployee(ctx context.Context, in *employee.NameId) (*employee.Employee, error) {
	db := conn()
	defer db.Close()
	empl := &employee.Employee{}
	stmt, err := db.Prepare("SELECT * FROM employee WHERE nama = ?")

	if err != nil {
		panic(err.Error())
	}

	err = stmt.QueryRow(in.Name).Scan(&empl.Id,&empl.Nama,&empl.Absen,&empl.Cuti,&empl.Username,&empl.Password,&empl.GajiPokok,&empl.TotalGaji,&empl.TunjanganMakan,&empl.TunjanganTransport,&empl.Role)

	if err != nil {
		return empl,nil
	}

	return empl,nil
}

func (s *server) UpdateEmployee(ctx context.Context, in *employee.Employee) (*employee.Employee, error) {
	db := conn()
	defer db.Close()

	stmt, err := db.Prepare("UPDATE employee SET nama=?, username=?, password=?, role=? WHERE id = ?")

	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(in.Nama, in.Username, in.Password,in.Role,in.Id)

	if err != nil {
		panic(err.Error())
	}
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
		Role:               in.Role,
	}, nil
}

func SelectAllEmployee() []*employee.Employee {
	db := conn()
	defer db.Close()

	var employeeList1 = []*employee.Employee{}

	results, err := db.Query("SELECT * FROM employee")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		empl := &employee.Employee{}
		// for each row, scan the result into our tag composite object
		err = results.Scan(&empl.Id,&empl.Nama,&empl.Absen,&empl.Cuti,&empl.Username,&empl.Password,&empl.GajiPokok,&empl.TotalGaji,&empl.TunjanganMakan,&empl.TunjanganTransport,&empl.Role)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		employeeList1 = append(employeeList1, empl)
	}
	return employeeList1
}

func (s *server) ReadAllEmployee(ctx context.Context, in *employee.Empty) (*employee.AllEmployee, error) {
	var employeeList1 = SelectAllEmployee()

	return &employee.AllEmployee{Employes: employeeList1},nil
}

func handUsername(username string) bool {
	db := conn()
	defer db.Close()

	results, err := db.Query("SELECT id,username FROM employee")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		empl := &employee.Employee{}
		err = results.Scan(&empl.Id,&empl.Username)

		if empl.Username == username {
			return true
		}
	}

	return false
}

func (s *server) DeleteEmployee(ctx context.Context, in *employee.NameId) (*employee.Result, error) {
	db := conn()
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM employee WHERE nama = ?")

	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(in.Name)

	if err != nil {
		panic(err.Error())
	}

	return &employee.Result{Response: "BERHASIl"}, nil
}

func (s *server) LaporanByUsername(ctx context.Context, input *employee.Username) (*employee.Employee, error) {
	var EmployeeByUsername *employee.Employee
	var employeeList1 = SelectAllEmployee()
	flag := true
	for i, empLoop := range employeeList1 {
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
			employeeList1[i] = EmployeeByUsername
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
	db := conn()
	defer db.Close()

	stmt, err := db.Prepare("UPDATE employee SET totalGaji=? WHERE id = ?")

	if err != nil {
		panic(err.Error())
	}

	var employeeList1 = SelectAllEmployee()
	for i, empLoop := range employeeList1 {
		tunjanganMakanTemp := float64(empLoop.TunjanganMakan) * (float64(empLoop.Absen) / float64(22))
		tunjanganTransportTemp := float64(empLoop.TunjanganTransport) * (float64(empLoop.Absen) / float64(22))
		gajiTotalTemp := empLoop.GajiPokok + int64(tunjanganMakanTemp) + int64(tunjanganTransportTemp)
		_, err = stmt.Exec(gajiTotalTemp,empLoop.Id)

		if err != nil {
			panic(err.Error())
		}
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
		employeeList1[i] = EmployeeTemp
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
	var employeeList1 = SelectAllEmployee()
	for i, empLoop := range employeeList1 {
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
