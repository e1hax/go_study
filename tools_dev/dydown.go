package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func downloadFile(filepath string, url string) error {
	//创建一个 http 请求，返回数据resp
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		fmt.Println("下载出错！")
		return err
	}

	//函数执行完毕后关闭http连接
	defer resp.Body.Close()

	//创建文件
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	//函数执行完毕后关闭out
	defer out.Close()

	//写入文件
	_, err = io.Copy(out, resp.Body)
	return err
}

func main() {
	fileurl := "https://v26-web.douyinvod.com/2d8f3639ade29b801832bded7a584b85/6346dae0/video/tos/cn/tos-cn-ve-15c001-alinc2/7e085e9f295f421890fba5dd38c008a7/?a=6383&ch=0&cr=0&dr=0&cd=0%7C0%7C0%7C0&cv=1&br=282&bt=282&cs=0&ds=4&ft=LjhJEL998xI7uEPmD0P5fQhlpPiXMviAOVJEShSqRCPD-Ani&mime_type=video_mp4&qs=0&rc=NzRpOGg6OTNmM2UzOjtpO0BpanM8dWU6Zmc2ZjMzNGkzM0A2YV5hNi9eNmAxMS8yNmIyYSNkL28xcjQwMXBgLS1kLS9zcw%3D%3D"
	filePath := "D:\\User\\e1hax\\视频\\new1.mp4"
	err := downloadFile(filePath, fileurl)
	if err == nil {
		fmt.Println("下载成功，文件地址为：", filePath)
	}
}
