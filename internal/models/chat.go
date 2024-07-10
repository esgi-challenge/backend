package models

type Message struct {
	GormModel
	Content   string `json:"content" gorm:"column:content"`
	ChannelId uint   `json:"channelId" gorm:"column:channel_id"`
	SenderId  uint   `json:"senderId" gorm:"column:sender_id"`
}

type Channel struct {
	GormModel
	FirstUserId  uint      `json:"firstUserId" gorm:"column:first_user_id"`
	SecondUserId uint      `json:"secondUserId" gorm:"column:second_user_id"`
	FirstUser    User      `json:"firstUser" gorm:"foreignKey:FirstUserId;references:ID"`
	SecondUser   User      `json:"secondUser" gorm:"foreignKey:SecondUserId;references:ID"`
	Messages     []Message `json:"messages" gorm:"foreignKey:ChannelId"`
}

type ChannelCreate struct {
	FirstUserId  uint `json:"firstUserId"`
	SecondUserId uint `json:"secondUserId"`
}
