func numUniqueEmails(emails []string) int {
	uniqueEmail := map[[2]string]struct{}{}

	for _, email := range emails {
		split := strings.Split(email, "@")
		var localName strings.Builder
		for i := 0; i < len(split[0]); i++ {
			if split[0][i] == '+' {
				break
			}
			if split[0][i] == '.' {
				continue
			}
			localName.WriteByte(split[0][i])
		}

		uniqueEmail[[2]string{localName.String(), split[1]}] = struct{}{}
	}

	return len(uniqueEmail)
}