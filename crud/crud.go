package crud

import (
	"fmt"
	"github.com/MarcosIgnacioo/personahtmx/initializers"
	"github.com/MarcosIgnacioo/personahtmx/models"
)

func Get(username string) *models.User {
    var user models.User;
    initializers.DB.First(&user,"username = ?", username)
    return &user
}

type result struct {
    ParentAuthor string `gorm:"column:p_author"`
    ParentContent string `gorm:"column:p_content"`
    ReplyAuthor string `gorm:"column:reply_author"`
    ReplyContent  string `gorm:"column:reply_content"`
}

func (r *result) String() string  {
    return fmt.Sprintf("PA: %v\nPC: %v\n///\nRA: %v\nRC: %v", r.ParentAuthor, r.ParentContent,r.ReplyAuthor ,r.ReplyContent)
}

func GetAll(offset int)  [] models.Tweet {
    var Tweets [] models.Tweet;
    //TODO modificar el struct de tweet para que guarde un Tweet y que este sea la ultima respuesstaaa
    var result result
    // var tl models.TweetLayout
    
    initializers.
        DB.Table("tweets").
        Select("tweets.content as p_content, tweets.author as p_author, tweet_replies.content as reply_content, tweet_replies.author as reply_author").
        Joins("RIGHT JOIN tweet_replies ON tweet_replies.parent_tweet_id = tweets.id").
        Order("tweets.updated_at DESC").
	    Limit(2).
	    Offset(offset).
        Scan(&result)

    initializers.DB.Order("updated_at DESC").Limit(2).Offset(offset).Find(&Tweets)

    return Tweets
}

func GetReplies()  {
    //TODO funcion para obtener las reespuestas en base a un tweet haciendo ese join
    // Le pasamos el tweet como parametro y boom
    // db.Joins("tweet_replies").Find(&tweet)
}
