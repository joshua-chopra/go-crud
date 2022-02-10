package models

// no need to specify JSON mappings since
// struct
type Book struct {
	ID     uint   `gorm:"primary_key"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
	Rating int    `json:"rating"`
}

func (b Book) isEmpty() bool {
	return b.ID == 0 && b.Title == ""
}
