// Code generated by gozz:doc. DO NOT EDIT.

package books_all_in_one

var (
	_types_doc = map[interface{}]map[string]string{
		(*BookService)(nil): _doc_BookService,

		(*QueryBook)(nil): _doc_QueryBook,
		(*FormBook)(nil):  _doc_FormBook,
		(*DataBook)(nil):  _doc_DataBook,
		(*ListBook)(nil):  _doc_ListBook,
	}

	_values_doc = map[string]map[interface{}]string{
		"book_type": map[interface{}]string{
			BookTypeAdventureStories:  "",
			BookTypeClassics:          "",
			BookTypeCrime:             "",
			BookTypeFantasy:           "",
			BookTypeHistoricalFiction: "",
			BookTypeHorror:            "",
			BookTypeHumourAndSatire:   "",
			BookTypeLiteraryFiction:   "",
			BookTypeMystery:           "",
			BookTypePoetry:            "",
			BookTypePlays:             "",
			BookTypeRomance:           "",
			BookTypeScienceFiction:    "",
			BookTypeShortStories:      "",
			BookTypeThrillers:         "",
			BookTypeWar:               "",
			BookTypeWomenFiction:      "",
			BookTypeYoungAdult:        "",
		},
	}
)

var _doc_BookService = map[string]string{
	"":       "BookService provide book management services",
	"List":   "List all books. return ListBook",
	"Get":    "Get book by book id. returns DataBook",
	"Create": "Create new book from FormBook. returns DataBook created",
	"Edit":   "Edit book by id from FormBook. returns DataBook edited",
}

var _doc_QueryBook = map[string]string{
	"":          "QueryBook represents struct for querying book list or book item",
	"Id":        "specify query id",
	"Title":     "specify query title keywords",
	"PageNo":    "specify query pagination page no. default: 1",
	"PageCount": "specify query pagination page count. default: 20",
}

func (QueryBook) FieldDoc(f string) string { return _doc_QueryBook[f] }

var _doc_FormBook = map[string]string{
	"":            "FormBook represents struct for creating or editing book",
	"Id":          "specify editing id. only works for editing",
	"Title":       "book title",
	"Type":        "book type",
	"Description": "book description",
	"Author":      "book author",
}

func (FormBook) FieldDoc(f string) string { return _doc_FormBook[f] }

var _doc_DataBook = map[string]string{
	"":          "DataBook represents struct for book item",
	"CreatedAt": "book created time",
	"CreatedBy": "book created username",
	"UpdatedAt": "book last updated time",
	"UpdatedBy": "book last updated username",
}

func (DataBook) FieldDoc(f string) string { return _doc_DataBook[f] }

var _doc_ListBook = map[string]string{
	"":      "ListBook represents struct form querying book list and total count of books in datastore",
	"List":  "query book results",
	"Total": "total count of books in datastore",
}

func (ListBook) FieldDoc(f string) string { return _doc_ListBook[f] }
