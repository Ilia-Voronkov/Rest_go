package messagesService

import "gorm.io/gorm"

type MessageRepository interface {
	// CreateMessage - создаем сообщение
	CreateMessage(message Message) (Message, error)
	// GetAllMessages - возвращаем все сообщения
	GetAllMessages() ([]Message, error)
	// UpdateMessageByID - обновляем сообщение по ID
	UpdateMessage(id int, newMessage Message) (Message, error)
	// DeleteMessageByID - удаляем сообщение по ID
	DeleteMessageByID(id int) error
}

type messageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *messageRepository {
	return &messageRepository{db: db}
}

func (r *messageRepository) CreateMessage(message Message) (Message, error) {
	err := r.db.Create(&message).Error
	if err != nil {
		return Message{}, err
	}
	return message, nil
}

func (r *messageRepository) GetAllMessages() ([]Message, error) {
	var messages []Message
	err := r.db.Find(&messages).Error
	return messages, err
}

func (r *messageRepository) UpdateMessage(id int, newMessage Message) (Message, error) {
	var message Message
	// Найдем сообщение по ID
	err := r.db.First(&message, id).Error
	if err != nil {
		return Message{}, err
	}

	// Обновим сообщение
	message.Text = newMessage.Text
	err = r.db.Save(&message).Error
	if err != nil {
		return Message{}, err
	}
	return message, nil
}

func (r *messageRepository) DeleteMessageByID(id int) error {
	var message Message
	err := r.db.First(&message, id).Error
	if err != nil {
		return err
	}
	return r.db.Delete(&message).Error
}
