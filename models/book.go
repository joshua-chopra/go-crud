package models

// no need to specify JSON mappings since
// struct
type Book struct {
	// need to be explicit w/ GORM such that if passed in struct
	// for creation has no ID key, we need to generate one
	// and use that for insertion.
	ID     uint   `gorm:"primary_key;auto_increment;not_null"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
	Rating int    `json:"rating"`
}

func (b Book) isEmpty() bool {
	return b.ID == 0 && b.Title == ""
}
