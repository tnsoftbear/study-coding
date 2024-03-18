// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"data-collector/internal/generated/restapi/operations"
	"data-collector/internal/generated/restapi/operations/article"
)

//go:generate swagger generate server --target ..\..\generated --name ArticleCollector --spec ..\..\..\api\spec.yml --principal interface{} --exclude-main

func configureFlags(api *operations.ArticleCollectorAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.ArticleCollectorAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()
	api.UrlformConsumer = runtime.DiscardConsumer
	api.XMLConsumer = runtime.XMLConsumer()

	api.JSONProducer = runtime.JSONProducer()
	api.XMLProducer = runtime.XMLProducer()

	// You may change here the memory limit for this multipart form parser. Below is the default (32 MB).
	// article.UpdateArticleWithFormMaxParseMemory = 32 << 20

	//api.ArticleAddArticleHandler = 

	if api.ArticleAddArticleHandler == nil {
		api.ArticleAddArticleHandler = article.AddArticleHandlerFunc(func(params article.AddArticleParams) middleware.Responder {
			return middleware.NotImplemented("operation article.AddArticle has not yet been implemented")
		})
	}
	if api.ArticleDeleteArticleHandler == nil {
		api.ArticleDeleteArticleHandler = article.DeleteArticleHandlerFunc(func(params article.DeleteArticleParams) middleware.Responder {
			return middleware.NotImplemented("operation article.DeleteArticle has not yet been implemented")
		})
	}
	if api.ArticleFindArticlesByIdsHandler == nil {
		api.ArticleFindArticlesByIdsHandler = article.FindArticlesByIdsHandlerFunc(func(params article.FindArticlesByIdsParams) middleware.Responder {
			return middleware.NotImplemented("operation article.FindArticlesByIds has not yet been implemented")
		})
	}
	if api.ArticleGetArticleByIDHandler == nil {
		api.ArticleGetArticleByIDHandler = article.GetArticleByIDHandlerFunc(func(params article.GetArticleByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation article.GetArticleByID has not yet been implemented")
		})
	}
	if api.ArticleUpdateArticleHandler == nil {
		api.ArticleUpdateArticleHandler = article.UpdateArticleHandlerFunc(func(params article.UpdateArticleParams) middleware.Responder {
			return middleware.NotImplemented("operation article.UpdateArticle has not yet been implemented")
		})
	}
	if api.ArticleUpdateArticleWithFormHandler == nil {
		api.ArticleUpdateArticleWithFormHandler = article.UpdateArticleWithFormHandlerFunc(func(params article.UpdateArticleWithFormParams) middleware.Responder {
			return middleware.NotImplemented("operation article.UpdateArticleWithForm has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
