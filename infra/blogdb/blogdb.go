package blogdb

import "github.com/worbridg/miniblog/domain"

// BlogDB is a temporary storage to store blogs.
type BlogDB struct {
	id    int
	blogs map[int]domain.Blog
}

func NewBlogDB() *BlogDB {
	return &BlogDB{
		id:    0,
		blogs: map[int]domain.Blog{},
	}
}

func (db *BlogDB) Store(blog domain.Blog) {
	db.id++
	db.blogs[db.id] = blog
}

func (db *BlogDB) Get(id int) (domain.Blog, bool) {
	blog, ok := db.blogs[id]
	return blog, ok
}

func (db *BlogDB) ID() int {
	return db.id
}

func (db *BlogDB) Blogs() map[int]domain.Blog {
	return db.blogs
}

func (db *BlogDB) Exist(id int) bool {
	_, ok := db.blogs[id]
	return ok
}

func (db *BlogDB) Update(id int, blog domain.Blog) {
	db.blogs[id] = blog
}

func (db *BlogDB) Delete(id int) {
	if db.Exist(id) {
		delete(db.blogs, id)
	}
}
