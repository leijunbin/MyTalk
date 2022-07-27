namespace go user

struct BaseResp {
    1:i32 status_code
    2:string status_message
}

struct RegisterUserRequest {
    1:required string username
    2:required string password
    3:required string phonenumber
}

struct RegisterUserResponse {
    1:required BaseResp base_resp
}

struct LoginUserRequest {
    1:required string username_or_phonenumber
    2:required string password
}

struct LoginUserResponse {
    1:required BaseResp base_resp
}

service UserService {
    RegisterUserResponse RegisterUser(1:RegisterUserRequest req)
    LoginUserResponse LoginUser(1:LoginUserRequest req)
}