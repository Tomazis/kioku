# Generated by Django 3.2.3 on 2021-06-10 17:45

from django.db import migrations, models
import django.db.models.deletion


class Migration(migrations.Migration):

    dependencies = [
        ('kanji', '0002_auto_20210610_1930'),
    ]

    operations = [
        migrations.AlterField(
            model_name='kanjialternative',
            name='kanji',
            field=models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, related_name='alternative', to='kanji.kanji'),
        ),
        migrations.CreateModel(
            name='Onyomi',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('name', models.CharField(max_length=10, verbose_name='Chinese reading')),
                ('kanji', models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, related_name='onyomi', to='kanji.kanji')),
            ],
        ),
        migrations.CreateModel(
            name='Kunyomi',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('name', models.CharField(max_length=10, verbose_name='Japanese reading')),
                ('kanji', models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, related_name='kunyomi', to='kanji.kanji')),
            ],
        ),
    ]
