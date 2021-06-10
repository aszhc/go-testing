package main

type Article struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type Blog struct {
	Articles []Article
}

func New() *Blog {
	return &Blog{}
}

// SaveArticle 保存文章
func (b *Blog) SaveArticle(article Article) {
	b.Articles = append(b.Articles, article)
}

// FetchAll 获取所有文章
func (b *Blog) FetchAll() []Article {
	return b.Articles
}
