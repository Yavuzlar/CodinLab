package response

import (
	"errors"

	service_errors "github.com/Yavuzlar/CodinLab/internal/errors"
	"github.com/Yavuzlar/CodinLab/pkg/validator_service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

/*
ResponseHandler
Author: Resul Çelik
Customized fiber error handler for response automatic response handling.
*/
type BaseResponse struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	DataCount  uint64      `json:"data_count,omitempty"`
	Errors     interface{} `json:"errors,omitempty"`
}

func (r *BaseResponse) Error() string {
	return r.Message
}

func ResponseHandler(c *fiber.Ctx, err error) error {
	base := &BaseResponse{}
	//BaseResponse
	if errors.As(err, &base) {
		return c.Status(err.(*BaseResponse).StatusCode).JSON(err)
	}

	//validation errors
	if errors.As(err, &validator.ValidationErrors{}) {
		errs := validator_service.ValidatorErrors(err)
		return c.Status(400).JSON(
			&BaseResponse{
				StatusCode: 400,
				Message:    "validation error",
				Errors:     errs,
			},
		)
	}

	//fiber errors
	fiberErr := &fiber.Error{}
	if errors.As(err, &fiberErr) {
		if fiberErr.Code == 404 {
			return c.Status(404).JSON(&BaseResponse{
				StatusCode: 404,
				Message:    "not found",
			})
		} else {
			return c.Status(err.(*fiber.Error).Code).JSON(&BaseResponse{
				StatusCode: err.(*fiber.Error).Code,
				Message:    err.(*fiber.Error).Message,
			})
		}
	}

	//service errors
	serviceErr := &service_errors.ServiceError{}
	if errors.As(err, &serviceErr) {
		resp := &BaseResponse{
			StatusCode: serviceErr.Code,
			Message:    serviceErr.Message,
		}
		if serviceErr.Error() != "" {
			resp.Errors = serviceErr.Error()
		}
		return c.Status(serviceErr.Code).JSON(resp)
	}

	//unknown errors
	return c.Status(500).JSON(&BaseResponse{
		StatusCode: 500,
		Message:    "Internal Server Error (Unknown)",
		Errors:     err.Error(),
	})
}

// Response function for create a new response.
func Response(statusCode int, message string, data interface{}, dataCount ...uint64) error {
	var count uint64
	if len(dataCount) > 0 {
		count = dataCount[0]
	}
	return &BaseResponse{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
		DataCount:  count,
	}
}
