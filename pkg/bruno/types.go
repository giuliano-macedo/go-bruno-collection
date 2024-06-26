package bruno

type Collection struct {
	Name                 string        `json:"name"`
	Version              string        `json:"version"`
	Items                []Item        `json:"items"`
	ActiveEnvironmentUid string        `json:"activeEnvironmentUid"`
	Environments         []Environment `json:"environments"`
	// Curently not supported in bruno while importing from json collections
	Docs string `json:"-"`
	Uid  string `json:"uid,omitempty"`
}

type Item struct {
	Type    string   `json:"type"`
	Name    string   `json:"name"`
	Seq     int      `json:"seq,omitempty"`
	Items   []Item   `json:"items,omitempty"`
	Request *Request `json:"request,omitempty"`
}

func (item *Item) IsFolder() bool {
	return len(item.Items) > 0
}

type Request struct {
	URL        string   `json:"url"`
	Method     string   `json:"method"`
	Headers    []Header `json:"headers"`
	Body       Body     `json:"body"`
	Auth       Auth     `json:"auth"`
	Script     Script   `json:"script"`
	Vars       Vars     `json:"vars"`
	Assertions []string `json:"assertions"`
	Tests      string   `json:"tests"`
	Docs       string   `json:"docs,omitempty"`
	Query      []string `json:"query,omitempty"`
	Params     []Param  `json:"params"`
}

type Param struct {
	Uid         string `json:"uid,omitempty"`
	Name        string `json:"name,omitempty"`
	Value       string `json:"value,omitempty"`
	Description string `json:"description,omitempty"`
	Enabled     bool   `json:"enabled,omitempty"`
}

type Header struct {
	Name    string `json:"name"`
	Value   string `json:"value"`
	Enabled bool   `json:"enabled"`
}

type Body struct {
	Mode           string   `json:"mode"`
	Json           string   `json:"json,omitempty"`
	FormUrlEncoded []string `json:"formUrlEncoded"`
	MultipartForm  []string `json:"multipartForm"`
}

type Auth struct {
	Mode   string `json:"mode"`
	Basic  Basic  `json:"basic"`
	Bearer Bearer `json:"bearer"`
}

type Basic struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Bearer struct {
	Token string `json:"token"`
}

type Script struct {
	Req string `json:"req,omitempty"`
	Res string `json:"res,omitempty"`
}

type Vars struct {
	Req []Variable `json:"req,omitempty"`
	Res []Variable `json:"res,omitempty"`
}

type Variable struct {
	Name    string `json:"name"`
	Value   string `json:"value"`
	Enabled bool   `json:"enabled"`
	Local   bool   `json:"local"`
}

type Environment struct {
	Uid       string                `json:"uid,omitempty"`
	Name      string                `json:"name"`
	Variables []EnvironmentVariable `json:"variables"`
}

type EnvironmentVariable struct {
	Name    string `json:"name"`
	Value   string `json:"value"`
	Enabled bool   `json:"enabled"`
	Secret  bool   `json:"secret"`
	Type    string `json:"type"`
}
