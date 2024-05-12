package main

//
//func genrateRequest(w http.ResponseWriter, r *http.Request, user database.User) {
//
//	// Create a new HTTP client
//	client := &http.Client{}
//	body := fmt.Sprintf(`{"feed_id": "%s", "user_id": %d}`, feedID, userID)
//
//	// Build the request for the second handler
//	req, err := http.NewRequest(
//		"GET",
//		"http://localhost:8080/v1/feed", nil,
//	)
//	if err != nil {
//		fmt.Fprintf(w, "Error creating request: %v", err)
//		return
//	}
//
//	req.Header.Set("Content-Type", "application/json")
//	req.Header.Set("Authorization", "")
//
//	// Send the request and handle the response
//	resp, err := client.Do(req)
//	if err != nil {
//		fmt.Fprintf(w, "Error sending request: %v", err)
//		return
//	}
//	defer resp.Body.Close()
//
//	// Process the response if needed (e.g., read body)
//}
