package main




func main() {
	hashtable := Init()
	list := []string{
		"Omid",
		"Ali",
		"Milad",
		"Kianoosh",
		"Mahdan",
		"Mehdi",
		"Alireza",
		"Amin",
		"Delaram",
	}
	for _, v := range list {
		hashtable.Insert(v)
	}
}


