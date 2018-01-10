package mycrypto

import (
	"c3m/common/lzjs"
	"c3m/common/mystring"

	"github.com/tidusant/c3m-common/log"

	//	"c3m/log"
	"encoding/base64"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func StringRand(strlen int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}
func NumRand(min, max int) int {
	if max <= min {
		max = min + 1
	}
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min) + min
}
func Base64fix(s string) string {
	n := 4 - len(s)%4
	if n < 4 {
		for i := 0; i < n; i++ {
			s += "="
		}
	}
	return s
}
func CampaignDecode(data string) string {
	code := data[:len(data)/2]
	code = strings.Replace(data, code, "", 1) + code
	strbytes, _ := base64.StdEncoding.DecodeString(Base64fix(code))
	return string(strbytes)
}

func Encode(data string, oddnumber int, div int) string {
	if oddnumber > 9 {
		oddnumber = 9
	}
	var x = NumRand(2, oddnumber)
	var x2 = base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(x)))
	x2 = strings.Replace(x2, "=", "", -1)
	//x2b := []byte(x2)
	if div == 2 {
		data = lzjs.CompressToBase64(data)

	} else {
		data = base64.StdEncoding.EncodeToString([]byte(data))
	}

	data = strings.Replace(data, "=", "", -1)
	xstr := StringRand(x)

	data += xstr

	b := []byte(data)

	var l = int(math.Floor(float64(len(data) / div)))

	var result1 []byte
	var result2 []byte

	for i := len(b) - 1; i >= 0; i-- {

		if i%x == 0 {
			result1 = append(result1, b[i]) // string([]rune(data)[i])

		} else {
			result2 = append(result2, b[i])
		}
	}

	strb64 := string(result1) + string(result2)
	strb64 = strb64[:l] + x2 + strb64[l:]

	return strb64
}

func DecodeBK(data string, keysalt string) string {
	keysalt = base64.StdEncoding.EncodeToString([]byte(keysalt))
	keysalt = strings.Replace(keysalt, "=", "", -1)
	data = strings.Replace(data, keysalt, "", 1)
	data = Base64fix(data)
	//byteDecode, _ := lzjs.DecompressFromBase64(data)
	byteDecode, _ := base64.StdEncoding.DecodeString(data)
	data = string(byteDecode)
	return data
}

func Decode(data string) string {
	if len(data) < 10 {
		log.Errorf("cannot decode %s", data)
		return data
	}
	x := 10
	xstr := data[:x]
	data = data[x:]
	data, _ = lzjs.DecompressFromBase64(data)
	xb64 := base64.StdEncoding.EncodeToString([]byte(xstr))
	xb64 = strings.Replace(xb64, "=", "", -1)
	data = strings.Replace(data, xb64, "", 1)
	return data
}

func DecodeOld(code string) string {
	if code == "" {
		return code
	}
	var rt string = ""
	key := code
	//key = "kZXUuYkRWzUgQk92YoNwRdh92Q3SZtFmb9Wa0NW"
	if key == rt {
		return rt
	}

	oddstr := "d"
	l := int(math.Floor((float64)(len(key)-2) / 2))
	num := key[l : l+2]

	key = key[:l] + key[l+2:]

	byteDecode, _ := base64.StdEncoding.DecodeString(Base64fix(num))
	num = string(byteDecode)

	floatNum, _ := strconv.ParseFloat(num, 64)
	intNum := (int)(floatNum)
	if intNum > 0 {
		//print_r($num);print_r("\r\n");
		//get odd string
		lf := math.Ceil((float64)(len(key)) / floatNum)
		oddstr = key[:int(lf)]
		ukey := strings.Replace(key, oddstr, "", 1)
		base64str := ""

		for i := len(oddstr) - 1; i >= 0; i-- {
			base64str += string(oddstr[len(oddstr)-1:])
			oddstr = oddstr[:len(oddstr)-1]
			if len(ukey)-intNum+1 > 0 {
				base64str += mystring.Reverse(string(ukey[len(ukey)-intNum+1:]))
			} else {
				base64str += mystring.Reverse(ukey)
			}
			if i > 0 {
				ukey = ukey[:len(ukey)-intNum+1]
			}
		}
		base64str = base64str[:len(base64str)-intNum]

		//log.Debugf("lzjs %s", base64str)
		//log.Debugf("lzjs %s", Base64fix(base64str))
		//byteDecode, _ := base64.StdEncoding.DecodeString(Base64fix(base64str))
		//base64str = Base64fix(base64str)
		//byteDecode, _ := lzjs.DecompressFromBase64(base64str)
		//data, _ := base64.StdEncoding.DecodeString(base64str)
		byteDecode, _ := lzjs.DecompressFromBase64(base64str)

		rt = string(byteDecode)
		//log.Debugf("data decompress %s", rt)
		//rt, _ = url.QueryUnescape(rt)
		//log.Debugf("data decompress %s", rt)
	}
	return rt
}

//encode for wapi
func EncodeW(data string) string {

	var x2 = base64.StdEncoding.EncodeToString([]byte(data))
	x2 = strings.Replace(x2, "=", "", -1)
	return x2
}

//decode for wapi
func DecodeW(code string) string {
	if code == "" {
		return code
	}
	var rt string = ""
	key := code
	//key = "kZXUuYkRWzUgQk92YoNwRdh92Q3SZtFmb9Wa0NW"
	if key == rt {
		return rt
	}

	oddstr := "d"
	l := int(math.Floor((float64)(len(key)-2) / 2))
	num := key[l : l+2]

	key = key[:l] + key[l+2:]

	byteDecode, _ := base64.StdEncoding.DecodeString(Base64fix(num))
	num = string(byteDecode)

	floatNum, _ := strconv.ParseFloat(num, 64)
	intNum := (int)(floatNum)
	if intNum > 0 {
		//print_r($num);print_r("\r\n");
		//get odd string
		lf := math.Ceil((float64)(len(key)) / floatNum)
		oddstr = key[:int(lf)]
		ukey := strings.Replace(key, oddstr, "", 1)
		base64str := ""

		for i := len(oddstr) - 1; i >= 0; i-- {
			base64str += string(oddstr[len(oddstr)-1:])
			oddstr = oddstr[:len(oddstr)-1]
			if len(ukey)-intNum+1 > 0 {
				base64str += mystring.Reverse(string(ukey[len(ukey)-intNum+1:]))
			} else {
				base64str += mystring.Reverse(ukey)
			}
			if i > 0 {
				ukey = ukey[:len(ukey)-intNum+1]
			}
		}
		base64str = base64str[:len(base64str)-intNum]

		byteDecode, _ := base64.StdEncoding.DecodeString(base64str)
		//byteDecode, _ := lzjs.DecompressFromBase64(base64str)

		rt = string(byteDecode)

	}
	return rt
}
