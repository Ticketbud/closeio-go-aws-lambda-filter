package main

import (
  "encoding/json"
  "fmt"
  "os"
  "time"
  "./closeio-go"
)

type Email struct {
  Email string
}

func main() {
  client := closeio.New("XXXXXXXXXXXXXX")

  var email Email
  err := json.Unmarshal([]byte(os.Args[1:][0]), &email)

  search := &closeio.LeadSearch{
    Query: "email:" + email.Email,
  }

  leads, err := client.Leads(search)
  if err != nil {
    fmt.Println(err)
  }
  if len(leads.Data) > 0 {
    for _, leadData := range leads.Data {
      if len(leadData.Custom["Lead_Source"]) > 0 {
        lead := map[string]string {
          "custom.lcf_PaiGpdkvtdMR1Xe6c7BoZEPeO3QTisMUXEIH7xHtGEO": time.Now().Format(time.RFC3339),
        }
        client.UpdateLead(leadData.Id, &lead)
      } else {
        lead := map[string]string {
          "custom.lcf_nvgL6lPLr1S1ExzrihHZXhpsCt27NiUTSUPYS1wf5Of": "Unbounce",
          "custom.lcf_PaiGpdkvtdMR1Xe6c7BoZEPeO3QTisMUXEIH7xHtGEO": time.Now().Format(time.RFC3339),
        }
        client.UpdateLead(leadData.Id, &lead)
      }
    }
  } else {
    lead := closeio.Lead {
      Contacts: &[]closeio.Contact{
        closeio.Contact{
          Emails: &[]closeio.Email{
            closeio.Email{
              Type: "main",
              Email: email.Email,
            },
          },
        },
      }, // Contacts
      Custom: map[string]string {
        "lcf_PaiGpdkvtdMR1Xe6c7BoZEPeO3QTisMUXEIH7xHtGEO": time.Now().Format(time.RFC3339),
        "lcf_nvgL6lPLr1S1ExzrihHZXhpsCt27NiUTSUPYS1wf5Of": "Unbounce",
      },
    }

    created, err := client.CreateLead(&lead)
    if err != nil {
      fmt.Println(err)
    }
    fmt.Printf("Created!", created)
  }
}

