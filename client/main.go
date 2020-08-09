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
		buatEmployee()
	default:
		fmt.Println("Mohon input sesuai Menu ")
		fmt.Println(" ")
		main()

	}
}

func buatEmployee() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	emp := employee.NewManageEmpClient(conn)

	r, err := emp.CreateEmployee(context.Background(), &employee.Employee{
		Id: 123,
		Absen:              12,
		Cuti:               12,
		Nama:               "staff",
		Username:           "staff",
		Password:           "1234",
		GajiPokok:          120000,
		TotalGaji:          120000,
		TunjanganMakan:     200000,
		TunjanganTransport: 200000,
		Status:             "MASUK",
		Role:               "STAFF",
	})
	if err != nil {
		log.Fatalf("Tidak bisa parkir : %v", err)
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
	//conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	//if err != nil {
	//	log.Fatalf("did not connect: %v", err)
	//}
	//defer conn.Close()
	//
	//emp := employee.NewManageEmpClient(conn)
	//
	//r, err := emp.ReadEmployee(context.Background(), &employee.NameId{
	//	Id:   "1234",
	//	Name: "staff",
	//})
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
