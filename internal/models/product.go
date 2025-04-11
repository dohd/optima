package models

import (
	"time"
)

type Product struct {
    ID                  int64     `gorm:"primaryKey;autoIncrement;not null"`
    Name                string    `gorm:"size:100;not null;"`
    Type                string    `gorm:"size:50;not null"` // Inventory, NonInventory, Service, OtherCharge
    Description         *string    `gorm:"size:255"`
    UnitPrice           float64    `gorm:"type:decimal(18,4);default:0;not null"`
    UnitCost            float64    `gorm:"type:decimal(18,4);default:0;not null"`
    SKU                 string    `gorm:"size:100;not null"`
    TrackQtyOnHand      bool      `gorm:"default:false;not null"`
    QtyOnHand           float64   `gorm:"type:decimal(10,2);default:0;not null"`
    ReorderPoint        int        `gorm:"default:0;not null"`
    InvStartDate        *time.Time `gorm:"type:date"`
    IncomeAccountID     *int64
    ExpenseAccountID    *int64
    AssetAccountID      *int64
    VendorID            *int64
    PurchaseDesc        *string    `gorm:"size:255"`
    PurchaseTaxIncluded bool       `gorm:"default:false;not null"`
    Taxable             bool      `gorm:"default:true;not null"`
    TaxRate             float64   `gorm:"type:decimal(10,2);default:0;not null"`
    Active              bool      `gorm:"default:true;not null"`
    SyncToken           *int64
    CreatedAt           time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
    UpdatedAt           time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
    DeletedAt           *time.Time
}
