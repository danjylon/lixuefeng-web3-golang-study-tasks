package dao

import (
	. "gorm_demo/config"
	. "gorm_demo/model"
)

func CreatePost(posts []Post) {
	GetDB().CreateInBatches(posts, len(posts))
}

func GetPosts(id int) (posts []Post) {
	GetDB().Preload("Comments").Where("user_id = ?", id).Find(&posts)
	return
}

/*
*
select posts.*, count(comments.id) comments_count from posts
left join comments on posts.id = comments.post_id
group by comments.post_id order by comments_count desc
*/
func GetMostComentsPost() (post Post) {
	GetDB().Model(&Post{}).Select("posts.*,count(comments.id) comments_count").
		Joins("left join comments on comments.post_id = posts.id").
		Group("comments.post_id").
		Order("comments_count desc").First(&post)
	return
}
