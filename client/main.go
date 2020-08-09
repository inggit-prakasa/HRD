package main

import (
	"context"
	"fmt"

	employee "github.com/inggit_prakasa/HRD/employee"
	"google.golang.org/grpc"
	"log"
)

const (
	address = "localhost:8080"
)

var emp *employee.Employee

func main() {

	fmt.Println("------------------------------------------------------------------------")
	fmt.Println("Welcome To Employee Menu, For All Employee are Our utmost Priority!!")
	fmt.Println("----------------------------------------------------------------")
	var a int
	fmt.Println("<< 1. Login STAFF >>")
	// Username :
	// Password :
	fmt.Println("<< 2. Login HRD   >>")
	// Username :
	// Password :
	fmt.Println("<< 3. Exit       >>")
	//IF Role staff -> MENUSTAFF
	//IF Role hrd -> MENUHRD
	fmt.Scan(&a)
	switch a {
	case 1:
		fmt.Println("--------------------------------------------------")
		fmt.Println("               MENU STAFF                         ")
		fmt.Println("--------------------------------------------------")
		fmt.Println("|| 1. Absen Staff                               ||")
		//Fungsi Absen()
		fmt.Println("|| 2. Lihat Profile Staff Anda                  ||")
		//Fungsi Profile()
		fmt.Println(" ")
		var b int
		fmt.Print("Input Pilihan Anda :")
		fmt.Scanln(&b)

	case 2:
		fmt.Println("--------------------------------------------------")
		fmt.Println("               MENU HRD                           ")
		fmt.Println("--------------------------------------------------")
		fmt.Println("|| 1. Manajemen Employee                        ||")
		//FUNGSI CRUD
		fmt.Println("|| 2. Print Laporan Employee                    ||")
		//FUNGSI LAPORAN
		fmt.Println(" ")
		var c int
		fmt.Print("Input Pilihan Anda :")
		fmt.Scanln(&c)
		switch c {
			case 1:
				fmt.Println("Buat Employee")
				buatEmployee()
			case 2:
				fmt.Println("Baca Employee")
				bacaEmployee()
			case 3:
				fmt.Println("Update Employee")
				updateEmployee()
			case 4:
				fmt.Println("Delete Employee")
				deleteEmployee()
			default:
				break
		}
	default:
		fmt.Println("Mohon input sesuai Menu ")
		fmt.Println(" ")
		main()

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
		Id: 				r1.Id,
		Nama:				nama,
		Username:           username,
		Password:           password,
		Status:             status,
		Role:               role,
	})

	if r.Id == 0 {
		fmt.Println("Tidak ada Karyawan")
	} else {
		fmt.Println("-------------")
		fmt.Printf("Id : %d\n",r.Id)
		fmt.Printf("Absen : %d\n",r.Absen)
		fmt.Printf("Cuti : %d\n",r.Cuti)
		fmt.Printf("Nama : %s\n",r.Nama)
		fmt.Printf("Username : %s\n",r.Username)
		fmt.Printf("Role : %s\n",r.Role)
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
		Nama:				nama,
		Username:           username,
		Password:           password,
		Status:             status,
		Role:               role,
	})

	if err != nil {
		log.Fatalf("Error : %v", err)
	}

	if r.Id == 0 {
		fmt.Println("Tidak ada Karyawan")
	} else {
		fmt.Println("-------------")
		fmt.Printf("Id : %d\n",r.Id)
		fmt.Printf("Absen : %d\n",r.Absen)
		fmt.Printf("Cuti : %d\n",r.Cuti)
		fmt.Printf("Nama : %s\n",r.Nama)
		fmt.Printf("Username : %s\n",r.Username)
		fmt.Printf("Role : %s\n",r.Role)
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
		fmt.Printf("Id : %d\n",r.Id)
		fmt.Printf("Absen : %d\n",r.Absen)
		fmt.Printf("Cuti : %d\n",r.Cuti)
		fmt.Printf("Nama : %s\n",r.Nama)
		fmt.Printf("Username : %s\n",r.Username)
		fmt.Printf("Role : %s\n",r.Role)
	}
}

// FUNGSI ABSEN
// Koneksi ke GRPC
//


//TIARA

//LOGIN
//PILIHAN STAFF
//MENUSTAFF || MENUHRD
//FUNGSISTAFF || HRD
//	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
//	if err != nil {
//		log.Fatalf("did not connect: %v", err)
//	}
//	defer conn.Close()
//	c := employee.NewHRDClient(conn)
//
//	variablelbuatloopmenu := true
//
//	for variablelbuatloopmenu {
//		fmt.Println("Laporan By Username")
//		var username string
//		fmt.Scan(&username)
//
//		input := &employee.GetEmpByUsername{
//			Username: username,
//		}
//		result := GetEmployeeByUsername(c, input)
//		fmt.Println(result)
//		variablelbuatloopmenu = false
//	}
//}

//func GetEmployeeByUsername(c employee.HRDClient, input *employee.GetEmpByUsername) string {
//	resp, err := c.LaporanByUsername(context.Background(), input)
//	if err != nil {
//		log.Fatalf(err.Error())
//	}
//
//	return "Id :" + resp.ID + "\n" +
//		"Nama :" + resp.Nama + "\n" +
//		"Total Gaji :" + resp.TotalGaji + "\n" +
//		"Status :" + resp.Status
//}
