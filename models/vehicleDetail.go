package models

import "go.mongodb.org/mongo-driver/bson/primitive"

/* type VehicleDetails struct {
	ID          primitive.ObjectID	`bson:"_id,omitempty" json:"id,omitempty"`
	Title       string				`bson:"title,omitempty" json:"title,omitempty"`
	Description string				`bson:"description,omitempty" json:"description,omitempty"`
} */
type VehicleDetails struct {
	ID            primitive.ObjectID	`bson:"_id,omitempty" json:"id,omitempty"`
	Vehicle       Vehicle				`bson:"vehicle,omitempty" json:"vehicle,omitempty"`
	Customer      Customer				`bson:"customer,omitempty" json:"customer,omitempty"`
	Insurance     Insurance				`bson:"insurance,omitempty" json:"insurance,omitempty"`
	
}
type Vehicle struct {
	RegistrationNumber	string   `bson:"registrationNumber,omitempty" json:"registrationNumber,omitempty"`
	VehicleMake	        string   `bson:"vehicleMake,omitempty" json:"vehicleMake,omitempty"`
	VehicleModel	    string   `bson:"vehicleModel,omitempty" json:"vehicleModel,omitempty"`
	VehicleVariant	    string   `bson:"vehicleVariant,omitempty" json:"vehicleVariant,omitempty"`
	RegistrationDate	string   `bson:"registrationDate,omitempty" json:"registrationDate,omitempty"`
	ManufactureYear	    string   `bson:"manufactureYear,omitempty" json:"manufactureYear,omitempty"`
	RtoCode	            string   `bson:"rtoCode,omitempty" json:"rtoCode,omitempty"`
	RtoLocation	        string   `bson:"rtoLocation,omitempty" json:"rtoLocation,omitempty"`
	FuelType	        string   `bson:"fuelType,omitempty" json:"fuelType,omitempty"`
	EngineNo	        string   `bson:"engineNo,omitempty" json:"engineNo,omitempty"`
	ChasisNo            string   `bson:"chasisNo,omitempty" json:"chasisNo,omitempty"`
}
type Individual struct {
	FirstName	string   `bson:"firstName,omitempty" json:"firstName,omitempty"`
	LastName	string   `bson:"lastName,omitempty" json:"lastName,omitempty"`
}
type Organisation struct {
	OrganisationName	string   `bson:"organisationName,omitempty" json:"organisationName,omitempty"`
}
type Customer struct {
	IndividualOrOrganisation	string   `bson:"individualOrOrganisation,omitempty" json:"individualOrOrganisation,omitempty"`
	Individual                  Individual   `bson:"individual,omitempty" json:"individual,omitempty"`
	Organisation                Organisation   `bson:"organisation,omitempty" json:"organisation,omitempty"`
	PermanentAddress	        string   `bson:"permanentAddress,omitempty" json:"permanentAddress,omitempty"`
	CommunicationAddress	    string   `bson:"communicationAddress,omitempty" json:"communicationAddress,omitempty"`
	PhoneNumber	                string   `bson:"phoneNumber,omitempty" json:"phoneNumber,omitempty"`
	EmailAddress	            string   `bson:"emailAddress,omitempty" json:"emailAddress,omitempty"`
	DateOfBirth	                string   `bson:"dateOfBirth,omitempty" json:"dateOfBirth,omitempty"`
	NomineeName	                string   `bson:"nomineeName,omitempty" json:"nomineeName,omitempty"`
	NomineeAge	                string   `bson:"nomineeAge,omitempty" json:"nomineeAge,omitempty"`
	NomineeRelationship	        string   `bson:"nomineeRelationship,omitempty" json:"nomineeRelationship,omitempty"`
}
type Insurance struct {
	InsuranceProvider	string   `bson:"insuranceProvider,omitempty" json:"insuranceProvider,omitempty"`
	PolicyNumber	    string   `bson:"policyNumber,omitempty" json:"policyNumber,omitempty"`
	PolicyType	        string   `bson:"policyType,omitempty" json:"policyType,omitempty"`
	PolicyTerm	        string   `bson:"policyTerm,omitempty" json:"policyTerm,omitempty"`
	RiskStartDate	    string   `bson:"riskStartDate,omitempty" json:"riskStartDate,omitempty"`
	RiskEndDate	        string   `bson:"riskEndDate,omitempty" json:"riskEndDate,omitempty"`
}
