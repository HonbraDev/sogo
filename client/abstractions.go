package client

import (
	"github.com/HonbraDev/sogo/models"
)

func (c *Client) GetAuthStatus() (bool, error) {
	var ok bool
	err := c.Get("/AuthorizationStatus", &ok)
	return ok, err
}

func (c *Client) GetUzivatelInfo(username string) (*models.UzivatelInfo, error) {
	if username == "" {
		username = c.Username
	}
	var userInfo models.UzivatelInfo
	err := c.Get("/UzivatelInfo/"+username, &userInfo)
	return &userInfo, err
}

func (c *Client) GetRozvrhoveUdalostiDay(date string) ([]models.RozvrhovaUdalost, error) {
	var w struct {
		Udalosti []models.RozvrhovaUdalost `json:"UDALOSTI"`
	}
	err := c.Get("/RozvrhoveUdalosti/"+date, &w)
	return w.Udalosti, err
}

func (c *Client) GetRozvrhoveUdalostiRange(d1, d2 string) ([]models.RozvrhovaUdalost, error) {
	var w struct {
		Udalosti []models.RozvrhovaUdalost `json:"UDALOSTI"`
	}
	err := c.Get("/RozvrhoveUdalosti/"+d1+"/"+d2, &w)
	return w.Udalosti, err
}
