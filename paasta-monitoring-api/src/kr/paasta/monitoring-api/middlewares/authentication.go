package middlewares

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"paasta-monitoring-api/connections"
	v1service "paasta-monitoring-api/services/api/v1"

	"github.com/labstack/echo/v4"
)

func CheckToken(conn connections.Connections) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			logger := c.Request().Context().Value("LOG").(*logrus.Entry)
			mode := os.Getenv("mode")
			if mode == "development" {
				c.Set("userId", "admin")  // For debug
				err := next(c)
				return err
			}

			// 1. 토큰 추출
			bearToken, err := v1service.ExtractToken(c)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}

			// 2. 토큰 검증 (signing method 검증, 서명 검증)
			token, err := v1service.VerifyToken(bearToken, "ACCESS_SECRET", c)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
			}

			// 3. 토큰 만료 검증 (동작 방식 확인 필요)
			err = v1service.TokenValid(token)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
			}

			// 4. 메타 데이터 추출 (메타데이터 이용한 Redis 확인)
			metadata, err := v1service.ExtractTokenMetadata(token, "ACCESS")
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			logger.Info(metadata)

			// 5. Redis 추출 (UUID를 이용한 userId 추출)
			userId, err := v1service.FetchAuth(metadata, conn.RedisInfo)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
			}
			logger.Info(userId + " is authorized")
			c.Set("userId", userId)

			// Continue
			if err := next(c); err != nil {
				return err
			}

			return nil
		}
	}
}
