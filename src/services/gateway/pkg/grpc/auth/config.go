package auth

type configParams struct {
	Port   string
	Domain string
}

func Config() configParams {
	return configParams{
		Port:   "3000",
		Domain: "auth", //localhost
	}
}
