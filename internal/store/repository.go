package store

type RequestRepository interface {
	Save(r *Request) error
	// FindById(id string) (*Request, error)
	List(offset, limit int) ([]*Request, error)
	// Load(id string) (*Request, error)
	// Update(r *Request) (*Request, error)
	// Delete(id string) error
}
