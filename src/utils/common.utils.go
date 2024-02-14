package utils

import (
	"encoding/base64"
	"math/rand"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func GetValueStr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func GetUUID() string {
	return uuid.New().String()
}

func CheckBirthDate(bd string) (bool, string) {

	format := "02/01/2006" // dd/mm/yyyy
	start, err := time.Parse(format, bd)
	if err != nil {
		return false, "birthdate incorrect format"
	}
	end := time.Now()

	if start.After(end) {
		return false, "birthdate incorrect format"
	}
	return true, ""
}

func ConvertIntTo2Decimal(value float64) float64 {
	return float64(int(value*100)) / 100
}

func CheckDecimal(v float64) int {
	s := strconv.FormatFloat(v, 'f', 2, 64)
	i := strings.IndexByte(s, '.')
	if i > -1 {
		return len(s) - i - 1
	}
	return 0
}

func SetDateFormatByDateStr(str_date string, new_format string, layout string) (string, time.Time) {
	str := str_date
	if layout == "" {
		layout = "2006-01-02 15:04:05"
	}
	if new_format == "" {
		new_format = "2006-01-02T15:04:05Z"
	}
	t, err := time.Parse(layout, str)

	if err != nil {
		return "", t
	}

	return t.Format(new_format), t
}

func RangeIn(low, hi int) int {
	return low + rand.Intn(hi-low)
}

func Int(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}

func String(i int64) string {
	s := strconv.FormatInt(i, 10)
	return s
}

func UrlDecoded(str string) string {
	u, err := url.QueryUnescape(str)
	if err != nil {
		return ""
	}
	return u
}

func Bool(str string) bool {
	b, _ := strconv.ParseBool(str)
	return b
}

func ConvDatetimeFormatLog(t time.Time) string {
	return t.Format("2006-01-02T15:04:05-0700")
}

func IfThenElse(condition bool, a interface{}, b interface{}) string {
	if condition {
		return a.(string)
	}
	return b.(string)
}

func Linearsearch(datalist []string, key string) bool {
	for _, item := range datalist {
		if item == key {
			return true
		}
	}
	return false
}

func ValidateBase64(b64 string) bool {
	_, err := base64.StdEncoding.DecodeString(b64)
	return err == nil
}

func DefaultValue(value string) string {
	if value == "" {
		value = "-"
	}
	return value
}

func ValidatorNew(s interface{}, fielderror *string) error {
	v := validator.New()
	err := v.Struct(s)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			if e != nil {
				descError := ""
				if e.Param() != "" {
					descError = "(" + e.Tag() + ":" + e.Param() + ")"
				}
				*fielderror += e.Field() + descError + ","
			}
		}
	}
	return err
}

func GetDateOffer(trContractTerm string) (string, string) {
	startDate := ""
	expireDate := ""

	now := time.Now()
	startDate = now.Format("2006-01-02 00:00:00")
	months := 0
	if trContractTerm != "" && trContractTerm != "0" {
		months, _ = strconv.Atoi(trContractTerm)
	}
	expireDate = now.AddDate(0, months, 0).Format("2006-01-02 00:00:00")

	return startDate, expireDate
}

func GetErrorResponse(err error) *[]ErrorParam {
	if err != nil {
		out := make([]ErrorParam, len(err.(validator.ValidationErrors)))
		for i, e := range err.(validator.ValidationErrors) {
			if e.Param() != "" {
				out[i] = ErrorParam{Param: e.Namespace(), Message: e.Tag() + " -> " + e.Param()}
			} else {
				out[i] = ErrorParam{Param: e.Namespace(), Message: e.Tag()}
			}
		}
		return &out
	}
	return nil
}

func ResponseSuccess(c *fiber.Ctx, message string, result interface{}) error {
	var response ResponseStandard

	if message == "" {
		message = "success"
	}

	response.Code = strconv.Itoa(fiber.StatusOK)
	response.Message = message
	response.Result = result
	response.CorrelationId = c.GetRespHeader("X-Request-Id")

	return c.Status(fiber.StatusOK).JSON(response)
}

func ResponseFailed(c *fiber.Ctx, code int, message string, err error) error {
	var response ResponseStandard
	if code == 0 {
		code = fiber.StatusOK
	}

	if message == "" {
		message = "failed"
	}

	response.Code = strconv.Itoa(code)
	response.Message = message
	response.Result = nil
	response.CorrelationId = c.GetRespHeader("X-Request-Id")

	if message == "Required Field" {
		response.Error = GetErrorResponse(err)
	} else if err != nil {
		response.ErrorMessage = err.Error()
	}

	return c.Status(code).JSON(response)
}
