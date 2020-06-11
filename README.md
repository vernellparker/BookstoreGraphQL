# BookstoreGraphQL

using a POST to "/login" in an api client like postman will get you a token. You must use 

{
	"username":"admin",
	"password":"admin"
}
 or 
 {
	"username":"test",
	"password":"test"
}

to get a token. 

The new query link will be: http://localhost:8080/auth/query

Add the token to the HTTP Hearders section like the following:

{
  "Authorization": "Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTE4NTA4MTMsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTU5MTg0NzIxM30.uBu6KegttF2xCtwQYfRE2fvtZEg7Qnd5dRu_nehu9veyEm23pdyi9R2p6p0M426R2GOqPoi1MKzq_ngNi8Ze9w"
}
