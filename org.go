package schema_org

type Organization struct {
	Thing
	Context         string           `json:"@context"`
	Type            string           `json:"@type"`
	ContactPoint    []ContactPoint   `json:"contactPoint,omitempty"`
	Email           string           `json:"email,omitempty"`
	Location        string           `json:"location,omitempty"` // address
	Telephone       string           `json:"telephone,omitempty"`
	Review          []Review         `json:"review,omitempty"`
	AggregateRating *AggregateRating `json:"aggregateRating,omitempty"`
}

type ContactPoint struct {
	Type        string `json:"@type"`
	Telephone   string `json:"telephone"`
	ContactType string `json:"contactType"`
}

type AggregateRating struct {
	Type        string `json:"@type"`
	RatingValue string `json:"ratingValue"`
	ReviewCount string `json:"reviewCount"`
}

func NewAggregateRating() *AggregateRating {
	return &AggregateRating{
		Type: "AggregateRating",
	}
}

/*"aggregateRating": {
	"@type": "AggregateRating",
	"ratingValue": "3.5",
	"reviewCount": "11"
},*/

func NewOrganization() *Organization {
	return &Organization{
		Context: "http://schema.org",
		Type:    "Organization",
	}
}

func (o *Organization) AddContact(number, info string) {
	o.ContactPoint = append(o.ContactPoint, ContactPoint{
		Type:        "ContactPoint",
		Telephone:   number,
		ContactType: info,
	})
}
