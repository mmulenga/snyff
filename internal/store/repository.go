package store

type RequestRepository interface {
	Save(r *Request) error
	// Load(id string) (*Request, error)
	// Update(r *Request) (*Request, error)
	// Delete(id string) error
}
