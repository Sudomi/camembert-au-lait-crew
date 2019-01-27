package soundcloud

import "fmt"

func (p *Playlist) IsMain() bool {
	switch p.PlaylistType {
	case "single", "album", "compilation", "ep":
		return true
	}
	return false
}

func (p *Playlist) Section() string {
	switch p.PlaylistType {
	case "album":
		return "albums"
	case "single", "ep":
		return "singles-eps"
	case "compilation":
		return "appears-on"
	}
	return "other"
}

func (p *Playlists) BySection(section string) []*Playlist {
	out := []*Playlist{}
	for _, playlist := range p.Playlists {
		if playlist.Section() == section {
			out = append(out, playlist)
		}
	}
	return out
}

func (p *Playlist) URL() string {
	return fmt.Sprintf("/album/%d", p.ID)
}

func (p *Playlist) IsExternal() bool { return true }

func (p *Playlist) ImageURL() string { return p.ArtworkUrl }