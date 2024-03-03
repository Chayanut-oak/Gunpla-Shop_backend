package omise

import (
	"log"
	"os"

	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/restModel"
	"github.com/joho/godotenv"
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
)

func GetOmiseClient() (*omise.Client, error) {
	godotenv.Load()
	publicKey := os.Getenv("OMISE_PUBLIC_KEY")
	secretKey := os.Getenv("OMISE_SECRET_KEY")
	client, e := omise.NewClient(publicKey, secretKey)
	if e != nil {
		return nil, e
	}
	return client, nil
}

func CreateToken(client *omise.Client, payment restModel.PaymentRestModel) (string, error) {
	result := &omise.Card{}

	err := client.Do(result, &operations.CreateToken{
		Name:            payment.Name,
		Number:          payment.Number,
		ExpirationMonth: payment.ExpirationMonth,
		ExpirationYear:  payment.ExpirationYear,
		SecurityCode:    payment.Cvc,
	})
	if err != nil {
		return "", err
	}
	log.Println(result)
	return result.ID, nil
}

func CreateChargeByToken(client *omise.Client, token string, orderId string, total int64) error {
	result := &omise.Charge{}
	err := client.Do(result, &operations.CreateCharge{
		Amount:      total * 100,
		Currency:    "thb",
		Description: "Gunpla Store customer charge",
		Card:        token,
		Metadata: map[string]interface{}{
			"OrderID": orderId,
		},
	})
	if err != nil {
		return err
	}
	log.Println(result)
	return nil
}
