package file

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"webb/service"
	"webb/service/AddUserMess"
)

func UploadHeadPhoto(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if err := r.ParseMultipartForm(10 << 20); err != nil {
			http.Error(w, "文件过大或格式错误", http.StatusBadRequest)
			return
		}

		// 限制内存使用并解析表单
		const maxUploadSize = 10 << 20
		if err := r.ParseMultipartForm(maxUploadSize); err != nil {
			fmt.Println("表单解析失败", err)
			http.Error(w, "文件过大或表单解析失败", http.StatusBadRequest)
			return
		}

		//处理头像上传
		file, header, err := r.FormFile("avatar")
		if err != nil {
			fmt.Println("name读取错误", err)
			return
		}
		defer func(file multipart.File) {
			err := file.Close()
			if err != nil {

			}
		}(file)
		fileName := header.Filename

		//生成唯一的文件名
		uniqueName := service.GenerateUniqueFileName(fileName)

		b, e := ioutil.ReadAll(file)
		if e != nil {
			fmt.Println("文件读取错误", err)
		}

		err = ioutil.WriteFile("./uploads/static/imgs/"+uniqueName, b, 0777)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("文件保存成功")
		//t.Execute(w, nil)

		//添加用户头像路径，保存到数据库当中
		c, err := r.Cookie("account")
		if err != nil {
			fmt.Println("获得当前账号失败", err)
			return
		}

		finalName := fmt.Sprintf("http://localhost:8080/uploads/static/imgs/%s", uniqueName)
		err = AddUserMess.AddAvatar(c.Value, finalName)
		if err != nil {
			fmt.Println("保存失败", err)
			return
		}
		fmt.Println("路径添加成功")

	}

}
