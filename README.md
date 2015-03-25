###Lightweight struct validator
I built this for personal use to validate struct. it only validates string and slice field.

adding validator:"require" to your struct property like the example below.

go get https://github.com/apisit/go-json-validator

  example.

    person := struct {
  		FirstName string `json:"first_name,omitempty" validator:"require"`
  		LastName  string `json:"last_name,omitempty" validator:"require"`
  	}{}
  	person.FirstName = "Apisit"
  	validateErr := validator.Validate(person)
  	if validateErr != nil {
  		log.Printf("validation error: %v", validateErr)
  		return
  	}
  	
or check these Go playground 
https://play.golang.org/p/ueU4VwrF-G
https://play.golang.org/p/T74unGIhvp
