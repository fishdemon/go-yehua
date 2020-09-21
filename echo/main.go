package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"io"
	"net/http"
	"os"
	"strconv"
)

type User struct {
	ID string
	Name string
	Age int8
	Username string
	Password string
}

func SaveUser(c echo.Context) error {
	// form
	username := c.FormValue("u")
	password := c.FormValue("p")

	fmt.Println(username + ":" + password)
	// string response
	return c.String(http.StatusOK, "success")
}

func GetUser(c echo.Context) error {
	// path
	id := c.Param("id")

	fmt.Println("id:" + id)

	u := User{"1235465432", "allen", 19, "allen", "123456"}
	return c.JSON(http.StatusOK, u)
}

func UpdateUser(c echo.Context) error {
	// path + json
	id := c.Param("id")
	u := new(User)
	if err := c.Bind(u); err != nil {
 		return c.NoContent(http.StatusBadRequest)
	}
	u.ID = id

	if jsonStr, err := json.Marshal(u); err == nil {
		fmt.Println(string(jsonStr))
	}
	return c.NoContent(http.StatusOK)
}

func DeleteUser(c echo.Context) error {
	// path
	id := c.Param("id")
	fmt.Println("delete : " + id)
	return c.NoContent(http.StatusOK)
}

func FindUser(c echo.Context) error {
	// query params
	name := c.QueryParam("name")
	ageStr := c.QueryParam("age")
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	fmt.Println(name)
	fmt.Println(age)
	// json response
	return c.JSON(http.StatusOK, name)
}

func request(c echo.Context) error {
	// request
	path := c.Path()
	fmt.Println("path: " + path)

	ip := c.RealIP()
	fmt.Println("ip: " + ip)

	scheme := c.Scheme()
	fmt.Println("scheme: " + scheme)

	cookies := c.Cookies()
	for _,v:=range cookies {
		str, err := json.Marshal(v)
		if err == nil {
			fmt.Println("json is err")
		}
		fmt.Println(string(str))
	}

	userAgent := c.Request().UserAgent()
	fmt.Println("user agent: " + userAgent)

	proto := c.Request().Proto
	fmt.Println("proto: " + proto)

	host := c.Request().Host
	fmt.Println("host: " + host)

	url := c.Request().URL
	fmt.Println("url: " + url.String())

	fmt.Println("---------headers-------")
	header := c.Request().Header
	for k, _ := range header {
		fmt.Println(k + ":" + header.Get(k))
	}

	return c.NoContent(http.StatusOK)
}

func upload(c echo.Context) error {
	// 通过FormFile函数获取客户端上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	//打开用户上传的文件
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// 创建目标文件，就是我们打算把用户上传的文件保存到什么地方
	// file.Filename 参数指的是我们以用户上传的文件名，作为目标文件名，也就是服务端保存的文件名跟用户上传的文件名一样
	userHome, _ := os.UserHomeDir()
	dst, err := os.Create(userHome + "/" + file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// 这里将用户上传的文件复制到服务端的目标文件
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.String(http.StatusOK, fmt.Sprintf("文件上传成功: %s", file.Filename))
}

func main() {
	fmt.Println("Starting Server")

	// 实例化echo对象
	e := echo.New()

	//注册一个Get请求, 路由地址为: /tizi365  并且绑定一个控制器函数, 这里使用的是闭包函数。
	e.GET("/tizi365", func(c echo.Context) error {
		//控制器函数直接返回一个字符串，http响应状态为http.StatusOK，就是200状态。
		return c.String(http.StatusOK, "欢迎访问tizi365.com")
	})

	e.POST("/users", SaveUser)
	e.GET("/users/:id", GetUser)
	e.PUT("/users/:id", UpdateUser)
	e.DELETE("/users/:id", DeleteUser)
	e.GET("/users", FindUser)
	e.GET("/request", request)
	e.POST("/file", upload)

	//启动http server, 并监听8080端口，冒号（:）前面为空的意思就是绑定网卡所有Ip地址，本机支持的所有ip地址都可以访问。
	e.Start(":8080")
}

