package customer

import (
	"fmt"
)

func showCus() {
	fmt.Println("----------------客户列表---------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱\t")
	show()
	fmt.Println("----------------客户列表完成------------")
	fmt.Println()
}

func updateCus() {
	var sel_update int
	var name string
	var sex string
	var age int
	var phone string
	var email string
	for {
		fmt.Println("请选择待修改客户端编号(-1退出)：")
		fmt.Scanln(&sel_update)
		if sel_update == -1 {
			return
		}
		if sel_update <= 0 || sel_update > len(cusSlice) {
			fmt.Println("没有查到该用户！！！请重新输入正确的编号！！！")
			continue
		}
		break
	}
	fmt.Printf("姓名(%v):", cusSlice[sel_update-1].Name)
	fmt.Scanln(&name)
	fmt.Printf("性别(%v):", cusSlice[sel_update-1].Sex)
	fmt.Scanln(&sex)
	fmt.Printf("年龄(%v):", cusSlice[sel_update-1].Age)
	fmt.Scanln(&age)
	fmt.Printf("电话(%v):", cusSlice[sel_update-1].Phone)
	fmt.Scanln(&phone)
	fmt.Printf("邮箱(%v):", cusSlice[sel_update-1].Email)
	fmt.Scanln(&email)
	if name == "" {
		name = cusSlice[sel_update-1].Name
	}
	if sex == "" {
		sex = cusSlice[sel_update-1].Sex
	}
	if age == 0 {
		age = cusSlice[sel_update-1].Age
	}
	if phone == "" {
		phone = cusSlice[sel_update-1].Phone
	}
	if email == "" {
		email = cusSlice[sel_update-1].Email
	}
	update(sel_update, name, sex, age, phone, email)

}

func deleteCus() {
	var del_id int
	var del_confirm string = ""
	fmt.Printf("请选择待删除客户编号：(-1退出)：")
	fmt.Scanln(&del_id)
	if del_id == -1 {
		return
	}
	fmt.Printf("确认删除？y/n\n")
	for {
		fmt.Scanln(&del_confirm)
		if del_confirm == "n" || del_confirm == "y" {
			break
		} else {
			fmt.Println("请重新输入！！！y/n")
		}
	}
	if del_confirm == "n" {
		return
	}
	if delCus(del_id) {
		fmt.Println("删除成功!!!")
	} else {
		fmt.Println("删除失败！！！")
	}

}

func addCus() {
	var name string
	var sex string
	var age int
	var phone string
	var email string

	fmt.Println("--------------添加客户------------")
	fmt.Printf("姓名：")
	fmt.Scanln(&name)
	fmt.Printf("性别：")
	fmt.Scanln(&sex)
	fmt.Printf("年龄：")
	fmt.Scanln(&age)
	fmt.Printf("电话：")
	fmt.Scanln(&phone)
	fmt.Printf("邮箱：")
	fmt.Scanln(&email)
	if add(name, sex, age, phone, email) {
		fmt.Println("---------------添加完成-----------")

	} else {
		fmt.Println("---------------添加失败-----------")

	}
}

func Menu() {
	loop := true
	var sel string
	for {
		fmt.Println("----------------客户信息管理软件--------------")
		fmt.Println()
		fmt.Println("                1.添 加 客 户                ")
		fmt.Println("                2.修 改 客 户                ")
		fmt.Println("                3.删 除 客 户                ")
		fmt.Println("                4.客 户 列 表                ")
		fmt.Println("                5.退      出                ")
		fmt.Println()
		fmt.Println("                请选择(1-5):                 ")
		fmt.Scanln(&sel)
		switch sel {
		case "1":
			addCus()
		case "2":
			updateCus()
		case "3":
			deleteCus()
		case "4":
			showCus()
		case "5":
			sel_exit := ""
			fmt.Println("确定要退出吗？y/n")
			for {
				fmt.Scanln(&sel_exit)
				if sel_exit == "y" || sel_exit == "n" {
					break
				} else {
					fmt.Println("输入错误！！！请重新输入！！！")
				}
			}
			if sel_exit == "y" {
				loop = false
			}

		default:
			fmt.Println("输入有误！！！")
			fmt.Println("请重新输入！！！")
		}
		if !loop {
			break
		}
	}
}
