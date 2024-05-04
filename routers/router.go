package routers

func (r *Router) Routing() {

	// basic routes
	r.Route.HandleFunc("/api/index/home", r.Handler.Home)
	r.Route.HandleFunc("/api/index/about", r.Handler.About)
	r.Route.HandleFunc("/api/index/contact", r.Handler.Contact)

	// Authentication and Authorization routes ...
	r.Route.HandleFunc("/api/auth/register", r.Handler.Register)
	r.Route.HandleFunc("/api/auth/login", r.Handler.Login)
	r.Route.HandleFunc("/api/auth/logout", r.Handler.Logout)
	r.Route.HandleFunc("/api/account/profile/delete", r.Handler.DeleteProfile)
	r.Route.Get("/api/auth/verify_email/{username}/{token}", r.Handler.VerifyEmail)
	r.Route.Post("/api/auth/password/update", r.Handler.UpdatePassword)
	r.Route.Post("/api/auth/password/reset", r.Handler.ResetPassword)
	r.Route.Post("/api/auth/password/reset/for/{username}/{token}", r.Handler.NewPassword)

	// product routes
	r.Route.HandleFunc("/api/product/create", r.Handler.CreateProduct)
	r.Route.HandleFunc("/api/product/{id}/read", r.Handler.ReadProduct)
	r.Route.HandleFunc("/api/product/{id}/update", r.Handler.UpdateProduct)
	r.Route.HandleFunc("/api/product/{id}/delete", r.Handler.DeleteProduct)
	r.Route.HandleFunc("/api/product/image/create", r.Handler.ImageHandler)

	// profile routes
	r.Route.HandleFunc("/api/account/profile/{username}/read", r.Handler.ReadProfile)
	r.Route.HandleFunc("/api/account/profile/update", r.Handler.UpdateProfile)

}
