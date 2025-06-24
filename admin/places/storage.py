from storages.backends.s3boto3 import S3Boto3Storage
from django.conf import settings

class CustomS3Boto3Storage(S3Boto3Storage):
    def url(self, name, parameters=None, expire=None):
        original_url = super().url(name, parameters=parameters, expire=expire)
        return original_url.replace("minio:9000", settings.HOST)