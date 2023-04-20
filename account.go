package gocloak

import (
	"context"
)

// GetAccountGroups get all groups of the account
func (g *GoCloak) GetAccountGroups(ctx context.Context, token, realm string) ([]*Group, error) {
	const errMessage = "could not get groups"

	var result []*Group

	resp, err := g.GetRequestWithBearerAuth(ctx, token).
		SetResult(&result).
		Get(g.getRealmURL(realm, "account", "groups"))

	if err := checkForError(resp, err, errMessage); err != nil {
		return nil, err
	}

	return result, nil
}
