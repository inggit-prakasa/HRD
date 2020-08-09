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
		fmt.Scanln(&c)switch c {
		case 1:
			fmt.Println("Manajemen employee")
		case 2:
			fmt.Println("--------------------------------------------------")
			fmt.Println("            Laporan Data Staff                    ")
			fmt.Println("--------------------------------------------------")
			fmt.Println("|| 1. Laporan Staff By Name                     ||")
			fmt.Println("|| 2. Laporan All Staff	                     ||")
			fmt.Println("|| 3. Back to Menu HRD				             ||")
			fmt.Println(" ")
			var d int
			fmt.Print("Input Pilihan Anda :")
			fmt.Scanln(&d)
			switch d {
			case 1:
				conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
				if err != nil {
					log.Fatalf("did not connect: %v", err)
				}
				defer conn.Close()
				c := employee.NewManageEmpClient(conn)
				var e string
				fmt.Print("Input username Anda :")
				fmt.Scanln(&e)
				input := &employee.GetEmpByUsername{
					Username: e,
				}
				result := GetEmployeeByUsername(c, input)
				fmt.Println(result)
			case 2:
				fmt.Println("Laporan All")
			default:
				main()
			}
		}
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
		Id:                 123,
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
		fmt.Printf("Id : %d\n", r.Id)
		fmt.Printf("Absen : %d\n", r.Absen)
		fmt.Printf("Cuti : %d\n", r.Cuti)
		fmt.Printf("Nama : %s\n", r.Nama)
		fmt.Printf("Username : %s\n", r.Username)
		fmt.Printf("Role : %s\n", r.Role)
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


func GetEmployeeByUsername(c employee.ManageEmpClient, input *employee.GetEmpByUsername) string {
	resp, err := c.LaporanByUsername(context.Background(), input)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return "Id : \t \t \t" + strconv.FormatInt(resp.Id, 10) + "\n" +
		"Nama : \t \t \t" + resp.Nama + "\n" +
		"Username : \t \t" + resp.Username + "\n" +
		"Gaji Pokok:\t \t" + fmt.Sprintf("%f", resp.GajiPokok) + "\n" +
		"Tunjangan makan:\t" + fmt.Sprintf("%f", resp.TunjanganMakan) + "\n" +
		"Tunjangan Transport\t:" + fmt.Sprintf("%f", resp.TunjanganTransport) + "\n" +
		"Total Gaji :\t\t" + fmt.Sprintf("%f", resp.TotalGaji) + "\n" +
		"Status :\t \t \t " + resp.Status + "\n" +
		"Role :\t \t \t" + resp.Role
}
// FUNGSI ABSEN
// Koneksi ke GRPC
//

//TIARA

//LOGIN
//PILIHAN STAFF
//MENUSTAFF || MENUHRD
//FUNGSISTAFF || HRD
