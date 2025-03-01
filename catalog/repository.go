package catalog

import (
	"context"
	"encoding/json"
	"errors"

	elastic "gopkg.in/olivere/elastic.v5"
)

var (
	ErrProductNotFound = errors.New("Product not found")
)

type Repository interface {
	Close()
	PutProduct(ctx context.Context, product Product) error
	GetProductById(ctx context.Context, id string) (*Product, error)
	ListProducts(ctx context.Context, skip uint64, take uint64) ([]Product, error)
	ListProductsWithIDs(ctx context.Context, ids []string) ([]Product, error)
	SearchProduct(ctx context.Context, query string, skip uint64, take uint64) ([]Product, error)
}

type elasticRepository struct {
	client *elastic.Client
}

type productDocument struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func NewElasticRepository(url string) (Repository, error) {
	client, err := elastic.NewClient(
		elastic.SetURL(url),
		elastic.SetSniff(false),
	)
	if err != nil {
		return nil, err
	}
	return &elasticRepository{client}, nil
}

func (r *elasticRepository) Close() {

}

func (r *elasticRepository) PutProduct(ctx context.Context, product Product) error {
	doc := productDocument{
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	}
	_, err := r.client.Index().
		Index("catalog").
		Type("product").
		Id(product.ID).
		BodyJson(doc).
		Do(ctx)
	return err
}

func (r *elasticRepository) GetProductById(ctx context.Context, id string) (*Product, error) {
	res, err := r.client.Get().
		Index("catalog").
		Type("product").
		Id(id).
		Do(ctx)
	if err != nil {
		return &Product{}, err
	}
	if !res.Found {
		return &Product{}, ErrProductNotFound
	}
	var doc productDocument
	if err := json.Unmarshal(*res.Source, &doc); err != nil {
		return &Product{}, err
	}
	return &Product{
		ID:          id,
		Name:        doc.Name,
		Description: doc.Description,
		Price:       doc.Price,
	}, nil
}

func (r *elasticRepository) ListProducts(ctx context.Context, skip uint64, take uint64) ([]Product, error) {
	res, err := r.client.Search().
		Index("catalog").
		Type("product").
		Query(elastic.NewMatchAllQuery()).
		From(int(skip)).
		Size(int(take)).
		Do(ctx)
	if err != nil {
		return nil, err
	}
	var products []Product
	for _, hit := range res.Hits.Hits {
		var doc productDocument
		if err := json.Unmarshal(*hit.Source, &doc); err != nil {
			return nil, err
		}
		products = append(products, Product{
			ID:          hit.Id,
			Name:        doc.Name,
			Description: doc.Description,
			Price:       doc.Price,
		})
	}
	return products, nil
}

func (r *elasticRepository) ListProductsWithIDs(ctx context.Context, ids []string) ([]Product, error) {
	items := []*elastic.MultiGetItem{}
	for _, id := range ids {
		items = append(items, elastic.NewMultiGetItem().Index("catalog").Type("product").Id(id))
	}
	res, err := r.client.MultiGet().Add(items...).Do(ctx)
	if err != nil {
		return nil, err
	}
	var products []Product
	for _, doc := range res.Docs {
		if doc.Found {
			var pd productDocument
			if err := json.Unmarshal(*doc.Source, &pd); err != nil {
				return nil, err
			}
			products = append(products, Product{
				ID:          doc.Id,
				Name:        pd.Name,
				Description: pd.Description,
				Price:       pd.Price,
			})
		}
	}
	return products, nil
}

func (r *elasticRepository) SearchProduct(ctx context.Context, query string, skip uint64, take uint64) ([]Product, error) {
	q := elastic.NewMultiMatchQuery(query, "name", "description")
	res, err := r.client.Search().
		Index("catalog").
		Type("product").
		Query(q).
		From(int(skip)).
		Size(int(take)).
		Do(ctx)
	if err != nil {
		return nil, err
	}
	var products []Product
	for _, hit := range res.Hits.Hits {
		var doc productDocument
		if err := json.Unmarshal(*hit.Source, &doc); err != nil {
			return nil, err
		}
		products = append(products, Product{
			ID:          hit.Id,
			Name:        doc.Name,
			Description: doc.Description,
			Price:       doc.Price,
		})
	}
	return products, nil
}
