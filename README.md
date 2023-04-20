# Overview

This repo was mainly created to standardize commit messages for the cloud content delivery team. It adds a pre commit hook that checks the structure of the commit. It can also be run in a cicd pipeline using the same regex validator, so there is no reason to manage the pipeline validation and local validation regex separately. It is currently being used in ccd's cicd pipelines.
