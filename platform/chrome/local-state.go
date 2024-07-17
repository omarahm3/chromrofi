package chrome

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type raw_profile struct {
	Profile profile `json:"profile"`
}

type profile struct {
	InfoCache map[string]profile_info `json:"info_cache"`
	Profiles  []string                `json:"profiles_order"`
}

type profile_info struct {
	Name string `json:"name"`
}

type LocalState struct {
	Profiles []Profile
}

type Profile struct {
	ID   string
	Name string
}

func GetLocalState(cpath string) (*LocalState, error) {
	path := filepath.Join(cpath, "Local State")
	source, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	p := &raw_profile{}
	if err := json.Unmarshal(source, p); err != nil {
		return nil, err
	}

	localstate := &LocalState{}

	for _, profile := range p.Profile.Profiles {
		name := p.Profile.InfoCache[profile].Name

		localstate.Profiles = append(localstate.Profiles, Profile{
			ID:   profile,
			Name: name,
		})
	}

	return localstate, nil
}

func (s *LocalState) HasProfile(id string) bool {
	for _, profile := range s.Profiles {
		if profile.ID == id || profile.Name == id {
			return true
		}
	}
	return false
}

func (s *LocalState) GetProfileKey(id string) string {
	for _, profile := range s.Profiles {
		if profile.ID == id || profile.Name == id {
			return profile.ID
		}
	}
	return ""
}
