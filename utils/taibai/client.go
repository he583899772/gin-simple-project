package taibai

import (
	"aiops/network-detect/utils"
	"errors"
	"net/url"

	"github.com/spf13/viper"

	"github.com/go-resty/resty/v2"
)

type userResp struct {
	Data   UserInfo `json:"data"`
	Status int      `json:"status"`
	ErrMsg string   `json:"errmsg"`
}

type UserInfo struct {
	Uid      int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Tel      string `json:"tel"`
	Mail     string `json:"mail"`
	Roles    []Role `json:"roles"`
}

type Role struct {
	Name string `json:"name"`
}

func GetUserInfo(username string) (*UserInfo, error) {
	var res userResp
	req := resty.New().R().SetResult(&res)
	auth, err := utils.GenerateAuth()
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", auth)
	payload := map[string]string{
		"operation": "query",
		"username":  username,
	}
	req.SetBody(payload)
	if res, err := req.Post(viper.GetString("dependSys.taiBaiURL") + "/v2/user-info"); err != nil {
		return nil, err
	} else if res.IsSuccess() {
		return &res.Result().(*userResp).Data, nil
	} else {
		return nil, errors.New(res.String())
	}
}

type teamResp struct {
	Data   []UserInfo `json:"data"`
	Status int        `json:"status"`
	ErrMsg string     `json:"errmsg"`
}

func GetTeamMembers(roleId string) ([]UserInfo, error) {
	var res teamResp
	req := resty.New().R().SetResult(&res)
	auth, err := utils.GenerateAuth()
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", auth)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	payload := url.Values{}
	payload.Set("role_id", roleId)
	req.SetFormDataFromValues(payload)
	if res, err := req.Post(viper.GetString("dependSys.taiBaiURL") + "/v1/gw/roleusers"); err != nil {
		return nil, err
	} else if res.IsSuccess() {
		return res.Result().(*teamResp).Data, nil
	} else {
		return nil, errors.New(res.String())
	}
}

func IsExistTeamMembers(mail, teamId string) (bool, error) {
	members, err := GetTeamMembers(teamId)
	if err != nil {
		return false, err
	}
	for _, member := range members {
		if member.Mail == mail {
			return true, nil
		}
	}
	return false, nil
}
