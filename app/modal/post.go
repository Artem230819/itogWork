package modal

type Post struct {
	Id       int
	Text     string
}

func GetAllPosts () (posts []Post, err error) {
	posts = []Post{
		{1,"Джон"},
		{2,"Говард"},
		{3,"Джек"},
		{4,"Лизель"},
		{5,"Джейн"},
		{6,"Мартин"},
		{7,"Джон"},
		{8,"Сэмвелл"},
		{9,"Гермиона"},
	}
	return
}
