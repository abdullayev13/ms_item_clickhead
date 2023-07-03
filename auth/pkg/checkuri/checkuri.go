package checkuri

import "regexp"

var userUris []*uri

type uri struct {
	regex   *regexp.Regexp
	methods map[string]struct{}
}

func CheckUriForUser(uri string, method string) bool {
	for _, u := range userUris {
		if u.regex.MatchString(uri) {
			_, ok := u.methods[method]
			return ok
		}
	}
	return false
}

func Add(expr string, methods ...string) error {
	regex, err := regexp.Compile(expr)
	if err != nil {
		return err
	}

	u := uri{regex: regex, methods: arrToSet(methods)}

	userUris = append(userUris, &u)
	return nil
}

func arrToSet(arr []string) map[string]struct{} {
	m := make(map[string]struct{}, len(arr))

	for _, s := range arr {
		m[s] = struct{}{}
	}

	return m
}

func init() {
	err := Add("^/api/auth/user-me", "GET", "POST", "PUT", "PATCH", "DELETE")
	if err != nil {
		panic(err)
	}
	err = Add("^/api/product/item/(list|[0-9])", "GET")
	if err != nil {
		panic(err)
	}

}
