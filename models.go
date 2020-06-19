package models

type Agent struct{
	ID              int         `json:"id"`
	Agent_Code       string      `json:"agent"`
	Agent_Name       string      `json:"name"`
	Working_Area     string      `json:"working_area"`
	Commission        int         `json:"commisssion"`
	Phone_no         int         `json:"phone_no"`
	Country          string      `json:"country"`
}

type ResponeAgent struct{
	Agent []Agent `json:"Agent"`
}

type Customer struct{
	ID               int         `json:"id"`
	Cust_Code       string      `json:"customer"`
	Cust_Name       string      `json:"name"`
	Working_Area     string      `json:"working_area"`
	Opening_Amt   int         `json:"opening_amt"`
	Receiving_Amt   int        `json:"receiving_amt"`
	Pending_Amt   int         `json:"pending_amt"`
	Commission        int         `json:"commisssion"`
	Phone_no         int         `json:"phone_no"`
	Country          string      `json:"country"`
}

type ResponseCustomer struct{
	Customer []Customer `json:"Customer"`
}

type Order struct{
	ID               int         `json:"id"`
	Ord_Num   int  `json:"ord_num"`
	Ord_Amt   int  `json:"ord_amt"`
	Adv_Amt   int  `json:"adv_amt"`
	Ord_Date  string  `json:"ord_date"`
	Cust_Code  string  `json:"customer"`
	Agent_Code  string  `json:"agent"`
}

type  ResponseOrder struct{
	Order []Order  `json:"Order"`
}
