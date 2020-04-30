package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

const   (
	VAL = 0x45FFFFFF
	INDEX = 0x0000003D
)
var(
	alphabet = []byte("qwertyuiopasdfghjklzxcvbnm1234567890QWERTYUIOPASDFGHJKLZXCVBNM")
)
func getMD5(str string)string  {
	m := md5.New()
	m.Write([]byte(str))
	c := m.Sum(nil)
	return hex.EncodeToString(c)
}

func TransForm(logurl string)([4]string,error)  {
	md5Str := getMD5(logurl)
	var tempVal int64
	var result [4]string
	var tempurl []byte
	for i:=0;i<4;i++{
		tempsubstr := md5Str[i*8:(i+1)*8]
		hexVal,err := strconv.ParseInt(tempsubstr,16,64)
		if err !=nil{
			return result,nil
		}
		tempVal = int64(VAL)&hexVal
		var index int64
		tempurl = []byte{}
		for i:=0;i<6;i++{
			index = INDEX&tempVal
			tempurl = append(tempurl,alphabet[index])
			tempVal = tempVal>>5
		}
		result[i] = string(tempurl)
	}
	return result,nil
}

func main()  {
	res, _ := TransForm("https://club.autohome.com.cn/bbs/thread/d59a41145c026a41/87249581-1.html#pvareaid=101650")
	fmt.Println(res)

	fmt.Println()
}
