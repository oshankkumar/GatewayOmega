package middlewares

import (
	"github.com/oshankkumar/GatewayOmega/services/auth"
	"github.com/oshankkumar/GatewayOmega/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
)

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.WithField("url", viper.GetString("services.auth.addr")).Infof("sending to auth service")
		authService := auth.NewAuthenticationService(http.DefaultClient)

		_, httpResp, err := authService.Account(utils.GetAuthTokenFromHeader(r.Header))
		if err != nil || httpResp.StatusCode != http.StatusOK {
			logrus.WithFields(logrus.Fields{
				"error":  err,
				"status": httpResp.Status,
			}).Errorf("user is not authenticated")
			return
		}

		logrus.WithField("status", httpResp.Status).Infof("user authentication successful")
		next.ServeHTTP(w, r)
	})
}
