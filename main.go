package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type RawResponse struct {
	Response struct {
		Count int `json:"count"`
		Items []struct {
			Attachments []struct {
				Photo struct {
					AccessKey string `json:"access_key"`
					AlbumID   int    `json:"album_id"`
					Date      int    `json:"date"`
					Height    int    `json:"height"`
					ID        int    `json:"id"`
					OwnerID   int    `json:"owner_id"`
					Photo130  string `json:"photo_130"`
					Photo604  string `json:"photo_604"`
					Photo75   string `json:"photo_75"`
					PostID    int    `json:"post_id"`
					Text      string `json:"text"`
					UserID    int    `json:"user_id"`
					Width     int    `json:"width"`
				} `json:"photo"`
				Type string `json:"type"`
			} `json:"attachments"`
			Comments struct {
				Count int `json:"count"`
			} `json:"comments"`
			Date     int `json:"date"`
			FromID   int `json:"from_id"`
			ID       int `json:"id"`
			IsPinned int `json:"is_pinned"`
			Likes    struct {
				Count int `json:"count"`
			} `json:"likes"`
			OwnerID  int    `json:"owner_id"`
			PostType string `json:"post_type"`
			Reposts  struct {
				Count int `json:"count"`
			} `json:"reposts"`
			Text string `json:"text"`
		} `json:"items"`
	} `json:"response"`
}

func apiCall(method string, params map[string]string) {
	queryString := url.Values{}
	for key, value := range params {
		queryString.Add(key, value)
	}
	queryString.Add("v", "5.37")
	res, err := http.Get(fmt.Sprintf("http://api.vk.com/method/%s?%s", method, queryString.Encode()))
	if err != nil {
		log.Fatal("Can't get response", err.Error())
		return
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Can't get body", err.Error())
		return
	}
	var rawResponse RawResponse
	err = json.Unmarshal(data, &rawResponse)
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, item := range rawResponse.Response.Items {
		fmt.Println(item.Text)
	}
	//fmt.Printf("%s\n", data)
}

func main() {
	fmt.Println("Starting...")
	params := map[string]string{"owner_id": "-44989697"}
	apiCall("wall.get", params)

}
