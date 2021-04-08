package db

var mcList []*MyCollections

func init() {
	mcList = []*MyCollections{
		&MyCollections{
			UserID:     "3c7c9e33-5d38-498c-88f5-77b125564b3b",
			Collection: "Избранное",
			BookID:     "dee44255-bf77-4e3c-8a2e-ca35ae05f860",
		},
		&MyCollections{
			UserID:     "3c7c9e33-5d38-498c-88f5-77b125564b3b",
			Collection: "Избранное",
			BookID:     "83431311-ceec-4637-aea1-0cfc4e5dd3ad",
		},
	}
}
