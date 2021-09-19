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

type Achievement struct {
	Id        string   `json:"id" db:"id"`
	Name      string   `json:"name" db:"name"`
	Subtitle  string   `json:"subtitle" db:"subtitle"`
	Desc      string   `json:"desc" db:"desc"`
	Icon      string   `json:"icon" db:"icon"`
	Link      string   `json:"link" db:"link"`
	Link_icon string   `json:"link_icon" db:"link_icon"`
	Github    string   `json:"github" db:"github"`
	Tech      []string `json:"tech" db:"tech"`
}

type Interest struct {
	Id   string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Icon string `json:"icon" db:"icon"`
	Desc string `json:"desc" db:"desc"`
}
