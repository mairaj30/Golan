func TestCreateMessageHandler(t *testing.T) {
	// Initialize the Gin router
	r := gin.Default()

	// Define the expected response
	expectedResponse := &Message{
		Id:        1,
		Title:     "test title",
		Body:      "test body",
		CreatedAt: time.Now(),
	}

	// Define the mock repository
	mockRepo := &MessageRepoMock{
		CreateFunc: func(msg *Message) (*Message, error) {
			return expectedResponse, nil
		},
	}

	// Set the mock repository as the repository for the handler
	r.POST("/messages", CreateMessageHandler(mockRepo))

	// Create a new request with the desired JSON body
	jsonBody := `{"title": "test title", "body": "test body"}`
	req, _ := http.NewRequest(http.MethodPost, "/messages", bytes.NewBufferString(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	// Create a new response recorder
	w := httptest.NewRecorder()

	// Call the handler function with the request and response recorder
	r.ServeHTTP(w, req)

	// Assert that the response has the correct status code and body
	assert.Equal(t, http.StatusOK, w.Code)
	var actualResponse Message
	err := json.Unmarshal(w.Body.Bytes(), &actualResponse)
	assert.Nil(t, err)
	assert.Equal(t, expectedResponse, &actualResponse)
}
