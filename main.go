package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type ConvertyQuery struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}

type ConvertionResult struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
	Result float64 `json:"result"`
	Rate   float64 `json:"rate"`
}

type ExchangeRateResponse struct {
	Bid string `json:"bid"`
}

func getExchangeRate(from string, to string) (float64, error) {
	if from == to {
		return 1.0, nil
	}

	url := fmt.Sprintf("https://economia.awesomeapi.com.br/json/last/%s-%s", from, to)

	req, err := http.Get(url)
	if err != nil {
		return 0.0, fmt.Errorf("falha ao conectar no serviço de cotação: %w", err)
	}
	defer req.Body.Close()

	if req.StatusCode != http.StatusOK {
		return 0.0, fmt.Errorf("par de moedas '%s-%s' não suportado", from, to)
	}

	var apiResult map[string]ExchangeRateResponse

	if err := json.NewDecoder(req.Body).Decode(&apiResult); err != nil {
		return 0.0, fmt.Errorf("erro ao ler a resposta do provedor: %w", err)
	}

	rateData, exists := apiResult[from+to]
	if !exists {
		return 0.0, fmt.Errorf("cotação para o par %s-%s não encontrada", from, to)
	}

	rate, err := strconv.ParseFloat(rateData.Bid, 64)
	if err != nil {
		return 0.0, fmt.Errorf("erro ao converter taxa: %w", err)
	}

	return rate, nil
}

func convertHandler(c *gin.Context) {
	var query ConvertyQuery
	if err := c.ShouldBindJSON(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	from := strings.ToUpper(query.From)
	to := strings.ToUpper(query.To)

	rate, err := getExchangeRate(from, to)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ConvertionResult{
		From:   from,
		To:     to,
		Amount: query.Amount,
		Result: query.Amount * rate,
		Rate:   rate,
	})
}

func main() {
	router := gin.Default()
	router.POST("/convert", convertHandler)
	router.Run(":8080")
}
