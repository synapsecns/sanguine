package db

import (
	"context"
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/synapsecns/sanguine/rfq/quoting-api/db/models"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// setupDatabase is a helper function to create an in-memory SQLite database for testing.
func setupDatabase(t *testing.T) *Database {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}
	db.AutoMigrate(&models.Quote{})
	return &Database{DB: db}
}

// TestInsertQuote tests the InsertQuote function.
func TestInsertQuote(t *testing.T) {
	// Arrange
	ctx := context.Background()
	database := setupDatabase(t)
	err := database.AutoMigrate(&models.Quote{})
	assert.NoError(t, err)

	originAmount, _ := decimal.NewFromString("200")
	destAmount, _ := decimal.NewFromString("100")

	// Create a dummy quote
	quote := models.Quote{
		Relayer:       "0xA",
		OriginChainId: 1,
		DestChainId:   2,
		OriginToken:   "0x1",
		DestToken:     "0x2",
		OriginAmount:  originAmount,
		DestAmount:    destAmount,
	}

	// Act
	id, err := database.InsertQuote(ctx, &quote)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, uint(1), id)

	var q models.Quote
	result := database.Last(&q)
	assert.NoError(t, result.Error)

	quote.ID = q.ID
	quote.Price, _ = decimal.NewFromString("0.5")
	quote.CreatedAt = q.CreatedAt
	quote.UpdatedAt = q.UpdatedAt
	quote.DeletedAt = q.DeletedAt
	assert.Equal(t, quote, q)
}

// TestUpdateQuote tests the UpdateQuote function.
func TestUpdateQuote(t *testing.T) {
	// Arrange
	ctx := context.Background()
	database := setupDatabase(t)
	err := database.AutoMigrate(&models.Quote{})
	assert.NoError(t, err)

	originAmount, _ := decimal.NewFromString("200")
	destAmount, _ := decimal.NewFromString("100")

	// Create a dummy quote
	quote := models.Quote{
		Relayer:       "0xA",
		OriginChainId: 1,
		DestChainId:   2,
		OriginToken:   "0x1",
		DestToken:     "0x2",
		OriginAmount:  originAmount,
		DestAmount:    destAmount,
	}
	_, err = database.InsertQuote(ctx, &quote)
	assert.NoError(t, err)

	// retrieve the stored quote
	result := database.Last(&quote)
	assert.NoError(t, result.Error)

	// update origin amount to decrease price
	originAmount, _ = decimal.NewFromString("400")
	quote.OriginAmount = originAmount

	// act
	err = database.UpdateQuote(ctx, &quote)

	// assert
	assert.NoError(t, err)

	var q models.Quote
	result = database.Where(&models.Quote{ID: quote.ID}).Find(&q)
	assert.NoError(t, result.Error)

	quote.Price, _ = decimal.NewFromString("0.25")
	quote.CreatedAt = q.CreatedAt
	quote.UpdatedAt = q.UpdatedAt
	quote.DeletedAt = q.DeletedAt
	assert.Equal(t, quote, q)
}

// TestGetQuote tests the GetQuote function.
func TestGetQuote(t *testing.T) {
	// Arrange
	ctx := context.Background()
	database := setupDatabase(t)
	err := database.AutoMigrate(&models.Quote{})
	assert.NoError(t, err)

	originAmount, _ := decimal.NewFromString("200")
	destAmount, _ := decimal.NewFromString("100")

	// Create a dummy quote
	quote := models.Quote{
		Relayer:       "0xA",
		OriginChainId: 1,
		DestChainId:   2,
		OriginToken:   "0x1",
		DestToken:     "0x2",
		OriginAmount:  originAmount,
		DestAmount:    destAmount,
	}
	_, err = database.InsertQuote(ctx, &quote)
	assert.NoError(t, err)

	// refetch
	result := database.Last(&quote)
	assert.NoError(t, result.Error)

	// act
	q, err := database.GetQuote(ctx, quote.ID)
	quote.CreatedAt = q.CreatedAt
	quote.UpdatedAt = q.UpdatedAt
	quote.DeletedAt = q.DeletedAt

	assert.NoError(t, err)
	assert.Equal(t, quote, q)
}

// TestGetQuotes tests the GetQuotes function.
func TestGetQuotes(t *testing.T) {
	// Arrange
	ctx := context.Background()
	database := setupDatabase(t)
	err := database.AutoMigrate(&models.Quote{})
	assert.NoError(t, err)

	originAmount, _ := decimal.NewFromString("200")
	destAmount, _ := decimal.NewFromString("100")

	// Create dummy quotes
	quotes := make([]models.Quote, 7)
	quotes[0] = models.Quote{
		Relayer:       "0xA",
		OriginChainId: 10,
		DestChainId:   20,
		OriginToken:   "0x1",
		DestToken:     "0x2",
		OriginAmount:  originAmount,
		DestAmount:    destAmount,
	}
	quotes[1] = models.Quote{
		Relayer:       "0xA",
		OriginChainId: 10,
		DestChainId:   20,
		OriginToken:   "0x1",
		DestToken:     "0x2",
		OriginAmount:  destAmount,
		DestAmount:    destAmount,
	}
	quotes[2] = models.Quote{
		Relayer:       "0xA",
		OriginChainId: 10,
		DestChainId:   30,
		OriginToken:   "0x1",
		DestToken:     "0x2",
		OriginAmount:  originAmount,
		DestAmount:    destAmount,
	}
	quotes[3] = models.Quote{
		Relayer:       "0xA",
		OriginChainId: 30,
		DestChainId:   20,
		OriginToken:   "0x1",
		DestToken:     "0x2",
		OriginAmount:  originAmount,
		DestAmount:    destAmount,
	}
	quotes[4] = models.Quote{
		Relayer:       "0xA",
		OriginChainId: 10,
		DestChainId:   20,
		OriginToken:   "0x3",
		DestToken:     "0x2",
		OriginAmount:  originAmount,
		DestAmount:    destAmount,
	}
	quotes[5] = models.Quote{
		Relayer:       "0xA",
		OriginChainId: 10,
		DestChainId:   20,
		OriginToken:   "0x1",
		DestToken:     "0x3",
		OriginAmount:  originAmount,
		DestAmount:    destAmount,
	}
	quotes[6] = models.Quote{
		Relayer:       "0xB",
		OriginChainId: 10,
		DestChainId:   20,
		OriginToken:   "0x1",
		DestToken:     "0x2",
		OriginAmount:  destAmount.Mul(decimal.NewFromFloat(0.9)),
		DestAmount:    destAmount,
	}
	for _, quote := range quotes {
		_, err = database.InsertQuote(ctx, &quote)
		assert.NoError(t, err)
	}

	// build the request
	req := models.Request{
		OriginChainId: 10,
		DestChainId:   20,
		OriginToken:   "0x1",
		DestToken:     "0x2",
		OriginAmount:  originAmount.Mul(decimal.NewFromFloat(0.5)),
	}
	qs, err := database.GetQuotes(ctx, &req)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(qs)) // filter should only produce two

	// should be first two quotes in reverse of insert order given price descending
	expectedQuotes := make([]models.Quote, 2)
	expectedQuotes[0] = models.Quote{
		ID:            qs[0].ID,
		Relayer:       "0xA",
		OriginChainId: 10,
		DestChainId:   20,
		OriginToken:   "0x1",
		DestToken:     "0x2",
		OriginAmount:  destAmount,
		DestAmount:    destAmount,
		Price:         decimal.NewFromFloat(1.0),
		CreatedAt:     qs[0].CreatedAt,
		UpdatedAt:     qs[0].UpdatedAt,
		DeletedAt:     qs[0].DeletedAt,
	}
	expectedQuotes[1] = models.Quote{
		ID:            qs[1].ID,
		Relayer:       "0xA",
		OriginChainId: 10,
		DestChainId:   20,
		OriginToken:   "0x1",
		DestToken:     "0x2",
		OriginAmount:  originAmount,
		DestAmount:    destAmount,
		Price:         decimal.NewFromFloat(0.5),
		CreatedAt:     qs[1].CreatedAt,
		UpdatedAt:     qs[1].UpdatedAt,
		DeletedAt:     qs[1].DeletedAt,
	}
	assert.Equal(t, expectedQuotes, qs)
}

// TestGetQuotesWhenUpdatedAtLast test the GetQuotes function with UpdatedAtLast time filter.
func TestGetQuotesWhenUpdatedAtLast(t *testing.T) {
	// Arrange
	ctx := context.Background()
	database := setupDatabase(t)
	err := database.AutoMigrate(&models.Quote{})
	assert.NoError(t, err)

	originAmount, _ := decimal.NewFromString("200")
	destAmount, _ := decimal.NewFromString("100")

	// Create dummy quotes
	quotes := make([]models.Quote, 7)
	quotes[0] = models.Quote{
		Relayer:       "0xA",
		OriginChainId: 10,
		DestChainId:   20,
		OriginToken:   "0x1",
		DestToken:     "0x2",
		OriginAmount:  originAmount,
		DestAmount:    destAmount,
	}
	quotes[1] = models.Quote{
		Relayer:       "0xA",
		OriginChainId: 10,
		DestChainId:   20,
		OriginToken:   "0x1",
		DestToken:     "0x2",
		OriginAmount:  destAmount,
		DestAmount:    destAmount,
	}
	quotes[2] = models.Quote{
		Relayer:       "0xA",
		OriginChainId: 10,
		DestChainId:   30,
		OriginToken:   "0x1",
		DestToken:     "0x2",
		OriginAmount:  originAmount,
		DestAmount:    destAmount,
	}
	quotes[3] = models.Quote{
		Relayer:       "0xA",
		OriginChainId: 30,
		DestChainId:   20,
		OriginToken:   "0x1",
		DestToken:     "0x2",
		OriginAmount:  originAmount,
		DestAmount:    destAmount,
	}
	quotes[4] = models.Quote{
		Relayer:       "0xA",
		OriginChainId: 10,
		DestChainId:   20,
		OriginToken:   "0x3",
		DestToken:     "0x2",
		OriginAmount:  originAmount,
		DestAmount:    destAmount,
	}
	quotes[5] = models.Quote{
		Relayer:       "0xA",
		OriginChainId: 10,
		DestChainId:   20,
		OriginToken:   "0x1",
		DestToken:     "0x3",
		OriginAmount:  originAmount,
		DestAmount:    destAmount,
	}
	quotes[6] = models.Quote{
		Relayer:       "0xB",
		OriginChainId: 10,
		DestChainId:   20,
		OriginToken:   "0x1",
		DestToken:     "0x2",
		OriginAmount:  destAmount.Mul(decimal.NewFromFloat(0.9)),
		DestAmount:    destAmount,
	}
	for _, quote := range quotes {
		_, err = database.InsertQuote(ctx, &quote)
		assert.NoError(t, err)
	}

	// build the request
	req := models.Request{
		OriginChainId: 10,
		DestChainId:   20,
		OriginToken:   "0x1",
		DestToken:     "0x2",
		OriginAmount:  decimal.NewFromFloat(100),
		UpdatedAtLast: time.Now().Add(1 * time.Second),
	}
	qs, err := database.GetQuotes(ctx, &req)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(qs)) // filter should produce zero

	// should be first two quotes in reverse of insert order given price descending
	expectedQuotes := make([]models.Quote, 0)
	assert.Equal(t, expectedQuotes, qs)
}

// TestDeleteQuote tests the DeleteQuote function.
func TestDeleteQuote(t *testing.T) {
	// Arrange
	ctx := context.Background()
	database := setupDatabase(t)
	err := database.AutoMigrate(&models.Quote{})
	assert.NoError(t, err)

	originAmount, _ := decimal.NewFromString("200")
	destAmount, _ := decimal.NewFromString("100")

	// Create a dummy quote
	quote := models.Quote{
		Relayer:       "0xA",
		OriginChainId: 1,
		DestChainId:   2,
		OriginToken:   "0x1",
		DestToken:     "0x2",
		OriginAmount:  originAmount,
		DestAmount:    destAmount,
	}
	_, err = database.InsertQuote(ctx, &quote)
	assert.NoError(t, err)

	// refetch
	result := database.Last(&quote)
	assert.NoError(t, result.Error)

	// act
	err = database.DeleteQuote(ctx, quote.ID)

	// assert
	assert.NoError(t, result.Error)

	_, err = database.GetQuote(ctx, quote.ID)
	assert.Error(t, err)
}

func TestBeforeCreateQuoteWhenIDNotZero(t *testing.T) {
	// Arrange
	ctx := context.Background()
	database := setupDatabase(t)
	err := database.AutoMigrate(&models.Quote{})
	assert.NoError(t, err)

	originAmount, _ := decimal.NewFromString("200")
	destAmount, _ := decimal.NewFromString("100")

	// Create a dummy quote
	quote := models.Quote{
		ID:            3,
		Relayer:       "0xA",
		OriginChainId: 1,
		DestChainId:   2,
		OriginToken:   "0x1",
		DestToken:     "0x2",
		OriginAmount:  originAmount,
		DestAmount:    destAmount,
	}

	// Act
	id, err := database.InsertQuote(ctx, &quote)

	// Assert
	assert.Error(t, err, "Invalid quote: created q.ID != 0")
	assert.Equal(t, id, uint(0))
}

func TestBeforeCreateQuoteWhenSameChainId(t *testing.T) {
	// Arrange
	ctx := context.Background()
	database := setupDatabase(t)
	err := database.AutoMigrate(&models.Quote{})
	assert.NoError(t, err)

	originAmount, _ := decimal.NewFromString("200")
	destAmount, _ := decimal.NewFromString("100")

	// Create a dummy quote
	quote := models.Quote{
		Relayer:       "0xA",
		OriginChainId: 1,
		DestChainId:   1,
		OriginToken:   "0x1",
		DestToken:     "0x2",
		OriginAmount:  originAmount,
		DestAmount:    destAmount,
	}

	// Act
	id, err := database.InsertQuote(ctx, &quote)

	// Assert
	assert.Error(t, err, "Invalid quote: q.OriginChainId == q.DestChainId")
	assert.Equal(t, id, uint(0))
}

func TestBeforeCreateQuoteWhenOriginTokenZero(t *testing.T) {
	// Arrange
	ctx := context.Background()
	database := setupDatabase(t)
	err := database.AutoMigrate(&models.Quote{})
	assert.NoError(t, err)

	originAmount, _ := decimal.NewFromString("200")
	destAmount, _ := decimal.NewFromString("100")

	// Create a dummy quote
	quote := models.Quote{
		Relayer:       "0xA",
		OriginChainId: 1,
		DestChainId:   1,
		OriginToken:   "0x0000000000000000000000000000000000000000",
		DestToken:     "0x2",
		OriginAmount:  originAmount,
		DestAmount:    destAmount,
	}

	// Act
	id, err := database.InsertQuote(ctx, &quote)

	// Assert
	assert.Error(t, err, "Invalid quote: q.Tokens == address(0)")
	assert.Equal(t, id, uint(0))
}

func TestBeforeCreateQuoteWhenDestTokenZero(t *testing.T) {
	// Arrange
	ctx := context.Background()
	database := setupDatabase(t)
	err := database.AutoMigrate(&models.Quote{})
	assert.NoError(t, err)

	originAmount, _ := decimal.NewFromString("200")
	destAmount, _ := decimal.NewFromString("100")

	// Create a dummy quote
	quote := models.Quote{
		Relayer:       "0xA",
		OriginChainId: 1,
		DestChainId:   1,
		OriginToken:   "0x1",
		DestToken:     "0x0000000000000000000000000000000000000000",
		OriginAmount:  originAmount,
		DestAmount:    destAmount,
	}

	// Act
	id, err := database.InsertQuote(ctx, &quote)

	// Assert
	assert.Error(t, err, "Invalid quote: q.Tokens == address(0)")
	assert.Equal(t, id, uint(0))
}

func TestBeforeCreateQuoteWhenOriginAmountZero(t *testing.T) {
	// Arrange
	ctx := context.Background()
	database := setupDatabase(t)
	err := database.AutoMigrate(&models.Quote{})
	assert.NoError(t, err)

	destAmount, _ := decimal.NewFromString("100")

	// Create a dummy quote
	quote := models.Quote{
		Relayer:       "0xA",
		OriginChainId: 1,
		DestChainId:   1,
		OriginToken:   "0x1",
		DestToken:     "0x2",
		OriginAmount:  decimal.Zero,
		DestAmount:    destAmount,
	}

	// Act
	id, err := database.InsertQuote(ctx, &quote)

	// Assert
	assert.Error(t, err, "Invalid quote: q.Amounts == 0")
	assert.Equal(t, id, uint(0))
}

func TestBeforeCreateQuoteWhenDestAmountZero(t *testing.T) {
	// Arrange
	ctx := context.Background()
	database := setupDatabase(t)
	err := database.AutoMigrate(&models.Quote{})
	assert.NoError(t, err)

	originAmount, _ := decimal.NewFromString("200")

	// Create a dummy quote
	quote := models.Quote{
		Relayer:       "0xA",
		OriginChainId: 1,
		DestChainId:   1,
		OriginToken:   "0x1",
		DestToken:     "0x2",
		OriginAmount:  originAmount,
		DestAmount:    decimal.Zero,
	}

	// Act
	id, err := database.InsertQuote(ctx, &quote)

	// Assert
	assert.Error(t, err, "Invalid quote: q.Amounts == 0")
	assert.Equal(t, id, uint(0))
}

func TestBeforeSaveQuoteWhenSameChainId(t *testing.T) {
	// Arrange
	ctx := context.Background()
	database := setupDatabase(t)
	err := database.AutoMigrate(&models.Quote{})
	assert.NoError(t, err)

	originAmount, _ := decimal.NewFromString("200")
	destAmount, _ := decimal.NewFromString("100")

	// Create a dummy quote
	quote := models.Quote{
		Relayer:       "0xA",
		OriginChainId: 1,
		DestChainId:   2,
		OriginToken:   "0x1",
		DestToken:     "0x2",
		OriginAmount:  originAmount,
		DestAmount:    destAmount,
	}
	_, err = database.InsertQuote(ctx, &quote)
	assert.NoError(t, err)

	// retrieve the stored quote
	result := database.Last(&quote)
	assert.NoError(t, result.Error)

	// update origin amount to decrease price
	originAmount, _ = decimal.NewFromString("400")
	quote.OriginAmount = originAmount
	quote.DestChainId = 1

	// act
	err = database.UpdateQuote(ctx, &quote)

	// Assert
	assert.Error(t, err, "Invalid quote: q.OriginChainId == q.DestChainId")
}

func TestBeforeSaveQuoteWhenOriginTokenZero(t *testing.T) {
	// Arrange
	ctx := context.Background()
	database := setupDatabase(t)
	err := database.AutoMigrate(&models.Quote{})
	assert.NoError(t, err)

	originAmount, _ := decimal.NewFromString("200")
	destAmount, _ := decimal.NewFromString("100")

	// Create a dummy quote
	quote := models.Quote{
		Relayer:       "0xA",
		OriginChainId: 1,
		DestChainId:   2,
		OriginToken:   "0x1",
		DestToken:     "0x2",
		OriginAmount:  originAmount,
		DestAmount:    destAmount,
	}
	_, err = database.InsertQuote(ctx, &quote)
	assert.NoError(t, err)

	// retrieve the stored quote
	result := database.Last(&quote)
	assert.NoError(t, result.Error)

	// update origin amount to decrease price
	originAmount, _ = decimal.NewFromString("400")
	quote.OriginAmount = originAmount
	quote.OriginToken = "0x0000000000000000000000000000000000000000"

	// act
	err = database.UpdateQuote(ctx, &quote)

	// Assert
	assert.Error(t, err, "Invalid quote: q.Tokens == address(0)")
}

func TestBeforeSaveQuoteWhenDestTokenZero(t *testing.T) {
	// Arrange
	ctx := context.Background()
	database := setupDatabase(t)
	err := database.AutoMigrate(&models.Quote{})
	assert.NoError(t, err)

	originAmount, _ := decimal.NewFromString("200")
	destAmount, _ := decimal.NewFromString("100")

	// Create a dummy quote
	quote := models.Quote{
		Relayer:       "0xA",
		OriginChainId: 1,
		DestChainId:   2,
		OriginToken:   "0x1",
		DestToken:     "0x2",
		OriginAmount:  originAmount,
		DestAmount:    destAmount,
	}
	_, err = database.InsertQuote(ctx, &quote)
	assert.NoError(t, err)

	// retrieve the stored quote
	result := database.Last(&quote)
	assert.NoError(t, result.Error)

	// update origin amount to decrease price
	originAmount, _ = decimal.NewFromString("400")
	quote.OriginAmount = originAmount
	quote.DestToken = "0x0000000000000000000000000000000000000000"

	// act
	err = database.UpdateQuote(ctx, &quote)

	// Assert
	assert.Error(t, err, "Invalid quote: q.Tokens == address(0)")
}

func TestBeforeSaveQuoteWhenOriginAmountZero(t *testing.T) {
	// Arrange
	ctx := context.Background()
	database := setupDatabase(t)
	err := database.AutoMigrate(&models.Quote{})
	assert.NoError(t, err)

	originAmount, _ := decimal.NewFromString("200")
	destAmount, _ := decimal.NewFromString("100")

	// Create a dummy quote
	quote := models.Quote{
		Relayer:       "0xA",
		OriginChainId: 1,
		DestChainId:   2,
		OriginToken:   "0x1",
		DestToken:     "0x2",
		OriginAmount:  originAmount,
		DestAmount:    destAmount,
	}
	_, err = database.InsertQuote(ctx, &quote)
	assert.NoError(t, err)

	// retrieve the stored quote
	result := database.Last(&quote)
	assert.NoError(t, result.Error)

	// update origin amount to decrease price
	quote.OriginAmount = decimal.Zero

	// act
	err = database.UpdateQuote(ctx, &quote)

	// Assert
	assert.Error(t, err, "Invalid quote: q.Amounts == 0")
}

func TestBeforeSaveQuoteWhenDestAmountZero(t *testing.T) {
	// Arrange
	ctx := context.Background()
	database := setupDatabase(t)
	err := database.AutoMigrate(&models.Quote{})
	assert.NoError(t, err)

	originAmount, _ := decimal.NewFromString("200")
	destAmount, _ := decimal.NewFromString("100")

	// Create a dummy quote
	quote := models.Quote{
		Relayer:       "0xA",
		OriginChainId: 1,
		DestChainId:   2,
		OriginToken:   "0x1",
		DestToken:     "0x2",
		OriginAmount:  originAmount,
		DestAmount:    destAmount,
	}
	_, err = database.InsertQuote(ctx, &quote)
	assert.NoError(t, err)

	// retrieve the stored quote
	result := database.Last(&quote)
	assert.NoError(t, result.Error)

	// update dest amount to decrease price
	quote.DestAmount = decimal.Zero

	// act
	err = database.UpdateQuote(ctx, &quote)

	// Assert
	assert.Error(t, err, "Invalid quote: q.Amounts == 0")
}
