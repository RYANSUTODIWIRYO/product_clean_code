package container

import (
	"github.com/product/pkg/adapter/api/grpc/dto"
	repoDBProduct "github.com/product/pkg/adapter/repository/db"
	ucProduct "github.com/product/pkg/usecase/product"
	ucCreateProduct "github.com/product/pkg/usecase/create_product"
	"github.com/product/internal/config"

	"github.com/sarulabs/di"
)

// Container :
type Container struct {
	ctn di.Container
}

// NewContainer :
func NewContainer() *Container {
	builder, _ := di.NewBuilder()
	// conf := cfg.GetConfig()
	// driver, err := db.NewInstanceDb(conf.Database.Nds.Mysql)
	// if err != nil {
	// 	panic(fmt.Sprintf("db connection error. %v", err))
	// }

	//dbDDServiceInstanceDriver, err := db.NewInstanceDb(conf.Database.Ddservice.Mysql)
	//if err != nil {
	//	panic(fmt.Sprintf("db connection error. %v", err))
	//}
	//dbDDServiceDriver = dbDDServiceInstanceDriver
	//dbDriver = driver
	//txDriver = driver.(db.Transactioner)

	_ = builder.Add([]di.Def{
		{Name: "ProductSvc", Build: ProductUseCase},
		{Name: "CreateProductSvc", Build: CreateProductUseCase},
	}...)
	return &Container{
		ctn: builder.Build(),
	}
}

// Resolve :
func (c *Container) Resolve(name string) interface{} {
	return c.ctn.Get(name)
}

func ProductUseCase(_ di.Container) (interface{}, error) {
	database := config.GetConfig().Database.Product_DB.MySql
	repoDBProduct := repoDBProduct.NewProductRepo(database)
	reportProductDto := &dto.ProductBuilder{}
	return ucProduct.NewProductInteractor(repoDBProduct, reportProductDto), nil
}

func CreateProductUseCase(_ di.Container) (interface{}, error) {
	database := config.GetConfig().Database.Product_DB.MySql
	repoDBCreateProduct := repoDBProduct.NewCreateProductRepo(database)
	reportCreateProductDto := &dto.CreateProductBuilder{}
	return ucCreateProduct.NewCreateProductInteractor(repoDBCreateProduct, reportCreateProductDto), nil
}