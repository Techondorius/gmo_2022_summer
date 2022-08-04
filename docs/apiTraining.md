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


## トレーニング追加 API

### リクエスト

```
POST /api/customeTR/add
```

| param        | type   | description |
| ------------ | ------ | ----------- |
| Name         | string | トレーニング名  |
| ConsumptingC | int    | 消費カロリー    |

CookieのIDとTokenを参照して、承認する。
承認した場合のみ200、承認していない場合403エラー

### レスポンス

#### 成功時

トレーニングリスト追加と同じ

#### 失敗時

##### Request bodyが不完全な時

400 Bad Request

##### Cookieでの承認が出来なかった場合

403 Forbidden


## トレーニング一覧 API

自作トレーニングとデフォルトトレーニングを両方まとめてインポート

### リクエスト

```
GET /api/customeTR/
```

CookieのIDとTokenを参照して、承認する。
承認した場合のみ200、承認していない場合403エラー

### レスポンス

#### 成功時

| column       | type    | details                       |
| ------------ | ------- | ----------------------------- |
| Detail[].ID           | int     | トレーニングID(固有)              |
| Detail[].Name         | string  | トレーニング名                    |
| Detail[].UserTR       | boolean | ユーザー自作TRならtrue            |
| Detail[].ConsumptingC | int     | 自作TRなら(kcal)、初期TRならメッツ値 |

```json
{
    "Detail": [
        {
            "ID": 107
            "Name": "スクワット",
            "UserTR": false
            "ConsumptingC": 1.5
        },{
            "ID": 216
            "Name": "平泳ぎ100m",
            "UserTR": True
            "ConsumptingC": 1500
        },
        ...
    ]
}

```

IDは完全固有ではなく、初期TRとUserTRそれぞれで一意に定まるものであり、初期TRのIDとUserTRのIDが重複する可能性はある。

#### 失敗時

##### Request bodyが不完全な時

400 Bad Request

##### Cookieでの承認が出来なかった場合

403 Forbidden


## トレーニング削除 API

### リクエスト

```
DELETE /api/customeTR/delete
```

| column | type | description     |
| ------ | ---- | --------------- |
| ID     | int  | 削除する自作TRのID |

```javascript
{
    "ID": 216
}
```

### レスポンス

#### 成功時

トレーニング一覧APIと同じ

##### Request bodyが不完全な時

400 Bad Request

##### Cookieでの承認が出来なかった場合

403 Forbidden