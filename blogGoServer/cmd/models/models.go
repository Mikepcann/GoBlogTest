package models

import(
	"github.com/google/uuid"
	"errors"
) 



type Post struct {
	Title string `json:"title"`
	Body string `json:"body"`
	ID uuid.UUID `json:"id"` 
}

type InMemoryPosts struct {
	List []Post
}

var Instance InMemoryPosts = InMemoryPosts{[]Post{}}


func GetMemory() *InMemoryPosts{
	return &Instance
}

// add post to List
func (i *InMemoryPosts) AddPost(p Post, id uuid.UUID, ){
	p.ID = id
	i.List = append(i.List,p )
}

// view all posts
func (i *InMemoryPosts) GetAllPosts() []Post {
	return i.List
}

// delete all posts
func (i *InMemoryPosts) DeleteAllPosts(){
	i.List = []Post{}
}

func (i *InMemoryPosts) DeleteById(id uuid.UUID) error{
	for index, el := range i.List {
		if el.ID == id {
			 i.List = append(i.List[:index], i.List[index+1:]...)
			return nil
		}
	}
	err := errors.New("ID Not found")
	return err 
}
