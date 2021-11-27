package models

import (
	"fmt"
	"time"
)

type Author struct {
	name     *string
	birthday *time.Time
	articles []*Article
}

func NewAuthor(name *string, birthday *time.Time) *Author {
	author := new(Author)
	author.name = name
	author.birthday = birthday
	return author
}

func NewAuthorWithAuthor(authorToClone *Author) *Author {
	author := new(Author)
	author.name = authorToClone.name
	author.birthday = authorToClone.birthday
	author.articles = authorToClone.articles
	return author
}

func (a Author) clone() *Author {
	return NewAuthorWithAuthor(&a)
}

func (a Author) String() string {
	return fmt.Sprintf("name: %s, birthday: %s", *a.name, a.birthday)
}

func (a *Author) AddArticle(article *Article)  {
	a.articles = append(a.articles, article)
}

func (a Author) GetArticles() []*Article {
	return a.articles
}
