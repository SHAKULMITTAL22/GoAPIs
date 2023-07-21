// Test generated by RoostGPT for test roost-test using AI Type Vertex AI and AI Model code-bison


    func TestGetBook_SuccessfulResponse(t *testing.T) {
        // Arrange
        req, err := http.NewRequest("GET", "/books/1234567890", nil)
        assert.NoError(t, err)
        
        // Act
        resp := &http.Response{}
        getBook(resp, req)
        
        // Assert
        assert.Equal(t, resp.StatusCode, http.StatusOK)
        assert.NotNil(t, resp.Body)
        var book Book
        err = json.NewDecoder(resp.Body).Decode(&book)
        assert.NoError(t, err)
        assert.Equal(t, book.ID, int64(1234567890))
    }
    
    func TestGetBook_InvalidParam(t *testing.T) {
        // Arrange
        req, err := http.NewRequest("GET", "/books/-1234567890", nil)
        assert.NoError(t, err)
        
        // Act
        resp := &http.Response{}
        getBook(resp, req)
        
        // Assert
        assert.Equal(t, resp.StatusCode, http.StatusBadRequest)
        assert.Empty(t, resp.Body)
        t.Logf("Received error response: %v", resp)
    }
    
    func TestGetBook_NonExistentBook(t *testing.T) {
        // Arrange
        req, err := http.NewRequest("GET", "/books/9999999999", nil)
        assert.NoError(t, err)
        
        // Act
        resp := &http.Response{}
        getBook(resp, req)
        
        // Assert
        assert.Equal(t, resp.StatusCode, http.StatusNotFound)
        assert.Empty(t, resp.Body)
        t.Logf("Received error response: %v", resp)
    }
    
    func TestGetBook_InternalServerError(t *testing.T) {
        // Arrange
        books = []Book{{ID: 1234567890}, {ID: 2345678901}}
        req, err := http.NewRequest("GET", "/books/2345678901", nil)
        assert.NoError(t, err)
        
        // Act
        resp := &http.Response{}
        getBook(resp, req)
        
        // Assert
        assert.Equal(t, resp.StatusCode, http.StatusInternalServerError)
        assert.Empty(t, resp.Body)
        t.Logf("Received error response: %v", resp)
    }
    
    func TestGetBook_JSONEncodingError(t *testing.T) {
        // Arrange
        books = []Book{{ID: 1234567890}, {ID: 2345678901}}
        req, err := http.NewRequest("GET", "/books/2345678901", nil)
        assert.NoError(t, err)
        
        // Act
        resp := &http.Response{}
        getBook(resp, req)
        
        // Assert
        assert.Equal(t, resp.StatusCode, http.StatusInternalServerError)
        assert.Empty(t, resp.Body)
        t.Logf("Received error response: %v", resp)
    }