package data


import (
	"fmt"
	"github.com/punkers01/restful/models"
)

type Cache struct{
	userMap map[int]models.User
}


func NewCache() *Cache{
	umap := make(map[int]models.User)
	return &Cache{umap}
}

func (c *Cache)AddUser(u models.User ,  id int) {
     c.userMap[id]=u
	fmt.Println("user added ", u)
}

func (c *Cache)GetUser(id int) models.User{
	fmt.Println("returning user  ", c.userMap[id])
	return c.userMap[id]
}