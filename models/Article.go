package models

import "fmt"

type Article struct {
	title   *string
	content *string
	author  *Author
}

func NewArticle(title *string, content *string, author *Author) *Article {
	article := new(Article)
	article.title = title
	article.content = content
	article.author = author
	return article
}

func (a Article) String() string {
	return fmt.Sprintf("title: %s, content: %s, author: %s", *a.title, *a.content, a.author.String())
}

func (a *Article) GetAuthor() *Author {
	return a.author
}
