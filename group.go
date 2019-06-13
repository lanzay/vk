package vk

import (
	"encoding/json"
	"strconv"
	"strings"
)

const (
	ENDPOINT = "https://api.vk.com/method"
	//https://api.vk.com/method/users.get?user_ids=210700286&fields=bdate&access_token=533bacf01e11f55b536a565b57531ac114461ae8736d6506a3&v=5.95
)

type VkGroup struct {
	vk *VK
	ID string
}

func (g *VkGroup) GetIdByUrl(u string) string {

	//https://vk.com/umnie_knizki?w=wall-44155346_79580
	//https://vk.com/club27669892?w=wall-27669892_6721
	//https://vk.com/public44155346?w=wall-44155346_66766

	tmp1 := strings.Split(u, "?")[0]
	tmp2 := strings.Split(tmp1, "/")
	tmp3 := strings.TrimPrefix(tmp2[len(tmp2)-1], "club")
	tmp4 := strings.TrimPrefix(tmp3, "public")
	return tmp4
}

func (g *VkGroup) GetInfo(groupIDs string) ([]Group, error) {
	return g.getById(groupIDs)
}

//groups.getById
//Возвращает информацию о заданном сообществе или о нескольких сообществах.
//https://vk.com/dev/groups.getById
func (g *VkGroup) getById(groupIDs string) ([]Group, error) {

	data, err := g.getByIdBody(groupIDs)
	if err != nil {
		return nil, err
	}

	var r []Group
	err = json.Unmarshal(data, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}
func (g *VkGroup) getByIdBody(groupIDs string) ([]byte, error) {

	fields := "members_count,place,site,status,verified,activity,age_limits,can_create_topic,can_message,can_post,can_upload_doc,wall,contacts,cover,description,links,main_album_id,market,public_date_label,trending,country,city,main_section,wiki_page,ban_info,can_see_all_posts,can_see_all_posts,public_date_label,counters"

	p := make(map[string]string)
	p["method"] = "groups.getById"
	p["v"] = "5.95"

	p["group_ids"] = groupIDs
	p["fields"] = fields

	return g.vk.GetBody(p, nil)
}

//groups.getMembers
//https://vk.com/dev/groups.getMembers
func (g *VkGroup) GetMembers(id int) ([]int, error) {

	var count int
	var members []int

	for {
		data, err := g.getMembers(id, count)
		if err != nil {
			return nil, err
		}
		var m ResponseInt
		err = json.Unmarshal(data, &m)
		members = append(members, m.Items...)
		count += len(m.Items)
		if count >= m.Count && len(m.Items) == 0 {
			break
		}
	}
	return members, nil
}

func (g *VkGroup) getMembers(id int, offset int) ([]byte, error) {

	p := make(map[string]string)
	p["method"] = "groups.getMembers"
	p["v"] = "5.95"

	p["group_id"] = strconv.Itoa(id)
	if offset != 0 {
		p["offset"] = strconv.Itoa(offset)
	}
	return g.vk.GetBody(p, nil)
}

//===============================================
type Group struct {
	ID             int           `json:"id"`
	Name           string        `json:"name"`
	ScreenName     string        `json:"screen_name"`
	IsClosed       int           `json:"is_closed"`
	Type           string        `json:"type"`
	IsAdmin        int           `json:"is_admin"`
	IsMember       int           `json:"is_member"`
	IsAdvertiser   int           `json:"is_advertiser"`
	MembersCount   int           `json:"members_count"`
	Site           string        `json:"site"`
	Status         string        `json:"status"`
	Verified       int           `json:"verified"`
	Activity       string        `json:"activity"`
	AgeLimits      int           `json:"age_limits"`
	CanCreateTopic int           `json:"can_create_topic"`
	CanPost        int           `json:"can_post"`
	CanUploadDoc   int           `json:"can_upload_doc"`
	Wall           int           `json:"wall"`
	Contacts       []interface{} `json:"contacts"`
	Description    string        `json:"description"`
	Links          []LinkGroup   `json:"links"`
	Market         Market        `json:"market"`
	Trending       int           `json:"trending"`
	Country        City          `json:"country"`
	City           City          `json:"city"`
	MainSection    int           `json:"main_section"`
	WikiPage       string        `json:"wiki_page"`
	CanSeeAllPosts int           `json:"can_see_all_posts"`
	Counters       Counters      `json:"counters"`
	CanMessage     int           `json:"can_message"`
	Cover          Cover         `json:"cover"`
	Photo50        string        `json:"photo_50"`
	Photo100       string        `json:"photo_100"`
	Photo200       string        `json:"photo_200"`
}

type Counters struct {
	Topics    int `json:"topics"`
	Videos    int `json:"videos"`
	Addresses int `json:"addresses"`
}

type Cover struct {
	Enabled int     `json:"enabled"`
	Images  []Image `json:"images"`
}

type LinkGroup struct {
	ID        int    `json:"id"`
	URL       string `json:"url"`
	Name      string `json:"name"`
	Desc      string `json:"desc"`
	Photo50   string `json:"photo_50"`
	Photo100  string `json:"photo_100"`
	EditTitle *int   `json:"edit_title,omitempty"`
}

type Market struct {
	Enabled int `json:"enabled"`
}
