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


## フレンド追加 API

### リクエスト

```
POST /api/friend/follow
```

| param     | type | description |
| --------- | ---- | ----------- |
| Following | int  | フォロー先ID   |

```javascript
{
    "Following": 2
}
```
CookieのIDとTokenを参照して、承認する。
承認した場合のみ200、承認していない場合403エラー

### レスポンス

#### 成功時

200 OK

| param            | type | description |
| ---------------- | ---- | ----------- |
| Detail.Following | int  | フォロー先ID   |

```javascript
{
    "Detail": {
        "Following": 2
    }
}
```

### 失敗時

#### Request bodyが不完全な時

400 Bad Request

#### Cookieでの承認が出来なかった場合

403 Forbidden


## フレンド追加 API

### リクエスト

```
POST /api/friend/follow
```

| param       | type | description    |
| ----------- | ---- | -------------- |
| Unfollowing | int  | フォロー外す先のID |

```javascript
{
    "Unfollowing": 2
}
```
CookieのIDとTokenを参照して、承認する。
承認した場合のみ200、承認していない場合403エラー

### レスポンス

#### 成功時

200 OK

| param              | type | description    |
| ------------------ | ---- | -------------- |
| Detail.Unfollowing | int  | フォロー外す先のID |

```javascript
{
    "Detail": {
        "Unfollowing": 2
    }
}
```

### 失敗時

#### Request bodyが不完全な時

400 Bad Request

#### Cookieでの承認が出来なかった場合

403 Forbidden

