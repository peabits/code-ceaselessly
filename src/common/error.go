package common

type Error string

func (err Error) Error() string {
	return string(err)
}

type IndexError string

func (err IndexError) Error() string {
	return string(err)
}
