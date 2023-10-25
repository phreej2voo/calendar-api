package jobs

import (
	"calendar-api/services"
	"fmt"

	"github.com/jrallison/go-workers"
)

func AddCrmLeads(message *workers.Msg) {
	phoneNumber := message.Args()
	params := map[string]interface{}{
		"phone": phoneNumber,
	}

	respBody := services.CrmPost("assignments/calendar_assign", params)
	fmt.Println(respBody)
}
