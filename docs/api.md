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
| --------- | ------ | -------------------- |
| ID        | string | ID(固有)              |
| Name      | string | 表示名                |
| Birthdate | date   | 誕生日                |
| Sex       | int    | 性別(男、女、その他で1~3) |
| Height    | int    | 身長                  |
| Weight    | int    | 体重                  |
| Password  | string | パスワード              |

```javascript
{
    "ID": "Pi",
    "Name": "ASDF",
    "Birthdate": 2002-1-1,
    "Sex": 1,
    "Height": 169,
    "Weight": 55,
    "Password": "Raspberry"
}
```

### レスポンス

#### 成功時

| param              | type   | description          |
| ------------------ | ------ | -------------------- |
| detail[].ID        | string | ID                   |
| detail[].Name      | string | 名前                  |
| detail[].Birthdate | string | 誕生日                |
| detail[].Objective | string | 目標消費カロリー         |
| detail[].Sex       | int    | 性別(男、女、その他で1~3) |

```javascript
{
    "detail": {
        "ID": "Pi",
        "Name": "ASDF",
        "Birthdate": 2002-1-1
        "Objective": 100,
        "Sex": 1
    }
}
```

目標消費カロリーは身長、体重、性別から算出する。


## ID重複チェック API

ユーザーIDは他のユーザーと重複できないため、既に同じIDのユーザーが居るか確認するためのエンドポイント

### リクエスト

```
GET /api/{id}
```

### レスポンス

#### 成功時

| param         | type    | description  |
| ------------- | ------- | ------------ |
| Detail.ID     | string  | ID           |
| Detail.Result | boolean | Trueで重複なし |

```javascript
{
    "Detail": {
        "ID": "Pia",
        "Result": "True"
    }
}
```

### 失敗時

#### Request bodyが不完全な時

400 Bad Request


## ユーザー編集 API

### リクエスト

```
PUT /api/register
```

| param     | type   | description          |
| --------- | ------ | -------------------- |
| ID        | string | ID(固有)              |
| Name      | string | 表示名                |
| Birthdate | date   | 誕生日                |
| Sex       | int    | 性別(男、女、その他で1~3) |
| Objective | int    | 目標消費カロリー         |
| Password  | string | パスワード              |

```javascript
{
    "ID": "Pi",
    "Name": "ASDF",
    "Birthdate": 2002-1-2
    "Sex": 1
    "Objective": 100,
    "Password": "Raspberry"
}
```

CookieのIDとTokenを参照して、承認する。
承認した場合のみ200、承認していない場合403エラー

### レスポンス

#### 成功時

| param              | type   | description          |
| ------------------ | ------ | -------------------- |
| Detail.ID        | string | ID                   |
| Detail.Name      | string | 名前                  |
| Detail.Birthdate | string | 誕生日                |
| Detail.Sex       | int    | 性別(男、女、その他で1~3) |
| Detail.Objective | string | 目標消費カロリー         |

```javascript
{
    "Message": "Created"
    "Detail": {
        "ID": "Pi",
        "Name": "ASDF",
        "Birthdate": 2002-1-2
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
GET /api/getUser
```

| param | type | description      |
| ----- | ---- | ---------------- |
| ID[]  | int  | 検索するID(最大10個)|

```javascript
{
    "ID": [
        1, 2, 3, ...
    ]
}
```
CookieのIDとTokenを参照して、承認する。
承認した場合のみ200、承認していない場合403エラー

### レスポンス

#### 成功時

| param                | type   | description          |
| -------------------- | ------ | -------------------- |
| Detail[].ID          | string | ID                   |
| Detail[].Name        | string | 名前                  |
| Detail[].Birthdate   | string | 誕生日                |
| Detail[].Sex         | int    | 性別(男、女、その他で1~3) |
| Detail[].ConsumptedC | int    | 日間消費カロリー         |

```javascript
{
    "Detail": [
        {
            "ID": 1,
            "Name": "Pi",
            "Birthdate": 2002-1-1
            "Sex": 1,
            "ConsumptedC": 500
        },{

        }
    ]
}
```

### 失敗時

#### Request bodyが不完全な時

400 Bad Request
