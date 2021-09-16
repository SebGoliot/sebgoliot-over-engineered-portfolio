package main

type Skill struct {
	Id    string `json:"id" db:"id"`
	Title string `json:"title" db:"title"`
	Icon  string `json:"icon" db:"icon"`
	Desc  string `json:"desc" db:"desc"`
	Like  bool   `json:"like" db:"like"`
}

type Author struct {
	Id       string `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Title    string `json:"title" db:"title"`
	Subtitle string `json:"subtitle" db:"subtitle"`
	About    string `json:"about" db:"about"`
	Github   string `json:"github" db:"github"`
	Linkedin string `json:"linkedin" db:"linkedin"`
}
