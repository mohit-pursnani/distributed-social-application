package models

type Post struct {
	UserId    int
	PostText  string
	FirstName string
	LastName  string
}

type UserPostList []Post

func CreateEmptyUserPostList() UserPostList {
	return make(UserPostList, 0)
}

func (upl *UserPostList) AddPost(userId int, postText string, firstName string, lastName string) Post {

	post := Post{UserId: userId, PostText: postText, FirstName: firstName, LastName: lastName}
	*upl = append(*upl, post)

	return post
}

func (upl *UserPostList) GetFolloweePosts(user UserStruct) []Post {
	followeePosts := CreateEmptyUserPostList()

	for _, userPost := range *upl {
		for _, user := range user.Following {
			if userPost.UserId == user.IdNum {
				var postListObj = Post{UserId: user.IdNum, PostText: userPost.PostText, FirstName: user.FirstName, LastName: user.LastName}
				followeePosts = append(followeePosts, postListObj)
			}
		}
	}
	return followeePosts
}
