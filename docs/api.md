# API Document

## ユーザー作成 API

### リクエスト

```
POST /api/regsiter
```

| param      | type   | description        |
| ----------- | ------ | ----------------- |
| メアド      | string | メアド              |
| ニックネーム | string | ページ内に表示する名前 |
| パスワード   | string | パスワード           |


### レスポンス

#### 成功時

| param                 | type   | description |
| --------------------- | ------ | ----------- |
| userid | number | ユーザーID    |

```javascript
{
    "detail": [
        {
            "ID": number,
            "FirstName": string,
            "LastName": string,
            "Hourly_wage": number
        },
        ...
    ]
}
```

## ユーザー作成 API

### リクエスト

```
PUT /userManage/newUser
```

| param       | type   | description   |
| ----------- | ------ | -----------   |
| FirstName   | string | 名前           |
| LastName    | string | 苗字           |
| Hourly_wage | number | 時給(Optional) |


```javascript
{
    "FirstName": string,
    "LastName": string
    "Hourly_wage": number
}
```

### レスポンス

#### 成功時

| param               | type   | description |
| ------------------- | ------ | ----------- |
| message             | string | Created     |
| detail[].ID         | number | ID          |
| detail[].FirstName  | string | 名前         |
| detail[].LastName   | string | 苗字         |
| detail[].Hourly_wage| string | 時給         |


```javascript
{
    "message": "Created"
    "detail": {
        "ID": number,
        "FirstName": string,
        "LastName": string
        "Hourly_wage": number,
    },
    ...
}
```

### 失敗時

#### Request bodyが不完全な時

400 Bad Request

## ユーザー編集 API

### リクエスト

```
PUT /userManage/editUser/{userID}
```

| param       | type   | description   |
| ----------- | ------ | ------------- |
| FirstName   | string | 名前(optional) |
| LastName    | string | 苗字(optional) |
| Hourly_wage | number | 時給(optional) |

```javascript
{
    "FirstName": string,
    "LastName": string
    "Hourly_wage": number,
}
```

### レスポンス

#### 成功時

| param               | type   | description |
| ------------------- | ------ | ----------- |
| message             | string | Created     |
| detail[].ID         | number | ID          |
| detail[].FirstName  | string | 名前         |
| detail[].LastName   | string | 苗字         |
| detail[].Hourly_wage| string | 時給         |

```javascript
{
    "message": "Created"
    "detail": {
        "ID": number,
        "FirstName": string,
        "LastName": string
        "Hourly_wage": number,
    },
    ...
}
```

### 失敗時

#### Request bodyが不完全な時

400 Bad Request

## ユーザー編集 API

### リクエスト

```
DELETE /userManage/deleteUser/{userID}
```

### レスポンス

#### 成功時

| param               | type   | description |
| ------------------- | ------ | ----------- |
| message             | string | Deleted     |

```javascript
{
    "message": "Deleted"
}
```

### 失敗時

#### Request bodyが不完全な時

400 Bad Request

## ユーザー編集 API

### リクエスト

```
POST /stamp/{userID}/in
```

### レスポンス

#### 成功時

| param                | type     | description          |
| -------------------- | -------- | -------------------- |
| message              | string   | Stamped successfully |
| detail[].ID          | number   | スタンプID            |
| detail[].UsersID     | number   | ユーザーID            |
| detail[].In_datetime | datetime | 出勤時刻              |
| detail[].Hourly_wage | number   | 時給                  |

```javascript
{
    "message": "Stamped successfully"
    "detail": {
        "ID": 4,
        "UsersID": 4,
        "Hourly_wage": 1000,
        "In_datetime": "2022/06/30 06:36"
    }
}
```

### 失敗時

#### Request bodyが不完全な時

400 Bad Request