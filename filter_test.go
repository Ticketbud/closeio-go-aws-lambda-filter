package filter

import (
	"testing"
	"github.com/Ticketbud/closeio-go"
)

func TestSearch(t *testing.T) {
	// using the first lead we found in close

	config := ReadConfig("./config.json")

	client := closeio.New(config.CloseKey)

	if leads := Search(client, "email", "sydney.rhea@gmail.com"); len(leads.Data) == 0 {
		t.Errorf("did not find the expected lead")
	}
}

func TestBuildCustomLeadSource(t *testing.T) {
	if leadUpdateMap := BuildCustomLeadSource("Unbounce"); len(leadUpdateMap) != 2 {
		t.Errorf("was not able to build the lead update map")
	}

}

func TestBuildCustomLeadSourceUpdates(t *testing.T) {
	
	config := ReadConfig("./config.json")
	
	client := closeio.New(config.CloseKey)

	leads := Search(client, "email", "sydney.rhea@gmail.com")

	if leadsBuilt := BuildCustomLeadSourceUpdates(leads, "Unbounce"); len(leadsBuilt) == 0 {
		t.Errorf("was not able to build update to leads")
	}
}

func TestBuildLeadCreate(t *testing.T) {
	if closeioLead := BuildLeadCreate("brandon@ticketbud.com", "Unbounce"); len(closeioLead.Custom) == 0 {
		t.Errorf("was not able to build a lead to create in close")
	}
}
