package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// RequireAuth ek middleware hai jo har request mein JWT token check karega
func RequireAuth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Client se 'Authorization' header maango
		authHeader := c.GetHeader("Authorization")

		// 2. Check karo ki header hai ya nahi, aur usme "Bearer " likha hai ya nahi
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Missing or invalid Bearer token"})
			c.Abort() // Request ko yahi rok do
			return
		}

		// 3. "Bearer " word ko hata kar asli token nikalo
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// 4. Token ko open karke check karo (Signature aur Expiry)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, http.ErrNotSupported
			}
			return []byte(secret), nil
		})

		// 5. Agar token fake hai ya expire ho chuka hai
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Invalid or expired token"})
			c.Abort()
			return
		}

		// 6. Token ke andar se user ki details nikalo aur Context me save kar do
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("user_id", claims["user_id"])
			c.Set("email", claims["email"])
			c.Set("role", claims["role"])
		}

		// 7. Sab kuch sahi hai, request ko aage badhne do
		c.Next()
	}
}