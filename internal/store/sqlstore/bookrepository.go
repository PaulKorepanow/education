package sqlstore

import (
	"bookLibrary/internal/model"
)

type BookRepository struct {
	store *SqlStore
}

func (b *BookRepository) Create(book *model.Book) error {
	if err := b.store.db.Create(book).Error; err != nil {
		return err
	}
	return nil
}

func (b *BookRepository) FindByTitle(title string) (*model.Book, error) {
	var book model.Book
	err := b.store.db.Where("title = ?", title).First(&book).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}
