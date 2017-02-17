import bson
import codecs
import pymongo

"""
Two movies (in bookings and showtimes lists) have broken ids
  57b37bf0377dd100054f9f92 should be 57b37ba8377dd100054f9f92
  57b37bf0377dd100054f9f91 should be 57b37b7c377dd100054f9f91
"""

def update_showtimes(client):
    db = client.showtimes
    for item in db.showtimes.find():
        print item
        moviesJson = str(item[u'movies'])
        if '57b37bf0377dd100054f9f92' in moviesJson:
            print "To be updated: ", moviesJson
            id = item[u'_id']
            print id

            dict = eval(moviesJson.replace('57b37bf0377dd100054f9f92', '57b37ba8377dd100054f9f92'))
            print  "Ready for update: ", dict

            db.showtimes.update(
                { "_id": id },
                {
                    "movies": dict,
                    "date": item[u'date'],
                    "createdon": item[u'createdon']
                }
            )

def update_bookings(client):
    db = client.bookings
    for item in db.bookings.find():
        print item
        moviesJson = str(item[u'movies'])
        if '57b37bf0377dd100054f9f91' in moviesJson:
            print "To be updated: ", moviesJson
            id = item[u'_id']
            print id

            dict = eval(moviesJson.replace('57b37bf0377dd100054f9f91', '57b37b7c377dd100054f9f91'))
            print  "Ready for update: ", dict

            db.bookings.update(
                { "_id": id },
                {
                    "movies": dict,
                    "userid": item[u'userid'],
                    "showtimeid": item[u'showtimeid']
                }
            )


client = pymongo.MongoClient("localhost", 27017)
update_showtimes(client)
