# Generated by Django 2.0.3 on 2018-04-09 08:04

from django.conf import settings
from django.db import migrations, models
import django.db.models.deletion
import libs.spec_validation
import uuid


class Migration(migrations.Migration):

    initial = True

    dependencies = [
        migrations.swappable_dependency(settings.AUTH_USER_MODEL),
        ('projects', '0001_initial'),
        ('repos', '0001_initial'),
    ]

    operations = [
        migrations.CreateModel(
            name='ExperimentGroup',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('description', models.TextField(blank=True, null=True)),
                ('created_at', models.DateTimeField(auto_now_add=True, db_index=True)),
                ('updated_at', models.DateTimeField(auto_now=True)),
                ('uuid', models.UUIDField(default=uuid.uuid4, editable=False, unique=True)),
                ('sequence', models.PositiveSmallIntegerField(editable=False, help_text='The sequence number of this group within the project.')),
                ('content', models.TextField(help_text='The yaml content of the polyaxonfile/specification.', validators=[libs.spec_validation.validate_spec_content])),
                ('code_reference', models.ForeignKey(blank=True, null=True, on_delete=django.db.models.deletion.SET_NULL, related_name='groups', to='repos.CodeReference')),
                ('project', models.ForeignKey(help_text='The project this polyaxonfile belongs to.', on_delete=django.db.models.deletion.CASCADE, related_name='experiment_groups', to='projects.Project')),
                ('user', models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, related_name='experiment_groups', to=settings.AUTH_USER_MODEL)),
            ],
            options={
                'ordering': ['sequence'],
            },
        ),
        migrations.AlterUniqueTogether(
            name='experimentgroup',
            unique_together={('project', 'sequence')},
        ),
    ]
