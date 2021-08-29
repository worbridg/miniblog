package domain

import "errors"

type Blog struct {
	Title   string
	Content string
}

func NewBlog(title, content string) (Blog, error) {
	if title == "" {
		return Blog{}, errors.New("no title")
	}
	if content == "" {
		return Blog{}, errors.New("no content")
	}

	return Blog{Title: title, Content: content}, nil
}
