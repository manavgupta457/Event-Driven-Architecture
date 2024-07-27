package main

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	// json_to_struct()
	json_to_struct2()

	// struct_to_json1()
	struct_to_json2()

}
