package sqlstore_test

import (
	"bookLibrary/internal/model"
	"bookLibrary/internal/store/sqlstore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBookRepository_Create(t *testing.T) {
	tb, teardown := sqlstore.TestStore(t, databaseURL)
	defer teardown("books")

	bookRep := tb.Book()
	err := bookRep.Create(model.TestBook(t))
	assert.NoError(t, err)
}

func TestBookRepository_FindByTitle(t *testing.T) {
	tb, teardown := sqlstore.TestStore(t, databaseURL)
	defer teardown("books")

	err := tb.Book().Create(model.TestBook(t))
	assert.NoError(t, err)

	book, err := tb.Book().FindByTitle("SimpleBook")
	assert.NoError(t, err)
	assert.Equal(t, "SimpleBook", book.Title)
}
