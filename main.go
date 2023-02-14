package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./index.html")
		t.Execute(w, nil)
	} else {
		//请求的是查询数据，那么执行查询的逻辑判断
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		var sename = strings.Join(r.Form["username"], "")
		var partname = strings.Join(r.Form["password"], "")
		db, err := sql.Open("mysql", "root:qhdwsx130324/mydb1?charset=utf8")
		infoErr(err)
		if sename != "" && partname != "" {
			var uid int
			var username string
			var password string
			//字符串拼接查询
			err := db.QueryRow("SELECT * FROM user WHERE username ='"+sename+"'AND password ='"+partname+"'").
				Scan(&uid, &username, &password)
			infoErr(err)
			//判断返回的数据是否为空
			if err == sql.ErrNoRows {
				fmt.Fprintf(w, "无该用户数据")
			} else {
				if (sename == username) && (partname == password) {
					fmt.Println(uid)
					fmt.Println(username)
					fmt.Println(password)
					t, _ := template.ParseFiles("./login.html")
					t.Execute(w, nil)
				}
			}
		} else if sename == "" || partname == "" {
			fmt.Fprintf(w, "错误，输入不能为空！")
		}

	}

}

func login_withoutSQL(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method) //获取请求的方法
    if r.Method == "GET" {
        t, _ := template.ParseFiles("D:/Golang/GoItem/go_ex/goSql/login.html")
        t.Execute(w, nil)
    } else {
        //请求的是查询数据，那么执行查询的逻辑判断
        r.ParseForm()
        fmt.Println("username:", r.Form["username"])
        var sename = strings.Join(r.Form["username"], "")
        var partname = strings.Join(r.Form["password"], "")
        db, err := sql.Open("mysql", "root:123456@/test?charset=utf8")
        checkErr(err)
        if sename != "" && partname != "" {
            var uid int
            var username string
            var password string
            //参数查询在一定程度上防止sql注入，参数化查询主要做了两件事:
            //1.参数过滤；2.执行计划重用
            //因为执行计划被重用，所以可以防止SQL注入。
            err := db.QueryRow("SELECT * FROM userinfo WHERE username = ? AND password = ?", sename, partname).
                Scan(&uid, &username, &password)
            //判断返回的数据是否为空
            if err == sql.ErrNoRows {
                fmt.Fprintf(w, "无该用户数据")
            } else {
                if (sename == username) && (partname == password) {
                    fmt.Println(uid)
                    fmt.Println(username)
                    fmt.Println(password)
                    t, _ := template.ParseFiles("D:/Golang/GoItem/go_ex/goSQL/success.html")
                    t.Execute(w, nil)
                }
            }
        } else if sename == "" || partname == "" {
            fmt.Fprintf(w, "错误，输入不能为空！")
        }

    }

}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

func infoErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/login", login)         //设置访问的路由     //设置访问的路由
	err := http.ListenAndServe(":9092", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
