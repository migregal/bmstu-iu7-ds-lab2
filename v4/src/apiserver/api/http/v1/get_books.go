package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type BooksRequest struct {
	PaginatedRequest `valid:"optional"`

	ShowAll   bool   `query:"showAll" valid:"optional"`
	LibraryID string `param:"id" valid:"uuidv4,required"`
}

type BooksResponse struct {
	PaginatedResponse

	Items []Book `json:"items"`
}

func (a *api) GetLibraryBooks(c echo.Context, req BooksRequest) error {
	books, err := a.core.GetLibraryBooks(c.Request().Context(), req.LibraryID, req.ShowAll, req.Page, req.Size)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	resp := BooksResponse{
		PaginatedResponse: PaginatedResponse{
			Page:     req.Page,
			PageSize: req.Size,
			Total:    books.Total,
		},
		Items: make([]Book, 0, len(books.Items)),
	}

	for _, book := range books.Items {
		resp.Items = append(resp.Items, Book(book))
	}

	return c.JSON(http.StatusOK, &resp)
}
