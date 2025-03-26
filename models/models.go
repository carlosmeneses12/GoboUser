package models //definir estructuras

type SecretRDSJson struct {
	Username            string `json:"username"` //`json:"username"` es un tag
	Password            string `json:"password"`
	Engine              string `json:"engine"`
	Host                string `json:"host"`
	Port                int    `json:"port"`
	DbClusterIdentifier string `json:"dbClusterIdentifier"`
}

type SignUp struct {
	Email  string `json:"email"`
	UserID string `json:"userID"`
}
