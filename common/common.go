package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/disintegration/imageorient"
	"golang.org/x/image/webp"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	"strings"
	"time"
)


func GetImage(url string) (image.Image, int,string, error) {


	c := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := c.Get(url)
	if err != nil {

		return nil,0,"", errors.New("请求错误:"+","+err.Error())
	}



	defer resp.Body.Close()


	if strings.Contains(url,"webp"){


		//body, err := ioutil.ReadAll(resp.Body)
		//
		//if err!=nil{
		//	return nil,0,"", err
		//}

		img, err := webp.Decode(resp.Body)

		if err!=nil{
			return nil,0,"", err
		}

		return img,int(resp.ContentLength),"webp", nil

	}else{

		img, itype, err := imageorient.Decode(resp.Body)


		if err!=nil{
			return nil,0,"", err
		}


		return img,int(resp.ContentLength),itype, nil

	}


}





func GetLocation() *time.Location  {
	var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
	return cstSh
}



func IntTimeToStr(timeUnix int64) string {
	var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
	formatTimeStr := time.Unix(timeUnix, 0).In(cstSh).Format("2006-01-02 15:04:05")
	return formatTimeStr
}


func Now() int64 {

	var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
	unix := time.Now().In(cstSh).Unix()
	return unix

}


func NowStr() string {

	var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
	unix := time.Now().In(cstSh).Unix()
	formatTimeStr := time.Unix(unix, 0).In(cstSh).Format("2006-01-02 15:04:05")
	return formatTimeStr

}




// Convert map json string
func MapToJson(m map[string]interface{}) (string, error) {
	jsonByte, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("Marshal with error: %+v\n", err)
		return "", nil
	}

	return string(jsonByte), nil
}

func JsonToMap(str string) map[string]interface{} {

	var tempMap map[string]interface{}

	err := json.Unmarshal([]byte(str), &tempMap)

	if err != nil {
		panic(err)
	}

	return tempMap
}