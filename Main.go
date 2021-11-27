package main

import "awesomeProject1/app"

/*
func main() {
	birthday := time.Date(2020, time.April,
		11, 21, 34, 01, 0, time.UTC)
	name := "value"
	author := models.NewAuthor(&name, &birthday)

	println(author.String())
	name = "other "

	articleName := ""
	articleContent := ""

	article := models.NewArticle(&articleName, &articleContent, author)

	author.AddArticle(article)
	article.GetAuthor().AddArticle(article)
	article.GetAuthor().AddArticle(article)

	for _, a := range article.GetAuthor().GetArticles() {
		println(a.String())
	}

	server()
}

func server() {
	server := server2.NewServer(1293)
	routeName := "/hello"

	handler := func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w,"")
	}
	server.AddRoute(routeName, handler)
	server.Start()
}*/

func main() {
	app.Start()
}
