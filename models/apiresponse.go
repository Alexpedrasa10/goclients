package models

type APIResponse struct {
    Code int         `json:"code"` 
    Data interface{} `json:"data"`
}
