package middleware

// import (
//         "time"
//
//         "github.com/dgrijalva/jwt-go"
//
//         "github.com/gofiber/fiber"
//         jwtware "github.com/gofiber/jwt"
// )
//
// type Authorizer struct {
//         secret []byte
// }
//
// func NewAuthorizer(secret []byte) *Authorizer {
//         return &Authorizer{secret}
// }
//
// func (a *Authorizer) Middleware() func(c *fiber.Ctx) {
//         return jwtware.New(jwtware.Config{
//                 SigningKey: a.secret,
//         })
// }
//
// func (a *Authorizer) Sign(c map[string]interface{}) (string, error) {
//         // Create token
//         token := jwt.New(jwt.SigningMethodHS256)
//
//         // Set claims
//         claims := token.Claims.(jwt.MapClaims)
//         claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
//         for k, v := range c {
//                 claims[k] = v
//         }
//
//         // Generate encoded token and send it as response.
//         t, err := token.SignedString(a.secret)
//         return t, err
// }
//
// func (a *Authorizer) ClaimsFromContext(c *fiber.Ctx) map[string]interface{} {
//         user := c.Locals("user").(*jwt.Token)
//         claims := user.Claims.(jwt.MapClaims)
//         return claims
// }
