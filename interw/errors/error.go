package main

type myError string

func (e myError) Error() string {
	return string(e)
}

func main() {
	println(handle())

}

func handle() error {
	return myError("error")
}
