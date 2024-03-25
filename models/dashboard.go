package models

type Period struct {
    YearMonth string `json:"yearmonth"`
    Label     string `json:"label"`
}

type CertificateByTenor struct {
    Tenor        int     `json:"tenor"`
    GrossPremium float64 `json:"gross_premium"`
    TSI          float64 `json:"tsi"`
    Certificate  int     `json:"certificate"`
}

type Gender struct {
    Gender string  `json:"gender"`
    Total  int     `json:"total"`
    Ratio  float64 `json:"ratio"`
}

type SummaryData struct {
    YearMonth         string                `json:"yearmonth"`
    CreatedAt         string                `json:"created_at"`
    GrossPremium      float64               `json:"gross_premium"`
    TSI               float64               `json:"tsi"`
    Certificate       int                   `json:"certificate"`
    PolicyRejected    int                   `json:"policy_rejected"`
    CertificateByTenor []CertificateByTenor `json:"certificate_by_tenor"`
    Gender            []Gender              `json:"gender"`
    Age               []interface{}         `json:"age"`
    Occupation        []interface{}         `json:"occupation"`
    GrossPremiumCancel int                   `json:"gross_premium_cancel"`
    Branch            []interface{}         `json:"branch"`
    SubmittedAmount   int                   `json:"submitted_amount"`
    Submitted         int                   `json:"submitted"`
    SubmittedDate     int                   `json:"submitted_date"`
    AcceptedAmount    int                   `json:"accepted_amount"`
    Accepted          int                   `json:"accepted"`
    AcceptedDate      string                `json:"accepted_date"`
    RejectedAmount    int                   `json:"rejected_amount"`
    Rejected          int                   `json:"rejected"`
    RejectedDate      string                `json:"rejected_date"`
    ClaimByTenor      []interface{}         `json:"claim_by_tenor"`
}

type QueryResult struct {
    YearMonth       string             `json:"yearmonth"`
    CreatedAt       string             `json:"created_at"`
    GrossPremium    float64            `json:"gross_premium"`
    TSI             float64            `json:"tsi"`
    Certificate     int                `json:"certificate"`
    PolicyRejected  int                `json:"policy_rejected"`
    CertificateByTenor []CertificateByTenor `json:"certificate_by_tenor"`
    Gender          []Gender           `json:"gender"`
    Age             []interface{}      `json:"age"`
    Occupation      []interface{}      `json:"occupation"`
    GrossPremiumCancel *int             `json:"gross_premium_cancel"`
    Branch          []interface{}      `json:"branch"`
    SubmittedAmount *int               `json:"submitted_amount"`
    Submitted       *int               `json:"submitted"`
    SubmittedDate   *int               `json:"submitted_date"`
    AcceptedAmount  *int               `json:"accepted_amount"`
    Accepted        *int               `json:"accepted"`
    AcceptedDate    *string            `json:"accepted_date"`
    RejectedAmount  *int               `json:"rejected_amount"`
    Rejected        *int               `json:"rejected"`
    RejectedDate    *string            `json:"rejected_date"`
    ClaimByTenor    []interface{}      `json:"claim_by_tenor"`
}

// func NewSummaryDataWithDefault() SummaryData {
//     return SummaryData{
//         GrossPremiumCancel: 0,
//         Branch:             []interface{}{},
//         SubmittedAmount:    0,
//         Submitted:          0,
//         SubmittedDate:      0,
//         AcceptedAmount:     0,
//         Accepted:           0,
//         AcceptedDate:       "",
//         RejectedAmount:     0,
//         Rejected:           0,
//         RejectedDate:       "", // Gunakan string kosong untuk menunjukkan bahwa tanggal penolakan tidak ada
//         ClaimByTenor:       []interface{}{},
//     }
// }
