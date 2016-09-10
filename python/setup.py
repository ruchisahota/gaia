# -*- coding: utf-8 -*-

from setuptools import setup

packages = ['gaia']
resources = []
api_version_path = "./gaia"

setup(
    name='gaia',
    version="1.0",
    url='',
    author='',
    author_email='',
    packages=packages,
    description='SDK for Aporeto REST API',
    include_package_data=True,
    install_requires=[line for line in open('requirements.txt')],
    dependency_links=[
      "git+https://github.com/aporeto-inc/pyelemental.git",
    ],
    data_files=resources
)