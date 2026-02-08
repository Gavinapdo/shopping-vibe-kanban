package product

import (
	"sort"
	"sync"
)

type Repository interface {
	List() []Product
	GetByID(id int64) (Product, bool)
	Create(p Product) Product
	Update(id int64, p Product) (Product, bool)
	Delete(id int64) bool
}

type InMemoryRepository struct {
	mu     sync.RWMutex
	items  map[int64]Product
	nextID int64
}

func NewInMemoryRepository(seed []Product) *InMemoryRepository {
	repo := &InMemoryRepository{
		items:  make(map[int64]Product, len(seed)),
		nextID: 1,
	}

	var maxID int64
	for _, p := range seed {
		repo.items[p.ID] = p
		if p.ID > maxID {
			maxID = p.ID
		}
	}
	if maxID > 0 {
		repo.nextID = maxID + 1
	}

	return repo
}

func (r *InMemoryRepository) List() []Product {
	r.mu.RLock()
	defer r.mu.RUnlock()

	products := make([]Product, 0, len(r.items))
	for _, p := range r.items {
		products = append(products, p)
	}

	sort.Slice(products, func(i, j int) bool {
		return products[i].ID < products[j].ID
	})

	return products
}

func (r *InMemoryRepository) GetByID(id int64) (Product, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	p, ok := r.items[id]
	return p, ok
}

func (r *InMemoryRepository) Create(p Product) Product {
	r.mu.Lock()
	defer r.mu.Unlock()

	p.ID = r.nextID
	r.items[p.ID] = p
	r.nextID++

	return p
}

func (r *InMemoryRepository) Update(id int64, p Product) (Product, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.items[id]; !ok {
		return Product{}, false
	}

	p.ID = id
	r.items[id] = p
	return p, true
}

func (r *InMemoryRepository) Delete(id int64) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.items[id]; !ok {
		return false
	}

	delete(r.items, id)
	return true
}
