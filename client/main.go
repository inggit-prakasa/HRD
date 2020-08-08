package main

import (
	"fmt"
)

const (
	address = "localhost:8080"
)

func main() {

	fmt.Println("------------------------------------------------------------------------")
	fmt.Println("Welcome To Employee Menu, For All Employee are Our utmost Priority!!")
	fmt.Println("----------------------------------------------------------------")
	var a int
	fmt.Println("<< 1. Login STAFF >>")
	fmt.Println("<< 2. Login HRD   >>")
	fmt.Println("<< 3. Exit       >>")
	fmt.Print("Input Login sesuai yang Anda inginkan: ")
	fmt.Scanln(&a)
	switch a {
	case 1:
		fmt.Println("--------------------------------------------------")
		fmt.Println("               MENU STAFF                         ")
		fmt.Println("--------------------------------------------------")
		fmt.Println("|| 1. Absen Staff                               ||")
		fmt.Println("|| 2. Lihat Profile Staff Anda                  ||")
		fmt.Println(" ")
		var b int
		fmt.Print("Input Pilihan Anda :")
		fmt.Scanln(&b)

	case 2:
		fmt.Println("--------------------------------------------------")
		fmt.Println("               MENU HRD                           ")
		fmt.Println("--------------------------------------------------")
		fmt.Println("|| 1. Manajemen Employee                        ||")
		fmt.Println("|| 2. Print Laporan Employee                    ||")
		fmt.Println(" ")
		var c int
		fmt.Print("Input Pilihan Anda :")
		fmt.Scanln(&c)
	default:
		fmt.Println("Mohon input sesuai Menu ")
		fmt.Println(" ")
		main()

	}
}

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
