package main

import (
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

// Define the Doctor type
var doctorType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Doctor",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"specialty": &graphql.Field{
				Type: graphql.String,
			},
			"contactInfo": &graphql.Field{
				Type: contactInfoType,
			},
		},
	},
)

// Define the Patient type
var patientType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Patient",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"age": &graphql.Field{
				Type: graphql.Int,
			},
			"gender": &graphql.Field{
				Type: graphql.String,
			},
			"contactInfo": &graphql.Field{
				Type: contactInfoType,
			},
			"medicalHistory": &graphql.Field{
				Type: graphql.NewList(medicalHistoryType),
			},
			"appointments": &graphql.Field{
				Type: graphql.NewList(appointmentType),
			},
			"medications": &graphql.Field{
				Type: graphql.NewList(medicationType),
			},
		},
	},
)

// Define the ContactInfo type
var contactInfoType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ContactInfo",
		Fields: graphql.Fields{
			"phone": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"address": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

// Define the MedicalHistory type
var medicalHistoryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "MedicalHistory",
		Fields: graphql.Fields{
			"date": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"doctor": &graphql.Field{
				Type: doctorType,
			},
		},
	},
)

// Define the Appointment type
var appointmentType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Appointment",
		Fields: graphql.Fields{
			"date": &graphql.Field{
				Type: graphql.String,
			},
			"time": &graphql.Field{
				Type: graphql.String,
			},
			"doctor": &graphql.Field{
				Type: doctorType,
			},
			"purpose": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

// Define the Medication type
var medicationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Medication",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"dosage": &graphql.Field{
				Type: graphql.String,
			},
			"frequency": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

// Sample data for patients
var patients = []map[string]interface{}{
	{
		"id":     "12345",
		"name":   "John Doe",
		"age":    30,
		"gender": "male",
		"contactInfo": map[string]interface{}{
			"phone":   "123-456-7890",
			"email":   "john.doe@example.com",
			"address": "123 Main St, Anytown, USA",
		},
		"medicalHistory": []map[string]interface{}{
			{
				"date":        "2022-01-01",
				"description": "Annual Checkup",
				"doctor": map[string]interface{}{
					"id":        "d1",
					"name":      "Dr. Smith",
					"specialty": "General Medicine",
				},
			},
			{
				"date":        "2023-01-01",
				"description": "Follow-up Visit",
				"doctor": map[string]interface{}{
					"id":        "d1",
					"name":      "Dr. Smith",
					"specialty": "General Medicine",
				},
			},
		},
		"appointments": []map[string]interface{}{
			{
				"date": "2023-06-20",
				"time": "10:00 AM",
				"doctor": map[string]interface{}{
					"id":        "d1",
					"name":      "Dr. Smith",
					"specialty": "General Medicine",
				},
				"purpose": "Routine Checkup",
			},
			{
				"date": "2023-07-20",
				"time": "02:00 PM",
				"doctor": map[string]interface{}{
					"id":        "d2",
					"name":      "Dr. Jones",
					"specialty": "Cardiology",
				},
				"purpose": "Cardiac Consultation",
			},
		},
		"medications": []map[string]interface{}{
			{
				"name":      "Medication A",
				"dosage":    "10mg",
				"frequency": "Once a day",
			},
			{
				"name":      "Medication B",
				"dosage":    "5mg",
				"frequency": "Twice a day",
			},
		},
	},
	{
		"id":     "67890",
		"name":   "Jane Smith",
		"age":    25,
		"gender": "female",
		"contactInfo": map[string]interface{}{
			"phone":   "987-654-3210",
			"email":   "jane.smith@example.com",
			"address": "456 Oak St, Anycity, USA",
		},
		"medicalHistory": []map[string]interface{}{
			{
				"date":        "2021-06-15",
				"description": "First Consultation",
				"doctor": map[string]interface{}{
					"id":        "d3",
					"name":      "Dr. Brown",
					"specialty": "Dermatology",
				},
			},
		},
		"appointments": []map[string]interface{}{
			{
				"date": "2023-05-10",
				"time": "11:00 AM",
				"doctor": map[string]interface{}{
					"id":        "d3",
					"name":      "Dr. Brown",
					"specialty": "Dermatology",
				},
				"purpose": "Skin Checkup",
			},
		},
		"medications": []map[string]interface{}{
			{
				"name":      "Medication C",
				"dosage":    "15mg",
				"frequency": "Once a day",
			},
		},
	},
}

// Sample data for doctors
var doctors = []map[string]interface{}{
	{
		"id":        "d1",
		"name":      "Dr. Smith",
		"specialty": "General Medicine",
		"contactInfo": map[string]interface{}{
			"phone":   "555-1234",
			"email":   "smith@hospital.com",
			"address": "123 Medical St, Medcity, USA",
		},
	},
	{
		"id":        "d2",
		"name":      "Dr. Jones",
		"specialty": "Cardiology",
		"contactInfo": map[string]interface{}{
			"phone":   "555-5678",
			"email":   "jones@hospital.com",
			"address": "456 Heart Ave, Cardiotown, USA",
		},
	},
	{
		"id":        "d3",
		"name":      "Dr. Brown",
		"specialty": "Dermatology",
		"contactInfo": map[string]interface{}{
			"phone":   "555-8765",
			"email":   "brown@hospital.com",
			"address": "789 Skin Blvd, Dermacity, USA",
		},
	},
}

// Define the Query type
var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"patient": &graphql.Field{
				Type: patientType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, ok := params.Args["id"].(string)
					if ok {
						for _, patient := range patients {
							if patient["id"] == id {
								return patient, nil
							}
						}
					}
					return nil, nil
				},
			},
			"doctor": &graphql.Field{
				Type: doctorType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, ok := params.Args["id"].(string)
					if ok {
						for _, doctor := range doctors {
							if doctor["id"] == id {
								return doctor, nil
							}
						}
					}
					return nil, nil
				},
			},
		},
	},
)

func main() {
	// Define the Schema
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})
	if err != nil {
		log.Fatalf("failed to create schema, error: %v", err)
	}

	// Create a new HTTP handler
	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	http.Handle("/graphql", h)

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
