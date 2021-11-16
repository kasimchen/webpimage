package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/image/webp"
	"image/jpeg"
	"net/http"
	"github.com/gin-gonic/gin"
	"os"
	"runtime"
	_ "image/jpeg"
	"webpimage/common"
	"github.com/gin-contrib/cors"
	"github.com/rs/xid"

)




func WebpToImageHandler(c *gin.Context) {


	urlStr:=c.Query("url")
	resp, err := http.Get(urlStr)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": "-1","data":"","msg":"下载异常"})
		return
	}

	//body, err := ioutil.ReadAll(resp.Body)
	//
	//if err!=nil{
	//	c.JSON(http.StatusOK, gin.H{"code": "-1","data":"","msg":"读取异常"})
	//	return
	//}

	img, err := webp.Decode(resp.Body)

	if err!=nil{
		c.JSON(http.StatusOK, gin.H{"code": "-1","data":"","msg":"解析异常"})
		return
	}

	var name =xid.New().String()
	f, err := os.Create("image_"+name+".jpg")
	defer f.Close()
	if err = jpeg.Encode(f, img, &jpeg.Options{Quality:100}); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": "-1","data":"","msg":"生成异常"})
		return
	}


	c.File("image.jpg")
}


func getImageInfoHandler(c *gin.Context) {



	jsonData:=c.PostForm("json")

	if jsonData==""{
		c.JSON(http.StatusOK, gin.H{"code": "-1","data":"","msg":"缺少json参数"})
		return
	}


	var imageList = make([]string,0)
	err:=json.Unmarshal([]byte(jsonData),&imageList)

	if err!=nil{
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{"code": "-1","data":"","msg":err.Error()})
		return

	}

	type imageData struct {
		url string
		height int
		width int
		size int
		imageType string
		err string
	}


	runtime.GOMAXPROCS(runtime.NumCPU())
	ch := make(chan imageData)
	var count = len(imageList)
	for _,v:=range imageList   {

		go func(v string) {
			image1,size,imageType, err := common.GetImage(v)

			var image_data imageData
			errStr:=""
			if err!=nil{
				errStr=err.Error()
				 image_data =  imageData{height:0,width:0,size:size,imageType:imageType,err:errStr,url:v}
			}else{
				 image_data =  imageData{height:image1.Bounds().Dy(),width:image1.Bounds().Dx(),size:size,imageType:imageType,err:errStr,url:v}
			}

			ch <- image_data // 协程结束，发出信号
		}(v)
	}


	var imageMapList =make([]map[string]interface{},count)

	for v:= range ch {
		// 每次从ch中接收数据，表明一个活动的协程结束


		m := make(map[string]interface{})
		m["size"] = v.size
		m["imageType"] = v.imageType
		m["error"] = v.err
		m["width"] = v.width
		m["height"] = v.height
		m["url"] = v.url

		imageMapList[count-1] = m

		count--
		// 当所有活动的协程都结束时，关闭管道
		if count == 0 {
			close(ch)
		}
	}


	c.JSON(http.StatusOK, gin.H{"code": "200","data":imageMapList,"msg":"success"})
}

func main() {

	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response


	r.Use(cors.Default())

	r.POST("get_image_infos",getImageInfoHandler)
	r.GET("webp_to_image",WebpToImageHandler)

	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(":8000")


}


