package messagesService

import "gorm.io/gorm"

// MessageRepository - интерфейс репозитория для сообщений
type MessageRepository interface {
	CreateMessage(message Message) (Message, error)
	GetAllMessages() ([]Message, error)
	UpdateMessage(id int, newMessage Message) (Message, error)
	DeleteMessageByID(id int) error
}

type messageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) MessageRepository {
	return &messageRepository{db: db}
}

func (r *messageRepository) CreateMessage(message Message) (Message, error) {
	err := r.db.Create(&message).Error
	return message, err
}

func (r *messageRepository) GetAllMessages() ([]Message, error) {
	var messages []Message
	err := r.db.Find(&messages).Error
	return messages, err
}

func (r *messageRepository) UpdateMessage(id int, newMessage Message) (Message, error) {
	var message Message
	err := r.db.First(&message, id).Error
	if err != nil {
		return Message{}, err
	}
	message.Text = newMessage.Text
	err = r.db.Save(&message).Error
	return message, err
}

func (r *messageRepository) DeleteMessageByID(id int) error {
	return r.db.Delete(&Message{}, id).Error
}
