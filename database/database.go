package database

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/glebarez/sqlite"
	"github.com/tlksimba519/ptcg_trainerweb/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	db, err := gorm.Open(sqlite.Open("pokemon.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}
	if err := models.MigrateAll(db); err != nil {
		log.Fatal("Migration failed:", err)
	}
	DB = db

}

func Seed() {
	users := []models.User{
		{
			Account:  "admin",
			Password: "admin123",
			Name:     "Administrator",
			Email:    "admin@example.com",
		},
		{
			Account:  "user1",
			Password: "password1",
			Name:     "Alice",
			Email:    "alice@example.com",
		},
		{
			Account:  "user2",
			Password: "password2",
			Name:     "Bob",
			Email:    "bob@example.com",
		},
		{
			Account:  "user3",
			Password: "password3",
			Name:     "Charlie",
			Email:    "charlie@example.com",
		},
	}
	cards := []models.DeckCardItem{}

	deck := models.Deck{
		Title: "Red Aggro",
		Cards: cards,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 檢查是否已有相同標題的 deck
	filter := bson.M{"title": deck.Title}

	var existing models.Deck
	err := DeckCollection.FindOne(ctx, filter).Decode(&existing)

	if err == mongo.ErrNoDocuments {
		// 沒有找到，插入新的
		_, insertErr := DeckCollection.InsertOne(ctx, deck)
		if insertErr != nil {
			log.Printf("❌ Failed to insert deck %s: %v", deck.Title, insertErr)
		} else {
			log.Printf("✅ Inserted deck: %s", deck.Title)
		}
	} else if err != nil {
		log.Printf("❌ Error checking existing deck: %v", err)
	} else {
		log.Printf("⚠️ Deck %s already exists, skipped", deck.Title)
	}

	for _, user := range users {
		// 加密密碼
		hashedPwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Printf("❌ Failed to hash password for user %s: %v", user.Account, err)
			continue
		}
		user.Password = string(hashedPwd)

		// 檢查是否已存在避免重複插入
		var existing models.User
		if err := DB.Where("account = ?", user.Account).First(&existing).Error; err != nil {
			if err := DB.Create(&user).Error; err != nil {
				log.Printf("❌ Failed to insert user %s: %v", user.Account, err)
			} else {
				log.Printf("✅ Inserted user: %s", user.Account)
			}
		} else {
			log.Printf("⚠️ User %s already exists, skipped", user.Account)
		}
	}

	rand.Seed(time.Now().UnixNano())
	gofakeit.Seed(time.Now().UnixNano())

	cardTypes := []string{"Pokémon", "Trainer", "Energy"}
	seriesList := []string{"Base", "Fossil", "Jungle", "Neo", "EX", "XY", "Sun & Moon", "Sword & Shield"}
	rarities := []string{"Common", "Uncommon", "Rare", "Holo Rare", "Ultra Rare"}
	energyTypes := []string{"Fire", "Water", "Electric", "Grass", "Psychic", "Fighting", "Darkness", "Steel", "Fairy"}
	trainerSubtypes := []string{"Supporter", "Item", "Stadium"}
	pokemonNames := []string{"Pikachu", "Charizard", "Bulbasaur", "Squirtle", "Eevee", "Gengar", "Mewtwo", "Snorlax", "Dragonite", "Jigglypuff"}

	for i := 1; i <= 50; i++ {
		cardType := cardTypes[rand.Intn(len(cardTypes))]

		card := models.Card{
			CardID:      "CARD" + gofakeit.LetterN(6),
			Name:        gofakeit.Name(),
			CardType:    cardType,
			Series:      seriesList[rand.Intn(len(seriesList))],
			Rarity:      rarities[rand.Intn(len(rarities))],
			ImageURL:    gofakeit.ImageURL(200, 300),
			Description: gofakeit.Sentence(10),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		switch cardType {
		case "Pokémon":
			pokeName := pokemonNames[rand.Intn(len(pokemonNames))]
			card.Name = pokeName
			hp := gofakeit.Number(30, 180)
			typ := energyTypes[rand.Intn(len(energyTypes))]
			weakness := energyTypes[rand.Intn(len(energyTypes))]
			resist := energyTypes[rand.Intn(len(energyTypes))]
			retreat := gofakeit.Number(0, 4)
			atk1Name := gofakeit.HipsterWord()
			atk2Name := gofakeit.HipsterWord()
			card.HP = hp
			card.Type = typ
			card.Weakness = weakness
			card.Resisitance = resist
			card.RetreatCost = retreat
			card.Attack1Name = atk1Name
			card.Attack1Cost = gofakeit.Number(1, 3)
			card.Attack1Damage = gofakeit.Number(10, 60)
			txt1 := gofakeit.Sentence(6)
			card.Attack1Text = txt1
			card.Attack2Name = atk2Name
			card.Attack2Cost = gofakeit.Number(2, 5)
			card.Attack2Damage = gofakeit.Number(30, 100)
			txt2 := gofakeit.Sentence(8)
			card.Attack2Text = txt2

		case "Trainer":
			st := trainerSubtypes[rand.Intn(len(trainerSubtypes))]
			card.TrainerSubtype = st

		case "Energy":
			et := energyTypes[rand.Intn(len(energyTypes))]
			sp := gofakeit.Bool()
			card.EnergyType = et
			card.IsSpecial = sp
		}
		var count int64
		if err := DB.Model(&models.Card{}).Count(&count).Error; err != nil {
			log.Fatal("❌ 無法取得卡片數量：", err)
		}
		if count >= 100 {
			log.Println("📦 已有超過 100 張卡片，跳過種子資料建立")
			return
		}
		if err := DB.Create(&card).Error; err != nil {
			log.Printf("❌ Failed to insert card %s: %v", card.CardID, err)
		} else {
			log.Printf("✅ Inserted card: %s (%s)", card.CardID, card.CardType)
		}
	}
}
