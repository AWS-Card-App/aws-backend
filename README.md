![image](https://github.com/AWS-Card-App/aws-backend/assets/73871477/f5d7fafc-d150-49bc-ab08-4c77e987cc69)

### API Docs
1. `GET /cards`
    - Query parameters
      - `name`: The name of the user to get the cards for
    - Returns: An array of cards objects with the same name as passed in query parameters
2. `PUT /card`
    - Query parameters
      - `name`: The name of the user to add the cards to
    - Body: Card content as raw text
    - Returns: ID of the new card added
3. `DELETE /card`
    - Query parameters
      - `name`: The name of the user the card belongs to
      - `id`: The ID of the card
4. `PATCH /card`
    - Query parameters
      - `name`: The name of the user the card belongs to
      - `id`: The id of the card
    - Body: New card content as raw text
