# Polybucket

## Why
Polybucket is simple version control system for Machine Learning Model.

DVC(https://dvc.org/) is nice project, but it's many parts depend on Git.
For example, you create model of version 1. 
If you wanna create new version model on another branch, you have to checkout to another branch on Git.
I don't wanna it. I think that model's version controlling should be independent from Git and more version controll system.

So I developed this project.
Polybucket has backend S3 or GCS. You can select each cloud service or local file system.

## Management Different
Polybucket managements model's different on Storage.
Management alogorithm is based on rsync.
It is not difficult.

I show the example on below.

![polybucket_repository_statement](https://user-images.githubusercontent.com/15085723/58231781-c10b5580-7d72-11e9-857f-bbb3c1a1c481.png)

This image indicates repository statement between each commit.
On first commit, you pushed a model file.
Then, you modify the model's parameter and push it.

On that time, Polybucket calculate diff between first model's file and second it.
A latest model file will be pushed completely, but first model's file transformed diff compared to the first.

If you restore the model file of first version, polybucket load a latest file and apply diff of first model.

Specifically, let's show simple case.

```
1: 0x00 0x00 0x00

1-2 diff = [0: 0x00, 3: None]

2: 0x01 0x00 0x00 0x01

2-3 diff = [1: 0x00, 3: 0x01]

3(latest): 0x01 0x01 0x00 0x00
```

You want to rollback the model version, Polybucket apply some paches on term of you want to return.