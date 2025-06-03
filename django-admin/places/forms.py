from django import forms
from .models import Photo
from uuid import uuid4
from django.core.files.storage import default_storage


class PhotoUploadForm(forms.ModelForm):
    upload = forms.FileField(required=False)

    class Meta:
        model = Photo
        fields = ['place', 'description', 'upload']

    def save(self, commit=True):
        instance = super().save(commit=False)
        file = self.cleaned_data.get('upload')
        if file:
            filename = default_storage.save(file.name, file)
            instance.image = filename  # сохраняем путь в поле image
        if commit:
            instance.save()
        return instance