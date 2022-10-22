| Endpoint           | Relative Path           | Method | Description                                                                     |
| ------------------ | ----------------------- | ------ | ------------------------------------------------------------------------------- |
| Create Event       | _/event/create_         | POST   | Endpoint to create new event                                                    |
| Create Location    | _/location/create_      | POST   | Endpoint to create new location                                                 |
| Create Ticket      | _/event/ticket/create_  | POST   | Endpoint to create new ticket type on one specific event                        |
| Get Event          | _/event/get_info_       | GET    | Endpoint to retrieve event information, including location data and ticket data |
| Purchase Ticket    | _/transaction/purchase_ | POST   | Endpoint to make a new purchase, customer data is sent via this API             |
| Transcation Detail | _/transaction/get_info_ | GET    | Endpoint to retrieve transaction created using endpoint _Purchase Ticket_       |
