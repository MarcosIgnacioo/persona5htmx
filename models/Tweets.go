package models
type Tweets struct {
    Tweets [] Tweet
    Replies [] TweetReply
    Start int
    Next int
    More  bool
    Template string
}
