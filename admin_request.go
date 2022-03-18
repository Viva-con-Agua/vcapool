package vcapool

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/Viva-con-Agua/vcago"
	"github.com/sonh/qs"
)

type AdminRequest struct {
	URL string
}

func NewAdminRequest() *AdminRequest {
	return &AdminRequest{
		URL: vcago.Config.GetEnvString("ADMIN_URL", "n", "http://172.4.5.3"),
	}
}

func (i *AdminRequest) GetUser(query *UserQuery) (r *UserList, err error) {
	encoder := qs.NewEncoder()
	var values url.Values
	if values, err = encoder.Values(query); err != nil {
		return
	}
	url := i.URL + "/admin/users?" + values.Encode()
	log.Print(url)
	request := new(http.Request)
	request, err = http.NewRequest("GET", url, nil)
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	client := &http.Client{}
	response := new(http.Response)
	response, err = client.Do(request)
	if err != nil {
		return nil, vcago.NewStatusInternal(err)
	}
	defer response.Body.Close()
	var bodyBytes []byte
	if response.StatusCode != 200 {
		if bodyBytes, err = ioutil.ReadAll(response.Body); err != nil {
			return nil, vcago.NewStatusInternal(err)
		}
		body := new(interface{})
		if err = json.Unmarshal(bodyBytes, body); err != nil {
			return nil, vcago.NewStatusInternal(err)
		}
		return nil, vcago.NewStatusInternal(err)
	}
	body := new(vcago.Response)
	if bodyBytes, err = ioutil.ReadAll(response.Body); err != nil {
		return nil, vcago.NewStatusInternal(err)
	}
	if err = json.Unmarshal(bodyBytes, body); err != nil {
		log.Print(err)
		return nil, vcago.NewStatusInternal(err)
	}
	if body.Payload != nil {
		bytes, _ := json.Marshal(&body.Payload)
		r = new(UserList)
		_ = json.Unmarshal(bytes, &r)
	}
	return
}
