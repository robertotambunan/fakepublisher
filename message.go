package main

func createPreMessage(source string) (msgs []Message) {
	switch source {
	case "random":
		for i := 0; i < 400; i++ {
			msg := random(10000, 100000)
			msgs = append(msgs, Message{
				ProductID: msg,
			})
		}
	case "file":
		fileTemp := readFile("file.txt")
		for msg := range fileTemp {
			msgs = append(msgs, Message{
				ProductID: msg,
			})
		}
	}
	return
}
