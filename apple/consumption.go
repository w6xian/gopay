package apple

import (
	"context"
	"fmt"
	"net/http"

	"github.com/w6xian/gopay"
)

// SendConsumptionInformation Send Consumption Information
// Doc: https://developer.apple.com/documentation/appstoreserverapi/send_consumption_information
func (c *Client) SendConsumptionInformation(ctx context.Context, transactionId string, bm gopay.BodyMap) (err error) {
	path := fmt.Sprintf(sendConsumptionInformation, transactionId)
	res, _, err := c.doRequestPut(ctx, path, bm)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusAccepted {
		return fmt.Errorf("http.stauts_code = %d", res.StatusCode)
	}
	return nil
}
