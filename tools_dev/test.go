package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 文件url需要修改成目标地址
	fileUrl := "https://www.douyin.com/video/7153160094094494989"
	err := DownloadFile("22.mp4", fileUrl)
	if err != nil {
		panic(err)
	}
	fmt.Println("Downloaded: " + fileUrl)
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println(resp.Body)

	// Create the file
	//out, err := os.Create(filepath)
	//if err != nil {
	//	return err
	//}
	//defer out.Close()

	// Write the body to file
	//_, err = io.Copy(out, resp.Body)
	return err

}
