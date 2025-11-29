package models


type User struct {
    ID       uint   `gorm:"primaryKey"`
    Username string `gorm:"unique;not null"`

    PublicKeys []PublicKey `gorm:"foreignKey:OwnerID;constraint:OnDelete:CASCADE;"`

}
