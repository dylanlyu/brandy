package util

import (
	"math"
	"math/rand"
	"net"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func PointerString(s string) *string     { return &s }
func PointerInt(i int) *int              { return &i }
func PointerInt64(i int64) *int64        { return &i }
func PointerBool(b bool) *bool           { return &b }
func PointerTime(t time.Time) *time.Time { return &t }

func GenerateRangeNum(min, max int) int {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(max-min) + min
	return randNum
}

func Round(x float64) int64 {
	return int64(math.Floor(x + 0.5))
}

func FirstUpper(s string) string {
	if s == "" {
		return ""
	}

	match, _ := regexp.MatchString("^[0-9a-zA-Z]+$", s)
	if match == false {
		return s
	}

	return strings.ToUpper(s[:1]) + strings.ToLower(s[1:])
}

func FirstLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}

func Int64Abs(n int64) int64 {
	y := n >> 63
	return (n ^ y) - y
}

func GetConstellation(month, day int) (star string) {
	if month <= 0 || month >= 13 {
		star = "-1"
	}

	if day <= 0 || day >= 32 {
		star = "-1"
	}

	if (month == 1 && day >= 20) || (month == 2 && day <= 18) {
		star = "Aquarius"
	}

	if (month == 2 && day >= 19) || (month == 3 && day <= 20) {
		star = "Pisces"
	}

	if (month == 3 && day >= 21) || (month == 4 && day <= 19) {
		star = "Aries"
	}

	if (month == 4 && day >= 20) || (month == 5 && day <= 20) {
		star = "Taurus"
	}

	if (month == 5 && day >= 21) || (month == 6 && day <= 21) {
		star = "Gemini"
	}

	if (month == 6 && day >= 22) || (month == 7 && day <= 22) {
		star = "Cancer"
	}

	if (month == 7 && day >= 23) || (month == 8 && day <= 22) {
		star = "Leo"
	}

	if (month == 8 && day >= 23) || (month == 9 && day <= 22) {
		star = "Virgo"
	}

	if (month == 9 && day >= 23) || (month == 10 && day <= 22) {
		star = "Libra"
	}

	if (month == 10 && day >= 23) || (month == 11 && day <= 21) {
		star = "Scorpio"
	}

	if (month == 11 && day >= 22) || (month == 12 && day <= 21) {
		star = "Sagittarius"
	}

	if (month == 12 && day >= 22) || (month == 1 && day <= 19) {
		star = "Capricorn"
	}

	return star
}

func RandShuffle(slice []any) {
	rand.NewSource(time.Now().UnixNano())
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

func IsChinese(str string) bool {
	var count int
	for _, v := range str {
		if unicode.Is(unicode.Han, v) {
			count++
			break
		}
	}
	return count > 0
}

func GetVerifyCode() string {
	seed := rand.NewSource(time.Now().UnixNano())
	randomNumber := rand.New(seed)

	verifyCode := strconv.Itoa(randomNumber.Intn(999999))
	verifyCode += "000000"

	return verifyCode[:6]
}

func AddSpace(s string) string {
	n := len(s)
	if n <= 0 {
		return s
	}

	return AddSpace(s[:n-1]) + " " + s[n-1:]
}

func IsZeroLimitAndOffset(limit, page int64) (int64, int64) {
	if limit <= 0 {
		limit = -1
	}

	if page <= 0 {
		page = -1
	}

	if page > 0 && limit > 0 {
		offset := (page - 1) * limit
		return limit, offset
	}

	return limit, page
}

func GetRemoteIP(req *http.Request) string {
	ip := req.Header.Get("XRealIP")
	if ip == "" {
		ip = req.Header.Get("XForwardedFor")
	}

	if ip == "" {
		if addr, _, err := net.SplitHostPort(req.RemoteAddr); err == nil {
			ip = addr
		} else {
			ip = req.RemoteAddr
		}
	}

	if idx := strings.IndexByte(ip, ','); idx >= 0 {
		ip = ip[:idx]
	}

	if ip == "::1" {
		ip = "127.0.0.1"
	}

	return ip
}
