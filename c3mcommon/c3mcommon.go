package c3mcommon

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/tidusant/c3m-common/log"
	"github.com/tidusant/c3m-common/mycrypto"
	"github.com/tidusant/c3m-common/mystring"

	"github.com/nfnt/resize"
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2"
)

var (
	db              map[string]*mgo.Database
	listCountryFlag map[string]string
	listLocale      map[string]string
	listCountry     map[string]string
)

func init() {
	log.Errorf("init common")
	//Config
	viper.SetConfigName("config") // no need to include file extension
	viper.AddConfigPath(".")      // set the  path of your config file

	err := viper.ReadInConfig()
	if !CheckError("Config file not found...", err) {
		os.Exit(1)
	}
	initListLocale()
	initCountryFlag()
	initListCountry()
}
func initListLocale() {
	listLocale = make(map[string]string)
	listLocale["af"] = "af_ZA"
	listLocale["ak"] = "ak_GH"
	listLocale["am"] = "am_ET"
	listLocale["ar"] = "ar_AR"
	listLocale["as"] = "as_IN"
	listLocale["ay"] = "ay_BO"
	listLocale["az"] = "az_AZ"
	listLocale["be"] = "be_BY"
	listLocale["bg"] = "bg_BG"
	listLocale["bn"] = "bn_IN"
	listLocale["br"] = "br_FR"
	listLocale["bs"] = "bs_BA"
	listLocale["ca"] = "ca_ES"
	listLocale["cb"] = "cb_IQ"
	listLocale["ck"] = "ck_US"
	listLocale["co"] = "co_FR"
	listLocale["cs"] = "cs_CZ"
	listLocale["cx"] = "cx_PH"
	listLocale["cy"] = "cy_GB"
	listLocale["da"] = "da_DK"
	listLocale["de"] = "de_DE"
	listLocale["el"] = "el_GR"
	listLocale["en"] = "en_GB"
	listLocale["eo"] = "eo_EO"

	listLocale["es"] = "es_ES"

	listLocale["et"] = "et_EE"
	listLocale["eu"] = "eu_ES"
	listLocale["fa"] = "fa_IR"
	listLocale["fb"] = "fb_LT"
	listLocale["ff"] = "ff_NG"
	listLocale["fi"] = "fi_FI"
	listLocale["fo"] = "fo_FO"
	listLocale["fr"] = "fr_FR"
	listLocale["fy"] = "fy_NL"
	listLocale["ga"] = "ga_IE"
	listLocale["gl"] = "gl_ES"
	listLocale["gn"] = "gn_PY"
	listLocale["gu"] = "gu_IN"
	listLocale["gx"] = "gx_GR"
	listLocale["ha"] = "ha_NG"
	listLocale["he"] = "he_IL"
	listLocale["hi"] = "hi_IN"
	listLocale["hr"] = "hr_HR"
	listLocale["ht"] = "ht_HT"
	listLocale["hu"] = "hu_HU"
	listLocale["hy"] = "hy_AM"
	listLocale["id"] = "id_ID"
	listLocale["ig"] = "ig_NG"
	listLocale["is"] = "is_IS"
	listLocale["it"] = "it_IT"
	listLocale["ja"] = "ja_JP"
	listLocale["jv"] = "jv_ID"
	listLocale["ka"] = "ka_GE"
	listLocale["kk"] = "kk_KZ"
	listLocale["km"] = "km_KH"
	listLocale["kn"] = "kn_IN"
	listLocale["ko"] = "ko_KR"
	listLocale["ku"] = "ku_TR"
	listLocale["ky"] = "ky_KG"
	listLocale["la"] = "la_VA"
	listLocale["lg"] = "lg_UG"
	listLocale["li"] = "li_NL"
	listLocale["ln"] = "ln_CD"
	listLocale["lo"] = "lo_LA"
	listLocale["lt"] = "lt_LT"
	listLocale["lv"] = "lv_LV"
	listLocale["mg"] = "mg_MG"
	listLocale["mi"] = "mi_NZ"
	listLocale["mk"] = "mk_MK"
	listLocale["ml"] = "ml_IN"
	listLocale["mn"] = "mn_MN"
	listLocale["mr"] = "mr_IN"
	listLocale["ms"] = "ms_MY"
	listLocale["mt"] = "mt_MT"
	listLocale["my"] = "my_MM"
	listLocale["nb"] = "nb_NO"
	listLocale["nd"] = "nd_ZW"
	listLocale["ne"] = "ne_NP"

	listLocale["nl"] = "nl_NL"
	listLocale["nn"] = "nn_NO"
	listLocale["ny"] = "ny_MW"
	listLocale["or"] = "or_IN"
	listLocale["pa"] = "pa_IN"
	listLocale["pl"] = "pl_PL"
	listLocale["ps"] = "ps_AF"
	listLocale["pt"] = "pt_PT"
	listLocale["qc"] = "qc_GT"
	listLocale["qu"] = "qu_PE"
	listLocale["rm"] = "rm_CH"
	listLocale["ro"] = "ro_RO"
	listLocale["ru"] = "ru_RU"
	listLocale["rw"] = "rw_RW"
	listLocale["sa"] = "sa_IN"
	listLocale["sc"] = "sc_IT"
	listLocale["se"] = "se_NO"
	listLocale["si"] = "si_LK"
	listLocale["sk"] = "sk_SK"
	listLocale["sl"] = "sl_SI"
	listLocale["sn"] = "sn_ZW"
	listLocale["so"] = "so_SO"
	listLocale["sq"] = "sq_AL"
	listLocale["sr"] = "sr_RS"
	listLocale["sv"] = "sv_SE"
	listLocale["sw"] = "sw_KE"
	listLocale["sy"] = "sy_SY"
	listLocale["sz"] = "sz_PL"
	listLocale["ta"] = "ta_IN"
	listLocale["te"] = "te_IN"
	listLocale["tg"] = "tg_TJ"
	listLocale["th"] = "th_TH"
	listLocale["tk"] = "tk_TM"

	listLocale["tl"] = "tl_ST"
	listLocale["tr"] = "tr_TR"
	listLocale["tt"] = "tt_RU"
	listLocale["tz"] = "tz_MA"
	listLocale["uk"] = "uk_UA"
	listLocale["ur"] = "ur_PK"
	listLocale["uz"] = "uz_UZ"
	listLocale["vi"] = "vi_VN"
	listLocale["wo"] = "wo_SN"
	listLocale["xh"] = "xh_ZA"
	listLocale["yi"] = "yi_DE"
	listLocale["yo"] = "yo_NG"
	listLocale["zh"] = "zh_CN"
	listLocale["zu"] = "zu_ZA"
	listLocale["zz"] = "zz_TR"
}
func initListCountry() {
	listCountry = make(map[string]string)
	listCountry["af"] = "Afrikaans"
	listCountry["ak"] = "Akan"
	listCountry["am"] = "Amharic"
	listCountry["ar"] = "Arabic"
	listCountry["as"] = "Assamese"
	listCountry["ay"] = "Aymara"
	listCountry["az"] = "Azerbaijani"
	listCountry["be"] = "Belarusian"
	listCountry["bg"] = "Bulgarian"
	listCountry["bn"] = "Bengali"
	listCountry["br"] = "Breton"
	listCountry["bs"] = "Bosnian"
	listCountry["ca"] = "Catalan"
	listCountry["cb"] = "Sorani Kurdish"
	listCountry["ck"] = "Cherokee"
	listCountry["co"] = "Corsican"
	listCountry["cs"] = "Czech"
	listCountry["cx"] = "Cebuano"
	listCountry["cy"] = "Welsh"
	listCountry["da"] = "Danish"
	listCountry["de"] = "German"
	listCountry["el"] = "Greek"
	listCountry["en"] = "English (US)"
	listCountry["eo"] = "Esperanto"
	listCountry["es"] = "Spanish (Venezuela)"
	listCountry["et"] = "Estonian"
	listCountry["eu"] = "Basque"
	listCountry["fa"] = "Persian"
	listCountry["fb"] = "Leet Speak"
	listCountry["ff"] = "Fulah"
	listCountry["fi"] = "Finnish"
	listCountry["fo"] = "Faroese"
	listCountry["fr"] = "French"
	listCountry["fy"] = "Frisian"
	listCountry["ga"] = "Irish"
	listCountry["gl"] = "Galician"
	listCountry["gn"] = "Guarani"
	listCountry["gu"] = "Gujarati"
	listCountry["gx"] = "Classical Greek"
	listCountry["ha"] = "Hausa"
	listCountry["he"] = "Hebrew"
	listCountry["hi"] = "Hindi"
	listCountry["hr"] = "Croatian"
	listCountry["ht"] = "Haitian Creole"
	listCountry["hu"] = "Hungarian"
	listCountry["hy"] = "Armenian"
	listCountry["id"] = "Indonesian"
	listCountry["ig"] = "Igbo"
	listCountry["is"] = "Icelandic"
	listCountry["it"] = "Italian"
	listCountry["ja"] = "Japanese"
	listCountry["jv"] = "Javanese"
	listCountry["ka"] = "Georgian"
	listCountry["kk"] = "Kazakh"
	listCountry["km"] = "Khmer"
	listCountry["kn"] = "Kannada"
	listCountry["ko"] = "Korean"
	listCountry["ku"] = "Kurdish (Kurmanji)"
	listCountry["ky"] = "Kyrgyz"
	listCountry["la"] = "Latin"
	listCountry["lg"] = "Ganda"
	listCountry["li"] = "Limburgish"
	listCountry["ln"] = "Lingala"
	listCountry["lo"] = "Lao"
	listCountry["lt"] = "Lithuanian"
	listCountry["lv"] = "Latvian"
	listCountry["mg"] = "Malagasy"
	listCountry["mi"] = "Māori"
	listCountry["mk"] = "Macedonian"
	listCountry["ml"] = "Malayalam"
	listCountry["mn"] = "Mongolian"
	listCountry["mr"] = "Marathi"
	listCountry["ms"] = "Malay"
	listCountry["mt"] = "Maltese"
	listCountry["my"] = "Burmese"
	listCountry["nb"] = "Norwegian (bokmal)"
	listCountry["nd"] = "Ndebele"
	listCountry["ne"] = "Nepali"

	listCountry["nl"] = "Dutch"
	listCountry["nn"] = "Norwegian (nynorsk)"
	listCountry["ny"] = "Chewa"
	listCountry["or"] = "Oriya"
	listCountry["pa"] = "Punjabi"
	listCountry["pl"] = "Polish"
	listCountry["ps"] = "Pashto"
	listCountry["pt"] = "Portuguese (Brazil)"
	listCountry["qc"] = "Quiché"
	listCountry["qu"] = "Quechua"
	listCountry["rm"] = "Romansh"
	listCountry["ro"] = "Romanian"
	listCountry["ru"] = "Russian"
	listCountry["rw"] = "Kinyarwanda"
	listCountry["sa"] = "Sanskrit"
	listCountry["sc"] = "Sardinian"
	listCountry["se"] = "Northern Sámi"
	listCountry["si"] = "Sinhala"
	listCountry["sk"] = "Slovak"
	listCountry["sl"] = "Slovenian"
	listCountry["sn"] = "Shona"
	listCountry["so"] = "Somali"
	listCountry["sq"] = "Albanian"
	listCountry["sr"] = "Serbian"
	listCountry["sv"] = "Swedish"
	listCountry["sw"] = "Swahili"
	listCountry["sy"] = "Syriac"
	listCountry["sz"] = "Silesian"
	listCountry["ta"] = "Tamil"
	listCountry["te"] = "Telugu"
	listCountry["tg"] = "Tajik"
	listCountry["th"] = "Thai"
	listCountry["tk"] = "Turkmen"

	listCountry["tl"] = "Klingon"
	listCountry["tr"] = "Turkish"
	listCountry["tt"] = "Tatar"
	listCountry["tz"] = "Tamazight"
	listCountry["uk"] = "Ukrainian"
	listCountry["ur"] = "Urdu"
	listCountry["uz"] = "Uzbek"
	listCountry["vi"] = "Tiếng Việt"
	listCountry["wo"] = "Wolof"
	listCountry["xh"] = "Xhosa"
	listCountry["yi"] = "Yiddish"
	listCountry["yo"] = "Yoruba"
	listCountry["zh"] = "Simplified Chinese (China)"
	listCountry["zu"] = "Zulu"
	listCountry["zz"] = "Zazaki"
}
func initCountryFlag() {
	listCountryFlag = make(map[string]string)
	listCountryFlag["af"] = "za"
	listCountryFlag["ak"] = "gh"
	listCountryFlag["am"] = "et"
	listCountryFlag["ar"] = "ar"
	listCountryFlag["as"] = "in"
	listCountryFlag["ay"] = "bo"
	listCountryFlag["az"] = "az"
	listCountryFlag["be"] = "by"
	listCountryFlag["bg"] = "bg"
	listCountryFlag["bn"] = "in"
	listCountryFlag["br"] = "fr"
	listCountryFlag["bs"] = "ba"
	listCountryFlag["ca"] = "es"
	listCountryFlag["cb"] = "iq"
	listCountryFlag["ck"] = "us"
	listCountryFlag["co"] = "fr"
	listCountryFlag["cs"] = "cz"
	listCountryFlag["cx"] = "ph"
	listCountryFlag["cy"] = "gb"
	listCountryFlag["da"] = "dk"
	listCountryFlag["de"] = "de"
	listCountryFlag["el"] = "gr"
	listCountryFlag["en"] = "gb"
	listCountryFlag["eo"] = "eo"
	listCountryFlag["es"] = "cl"
	listCountryFlag["es"] = "es"
	listCountryFlag["et"] = "ee"
	listCountryFlag["eu"] = "es"
	listCountryFlag["fa"] = "ir"
	listCountryFlag["fb"] = "lt"
	listCountryFlag["ff"] = "ng"
	listCountryFlag["fi"] = "fi"
	listCountryFlag["fo"] = "fo"
	listCountryFlag["fr"] = "fr"
	listCountryFlag["fy"] = "nl"
	listCountryFlag["ga"] = "ie"
	listCountryFlag["gl"] = "es"
	listCountryFlag["gn"] = "py"
	listCountryFlag["gu"] = "in"
	listCountryFlag["gx"] = "gr"
	listCountryFlag["ha"] = "ng"
	listCountryFlag["he"] = "il"
	listCountryFlag["hi"] = "in"
	listCountryFlag["hr"] = "hr"
	listCountryFlag["ht"] = "ht"
	listCountryFlag["hu"] = "hu"
	listCountryFlag["hy"] = "am"
	listCountryFlag["id"] = "id"
	listCountryFlag["ig"] = "ng"
	listCountryFlag["is"] = "is"
	listCountryFlag["it"] = "it"
	listCountryFlag["ja"] = "jp"
	listCountryFlag["jv"] = "id"
	listCountryFlag["ka"] = "ge"
	listCountryFlag["kk"] = "kz"
	listCountryFlag["km"] = "kh"
	listCountryFlag["kn"] = "in"
	listCountryFlag["ko"] = "kr"
	listCountryFlag["ku"] = "tr"
	listCountryFlag["ky"] = "kg"
	listCountryFlag["la"] = "va"
	listCountryFlag["lg"] = "ug"
	listCountryFlag["li"] = "nl"
	listCountryFlag["ln"] = "cd"
	listCountryFlag["lo"] = "la"
	listCountryFlag["lt"] = "lt"
	listCountryFlag["lv"] = "lv"
	listCountryFlag["mg"] = "mg"
	listCountryFlag["mi"] = "nz"
	listCountryFlag["mk"] = "mk"
	listCountryFlag["ml"] = "in"
	listCountryFlag["mn"] = "mn"
	listCountryFlag["mr"] = "in"
	listCountryFlag["ms"] = "my"
	listCountryFlag["mt"] = "mt"
	listCountryFlag["my"] = "mm"
	listCountryFlag["nb"] = "no"
	listCountryFlag["nd"] = "zw"
	listCountryFlag["ne"] = "np"
	listCountryFlag["nl"] = "nl"
	listCountryFlag["nn"] = "no"
	listCountryFlag["ny"] = "mw"
	listCountryFlag["or"] = "in"
	listCountryFlag["pa"] = "in"
	listCountryFlag["pl"] = "pl"
	listCountryFlag["ps"] = "af"
	listCountryFlag["pt"] = "pt"
	listCountryFlag["qc"] = "gt"
	listCountryFlag["qu"] = "pe"
	listCountryFlag["rm"] = "ch"
	listCountryFlag["ro"] = "ro"
	listCountryFlag["ru"] = "ru"
	listCountryFlag["rw"] = "rw"
	listCountryFlag["sa"] = "in"
	listCountryFlag["sc"] = "it"
	listCountryFlag["se"] = "no"
	listCountryFlag["si"] = "lk"
	listCountryFlag["sk"] = "sk"
	listCountryFlag["sl"] = "si"
	listCountryFlag["sn"] = "zw"
	listCountryFlag["so"] = "so"
	listCountryFlag["sq"] = "al"
	listCountryFlag["sr"] = "rs"
	listCountryFlag["sv"] = "se"
	listCountryFlag["sw"] = "ke"
	listCountryFlag["sy"] = "sy"
	listCountryFlag["sz"] = "pl"
	listCountryFlag["ta"] = "in"
	listCountryFlag["te"] = "in"
	listCountryFlag["tg"] = "tj"
	listCountryFlag["th"] = "th"
	listCountryFlag["tk"] = "tm"
	listCountryFlag["tl"] = "st"
	listCountryFlag["tr"] = "tr"
	listCountryFlag["tt"] = "ru"
	listCountryFlag["tz"] = "ma"
	listCountryFlag["uk"] = "ua"
	listCountryFlag["ur"] = "pk"
	listCountryFlag["uz"] = "uz"
	listCountryFlag["vi"] = "vn"
	listCountryFlag["wo"] = "sn"
	listCountryFlag["xh"] = "za"
	listCountryFlag["yi"] = "de"
	listCountryFlag["yo"] = "ng"
	listCountryFlag["zh"] = "cn"
	listCountryFlag["zu"] = "za"
	listCountryFlag["zz"] = "tr"
}

func ConnectDB(dbname string) (db *mgo.Database, strErr string) {
	if viper.GetString("db"+dbname+".h") == "" {
		strErr = "invalid db " + dbname + ", viper"
	}
	mongoDBDialInfo := mgo.DialInfo{
		Addrs:    []string{viper.GetString("db" + dbname + ".h")},
		Timeout:  60 * time.Second,
		Database: viper.GetString("db" + dbname + ".d"),
		Username: viper.GetString("db" + dbname + ".u"),
		Password: viper.GetString("db" + dbname + ".p"),
	}

	mongoSession, err := mgo.DialWithInfo(&mongoDBDialInfo)

	if CheckError("error when connect db: %s\n", err) {
		mongoSession.SetMode(mgo.Monotonic, true)
		db = mongoSession.DB(viper.GetString("db" + dbname + ".d"))
	}
	if err != nil {
		strErr = err.Error()
	}
	return db, strErr
}

//func CheckSession(s string) bool {
//	col := db["session"].C("sessions")

//	var result models.Session
//	err2 := col.Find(bson.M{"uid": s}).One(&result)

//	if err2 != nil {
//		log.Infof("Session not found uid '%s': %s\n", s, err2)
//	} else {
//		if result.Expired > time.Now().Unix() {
//			return true
//		} else {
//			log.Infof("Session expired: uid '%s'", s)
//			return false
//		}

//	}
//	return false
//}
//func InitDb(dbnames []string) bool {
//	for _, dbname := range dbnames {
//		if viper.GetString("db"+dbname+".h") == "" {
//			log.Errorf("invalid db " + dbname)
//			return false
//		}
//		mongoDBDialInfo := mgo.DialInfo{
//			Addrs:    []string{viper.GetString("db" + dbname + ".h")},
//			Timeout:  60 * time.Second,
//			Database: viper.GetString("db" + dbname + ".d"),
//			Username: viper.GetString("db" + dbname + ".u"),
//			Password: viper.GetString("db" + dbname + ".p"),
//		}

//		mongoSession, err := mgo.DialWithInfo(&mongoDBDialInfo)

//		if CheckError("error when connect db: %s\n", err) {
//			mongoSession.SetMode(mgo.Monotonic, true)
//			db[dbname] = mongoSession.DB(viper.GetString("db" + dbname + ".d"))

//		} else {
//			return false
//		}

//	}
//	return true
//}
func ReturnJsonMessage(status, strerr, strmsg, data string) string {
	if data == "" {
		data = "{}"
	}
	return "{\"status\":\"" + status + "\",\"error\":\"" + strerr + "\",\"message\":\"" + strmsg + "\",\"data\":" + data + "}"
}
func FileCount(path string) int {
	i := 0
	files, err := ioutil.ReadDir(path)
	if err != nil {
		CheckError(path+" filecount error", err)
		return 0
	}
	for _, file := range files {
		if file.IsDir() {
			i += FileCount(path + "/" + file.Name())
		} else {
			i++
		}

	}
	return i
}

//func CheckRequest(uri, useragent, referrer, remoteAddress string) bool {

//	col := db["session"].C("requestUrls")
//	log.Printf("now: %d , check: %d", int(time.Now().Unix()), int(time.Now().Unix())-10)
//	urlcount, err := col.Find(bson.M{"uri": uri, "created": bson.M{"$gt": int(time.Now().Unix()) - 1}}).Count()
//	if CheckError("checkRequest", err) {
//		if urlcount == 0 {
//			//check ip in 3 sec
//			urlcount, err := col.Find(bson.M{"remoteAddress": remoteAddress, "created": bson.M{"$gt": int(time.Now().Unix()) - 3}}).Count()
//			if urlcount < 50 {
//				err = col.Insert(bson.M{"uri": uri, "created": int(time.Now().Unix()), "user-agent": useragent, "referer": referrer, "remoteAddress": remoteAddress})
//				CheckError("checkRequest Insert", err)
//				return true
//			}

//		}
//	}
//	return false
//}
func CheckError(msg string, err error) bool {
	if err != nil {
		log.Debugf(msg+": %s", err.Error())
		return false
	}
	return true
}
func ImgResize(imagebytes []byte, w, h uint) ([]byte, string) {
	filetype := http.DetectContentType(imagebytes[:512])
	r := bytes.NewReader(imagebytes)
	imagecontent, _, err := image.Decode(r)
	m := resize.Resize(w, h, imagecontent, resize.NearestNeighbor)
	if err != nil {
		return nil, ""
	}
	var buf bytes.Buffer
	wr := io.Writer(&buf)

	if filetype == "image/jpeg" {
		jpeg.Encode(wr, m, nil)
	} else if filetype == "image/gif" {
		gif.Encode(wr, m, nil)
	} else if filetype == "image/png" {
		png.Encode(wr, m)
	}

	return buf.Bytes(), strings.Replace(filetype, "image/", "", 1)
}

//func GetShop(userid, shopid string) models.Shop {
//	coluser := db["cuahang"].C("addons_shops")
//	var shop models.Shop
//	coluser.Find(bson.M{"_id": bson.ObjectIdHex(shopid), "clientid": bson.ObjectIdHex(userid)}).One(&shop)
//	return shop
//}

//func UpdateAlbum(shop models.Shop) models.Shop {
//	coluser := db["cuahang"].C("addons_shops")

//	cond := bson.M{"_id": shop.ID}
//	change := bson.M{"$set": bson.M{"albums": shop.Albums}}

//	coluser.Update(cond, change)
//	return shop
//}

func CheckDomain(requestDomain string) string {

	domainallow := strings.Split(viper.GetString("config.domainallow"), ",")
	for i := 0; i < len(domainallow); i++ {
		log.Debugf("%s - %s", domainallow[i], requestDomain)
		if domainallow[i] == requestDomain {
			return requestDomain
			break
		}
	}
	return ""
}

func Fake64() string {
	return mystring.RandString(100)
}

func Code2Flag(code string) string {
	return listCountryFlag[code]
}

// Check if a port is available
func CheckPort(port int) (status bool) {

	// Concatenate a colon and the port
	host := ":" + strconv.Itoa(port)

	// Try to create a server with the port
	server, err := net.Listen("tcp", host)

	// if it fails then the port is likely taken
	if err != nil {
		return false
	}

	// close the server
	server.Close()

	// we successfully used and closed the port
	// so it's now available to be used again
	return true
}

func FolderExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func CopyFile(source string, dest string) (err error) {
	sourcefile, err := os.Open(source)
	if err != nil {
		return err
	}

	defer sourcefile.Close()

	destfile, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer destfile.Close()

	_, err = io.Copy(destfile, sourcefile)
	if err == nil {
		sourceinfo, err := os.Stat(source)
		if err != nil {
			err = os.Chmod(dest, sourceinfo.Mode())
		}

	}

	return
}

func CopyDir(source string, dest string) (err error) {

	// get properties of source dir
	sourceinfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	// create dest dir

	err = os.MkdirAll(dest, sourceinfo.Mode())
	if err != nil {
		return err
	}

	directory, _ := os.Open(source)

	objects, err := directory.Readdir(-1)

	for _, obj := range objects {

		sourcefilepointer := source + "/" + obj.Name()

		destinationfilepointer := dest + "/" + obj.Name()

		if obj.IsDir() {
			// create sub-directories - recursively
			err = CopyDir(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			// perform copy
			err = CopyFile(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
	return
}

func JSMinify(content string) string {
	data := url.Values{}
	data.Add("data", content)
	rtstr := RequestUrl(viper.GetString("config.minify"), "POST", data)
	if strings.Index(rtstr, "ERROR!!!") >= 0 {
		log.Debugf("JSMinify Fail: %s", rtstr)
		return ""
	}
	return rtstr

}

func RequestUrl(url, method string, data url.Values) string {
	var rsp *http.Response
	var err error
	if strings.ToLower(method) == "post" {
		rsp, err = http.PostForm(url, data)
		if !CheckError("request api", err) {
			return ""
		}
	} else {
		rsp, err = http.Get(url + "?" + data.Encode())
		if !CheckError("request api", err) {
			return ""
		}
	}

	defer rsp.Body.Close()
	rtbyte, err := ioutil.ReadAll(rsp.Body)
	CheckError("request read data", err)
	rtstr := string(rtbyte)
	return rtstr
}

func RequestService(serviceurl string, data url.Values) string {
	payloadBytes, err := json.Marshal(data)
	body := bytes.NewReader(payloadBytes)
	log.Debugf("request service url: %s", viper.GetString("config.ServiceUrl")+serviceurl)
	req, err := http.NewRequest("POST", viper.GetString("config.ServiceUrl")+serviceurl, body)
	if !CheckError("request api", err) {
		return ""
	}
	//req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if !CheckError("request api", err) {

		return ""
	}

	defer resp.Body.Close()

	bodyresp, _ := ioutil.ReadAll(resp.Body)
	bodystr := string(bodyresp)

	if bodystr == "" {
		return ""
	}

	log.Debugf("response: %s", bodystr)
	bodystr = mycrypto.Decode4(bodystr)
	return bodystr
}

func RemoveHTMLComments(content []byte) []byte {
	// https://www.google.com/search?q=regex+html+comments
	// http://stackoverflow.com/a/1084759
	htmlcmt := regexp.MustCompile(`<!--[^>]*-->`)
	return htmlcmt.ReplaceAll(content, []byte(""))
}

func MinifyHTML(html []byte) string {
	// read line by line
	minifiedHTML := ""
	scanner := bufio.NewScanner(bytes.NewReader(RemoveHTMLComments(html)))
	for scanner.Scan() {
		// all leading and trailing white space of each line are removed
		lineTrimmed := strings.TrimSpace(scanner.Text())
		minifiedHTML += lineTrimmed
		if len(lineTrimmed) > 0 {
			// in case of following trimmed line:
			// <div id="foo"
			minifiedHTML += " "
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return minifiedHTML
}

//minify css
func RemoveCStyleComments(content []byte) []byte {
	// http://blog.ostermiller.org/find-comment
	ccmt := regexp.MustCompile(`/\*([^*]|[\r\n]|(\*+([^*/]|[\r\n])))*\*+/`)
	return ccmt.ReplaceAll(content, []byte(""))
}

func RemoveCppStyleComments(content []byte) []byte {
	cppcmt := regexp.MustCompile(`//.*`)
	return cppcmt.ReplaceAll(content, []byte(""))
}

func MinifyCSS(csscontent []byte) string {

	cssAllNoComments := RemoveCStyleComments(csscontent)

	// read line by line
	minifiedCss := ""
	scanner := bufio.NewScanner(bytes.NewReader(cssAllNoComments))
	for scanner.Scan() {
		// all leading and trailing white space of each line are removed
		minifiedCss += strings.TrimSpace(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return minifiedCss
}

func CreateImageFile(path, b64content string) error {
	unbased, err := base64.StdEncoding.DecodeString(b64content)
	if err != nil {
		log.Debugf("Cannot decode b64  %s", err)
		return err
	}
	r := bytes.NewReader(unbased)

	if filepath.Ext(path) == ".png" {
		im, err := png.Decode(r)
		if err != nil {
			log.Debugf("Bad png  %s", err)
			return err
		}

		f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {

			log.Debugf("Cannot open file  %s", err)
			return err
		}
		png.Encode(f, im)
		f.Close()
	} else if filepath.Ext(path) == ".jpg" || filepath.Ext(path) == ".jpeg" {
		im, err := jpeg.Decode(r)
		if err != nil {
			log.Debugf("Bad jpg  %s", err)
			return err
		}

		f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			log.Debugf("Cannot open file  %s", err)
			return err
		}
		jpeg.Encode(f, im, nil)
		f.Close()
	} else if filepath.Ext(path) == ".gif" {
		im, err := gif.Decode(r)
		if err != nil {
			log.Debugf("Bad gif  %s", err)
			return err
		}

		f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {

			log.Debugf("Cannot open file  %s", err)
			return err
		}
		gif.Encode(f, im, nil)
		f.Close()
	}
	return nil
}

const (
	encodePath encoding = 1 + iota
	encodeHost
	encodeUserPassword
	encodeQueryComponent
	encodeFragment
)

type encoding int
type EscapeError string

func (e EscapeError) Error() string {
	return "invalid URL escape " + strconv.Quote(string(e))
}

func ishex(c byte) bool {
	switch {
	case '0' <= c && c <= '9':
		return true
	case 'a' <= c && c <= 'f':
		return true
	case 'A' <= c && c <= 'F':
		return true
	}
	return false
}

func unhex(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}
	return 0
}

// Return true if the specified character should be escaped when
// appearing in a URL string, according to RFC 3986.
//
// Please be informed that for now shouldEscape does not check all
// reserved characters correctly. See golang.org/issue/5684.
func shouldEscape(c byte, mode encoding) bool {
	// §2.3 Unreserved characters (alphanum)
	if 'A' <= c && c <= 'Z' || 'a' <= c && c <= 'z' || '0' <= c && c <= '9' {
		return false
	}

	if mode == encodeHost {
		// §3.2.2 Host allows
		//  sub-delims = "!" / "$" / "&" / "'" / "(" / ")" / "*" / "+" / "," / ";" / "="
		// as part of reg-name.
		// We add : because we include :port as part of host.
		// We add [ ] because we include [ipv6]:port as part of host
		switch c {
		case '!', '$', '&', '\'', '(', ')', '*', '+', ',', ';', '=', ':', '[', ']':
			return false
		}
	}

	switch c {
	case '-', '_', '.', '~': // §2.3 Unreserved characters (mark)
		return false

	case '$', '&', '+', ',', '/', ':', ';', '=', '?', '@': // §2.2 Reserved characters (reserved)
		// Different sections of the URL allow a few of
		// the reserved characters to appear unescaped.
		switch mode {
		case encodePath: // §3.3
			// The RFC allows : @ & = + $ but saves / ; , for assigning
			// meaning to individual path segments. This package
			// only manipulates the path as a whole, so we allow those
			// last two as well. That leaves only ? to escape.
			return c == '?'

		case encodeUserPassword: // §3.2.1
			// The RFC allows ';', ':', '&', '=', '+', '$', and ',' in
			// userinfo, so we must escape only '@', '/', and '?'.
			// The parsing of userinfo treats ':' as special so we must escape
			// that too.
			return c == '@' || c == '/' || c == '?' || c == ':'

		case encodeQueryComponent: // §3.4
			// The RFC reserves (so we must escape) everything.
			return true

		case encodeFragment: // §4.1
			// The RFC text is silent but the grammar allows
			// everything, so escape nothing.
			return false
		}
	}

	// Everything else must be escaped.
	return true
}

func escape(s string, mode encoding) string {
	spaceCount, hexCount := 0, 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if shouldEscape(c, mode) {
			if c == ' ' && mode == encodeQueryComponent {
				spaceCount++
			} else {
				hexCount++
			}
		}
	}

	if spaceCount == 0 && hexCount == 0 {
		return s
	}

	t := make([]byte, len(s)+2*hexCount)
	j := 0
	for i := 0; i < len(s); i++ {
		switch c := s[i]; {
		case c == ' ' && mode == encodeQueryComponent:
			t[j] = '+'
			j++
		case shouldEscape(c, mode):
			t[j] = '%'
			t[j+1] = "0123456789ABCDEF"[c>>4]
			t[j+2] = "0123456789ABCDEF"[c&15]
			j += 3
		default:
			t[j] = s[i]
			j++
		}
	}
	return string(t)
}

// unescape unescapes a string; the mode specifies
// which section of the URL string is being unescaped.
func unescape(s string, mode encoding) (string, error) {
	// Count %, check that they're well-formed.
	n := 0
	hasPlus := false
	for i := 0; i < len(s); {
		switch s[i] {
		case '%':
			n++
			if i+2 >= len(s) || !ishex(s[i+1]) || !ishex(s[i+2]) {
				s = s[i:]
				if len(s) > 3 {
					s = s[:3]
				}
				return "", EscapeError(s)
			}
			i += 3
		case '+':
			hasPlus = mode == encodeQueryComponent
			i++
		default:
			i++
		}
	}

	if n == 0 && !hasPlus {
		return s, nil
	}

	t := make([]byte, len(s)-2*n)
	j := 0
	for i := 0; i < len(s); {
		switch s[i] {
		case '%':
			t[j] = unhex(s[i+1])<<4 | unhex(s[i+2])
			j++
			i += 3
		case '+':
			if mode == encodeQueryComponent {
				t[j] = ' '
			} else {
				t[j] = '+'
			}
			j++
			i++
		default:
			t[j] = s[i]
			j++
			i++
		}
	}
	return string(t), nil
}

func EncodeUriComponent(rawString string) string {
	return escape(rawString, encodeFragment)
}

func DecodeUriCompontent(encoded string) (string, error) {
	return unescape(encoded, encodeQueryComponent)
}
