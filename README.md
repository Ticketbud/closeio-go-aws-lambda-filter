# Close.io Filter AWS Lambda Golang

Accepts a JSON string of an email address (from Zapier, most likely):

```
{ "email": "nathaniel@ticketbud.com" }
``` 

If the contact doesn't exist, we create it, and tag it with "Unbounce" and the currrent time.

If the contact does exist, we tag `Lead_Source_Unbounce` with the current time. If the contact doesn't already have a Lead Source, we go ahead and add Unbounce.

Would probably need a bit of hacking to get working with your setup, but hopefully serves as a good starting point.

# Testing

To run the tests, create a config.json file:

```
echo "{"closekey":"XXXXXXXXXXXXXXX"}" >> config.json
```

Then just:

```
go test
```
