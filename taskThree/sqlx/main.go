package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

/*
*
题目1：使用SQL扩展库进行查询
假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
要求 ：
编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。

题目2：实现类型安全映射
假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
要求 ：
定义一个 Book 结构体，包含与 books 表对应的字段。
编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
*/
type Employee struct {
	ID         uint    `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}
type Book struct {
	ID     uint    `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

var DB *sqlx.DB

func initDB() (err error) {
	dsn := "root:root@1234@tcp(127.0.0.1:3306)/sqlx?charset=utf8mb4&parseTime=True&loc=Local" //loc=Asia%2FShanghai
	// 也可以使用sqlx.MustConnect("mysql", dsn)，如果连接失败，就panic
	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Println("连接数据库失败：", err)
		return
	}
	DB.SetMaxOpenConns(50)
	DB.SetMaxIdleConns(10)
	return
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Println("初始化数据库失败：", err)
	}
	fmt.Println("初始化数据库成功")
	//namedBatchInsertEmployee([]Employee{
	//	{Name: "张三", Department: "技术部", Salary: 10000},
	//	{Name: "李四", Department: "人资部", Salary: 5000},
	//	{Name: "王五", Department: "技术部", Salary: 15000},
	//	{Name: "赵六", Department: "技术部", Salary: 12000},
	//	{Name: "Jack", Department: "综合部", Salary: 6000},
	//	{Name: "Tom", Department: "财务部", Salary: 8000},
	//})
	//编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
	queryEmps := namedQueryEmps(Employee{Department: "技术部"})
	fmt.Printf("queryEmps: %#v\n", queryEmps)
	//编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
	emp := queryMostWellPaidEmp()
	fmt.Printf("queryMostWellPaidEmp: %#v\n", emp)
	//编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
	//namedBatchInsertBook([]Book{
	//	{Title: "《活着》", Author: "余华", Price: 30},
	//	{Title: "《老人与海》", Author: "海明威", Price: 40},
	//	{Title: "《三国演义》", Author: "罗贯中", Price: 50},
	//	{Title: "《西游记》", Author: "吴承恩", Price: 60},
	//	{Title: "《水浒传》", Author: "施耐庵", Price: 70},
	//	{Title: "《狂人日记》", Author: "鲁迅", Price: 20},
	//})
	books := queryBooks(50)
	fmt.Printf("queryBooks: %#v\n", books)

}

func queryMostWellPaidEmp() Employee {
	sqlStr := "SELECT * FROM employees order by salary desc limit 1"
	var emp Employee
	err := DB.Get(&emp, sqlStr)
	if err != nil {
		fmt.Println("根据id查询用户失败：", err)
	}
	//fmt.Println("emp:", emp)
	return emp
}

// 批量插入
func namedBatchInsertEmployee(emps []Employee) {
	fmt.Printf("emps: %#v\n", emps)
	sqlStr := "insert into employees(name, department, salary) values(:name, :department, :salary)"
	exec, err := DB.NamedExec(sqlStr, emps)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("exec:", exec)
	id, err := exec.LastInsertId()
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("id:", id)
}

// NamedQuery不使用？作为占位符，使用字段名称或key作为占位符
func namedQueryEmps(emp Employee) []Employee {
	sqlStr := "select * from employees where department = :department"
	//exec, err := DB.NamedExec(sqlStr, map[string]any{"name": emp.Name})
	rows, err := DB.NamedQuery(sqlStr, emp)
	if err != nil {
		fmt.Println("err:", err)
	}
	//fmt.Println("rows:", rows)
	var emps []Employee
	//scans, err := rows.SliceScan()
	//fmt.Println("scans:", scans)
	for rows.Next() {
		var emp Employee
		//scan操作
		err = rows.StructScan(&emp)
		//err = rows.MapScan(&emp)
		if err != nil {
			fmt.Println("err:", err)
		}
		//fmt.Println("emp:", emp)
		emps = append(emps, emp)
	}
	//fmt.Println("emps:", emps)
	return emps
}

// 批量插入
func namedBatchInsertBook(books []Book) {
	fmt.Printf("books: %#v\n", books)
	sqlStr := "insert into books(title, author, price) values(:title, :author, :price)"
	exec, err := DB.NamedExec(sqlStr, books)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("exec:", exec)
	id, err := exec.LastInsertId()
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("id:", id)
}

// 查询
func queryBooks(price float64) []Book {
	sqlStr := "SELECT * FROM books WHERE price > ?"
	var books []Book
	err := DB.Select(&books, sqlStr, price)
	if err != nil {
		fmt.Println("根据price查询books失败：", err)
	}
	return books
}
