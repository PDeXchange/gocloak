package gocloak

import "context"

// GetUser returns a user with the given userid
func (g *GoCloak) GetUser(ctx context.Context, token, realm string, userID string) (*User, error) {
	const errMessage = "could not get user"

	var result *User

	resp, err := g.GetRequestWithBearerAuth(ctx, token).
		SetResult(&result).
		Get(g.getAdminRealmURL(realm, "users", userID))

	if err := checkForError(resp, err, errMessage); err != nil {
		return nil, err
	}

	return result, nil
}
