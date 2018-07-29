package middlewares

import (
	"github.com/oshankkumar/GatewayOmega/services/auth"
	"github.com/oshankkumar/GatewayOmega/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"github.com/opentracing/opentracing-go"
)

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		span,ctx := opentracing.StartSpanFromContext(r.Context(),"authentication")
		ctx = opentracing.ContextWithSpan(ctx,span)
		r = r.WithContext(ctx)
		logrus.WithField("url", viper.GetString("services.auth.addr")).Infof("sending to auth service")
		authService := auth.NewAuthenticationService(http.DefaultClient)

		_, httpResp, err := authService.Account(utils.GetAuthTokenFromHeader(r.Header))
		if err != nil || httpResp.StatusCode != http.StatusOK {
			span.LogKV("error",err,"authentication","unsuccessful")
			logrus.WithFields(logrus.Fields{
				"error":  err,
				"status": httpResp.Status,
			}).Errorf("user is not authenticated")
			return
		}

		logrus.WithField("status", httpResp.Status).Infof("user authentication successful")
		span.LogKV("http-status", httpResp.Status,"authentication","successful")
		span.Finish()
		next.ServeHTTP(w, r)
	})
}
