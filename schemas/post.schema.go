package schemas

type CreatePost struct {
	//Эта структура будет использоваться Gin Gonic для проверки
	//полезной нагрузки запроса при добавлении новых записей в базу данных.
	Title    string `json:"title" binding:"required"`
	Category string `json:"category" binding:"required"`
	Content  string `json:"content" binding:"required"`
}

type UpdatePost struct {
	// Эта структура будет использоваться Gin Gonic
	//для проверки полезной нагрузки запроса при обновлении записей в базе данных.
	Title    string `json:"title" binding:"required"`
	Category string `json:"category" binding:"required"`
	Content  string `json:"content" binding:"required"`
}
