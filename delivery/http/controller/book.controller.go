package controller

import (
	"favorite-book/delivery/http/dto/request"
	"favorite-book/domain"
	"favorite-book/shared/util"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type BookController struct {
	domain domain.Domain
}

func NewBookController(domain domain.Domain) BookController {
	return BookController{
		domain: domain,
	}
}

func (this *BookController) SaveBook(ctx *fiber.Ctx) error {
	var book request.RequestBookDTO

	if err := ctx.BodyParser(&book); err != nil {
		resp, statusCode := util.ConstructResponseError(fiber.StatusBadRequest, "Invalid request body")
		return ctx.Status(statusCode).JSON(resp)
	}

	if err := this.domain.BookUsecase.SaveBook(book.ToBookEntity()); err != nil {
		resp, statusCode := util.ConstructResponseError(fiber.StatusBadRequest, "Failed to save book")
		return ctx.Status(statusCode).JSON(resp)
	}

	return ctx.Redirect("/")
}

func (this *BookController) GetAllBook(ctx *fiber.Ctx) error {
	book, err := this.domain.BookUsecase.GetAllBook()
	fmt.Println("list of book -> ", book)

	if err != nil {
		resp, statusCode := util.ConstructResponseError(fiber.StatusBadRequest, "Failed to fetch book")
		return ctx.Status(statusCode).JSON(resp)
	}

	return ctx.Render("index", fiber.Map{
		"Books": *book,
	})
}
