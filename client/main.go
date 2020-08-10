package main

import (
	"context"
	"fmt"
	"log"

	employee "github.com/inggit_prakasa/HRD/employee"
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
				absenStaff(emp.Id)
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
	fmt.Println(r)
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

	for i := 0; i < len(r.Employes); i++ {
		fmt.Println(r.Employes[i].Id)
		fmt.Println(r.Employes[i].Nama)
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

	fmt.Print("nama : ")
	fmt.Scan(&nama)

	r, err := emp.DeleteEmployee(context.Background(), &employee.NameId{
		Name: nama,
	})

	fmt.Println(r.Response)
}

func updateEmployee() {
	var (
		nama, username, password, status, role string
	)

	fmt.Print("nama : ")
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
		fmt.Println("Nama salah")
		return
	}

	fmt.Print("Update nama")
	fmt.Print("Nama : ")
	fmt.Scan(&nama)
	fmt.Print("username : ")
	fmt.Scan(&username)
	fmt.Print("password : ")
	fmt.Scan(&password)
	fmt.Print("status : ")
	fmt.Scan(&status)
	fmt.Print("role : ")
	fmt.Scan(&role)

	r, err := emp.UpdateEmployee(context.Background(), &employee.Employee{
		Id:       r1.Id,
		Nama:     nama,
		Username: username,
		Password: password,
		Status:   status,
		Role:     role,
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

func buatEmployee() {
	var (
		nama, username, password, status, role string
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
	fmt.Print("status : ")
	fmt.Scan(&status)
	fmt.Print("role : ")
	fmt.Scan(&role)

	r, err := emp.CreateEmployee(context.Background(), &employee.Employee{
		Nama:     nama,
		Username: username,
		Password: password,
		Status:   status,
		Role:     role,
	})

	if err != nil {
		log.Fatalf("Error : %v", err)
	}

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

func bacaEmployee() {
	var (
		nama string
	)

	fmt.Print("nama : ")
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
	fmt.Println("|| 1. Laporan Staff By Name\t\t\t||")
	fmt.Println("|| 2. Laporan All Staff\t\t\t||")
	fmt.Println("|| 3. Exit\t\t\t\t\t\t||")
	fmt.Println("------------ Please Press [1/2/3] --------------- ")
	fmt.Println(" ")
	var d int
	fmt.Print("Input Pilihan Anda :")
	fmt.Scanln(&d)
	switch d {
	case 1:
		singleStaffReport()
		laporanEmployee()
	case 2:
		allStaffReport()
		laporanEmployee()
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
	fmt.Scanln(&e)
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
	fmt.Scanln(&inp)
	input := &employee.FileName{
		File: inp,
	}
	result := LaporanAllEmployee(c, input)
	fmt.Println(result)
}
