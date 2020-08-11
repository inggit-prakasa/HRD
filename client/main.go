package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	employee "github.com/inggit_prakasa/HRD/employee"
	//"HRD/employee"

	"google.golang.org/grpc"
)

const (
	address = "localhost:8080"
)

var emp *employee.Employee

func main() {

	fmt.Println("------------------------------------------------------------------------")
	fmt.Println("Welcome To Employee Menu, For All Employee are Our utmost Priority!!")
	fmt.Println("----------------------------------------------------------------")
	var username, password string
	fmt.Println("<<  Login Page >>")
	fmt.Print("username : ")
	fmt.Scan(&username)
	fmt.Print("password : ")
	fmt.Scan(&password)

	emp := login(username, password)
	switch emp.Role {
	case "STAFF":
		var b int
		handlAbsen := true
		flag := true
		for flag {
			fmt.Println("--------------------------------------------------")
			fmt.Println("                    MENU STAFF                    ")
			fmt.Println("--------------------------------------------------")
			fmt.Println("|| 1. Absen Staff                               ||")
			fmt.Println("|| 2. Lihat Profile Staff Anda                  ||")
			fmt.Println("|| 3. Log Out					||")
			fmt.Print("Input Pilihan Anda : ")
			fmt.Scan(&b)
			switch b {
			case 1:
				if !handlAbsen {
					fmt.Println("Anda sudah absen hari ini")
					break
				}
				absenStaff(emp.Id)
				handlAbsen = false
			case 2:
				lihatProfile(emp.Nama)
			case 3:
				flag = false
				break
			default:
				fmt.Println("Pilihan anda salah coba lagi")
			}
		}
		main()
	case "HRD":
		flag := true
		for flag {
			fmt.Println("--------------------------------------------------")
			fmt.Println("                     MENU HRD                     ")
			fmt.Println("--------------------------------------------------")
			fmt.Println("|| 1. Create Employee    \t\t\t||")
			fmt.Println("|| 2. Find Employee      \t\t\t||")
			fmt.Println("|| 3. Edit Employee      \t\t\t||")
			fmt.Println("|| 4. Delete Employee    \t\t\t||")
			fmt.Println("|| 5. All Employee       \t\t\t||")
			fmt.Println("|| 6. Print Laporan Employee \t\t\t||")
			fmt.Println("|| 7. LOG OUT            \t\t\t||")
			//FUNGSI LAPORAN
			fmt.Println(" ")
			fmt.Print("Input Pilihan Anda : ")
			var c int
			fmt.Scan(&c)
			switch c {
			case 1:
				buatEmployee()
			case 2:
				bacaEmployee()
			case 3:
				updateEmployee()
			case 4:
				deleteEmployee()
			case 5:
				bacaAllEmployee()
			case 6:
				laporanEmployee()
			case 7:
				flag = false
				break
			}
		}
		main()
	default:
		fmt.Println("Login gagal")
		main()

	}
}

func absenStaff(id int64) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	emp := employee.NewManageEmpClient(conn)

	r, err := emp.Absen(context.Background(), &employee.Employeeid{
		Id: id,
	})

	fmt.Println("Anda telah absen : ", r.Response)
}

func lihatProfile(nama string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	emp := employee.NewManageEmpClient(conn)

	r, err := emp.ReadEmployee(context.Background(), &employee.NameId{
		Name: nama,
	})
	fmt.Println("----PROFILE-----")
	fmt.Println("Nama \t\t:", r.Nama)
	fmt.Println("Absen \t\t:", r.Absen)
	fmt.Println("Cuti \t\t:", r.Cuti)
	fmt.Println("Gaji Pokok \t:", r.GajiPokok)
	fmt.Println("Role \t\t:", r.Role)
	fmt.Println("Total Gaji \t:", r.TotalGaji)
}

func login(username string, password string) *employee.Employee {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	emp := employee.NewManageEmpClient(conn)

	r, err := emp.Login(context.Background(), &employee.DataLogin{
		Username: username,
		Password: password,
	})

	return r
}

func bacaAllEmployee() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	emp := employee.NewManageEmpClient(conn)

	r, err := emp.ReadAllEmployee(context.Background(), &employee.Empty{})
	fmt.Println("ID \t Nama \t Username \t Absen \t Role")
	for i:=0; i < len(r.Employes) ;i++ {
		fmt.Printf("%d \t %s \t %s \t\t %d \t %s \n", r.Employes[i].Id,r.Employes[i].Nama,r.Employes[i].Username,r.Employes[i].Absen,r.Employes[i].Role)
	}
}

func deleteEmployee() {
	var nama string
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	emp := employee.NewManageEmpClient(conn)

	fmt.Print("username/ID : ")
	fmt.Scan(&nama)

	r, err := emp.DeleteEmployee(context.Background(), &employee.NameId{
		Name: nama,
	})

	fmt.Println(r.Response)
}

func updateEmployee() {
	var (
		pilRole int
		nama, username, password, role string
	)

	fmt.Print("Username/ID : ")
	fmt.Scan(&nama)

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	emp := employee.NewManageEmpClient(conn)

	r1, err := emp.ReadEmployee(context.Background(), &employee.NameId{
		Name: nama,
	})

	if r1.Id == 0 {
		fmt.Println("Username/id salah")
		return
	}

	fmt.Println("--- EDIT --- ")
	fmt.Print("Nama : ")
	fmt.Scan(&nama)
	fmt.Print("username : ")
	fmt.Scan(&username)
	fmt.Print("password : ")
	fmt.Scan(&password)
	fmt.Print("role [1.STAFF/2.HRD] : ")
	fmt.Scan(&pilRole)
	if pilRole == 1 {
		role = "STAFF"
	} else if pilRole == 2 {
		role = "HRD"
	} else {
		fmt.Println("Pilihan salah, Coba Lagi")
		return
	}

	r, err := emp.UpdateEmployee(context.Background(), &employee.Employee{
		Id:       r1.Id,
		Nama:     nama,
		Username: username,
		Password: password,
		Role:     role,
	})

	if r.Id == 0 {
		fmt.Println("Username sudah digunakan, Coba lagi")
	} else {
		fmt.Println("-------------")
		fmt.Printf("Id : %d\n", r.Id)
		fmt.Printf("Absen : %d\n", r.Absen)
		fmt.Printf("Cuti : %d\n", r.Cuti)
		fmt.Printf("Nama : %s\n", r.Nama)
		fmt.Printf("Username : %s\n", r.Username)
		fmt.Printf("Role : %s\n", r.Role)
	}
}

func buatEmployee() {
	var (
		pilRole int
		nama, username, password, role string
	)

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	emp := employee.NewManageEmpClient(conn)

	fmt.Print("Nama : ")
	fmt.Scan(&nama)
	fmt.Print("username : ")
	fmt.Scan(&username)
	fmt.Print("password : ")
	fmt.Scan(&password)
	fmt.Print("role [1.STAFF/2.HRD] : ")
	fmt.Scan(&pilRole)
	if pilRole == 1 {
		role = "STAFF"
	} else if pilRole == 2 {
		role = "HRD"
	} else {
		fmt.Println("Pilihan salah, Coba Lagi")
		return
	}

	r, err := emp.CreateEmployee(context.Background(), &employee.Employee{
		Nama:     nama,
		Username: username,
		Password: password,
		Role:     role,
	})

	if err != nil {
		log.Fatalf("Error : %v", err)
	}

	if r.Id == 0 {
		fmt.Println("Username sudah dipakai")
	} else {
		fmt.Println("-------------")
		fmt.Printf("Id : %d\n", r.Id)
		fmt.Printf("Absen : %d\n", r.Absen)
		fmt.Printf("Cuti : %d\n", r.Cuti)
		fmt.Printf("Nama : %s\n", r.Nama)
		fmt.Printf("Username : %s\n", r.Username)
		fmt.Printf("Role : %s\n", r.Role)
	}

}

func bacaEmployee() {
	var (
		nama string
	)

	fmt.Print("username/ID : ")
	fmt.Scan(&nama)
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	emp := employee.NewManageEmpClient(conn)

	r, err := emp.ReadEmployee(context.Background(), &employee.NameId{
		Name: nama,
	})

	if r.Id == 0 {
		fmt.Println("Tidak ada Karyawan")
	} else {
		fmt.Println("-------------")
		fmt.Printf("Id : %d\n", r.Id)
		fmt.Printf("Absen : %d\n", r.Absen)
		fmt.Printf("Cuti : %d\n", r.Cuti)
		fmt.Printf("Nama : %s\n", r.Nama)
		fmt.Printf("Username : %s\n", r.Username)
		fmt.Printf("Role : %s\n", r.Role)
	}
}

func laporanEmployee() {
	fmt.Println("--------------------------------------------------")
	fmt.Println("            Laporan Data Staff                    ")
	fmt.Println("--------------------------------------------------")
	fmt.Println("|| 1. Laporan Staff By Username\t\t\t||")
	fmt.Println("|| 2. Laporan All Staff\t\t\t\t||")
	fmt.Println("|| 3. Logout\t\t\t\t\t||")
	fmt.Println("------------ Please Press [1/2/3] --------------- ")
	fmt.Println(" ")
	var d int
	fmt.Print("Input Pilihan Anda :")
	fmt.Scan(&d)
	switch d {
	case 1:
		singleStaffReport()
	case 2:
		allStaffReport()
	case 3:
		main()
	default:
		fmt.Println("[Warning !!!]Wrong Input, Please Press [1/2/3] = ")
		laporanEmployee()

	}
}

func singleStaffReport() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := employee.NewManageEmpClient(conn)
	var e string
	fmt.Print("Input username Anda :")
	fmt.Scan(&e)
	input := &employee.Username{
		Username: e,
	}
	result := GetEmployeeByUsername(c, input)
	fmt.Println(result)
}

func allStaffReport() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := employee.NewManageEmpClient(conn)
	var inp string
	fmt.Print("Input Nama File Laporan :")
	fmt.Scan(&inp)
	input := &employee.FileName{
		File: inp,
	}
	result := LaporanAllEmployee(c, input)
	fmt.Println(result)
}

func GetEmployeeByUsername(c employee.ManageEmpClient, input *employee.Username) string {
	resp, err := c.LaporanByUsername(context.Background(), input)
	if err != nil {
		log.Fatalf(err.Error())
	}

	strmsg := "Username Not Found, Please Insert Your Username"
	if strings.EqualFold(resp.Message, strmsg) {
		return strmsg
	} else {
		return "===Laporan Employee" + resp.Nama + "===\n" +
			"Id  \t \t \t:" + strconv.FormatInt(resp.Id, 10) + "\n" +
			"Nama  \t \t \t:" + resp.Nama + "\n" +
			"Username  \t \t:" + resp.Username + "\n" +
			"Gaji Pokok\t \t:" + strconv.FormatInt(resp.GajiPokok, 10) + "\n" +
			"Tunjangan makan\t\t:" + strconv.FormatInt(resp.TunjanganMakan, 10) + "\n" +
			"Tunjangan Transport\t:" + strconv.FormatInt(resp.TunjanganTransport, 10) + "\n" +
			"Total Gaji \t\t:" + strconv.FormatInt(resp.TotalGaji, 10) + "\n" +
			"Role \t \t \t:" + resp.Role + "\n"
	}

}

func LaporanAllEmployee(c employee.ManageEmpClient, input *employee.FileName) string {
	resp, err := c.LaporanAll(context.Background(), input)
	if err != nil {
		log.Fatalf(err.Error())
	}
	if len(resp.Employes) == 0 {
		return "Belum ada Staff yang terdaftar di Aplikasi"
	}
	var AllEmp string
	for _, empLoop := range resp.Employes {
		AllEmp += "===Laporan Employee" + empLoop.Nama + "===\n" +
			"Id  \t \t \t:" + strconv.FormatInt(empLoop.Id, 10) + "\n" +
			"Nama  \t \t \t:" + empLoop.Nama + "\n" +
			"Username  \t \t:" + empLoop.Username + "\n" +
			"Gaji Pokok\t \t:" + strconv.FormatInt(empLoop.GajiPokok, 10) + "\n" +
			"Tunjangan makan\t\t:" + strconv.FormatInt(empLoop.TunjanganMakan, 10) + "\n" +
			"Tunjangan Transport\t:" + strconv.FormatInt(empLoop.TunjanganTransport, 10) + "\n" +
			"Total Gaji \t\t:" + strconv.FormatInt(empLoop.TotalGaji, 10) + "\n" +
			"Role \t \t \t:" + empLoop.Role + "\n"
	}

	return AllEmp
}
