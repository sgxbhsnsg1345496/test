package model

type Category struct {
	BaseModel
	Name             string `gorm:"type:varchar(20);not null"`
	ParentCategoryID int32
	ParentCategory   *Category
	Level            int32 `gorm:"type:int;not null;default:1"`
	IsTab            bool  `gorm:"default:false;not null"`
}

type Brands struct {
	BaseModel
	Name string `gorm:"type:varchar(20);not null"`
	Logo string `gorm:"type:varchar(200);default:'';not null"`
}

// 表之间建立多对多的关系
type GoodsCategoryBrand struct {
	BaseModel
	CategoryId int32 `gorm:"type:int;index:idx_category_brand,unique"`
	Category   Category

	BrandId int32 `gorm:"type:int;index:idx_category_brand,unique"`
	Brands  Brands
}

func (GoodsCategoryBrand) TableName() string {
	return "goodscategorybrand"
}

// 轮播图,一般是商业行为
type Banner struct {
	BaseModel
	Image string `gorm:"type:varchar(200);not null"`
	Url   string `gorm:"type:varchar(200);not null"`
	Index int32  `gorm:"type:int;default:1;not null"`
}

// 商品信息表
type Goods struct {
	BaseModel
	CategoryId int32 `gorm:"type:int;not null"`
	Category   Category

	BrandId int32 `gorm:"type:int;not null"`
	Brands  Brands

	OnSale   bool `gorm:"default:false;not null"`
	ShipFree bool `gorm:"default:false;not null"`
	IsNew    bool `gorm:"default:false;not null"`
	IsHot    bool `gorm:"default:false;not null"`

	Name             string   `gorm:"type:varchar(50);not null"`
	GoodsSn          string   `gorm:"type:varchar(50);not null"`
	ClickNum         int32    `gorm:"type:int;default:0;not null"`
	SoldNum          int32    `gorm:"type:int;default:0;not null"`
	FavNum           int32    `gorm:"type:int;default:0;not null"`
	MarkerPrice      float32  `gorm:"not null"`
	ShopPrice        float32  `gorm:"not null"`
	GoodsBrief       string   `gorm:"type:varchar(100);not null"`
	Images           GormList `gorm:"type:varchar(1000);not null"`
	DescImages       GormList `gorm:"type:varchar(1000);not null"`
	GoodsFrontImages string   `gorm:"type:varchar(200);not null"`
}
