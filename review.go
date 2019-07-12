package schema_org

import "time"

type Review struct {
	CreativeWork
	Thing
	Type         string        `json:"@type"`
	ItemReviewed *Thing        `json:"itemReviewed,omitempty"`
	ReviewBody   string        `json:"reviewBody"`
	Review       []Review      `json:"review,omitempty"`
	ReviewRating *ReviewRating `json:"reviewRating"`
}

type Thing struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Url         string `json:"url,omitempty"`
}

type CreativeWork struct {
	DateCreated   time.Time `json:"dateCreated"`
	DatePublished time.Time `json:"datePublished"`
}

func NewReview() Review {
	return Review{
		Type: "Review",
	}
}

type ReviewRating struct {
	Type        string `json:"@type"`
	BestRating  string `json:"bestRating"`
	RatingValue string `json:"ratingValue"`
	WorstRating string `json:"worstRating"`
}

func NewReviewRating() *ReviewRating {
	return &ReviewRating{
		Type:        "Rating",
		BestRating:  "5",
		WorstRating: "1",
	}
}

/*
"reviewRating": {
	"@type": "Rating",
	"bestRating": "5",
	"ratingValue": "1",
	"worstRating": "1"
}*/
