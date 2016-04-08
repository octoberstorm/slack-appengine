package slack

import (
	"encoding/json"
	"errors"
)

//
// types
//

// slack user type
type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Deleted  bool   `json:"deleted"`
	Color    string `json:"color"`
	Profile  *ProfileInfo
	IsAdmin  bool `json:"is_admin"`
	IsOwner  bool `json:"is_owner"`
	Has2fa   bool `json:"has_2fa"`
	HasFiles bool `json:"has_files"`
}

// slack user profile type
type ProfileInfo struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	RealName  string `json:"real_name"`
	Email     string `json:"email"`
	Skype     string `json:"skype"`
	Phone     string `json:"phone"`
	Image25   string `json:"image_24"`
	Image32   string `json:"image_32"`
	Image48   string `json:"image_48"`
	Image72   string `json:"image_72"`
	Image192  string `json:"image_192"`
}

type UserData []*User

// implement the sort interface
func (u UserData) Len() int {
	return len(u)
}
func (u UserData) Less(i, j int) bool {
	return u[i].Profile.FirstName < u[j].Profile.FirstName
}
func (u UserData) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

// presence data
type UserPresence struct {
	Ok              string `json:"ok,omitempty"`
	Presence        string `json:"presence,omitempty"`
	Online          bool   `json:"online,omitempty"`
	AutoAway        bool   `json:"auto_away,omitempty"`
	ManualAway      bool   `json:"manual_away,omitempty"`
	ConnectionCount int    `json:"connection_count,omitempty"`
	LastActivity    int64  `json:"last_activity,omitempty"`
}

// response type of `users.info` api
type UsersInfoAPIResponse struct {
	BaseAPIResponse
	User *User `json:"user"`
}

// response type of `users.list` api
type UsersListAPIResponse struct {
	BaseAPIResponse
	RawMembers json.RawMessage `json:"members"`
}

// user presence response type of `users.getpresence` api
type UserPresenceAPIResponse struct {
	BaseAPIResponse
	Presence *UserPresence
}

//
// funcs
//

// API users.list: Lists all users in a Slack team.
func (sl *Slack) UsersList() (UserData, error) {
	uv := sl.urlValues()
	body, err := sl.GetRequest(usersListApiEndpoint, uv)
	if err != nil {
		return nil, err
	}
	res := new(UsersListAPIResponse)
	err = json.Unmarshal(body, res)
	if err != nil {
		return nil, err
	}
	if !res.Ok {
		return nil, errors.New(res.Error)
	}
	return res.Members()
}

func (res *UsersListAPIResponse) Members() (UserData, error) {
	var members UserData
	err := json.Unmarshal(res.RawMembers, &members)
	if err != nil {
		return nil, err
	}
	return members, nil
}

// FindUser returns a user object that satisfy conditions specified.
func (sl *Slack) FindUser(cb func(*User) bool) (*User, error) {
	members, err := sl.UsersList()
	if err != nil {
		return nil, err
	}
	for _, member := range members {
		if cb(member) {
			return member, nil
		}
	}
	return nil, errors.New("No such user.")
}

// FindUserByName returns a user object that matches name specified.
func (sl *Slack) FindUserByName(name string) (*User, error) {
	return sl.FindUser(func(user *User) bool {
		return user.Name == name
	})
}

// get user's presence
func (sl *Slack) GetUserPresence(name string) (*UserPresence, error) {
	uv := sl.urlValues()
	uv.Add("user", name)
	body, err := sl.GetRequest(userPresenceApiEndpoint, uv)
	if err != nil {
		return nil, err
	}
	res := new(UserPresenceAPIResponse)
	err = json.Unmarshal(body, res)
	if err != nil {
		return nil, err
	}
	if !res.Ok {
		return nil, errors.New(res.Error)
	}
	return res.Presence, nil
}

// API users.info: Gets information about a user.
func (sl *Slack) UsersInfo(userId string) (*User, error) {
	uv := sl.urlValues()
	uv.Add("user", userId)
	body, err := sl.GetRequest(usersInfoApiEndpoint, uv)
	if err != nil {
		return nil, err
	}
	res := new(UsersInfoAPIResponse)
	err = json.Unmarshal(body, res)
	if err != nil {
		return nil, err
	}
	if !res.Ok {
		return nil, errors.New(res.Error)
	}
	return res.User, nil
}
