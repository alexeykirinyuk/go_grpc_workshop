package product_service

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func setup(t *testing.T) (*Service, *MockIRepository, *MockICategoryClient) {
	ctrl := gomock.NewController(t)
	repo := NewMockIRepository(ctrl)
	client := NewMockICategoryClient(ctrl)

	service := &Service{
		repo:   repo,
		client: client,
	}

	return service, repo, client
}

func TestCreateProduct_Success_ReturnName(t *testing.T) {
	serv, repo, client := setup(t)

	client.EXPECT().
		IsCategoryExists(context.Background(), int64(312)).
		Return(true, nil)

	repo.EXPECT().
		SaveProduct(gomock.Any(), gomock.Any()).
		DoAndReturn(func(ctx context.Context, product *Product) error {
			product.ID = 124
			return nil
		})

	prod, err := serv.CreateProduct(context.Background(), "test-product", 312)
	require.Nil(t, err)
	require.Equal(t, "test-product", prod.Name)
}

func TestCreateProduct_Success_ReturnCategoryID(t *testing.T) {
	serv, repo, client := setup(t)

	client.EXPECT().
		IsCategoryExists(context.Background(), int64(312)).
		Return(true, nil)

	repo.EXPECT().
		SaveProduct(gomock.Any(), gomock.Any()).
		DoAndReturn(func(ctx context.Context, product *Product) error {
			product.ID = 124
			return nil
		})

	prod, err := serv.CreateProduct(context.Background(), "test-product", 312)
	require.Nil(t, err)
	require.Equal(t, int64(312), prod.CategoryId)
}

func TestCreateProduct_Success_Return_ID(t *testing.T) {
	serv, repo, client := setup(t)

	client.EXPECT().
		IsCategoryExists(context.Background(), int64(312)).
		Return(true, nil)

	repo.EXPECT().
		SaveProduct(gomock.Any(), gomock.Any()).
		DoAndReturn(func(ctx context.Context, product *Product) error {
			product.ID = 124
			return nil
		})

	prod, err := serv.CreateProduct(context.Background(), "test-product", 312)
	require.Nil(t, err)
	require.Equal(t, int64(124), prod.ID)
}

func TestCreateProduct_CategoryDoesNotExists(t *testing.T) {
	serv, _, client := setup(t)

	client.EXPECT().
		IsCategoryExists(context.Background(), int64(312)).
		Return(false, nil)

	_, err := serv.CreateProduct(context.Background(), "test-product", 312)
	require.NotNil(t, err)
}
