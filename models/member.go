package models

type (
	// Member
	Member struct {
		ID        int    `json:"ID" body:"ID"`
		USERNAME  string `json:"USERNAME" body:"USERNAME"`
		GENDER    string `json:"GENDER" body:"GENDER"`
		SKINTYPE  string `json:"SKINTYPE" body:"SKINTYPE"`
		SKINCOLOR string `json:"SKINCOLOR" body:"SKINCOLOR"`
	}

	Product struct {
		ID           int    `json:"ID" body:"ID"`
		NAME_PRODUCT string `json:"NAME_PRODUCT" body:"NAME_PRODUCT"`
		PRICE        string `json:"PRICE" body:"PRICE"`
	}

	Review struct {
		ID                 int    `json:"ID" body:"ID"`
		ID_PRODUCT         int    `json:"ID_PRODUCT" body:"ID_PRODUCT"`
		ID_MEMBER          int    `json:"ID_MEMBER" body:"ID_MEMBER"`
		DESC_REVIEW        string `json:"DESC_REVIEW" body:"DESC_REVIEW"`
		JUMLAH_LIKE_REVIEW int    `json:"JUMLAH_LIKE_REVIEW" body:"JUMLAH_LIKE_REVIEW"`
	}

	Like_Review struct {
		ID_REVIEW int `json:"ID_REVIEW" body:"ID_REVIEW"`
		ID_MEMBER int `json:"ID_MEMBER" body:"ID_MEMBER"`
	}
)
