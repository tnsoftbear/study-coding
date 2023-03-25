curl -H 'api-key:<api-key>' \
-X POST -d '{
"name":"Campaign sent via the API",
"subject":"My subject",
"sender": { "name": "From name", "email":"fake@gmail.com" },
"type": "classic",
# Content that will be sent
"htmlContent": "Congratulations! You successfully sent this example campaign via the Sendinblue API.",
# Select the recipients
"recipients": { "listIds": [2,7] },
# Schedule the sending in one hour
"scheduledAt": "2023-03-26 00:00:01",
}' \
'https://api.sendinblue.com/v3/emailCampaigns'
