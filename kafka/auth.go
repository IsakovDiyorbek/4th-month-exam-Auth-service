package kafka

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Exam4/4th-month-exam-Auth-service/genproto/auth"
	"github.com/Exam4/4th-month-exam-Auth-service/service"
)

func AuhtRegister(Auth *service.AuthService) func(message []byte) {
	return func(message []byte) {
		var eval auth.RegisterRequest
		if err := json.Unmarshal(message, &eval); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}
		respEval, err := Auth.Register(context.Background(), &eval)
		if err != nil {
			log.Printf("Cannot user register via Kafka: %v", err)
			return
		}
		log.Printf("Register user via Kafka: %+v", respEval)
	}
}
