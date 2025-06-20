from django.contrib import admin
from leaflet.admin import LeafletGeoAdmin
from places.models import Place, City, Photo, PlaceType, AppUser, Favorite, VisitedPlace, Country
from django.utils.html import format_html
from django.conf import settings


class PhotoInline(admin.TabularInline):
    model = Photo
    extra = 1  # сколько пустых форм показывать
    readonly_fields = ('image_tag',)
    fields = ('image', 'description', 'image_tag')  # что видно в инлайне

    def image_tag(self, obj):
        if obj.image:
            url = obj.image.url
            return format_html('<img src="{}" style="max-height: 100px;" />', url)
        return "—"

    image_tag.short_description = 'Preview'

admin.site.register(AppUser)

@admin.register(PlaceType)
class PlaceAdmin(LeafletGeoAdmin):
    list_display = ("name",)

@admin.register(Country)
class PlaceAdmin(LeafletGeoAdmin):
    list_display = ("name",)

@admin.register(City)
class PlaceAdmin(LeafletGeoAdmin):
    list_display = ("name",)

@admin.register(Place)
class PlaceAdmin(LeafletGeoAdmin):
    list_display = ("name", "geom")
    inlines = [PhotoInline]

@admin.register(Photo)
class PhotoAdmin(admin.ModelAdmin):
    list_display = ("place", "image_tag")
    readonly_fields = ("image_tag",)

admin.site.register(Favorite)

# admin.site.register(VisitedPlace)