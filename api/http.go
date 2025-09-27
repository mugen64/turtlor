package api

import (
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"github.com/go-chi/httplog/v3"
	"github.com/mugen64/turtlor/api/handlers"
	"github.com/mugen64/turtlor/configs"
	"github.com/mugen64/turtlor/pkg/apperrors"
	"github.com/mugen64/turtlor/pkg/logger"
	"github.com/mugen64/turtlor/pkg/utils"
)

func NewServer(cfg *configs.Config, logger *logger.Logger) http.Handler {

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(httplog.RequestLogger(logger.GetLogger(), &httplog.Options{
		// Level defines the verbosity of the request logs:
		// slog.LevelDebug - log all responses (incl. OPTIONS)
		// slog.LevelInfo  - log responses (excl. OPTIONS)
		// slog.LevelWarn  - log 4xx and 5xx responses only (except for 429)
		// slog.LevelError - log 5xx responses only
		Level: logger.GetLogLevel(),

		// Set log output to Elastic Common Schema (ECS) format.
		Schema: httplog.SchemaECS,

		// RecoverPanics recovers from panics occurring in the underlying HTTP handlers
		// and middlewares. It returns HTTP 500 unless response status was already set.
		//
		// NOTE: Panics are logged as errors automatically, regardless of this setting.
		RecoverPanics: true,

		// Optionally, filter out some request logs.
		Skip: func(req *http.Request, respStatus int) bool {
			return respStatus == 404 || respStatus == 405
		},

		// Optionally, log selected request/response headers explicitly.
		LogRequestHeaders:  []string{"Origin"},
		LogResponseHeaders: []string{},

		// Optionally, enable logging of request/response body based on custom conditions.
		// Useful for debugging payload issues in development.
		LogRequestBody:  utils.IsDebugHeaderSet,
		LogResponseBody: utils.IsDebugHeaderSet,
	}))
	r.Use(middleware.Recoverer)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// r.Get("/swagger/*", httpSwagger.Handler(
	// 	httpSwagger.URL(cfg.Docs.URL), //The url pointing to API definition
	// ))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		utils.WriteJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})

	handlers.NewHttpHandler(cfg, logger, r)

	//r.Handle("/f", http.StripPrefix("/", http.FileServer(http.FS(contentStatic))))
	FileServer(r, "/static", http.Dir("./static"))

	r.NotFound(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := utils.WriteErrorResponse(w, apperrors.RouteNotFound(r.RequestURI))
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}))

	r.MethodNotAllowed(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := utils.WriteErrorResponse(w, apperrors.MethodNotAllowed(r.RequestURI, r.Method))
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}))

	return r
}

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	r.Get(path+"/*", func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	})
}
