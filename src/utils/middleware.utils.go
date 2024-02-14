package utils

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	jwtware "github.com/gofiber/jwt/v2"
)

type Middleware struct{}

type IMiddleware interface {
	AuthorizationRequired(jwtSecret string) fiber.Handler
	InitialMiddlewareRequestId(header string) (iRequestId fiber.Handler)
	InitialLimiter(numOfLimit int) (iLimiter fiber.Handler)
}

func NewMiddleware() IMiddleware {
	return &Middleware{}
}

func (r *Middleware) InitialMiddlewareRequestId(header string) (iRequestId fiber.Handler) {
	if header != "" {
		iRequestId = requestid.New(requestid.Config{
			Header: header,
			Generator: func() string {
				return GetUUID()
			},
		})
	} else {
		iRequestId = requestid.New()
	}
	return
}

func (r *Middleware) InitialLimiter(numOfLimit int) (iLimiter fiber.Handler) {
	if numOfLimit == 0 {
		numOfLimit = 20
	}
	iLimiter = limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1"
		},
		Max:        numOfLimit,
		Expiration: 30 * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.Get("x-forwarded-for")
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Rate Limit Too Many Requests"})
		},
	})
	return
}

func (r *Middleware) AuthorizationRequired(jwtSecret string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SuccessHandler: AuthSuccess,
		ErrorHandler:   AuthError,
		SigningKey:     []byte(jwtSecret),
		SigningMethod:  "HS256",
	})
}

func AuthSuccess(c *fiber.Ctx) error {
	c.Next()
	return nil
}

func AuthError(c *fiber.Ctx, e error) error {
	c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "Unauthorized",
		"msg":   e.Error(),
	})
	return nil
}
