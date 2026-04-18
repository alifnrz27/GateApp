package validation

import (
	"GateApp/backend/utils"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type TriggerGateRequest struct {
	GateUUID string `json:"gate_uuid" validate:"required"`
	Trigger  string `json:"trigger" validate:"required,oneof=open"`
	Relay    int    `json:"relay"`
}

var validationMessages = map[string]string{
	"gate_uuid.required": "Gate UUID wajib diisi",

	"trigger.required": "Trigger wajib diisi",
	"trigger.oneof":    "Trigger hanya boleh 'open'",
}

func ValidateTriggerGate(req *TriggerGateRequest) map[string]string {
	return utils.ValidateStruct(validate, req, validationMessages)
}
