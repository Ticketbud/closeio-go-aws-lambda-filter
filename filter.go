package filter

import (
	"fmt"
	"time"
	"os"
	"encoding/json"
	"io/ioutil"
	"github.com/Ticketbud/closeio-go"
)

type Email struct {
	Email string
}

type LeadUpdate struct {
	id string
	updateFields map[string]string
}

type Config struct {
	CloseKey string `json:"closekey"`
}
// Search for a given lead in closeio based on the incoming key, value query
func Search(client *closeio.Closeio, key, value string) (l *closeio.Leads) {
	
	search := &closeio.LeadSearch{
		Query: key + ":" + value,
	}
	
	leads, err := client.Leads(search)

	if err != nil {
		fmt.Println(err)
	}

	return leads
}

// Update returned leads from a search for those leads in our system
func BuildCustomLeadSourceUpdates(leads *closeio.Leads, leadSource string) (leadUpdates []LeadUpdate ) {
	for _, leadData := range leads.Data {

		if len(leadData.Custom["Lead_Source"]) == 0 {
			leadUpdates = append(leadUpdates, LeadUpdate{
				id: leadData.Id,
				updateFields: BuildCustomLeadSource(leadSource),
			})
		}

	}
	return leadUpdates
}

func BuildCustomLeadSource(source string) (l map[string]string) {
	return map[string]string {
		"custom.lcf_nvgL6lPLr1S1ExzrihHZXhpsCt27NiUTSUPYS1wf5Of" : source,
		"custom.lcf_PaiGpdkvtdMR1Xe6c7BoZEPeO3QTisMUXEIH7xHtGEO" : time.Now().Format(time.RFC3339),
	}
}

func BuildLeadCreate(email, source string) (l closeio.Lead) {
	return closeio.Lead {
		Contacts: &[]closeio.Contact {
			closeio.Contact{
				Emails: &[]closeio.Email{
					closeio.Email{
						Type: "main",
						Email: email,
					},
				},
			},
		},
		Custom: BuildCustomLeadSource(source),		
	}
}

func ReadConfig(file string) (config Config) {
	raw, err := ioutil.ReadFile(file)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	
	json.Unmarshal(raw, &config)

	return config
}

func main() {

	config := ReadConfig("./config.json")

	client := closeio.New(config.CloseKey)
	var email Email

	err := json.Unmarshal([]byte(os.Args[1:][0]), &email)

	if err != nil {
		fmt.Println(err)
	}
	
	emailLeads := Search(client, "email", email.Email)
	
	if len(emailLeads.Data) > 0 {
		for _, leadUpdate  := range BuildCustomLeadSourceUpdates(emailLeads, "Unbounce") {
			client.UpdateLead(leadUpdate.id, &leadUpdate.updateFields)
		}

	} else {
		newLead := BuildLeadCreate(email.Email, "Unbounce")
		client.CreateLead(&newLead)
	}
}



