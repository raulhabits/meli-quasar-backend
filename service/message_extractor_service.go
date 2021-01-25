package service

import "errors"

func GetMessage(messages [][]string) (msg string, err error) {
	if len(messages) != 3 {
		return "", errors.New("The message could not be retrieved, there is not enough data")
	}

	previous := ""
	wasPreviousWordEmpty := false
	var message string

	var index int
	for index = 0;; {
		found:= false
		for _, messageList := range messages {
			if index >= len(messageList) {
				return message, nil
			}
			word := messageList[index]
			if word != "" && word != previous {
				if len(message) != 0 && !wasPreviousWordEmpty {
					message += " "
				}
				message = message + word
				found = true
				index++
				previous=word
				wasPreviousWordEmpty = false
				break
			}
		}
		if !found {
			message = message + " "
			index++
			wasPreviousWordEmpty = true
		}
	}
}