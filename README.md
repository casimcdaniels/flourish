### Flourish API Coding Assignment

| Method        | Endpoint          | Description
| ------------- |:-------------:|:-----------:|
| POST      | /strains      |  Create a new strain { name: "", race: "", flavors: [], effects: [] } |
| GET      | /strains/123 |  Fetch a record in the strains table |
| PATCH | /strains/123      |  Updates a strain { name: "", race: "", flavors: [], effects: [] }  |
| DELETE      | /strains/123 |  Deletes the requested strain |
| GET     | /strains/search    |  Search and filter the strains e.g. strains/search?name=StrainA&treatment=Depression&effect=Sleepy&race=indica?flavor=Pine |
