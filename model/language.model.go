package model

type Language struct {
	Id uint `gorm:"primary_key;auto_increment"`
	Iso string `gorm:"type:varchar(2);not null;unique"`
	Name string `gorm:"type:varchar(35);not null"`
	IsDefault bool `gorm:"not null;default:false"`
	Position int `gorm:"not null;default:0"`
}

type LanguageModel struct {
}

func (model LanguageModel) Get() ([]Language, Language) {
	var languages []Language
	var currentLanguage Language

	Db.Order("position").Find(&languages)

	for _, item := range languages {
        if item.IsDefault {
            currentLanguage = item
        }
    }

	return languages, currentLanguage
}

func (model LanguageModel) Check(iso string, languages []Language) (int) {
	for i, item := range languages {
        if item.Iso == iso {
            return i
        }
    }

	return -1
}

func (model LanguageModel) Init() {
	var (
		items = []Language{
			{Iso: "en", Name: "English", IsDefault: true, Position: 10},
			{Iso: "et", Name: "Eesti", Position: 30},
			{Iso: "ru", Name: "Русский", Position: 20},
		}
	)

	for index := range items {
		Db.Create(&items[index])
	}
}