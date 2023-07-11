package user

func RemoveDuplicates(inputUsers []*User) []*User {
	keys := make(map[string]bool)
	users := []*User{}

	for _, u := range inputUsers {
		if keys[u.GetSecurityUUID()] {
			continue
		}

		keys[u.GetSecurityUUID()] = true
		users = append(users, u)
	}

	return users
}
