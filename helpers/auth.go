package helpers

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"code.google.com/p/go.crypto/bcrypt"
	"github.com/elcct/amazingwebsite/models"
)

func Login(c *mgo.Database, email string, password string) (user *models.User, err error) {
	err = c.C("users").Find(bson.M{"e": email}).One(&user)
	if err != nil {
		return
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	if err != nil {
		user = nil	
	}
	return
}
