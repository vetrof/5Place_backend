from django.contrib import admin
from leaflet.admin import LeafletGeoAdmin

from places.models import Place, City

@admin.register(Place)
class PlaceAdmin(LeafletGeoAdmin):
    list_display = ("city_name", "name", "geom")

@admin.register(City)
class PlaceAdmin(admin.ModelAdmin):
    list_display = ("name",)
