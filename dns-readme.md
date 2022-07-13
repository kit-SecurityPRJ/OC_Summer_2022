# dnsの構築説明

## 必要ファイル
- coredns
- hosts
- Corefile
- Dockerfile

### coredns

- corednsの生成
    - 利用時にはdocker が必要。ただしほかの方法もある。

``` 
$ docker run --rm -i -t -v $PWD:/v -w /v golang:1.17 make 
```

### hosts

- IPアドレスとドメイン名との対応表

### Corefile

- corednsの設定ファイル

### Dockerfile

- サーバー構築にDockerを利用したため,その設定ファイル。