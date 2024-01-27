# Composable Interfaces

This is an experiment/example of how to create sub-interfaces while being able 
to compose data-types together. In this repo I've used networking as an example,
but this could be applied to _any_ sub-interface approach.

I've often had to use this pattern to enable support for multiple types of 
database where you want to split an interface based on the entity type.

For example, you might have `User` and `Post` entities. You might want
to have a `UserRepository` and a `PostRepository` that can be composed into a
`Repository` interface.

This is a small example of how to do that:

```go
// datastore/datastore.go
type Repository interface {
    Post() PostRepository
    User() UserRepository
}

type PostRepository interface {
    Find(id int) (*Post, error)
    Save(post *Post) error
}

type UserRepository interface {
    Find(id int) (*User, error)
    Save(user *User) error
}

type repository struct {
    post PostRepository
    user UserRepository
}

// datastore/mongo/mongo.go
var _ datastore.Repository = (*MongoRepository)(nil)

func New() *MongoRepository {
    return &MongoRepository{
        post: NewMongoPostRepository(),
        user: NewMongoUserRepository(),
    }
}

type MongoRepository struct {
	// ...
}

// datastore/mysql/mysql.go
var _ datastore.Repository = (*MongoRepository)(nil)

func New() *MongoRepository {
    return &MongoRepository{
        post: NewMySQLPostRepository(),
        user: NewMySQLUserRepository(),
    }
}

type MySQLRepository struct {
	// ...
}
```

The repo example also shows how to share common expected data with "TypeData"
structs. I've not really thought of a cleaner solution idea for this yet, but
I'd be happy to hear any ideas.

- [main.go](cmd/network/main.go) - The entry point showing how we can compose multiple implementations to the same interface.
- [network](pkg/network) - provides the top level "domain" for networks, interfaces and transmission types.
- [interface](pkg/interface) - interface provides two concrete implementations - physical and vlan.
- [transmission](pkg/transmission) - transmission provides two concrete implementation - bluetooth and wifi.

Which ends up looking like this:

```go
func main() {}
    nic := network.New(
        bluetooth.New(bluetooth.BLEEnabled), // transmission method (TransmissionTypeable)
        physical.New("bnet0", "0xB10E70074"), // hardware interface (InterfaceTypeable)
    )
    
    // ...
}
```

The naming conventions I used have come from Hashicorp, which I've found to 
make clear what is an interface, vs the concrete composable implementation of 
said interface.

- `XTypeable` = interface
- `XType` = concrete implementation

Where `X` is the name of the thing you're wanting to have multiple 
implementations of.