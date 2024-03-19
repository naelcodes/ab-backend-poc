package api

import (
	"github.com/gofiber/fiber/v2"
	"neema.co.za/rest/utils/payloads"

	CustomErrors "neema.co.za/rest/utils/errors"
)

func (api *Api) EmailSignInHandler(c *fiber.Ctx) error {

	payload := c.Locals("payload").(*payloads.AuthSignInPayload)

	_, err := api.Service.EmailSignInService(&payload.User)

	if customError, isError := err.(*CustomErrors.CustomError); isError {
		if customError.Type == "NotFoundError" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials, user not found"})
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server error, sign up failed"})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Sign in successful"})
}

func (api *Api) EmailSignUpHandler(c *fiber.Ctx) error {

	payload := c.Locals("payload").(*payloads.AuthSignUpPayload)

	ok, err := api.Service.BeginEmailSignUpService(&payload.User)

	if customError, isError := err.(*CustomErrors.CustomError); isError {
		if customError.Type == "Domain Validation Error" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User with this email already exists"})
		}
	}

	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server error, sign up failed"})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"message": "Sign up successful"})
}

func (api *Api) CodeVerificationHandler(c *fiber.Ctx) error {

	code, _ := c.ParamsInt("code")
	user, err := api.Service.CompleteEmailSignUpService(code)
	if customError, isError := err.(*CustomErrors.CustomError); isError {
		if customError.Type == "NotFoundError" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "verification code expired"})
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server error, sign up failed"})
		}
	}
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"message": "Sign up successful", "user": user})
}

func (api *Api) FacebookAuthHandler(c *fiber.Ctx) error {
	url := api.Service.FacebookAuthService()
	return c.Status(fiber.StatusOK).Redirect(url, fiber.StatusTemporaryRedirect)
}

func (api *Api) FacebookAuthRedirectHandler(c *fiber.Ctx) error {

	state := c.Query("state")
	code := c.Query("code")

	isNewUser, user, err := api.Service.FacebookAuthRedirectService(code, state)

	if err != nil {
		return err
	}

	if isNewUser {
		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"user": user})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"message": "Sign in successful"})

}

func (api *Api) GoogleAuthHandler(c *fiber.Ctx) error {
	url := api.Service.GoogleAuthService()
	return c.Status(fiber.StatusOK).Redirect(url, fiber.StatusTemporaryRedirect)
}

func (api *Api) GoogleAuthRedirectHandler(c *fiber.Ctx) error {
	state := c.Query("state")
	code := c.Query("code")
	isNewUser, user, err := api.Service.GoogleAuthRedirectService(code, state)

	if err != nil {
		return err
	}

	if isNewUser {
		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"user": user})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"message": "Sign in successful"})

}
