package soundcloud

import (
	"fmt"
	"net/url"
)

type endpoint struct {
	api     *Api
	base    string
	authReq bool
}

// Creates a new endpoint at the join of all dirs, where dir can be string or uint64 (ids).
// Last param can be bool, which would make an authentication required endpoint
func (api *Api) newEndpoint(dirs ...interface{}) endpoint {
	path := ""
	authReq := false
	for i, dir := range dirs {
		if sdir, ok := dir.(string); ok {
			path += "/" + sdir
			continue
		}
		if idir, ok := dir.(uint64); ok {
			path += fmt.Sprintf("/%d", idir)
			continue
		}
		if bdir, ok := dir.(bool); ok && i+1 == len(dirs) {
			authReq = bdir
			continue
		}
		panic("Only strings and uint64 accepted; last can be bool")
	}
	return endpoint{api, path, authReq}
}

type userEndpoint struct {
	endpoint
}

func (api *Api) newUserEndpoint(dirs ...interface{}) *userEndpoint {
	return &userEndpoint{api.newEndpoint(dirs...)}
}

func (u *userEndpoint) Get(params url.Values) (*User, error) {
	ret := new(User)
	err := u.api.get(u.base, params, ret, u.authReq)
	return ret, err
}

// func (u *userEndpoint) Put(params url.Values) (*User, error) {
//   ret := new(User)
//   err := u.api.put(u.base, params, ret)
//   return ret, err
// }

// func (u *userEndpoint) Delete(params url.Values) (*User, error) {
//   ret := new(User)
//   err := u.api.delete(u.base, params, ret)
//   return ret, err
// }

type trackEndpoint struct {
	endpoint
}

func (api *Api) newTrackEndpoint(dirs ...interface{}) *trackEndpoint {
	return &trackEndpoint{api.newEndpoint(dirs...)}
}

func (t *trackEndpoint) Get(params url.Values) (*Track, error) {
	ret := new(Track)
	err := t.api.get(t.base, params, ret)
	return ret, err
}

type commentEndpoint struct {
	endpoint
}

func (api *Api) newCommentEndpoint(dirs ...interface{}) *commentEndpoint {
	return &commentEndpoint{api.newEndpoint(dirs...)}
}

func (t *commentEndpoint) Get(params url.Values) (*Comment, error) {
	ret := new(Comment)
	err := t.api.get(t.base, params, ret)
	return ret, err
}

type playlistEndpoint struct {
	endpoint
}

func (api *Api) newPlaylistEndpoint(dirs ...interface{}) *playlistEndpoint {
	return &playlistEndpoint{api.newEndpoint(dirs...)}
}

func (t *playlistEndpoint) Get(params url.Values) (*Playlist, error) {
	ret := new(Playlist)
	err := t.api.get(t.base, params, ret)
	return ret, err
}

type groupEndpoint struct {
	endpoint
}

func (api *Api) newGroupEndpoint(dirs ...interface{}) *groupEndpoint {
	return &groupEndpoint{api.newEndpoint(dirs...)}
}

func (t *groupEndpoint) Get(params url.Values) (*Group, error) {
	ret := new(Group)
	err := t.api.get(t.base, params, ret)
	return ret, err
}

type appEndpoint struct {
	endpoint
}

func (api *Api) newAppEndpoint(dirs ...interface{}) *appEndpoint {
	return &appEndpoint{api.newEndpoint(dirs...)}
}

func (t *appEndpoint) Get(params url.Values) (*App, error) {
	ret := new(App)
	err := t.api.get(t.base, params, ret)
	return ret, err
}
