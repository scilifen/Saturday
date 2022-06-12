package router_test

import (
	"log"
	"saturday/util"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetEventById(t *testing.T) {
	rawAPITestCase, err := util.GetCsvMap("testdata/get_event_by_id.csv")
	if err != nil {
		t.Error(err)
	}
	for _, rawCase := range rawAPITestCase {
		t.Run(rawCase["CaseId"], func(t *testing.T) {
			code, _ := strconv.Atoi(rawCase["code"])
			// auth := rawCase["Authorization"]
			testCase := APITestCase{
				"success",
				Request{
					"GET",
					"/events/" + rawCase["event_id"],
					// GenToken(auth, "2333333333"),
					"",
					gin.H{},
				},
				Response{
					code,
					gin.H{
						"event_id":  1,
						"client_id": 1,
						"model":     "7590",
						"problem":   "hackintosh",
						"member": gin.H{
							"member_id":    "2333333333",
							"alias":        "滑稽",
							"name":         "滑稽",
							"section":      "计算机233",
							"role":         "member",
							"profile":      "relaxing",
							"phone":        "12356839487",
							"qq":           "123456",
							"avatar":       "",
							"created_by":   "0000000000",
							"gmt_create":   "2022-04-23 15:49:59",
							"gmt_modified": "2022-04-30 17:29:46",
						},
						"closed_by": gin.H{
							"member_id":    "0000000000",
							"alias":        "管理",
							"name":         "管理",
							"section":      "计算机000",
							"role":         "admin",
							"profile":      "",
							"phone":        "",
							"qq":           "",
							"avatar":       "",
							"created_by":   "",
							"gmt_create":   "2022-04-30 17:28:42",
							"gmt_modified": "2022-04-30 17:28:44",
						},
						"status":       "accepted",
						"logs":         "IGNORE",
						"gmt_create":   "2022-05-10 10:23:54",
						"gmt_modified": "2022-06-02 16:18:52",
					},
				},
			}
			if rawCase["success"] != "TRUE" {
				testCase.Response.Body = gin.H{
					"message": rawCase["error_message"],
				}
			}
			err = testCase.Test()
			if err != nil {
				t.Error(err)
			}
		})
	}
}

func TestCreateEvent(t *testing.T) {
	rawAPITestCase, err := util.GetCsvMap("testdata/create_event.csv")
	if err != nil {
		t.Error(err)
	}
	for _, rawCase := range rawAPITestCase {
		log.Println(rawCase["CaseId"])
		log.Println(rawCase["code"])
		t.Run(rawCase["CaseId"], func(t *testing.T) {
			code, _ := strconv.Atoi(rawCase["code"])
			auth := rawCase["Authorization"]
			testCase := APITestCase{
				"success",
				Request{
					"POST",
					"/client/events",
					GenToken(auth, "1"),
					gin.H{
						"model":              rawCase["model"],
						"phone":              rawCase["phone"],
						"qq":                 rawCase["qq"],
						"contact_preference": rawCase["contact_preference"],
						"problem":            rawCase["problem"],
					},
				},
				Response{
					code,
					gin.H{
						"event_id":           "IGNORE",
						"client_id":          1,
						"model":              rawCase["model"],
						"phone":              rawCase["phone"],
						"qq":                 rawCase["qq"],
						"contact_preference": rawCase["contact_preference"],
						"problem":            rawCase["problem"],
						"member_id":          "",
						"closed_by":          "",
						"status":             "open",
						"logs":               "IGNORE",
						"gmt_create":         "2022-05-10 10:23:54",
						"gmt_modified":       "2022-05-12 23:22:44",
					},
				},
			}
			if rawCase["success"] != "TRUE" {
				testCase.Response.Body = gin.H{
					"message": "Validation Failed",
				}
			}
			err := testCase.Test()
			if err != nil {
				t.Error(err)
			}

		})
	}
}
