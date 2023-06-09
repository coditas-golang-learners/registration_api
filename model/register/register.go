package register

type Register struct {
	ID        int    `json:"id"`
	Username  string `json:"Username"`
	Firstname string `json:"Firstname"`
	Lastname  string `json:"Lastname"`
	Email     string `json:"Email"`
	Pan       string `json:"Pan"`
	Adhar     string `json:"Adhar"`
	Mobile    string `json:"Mobile"`
	Password  string `json:"Password"`
}
