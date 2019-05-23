# Polybucket

## Why
Polybucket is simple version control system for Machine Learning Model.

DVC(https://dvc.org/) is nice project, but it's many parts depend on Git.
For example, you create model of version 1. 
If you wanna create new version model on another branch, you have to checkout to another branch on Git.
I don't wanna it. I think that model's version controlling should be independent from Git and more version controll system.

So I developed this project.
Polybucket has backend S3 or GCS. You can select each cloud service or local file system.

## Data structure
