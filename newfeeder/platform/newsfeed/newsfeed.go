package newsfeed

import "fmt"

type Item struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

type Repo struct {
	Items []Item
}

func New() *Repo {

	return &Repo{}
}

func (r *Repo) Add(item Item) {
	fmt.Println(item, "received Item")
	r.Items = append(r.Items, item)
	fmt.Println(r.Items, "after append")
}

func (r *Repo) GetAll() []Item {
	return r.Items
}
