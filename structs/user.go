package structs

type ClothingEnum string

const (
	ClothingTypeMale   ClothingEnum = "MALE"
	ClothingTypeFemale ClothingEnum = "FEMALE"
)

type Users struct {
	ID           string
	ClothingType ClothingEnum
}
