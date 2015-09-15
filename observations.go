package ptgen

import (
	"math/rand"

	"github.com/intervention-engine/fhir/models"
)

func GenerateBP(ctx Context) []models.Observation {
	sys, dia := models.Observation{}, models.Observation{}
	sys.Code = &models.CodeableConcept{Coding: []models.Coding{{Code: "271649006", System: "http://snomed.info/sct"}}, Text: "Systolic Blood Pressure"}
	dia.Code = &models.CodeableConcept{Coding: []models.Coding{{Code: "271650006", System: "http://snomed.info/sct"}}, Text: "Diastolic Blood Pressure"}
	switch ctx.Hypertention {
	case "Normal":
		sys.ValueQuantity = GenerateQuantity(100, 120)
		dia.ValueQuantity = GenerateQuantity(65, 80)
	case "Pre-hypertension":
		sys.ValueQuantity = GenerateQuantity(120, 140)
		dia.ValueQuantity = GenerateQuantity(80, 90)
	case "Hypertension":
		sys.ValueQuantity = GenerateQuantity(140, 180)
		dia.ValueQuantity = GenerateQuantity(90, 120)
	}
	sys.ValueQuantity.Units = "mmHg"
	dia.ValueQuantity.Units = "mmHg"

	return []models.Observation{sys, dia}
}

func GenerateCholesterol(ctx Context) []models.Observation {
	ldl, hdl, tri := models.Observation{}, models.Observation{}, models.Observation{}
	ldl.Code = &models.CodeableConcept{Coding: []models.Coding{{Code: "314036004", System: "http://snomed.info/sct"}}, Text: "Plasma LDL Cholesterol Measurement"}
	hdl.Code = &models.CodeableConcept{Coding: []models.Coding{{Code: "314035000", System: "http://snomed.info/sct"}}, Text: "Plasma HDL Cholesterol Measurement"}
	tri.Code = &models.CodeableConcept{Coding: []models.Coding{{Code: "167082000", System: "http://snomed.info/sct"}}, Text: "Plasma Triglyceride Measurement"}

	switch ctx.Cholesterol {
	case "Optimal":
		ldl.ValueQuantity = GenerateQuantity(80, 100)
		hdl.ValueQuantity = GenerateQuantity(60, 70)
		tri.ValueQuantity = GenerateQuantity(100, 140)
	case "Near Optimal":
		ldl.ValueQuantity = GenerateQuantity(100, 130)
		hdl.ValueQuantity = GenerateQuantity(50, 60)
		tri.ValueQuantity = GenerateQuantity(140, 160)
	case "Borderline":
		ldl.ValueQuantity = GenerateQuantity(130, 150)
		hdl.ValueQuantity = GenerateQuantity(40, 60)
		tri.ValueQuantity = GenerateQuantity(160, 200)
	case "High":
		ldl.ValueQuantity = GenerateQuantity(160, 200)
		hdl.ValueQuantity = GenerateQuantity(40, 50)
		tri.ValueQuantity = GenerateQuantity(200, 300)
	case "Very High":
		ldl.ValueQuantity = GenerateQuantity(190, 220)
		hdl.ValueQuantity = GenerateQuantity(30, 40)
		tri.ValueQuantity = GenerateQuantity(300, 400)
	}

	ldl.ValueQuantity.Units = "mg/dL"
	hdl.ValueQuantity.Units = "mg/dL"
	tri.ValueQuantity.Units = "mg/dL"

	return []models.Observation{ldl, hdl, tri}
}

func GenerateWeightAndHeight(patient models.Patient) []models.Observation {
	w, h := models.Observation{}, models.Observation{}
	if patient.Gender == "male" {
		w.ValueQuantity = GenerateQuantity(100, 300)
		h.ValueQuantity = GenerateQuantity(60, 80)
	} else {
		w.ValueQuantity = GenerateQuantity(80, 250)
		h.ValueQuantity = GenerateQuantity(55, 75)
	}

	w.ValueQuantity.Units = "lbs"
	h.ValueQuantity.Units = "in"

	return []models.Observation{w, h}
}

func GenerateQuantity(min, max int) *models.Quantity {
	q := float64(min + rand.Intn(max-min))
	return &models.Quantity{Value: &q}
}
