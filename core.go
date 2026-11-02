package core

import "errors"

type Provider struct {
	Name             string
	Models           []string
	Healthy          bool
	Active, Capacity int
	CostMicros       int64
}

func Choose(providers []Provider, model string) (Provider, error) {
	var chosen Provider
	found := false
	for _, p := range providers {
		if !p.Healthy || p.Active < 0 || p.Capacity <= p.Active || p.CostMicros < 0 {
			continue
		}
		supports := false
		for _, m := range p.Models {
			if m == model {
				supports = true
			}
		}
		if supports && (!found || p.CostMicros < chosen.CostMicros || (p.CostMicros == chosen.CostMicros && p.Name < chosen.Name)) {
			chosen = p
			found = true
		}
	}
	if !found {
		return Provider{}, errors.New("no eligible provider")
	}
	return chosen, nil
}
