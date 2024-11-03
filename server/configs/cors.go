package configs

import "github.com/spf13/viper"

func GetCorsAllowedOrigins() []string {
	corsAllowedOrigins := viper.GetStringSlice("CORS_ALLOWED_ORIGINS")
	if len(corsAllowedOrigins) == 0 {
		return []string{"https://*", "http://*", "ws://*", "wss://*"} // Fallback setting CORs for all domain
	}
	return corsAllowedOrigins
}

func GetCorsAllowedMethods() []string {
	corsAllowedMethods := viper.GetStringSlice("CORS_ALLOWED_METHODS")
	if len(corsAllowedMethods) == 0 {
		return []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	}
	return corsAllowedMethods
}

func GetCorsAllowedHeaders() []string {
	corsAllowedHeader := viper.GetStringSlice("CORS_ALLOWED_HEADERS")
	if len(corsAllowedHeader) == 0 {
		return []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"}
	}
	return corsAllowedHeader
}

func GetCorsExposedHeaders() []string {
	corsExposedHeader := viper.GetStringSlice("CORS_EXPOSED_HEADERS")
	if len(corsExposedHeader) == 0 {
		return []string{"Link"}
	}
	return corsExposedHeader
}

func GetCorsMaxAge() int {
	corsMaxAge := viper.GetInt("CORS_MAX_AGE")
	if corsMaxAge == 0 {
		return 300
	}
	return corsMaxAge
}
