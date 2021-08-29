package domain

type BlogDB interface {
	Store(blog Blog)
	Get(id int) (Blog, bool)
	ID() int
	Blogs() map[int]Blog
	Update(id int, blog Blog)
	Exist(id int) bool
	Delete(id int)
}
