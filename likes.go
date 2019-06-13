package vk

import (
	"encoding/json"
	"strconv"
)

//https://vk.com/dev/likes
type VkLike struct {
	vk *VK
}

//TODO GetListFromComment
func (l *VkLike) GetListFromComment(owner_id, item_id int) ([]int, error) {
	return nil, nil
}

//likes.getList
//Получает список идентификаторов пользователей, которые добавили заданный объект в свой список Мне нравится.
//https://vk.com/dev/likes.getList
func (l *VkLike) GetList(owner_id, item_id int) ([]int, error) {

	r, err := l.getList(owner_id, item_id, 0)
	if err != nil {
		return nil, err
	}

	if len(r.Items) == r.Count {
		return r.Items, nil
	}

	all := make([]int, 0, r.Count)
	all = append(all, r.Items...)
	for {
		r, err := l.getList(owner_id, item_id, len(all))
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

func (l *VkLike) getList(owner_id, item_id, offset int) (*ResponseInt, error) {

	data, err := l.getListBody(owner_id, item_id, offset)
	if err != nil {
		return nil, err
	}

	var r ResponseInt
	err = json.Unmarshal(data, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (l *VkLike) getListBody(owner_id, item_id int, offset int) ([]byte, error) {

	p := make(map[string]string)
	p["method"] = "likes.getList"
	p["v"] = "5.95"

	p["type"] = "post" //TODO "comment"
	p["owner_id"] = strconv.Itoa(owner_id)
	p["item_id"] = strconv.Itoa(item_id)
	p["filter"] = "likes"
	p["count"] = "1000"
	if offset != 0 {
		p["offset"] = strconv.Itoa(offset)
	}
	return l.vk.GetBody(p, nil)
}
