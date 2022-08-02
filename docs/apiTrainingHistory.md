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


## トレーニング履歴一覧 API

### リクエスト

```
GET /api/training/
```
CookieのIDとTokenを参照して、承認する。
承認した場合のみ200、承認していない場合403エラー

### レスポンス

#### 成功時

200 OK

| param                 | type     | description    |
| --------------------- | -------- | -------------- |
| Detail[].ID           | int      | トレーニング履歴ID |
| Detail[].Time         | datetime | TR終了時刻       |
| Detail[].TName        | string   | トレーニング名     |
| Detail[].TLength      | int      | TR時間          |
| Detail[].ConsumptingC | int      | 消費カロリー       |

```javascript
{
    "Detail": [
        {
            "ID": 1004
            "Time": 2022-01-01 13:59:59.12311,
            "TName": "スクワット",
            "TLength": 120,
            "ConsumptingC": 1500
        },{
            ...
        }
    ]
}
```

### 失敗時

#### Request bodyが不完全な時

400 Bad Request

#### Cookieでの承認が出来なかった場合

403 Forbidden


## トレーニング履歴追加 API

### リクエスト

```
POST /api/training/add
```

| column     | type    | description              |
| ---------- | ------- | ------------------------ |
| ID         | int     | トレーニングID              |
| IsCustomed | boolean | 真なら自作TR               |
| TLength    | int     | TR時間(初期TRの場合のみ必要) |

```javascript
{
    "ID": 107,
    "IsCustom": false,
    "TLength": 120
}
```

```javascript
// 例2
{
    "ID": 216,
    "IsCustomed": true
}
```

### レスポンス

### 成功時

トレーニング履歴一覧と同じ

### 失敗時

#### Request bodyが不完全な時

400 Bad Request

#### Cookieでの承認が出来なかった場合

403 Forbidden


## トレーニング履歴編集 API

### リクエスト

```
PUT /api/training/edit
```

| column      | type    | description              |
| ----------- | ------- | ------------------------ |
| TRHistoryID | int     | トレーニング履歴ID           |
| ID          | int     | トレーニングID              |
| IsCustomed  | boolean | 真なら自作TR               |
| TLength     | int     | TR時間(初期TRの場合のみ必要) |

```javascript
{
    "TRHistoryID": 1007,
    "ID": 107,
    "IsCustom": false,
    "TLength": 120
}
```

```javascript
// 例2
{
    "TRHistoryID": 1008,
    "ID": 216,
    "IsCustomed": true
}
```

### レスポンス

### 成功時

トレーニング履歴一覧と同じ

### 失敗時

#### Request bodyが不完全な時

400 Bad Request

#### Cookieでの承認が出来なかった場合

403 Forbidden


## トレーニング履歴削除 API

### リクエスト

```
DELETE /api/training/delete
```

| column      | type    | description              |
| ----------- | ------- | ------------------------ |
| TRHistoryID | int     | トレーニング履歴ID           |

```javascript
{
    "TRHistoryID": 1007
}
```

### レスポンス

### 成功時

トレーニング履歴一覧と同じ

### 失敗時

#### Request bodyが不完全な時

400 Bad Request

#### Cookieでの承認が出来なかった場合

403 Forbidden

