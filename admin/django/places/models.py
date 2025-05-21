from django.contrib.gis.db import models

class City(models.Model):
    name = models.TextField(primary_key=True)

    class Meta:
        managed = False
        db_table = 'city'

    def __str__(self):
        return self.name

class Place(models.Model):
    city_name = models.ForeignKey(City, models.DO_NOTHING, db_column='city_name', blank=True, null=True)
    name = models.TextField(blank=True, null=True)
    geom = models.PointField(geography=True, blank=True, null=True)
    descr = models.TextField(blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'place'

    def __str__(self):
        return self.name





