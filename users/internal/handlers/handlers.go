package handlers

// - `POST /auth/register`: Create a new user account.
// - `POST /auth/login`: Authenticate user, return JWT.
// - `POST /auth/refresh`: Refresh JWT token.
// - `GET /auth/user`: Get authenticated user’s profile.
// - `POST /auth/oauth/:provider`: Initiate OAuth flow (e.g., Google, GitHub).
// - `POST /auth/oauth/:provider/callback`: Handle OAuth callback.
// - `DELETE /auth/logout`: Invalidate session.

/*
func RegisterAuthHandlers(app *gofr.App) error {
	pb.RegisterAuthServiceServerWithGofr(app, &pb.AuthServiceGoFrServer{})
	return nil
}
*/
