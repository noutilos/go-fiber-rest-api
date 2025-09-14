package factory

type Service interface {
	Create() // Generic type should be there as input
	// Update(ctx *fiber.Ctx) error
	// Delete(ctx *fiber.Ctx) error
	// Get(ctx *fiber.Ctx) error
	// GetAll(ctx *fiber.Ctx) error
}