# -*- coding: utf-8 -*-
from __future__ import absolute_import, division, print_function

from marshmallow import Schema, fields, post_load, post_dump
from marshmallow import validate

from polyaxon_schemas.base import BaseConfig
from polyaxon_schemas.experiment import ExperimentSchema
from polyaxon_schemas.utils import UUID


class ExperimentGroupSchema(Schema):
    uuid = UUID(allow_none=True)
    name = fields.Str()
    description = fields.Str(allow_none=True)
    content = fields.Str()
    project = UUID(allow_none=True)
    num_experiments = fields.Int(allow_none=True)
    experiments = fields.Nested(ExperimentSchema, many=True, allow_none=True)

    class Meta:
        ordered = True

    @post_load
    def make(self, data):
        return ExperimentGroupConfig(**data)

    @post_dump
    def unmake(self, data):
        return ExperimentGroupConfig.remove_reduced_attrs(data)


class ExperimentGroupConfig(BaseConfig):
    SCHEMA = ExperimentGroupSchema
    IDENTIFIER = 'experiment_group'
    REDUCED_ATTRIBUTES = ['uuid', 'project', 'experiments', 'description']

    def __init__(self,
                 name,
                 description=None,
                 content=None,
                 uuid=None,
                 project=None,
                 num_experiments=0,
                 experiments=None,):
        self.name = name
        self.description = description
        self.content = content
        self.uuid = uuid
        self.project = project
        self.num_experiments = num_experiments
        self.experiments = experiments


class ProjectSchema(Schema):
    name = fields.Str(validate=validate.Regexp(regex=r'^[-a-zA-Z0-9_]+\Z'))  # TODO: must be slug
    uuid = UUID(allow_none=True)
    description = fields.Str(allow_none=True)
    is_public = fields.Boolean(allow_none=True)
    num_experiments = fields.Int(allow_none=True)
    num_experiment_groups = fields.Int(allow_none=True)
    experiment_groups = fields.Nested(ExperimentGroupSchema, many=True, allow_none=True)
    experiments = fields.Nested(ExperimentSchema, many=True, allow_none=True)

    class Meta:
        ordered = True

    @post_load
    def make(self, data):
        return ProjectConfig(**data)

    @post_dump
    def unmake(self, data):
        return ProjectConfig.remove_reduced_attrs(data)


class ProjectConfig(BaseConfig):
    SCHEMA = ProjectSchema
    IDENTIFIER = 'project'
    REDUCED_ATTRIBUTES = ['uuid', 'description', 'experiments', 'experiment_groups']

    def __init__(self,
                 name,
                 uuid=None,
                 description=None,
                 is_public=True,
                 num_experiments=0,
                 num_experiment_groups=0,
                 experiments=None,
                 experiment_groups=None):
        self.name = name
        self.uuid = uuid
        self.description = description
        self.is_public = is_public
        self.num_experiments = num_experiments
        self.num_experiment_groups = num_experiment_groups
        self.experiments = experiments
        self.experiment_groups = experiment_groups
