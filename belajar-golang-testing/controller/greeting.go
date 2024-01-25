package controller

func GetGreeting(name string, time int) (greeting string) {
	if time <= 0 {
		time = 1
	} else if time > 24 {
		time = 25
	}

	switch {
	case time >= 1 && time <= 10:
		greeting += "Good morning"
	case time >= 11 && time <= 14:
		greeting += "Good afternoon"
	case time >= 15 && time <= 24:
		greeting += "Good evening"
	}

	greeting += " " + name

	return
}
