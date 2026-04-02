package migrations

import (
    "encoding/json"
    "log"
    "os"
    "path/filepath"
    
    "gorm.io/gorm"
    
    "github.com/igor-izvekov/crm/pkg/models"
)

type Loader interface {
    GetModels() interface{}
    GetTableName() string
}

type GenericLoader[T any] struct {
    FileName  string
    TableName string
    DataKey   string
}

func (l *GenericLoader[T]) GetModels() interface{} {
    var wrapper struct {
        Data []T `json:"data"`
    }
    return &wrapper
}

func (l *GenericLoader[T]) GetTableName() string {
    return l.TableName
}

func loadFromJSON[T any](db *gorm.DB, filePath, dataKey string) ([]T, error) {
    data, err := os.ReadFile(filePath)
    if err != nil {
        log.Printf("Ошибка чтения файла %s: %v", filePath, err)
        return nil, err
    }
    
    var wrapper struct {
        Data []T `json:"data"`
    }
    
    if err := json.Unmarshal(data, &wrapper); err != nil {
        log.Printf("Ошибка парсинга JSON %s: %v", filePath, err)
        return nil, err
    }
    
    return wrapper.Data, nil
}

func saveBatch[T any](db *gorm.DB, items []T, name string) error {
    if len(items) == 0 {
        log.Printf("Нет данных для %s", name)
        return nil
    }
    
    for _, item := range items {
        if err := db.Create(&item).Error; err != nil {
            log.Printf("Ошибка вставки %s: %v", name, err)
            return err
        }
    }
    
    log.Printf("Загружено %s: %d записей", name, len(items))
    return nil
}

func loadAndSaveClients(db *gorm.DB, dataDir string) error {
    clients, err := loadFromJSON[models.Client](
        db,
        filepath.Join(dataDir, "clients.json"),
        "clients",
    )
    if err != nil {
        return err
    }
    return saveBatch(db, clients, "клиентов")
}

func loadAndSaveManagers(db *gorm.DB, dataDir string) error {
    managers, err := loadFromJSON[models.Manager](
        db,
        filepath.Join(dataDir, "managers.json"),
        "managers",
    )
    if err != nil {
        return err
    }
    return saveBatch(db, managers, "менеджеров")
}

func loadAndSaveObjects(db *gorm.DB, dataDir string) error {
    objects, err := loadFromJSON[models.RealEstateObject](
        db,
        filepath.Join(dataDir, "objects.json"),
        "objects",
    )
    if err != nil {
        return err
    }
    return saveBatch(db, objects, "объектов")
}

func loadAndSaveDeals(db *gorm.DB, dataDir string) error {
    deals, err := loadFromJSON[models.Deal](
        db,
        filepath.Join(dataDir, "deals.json"),
        "deals",
    )
    if err != nil {
        return err
    }
    return saveBatch(db, deals, "сделок")
}

func loadAndSaveCommissions(db *gorm.DB, dataDir string) error {
    commissions, err := loadFromJSON[models.Commission](
        db,
        filepath.Join(dataDir, "commissions.json"),
        "commissions",
    )
    if err != nil {
        return err
    }
    return saveBatch(db, commissions, "комиссий")
}

func loadAndSavePayments(db *gorm.DB, dataDir string) error {
    payments, err := loadFromJSON[models.Payment](
        db,
        filepath.Join(dataDir, "payments.json"),
        "payments",
    )
    if err != nil {
        return err
    }
    return saveBatch(db, payments, "оплат")
}

func loadAndSaveDocuments(db *gorm.DB, dataDir string) error {
    documents, err := loadFromJSON[models.Document](
        db,
        filepath.Join(dataDir, "documents.json"),
        "documents",
    )
    if err != nil {
        return err
    }
    return saveBatch(db, documents, "документов")
}

type SeedConfig struct {
    DataDir string
}

func NewSeedConfig() *SeedConfig {
    return &SeedConfig{
        DataDir: "migrations/data",
    }
}

func Seed(db *gorm.DB) error {
    return SeedWithConfig(db, NewSeedConfig())
}

func SeedWithConfig(db *gorm.DB, config *SeedConfig) error {
    log.Println("Заполняем тестовыми данными из JSON файлов")
    
    loaders := []func(*gorm.DB, string) error{
        loadAndSaveClients,
        loadAndSaveManagers,
        loadAndSaveObjects,
        loadAndSaveDeals,
        loadAndSaveCommissions,
        loadAndSavePayments,
        loadAndSaveDocuments,
    }
    
    for _, loader := range loaders {
        if err := loader(db, config.DataDir); err != nil {
            return err
        }
    }

    return nil
}
