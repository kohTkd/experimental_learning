package environment

type EnvName string

// Application Environment names
const (
	development EnvName = "development"
	test        EnvName = "test"
	production  EnvName = "production"
)

var envs = []EnvName{development, test, production}

func Development() EnvName {
	return development
}

func Test() EnvName {
	return test
}

func (e EnvName) IsValid() bool {
	return e == development || e == test || e == production
}

func (e *EnvName) String() string {
	return string(*e)
}

func Get(s string) *EnvName {
	for _, e := range envs {
		if s == e.String() {
			return &e
		}
	}

	return nil
}
