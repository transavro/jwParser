package jwParser

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"net/http"
)

var(
	jwMonetizeType []string
	jwProviderArray []string
	jwProvidersMap = make(map[string]int64)
	jwGenre = make(map[int64]string)
	jwProviderSourceMap = make(map[string]string)
	jwCategories []string
)

func init() {
	log.Println("JW ASSETS INIT ")
	// monetize_type
	jwMonetizeType = []string{"free", "flatrate", "ads", "rent'", "buy", "5D"}
	//categories
	jwCategories = []string{"movie", "show"}
}

func getGenreFromJW()  {
	resp , err := http.Get("https://apis.justwatch.com/content/genres/locale/en_IN")
	if err != nil {
		log.Fatal("Error while fetch JW Genre",err)
	}
	if resp.StatusCode != 200 {
		log.Fatal("Error while fetching JW Genre ",resp.StatusCode)
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		log.Fatal("Io error ", err)
	}
	for _, r := range gjson.ParseBytes(bytes).Array() {
		jwGenre[r.Get("id").Int()] = r.Get("translation").String()
	}
}

func getProviderFromJW()  {
	resp , err := http.Get("https://apis.justwatch.com/content/providers/locale/en_IN")
	if err != nil {
		log.Fatal("Error while fetch JW Provider",err)
	}
	if resp.StatusCode != 200 {
		log.Fatal("Error while fetching JW Provider ",resp.StatusCode)
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		log.Fatal("Io error ", err)
	}
	for _, r := range gjson.ParseBytes(bytes).Array() {
		jwProvidersMap[r.Get("short_name").String()] = r.Get("id").Int()
		jwProviderSourceMap[r.Get("short_name").String()] = r.Get("clear_name").String()
	}
}

func GetPackageNameForJW(sourceCode interface{}) string {
	sources := fmt.Sprint(sourceCode)
	switch sources {
	case "nfx":
		{
			return "com.netflix.mediaclient"
		}
	case "hoo":
		{
			return "tv.hooq.android"
		}
	case "prv":
		{
			return "com.amazon.avod.thirdpartyclient"
		}
	case "hst":
		{
			return "in.startv.hotstar"
		}
	case "voo":
		{
			return "com.tv.v18.viola"
		}
	case "viu":
		{
			return "com.vuclip.viu"
		}
	case "jio":
		{
			return "com.jio.media.ondemand"
		}
	case "zee":
		{
			return "com.graymatrix.did"
		}
	case "ern":
		{
			return "com.erosnow "
		}
	case "ply":
		{
			return "com.android.vending"
		}
	case "mbi":
		{
			return "com.mubi"
		}
	case "snl":
		{
			return "com.sonyliv"
		}
	case "yot":
		{
			return "com.google.android.youtube.tv"
		}
	case "nfk":
		{
			return "com.netflix.mediaclient"
		}
	case "tbv":
		{
			return "com.tubitv"
		}
	case "ytv":
		{
			return "com.tru"
		}
	case "snx":
		{
			return "com.suntv.sunnxt"
		}
	case "cru":
		{
			return "Crunchyroll"
		}
	case "hoc":
		{
			return "com.viewlift.hoichoi"
		}
	case "abj":
		{
			return "com.balaji.alt"
		}
	}
	return ""
}

func isoCodeToLanguage(langCode string) string {
	switch langCode {
	case "en":
		{
			return "English"
		}
	case "hi":
		{
			return "Hindi"
		}
	case "ta":
		{
			return "Tamil"
		}
	case "ko":
		{
			return "Korean"
		}
	case "te":
		{
			return "Telugu"
		}
	case "ka":
		{
			return "Georgian"
		}
	case "kn":
		{
			return "Kannada"
		}
	case "mr":
		{
			return "Marathi"
		}
	case "bn":
		{
			return "Bangla"
		}
	case "ml":
		{
			return "Malayalam"
		}
	case "zh":
		{
			return "Chinese"
		}
	case "or":
		{
			return "Oriya"
		}
	case "pa":
		{
			return "Punjabi"
		}
	case "ja":
		{
			return "Japanese"
		}
	case "fr":
		{
			return "French"
		}
	}
	return ""
}

func itemExists(a []*ContentAvailable, x *ContentAvailable) bool {
	for _, avaliable := range a {
		if avaliable.Source == x.Source {
			return true
		}
	}
	return false
}
