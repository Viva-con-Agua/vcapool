package vcapool

/*
type AdminRequest struct {
	URL string
}

func NewAdminRequest() *AdminRequest {
	return &AdminRequest{
		URL: vcago.Config.GetEnvString("ADMIN_URL", "n", "http://172.4.5.3"),
	}
}

func (i *AdminRequest) Get(path string) (r *vcago.Response, err error) {
	url := i.URL + path
	request := new(http.Request)
	request, err = http.NewRequest("GET", url, nil)
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	client := &http.Client{}
	response := new(http.Response)
	response, err = client.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	var bodyBytes []byte
	if response.StatusCode != 200 {
		if bodyBytes, err = ioutil.ReadAll(response.Body); err != nil {
			return
		}
		body := new(interface{})
		if err = json.Unmarshal(bodyBytes, body); err != nil {
			return
		}
		return nil, errors.New(response.Status)
	}
	r = new(vcago.Response)
	if bodyBytes, err = ioutil.ReadAll(response.Body); err != nil {
		return
	}
	if err = json.Unmarshal(bodyBytes, r); err != nil {
		return
	}
	return

}

func (i *AdminRequest) Post(path string, data interface{}) (r *vcago.Response, err error) {
	var jsonData []byte
	if jsonData, err = json.Marshal(data); err != nil {
		return
	}
	url := i.URL + path
	request := new(http.Request)
	request, err = http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	client := &http.Client{}
	response := new(http.Response)
	response, err = client.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	var bodyBytes []byte
	if response.StatusCode != 201 {
		if bodyBytes, err = ioutil.ReadAll(response.Body); err != nil {
			return
		}
		body := new(interface{})
		if err = json.Unmarshal(bodyBytes, body); err != nil {
			return
		}
		return nil, errors.New(response.Status)
	}
	r = new(vcago.Response)
	if bodyBytes, err = ioutil.ReadAll(response.Body); err != nil {
		return
	}
	if err = json.Unmarshal(bodyBytes, r); err != nil {
		return
	}
	return

}

func (i *AdminRequest) GetUser(query *UserQuery) (r *UserList, err error) {
	uRL := "/admin/users"
	if query != nil {
		encoder := qs.NewEncoder()
		var values url.Values
		if values, err = encoder.Values(query); err != nil {
			return
		}
		uRL = uRL + "?" + values.Encode()
	}
	response := new(vcago.Response)
	if response, err = i.Get(uRL); err != nil {
		return
	}
	if response.Payload != nil {
		bytes, _ := json.Marshal(&response.Payload)
		r = new(UserList)
		_ = json.Unmarshal(bytes, &r)
	}
	return
}

func (i *AdminRequest) GetCrew(query *CrewQuery) (r *CrewList, err error) {
	uRL := "/admin/crews"
	if query != nil {
		encoder := qs.NewEncoder()
		var values url.Values
		if values, err = encoder.Values(query); err != nil {
			return
		}
		uRL = uRL + "?" + values.Encode()
	}
	response := new(vcago.Response)
	if response, err = i.Get(uRL); err != nil {
		return
	}
	if response.Payload != nil {
		bytes, _ := json.Marshal(&response.Payload)
		r = new(CrewList)
		_ = json.Unmarshal(bytes, &r)
	}
	return
}*/
