package gravatar_recon

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type GravatarProfile struct {
	Hash          string `json:"hash"`
	ProfileURL    string `json:"profileUrl"`
	PreferredUser string `json:"preferredUsername"`
	ThumbnailURL  string `json:"thumbnailUrl"`

	Photos []struct {
		Value string `json:"value"`
		Type  string `json:"type"`
	} `json:"photos"`

	DisplayName     string `json:"displayName"`
	Pronunciation   string `json:"pronunciation"`
	Pronouns        string `json:"pronouns"`
	AboutMe         string `json:"aboutMe"`
	CurrentLocation string `json:"currentLocation"`
	JobTitle        string `json:"job_title"`
	Company         string `json:"company"`

	PhoneNumbers []struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"phoneNumbers"`

	ContactInfo []struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"contactInfo"`

	Emails []struct {
		Primary string `json:"primary"`
		Value   string `json:"value"`
	} `json:"emails"`

	Accounts []struct {
		Domain    string `json:"domain"`
		Display   string `json:"display"`
		URL       string `json:"url"`
		IconURL   string `json:"iconUrl"`
		IsHidden  bool   `json:"is_hidden"`
		Username  string `json:"username"`
		Verified  bool   `json:"verified"`
		Name      string `json:"name"`
		Shortname string `json:"shortname"`
	} `json:"accounts"`

	ProfileBackground struct {
		Color        string  `json:"color"`
		Opacity      float64 `json:"opacity"`
		Size         int     `json:"size"`
		PrimaryColor string  `json:"primary_color"`
		SurfaceColor string  `json:"surface_color"`
	} `json:"profileBackground"`
}

type gravatarResponse struct {
	Entry []GravatarProfile `json:"entry"`
}

func md5Hash(email string) string {
	email = strings.ToLower(strings.TrimSpace(email))
	hash := md5.Sum([]byte(email))
	return hex.EncodeToString(hash[:])
}

func GetGravatarProfiles(email string) (*[]GravatarProfile, error) {
	hash := md5Hash(email)
	url := fmt.Sprintf("https://www.gravatar.com/%s.json", hash)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Go-OSINT-Client/1.0")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return &[]GravatarProfile{}, nil
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("no gravatar profile found (status %d)", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var profiles gravatarResponse
	if err := json.Unmarshal(body, &profiles); err != nil {
		return nil, err
	}

	return &profiles.Entry, nil
}
