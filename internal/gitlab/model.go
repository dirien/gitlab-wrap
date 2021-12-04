package gitlab

import "time"

type Projects []struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Name        string    `json:"name"`
	CreatedAt   time.Time `json:"created_at"`
}

type User struct {
	ID              int       `json:"id"`
	Username        string    `json:"username"`
	Name            string    `json:"name"`
	State           string    `json:"state"`
	AvatarURL       string    `json:"avatar_url"`
	WebURL          string    `json:"web_url"`
	CreatedAt       time.Time `json:"created_at"`
	Bio             string    `json:"bio"`
	Location        string    `json:"location"`
	PublicEmail     string    `json:"public_email"`
	Skype           string    `json:"skype"`
	Linkedin        string    `json:"linkedin"`
	Twitter         string    `json:"twitter"`
	WebsiteURL      string    `json:"website_url"`
	Organization    string    `json:"organization"`
	JobTitle        string    `json:"job_title"`
	Pronouns        string    `json:"pronouns"`
	Bot             bool      `json:"bot"`
	WorkInformation string    `json:"work_information"`
	Followers       int       `json:"followers"`
	Following       int       `json:"following"`
	LocalTime       string    `json:"local_time"`
}

type Users []User

type WrapStats struct {
	User             *User
	ProjectSum       int
	IssuesSum        int
	MergeRequestsSum int
	StarredProjects  int
}
