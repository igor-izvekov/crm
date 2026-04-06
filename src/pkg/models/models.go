package models

import (
	"time"
)

type Client struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Name      string     `gorm:"not null;size:150" json:"name"`      // ФИО
	Email     string     `gorm:"size:100" json:"email"`              // Email
	Phone     string     `gorm:"not null;size:20" json:"phone"`      // Телефон
	Comment   string     `gorm:"size:1000" json:"comment"`           // Комментарий

	// Связи
	Deals []Deal `gorm:"foreignKey:SellerID" json:"deals,omitempty"`
}

type ObjectStatus string

const (
	ObjectStatusActive    ObjectStatus = "active"     // активен
	ObjectStatusSold      ObjectStatus = "sold"       // продан
	ObjectStatusReserved  ObjectStatus = "reserved"   // забронирован
	ObjectStatusCancelled ObjectStatus = "cancelled"  // снят с продажи
)

type RealEstateObject struct {
	ID          uint         `gorm:"primaryKey" json:"id"`
	Address     string       `gorm:"not null;size:300;index" json:"address"` // адрес
	ObjectType  string       `gorm:"size:50" json:"object_type"`             // квартира, дом, коммерция
	Area        float64      `json:"area"`                                   // площадь
	Cost        float64      `gorm:"not null" json:"cost"`                   // стоимость
	Description string       `gorm:"size:2000" json:"description"`           // описание
	Status      ObjectStatus `gorm:"default:active;size:20" json:"status"`   // статус объекта
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`

	// Связи
	Deals  []Deal           `gorm:"foreignKey:RealEstateObjectID" json:"deals,omitempty"`
}

type DealStage string

const (
	DealStageNew            DealStage = "new"              // новая сделка
	DealStageDocumentPrep   DealStage = "document_prep"    // подготовка документов
	DealStageWaitingPayment DealStage = "waiting_payment"  // ожидание оплаты
	DealStageCompleted      DealStage = "completed"        // завершена
	DealStageCancelled      DealStage = "cancelled"        // отменена
)

type Manager struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"not null;size:150" json:"name"`  // ФИО
	Email     string    `gorm:"size:100" json:"email"`          // Email
	Phone     string    `gorm:"size:20" json:"phone"`           // Телефон

	// Связи
	Deals []Deal `gorm:"foreignKey:ManagerID" json:"deals,omitempty"`
}

type Deal struct {
	ID                 uint       `gorm:"primaryKey" json:"id"`
	SellerID           uint       `gorm:"not null;index" json:"seller_id"`               // продавец
	BuyerID            uint       `gorm:"not null;index" json:"buyer_id"`                // покупатель
	RealEstateObjectID uint       `gorm:"not null;index" json:"real_estate_object_id"`   // объект
	Amount             float64    `gorm:"not null" json:"amount"`                        // сумма сделки
	Stage              DealStage  `gorm:"default:new;size:30" json:"stage"`              // этап сделки
	ManagerID          uint       `gorm:"not null;index" json:"manager_id"`              // ответственный менеджер
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at"`

	// Связи
	Seller           Client           `gorm:"foreignKey:SellerID" json:"seller,omitempty"`
	Buyer            Client           `gorm:"foreignKey:BuyerID" json:"buyer,omitempty"`
	Manager          Manager          `gorm:"foreignKey:ManagerID" json:"manager,omitempty"`
	RealEstateObject RealEstateObject `gorm:"foreignKey:RealEstateObjectID" json:"real_estate_object,omitempty"`
	Documents        []Document       `gorm:"foreignKey:DealID" json:"documents,omitempty"`
	Payments         []Payment        `gorm:"foreignKey:DealID" json:"payments,omitempty"`
	Commission       *Commission      `gorm:"foreignKey:DealID" json:"commission,omitempty"`
}

type Document struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	DealID     uint      `gorm:"not null;index" json:"deal_id"`       // привязка к сделке
	Name       string    `gorm:"not null;size:200" json:"name"`       // название файла
	FilePath   string    `gorm:"not null;size:500" json:"file_path"`  // путь к файлу
	FileSize   int64     `json:"file_size"`                           // размер в байтах
	MimeType   string    `gorm:"size:100" json:"mime_type"`           // тип файла
	UploadedBy uint      `json:"uploaded_by"`                         // ID менеджера
	UploadedAt time.Time `json:"uploaded_at"`

	// Связи
	Deal    Deal    `gorm:"foreignKey:DealID" json:"deal,omitempty"`
	Manager Manager `gorm:"foreignKey:UploadedBy" json:"manager,omitempty"`
}

type PaymentType string

const (
	PaymentTypeCash     PaymentType = "cash"      // наличные
	PaymentTypeCard     PaymentType = "card"      // карта
	PaymentTypeTransfer PaymentType = "transfer"  // перевод
	PaymentTypeMortgage PaymentType = "mortgage"  // ипотека
)

type Payment struct {
	ID        uint        `gorm:"primaryKey" json:"id"`
	DealID    uint        `gorm:"not null;index" json:"deal_id"` // привязка к сделке
	Date      time.Time   `gorm:"not null" json:"date"`          // дата оплаты
	Amount    float64     `gorm:"not null" json:"amount"`        // сумма
	Type      PaymentType `gorm:"size:30" json:"type"`           // тип оплаты
	CreatedAt time.Time   `json:"created_at"`

	// Связь
	Deal Deal `gorm:"foreignKey:DealID" json:"deal,omitempty"`
}
