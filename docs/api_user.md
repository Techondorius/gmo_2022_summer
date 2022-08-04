# API Document

すべてのリクエストのレスポンスには状態メッセージ("OK"、403など)、タイムスタンプ(ISO8601準拠)などが記述される。以下例。

```javascript
{
    {
        "ret_msg": "OK",
        "result": {
    ...
    }
        "time_now": "1567108756.834357"
    }
}
```


## ユーザー作成 API

### リクエスト

```
POST /api/regsiter
```

| param     | type   | description          |
|-----------|--------|----------------------|
| ID        | string | ID(固有, 最小2文字、最大20文字) |
| Name      | string | 表示名                  |
| Birthdate | int    | 誕生日                  |
| Sex       | int    | 性別(男、女、その他で1~3)      |
| Height    | int    | 身長                   |
| Weight    | int    | 体重                   |
| Password  | string | パスワード                |

```javascript
{
    "ID": "Pi",
    "Name": "ASDF",
    "Birthdate": 12341234,
    "Sex": 1,
    "Height": 169,
    "Weight": 55,
    "Password": "Raspberry"
}
```

### レスポンス

#### 成功時

| param            | type   | description |
|------------------|--------|-------------|
| Detail.ID        | string | ID          |
| Detail.Name      | string | 名前          |
| Detail.Birthdate | int    | 誕生日(UNIX)   |
| Detail.Sex       | int    | 性別(男、女で1/2) |
| Detail.Height    | int    | 身長          |
| Detail.Weight    | int    | 体重          |
| Detail.Objective | string | 目標消費カロリー    |

```javascript
{
    "detail": {
        "ID": "Pi",
        "Name": "ASDF",
        "Birthdate": 12341234,
        "Sex": 1,
        "Height": 169,
        "Weight": 55,
        "Objective": 100
    }
}
```
成功時のみCookieを設定する
目標消費カロリーは身長、体重、性別から算出する。


## ログイン API

### リクエスト

```
POST /api/login
```

| param    | type   | description |
| -------- | ------ | ------------|
| ID       | string | ID          |
| Password | string | パスワード     |

```javascript
{
    "ID": "Pi",
        "Password": "Raspberry"
}
```

CookieのIDとTokenを参照して、承認する。
承認した場合のみ200、承認していない場合403エラー

### レスポンス

#### 成功時

| param         | type    | description |
| ------------- | ------- | ----------- |
| Detail.Result | Boolean | trueなら承認  |

```javascript
{
    "Detail": {
        "Result": true
    }
}
```
成功時のみCookieを設定する

### 失敗時

#### Request bodyが不完全な時

400 Bad Request

#### Cookieでの承認が出来なかった場合

403 Forbidden


## ユーザー編集 API

### リクエスト

```
PUT /api/users/editUser
```

| param     | type   | description          |
|-----------|--------|----------------------|
| ID        | string | ID(固有)               |
| Name      | string | 表示名                  |
| Birthdate | int    | 誕生日(UNIX)            |
| Sex       | int    | 性別(男、女で1/2)          |
| Height    | int    | 身長                   |
| Weight    | int    | 体重                   |
| Objective | int    | 目標消費カロリー             |
| Password  | string | 旧パスワード(not required) |
| NPassword | string | 新パスワード(not required) |

```javascript
{
    "ID": "Pi",
        "Name": "ASDF",
        "Birthdate": 12341234,
        "Sex": 1,
        "Height": 169,
        "Weight": 55,
        "Objective": 100,
        "Password": "Raspberry",
        "NPassword": "R4spberry"
}
```
パスワードを変更する際にはPasswordとNPasswordは必要となる。

CookieのIDとTokenを参照して、承認する。
承認した場合のみ200、承認していない場合403エラー

### レスポンス

#### 成功時

| param            | type   | description |
|------------------|--------|-------------|
| Detail.ID        | string | ID          |
| Detail.Name      | string | 名前          |
| Detail.Birthdate | int    | 誕生日         |
| Detail.Sex       | int    | 性別(男、女で1/2) |
| Detail.Objective | string | 目標消費カロリー    |

```javascript
{
    "Message": "Created"
    "Detail": {
        "ID": "Pi",
            "Name": "ASDF",
            "Birthdate": 12341234
        "Sex": 1
        "Objective": 100,
    }
}
```

### 失敗時

#### Request bodyが不完全な時

400 Bad Request

#### Cookieでの承認が出来なかった場合

403 Forbidden


## ユーザーデータ取得 API

### リクエスト

```
GET /api/users/getUser
```

| param     | type | description |
|-----------|------|-------------|
| ID        | int  | 検索するID      |
| StartTime | int  | 開始時刻(UNIX)  |
| EndTime   | int  | 終了時刻(UNIX)  |

```javascript
{
    "ID": 1,
        "StartTime": 1656601200,
        "EndTime": 1659538800,
}
```
CookieのIDとTokenを参照して、承認する。
承認した場合のみ200、承認していない場合403エラー

### レスポンス

#### 成功時

| param              | type   | description |
|--------------------|--------|-------------|
| Detail.ID          | string | ID          |
| Detail.Name        | string | 名前          |
| Detail.Birthdate   | int    | 誕生日(UNIX)   |
| Detail.Sex         | int    | 性別(男、女で1/2) |
| Detail.Height      | int    | 身長          |
| Detail.Weight      | int    | 体重          |
| Detail.ConsumptedC | int    | 日間目標消費カロリー  |

```javascript
{
    "Detail": {
        "ID": "Pi",
        "Name": "ASDF",
        "Birthdate": 12341234,
        "Sex": 1,
        "Height": 169,
        "Weight": 55,
        "Objective": 100
    }
}
```

### 失敗時

#### Request bodyが不完全な時

400 Bad Request

#### Cookieでの承認が出来なかった場合

403 Forbidden