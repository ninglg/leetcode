func validIPAddress(IP string) string {
	ips := strings.Split(IP, ".")
	if len(ips) == 4 {
		for _, x := range ips {
			if (len(x) == 0) || (len(x) > 1 && x[0] == '0') || (len(x) > 1 && x[0] == '-') {
				return "Neither"
			}

			y, err := strconv.Atoi(x)
			if err != nil {
				return "Neither"
			}

			if y < 0 || y > 255 {
				return "Neither"
			}
		}
		return "IPv4"
	}

	ips = strings.Split(IP, ":")
	if len(ips) == 8 {
		for _, x := range ips {
			if len(x) == 0 || len(x) > 4 {
				return "Neither"
			}

			flag := true
			for i := 0; i < len(x); i++ {
				if (x[i] >= '0' && x[i] <= '9') || (x[i] >= 'a' && x[i] <= 'f') || (x[i] >= 'A' && x[i] <= 'F') {
					continue
				}
				flag = false
			}

			if !flag {
				return "Neither"
			}

		}
		return "IPv6"
	}

	return "Neither"
}

