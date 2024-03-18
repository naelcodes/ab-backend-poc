package api

import "github.com/gofiber/fiber/v2"

func (api *Api) EmailSignInHandler(c *fiber.Ctx) error {
	return nil
}

func (api *Api) EmailSignUpHandler(c *fiber.Ctx) error {
	return nil
}

func (api *Api) EmailVerificationHandler(c *fiber.Ctx) error {
	return nil
}

func (api *Api) FacebookAuthHandler(c *fiber.Ctx) error {
	url := api.Service.FacebookAuthService()
	return c.Status(fiber.StatusOK).Redirect(url, fiber.StatusTemporaryRedirect)
}

func (api *Api) FacebookAuthRedirectHandler(c *fiber.Ctx) error {

	state := c.Query("state")
	code := c.Query("code")

	result, err := api.Service.FacebookAuthRedirectService(code, state)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"result": result})
}

func (api *Api) GoogleAuthHandler(c *fiber.Ctx) error {
	url := api.Service.GoogleAuthService()
	return c.Status(fiber.StatusOK).Redirect(url, fiber.StatusTemporaryRedirect)
}

func (api *Api) GoogleAuthRedirectHandler(c *fiber.Ctx) error {
	state := c.Query("state")
	code := c.Query("code")
	result, err := api.Service.GoogleAuthRedirectService(code, state)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"result": result})

}
