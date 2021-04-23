package shortener

// RedirectRepository is responsible for the methods that are used to redirect a url.
type RedirectRepository interface {
	Find(code string) (*Redirect, error)
	Store(redirect *Redirect) error
}
