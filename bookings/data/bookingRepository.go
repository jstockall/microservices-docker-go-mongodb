package data

import (
	"github.com/mmorejon/cinema/bookings/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type BookingRepository struct {
	C *mgo.Collection
}

func (r *BookingRepository) Create(booking *models.Booking) error {
	obj_id := bson.NewObjectId()
	booking.Id = obj_id
	err := r.C.Insert(&booking)
	return err
}

func (r *BookingRepository) GetAll() []models.Booking {
	var bookings []models.Booking
	iter := r.C.Find(nil).Iter()
	result := models.Booking{}
	for iter.Next(&result) {
		bookings = append(bookings, result)
	}
	return bookings
}

func (r *BookingRepository) GetBy(key string, value string) []models.Booking {
	var bookings []models.Booking
	iter := r.C.Find(bson.M{ key:value}).Iter()
	result := models.Booking{}
	for iter.Next(&result) {
		bookings = append(bookings, result)
	}
	return bookings
}

func (r *BookingRepository) Delete(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

func (r *BookingRepository) GetById(id string) (booking models.Booking, err error) {
	err = r.C.FindId(bson.ObjectIdHex(id)).One(&booking)
	return
}
