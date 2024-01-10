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
func GetAll(offset int)  [] models.Tweet {
    fmt.Println("asdfdasfasdfasdfasdfhjasdhjfhajsdf")
    var Tweets [] models.Tweet;
    //TODO modificar el struct de tweet para que guarde un Tweet y que este sea la ultima respuesstaaa
    var result []struct {
        ParentContent string `gorm:"column:parent_content"`
        ReplyContent  string `gorm:"column:reply_content"`
    }

    initializers.DB.Table("tweets").
        Select("tweets.content as parent_content, tweet_replies.content as reply_content").
        Joins("RIGHT JOIN tweet_replies ON tweet_replies.parent_tweet_id = tweets.id").
	    Limit(2).
	    Offset(offset).
        Scan(&result)

    initializers.DB.Order("updated_at desc").Limit(2).Offset(offset).Find(&Tweets)

    fmt.Println("/////")
    fmt.Println(result)
    fmt.Println("/////")
    return Tweets
}

func GetReplies()  {
    //TODO funcion para obtener las reespuestas en base a un tweet haciendo ese join
    // Le pasamos el tweet como parametro y boom
    // db.Joins("tweet_replies").Find(&tweet)
}
