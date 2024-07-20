package browser

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type raw_profile struct {
	Profile profile_path `json:"profile"`
}

type profile_path struct {
	InfoCache map[string]profile_info `json:"info_cache"`
	Profiles  []string                `json:"profiles_order"`
}

type profile_info struct {
	Name string `json:"name"`
}

type ChromiumLocalState struct {
	Profiles []profile
}

type profile struct {
	ID   string
	Name string
}

func GetLocalState(cpath string) (*ChromiumLocalState, error) {
	path := filepath.Join(cpath, "Local State")
	source, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	p := &raw_profile{}
	if err := json.Unmarshal(source, p); err != nil {
		return nil, err
	}

	localstate := &ChromiumLocalState{}

	for _, id := range p.Profile.Profiles {
		name := p.Profile.InfoCache[id].Name

		localstate.Profiles = append(localstate.Profiles, profile{
			ID:   id,
			Name: name,
		})
	}

	return localstate, nil
}

func (s *ChromiumLocalState) HasProfile(id string) bool {
	for _, profile := range s.Profiles {
		if profile.ID == id || profile.Name == id {
			return true
		}
	}
	return false
}

func (s *ChromiumLocalState) GetProfileKey(id string) string {
	for _, profile := range s.Profiles {
		if profile.ID == id || profile.Name == id {
			return profile.ID
		}
	}
	return ""
}
