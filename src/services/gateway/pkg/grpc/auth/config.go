package auth

type configParams struct {
	Port   string
	Domain string
}

func Config() configParams {
	return configParams{
		Port:   "9000",
		Domain: "localhost", //localhost
	}
}
