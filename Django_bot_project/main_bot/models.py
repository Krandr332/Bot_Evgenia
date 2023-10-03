from django.contrib import admin
from django.contrib.auth.models import User
from django.db import models


# Создаем модель расширения профиля пользователя
class UserProfile(models.Model):
    user = models.OneToOneField(User, on_delete=models.CASCADE)
    middle_name = models.CharField(max_length=255)
    registration_date = models.DateField()
    phone_number = models.CharField(max_length=15)
    access_rights = models.CharField(max_length=255)
    tg_id = models.CharField(max_length=255)
    region = models.CharField(max_length=255)


# Создаем модель для постов
class Post(models.Model):
    post_id = models.AutoField(primary_key=True)
    region = models.CharField(max_length=255)
    status = models.CharField(max_length=255)
    text = models.TextField()
    image = models.ImageField(upload_to='posts/', null=True, blank=True)
    date_added = models.DateTimeField(auto_now_add=True)
    date_published = models.DateTimeField()
    user = models.ForeignKey(User, on_delete=models.CASCADE)

# Регистрируем модели в административной панели Django
