package vk

import (
	"encoding/json"
	"strconv"
)

//https://vk.com/dev/wall
type VkWall struct {
	vk *VK
}

//wall.get
//Возвращает список записей со стены пользователя или сообщества.
//https://vk.com/dev/wall.get
func (w *VkWall) GetPosts(id int) ([]Post, error) {

	//TODO сейчас только первые 100 записей
	data, err := w.getPosts(id, 0)
	if err != nil {
		return nil, err
	}

	var r WallResponse
	err = json.Unmarshal(data, &r)

	return r.Items, nil

}

func (w *VkWall) getPosts(id int, offset int) ([]byte, error) {

	p := make(map[string]string)
	p["method"] = "wall.get"
	p["v"] = "5.95"

	p["owner_id"] = strconv.Itoa(id)
	p["count"] = "100"
	if offset != 0 {
		p["offset"] = strconv.Itoa(offset)
	}
	return w.vk.GetBody(p, nil)
}

//wall.getComments
//Возвращает список комментариев к записи на стене.
//https://vk.com/dev/wall.getComments
func (w *VkWall) GetComments(owner_id, item_id int) ([]Comment, error) {

	r, err := w.getComments(owner_id, item_id, 0)
	if err != nil {
		return nil, err
	}

	if len(r.Items) == r.Count {
		return r.Items, nil
	}

	all := make([]Comment, 0, r.Count)
	all = append(all, r.Items...)
	for {
		r, err := w.getComments(owner_id, item_id, len(all))
		if err != nil {
			return all, err
		}
		all = append(all, r.Items...)
		if len(all) >= r.Count || len(r.Items) == 0 {
			break
		}
	}
	return all, err
}

func (w *VkWall) getComments(owner_id, item_id, offset int) (*CommentResp, error) {

	data, err := w.getCommentsBody(owner_id, item_id, offset)
	if err != nil {
		return nil, err
	}

	var r CommentResp
	err = json.Unmarshal(data, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (w *VkWall) getCommentsBody(owner_id, post_id int, offset int) ([]byte, error) {

	p := make(map[string]string)

	p["method"] = "wall.getComments"
	p["v"] = "5.95"

	p["owner_id"] = strconv.Itoa(owner_id)
	p["post_id"] = strconv.Itoa(post_id)
	p["need_likes"] = "1" //1 — возвращать информацию о лайках.
	p["sort"] = "desc"
	//p["comment_id"] = ""	//TODO
	p["thread_items_count"] = "10"
	p["preview_length"] = "0" //количество символов, по которому нужно обрезать текст комментария. Укажите 0, если Вы не хотите обрезать текст.
	p["count"] = "100"
	if offset != 0 {
		p["offset"] = strconv.Itoa(offset)
	}

	return w.vk.GetBody(p, nil)

}

//=================

type WallResponse struct {
	Count int    `json:"count"`
	Items []Post `json:"items"`
}
type Post struct {
	ID          int          `json:"id"`
	FromID      int          `json:"from_id"`
	OwnerID     int          `json:"owner_id"`
	Date        int64        `json:"date"`
	MarkedAsAds int          `json:"marked_as_ads"`
	PostType    string       `json:"post_type"`
	Text        string       `json:"text"`
	IsPinned    int          `json:"is_pinned"`
	Attachments []Attachment `json:"attachments"`
	PostSource  PostSource   `json:"post_source"`
	Comments    Comments     `json:"comments"`
	Likes       Likes        `json:"likes"`
	Reposts     Reposts      `json:"reposts"`
	Views       Views        `json:"views"`
	IsFavorite  bool         `json:"is_favorite"`
}

type Attachment struct {
	Type  string `json:"type"`
	Photo *Photo `json:"photo,omitempty"`
	Doc   *Doc   `json:"doc,omitempty"`
}

type Doc struct {
	ID        int    `json:"id"`
	OwnerID   int    `json:"owner_id"`
	Title     string `json:"title"`
	Size      int    `json:"size"`
	EXT       string `json:"ext"`
	URL       string `json:"url"`
	Date      int64  `json:"date"`
	Type      int    `json:"type"`
	AccessKey string `json:"access_key"`
}

type Photo struct {
	ID        int    `json:"id"`
	AlbumID   int    `json:"album_id"`
	OwnerID   int    `json:"owner_id"`
	UserID    int    `json:"user_id"`
	Sizes     []Size `json:"sizes"`
	Text      string `json:"text"`
	Date      int64  `json:"date"`
	AccessKey string `json:"access_key"`
}

type Size struct {
	Type   string `json:"type"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type Comments struct {
	Count         int  `json:"count"`
	CanPost       int  `json:"can_post"`
	GroupsCanPost bool `json:"groups_can_post"`
}

type Likes struct {
	Count      int `json:"count"`
	UserLikes  int `json:"user_likes"`
	CanLike    int `json:"can_like"`
	CanPublish int `json:"can_publish"`
}

type PostSource struct {
	Type string `json:"type"`
}

type Reposts struct {
	Count        int `json:"count"`
	UserReposted int `json:"user_reposted"`
}

type Views struct {
	Count int `json:"count"`
}

//=== CommentResp ==============================================================
type CommentResp struct {
	Count             int       `json:"count"`
	Items             []Comment `json:"items"`
	CurrentLevelCount int       `json:"current_level_count"`
	CanPost           bool      `json:"can_post"`
	ShowReplyButton   bool      `json:"show_reply_button"`
	GroupsCanPost     bool      `json:"groups_can_post"`
}

type Comment struct {
	ID           int                `json:"id"`
	FromID       int                `json:"from_id"`
	PostID       int                `json:"post_id"`
	OwnerID      int                `json:"owner_id"`
	ParentsStack []interface{}      `json:"parents_stack"`
	Date         int64              `json:"date"`
	Text         string             `json:"text"`
	Likes        Likes              `json:"likes"`
	Thread       Thread             `json:"thread"`
	Attachments  []PurpleAttachment `json:"attachments"`
}

type PurpleAttachment struct {
	Type  string `json:"type"`
	Photo *Photo `json:"photo,omitempty"`
	Link  *Link  `json:"link,omitempty"`
}

type Thread struct {
	Count           int64        `json:"count"`
	Items           []ThreadItem `json:"items"`
	CanPost         bool         `json:"can_post"`
	ShowReplyButton bool         `json:"show_reply_button"`
	GroupsCanPost   bool         `json:"groups_can_post"`
}

type ThreadItem struct {
	ID             int                `json:"id"`
	FromID         int                `json:"from_id"`
	PostID         int                `json:"post_id"`
	OwnerID        int                `json:"owner_id"`
	ParentsStack   []int              `json:"parents_stack"`
	Date           int64              `json:"date"`
	Text           string             `json:"text"`
	Likes          Likes              `json:"likes"`
	ReplyToUser    int                `json:"reply_to_user"`
	ReplyToComment int                `json:"reply_to_comment"`
	Attachments    []FluffyAttachment `json:"attachments"`
}

type FluffyAttachment struct {
	Type  string `json:"type"`
	Video Video  `json:"video"`
}

type Video struct {
	ID          int    `json:"id"`
	OwnerID     int    `json:"owner_id"`
	Title       string `json:"title"`
	Duration    int    `json:"duration"`
	Description string `json:"description"`
	Date        int64  `json:"date"`
	Comments    int    `json:"comments"`
	Views       int    `json:"views"`
	LocalViews  int    `json:"local_views"`
	Photo130    string `json:"photo_130"`
	Photo320    string `json:"photo_320"`
	Photo800    string `json:"photo_800"`
	AccessKey   string `json:"access_key"`
	UserID      int    `json:"user_id"`
	Platform    string `json:"platform"`
	CanAdd      int    `json:"can_add"`
	TrackCode   string `json:"track_code"`
}
