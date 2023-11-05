# 灵嘚用户统一认证服务（LingDei-AuthServer）

灵嘚用户认证服务接口是一个基于 Fiber 的 RESTful API 服务，用于提供灵嘚（LingDei）的用户认证服务。

## 配置

config/config.yaml

```yaml
# 数据库
mysql:
  DSN: 数据库连接字符串

QINIU:
  access_key: 七牛云access_key
  secret_key: 七牛云secret_key
  bucket: 七牛云存储空间名称
```

## 运行

```bash
go mod tidy
go run .
````

## 接口文档

### Security

**ApiKeyAuth**

| apiKey | _API Key_     |
| ------ | ------------- |
| In     | header        |
| Name   | Authorization |

---

### /profile/avatar

#### POST

##### Summary

上传头像

##### Description

上传头像

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| file | formData   | 头像文件    | Yes      | file   |

##### Responses

| Code | Description | Schema                                     |
| ---- | ----------- | ------------------------------------------ |
| 200  | OK          | [model.AvatorResp](#modelavatorresp)       |
| 400  | Bad Request | [model.OperationResp](#modeloperationresp) |

##### Security

| Security Schema | Scopes |
| --------------- | ------ |
| ApiKeyAuth      |        |

### /profile/delete

#### DELETE

##### Summary

删除用户 Profile🦸‍♂️

##### Description

删除用户 Profile，本接口对管理员权限的要求为等级 7。一般不需要用到本接口。

##### Parameters

| Name | Located in | Description              | Required | Schema |
| ---- | ---------- | ------------------------ | -------- | ------ |
| id   | query      | 被删除 Profile 的用户 id | Yes      | string |

##### Responses

| Code | Description | Schema                                     |
| ---- | ----------- | ------------------------------------------ |
| 200  | OK          | [model.OperationResp](#modeloperationresp) |
| 400  | Bad Request | [model.OperationResp](#modeloperationresp) |

##### Security

| Security Schema | Scopes |
| --------------- | ------ |
| ApiKeyAuth      |        |

### /profile/get

#### GET

##### Summary

获取用户 Profile

##### Description

获取用户 Profile

##### Parameters

| Name      | Located in | Description | Required | Schema |
| --------- | ---------- | ----------- | -------- | ------ |
| user_uuid | query      | 用户 UUID   | No       | string |

##### Responses

| Code | Description | Schema                                     |
| ---- | ----------- | ------------------------------------------ |
| 200  | OK          | [model.ProfileResp](#modelprofileresp)     |
| 400  | Bad Request | [model.OperationResp](#modeloperationresp) |

### /profile/my

#### GET

##### Summary

获取用户 Profile

##### Description

获取用户 Profile，用户 id 通过 token 取得，不需要单独传参

##### Parameters

| Name      | Located in | Description | Required | Schema |
| --------- | ---------- | ----------- | -------- | ------ |
| user_uuid | query      | 用户 UUID   | No       | string |

##### Responses

| Code | Description | Schema                                     |
| ---- | ----------- | ------------------------------------------ |
| 200  | OK          | [model.ProfileResp](#modelprofileresp)     |
| 400  | Bad Request | [model.OperationResp](#modeloperationresp) |

##### Security

| Security Schema | Scopes |
| --------------- | ------ |
| ApiKeyAuth      |        |

### /profile/update

#### POST

##### Summary

更新用户个人资料

##### Description

更新用户个人资料，用户 id 通过 token 取得，不需要单独传参。若 Profile 不存在则会自动创建新 Profile。若参数为空则代表不对该项进行修改。

##### Parameters

| Name      | Located in | Description | Required | Schema |
| --------- | ---------- | ----------- | -------- | ------ |
| nickname  | query      | 用户昵称    | No       | string |
| avatarurl | query      | 头像地址    | No       | string |
| email     | query      | 邮箱        | No       | string |

##### Responses

| Code | Description | Schema                                     |
| ---- | ----------- | ------------------------------------------ |
| 200  | OK          | [model.OperationResp](#modeloperationresp) |
| 400  | Bad Request | [model.OperationResp](#modeloperationresp) |

##### Security

| Security Schema | Scopes |
| --------------- | ------ |
| ApiKeyAuth      |        |

---

### /token/get

#### POST

##### Summary

获取 Token

##### Description

通过用户名和密码获取 token

##### Parameters

| Name     | Located in | Description | Required | Schema |
| -------- | ---------- | ----------- | -------- | ------ |
| username | query      | UserName    | Yes      | string |
| password | query      | Password    | Yes      | string |

##### Responses

| Code | Description | Schema                                     |
| ---- | ----------- | ------------------------------------------ |
| 200  | OK          | [model.TokenResp](#modeltokenresp)         |
| 400  | Bad Request | [model.OperationResp](#modeloperationresp) |

---

### /user/changepwd

#### POST

##### Summary

修改用户密码

##### Description

修改用户密码，username 通过 token 取得，只需传入新密码、旧密码两个参数即可

##### Parameters

| Name         | Located in | Description | Required | Schema |
| ------------ | ---------- | ----------- | -------- | ------ |
| old_password | query      | 旧密码      | Yes      | string |
| new_password | query      | 新密码      | Yes      | string |

##### Responses

| Code | Description | Schema                                     |
| ---- | ----------- | ------------------------------------------ |
| 200  | OK          | [model.OperationResp](#modeloperationresp) |
| 400  | Bad Request | [model.OperationResp](#modeloperationresp) |

##### Security

| Security Schema | Scopes |
| --------------- | ------ |
| ApiKeyAuth      |        |

### /user/delete

#### DELETE

##### Summary

删除用户 🦸‍♂️

##### Description

删除用户，本接口对管理员权限的要求为等级 7。仅能删除权限比自己低的管理员，普通用户可以被随意删除。

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| id   | query      | 用户 id     | Yes      | string |

##### Responses

| Code | Description | Schema                                     |
| ---- | ----------- | ------------------------------------------ |
| 200  | OK          | [model.OperationResp](#modeloperationresp) |
| 400  | Bad Request | [model.OperationResp](#modeloperationresp) |

##### Security

| Security Schema | Scopes |
| --------------- | ------ |
| ApiKeyAuth      |        |

### /user/exist

#### GET

##### Summary

检查用户名是否已经存在

##### Description

通过用户名检查用户名是否已经存在

##### Parameters

| Name     | Located in | Description | Required | Schema |
| -------- | ---------- | ----------- | -------- | ------ |
| username | query      | UserName    | Yes      | string |

##### Responses

| Code | Description | Schema                                     |
| ---- | ----------- | ------------------------------------------ |
| 200  | OK          | [model.OperationResp](#modeloperationresp) |
| 400  | Bad Request | [model.OperationResp](#modeloperationresp) |

##### Security

| Security Schema | Scopes |
| --------------- | ------ |
| ApiKeyAuth      |        |

### /user/getid

#### GET

##### Summary

获取用户 id

##### Description

通过用户名获取用户 id

##### Parameters

| Name     | Located in | Description | Required | Schema |
| -------- | ---------- | ----------- | -------- | ------ |
| username | query      | UserName    | Yes      | string |

##### Responses

| Code | Description | Schema                                     |
| ---- | ----------- | ------------------------------------------ |
| 200  | OK          | [model.UIDResp](#modeluidresp)             |
| 400  | Bad Request | [model.OperationResp](#modeloperationresp) |

##### Security

| Security Schema | Scopes |
| --------------- | ------ |
| ApiKeyAuth      |        |

### /user/getusers

#### GET

##### Summary

获取用户列表 🦸‍♂️

##### Description

获取用户列表，可选参数 username 的作用是搜索所有包括该参数的用户名，本接口对管理员权限的要求为等级 6

##### Parameters

| Name     | Located in | Description | Required | Schema |
| -------- | ---------- | ----------- | -------- | ------ |
| username | query      | UserName    | No       | string |

##### Responses

| Code | Description | Schema                                     |
| ---- | ----------- | ------------------------------------------ |
| 200  | OK          | [model.UsersResp](#modelusersresp)         |
| 400  | Bad Request | [model.OperationResp](#modeloperationresp) |

##### Security

| Security Schema | Scopes |
| --------------- | ------ |
| ApiKeyAuth      |        |

### /user/logout

#### DELETE

##### Summary

Token 注销

##### Description

用户登陆注销，本接口暂未实现实际功能，请在前端清除 token

##### Responses

| Code | Description | Schema                                     |
| ---- | ----------- | ------------------------------------------ |
| 200  | OK          | [model.OperationResp](#modeloperationresp) |
| 400  | Bad Request | [model.OperationResp](#modeloperationresp) |

##### Security

| Security Schema | Scopes |
| --------------- | ------ |
| ApiKeyAuth      |        |

### /user/register

#### POST

##### Summary

用户注册

##### Description

新用户注册

##### Parameters

| Name     | Located in | Description | Required | Schema |
| -------- | ---------- | ----------- | -------- | ------ |
| username | query      | UserName    | Yes      | string |
| password | query      | Password    | Yes      | string |

##### Responses

| Code | Description | Schema                                     |
| ---- | ----------- | ------------------------------------------ |
| 200  | OK          | [model.OperationResp](#modeloperationresp) |
| 400  | Bad Request | [model.OperationResp](#modeloperationresp) |

### /user/update

#### POST

##### Summary

修改任意用户 🦸‍♂️

##### Description

修改任意用户,管理员权限要求为 9

##### Parameters

| Name     | Located in | Description  | Required | Schema  |
| -------- | ---------- | ------------ | -------- | ------- |
| username | query      | 用户名       | Yes      | string  |
| password | query      | 新密码       | Yes      | string  |
| role     | query      | 权限等级 1~9 | Yes      | integer |

##### Responses

| Code | Description | Schema                                     |
| ---- | ----------- | ------------------------------------------ |
| 200  | OK          | [model.OperationResp](#modeloperationresp) |
| 400  | Bad Request | [model.OperationResp](#modeloperationresp) |

##### Security

| Security Schema | Scopes |
| --------------- | ------ |
| ApiKeyAuth      |        |

---

### Models

#### model.AvatorResp

| Name | Type    | Description | Required |
| ---- | ------- | ----------- | -------- |
| code | integer |             | No       |
| url  | string  |             | No       |

#### model.OperationResp

| Name | Type    | Description | Required |
| ---- | ------- | ----------- | -------- |
| code | integer |             | No       |
| msg  | string  |             | No       |

#### model.Profile

| Name       | Type   | Description | Required |
| ---------- | ------ | ----------- | -------- |
| avatar_url | string |             | No       |
| email      | string |             | No       |
| id         | string |             | Yes      |
| nickname   | string |             | Yes      |

#### model.ProfileResp

| Name    | Type                           | Description | Required |
| ------- | ------------------------------ | ----------- | -------- |
| code    | integer                        |             | No       |
| profile | [model.Profile](#modelprofile) |             | No       |

#### model.TokenResp

| Name  | Type    | Description | Required |
| ----- | ------- | ----------- | -------- |
| code  | integer |             | No       |
| role  | integer |             | No       |
| token | string  |             | No       |

#### model.UIDResp

| Name | Type    | Description | Required |
| ---- | ------- | ----------- | -------- |
| code | integer |             | No       |
| id   | string  |             | No       |

#### model.User

| Name     | Type                           | Description | Required |
| -------- | ------------------------------ | ----------- | -------- |
| id       | string                         |             | Yes      |
| profile  | [model.Profile](#modelprofile) | Profile     | No       |
| role     | integer                        |             | No       |
| username | string                         |             | Yes      |

#### model.UsersResp

| Name  | Type                         | Description | Required |
| ----- | ---------------------------- | ----------- | -------- |
| code  | integer                      |             | No       |
| users | [ [model.User](#modeluser) ] |             | No       |
