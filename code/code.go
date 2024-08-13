package code

const uniqueCodeCharacters = "1234567890qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"

func GenerateUniqueCodeByNumber(id int64) string {
	code := ""
	for {
		if id < 62 {
			break
		}
		code = string(uniqueCodeCharacters[id%62]) + code
		id = id / 62
	}
	if id >= 0 {
		code = string(uniqueCodeCharacters[id]) + code
	}

	return code
}
