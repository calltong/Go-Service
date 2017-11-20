package notify

import (
  //"fmt"
	"encoding/json"
	"net/http"
  //"time"
  //"bytes"
  "errors"
)

func SendToUser(company, id string, data []byte) error {
  url := "/user/notify/" + company + "/" + id

	return send(url, data)
}

func SendToAdmin(company, id string, data []byte) error {
  url := "/admin/notify/" + company + "/" + id

	return send(url, data)
}

func send(url, data []byte) error {
  res, err := http.Post(url, "application/json", data)
  defer res.Body.Close()
  if err != nil {
    panic(err)
    return err
  }

  if res.StatusCode != http.StatusOK {
    decoder := json.NewDecoder(res.Body)
    var resMsg ResMessage
  	err = decoder.Decode(&resMsg)
    return errors.New(resMsg.Message)
  }

	return nil
}
