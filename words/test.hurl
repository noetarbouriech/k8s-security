# Test the /nouns endpoint
GET http://localhost:8080/nouns
HTTP/1.1 200
[Asserts]
header "Content-Type" == "text/plain"
body != ""  # Ensure the body is not empty

# Test the /verbs endpoint
GET http://localhost:8080/verbs
HTTP/1.1 200
[Asserts]
header "Content-Type" == "text/plain"
body != ""  # Ensure the body is not empty