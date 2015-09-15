package ptgen

import (
	"math/rand"
	"time"

	"github.com/icrowley/fake"
	"github.com/intervention-engine/fhir/models"
	"github.com/jmcvetta/randutil"
)

// Context contains information about the patient that can be used when
// generating information
type Context struct {
	Smoker       string
	Hypertention string
	Alcohol      string
	Cholesterol  string
}

func GenerateDemographics() models.Patient {
	patient := models.Patient{}
	patient.Gender = fake.Gender()
	name := models.HumanName{}
	var firstName string
	if patient.Gender == "male" {
		firstName = fake.MaleFirstName()
	} else {
		firstName = fake.FemaleFirstName()
	}
	name.Given = []string{firstName}
	name.Family = []string{fake.LastName()}
	patient.Name = []models.HumanName{name}
	patient.BirthDate = &models.FHIRDateTime{Time: RandomBirthDate(), Precision: models.Date}
	patient.Address = []models.Address{GenerateAddress()}
	return patient
}

// RandomBirthDate generates a random birth date between 65 and 85 years ago
func RandomBirthDate() time.Time {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomYears := r.Intn(20)
	yearsAgo := randomYears + 65
	randomMonth := r.Intn(11)
	randomDay := r.Intn(28)
	t := time.Now()
	return t.AddDate(-yearsAgo, -randomMonth, -randomDay).Truncate(time.Hour * 24)
}

func GenerateAddress() models.Address {
	address := models.Address{}
	address.Line = []string{fake.Street()}
	address.City = fake.City()
	address.State = fake.StateAbbrev()
	address.PostalCode = fake.Zip()
	return address
}

// NewContext generates a new context with randomly populated content
func NewContext() Context {
	ctx := Context{}
	smokingChoices := []randutil.Choice{
		{2, "Smoker"},
		{3, "Non-smoker"},
		{1, "Ex-smoker"}}
	sc, _ := randutil.WeightedChoice(smokingChoices)
	ctx.Smoker = sc.Item.(string)

	alcoholChoices := []randutil.Choice{
		{2, "Occasional"},
		{1, "Heavy"},
		{1, "None"}}
	ac, _ := randutil.WeightedChoice(alcoholChoices)
	ctx.Alcohol = ac.Item.(string)

	cholesterolChoices := []randutil.Choice{
		{3, "Optimal"},
		{1, "Near Optimal"},
		{2, "Borderline"},
		{1, "High"},
		{2, "Very High"}}
	cc, _ := randutil.WeightedChoice(cholesterolChoices)
	ctx.Cholesterol = cc.Item.(string)

	hc, _ := randutil.ChoiceString([]string{"Normal", "Pre-hypertension", "Hypertension"})
	ctx.Hypertention = hc
	return ctx
}
