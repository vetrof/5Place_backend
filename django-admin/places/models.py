from django.contrib.gis.db import models
from django.utils.html import format_html
from .storage import CustomS3Boto3Storage
from django.conf import settings


class AppUser(models.Model):
    uuid = models.TextField()
    name = models.TextField()
    email = models.TextField()
    firebase_token = models.TextField()
    telegram_id = models.TextField()
    jwt_token = models.TextField()
    author = models.BooleanField(default=False)
    staff = models.BooleanField(default=False)
    superuser = models.BooleanField(default=False)
    created_at = models.DateTimeField(auto_now_add=True)

    class Meta:
        managed = False
        db_table = 'app_user'

    def __str__(self):
        return self.name

class PlaceType(models.Model):
    name = models.TextField(unique=True)

    class Meta:
        managed = False
        db_table = 'app_place_type'

    def __str__(self):
        return self.name


class Country(models.Model):
    name = models.TextField(unique=True)

    class Meta:
        managed = False
        db_table = 'app_country'

    def __str__(self):
        return self.name

class City(models.Model):
    name = models.TextField(unique=True)
    geom = models.PointField(geography=True)
    country = models.ForeignKey(Country, models.DO_NOTHING)

    class Meta:
        managed = False
        db_table = 'app_city'

    def __str__(self):
        return self.name


class Place(models.Model):
    type = models.ForeignKey('PlaceType', models.DO_NOTHING)
    city = models.ForeignKey(City, models.DO_NOTHING)
    name = models.CharField(max_length=255)
    descr = models.TextField()
    geom = models.PointField(geography=True)

    class Meta:
        managed = False
        db_table = 'app_place'

    def __str__(self):
       return self.name


class Photo(models.Model):
    place = models.ForeignKey('Place', models.DO_NOTHING, blank=True, null=True)
    image = models.ImageField(upload_to="places_photo/", storage=CustomS3Boto3Storage())
    description = models.CharField(max_length=255, blank=True, null=True)

    def image_tag(self):
        if self.image:
            url = self.image.url
            return format_html('<img src="{}" style="max-height: 200px;" />', url)
        return "—"

    image_tag.short_description = 'Preview'
    image_tag.allow_tags = True

    class Meta:
        managed = False
        db_table = 'app_photo'

    def __str__(self):
        return f"{self.place}"


class Favorite(models.Model):
    user = models.ForeignKey("AppUser", models.DO_NOTHING, blank=True, null=True)
    place = models.ForeignKey("Place", models.DO_NOTHING, blank=True, null=True)
    created_at = models.DateTimeField(auto_now_add=True)

    class Meta:
        managed = False
        db_table = 'app_favorite'


class VisitedPlace(models.Model):
    user_id = models.ForeignKey("AppUser", models.DO_NOTHING, blank=True, null=True)
    place_id = models.ForeignKey("Place", models.DO_NOTHING, blank=True, null=True)
    visited_at = models.DateTimeField(auto_now_add=True)

    class Meta:
        managed = False
        db_table = 'app_visited_place'



